package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"sort"
	"strings"
	"time"

	"github.com/Jeswyrne/chlnge/api/models"
	"github.com/Jeswyrne/chlnge/api/response"
	customPkg "github.com/Jeswyrne/chlnge/pkg/user"
	"github.com/patrickmn/go-cache"
)

const GithubApiUrl = "https://api.github.com/users/"

type User struct {
	Cache *cache.Cache
}

func NewUser(cache *cache.Cache) *User {
	return &User{
		Cache: cache,
	}
}

func (u *User) Handler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	users := r.URL.Query().Get("users")
	userList := strings.Split(users, ",")

	if len(users) == 0 {
		response.Error(w,
			http.StatusNotFound,
			errors.New("not found"),
		)
		return
	}

	var userInfoList customPkg.InfoList
	for _, user := range userList {
		var userInfo models.UserInformation
		identifier := strings.ToLower(user)

		data, expiration, isFound := u.Cache.GetWithExpiration(identifier)
		if isFound && expiration.After(time.Now()) {
			userInfo = data.(models.UserInformation)
		} else {
			name := fmt.Sprintf("%v", user)

			resp, err := http.Get(GithubApiUrl + name)
			if err != nil {
				response.Error(w, http.StatusInternalServerError, err)
				return
			}

			defer resp.Body.Close()

			body, err := io.ReadAll(resp.Body)
			if err != nil {
				response.Error(w, http.StatusInternalServerError, err)
				return
			}

			err = json.Unmarshal(body, &userInfo)
			if err != nil {
				response.Error(w, http.StatusInternalServerError, err)
				return
			}

			u.Cache.Set(identifier, userInfo, time.Duration(2)*time.Minute)
		}

		userInfoList = append(userInfoList, customPkg.Info{&userInfo})
	}

	sort.Sort(userInfoList)
	response.ToJSON(w, http.StatusOK, userInfoList)
}

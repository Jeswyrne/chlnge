package models

type UserInformation struct {
	Name       string `json:"name"`
	Login      string `json:"login"`
	Company    string `json:"company"`
	Followers  int    `json:"followers"`
	PublicRepo int    `json:"public_repos"`
}

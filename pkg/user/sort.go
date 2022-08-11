package user

import "github.com/Jeswyrne/chlnge/api/models"

type Info struct {
	*models.UserInformation
}

type InfoList []Info

func (i InfoList) Len() int {
	return len(i)
}

func (i InfoList) Less(j, k int) bool {
	return i[j].Name < i[k].Name
}

func (i InfoList) Swap(j, k int) {
	i[j], i[k] = i[k], i[j]
}

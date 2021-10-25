package profile

import (
	"encoding/json"
	profile "github.com/alicfeng/gin_template/app/component/profile/drivers"
	"github.com/alicfeng/gin_template/app/repository"
	"github.com/kataras/golog"
)

func New(code int) *Profile {
	return &Profile{Code: code}
}

func Working() *Profile {
	return New(CodeWorking).query()
}

func (profile *Profile) query() *Profile {
	database := repository.GetMySQL()
	result := database.Where("code = ?", profile.Code).First(profile)

	if 0 == result.RowsAffected {
		panic("profile error")
	}

	return profile
}

func (profile *Profile) ChooseWorking() (working profile.Working) {
	err := json.Unmarshal([]byte(profile.Content), &working)
	if nil != err {
		panic("profile content error")
	}

	return working
}

func (profile *Profile) Save(bean interface{}) bool {
	jsonBytes, err := json.Marshal(bean)

	if nil != err {
		golog.Error("save profile failure about marshal")
		return false
	}

	profile.Content = string(jsonBytes)
	result := repository.GetMySQL().Save(&profile)
	if nil != result.Error {
		golog.Error("save profile failure about database")
		return false
	}

	return true
}

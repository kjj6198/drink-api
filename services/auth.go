package services

import (
	"drink-api/utils"
	"encoding/json"
	"errors"
)

const (
	tokenEndpoint = "https://www.googleapis.com/oauth2/v3/tokeninfo"
)

type UserInfo struct {
	Email   string `json:"email"`
	Picture string `json:"picture"`
	Name    string `json:"name"`
}

func Auth(idToken string) (info *UserInfo, err error) {
	result, err := utils.DoRequest(tokenEndpoint, &utils.Options{
		Method: "GET",
		Params: map[string]interface{}{
			"id_token": idToken,
		},
	})
	user := &UserInfo{}
	err = json.Unmarshal(result, user)
	if err != nil {
		return nil, errors.New("can not receive info from google")
	}

	return user, nil
}

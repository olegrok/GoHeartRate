package database

import (
	"fmt"
	"net/http"
	"strconv"
)

func IsAuthorizedUser(cookies []*http.Cookie) bool {
	fmt.Println(cookies)
	var token, uid string
	for _, c := range cookies {
		switch c.Name {
		case "token":
			token = c.Value
		case "uid":
			uid = c.Value
		}
	}

	fmt.Printf("used_id = %s; valid = %t\n", uid, !DB.Where("user_id = ? AND token = ?", uid, token).First(&UserSession{}).RecordNotFound())
	return !DB.Where("user_id = ? AND token = ?", uid, token).First(&UserSession{}).RecordNotFound()
}

func SaveResult(id string, result float64) error {
	uid, _ := strconv.ParseUint(id, 10, 64)
	return DB.Create(&UserResult{
		UserID: uid,
		Result: fmt.Sprint(result),
	}).Error
}

func GetResults(id string) (*[]UserResult, error) {
	var res []UserResult
	var err error
	if err = DB.Where("user_id = ?", id).Order("created_at desc").Find(&res).Limit(10).Error; err != nil {
		return nil, err
	}
	for i, j := range res {
		fmt.Println(i, j.CreatedAt, j.Result)
	}
	return &res, nil
}

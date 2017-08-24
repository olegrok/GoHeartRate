package database

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// IsAuthorizedUser checks the validity of user's cookie
func IsAuthorizedUser(cookies []*http.Cookie) (*string, bool) {
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
	return &uid, !DB.Where("user_id = ? AND token = ?", uid, token).First(&UserSession{}).RecordNotFound()
}

// SaveResult saves user's result in database
func SaveResult(id string, result float64) error {
	uid, _ := strconv.ParseUint(id, 10, 64)
	if err := DB.Create(&UserResult{
		UserID: uid,
		Result: fmt.Sprint(result),
	}).Error; err != nil {
		log.Printf("result save error: %s", err)
		return err
	}
	return nil
}

// GetResults gets 10 last results by userID
func GetResults(id string) (*[]UserResult, error) {
	var res []UserResult
	var err error
	if err = DB.Where("user_id = ?", id).Order("created_at desc").Limit(10).Find(&res).Error; err != nil {
		return nil, err
	}
	for i, j := range res {
		fmt.Println(i, j.CreatedAt, j.Result)
	}
	return &res, nil
}

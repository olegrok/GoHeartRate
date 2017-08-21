package database

import (
	"fmt"
	"net/http"
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

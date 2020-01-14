package domain

import (
	"fmt"
	"github.com/silentlight/golang-microservices/mvc/utils"
	"net/http"
)

var(
	users = map[int64]*User {
		123: {Id: 1, FirstName: "Bartosz", LastName: "Hejman"},
	}
)

func GetUser(userId int64) (*User, *utils.ApplicationError) {
	user := users[userId]
	if user == nil {
		return nil, &utils.ApplicationError{
			Message:    fmt.Sprintf("user %v was not found", userId),
			StatusCode: http.StatusNotFound,
			Code:       "not_found",
		}
	}
	return user, nil
}

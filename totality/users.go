package totality

/**
This package will contains all interface and entities definations.
**/

import "context"

// User represents User.
type User struct {
	ID      int64   `json:"id"`
	Name    string  `json:"name"`
	City    string  `json:"city"`
	Phone   int64   `json:"phone"`
	Height  float32 `json:"height"`
	Married bool    `json:"married"`
}

// UserManager manages users.
type UserManager interface {
	GetUserByID(context.Context, int64) (User, error)
	GetUsers(context.Context, []int64) ([]User, error)
}

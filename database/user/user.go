package user

import "sync"

type (
	User struct {
		UserID   string `json:"id"`
		Username string `json:"name"`
	}
	Users []User
)

var (
	once  sync.Once
	users *Users
)

func Connect() {
	once.Do(func() {
		users = new(Users)
	})
}

func AddRecord(newUser User) {
	*users = append(*users, newUser)
}

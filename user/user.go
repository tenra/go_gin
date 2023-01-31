package user

import (
    "fmt"

    "github.com/tenra/go_gin/lib"
)

type User struct {
	Id   uint   `json:"id" binding:"required"`
    Name string `json:"name"`
    Password string `json:"password"`
}

type Users struct {
    Items []User
}

func New() *Users {
    return &Users{}
}

func (r *Users) Add(a User) {
    r.Items = append(r.Items, a)
    db := lib.GetDBConn().DB
    if err := db.Create(a).Error; err != nil {
        fmt.Println("err!")
    }
}

func (r *Users) GetAll() []User {
    db := lib.GetDBConn().DB
    var users []User
    if err := db.Find(&users).Error; err != nil {
        return nil
    }
    return users
}

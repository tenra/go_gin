package promotion

import (
    "fmt"

    "github.com/tenra/go_gin/lib"
)

type Promotion struct {
	Id      uint   `json:"id" binding:"required"`
    Title       string `json:"title"`
    Content string `json:"content"`
}

type Promotions struct {
    Items []Promotion
}

func New() *Promotions {
    return &Promotions{}
}

func (r *Promotions) Add(a Promotion) {
    r.Items = append(r.Items, a)
    db := lib.GetDBConn().DB
    if err := db.Create(a).Error; err != nil {
        fmt.Println("err!")
    }
}

func (r *Promotions) GetAll() []Promotion {
    db := lib.GetDBConn().DB
    var promotions []Promotion
    if err := db.Find(&promotions).Error; err != nil {
        return nil
    }
    return promotions
}

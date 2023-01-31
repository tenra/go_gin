package handler

import (
    "net/http"

    "github.com/tenra/go_gin/user"
    "github.com/gin-gonic/gin"
)

func UsersGet(users *user.Users) gin.HandlerFunc {
    return func(c *gin.Context) {
        result := users.GetAll()
        c.JSON(http.StatusOK, result)
    }
}

type User struct {
    Id   uint   `json:"id" binding:"required"`
    Name string `json:"name" binding:"required"`
    //Password string `json:"password"`
}

func UserPost(post *user.Users) gin.HandlerFunc {
    return func(c *gin.Context) {
        requestBody := User{}
        c.Bind(&requestBody)

        item := user.User{
            Id: requestBody.Id,
            Name: requestBody.Name,
            //Password: requestBody.Password,
        }
        post.Add(item)

        c.Status(http.StatusNoContent)
    }
}

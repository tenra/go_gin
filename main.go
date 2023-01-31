package main

import (
    "os"
    "time"

    "github.com/gin-gonic/gin"
    "github.com/tenra/go_gin/promotion"
    "github.com/tenra/go_gin/handler"
    "github.com/tenra/go_gin/lib"
    "github.com/tenra/go_gin/user"
    "github.com/joho/godotenv"

    "github.com/gin-contrib/cors"
)

func main() {
    if os.Getenv("USE_HEROKU") != "1" {
        err := godotenv.Load()
        if err != nil {
            panic(err)
        }
    }

    promotion := promotion.New()
    user := user.New()

    lib.DBOpen()
    defer lib.DBClose()

    r := gin.Default()

    r.Use(cors.New(cors.Config{
        AllowOrigins: []string{
			os.Getenv("FRONT_URL"),
        },
        AllowMethods: []string{
            "POST",
            "GET",
            "OPTIONS",
        },
        AllowHeaders: []string{
            "Access-Control-Allow-Credentials",
            "Access-Control-Allow-Headers",
            "Content-Type",
            "Content-Length",
            "Accept-Encoding",
            "Authorization",
        },
        AllowCredentials: true,
        MaxAge:           24 * time.Hour,
    }))

    r.GET("/promotions", handler.Index(promotion))
    r.POST("/promotions", handler.Create(promotion))
    r.POST("/user/login", handler.UserPost(user))

    r.Run(os.Getenv("HTTP_HOST") + ":" + os.Getenv("HTTP_PORT"))
}

package main

import (
    "os"
    "time"
    //"log"
	"net/http"

    "github.com/gin-gonic/gin"
    "github.com/tenra/go_gin/promotion"
    "github.com/tenra/go_gin/handler"
    "github.com/tenra/go_gin/lib"
    "github.com/tenra/go_gin/user"
    "github.com/joho/godotenv"

    "github.com/gin-contrib/cors"

    "github.com/tenra/go_gin/middleware"
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

    router := http.NewServeMux()

	// This route is always accessible.
	router.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"Hello from a public endpoint! You don't need to be authenticated to see this."}`))
	}))

	// This route is only accessible if the user has a valid access_token.
	router.Handle("/mypage", middleware.EnsureValidToken()(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// CORS Headers.
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			w.Header().Set("Access-Control-Allow-Origin", os.Getenv("FRONT_URL"))
			w.Header().Set("Access-Control-Allow-Headers", "Authorization")

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"message":"Hello from a private endpoint! You need to be authenticated to see this."}`))
		}),
	))

    r.GET("/promotions", handler.Index(promotion))
    r.POST("/promotions", handler.Create(promotion))
    r.POST("/user/login", handler.UserPost(user))

    r.Run(os.Getenv("HTTP_HOST") + ":" + os.Getenv("HTTP_PORT"))
}

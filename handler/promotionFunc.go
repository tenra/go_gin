package handler

import (
    "net/http"

    "github.com/tenra/go_gin/promotion"
    "github.com/gin-gonic/gin"
)

type PromotionPostRequest struct {
    Id      uint   `json:"id" binding:"required"`
    Title   string `json:"title"`
    Content string `json:"content" binding:"required"`
    //Description string `json:"description"`
}

// Index action: GET /promotions
func Index(promotions *promotion.Promotions) gin.HandlerFunc {
    return func(c *gin.Context) {
        result := promotions.GetAll()
        c.JSON(http.StatusOK, result)
    }
}

// Create action: POST /promotions
func Create(post *promotion.Promotions) gin.HandlerFunc {
    return func(c *gin.Context) {
        requestBody := PromotionPostRequest{}
        c.Bind(&requestBody)

        item := promotion.Promotion{
            Id: requestBody.Id,
            Title:       requestBody.Title,
            Content: requestBody.Content,
        }
        post.Add(item)

        c.Status(http.StatusNoContent)
    }
}

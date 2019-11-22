package middleware

import (
  "github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
  return func(c *gin.Context) {
    c.Set("uid", int64(100000007))
    c.Set("group_id", uint(1))
    c.Next()
  }
}
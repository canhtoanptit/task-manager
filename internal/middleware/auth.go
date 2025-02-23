package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"strings"
	"task-management-api/internal/config"
	"task-management-api/internal/models"
)

func Authenticate(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "No token provided"})
		c.Abort()
		return
	}

	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.AppConfig.JWTSecret), nil
	})

	if err != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		c.Abort()
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
		c.Abort()
		return
	}

	c.Set("userID", uint(claims["id"].(float64)))
	c.Set("role", claims["role"])
	c.Next()
}

func EmployerRequired(c *gin.Context) {
	if c.GetString("role") != string(models.Employer) {
		c.JSON(http.StatusForbidden, gin.H{"error": "Employer privileges required"})
		c.Abort()
		return
	}
	c.Next()
}

func EmployeeRequired(c *gin.Context) {
	if c.GetString("role") != string(models.Employee) {
		c.JSON(http.StatusForbidden, gin.H{"error": "Employee privileges required"})
		c.Abort()
		return
	}
	c.Next()
}

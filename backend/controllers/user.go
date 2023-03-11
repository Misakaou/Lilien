package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"

	"lilien-server/database"
)

type UserController struct{}

func (u UserController) Retrieve(c *gin.Context) {
	if c.Param("id") != "" {
		user := database.GetDatabaseClient().Database("golangAPI").Collection("test").FindOne(c, bson.M{"_id": c.Param("id")})
		if user != nil {
			c.JSON(http.StatusOK, user)
			return
		}
	}
	c.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
	c.Abort()
}

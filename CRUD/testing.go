package CRUD

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gym/models"
	"net/http"
)

func GetUsersDetail(c *gin.Context) {
	c.JSON(http.StatusOK, models.Users)
}

func GetUsers(c *gin.Context) {

	var param struct {
		Name string `uri:"name"`
	}

	if err := c.ShouldBindUri(&param); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	for _, u := range models.Users {
		if u.Name == param.Name {
			c.JSON(http.StatusOK, u)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{
		"error": fmt.Sprintf("user %s not found", param.Name),
	})

}

func GetAgeUsers(c *gin.Context) {
	var query struct {
		Uid int `form:"ID"`
	}

	if err := c.ShouldBindQuery(&query); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("user %s not found", query.Uid),
		})
		return
	}

	for _, u := range models.Users {
		if u.ID == query.Uid {
			c.JSON(http.StatusOK, u)
		}
		return
	}
	c.JSON(http.StatusNotFound, gin.H{
		"error": fmt.Sprintf("user %s not found", query.Uid),
	})
}

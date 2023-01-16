package CRUD

import (
	"github.com/gin-gonic/gin"
	"github.com/gym/ConnectDB"
	"github.com/gym/models"
	"net/http"
)

func CreateUser(c *gin.Context) {
	var Users models.DataUser

	db, err := ConnectDB.GetDBUsers()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	if err = c.ShouldBindJSON(&Users); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	db.Create(&Users)
}

func CreateLocation(c *gin.Context) {
	var Location models.Location

	db, err := ConnectDB.GetDBLocation()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	if err = c.ShouldBindJSON(&Location); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	db.Create(&Location)
}

func CreateSchool(c *gin.Context) {
	var School models.School

	db, err := ConnectDB.GetDBSchool()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	if err = c.ShouldBindJSON(&School); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	db.Create(&School)
}

func FindUsers(c *gin.Context) {
	var Users models.DataUser

	db, err := ConnectDB.GetDBUsers()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	db.Find(&Users)
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}
func FindLocation(c *gin.Context) {
	var Location models.Location

	db, err := ConnectDB.GetDBLocation()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	db.Find(&Location)
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}
func FindSchool(c *gin.Context) {
	var School models.School

	db, err := ConnectDB.GetDBSchool()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	db.Find(&School)
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}

func UpdateUsers(c *gin.Context) {

	var Users models.DataUser
	var UpdateUsers models.UpdateDataUser

	db, err := ConnectDB.GetDBUsers()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	if err = db.Where("id = ?", c.Param("id")).First(&Users).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	if err = c.ShouldBindJSON(&UpdateUsers); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	db.Model(&Users).Updates(UpdateUsers)
}

func UpdateLocation(c *gin.Context) {

	var Location models.Location
	var UpdateLocation models.UpdateLocation

	db, err := ConnectDB.GetDBLocation()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	if err = db.Where("id = ?", c.Param("id")).First(&Location).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	if err = c.ShouldBindJSON(&UpdateLocation); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	db.Model(&Location).Updates(UpdateLocation)
}

func UpdateSchool(c *gin.Context) {

	var School models.School
	var UpdateSchool models.UpdateSchool

	db, err := ConnectDB.GetDBSchool()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	if err = db.Where("id = ?", c.Param("id")).First(&School).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	if err = c.ShouldBindJSON(&UpdateSchool); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	db.Model(&School).Updates(UpdateSchool)
}

func DeleteUsers(c *gin.Context) {

	var Users models.DataUser

	db, err := ConnectDB.GetDBUsers()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	if err = db.Where("id = ?", c.Param("id")).First(&Users).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	db.Delete(&Users)
}

func DeleteLocation(c *gin.Context) {

	var Location models.Location

	db, err := ConnectDB.GetDBLocation()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	if err = db.Where("id = ?", c.Param("id")).First(&Location).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	db.Delete(&Location)
}

func DeleteSchool(c *gin.Context) {

	var School models.School

	db, err := ConnectDB.GetDBSchool()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	if err = db.Where("id = ?", c.Param("id")).First(&School).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	db.Delete(&School)
}

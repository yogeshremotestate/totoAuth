package controllers

import (
	"LearnGo-todoAuth/initializers"
	"LearnGo-todoAuth/models"

	"github.com/gin-gonic/gin"
)

func NoteCreate(c *gin.Context) {
	// get body

	var body struct {
		Title string
		Body  string
	}
	c.Bind(&body) // pas the

	// create Note

	post := models.Note{Title: body.Title, Body: body.Body}
	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	//return
	c.JSON(200, gin.H{
		"Post": post, // can not return result var as it does not show anything on response
	})
}

func GetAllNote(c *gin.Context) {

	// get all Notes

	var posts []models.Note

	initializers.DB.Find(&posts)

	//return
	c.JSON(200, gin.H{
		"Posts": posts,
	})
}

func GetNote(c *gin.Context) {
	//get id
	id := c.Params
	// fmt.Println(id)

	// get all Notes

	var post models.Note

	initializers.DB.First(&post, id[0].Value)

	//return
	c.JSON(200, gin.H{
		"Post": post,
	})
}

func UpdateNote(c *gin.Context) {
	// get body and id

	id := c.Params

	var body struct {
		Title string
		Body  string
	}
	c.Bind(&body) // pass the boddy

	// find and update

	var post models.Note

	initializers.DB.First(&post, id[0].Value)

	result := initializers.DB.Model(&post).Updates(models.Note{Title: body.Title, Body: body.Body})

	if result.Error != nil {
		c.Status(400)
		return
	}

	//return
	c.JSON(200, gin.H{
		"Post": post, // can not return result var as it does not show anything on response
	})
}

func DeleteNote(c *gin.Context) {
	//get id
	id := c.Params
	// fmt.Println(id)

	// get all Notes

	initializers.DB.Delete(&models.Note{}, id[0].Value)

	//return
	c.Status(200)
}

package v1

import "github.com/gin-gonic/gin"

type User struct{}

func NewUser() User {
	return User{}
}

func (u User) Get(c *gin.Context) {

	c.JSON(200, gin.H{"msg": "success"})
}

func (u User) List(c *gin.Context) {

	c.JSON(200, gin.H{"msg": "success"})
}

func (u User) Create(c *gin.Context) {

	c.JSON(200, gin.H{"msg": "success"})
}

func (u User) Update(c *gin.Context) {

	c.JSON(200, gin.H{"msg": "success"})
}

func (u User) Delete(c *gin.Context) {

	c.JSON(200, gin.H{"msg": "success"})
}

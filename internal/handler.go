package internal

import (
	"my_local_communitate/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Upload(c *gin.Context) {
	if err := upload(c); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 1, "msg": err.Error()})
	}
}

func upload(c *gin.Context) error {
	var uploadRequest model.UploadRequest
	if err := c.ShouldBind(&uploadRequest); err != nil {
		return err
	}

	file, err := c.FormFile("data")
	if err != nil {
		return err
	}

	c.SaveUploadedFile(file, "")

	return nil
}

func KeyGen(c *gin.Context) {
	if err := keyGen(c); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 1, "msg": err.Error()})
	}
}

func keyGen(c *gin.Context) error {
	var keyGenRequest model.KeyGenRequest
	if err := c.ShouldBind(&keyGenRequest); err != nil {
		return err
	}

	// 责任链模式
	switch keyGenRequest.Step {
	case 1:

	case 2:

	}

	return nil
}

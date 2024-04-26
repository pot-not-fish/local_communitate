package internal

import (
	"bytes"
	"fmt"
	"io"
	"my_local_communitate/model"
	"my_local_communitate/pkg/crypto/crypto_md5"
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

	// md5校验
	fd, err := file.Open()
	if err != nil {
		return err
	}
	buf := new(bytes.Buffer)
	io.Copy(buf, fd)
	md5 := crypto_md5.NewHashMD5()
	if !md5.IsValid(buf.Bytes(), uploadRequest.Signature) {
		return fmt.Errorf("invalid file")
	}

	// DES解密文件

	c.JSON(http.StatusOK, model.UploadResponse{Code: 0, Msg: "OK"})
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

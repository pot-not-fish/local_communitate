package internal

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"my_local_communitate/model"
	"my_local_communitate/pkg/cache/group"
	"my_local_communitate/pkg/crypto/crypto_des"
	"my_local_communitate/pkg/crypto/crypto_md5"
	"my_local_communitate/pkg/crypto/diffie_hellman"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

var (
	CTX context.Context
)

func Upload(c *gin.Context) {
	if err := upload(c); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 1, "msg": err.Error()})
	}
}

func upload(c *gin.Context) error {
	var uploadRequest model.UploadRequest
	if err := c.ShouldBind(&uploadRequest); err != nil {
		fmt.Printf("0 ")
		fmt.Println(err.Error())
		return err
	}

	file, err := c.FormFile("data")
	if err != nil {
		fmt.Printf("1 ")
		fmt.Println(err.Error())
		return err
	}

	fd, err := file.Open()
	if err != nil {
		fmt.Printf("2 ")
		fmt.Println(err.Error())
		return err
	}
	buf := new(bytes.Buffer)
	io.Copy(buf, fd)

	// DES解密文件 先解密再校验，不然有被篡改的风险
	secretKey, err := group.Get("symmetric_key", uploadRequest.URL)
	if err != nil {
		fmt.Print("3 ")
		fmt.Println(err.Error())
		return err
	}
	des := crypto_des.NewCrypto(string(secretKey))

	fmt.Println("server symmetric key ", string(secretKey))

	cipher_data, err := des.Decrypto(buf.Bytes())
	if err != nil {
		fmt.Printf("4 ")
		fmt.Println(err.Error())
		return err
	}

	// md5校验
	md5 := crypto_md5.NewHashMD5()
	if !md5.IsValid(cipher_data, uploadRequest.Signature) {
		return fmt.Errorf("invalid file")
	}

	os.WriteFile(file.Filename, cipher_data, 0644)

	runtime.EventsEmit(CTX, "upload_list", fmt.Sprintf("%d-%d-%d  %d:%d:%d", time.Now().Year(), time.Now().Month(), time.Now().Day(), time.Now().Hour(), time.Now().Minute(), time.Now().Second()), file.Filename, "接收成功")

	group.Del("asymmetric_key", uploadRequest.URL)
	group.Del("symmetric_key", uploadRequest.URL)

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

	switch keyGenRequest.Step {
	case 1:
		if len(keyGenRequest.Share) != 1 {
			return fmt.Errorf("invalid share first step")
		}
		diffie := diffie_hellman.NewCrypto()

		fmt.Println("server g^a ", keyGenRequest.Share[0])
		fmt.Println("server r", diffie.Get_R())

		symmetric_key := diffie_hellman.QuickPower(keyGenRequest.Share[0], diffie.Get_R())
		symmetric_key_str := fmt.Sprintf("%d", symmetric_key)
		group.Set("asymmetric_key", keyGenRequest.URL, []byte(symmetric_key_str))
		c.JSON(http.StatusOK, model.KeyGenResponse{Code: 0, Msg: "OK", Share: []int64{diffie.Get_GR()}})
	case 2:
		if len(keyGenRequest.Share) != 1 {
			return fmt.Errorf("invalid share second step")
		}
		val, err := group.Get("asymmetric_key", keyGenRequest.URL)
		if err != nil {
			return err
		}
		data, _ := strconv.ParseInt(string(val), 10, 64)
		symmetric_key := (diffie_hellman.QuickPower(data, diffie_hellman.GetP()-2) * keyGenRequest.Share[0]) % diffie_hellman.GetP()
		group.Set("symmetric_key", keyGenRequest.URL, []byte(fmt.Sprintf("%d", symmetric_key)))

		fmt.Println("keygen symmetric key ", symmetric_key)

		c.JSON(http.StatusOK, model.KeyGenResponse{Code: 0, Msg: "OK"})
	}
	return nil
}

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "OK"})
}

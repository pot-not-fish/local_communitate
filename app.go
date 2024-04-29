package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"my_local_communitate/internal"
	"my_local_communitate/model"
	"my_local_communitate/pkg/crypto/crypto_des"
	"my_local_communitate/pkg/crypto/crypto_md5"
	"my_local_communitate/pkg/crypto/diffie_hellman"
	"net"
	"net/http"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	internal.CTX = ctx
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) Send(path string, name string, data []byte) string {
	// 检查地址是否有效
	path = fmt.Sprintf("http://%s:5000", path)

	// 获取本机IP
	ipv4, err := GetLocalIP()
	if err != nil {
		return err.Error()
	}

	// 密钥交换
	// 发送g^a 返回g^b
	// 发送x*g^ab 返回确认
	diffie := diffie_hellman.NewCrypto()
	step_one_request := model.KeyGenRequest{URL: ipv4, Step: 1, Share: []int64{diffie.Get_GR()}}
	step_one_response, err := KeyGenRequest(path, step_one_request)
	if err != nil {
		return err.Error()
	}
	x := diffie_hellman.GetRandom()
	diffie.SetGAB(diffie_hellman.QuickPower(step_one_response.Share[0], diffie.Get_R()))

	fmt.Println("client symmetric key ", x, " a ", diffie.Get_R(), " g^a ", diffie.Get_GR())

	cipher_data := diffie.Encrypto(x)
	step_two_request := model.KeyGenRequest{URL: ipv4, Step: 2, Share: []int64{cipher_data}}
	_, err = KeyGenRequest(path, step_two_request)
	if err != nil {
		return err.Error()
	}

	// 校验文件
	signature := crypto_md5.GenerateHash(data)

	// 加密文件 发送文件
	des := crypto_des.NewCrypto(fmt.Sprintf("%d", x))
	cipher_file, err := des.Encrypto(data)
	if err != nil {
		return err.Error()
	}

	_, err = UploadRequest(path, cipher_file, name, signature, ipv4)
	if err != nil {
		return err.Error()
	}

	return "OK"
}

// 获取本机网卡IP
func GetLocalIP() (ipv4 string, err error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return
	}
	// 取第一个非lo的网卡IP
	for _, addr := range addrs {
		// 这个网络地址是IP地址: ipv4, ipv6
		ipNet, isIpNet := addr.(*net.IPNet)
		if isIpNet && !ipNet.IP.IsLoopback() {
			// 跳过IPV6
			if ipNet.IP.To4() != nil {
				ipv4 = ipNet.IP.String() // 192.168.1.1
				return
			}
		}
	}
	return
}

func UploadRequest(path string, data []byte, name string, signature string, origin string) (response *model.UploadResponse, err error) {
	path = fmt.Sprintf("%s/upload", path)

	buf := new(bytes.Buffer)
	bw := multipart.NewWriter(buf) // body writer

	sig, _ := bw.CreateFormField("signature")
	sig.Write([]byte(signature))

	url, _ := bw.CreateFormField("url")
	url.Write([]byte(origin))

	// 要使用文件的格式
	file, _ := bw.CreateFormFile("data", name)
	reader := bytes.NewReader(data)
	io.Copy(file, reader)
	// 写完要关闭，否则，请求体会缺少结束边界
	bw.Close()

	req, err := http.NewRequest("POST", path, buf)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", bw.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	_, err = io.Copy(buf, resp.Body)
	if err != nil {
		return nil, err
	}

	response = new(model.UploadResponse)
	err = json.Unmarshal(buf.Bytes(), response)
	if err != nil {
		return nil, err
	}

	if response.Code != 0 {
		return nil, fmt.Errorf(response.Msg)
	}
	return response, nil
}

func KeyGenRequest(path string, request model.KeyGenRequest) (response *model.KeyGenResponse, err error) {
	path = fmt.Sprintf("%s/keygen", path)

	request_json, _ := json.Marshal(request)

	resp, err := http.Post(path, "application/json;charset=utf-8", bytes.NewBuffer(request_json))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	buf := new(bytes.Buffer)
	_, err = io.Copy(buf, resp.Body)
	if err != nil {
		return nil, err
	}

	fmt.Println(buf.String())
	response = new(model.KeyGenResponse)
	err = json.Unmarshal(buf.Bytes(), response)
	if err != nil {
		return nil, err
	}

	if response.Code != 0 {
		return nil, fmt.Errorf(response.Msg)
	}
	return response, nil
}

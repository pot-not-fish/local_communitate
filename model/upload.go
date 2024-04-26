package model

type UploadRequest struct {
	URL       string `form:"url" json:"url" xml:"url"`
	Signature string `form:"signature" json:"signature" xml:"signature"`
}

type UploadResponse struct {
	Code int    `form:"code" json:"code" xml:"code"`
	Msg  string `form:"msg" json:"msg" xml:"msg"`
}

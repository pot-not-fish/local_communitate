package model

type UploadRequest struct {
	URL string `json:"url"`
}

type UploadResponse struct {
	Code int    `form:"code" json:"code" xml:"code"`
	Msg  string `form:"msg" json:"msg" xml:"msg"`
}

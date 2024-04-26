package model

type KeyGenRequest struct {
	URL   string   `form:"url" json:"url" xml:"url"`
	Step  int      `form:"step" json:"step" xml:"step"`
	Share []string `form:"share" json:"share" xml:"share"`
}

type KeyGenResponse struct {
	Code int    `form:"code" json:"code" xml:"code"`
	Msg  string `form:"msg" json:"msg" xml:"msg"`

	Step  int      `form:"step" json:"step" xml:"step"`
	Share []string `form:"share" json:"share" xml:"share"`
}

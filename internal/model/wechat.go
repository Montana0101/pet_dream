package model

type WechatLogin struct {
	Errcode    int    `json:"errcode"`
	Errmsg     string `json:"errmsg"`
	Sessionkey string `json:"session_key"`
	Openid     string `json:"openid"`
}

package model

// User 用户
type User struct {
	Id        int      `json:"id" `
	NickName  *string  `json:"nick_name"'`
	RealName  *string  `json:"real_name"`
	Identity  *string  `json:"identity"`
	Age       *int     `json:"age"`
	Gender    *int     `json:"gender"`
	Longitude *float64 `json:"longitude"`
	Latitude  *float64 `json:"latitude"`
	Openid    *string  `json:"openid"`
	Code      string   `json:"code"`
	City      *string  `json:"city"`
	District  *string  `json:"district"`
	AvatarUrl *string  `json:"avatar_url"`
}

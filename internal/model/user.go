package model

// User 用户
type User struct {
	Id         int      `json:"id" `
	NickName   string   `json:"nick_name"'`
	RealName   *string  `json:"real_name"`
	IdentityId *string  `json:"identity_id"`
	Age        *int     `json:"age"`
	Gender     *int     `json:"gender"`
	Longitude  *float64 `json:"longitude"`
	Latitude   *float64 `json:"latitude"`
}

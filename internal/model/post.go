package model

// 贴文
type Post struct {
	Id         *int    `json:"id"`
	Title      *string `json:"title"`
	UserId     *int    `json:"user_id"`
	Content    *string `json:"content"`
	CreateTime *string `json:"create_time"`
	UpdateTime *string `json:"update_time"`
}

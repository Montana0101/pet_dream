package model

// Post 贴文
type Post struct {
	Id         *int    `json:"id"`
	Title      *string `json:"title"`
	UserId     *int    `json:"user_id"`
	PetId      *int    `json:"pet_id"`
	Content    *string `json:"content"`
	CreateTime *string `json:"create_time"`
	UpdateTime *string `json:"update_time"`
}

// 首页推荐
type Recommand struct {
	PageNo   *string `json:"page_no"`
	PageSize *string `json:"page_size"`
}

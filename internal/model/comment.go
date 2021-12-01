package model

// Comment 评论
type Comment struct {
	Id         *int    `json:"id"`
	PostId     *int    `json:"post_id"`
	UserId     *int    `json:"user_id"`
	Content    *string `json:"content"`
	CreateTime *string `json:"create_time"`
	UpdateTime *string `json:"update_time"`
	Source     *int    `json:"source"`
	ReplyId    *int    `json:"reply_id"`
	ParentId   int     `json:"parent_id"`
}

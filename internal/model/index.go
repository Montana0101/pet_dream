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
	Media      []struct {
		Url  *string `json:"url"`
		Type *string `json:"type"`
	} `json:"media"`
}

// Diary 猫咪日志
type Diary struct {
	Id         *int    `json:"id"`
	Title      *string `json:"title"`
	UserId     *int    `json:"user_id"`
	PetId      *int    `json:"pet_id"`
	Content    *string `json:"content"`
	CreateTime *string `json:"create_time"`
	UpdateTime *string `json:"update_time"`
	Media      []struct {
		Url  *string `json:"url"`
		Type *string `json:"type"`
	} `json:"media"`
}

type Media struct {
	Url  *string `json:"url"`
	Type *string `json:"type"`
}

// 首页推荐
type Recommand struct {
	PageNo   *string `json:"page_no"`
	PageSize *string `json:"page_size"`
}

// Pet 宠物
type Pet struct {
	Id     *int    `json:"id"`
	UserId *int    `json:"user_id"'`
	Name   *string `json:"name"'`
	Age    *int    `json:"age"'`
	Gender *int    `json:"gender"'`
	Breed  *int    `json:"breed"`
	Intro  *string `json:"intro"`
	Photo  *string `json:"photo"`
	//source float64 `json:"source"`
}

type Tag struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Value int    `json:"value"`
}

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

type WechatLogin struct {
	Errcode    int    `json:"errcode"`
	Errmsg     string `json:"errmsg"`
	Sessionkey string `json:"session_key"`
	Openid     string `json:"openid"`
}

// Comment 评论
type Comment struct {
	Id         int     `json:"id"`
	PostId     *int    `json:"post_id"`
	UserId     *int    `json:"user_id"`
	Content    *string `json:"content"`
	CreateTime *string `json:"create_time"`
	UpdateTime *string `json:"update_time"`
	Source     *int    `json:"source"`
	ReplyId    *int    `json:"reply_id"`
	ParentId   int     `json:"parent_id"`
	ReplyName  *string `json:"reply_name"`
}

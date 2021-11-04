package model

// Pet 宠物
type Pet struct {
	Id     *int    `json:"id"`
	UserId *int    `json:"user_id"'`
	Name   *string `json:"name"'`
	Age    *int    `json:"age"'`
	Gender *int    `json:"gender"'`
	Breed  *int    `json:"breed"`
	Intro  *string `json:"intro"`
	//source float64 `json:"source"`
}

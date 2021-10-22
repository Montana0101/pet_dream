package model

type Pet struct {
	Id *int `json:"id"`
	UserId *int `json:"user_id"'`
	Name *string `json:"name"'`
	Age *int `json:"age"'`
	Gender *int `json:"gender"'`
	//source float64 `json:"source"`
}

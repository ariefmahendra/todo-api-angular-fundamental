package model

type Todo struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	IsCompleted bool   `json:"isCompleted"`
	CreatedAt   string `json:"createdAt"`
}

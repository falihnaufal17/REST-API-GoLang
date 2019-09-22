package main

type Users struct {
	UserId   string `form:"userid" json:"userid"`
	Name     string `form:"name" json:"name"`
	Phone    string `form:"phone" json:"phone"`
	Location string `form:"location" json:"location"`
}

type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []Users
}

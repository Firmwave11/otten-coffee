package models

import "time"

type Data struct {
	Status struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	} `json:"status"`
	Data struct {
		Receivedby string      `json:"receivedBy"`
		Histories  []Historeis `json:"histories"`
	} `json:"data"`
}

type Historeis struct {
	Description string    `json:"description"`
	Createdat   time.Time `json:"createdAt"`
	Formatted   struct {
		Createdat string `json:"createdAt"`
	} `json:"formatted"`
}

package main

import "gorm.io/gorm"

type User struct {
	gorm.Model
	DisplayName string `json:"displayname"`
	Email       string `json:"email"`
	Username    string `json:"username"`
	Phone       string `json:"phone"`
	Address     string `json:"address"`
	Age         int    `json:"age"`
	Avatar      string `json:"avatar"`
	Flag        bool   `json:"flag"`
}

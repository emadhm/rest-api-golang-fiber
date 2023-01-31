package models

import "github.com/jinzhu/gorm"

type Auths struct {

	gorm.Model // by default create ID CreatedAt UpdatedAt DeletedAt
    
	User_id   int `json:"user_id"`
	Token string `json:"token"`
	

}
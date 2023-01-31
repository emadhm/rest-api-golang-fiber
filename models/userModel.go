package models
 
import "github.com/jinzhu/gorm"

type Users struct {

	gorm.Model // by default create ID CreatedAt UpdatedAt DeletedAt
    
	Name   string `json:"name"`
	Email string  `json:"email"`
	Password string `json:"password"`
	Role string `json:"role"`
	

}
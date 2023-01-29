package models
 
import "github.com/jinzhu/gorm"

type Users struct {

	gorm.Model // by default create ID CreatedAt UpdatedAt DeletedAt
    
	Name   string `json:"name"`
	Person string `json:"person"`
	

}
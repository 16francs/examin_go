package model

type User struct {
	Base
	LoginID  string `json:"login_id" gorm:"unique; not null"`
	Password string `json:"-"`
	Name 		 string `json:"name"`	
	School 	 string `json:"school"`
	Role 		 int 		`json:"role"`
}

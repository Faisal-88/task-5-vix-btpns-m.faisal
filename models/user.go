package models

type User struct {
	ID       uint64 `gorm:"primary_key:auto_increment" json:"id"`
	Username string `gorm:"type:varchar(255)" json:"username"`
	Email    string `gorm:"type:varchar(255)" json:"email"`
	Password string `gorm:"type:varchar(255)" json:"password"`
	// Token    string `gorm:"-" json:"token,omitempty"`
}

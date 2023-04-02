package models

type Photos struct {
	ID       uint64 `gorm:"primary_key:auto_increment" json:"id"`
	Title    string `gorm:"type:varchar(255)" json:"title"`
	Caption  string `gorm:"type:varchar(255)" json:"caption"`
	PhotoUrl string `gorm:"type:varchar(255)" json:"photo-url"`
	// Token    string `gorm:"-" json:"token,omitempty"`
	UserId uint64 `gorm:"not null" json:"-"`
	// User   User   `gorm:"foreignKey:UserID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"user"`
}

func GetAllPhotos() ([]Photos, error) {
	var photos []Photos
	if err := DB.Find(&photos).Error; err != nil {
		return nil, err
	}
	return photos, nil
}

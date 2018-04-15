package catalog

type Publication struct {
	PublicationId int `gorm:"primary_key"`
	PublicationFrequency int8 // in weeks
	Publisher string `gorm:"size:128"`
	Genre string `gorm:"size:128"`
}

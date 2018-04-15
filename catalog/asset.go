package catalog

type Asset struct {
	AssetId   int `gorm:"primary_key" json:"id"`
	AssetType int `gorm:"default:1" json:"type"` // default is magazine(type:1)
	MagazineId int `gorm:"type:bigint REFERENCES magazines(magazine_id)" json:"-"`
}
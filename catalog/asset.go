package catalog

type Asset struct {
	AssetId int `gorm:"primary_key"`
	AssetType int `gorm:"default:1"` // default is magazine(type:1)

	MagazineId int `gorm:"type:bigint REFERENCES magazines(magazine_id)"`
}





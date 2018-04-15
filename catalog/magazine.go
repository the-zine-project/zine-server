package catalog

type Magazine struct {
	MagazineId int `gorm:"primary_key"`
	Title string `gorm:"size:128"`
	Description string `gorm:"size:2048"`
	IssueNumber int
	IssueTitle string `gorm:"size:128"`
	Cost int
	// coverImage

	PublicationId int `gorm:"type:bigint REFERENCES publications(publication_id)"`
}


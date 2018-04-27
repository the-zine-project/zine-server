package catalog

import (
	"github.com/graphql-go/graphql"
	"zine/db"
)

type Magazine struct {
	MagazineId    int    `gorm:"primary_key" json:"id"`
	Title         string `gorm:"size:128"`
	Description   string `gorm:"size:2048"`
	IssueNumber   int
	IssueTitle    string `gorm:"size:128"`
	Cost          int
	CoverImageURL string `gorm:"size:1024"`

	PublicationId int `gorm:"type:bigint REFERENCES publications(publication_id)"`
}

var MagazineType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Magazine",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"title": &graphql.Field{
				Type: graphql.String,
			},
			"issueNumber": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var MagazinesField = graphql.Field{
	Type: graphql.NewList(MagazineType),
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		var magazines []Magazine
		if id, ok := p.Args["id"]; ok {
			db.DBCon.Where(&Magazine{MagazineId: id.(int)}).Find(&magazines)
		} else {
			db.DBCon.Find(&magazines)
		}
		return magazines, nil
	},
}

package catalog

import (
	"github.com/graphql-go/graphql"
	"github.com/the-zine-project/zine-server/db"
)

type Publication struct {
	PublicationId        int    `gorm:"primary_key" json:"id"`
	PublicationFrequency int8   // in weeks
	Publisher            string `gorm:"size:128"`
	Genre                string `gorm:"size:128"`

	Magazines []Magazine
}

var PublicationType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Magazine",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"genre": &graphql.Field{
				Type: graphql.String,
			},
			"publicationFrequency": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var PublicationField = graphql.Field{
	Type: graphql.NewList(PublicationType),
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		var publications []Publication
		if id, ok := p.Args["id"]; ok {
			db.DBCon.Where(&Publication{PublicationId: id.(int)}).Find(&publications)
		} else {
			db.DBCon.Find(&publications)
		}
		return publications, nil
	},
}

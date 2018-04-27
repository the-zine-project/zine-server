package catalog

import (
	"github.com/graphql-go/graphql"
	"zine/db"
)

type Asset struct {
	AssetId    int `gorm:"primary_key" json:"id"`
	AssetType  int `gorm:"default:1" json:"type"` // default is magazine(type:1)
	MagazineId int `gorm:"type:bigint REFERENCES magazines(magazine_id)" json:"-"`
}

var AssetType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Asset",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"type": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var AssetsField = graphql.Fields{
	"assets": &graphql.Field{
		Type: graphql.NewList(AssetType),
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"type": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			var assets []Asset
			if id, ok := p.Args["id"]; ok {
				db.DBCon.Where(&Asset{AssetId: id.(int)}).Find(&assets)
			} else {
				db.DBCon.Find(&assets)
			}
			return assets, nil
		},
	},
}

package config

import (
	"@dragon-cli-template/apps/api/db"
	"@dragon-cli-template/apps/api/routes/content"
	contentmodel "@dragon-cli-template/apps/api/routes/contentModel"
	"@dragon-cli-template/apps/api/routes/field"
	fieldtype "@dragon-cli-template/apps/api/routes/fieldType"
)

func Init() {
	db.DB.AutoMigrate(&contentmodel.ContentModel{})
	db.DB.AutoMigrate(&fieldtype.FieldType{}, &field.Field{}, &content.Content{})
}

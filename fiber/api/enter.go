package api

import (
	"studyfiber/api/admin"
	"studyfiber/api/test"
	"studyfiber/api/upload"
)

type ApiGroup struct {
	TestApiGroup   test.ApiGroup
	AdminApiGroup  admin.ApiGroup
	UplaodApiGroup upload.ApiGroup
}

var ApiGroupApp = new(ApiGroup)

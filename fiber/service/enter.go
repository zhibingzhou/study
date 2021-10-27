package service

import (
	"studyfiber/service/admin"
	"studyfiber/service/upload"
)

type ServiceGroup struct {
	AdminServiceGroup  admin.ServiceGroup
	UploadServiceGroup upload.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)

package upload

import "studyfiber/service"

type ApiGroup struct {
	SimpleUplader
}

var uploadService = service.ServiceGroupApp.UploadServiceGroup

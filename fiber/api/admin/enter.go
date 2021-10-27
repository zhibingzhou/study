package admin

import "studyfiber/service"

type ApiGroup struct {
	FileHistory
}

var adminService = service.ServiceGroupApp.AdminServiceGroup

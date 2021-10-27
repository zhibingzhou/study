package router

import "studyfiber/api"

func AdminRouterInit() {
	var admin_historyrouter = api.ApiGroupApp.AdminApiGroup.FileHistory
	Router.Post("/admin/del_history.do", admin_historyrouter.HistoryDel)
}

func TestRouterInit() {
	var test_indexrouter = api.ApiGroupApp.TestApiGroup.IndexApi
	Router.Get("/test/test.do", test_indexrouter.Test)
}

func UploadRouterInit() {
	var upload_router = api.ApiGroupApp.UplaodApiGroup.SimpleUplader
	uploadRouterGroup := Router.Group("/simpleUploader")
	uploadRouterGroup.Post("/upload", upload_router.SimpleUploaderUpload) // 上传功能
	uploadRouterGroup.Get("/checkFileMd5", upload_router.CheckFileMd5)    // 文件完整度验证
	uploadRouterGroup.Get("/mergeFileMd5", upload_router.MergeFileMd5)    // 合并文件
}

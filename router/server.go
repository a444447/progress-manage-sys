package router

import (
	v1 "progress-manage-system/api/v1"
	"progress-manage-system/model"
	"progress-manage-system/repository"
	"progress-manage-system/service"
)

var (
	db, _            = model.InitDB()
	thesisRepository = repository.NewThesisRepository(db)
	fileRepository   = repository.NewFileRepository(db)
	thesisService    = service.NewThesisService(thesisRepository)
	cosService       = service.NewCosService()
	fileService      = service.NewFileService(cosService, fileRepository)
	thesisController = v1.NewThesisController(thesisService)
	fileController   = v1.NewFileService(fileService)
)

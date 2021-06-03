package service

import (
	"errors"
	"mime/multipart"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model"
	"github.com/snowlyg/go-tenancy/model/request"
	"github.com/snowlyg/go-tenancy/utils/upload"
	"github.com/snowlyg/multi"
)

func Upload(file model.TenancyMedia) (model.TenancyMedia, error) {
	err := g.TENANCY_DB.Create(&file).Error
	return file, err
}

func FindFile(id uint) (model.TenancyMedia, error) {
	var file model.TenancyMedia
	err := g.TENANCY_DB.Where("id = ?", id).First(&file).Error
	return file, err
}

func DeleteFile(file model.TenancyMedia) error {
	fileFromDb, err := FindFile(file.ID)
	if err != nil {
		return err
	}
	oss := upload.NewOss()
	if err = oss.DeleteFile(fileFromDb.Key); err != nil {
		return errors.New("文件删除失败")
	}
	err = g.TENANCY_DB.Where("id = ?", file.ID).Unscoped().Delete(&file).Error
	return err
}

func GetFileRecordInfoList(info request.PageInfo, ctx *gin.Context) (interface{}, int64, error) {
	var total int64
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := g.TENANCY_DB
	if !multi.IsAdmin(ctx) {
		db = db.Where("sys_tenancy_id = ?", multi.GetTenancyId(ctx))
	}
	var fileLists []model.TenancyMedia
	err := db.Find(&fileLists).Count(&total).Error
	if err != nil {
		return fileLists, total, err
	}
	err = db.Limit(limit).Offset(offset).Order("updated_at desc").Find(&fileLists).Error
	return fileLists, total, err
}

func UploadFile(header *multipart.FileHeader, noSave, path string, ctx *gin.Context) (model.TenancyMedia, error) {
	oss := upload.NewOss()
	filePath, key, uploadErr := oss.UploadFile(header)
	if uploadErr != nil {
		panic(uploadErr)
	}
	if path != "" {
		filePath = path + filePath
	}

	var media model.TenancyMedia
	if noSave == "0" {
		s := strings.Split(header.Filename, ".")
		media.Url = filePath
		media.Name = header.Filename
		media.Tag = s[len(s)-1]
		media.Key = key
		media.SysTenancyID = multi.GetTenancyId(ctx)
		return Upload(media)
	}
	return media, nil
}

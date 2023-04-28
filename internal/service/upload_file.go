package service

import (
	"bytes"
	log "elab-backend/logger"
)

func (s *Service) UploadFile(file []byte, filename string, fileType string) bool {
	var fileSuffix string
	if fileType == "video" {
		fileSuffix = ".mp4"
	} else if fileType == "picture" {
		fileSuffix = ".jpg"
	} else {
		log.Logger.Error("无法上传" + fileType + "类型文件")
		return false
	}
	err := s.OssBucket.PutObject("video/"+filename+fileSuffix, bytes.NewReader(file))
	if err != nil {
		log.Logger.Error("上传文件失败" + err.Error())
		return false
	} else {
		return true
	}
}

package util

import (
	"BookMall/config"
	"BookMall/pkg/e"
	"context"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"mime/multipart"
)

func UploadToQiniu(file multipart.File, fileSize int) (int, string) {
	var AccessKey = config.AccessKey
	var Secretkey = config.Secretkey
	var Bucket = config.Bucket
	var ImgUrl = config.QiniuServer
	putPolicy := storage.PutPolicy{
		Scope: Bucket,
	}
	mac := qbox.NewMac(AccessKey, Secretkey)
	upToken := putPolicy.UploadToken(mac)
	cfg := storage.Config{
		Zone:          &storage.ZoneHuanan,
		UseCdnDomains: false,
		UseHTTPS:      false,
	}
	//putExtra用于细化上传功能
	putExtra := &storage.PutExtra{}
	//formUploader是一个表单上传器
	formUploader := storage.NewFormUploader(&cfg)
	//ret用于存储上传后的返回信息
	ret := storage.PutRet{}
	err := formUploader.PutWithoutKey(context.Background(), &ret, upToken, file, int64(fileSize), putExtra)
	if err != nil {
		code := e.ErrorUploadPictureToQiniu
		return code, err.Error()
	}
	url := ret.Key + ImgUrl
	return 200, url
}

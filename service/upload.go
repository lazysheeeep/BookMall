package service

import (
	"BookMall/config"
	"io/ioutil"
	"mime/multipart"
	"os"
	"strconv"
)

func UploadAvatarToLocalStatic(file multipart.File, ID uint, userName string) (filePath string, err error) {
	bID := strconv.Itoa(int(ID))
	basePath := "." + config.AvatarPath + "user" + bID + "/"
	if !PathExistOrNot(basePath) {
		CreateDir(basePath)
	}
	avatarPath := basePath + userName + ".JPG"
	content, err := ioutil.ReadAll(file)
	if err != nil {
		return "", nil
	}
	err = ioutil.WriteFile(avatarPath, content, 0666)
	if err != nil {
		return "", err
	}
	return "product" + bID + "/" + userName + ".JPG", nil
}

func PathExistOrNot(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

func CreateDir(path string) bool {
	err := os.MkdirAll(path, 755)
	if err != nil {
		return false
	}
	return true
}

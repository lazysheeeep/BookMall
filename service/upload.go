package service

import (
	"BookMall/config"
	"io"
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
	content, err := io.ReadAll(file)
	if err != nil {
		return "", nil
	}
	err = os.WriteFile(avatarPath, content, 0666)
	if err != nil {
		return "", err
	}
	return "product" + bID + "/" + userName + ".JPG", nil
}

func UploadBookToLocalStatic(file multipart.File, ID uint, bookName string) (filePath string, err error) {
	bID := strconv.Itoa(int(ID))
	basePath := "." + config.BookPath + "book" + bID + "/"
	if !PathExistOrNot(basePath) {
		CreateDir(basePath)
	}
	bookPath := basePath + bookName + ".JPG"
	content, err := io.ReadAll(file)
	if err != nil {
		return "", nil
	}
	err = os.WriteFile(bookPath, content, 0666)
	if err != nil {
		return "", err
	}
	return "book" + bID + "/" + bookName + ".JPG", nil
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

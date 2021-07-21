package upload

import (
	"errors"
	"go-admin/global"
	"go-admin/utils"
	"io"
	"io/ioutil"
	"mime/multipart"
	"os"
	"path"
	"strings"
	"time"
)

func GetFileName(name string) string {
	ext := GetFileExt(name)
	fileName := strings.TrimSuffix(name, ext)
	fileName = utils.EncodeMD5(fileName)
	return fileName + ext
}

//获取后缀名
func GetFileExt(name string) string {
	return path.Ext(name)
}

//获取配置文件中的上传目录
func GetSavePath() string {
	return global.Upload.UploadPath + "/" + time.Now().Format("20060102")
}

//检测上传路径是否存在
func CheckSavePath(dst string) bool {
	_, err := os.Stat(dst)
	return os.IsNotExist(err)
}

//检测上传文件的后缀名是否在允许上传的后缀名中
func CheckExt(name string) bool {
	ext := GetFileExt(name)
	for _, allow := range global.Upload.AllowExts {
		if strings.ToUpper(allow) == strings.ToUpper(ext) {
			return true
		}
	}
	return false
}

//检测上传文件的大小
func CheckSize(f multipart.File) bool {
	content, _ := ioutil.ReadAll(f)
	size := len(content)
	if size > global.Upload.UploadMaxSize*1024*1024 {
		return true
	}
	return false
}

func CheckPermission(dst string) bool {
	_, err := os.Stat(dst)
	return os.IsPermission(err)
}

func CreateSavePath(dst string, perm os.FileMode) error {
	err := os.MkdirAll(dst, perm)
	if err != nil {
		return err
	}
	return nil
}

func SaveFile(file *multipart.FileHeader, dst string) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()
	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()
	_, err1 := io.Copy(out, src)
	return err1
}

func UploadFile(file multipart.File, fileHeader *multipart.FileHeader) (string, error) {
	fileName := GetFileName(fileHeader.Filename)
	uploadPath := GetSavePath()
	dst := uploadPath + "/" + fileName
	if !CheckExt(fileName) {
		return "", errors.New("该文件类型不可上传")
	}
	if CheckSavePath(uploadPath) {
		if err := CreateSavePath(uploadPath, os.ModePerm); err != nil {
			return "", errors.New("创建文件保存目录失败")
		}
	}
	if CheckSize(file) {
		return "", errors.New("文件大小超出限制")
	}
	if CheckPermission(uploadPath) {
		return "", errors.New("上传的文件没有权限")
	}
	if err := SaveFile(fileHeader, dst); err != nil {
		return "", err
	}
	return global.Upload.Url + "/" + time.Now().Format("20060102") + "/" + fileName, nil
}

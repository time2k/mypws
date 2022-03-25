package mylibs

import (
	"fmt"
	"image"
	"io/ioutil"
	"log"
	"mime/multipart"
	myconfig "mypws/config"
	"os"
	"strings"
)

//CreateDir 调用os.MkdirAll递归创建文件夹
func CreateDir(filePath string) error {
	if !isExist(filePath) {
		err := os.MkdirAll(filePath, os.ModePerm)
		return err
	}
	return nil
}

// 判断所给路径文件/文件夹是否存在(返回true是存在)
func isExist(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

//UploadPicFile 处理上传逻辑
func UploadPicFile(file *multipart.FileHeader, maxwidth int, maxheight int, fileindex string) (string, error) {
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	imgsrc, filetype, err := image.Decode(src)
	if err != nil {
		log.Printf("image decode error.filetype %s, error: %s", filetype, err.Error())
		return "", err
	}

	// Get height and width
	height := imgsrc.Bounds().Dy()
	width := imgsrc.Bounds().Dx()

	if height > maxheight || width > maxwidth {
		return "", fmt.Errorf("image width or height oversize")
	}

	//文件指针回0
	src.Seek(0, 0)

	srcbt, err := ioutil.ReadAll(src)
	if err != nil {
		return "", fmt.Errorf("read error: %s", err.Error())
	}

	filex := strings.Split(file.Filename, ".")
	filesuffix := filex[len(filex)-1]
	newfilename := "/" + fileindex + "." + filesuffix

	err = ioutil.WriteFile(myconfig.FILE_UPLOAD_DIR+newfilename, srcbt, 0644)
	if err != nil {
		return "", fmt.Errorf("write error: %s", err.Error())
	}

	return newfilename, nil
}

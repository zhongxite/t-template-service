package utils

import (
	"errors"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"log"
	"mime/multipart"
	"path/filepath"
	"time"
)

func UploadOss(file *multipart.FileHeader) (string, error) {
	fileExt := filepath.Ext(file.Filename) // 获取后缀名
	allowExts := []string{".jpg", ".png", ".gif", ".jpeg", ".doc", ".docx", ".ppt", ".pptx", ".xls", ".xlsx", ".pdf"}
	allowFlag := false
	log.Println(1)
	for _, ext := range allowExts {
		if ext == fileExt {
			allowFlag = true
			break
		}
	}
	if !allowFlag {
		return "", errors.New("上传的格式有误")
	}
	now := time.Now()
	fileDir := fmt.Sprintf("test/%s", now.Format("20060102")) //文件存放路径
	timeStamp := now.Unix()
	fileName := fmt.Sprintf("%d-%s", timeStamp, file.Filename) //文件名称
	fileKey := filepath.Join(fileDir, fileName)                //拼接文件名
	src, err := file.Open()
	defer src.Close()
	if err != nil {
		return "", errors.New("上传失败")
	}

	client, err := oss.New("oss-cn-beijing.aliyuncs.com", "LTAI5tQP5qwtHTSEqWHf9YSX", "LrmAWEjoTTlYpMpX0RUqrTGaiXYJnz")
	if err != nil {
		return "", errors.New("上传失败")
	}
	bucketName := "use-demo"
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		return "", errors.New("上传失败")
	}
	err = bucket.PutObject(fileKey, src)
	if err != nil {
		return "", errors.New("上传失败")
	}
	url := fmt.Sprintf("https://%s.%s/%s", bucketName, "oss-cn-beijing.aliyuncs.com", fileKey)
	return url, nil
}

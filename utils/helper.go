package utils

import (
	"errors"
	"fmt"
	"mime/multipart"
	"time"

	storage_go "github.com/supabase-community/storage-go"
)

const (
	MaxFileSize = 5 * 1024 * 1024
)

func UploadImage(fileInput *multipart.FileHeader) (string, error) {
	if fileInput.Size > 6*1024*1024 {
		return "", errors.New("file size is too big")
	}
	file, err := fileInput.Open()
	if err != nil {
		return "", err
	}
	defer file.Close()
	fileName := fmt.Sprintf("data%s%s", time.Now().Format("20060102150405"), fileInput.Filename)
	client := storage_go.NewClient("https://jgjyjvyldoamqndazixl.supabase.co/storage/v1", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJzdXBhYmFzZSIsInJlZiI6ImpnanlqdnlsZG9hbXFuZGF6aXhsIiwicm9sZSI6ImFub24iLCJpYXQiOjE2NDc4MzQ0MDQsImV4cCI6MTk2MzQxMDQwNH0.WVMjJIRoK_cnyfRXdYvTokNWBCCqLWfbeu7xXeZrs6I", nil)
	client.UploadFile("foto-proker", fileName, file)
	linkImage := GenerateLinkImage(fileName)
	return linkImage, nil
}

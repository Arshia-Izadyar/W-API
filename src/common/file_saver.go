package common

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"strings"

	"github.com/google/uuid"
)

func SaveFile(file *multipart.FileHeader, directory string) (string, error) {
	randFilename := uuid.New()
	err := os.MkdirAll(directory, os.ModePerm)
	if err != nil {
		return "", err
	}
	fileNameArr := strings.Split(file.Filename, ".")
	ext := fileNameArr[len(fileNameArr)-1]
	filename := fmt.Sprintf("%s.%s", randFilename, ext)
	dest := fmt.Sprintf("%s/%s", directory, filename)

	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	out, err := os.Create(dest)
	if err != nil {
		return "", err
	}
	defer out.Close()
	_, err = io.Copy(out, src)
	if err != nil {
		return "", err
	}
	return filename, nil
}

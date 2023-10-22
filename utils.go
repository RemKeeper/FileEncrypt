package main

import (
	"fyne.io/fyne/v2/dialog"
	"os"
	"path/filepath"
	"strings"
)

func ReadFIle(FileURI string) (FileByte []byte, FileName string, err error) {
	file, err := os.ReadFile(FileURI)
	if err != nil {
		return nil, "", err
	}
	return file, filepath.Base(FileURI), nil
}

func WriteFile(WritePath string, FileByte []byte) error {
	err := os.WriteFile(WritePath, FileByte, 0644)
	if err != nil {
		return err
	}
	return nil
}

func RemoveEncryptionSuffix(FileURI string) string {
	SuffixIndex := strings.LastIndex(FileURI, ".")
	return FileURI[:SuffixIndex]
}

func RemoveFileName(FileURI string) string {
	SuffixIndex := strings.LastIndex(FileURI, "/")
	return FileURI[:SuffixIndex+1]
}

func DisplayErrorDialog(err error, Message string) {
	dialog.NewCustom("错误", Message+" "+err.Error(), nil, Windows).Show()
}

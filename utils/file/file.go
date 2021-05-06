package file

import (
	"io/ioutil"
	"mime/multipart"
	"os"
)

func ReadFile(filepath string) ([]byte, error) {
	content, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	return content, nil
}

func ReadAll(file multipart.File) ([]byte, error) {
	content, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	return content, nil
}

func Write(filePath string, data []byte) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(file.Name(), data, 0664)
	if err != nil {
		return err
	}
	return nil
}
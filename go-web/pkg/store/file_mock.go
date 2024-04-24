package store

import (
	"encoding/json"
	"os"
)

type FileStoreMock struct {
	FileName string
	Mock     *Mock
}

type Mock struct {
	Data          []byte
	Err           error
	HasCalledRead bool
}

func (fs *FileStoreMock) AddMock(mock *Mock) {
	fs.Mock = mock
}
func (fs *FileStoreMock) ClearMock() {
	fs.Mock = nil
}

func (fs *FileStoreMock) Write(data interface{}) error {
	if fs.Mock != nil {
		if fs.Mock.Err != nil {
			return fs.Mock.Err
		}
		return nil
	}
	fileData, err := json.MarshalIndent(data, "", "")
	if err != nil {
		return err
	}
	return os.WriteFile(fs.FileName, fileData, 0644)
}

func (fs *FileStoreMock) Read(data interface{}) error {
	fs.Mock.HasCalledRead = true
	if fs.Mock != nil {
		if fs.Mock.Err != nil {
			return fs.Mock.Err
		}
		return json.Unmarshal(fs.Mock.Data, data)
	}
	file, err := os.ReadFile(fs.FileName)
	if err != nil {
		return err
	}
	return json.Unmarshal(file, data)
}

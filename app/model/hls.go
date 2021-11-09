package model

import (
	"bytes"
	"io"
	"ms-stream-api/app/util"
	"os"
)

type HlsFile struct {
	File io.ReadCloser
}

func NewHlsFile(file *os.File) HlsFile {
	return HlsFile{
		File: file,
	}
}

func HlsFileGet(videoId int, segName string) (*HlsFile, error) {
	mediaFilePath := util.FileUtilHlsFilePath(videoId, segName)
	file, err := os.Open(mediaFilePath)
	if err != nil {
		return nil, err
	}
	data := NewHlsFile(file)
	return &data, nil
}

func (t HlsFile) Bytes() []byte {
	buffer := new(bytes.Buffer)
	io.Copy(buffer, t.File)
	return buffer.Bytes()
}

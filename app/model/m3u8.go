package model

import (
	"bytes"
	"io"
	"ms-stream-api/app/util"
	"os"
)

type M3u8File struct {
	File io.ReadCloser
}

func NewM3u8File(file *os.File) M3u8File {
	return M3u8File{
		File: file,
	}
}

func M3u8FileGet(videoId int) (*M3u8File, error) {
	mediaFilePath := util.FileUtilM3u8FilePath(videoId)
	file, err := os.Open(mediaFilePath)
	if err != nil {
		return nil, err
	}
	data := NewM3u8File(file)
	return &data, nil
}

func (t M3u8File) Bytes() []byte {
	buffer := new(bytes.Buffer)
	io.Copy(buffer, t.File)
	return buffer.Bytes()
}

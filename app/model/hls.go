package model

import (
	"bytes"
	"io"
	"ms-stream-api/app/util"
	"os"
)

type HlsFile struct {
	Closer io.ReadCloser
}

func NewHlsFile(f *os.File) HlsFile {
	return HlsFile{
		Closer: f,
	}
}

func HlsFileGet(videoId int, segName string) (*HlsFile, error) {
	mediaFilePath := util.FileUtil.HlsFilePath(videoId, segName)
	f, err := os.Open(mediaFilePath)
	if err != nil {
		return nil, err
	}
	data := NewHlsFile(f)
	return &data, nil
}

func (t HlsFile) Bytes() []byte {
	buffer := new(bytes.Buffer)
	io.Copy(buffer, t.Closer)
	return buffer.Bytes()
}

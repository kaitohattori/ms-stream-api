package model

import (
	"bytes"
	"io"
	"ms-stream-api/app/util"
	"os"
)

type M3u8File struct {
	Closer io.ReadCloser
}

func NewM3u8File(f *os.File) M3u8File {
	return M3u8File{
		Closer: f,
	}
}

func M3u8FileGet(videoId int) (*M3u8File, error) {
	mediaFilePath := util.FileUtilM3u8FilePath(videoId)
	f, err := os.Open(mediaFilePath)
	if err != nil {
		return nil, err
	}
	data := NewM3u8File(f)
	return &data, nil
}

func (t M3u8File) Bytes() []byte {
	buffer := new(bytes.Buffer)
	io.Copy(buffer, t.Closer)
	return buffer.Bytes()
}

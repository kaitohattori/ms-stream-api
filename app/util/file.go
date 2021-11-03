package util

import (
	"fmt"
)

var FileUtil FileUtilFuncs

type FileUtilFuncs struct {
	MediaRoot    string
	M3u8FileName string
}

func init() {
	FileUtil = FileUtilFuncs{
		MediaRoot:    "assets/media",
		M3u8FileName: "video.m3u8",
	}
}

func (u FileUtilFuncs) M3u8FilePath(videoId int) string {
	basePath := fmt.Sprintf("%s/%d", u.MediaRoot, videoId)
	return fmt.Sprintf("%s/hls/%s", basePath, u.M3u8FileName)
}

func (u FileUtilFuncs) HlsFilePath(videoId int, segName string) string {
	basePath := fmt.Sprintf("%s/%d", u.MediaRoot, videoId)
	return fmt.Sprintf("%s/hls/%s", basePath, segName)
}

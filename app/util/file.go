package util

import (
	"fmt"
	"ms-stream-api/config"
)

var FileUtil FileUtilFuncs

type FileUtilFuncs struct {
}

func init() {
	FileUtil = FileUtilFuncs{}
}

func FileUtilM3u8FilePath(videoId int) string {
	basePath := fmt.Sprintf("%s/%d", config.Config.AssetsDirPath, videoId)
	return fmt.Sprintf("%s/hls/%s", basePath, config.Config.AssetsM3u8FileName)
}

func FileUtilHlsFilePath(videoId int, segName string) string {
	basePath := fmt.Sprintf("%s/%d", config.Config.AssetsDirPath, videoId)
	return fmt.Sprintf("%s/hls/%s", basePath, segName)
}

package controller

import (
	"log"
	"ms-stream-api/app/model"
	"ms-stream-api/app/util"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type StreamController struct {
}

func NewStreamController() *StreamController {
	return &StreamController{}
}

// StreamController Stream docs
// @Summary Stream video
// @Description return M3u8 file
// @Tags Stream
// @Produce application/x-mpegURL
// @Success 200 {object} string ""
// @Failure 400 {object} util.HTTPError
// @Failure 404 {object} util.HTTPError
// @Failure 500 {object} util.HTTPError
// @Router /video/{id}/stream [get]
func (c *StreamController) GetM3u8File(ctx *gin.Context) {
	videoIdStr := ctx.Param("id")
	videoId, err := strconv.Atoi(videoIdStr)
	if err != nil {
		log.Println(err)
		util.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	m3u8File, err := model.M3u8FileGet(videoId)
	if err != nil {
		log.Println(err)
		util.NewError(ctx, http.StatusNotFound, err)
		return
	}
	ctx.Data(http.StatusOK, "application/x-mpegURL", m3u8File.Bytes())
}

// StreamController Stream docs
// @Summary Get hls file
// @Description return hls file
// @Tags Stream
// @Produce video/MP2T
// @Success 200 {object} string ""
// @Failure 400 {object} util.HTTPError
// @Failure 404 {object} util.HTTPError
// @Failure 500 {object} util.HTTPError
// @Router /video/{id}/stream/{segName} [get]
func (c *StreamController) GetHlsFile(ctx *gin.Context) {
	segName := ctx.Param("segName")
	videoIdStr := ctx.Param("id")
	videoId, err := strconv.Atoi(videoIdStr)
	if err != nil {
		log.Println(err)
		util.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	hlsFile, err := model.HlsFileGet(videoId, segName)
	if err != nil {
		log.Println(err)
		util.NewError(ctx, http.StatusNotFound, err)
		return
	}
	ctx.Data(http.StatusOK, "video/MP2T", hlsFile.Bytes())
}

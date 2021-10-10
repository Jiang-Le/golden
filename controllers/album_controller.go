package controllers

import (
	"time"
	
	"github.com/gin-gonic/gin"
	"github.com/jiel/golden/services"
)

func NewAlbumController(service services.AlbumService) *AlbumController {
	return &AlbumController{
		AlbumService: service,
	}
}

type AlbumController struct {
	BaseController
	AlbumService services.AlbumService
}

type GetAlbumsRequest struct {
	Offset string   `json:"offset"`
	Size   int32    `json:"size"`
	Tags   []string `json:"tags"`
}

type AlbumResponse struct {
	ID        string   `json:"id"`
	Name      string   `json:"name"`
	Tags      []string `json:"tags"`
	Cover     string   `json:"cover"`
	Origin    string   `json:"origin"`
	Timestamp int64    `json:"timestamp"`
	Likes     int64    `json:"likes"`
	Unlikes   int64    `json:"unlikes"`
	Pics      []string `json:"pics"`
}

func (c *AlbumController) GetAlbums(ctx *gin.Context) {
	var request GetAlbumsRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		c.serveError(ctx, ErrInvalidRequest)
		return
	}
	albums, err := c.AlbumService.List(request.Offset, request.Size, request.Tags)
	if err != nil {
		c.serveError(ctx, ErrAlbumsQueryFail)
		return
	}
	albumsResponse := make([]*AlbumResponse, 0, len(albums))
	for _, album := range albums {
		albumsResponse = append(albumsResponse, &AlbumResponse{
			ID:        album.ID,
			Name:      album.Name,
			Tags:      album.Tags,
			Cover:     album.Cover,
			Origin:    album.Origin,
			Timestamp: album.Timestamp,
			Likes:     album.Likes,
			Unlikes:   album.Unlikes,
			Pics:      album.Pics,
		})
	}
	c.serveResponse(ctx, albumsResponse)
}

type CreateAlbumRequest struct {
	Name   string   `json:"name"`
	Tags   []string `json:"tags"`
	Cover  string   `json:"cover"`
	Origin string   `json:"origin"`
	Pics   []string `json:"pics"`
}

func (c *AlbumController) CreateAlbum(ctx *gin.Context) {
	var request CreateAlbumRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		c.serveError(ctx, ErrInvalidRequest)
		return
	}
	album := services.Album{
		Name:   request.Name,
		Tags:   request.Tags,
		Origin: request.Origin,
		Pics:   request.Pics,
	}
	if len(request.Cover) == 0 {
		album.Cover = album.Pics[0]
	}
	album.Timestamp = time.Now().Unix()
	if err := c.AlbumService.Create(&album); err != nil {
		c.serveError(ctx, ErrAlbumCreateFail)
		return
	}
	c.serveResponse(ctx, nil, "创建成功")
}

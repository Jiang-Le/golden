package dao

import (
	"context"

	"github.com/jiel/golden/models"
	"github.com/pkg/errors"
)

var (
	ErrNoAlbum        = errors.New("no album")
	ErrAlbumNotExist  = errors.New("album not exist")
	ErrInvalidAlbumID = errors.New("invalid album id")
)

type AlbumListQueryParam struct {
	Offset string
	Size   int32
	Tags   []string
}

type AlbumRepo interface {
	CreateAlbum(ctx context.Context, album models.Album) (string, error)
	CreateAlbums(ctx context.Context, albums []models.Album) ([]string, error)
	DeleteAlbum(ctx context.Context, albumID string) error
	DeleteAlbums(ctx context.Context, albumIDs []string) error
	ListRecentlyAlbum(ctx context.Context, param AlbumListQueryParam) ([]models.Album, error)
}

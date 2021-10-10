package impl

import (
	"context"
	"time"

	"github.com/jiel/golden/dao"
	"github.com/jiel/golden/models"
	"github.com/jiel/golden/services"
	"github.com/pkg/errors"
)

func NewAlbumService(albumRepo dao.AlbumRepo) services.AlbumService {
	return &albumServiceImpl{
		AlbumRepo: albumRepo,
	}
}

type albumServiceImpl struct {
	AlbumRepo dao.AlbumRepo
}

func (c *albumServiceImpl) List(offset string, size int32, tags []string) ([]*services.Album, error) {
	param := dao.AlbumListQueryParam{
		Offset: offset,
		Size:   size,
		Tags:   tags,
	}
	ctx := context.Background()
	rets, err := c.AlbumRepo.ListRecentlyAlbum(ctx, param)
	if err != nil {
		if err == dao.ErrNoAlbum {
			return []*services.Album{}, nil
		}
		return nil, errors.WithStack(err)
	}
	albums := make([]*services.Album, 0, len(rets))
	for _, item := range rets {
		albums = append(albums, &services.Album{
			ID:        item.ID.Hex(),
			Name:      item.Name,
			Tags:      item.Tags,
			Cover:     item.Cover,
			Origin:    item.Origin,
			Timestamp: item.Timestamp,
			Likes:     item.Likes,
			Unlikes:   item.Unlikes,
			Pics:      item.Pics,
		})
	}
	return albums, nil
}

func (c *albumServiceImpl) Create(album *services.Album) error {
	ctx := context.Background()
	albumModel := models.Album{
		Name:      album.Name,
		Tags:      album.Tags,
		Cover:     album.Cover,
		Origin:    album.Origin,
		Timestamp: album.Timestamp,
		Likes:     0,
		Unlikes:   0,
		Pics:      album.Pics,
	}
	if _, err := c.AlbumRepo.CreateAlbum(ctx, albumModel); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (c *albumServiceImpl) BatchCreate(albums []*services.Album) error {
	ctx := context.Background()
	albumModels := make([]models.Album, 0, len(albums))
	for _, album := range albums {
		albumModels = append(albumModels, models.Album{
			Name:      album.Name,
			Tags:      album.Tags,
			Cover:     album.Cover,
			Origin:    album.Origin,
			Timestamp: time.Now().Unix(),
			Likes:     0,
			Unlikes:   0,
			Pics:      album.Pics,
		})
	}
	if _, err := c.AlbumRepo.CreateAlbums(ctx, albumModels); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (c *albumServiceImpl) DeleteAlbum(albumID string) error {
	ctx := context.Background()
	if err := c.AlbumRepo.DeleteAlbum(ctx, albumID); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (c *albumServiceImpl) DeleteAlbums(albumIDs []string) error {
	ctx := context.Background()
	if err := c.AlbumRepo.DeleteAlbums(ctx, albumIDs); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

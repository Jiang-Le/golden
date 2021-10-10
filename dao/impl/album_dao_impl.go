package impl
//
//import (
//	"context"
//	"github.com/jiel/golden/dao"
//	"github.com/jiel/golden/models"
//	"github.com/pkg/errors"
//	"gorm.io/gorm"
//	"gorm.io/gorm/clause"
//)
//
//func NewAlbumRepo() dao.AlbumRepo {
//
//}
//
//type albumRepo struct {
//	db *gorm.DB
//}
//
//func (r *albumRepo) CreateAlbum(ctx context.Context, album models.Album) (uint, error) {
//	err := r.db.Transaction(func(tx *gorm.DB) error {
//		if err := r.db.Clauses(clause.OnConflict{Columns: []clause.Column{{Name: "Name"}}, DoNothing: true}).
//			Create(&album.Tags).Error; err != nil {
//			return err
//		}
//		if err := r.db.Omit("Tags").Create(&album).Error; err != nil {
//			return err
//		}
//		if err := r.db.Model(&album).Association("Tags").Append(album.Tags); err != nil {
//			return err
//		}
//		return nil
//	})
//	if err != nil {
//		return 0, errors.WithStack(err)
//	}
//	return album.ID, nil
//}
//
//func (r *albumRepo) CreateAlbums(ctx context.Context, albums []models.Album) ([]int64, error) {
//	r.db.Transaction(func(tx *gorm.DB) error {
//		if err := r.db.Omit("Tags").Create(&albums).Error; err != nil {
//			return err
//		}
//		allTags := func() []models.Tag {
//			ret := make([]models.Tag, 0)
//			for _, album := range albums {
//				ret = append(ret, album.Tags...)
//			}
//			return ret
//		}()
//
//
//	})
//}
//
//func (r *albumRepo) DeleteAlbum(ctx context.Context, albumID uint) error {
//	panic("implement me")
//}
//
//func (r *albumRepo) DeleteAlbums(ctx context.Context, albumIDs []uint) error {
//	panic("implement me")
//}
//
//func (r *albumRepo) ListRecentlyAlbum(ctx context.Context, param dao.AlbumListQueryParam) ([]models.Album, error) {
//	panic("implement me")
//}

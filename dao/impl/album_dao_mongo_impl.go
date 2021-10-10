package impl

import (
	"context"

	"github.com/jiel/golden/dao"
	"github.com/jiel/golden/models"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongoAlbumRepo(client *mongo.Client, dbName, colName string) dao.AlbumRepo {
	return &albumMongoRepo{
		db: client.Database(dbName).Collection(colName),
	}
}

type albumMongoRepo struct {
	db *mongo.Collection
}

func (r *albumMongoRepo) CreateAlbum(ctx context.Context, album models.Album) (string, error) {
	ret, err := r.db.InsertOne(ctx, &album)
	if err != nil {
		return "", errors.WithStack(err)
	}
	id := ret.InsertedID.(primitive.ObjectID).Hex()
	return id, nil
}

func (r *albumMongoRepo) CreateAlbums(ctx context.Context, albums []models.Album) ([]string, error) {
	values := make([]interface{}, 0, len(albums))
	for album := range albums {
		values = append(values, album)
	}
	ret, err := r.db.InsertMany(ctx, values)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	ids := make([]string, 0, len(ret.InsertedIDs))
	for _, id := range ret.InsertedIDs {
		ids = append(ids, id.(primitive.ObjectID).Hex())
	}
	return ids, nil
}

func (r *albumMongoRepo) DeleteAlbum(ctx context.Context, albumID string) error {
	id, err := primitive.ObjectIDFromHex(albumID)
	if err != nil {
		return dao.ErrInvalidAlbumID
	}
	filter := bson.M{
		"_id": id,
	}
	ret, err := r.db.DeleteOne(ctx, filter)
	if err != nil {
		return errors.WithStack(err)
	}
	if ret.DeletedCount == 0 {
		return dao.ErrAlbumNotExist
	}
	return nil
}

func (r *albumMongoRepo) DeleteAlbums(ctx context.Context, albumIDs []string) error {
	ids := make([]primitive.ObjectID, 0, len(albumIDs))
	for _, albumID := range albumIDs {
		id, err := primitive.ObjectIDFromHex(albumID)
		if err != nil {
			return dao.ErrInvalidAlbumID
		}
		ids = append(ids, id)
	}
	filter := bson.M{
		"_id": bson.M{
			"$in": ids,
		},
	}
	if _, err := r.db.DeleteMany(ctx, filter); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (r *albumMongoRepo) ListRecentlyAlbum(ctx context.Context, param dao.AlbumListQueryParam) ([]models.Album, error) {
	opt := &options.FindOptions{
		Sort: bson.M{
			"_id": -1,
		},
	}
	opt.SetLimit(int64(param.Size))
	filter := bson.M{}
	if len(param.Offset) > 0 {
		offset, err := primitive.ObjectIDFromHex(param.Offset)
		if err != nil {
			return nil, dao.ErrInvalidAlbumID
		}
		listAlbumFilterWithOffset(filter, offset)
	}
	if len(param.Tags) > 0 {
		listAlbumFilterWithTags(filter, param.Tags)
	}
	cursor, err := r.db.Find(ctx, filter, opt)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, dao.ErrNoAlbum
		}
		return nil, errors.WithStack(err)
	}
	defer cursor.Close(ctx)
	var ret []models.Album
	for cursor.Next(ctx) {
		var album models.Album
		if err := cursor.Decode(&album); err != nil {
			return nil, errors.WithStack(err)
		}
		ret = append(ret, album)
	}
	return ret, nil
}

func listAlbumFilterWithOffset(filter bson.M, offset primitive.ObjectID) bson.M {
	filter["_id"] = bson.M{
		"$lt": offset,
	}
	return filter
}

func listAlbumFilterWithTags(filter bson.M, tags []string) bson.M {
	filter["tags"] = bson.M{
		"$in": tags,
	}
	return filter
}

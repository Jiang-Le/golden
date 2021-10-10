package services

type AlbumService interface {
	List(offset string, size int32, tags []string) ([]*Album, error)
	Create(album *Album) error
	BatchCreate(albums []*Album) error
	DeleteAlbum(albumID string) error
	DeleteAlbums(albumIDs []string) error
}

type Album struct {
	ID        string
	Name      string
	Tags      []string
	Cover     string
	Origin    string
	Timestamp int64
	Likes     int64
	Unlikes   int64
	Pics      []string
}

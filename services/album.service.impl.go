package services

import (
	"api/ETOM/albums/models"
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type AlbumServiceImpl struct {
	albumCollection *mongo.Collection
	ctx context.Context
}

func NewAlbumService(albumCollection *mongo.Collection, ctx context.Context) AlbumService {
 return &AlbumServiceImpl {
	albumCollection: albumCollection,
	ctx: ctx,
 }
}

func (a *AlbumServiceImpl) CreateAlbum(album *models.Album) error{
	_, err := a.albumCollection.InsertOne(a.ctx, album)
	return err
}

func (a *AlbumServiceImpl) GetAlbum(name *string) (*models.Album, error){
	var album *models.Album
	query := bson.D{bson.E{Key:"album_name", Value:name}}
	err:= a.albumCollection.FindOne(a.ctx, query).Decode(&album)
	return album, err
} 

func (a *AlbumServiceImpl) GetAll()( []*models.Album, error){
	var albums []*models.Album
	cursor , err := a.albumCollection.Find(a.ctx, bson.D{})
	if err != nil{
		return nil, err
	}
	for cursor.Next(a.ctx) {
		var album models.Album
		err := cursor.Decode(&album)
		if err != nil{
			return 	nil , err
		}
		albums = append(albums, &album)
	}
	if err:=cursor.Err(); err!= nil{
		return nil, err
	}
	cursor.Close(a.ctx)
	if len(albums) == 0 {
		return nil , errors.New("empty, nothing was found")
	}
	return albums, nil
} 

func (a *AlbumServiceImpl) UpdateAlbum(album *models.Album) error {
	filterQuery := bson.D{bson.E{Key:"album_name", Value:album.Name}}
	updateQuery := bson.D{bson.E{Key:"$set", Value:bson.D{bson.E{Key: "album_name", Value: album.Name}, bson.E{Key:"user_age", Value: album.Age}, bson.E{Key:"user_adress", Value: album.Adress}}}}
	result, _:= a.albumCollection.UpdateOne(a.ctx, filterQuery, updateQuery)
	if result.MatchedCount != 1 {
		return errors.New("no match Album find for update")
	}
	return nil
} 

func (a *AlbumServiceImpl) DeleteAlbum(name *string) error {
	filterQuery := bson.D{bson.E{Key:"album_name", Value:name}}
	result, _ := a.albumCollection.DeleteOne(a.ctx, filterQuery)
	if result.DeletedCount != 1 {
		return errors.New("no match Album find for update")
	}
	return nil
} 
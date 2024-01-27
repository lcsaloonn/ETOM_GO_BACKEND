package services

import (
	"api/ETOM/albums/models"
	"context"

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

func (u *AlbumServiceImpl) CreateAlbum(album *models.Album) error{
	return nil
}

func (u *AlbumServiceImpl) GetAlbum(name *string) (*models.Album, error){
	return nil, nil
} 

func (u *AlbumServiceImpl) GetAll()( []*models.Album, error){
	return nil, nil
} 

func (u *AlbumServiceImpl) UpdateAlbum(album *models.Album) error {
	return nil
} 

func (u *AlbumServiceImpl) DeleteAlbum(name *string) error {
	return nil
} 
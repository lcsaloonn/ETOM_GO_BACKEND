package services

import "api/ETOM/albums/models"

type AlbumService interface {
	CreateAlbum(*models.Album) error
	GetAlbum(*string) (*models.Album, error)
	GetAll() ([]*models.Album, error)
	UpdateAlbum(*models.Album) error
	DeleteAlbum(*string) error
}
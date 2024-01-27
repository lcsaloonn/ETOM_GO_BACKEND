package controllers

import (
	"api/ETOM/albums/services"

	"github.com/gin-gonic/gin"
)

type AlbumController struct {
	AlbumService services.AlbumService
}

func New(albumService services.AlbumService) AlbumController{
	return AlbumController{
		AlbumService: albumService,
	}	
}

func (c *AlbumController) CreateAlbum(ctx *gin.Context) {
	 ctx.JSON(200, "")
}	

func (c *AlbumController) GetAlbum(ctx *gin.Context) {
	ctx.JSON(200, "")
} 

func (c *AlbumController) GetAll(ctx *gin.Context){
	ctx.JSON(200, "")
} 

func (c *AlbumController) UpdateAlbum(ctx *gin.Context)  {
	ctx.JSON(200, "")
} 

func (c *AlbumController) DeleteAlbum(ctx *gin.Context)  {
	ctx.JSON(200, "")
} 

func (c *AlbumController) RegisterAlbumRoutes(rg *gin.RouterGroup){
	userRoute := rg.Group("/album")
	userRoute.POST("/create", c.CreateAlbum)
	userRoute.GET("/get/:name", c.GetAlbum)
	userRoute.GET("/get", c.GetAlbum)
	userRoute.DELETE("/delete/:name", c.DeleteAlbum)
	userRoute.PATCH("/update", c.UpdateAlbum)
}
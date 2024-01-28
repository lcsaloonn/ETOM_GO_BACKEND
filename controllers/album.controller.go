package controllers

import (
	"api/ETOM/albums/models"
	"api/ETOM/albums/services"
	"net/http"

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
	var album models.Album
	if err := ctx.ShouldBindJSON(&album); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message":err.Error()})
	}
	err:=c.AlbumService.CreateAlbum(&album)
	if err != nil{
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	 ctx.JSON(http.StatusCreated, gin.H{"message":"success"})
}	

func (c *AlbumController) GetAlbum(ctx *gin.Context) {
	albumName:=ctx.Param("name")
	album, err := c.AlbumService.GetAlbum(&albumName)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message":err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, album)
} 

func (c *AlbumController) GetAll(ctx *gin.Context){
	albums, err:=c.AlbumService.GetAll()
	if err != nil{
		ctx.JSON(http.StatusBadGateway, gin.H{"message":err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, albums)
} 

func (c *AlbumController) UpdateAlbum(ctx *gin.Context)  {
	var album models.Album
	if err:= ctx.ShouldBindJSON(&album); err != nil{
		ctx.JSON(http.StatusBadRequest , gin.H{"message": err.Error()})
		return
	}
	err:=c.AlbumService.UpdateAlbum(&album)
	if err != nil{
		ctx.JSON(http.StatusBadGateway, gin.H{"message":err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message":"success"})
} 

func (c *AlbumController) DeleteAlbum(ctx *gin.Context)  {
	userName := ctx.Param("name")
	err:= c.AlbumService.DeleteAlbum(&userName)
	if err != nil{
		ctx.JSON(http.StatusBadGateway, gin.H{"message":err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message":"success"})
} 

func (c *AlbumController) RegisterAlbumRoutes(rg *gin.RouterGroup){
	userRoute := rg.Group("/album")
	userRoute.POST("/create", c.CreateAlbum)
	userRoute.GET("/get/:name", c.GetAlbum)
	userRoute.GET("/get", c.GetAlbum)
	userRoute.DELETE("/delete/:name", c.DeleteAlbum)
	userRoute.PATCH("/update", c.UpdateAlbum)
}
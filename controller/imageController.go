package controller

import (
	"bytes"
	"fmt"
	"github.com/TiagoNora/GoCRUDV2/config/minioClient"
	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
	"github.com/rs/zerolog/log"
)

type ImageController interface {
	CreateImage(c *gin.Context)
	GetImage(c *gin.Context)
}

type imageController struct {
	minio *minio.Client
}

// @Summary Upload an image to MinIO
// @Description Upload an image to MinIO bucket "imagens"
// @Tags Image
// @Accept multipart/form-data
// @Param file formData file true "File to upload"
// @Success 200 {object} map[string]string "Success message with file name"
// @Failure 400 {object} map[string]string "Error message"
// @Router /image [post]
func (i imageController) CreateImage(c *gin.Context) {
	file, fileHeader, err := c.Request.FormFile("file")
	if err != nil {
		log.Error().Err(err).Msg("Erro ao obter o arquivo")
		return
	}
	defer file.Close()

	buffer := make([]byte, fileHeader.Size)
	_, err = file.Read(buffer)
	if err != nil {
		log.Error().Err(err).Msg("Erro ao obter o arquivo")
		return
	}

	bucketName := "imagens"
	objectName := fileHeader.Filename
	contentType := fileHeader.Header.Get("Content-Type")

	err = i.minio.MakeBucket(c, bucketName, minio.MakeBucketOptions{})
	if err != nil {
		exists, errBucketExists := i.minio.BucketExists(c, bucketName)
		if errBucketExists != nil || !exists {
			log.Error().Err(err).Msg("Erro ao obter o arquivo")
			return
		}
	}

	reader := bytes.NewReader(buffer)
	_, err = i.minio.PutObject(c, bucketName, objectName, reader, fileHeader.Size, minio.PutObjectOptions{
		ContentType: contentType,
	})
	if err != nil {
		log.Error().Err(err).Msg("Erro ao obter o arquivo")
		return
	}

	sendSuccess(c, "Upload", objectName)
}

// @Summary Get an image from MinIO
// @Description Get an image from MinIO bucket "imagens"
// @Tags Image
// @Param fileName path string true "File name to retrieve"
// @Success 200 {file} file "Image file"
// @Failure 400 {object} map[string]string "Error message"
// @Router /image/{fileName} [get]
func (i imageController) GetImage(c *gin.Context) {
	fileName := c.Param("fileName")
	if fileName == "" {
		log.Error().Msg("Erro ao obter o arquivo")
		return
	}

	bucketName := "imagens"

	object, err := i.minio.GetObject(c, bucketName, fileName, minio.GetObjectOptions{})
	if err != nil {
		log.Error().Err(err).Msg("Erro ao obter o arquivo")
		return
	}
	defer object.Close()

	stat, err := object.Stat()
	if err != nil {
		log.Error().Err(err).Msg("Erro ao obter o arquivo")
		return
	}

	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Disposition", "inline; filename="+fileName)
	c.Header("Content-Type", stat.ContentType)
	c.Header("Content-Length", fmt.Sprintf("%d", stat.Size))

	c.DataFromReader(200, stat.Size, stat.ContentType, object, nil)
}

func NewImageController() ImageController {
	client := minioClient.GetMinioClient()
	return &imageController{minio: client}
}

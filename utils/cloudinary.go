package utils

import (
	"context"
	"errors"
	"strings"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/spf13/viper"
)

func UploadImage(base64 string) (*uploader.UploadResult, error) {

	if len(base64) <= 10 || base64[0:10] != "data:image" {
		return nil, errors.New("image is not a valid base64")
	}

	storage, err := cloudinary.NewFromParams(
		viper.GetString("cloudinary.name"),
		viper.GetString("cloudinary.apiKey"),
		viper.GetString("cloudinary.apiSecret"))

	if err != nil {
		return nil, err
	}

	result, err := storage.Upload.Upload(context.Background(), base64, uploader.UploadParams{Folder: viper.GetString("cloudinary.folder")})

	return result, err
}

func GetImageNameFromUrl(imageUrl string) string {

	imageUrlArr := strings.Split(imageUrl, "/")
	imageName := imageUrlArr[len(imageUrlArr)-1]
	return imageName
}

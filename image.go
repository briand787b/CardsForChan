package main

import (
	"time"
	"github.com/disintegration/imaging"
	"image"
)

type Image struct {
	ID 		int
	MD5Sum		string
	Location 	string
	Size 		int64
	CreatedAt 	time.Time
}

type ImageStore interface {
	Find(id int) (*Image, error)
	FindAll(offset int) ([]Image, error)
	FindAllByUser(user *User, offset int) ([]Image, error)
}


var thumbnailWidth = 400
var widthPreview = 800

func (image *Image) CreateResizedImages() error {
	// Generate an image from a file
	srcImage, err := imaging.Open("./data/images/" + image.Location)
	if err != nil {
		return err
	}

	// Create a channel to send errors on
	errChan := make(chan error)

	// Process each size
	go image.resizePreview(errChan, srcImage)
	go image.resizeThumbnail(errChan, srcImage)

	// Wait for images to finish resizing
	for i := 0; i < 2; i++ {
		err = <- errChan
		if err != nil {
			return err
		}
	}

	return nil
}

func (image *Image) resizeThumbnail(errChan chan error, srcImage image.Image) {
	dstImage := imaging.Thumbnail(srcImage, thumbnailWidth, thumbnailWidth, imaging.Lanczos)
	destination := "./data/images/thumbnail/" + image.Location
	errChan <- imaging.Save(dstImage, destination)
}

func (image *Image) resizePreview(errChan chan error, srcImage image.Image) {
	size := srcImage.Bounds().Size()
	ratio := float64(size.Y) / float64(size.X)
	targetHeight := int(float64(widthPreview) * ratio)

	dstImage := imaging.Resize(srcImage, widthPreview, targetHeight, imaging.Lanczos)
	destination := "./data/images/preview/" + image.Location

	errChan <- imaging.Save(dstImage, destination)
}

func (image *Image) StaticRoute() string {
	return "/im/" + image.Location
}

func (image *Image) ShowRoute() string {
	return "/image/" + string(image.ID)
}

func (image *Image) StaticThumbnailRoute() string {
	return "/im/thumbnail/" + image.Location
}

func (image *Image) StaticPreviewRoute() string {
	return "/im/preview/" + image.Location
}

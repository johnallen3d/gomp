package main

import (
	"image"
	"os"

	"github.com/dhowden/tag"
	"github.com/nfnt/resize"
	"github.com/spf13/viper"
)

/*
	Gets the Image Path from the uri to the string passed
	if embedded image is found the path to that Image is returned else
	path to default image is sent.
*/
func getAlbumArt(uri string) string {
	var path string = "/H/code/goMP/default.jpg"
	f, err := os.Open(uri)
	if err != nil {
		panic(err)
	}
	m, err := tag.ReadFrom(f)
	if err != nil {
		panic(err)
	}
	albumCover := m.Picture()
	if albumCover != nil {
		b, err := os.Create("/H/code/goMP/thumb.jpg")
		if err != nil {
			panic(err)
		}
		defer b.Close()
		b.Write(albumCover.Data)
		path = "/H/code/goMP/thumb.jpg"
		b.Close()
	}
	f.Close()
	return path
}

func getImg(uri string) (image.Image, error) {

	f, err := os.Open(uri)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		return nil, err
	}

	fw, fh := getFontWidth()
	img = resize.Resize(
		uint(float32(IMG_W)*(fw+float32(viper.GetFloat64("IMAGE_WIDTH_EXTRA_X")))), uint(float32(IMG_H)*(fh+float32(viper.GetFloat64("IMAGE_WIDTH_EXTRA_Y")))),
		img,
		resize.Bilinear,
	)

	return img, nil
}

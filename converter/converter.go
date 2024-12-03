package converter

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"os"
	"strings"

	"github.com/chai2010/webp"
	"golang.org/x/image/tiff"
)

// GetDecoder returns the appropriate decoder for an image type.
func GetDecoder(format string) (func(io.Reader) (image.Image, error), error) {
	switch format {
	case "png":
		return png.Decode, nil
	case "jpeg", "jpg":
		return jpeg.Decode, nil
	case "tiff":
		return tiff.Decode, nil
	case "webp":
		return webp.Decode, nil
	default:
		return nil, fmt.Errorf("unsupported input format: %s", format)
	}
}

// GetEncoder returns the appropriate encoder for an image type.
func GetEncoder(format string) (func(*os.File, image.Image) error, error) {
	switch format {
	case "png":
		return func(outFile *os.File, img image.Image) error {
			return png.Encode(outFile, img)
		}, nil
	case "jpeg", "jpg":
		return func(outFile *os.File, img image.Image) error {
			return jpeg.Encode(outFile, img, &jpeg.Options{Quality: 90})
		}, nil
	case "tiff":
		return func(outFile *os.File, img image.Image) error {
			return tiff.Encode(outFile, img, nil)
		}, nil
	case "webp":
		return func(outFile *os.File, img image.Image) error {
			return webp.Encode(outFile, img, &webp.Options{Lossless: true})
		}, nil
	default:
		return nil, fmt.Errorf("unsupported output format: %s", format)
	}
}

func GetFormat(format string) string {
	return strings.TrimPrefix(strings.ToLower(strings.Split(format, ".")[1]), ".")
}

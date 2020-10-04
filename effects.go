package composer

import (
	"github.com/anthonynsimon/bild/adjust"
	"github.com/anthonynsimon/bild/blur"
	"github.com/anthonynsimon/bild/effect"
	"github.com/anthonynsimon/bild/paint"
	"github.com/anthonynsimon/bild/transform"
	"github.com/getevo/evo/lib"
	"image"
	"image/color"
	"regexp"
	"strings"
)
var effectRegex = regexp.MustCompile(`(?m)(\w+)\((.*?)\)`)
func ApplyEffect(apply string, im image.Image) image.Image {
	var effects = effectRegex.FindAllStringSubmatch(apply,-1)
	for _,item := range effects{
		var params = strings.Split(item[2],",")
		switch strings.ToLower(item[1]) {
		case "brightness":
			im = adjust.Brightness(im,lib.ParseSafeFloat(params[0]))
		case "contrast":
			im = adjust.Contrast(im,lib.ParseSafeFloat(params[0]))
		case "gamma":
			im = adjust.Gamma(im,lib.ParseSafeFloat(params[0]))
		case "hue":
			im = adjust.Hue(im,lib.ParseSafeInt(params[0]))
		case "saturation":
			im = adjust.Saturation(im,lib.ParseSafeFloat(params[0]))
		case "dilate":
			im = effect.Dilate(im,lib.ParseSafeFloat(params[0]))
		case "blur.box":
			im = blur.Box(im,lib.ParseSafeFloat(params[0]))
		case "blur.gaussian":
			im = blur.Gaussian(im,lib.ParseSafeFloat(params[0]))
		case "edge":
			im = effect.EdgeDetection(im,lib.ParseSafeFloat(params[0]))
		case "emboss":
			im = effect.Emboss(im)
		case "erode":
			im = effect.Erode(im,lib.ParseSafeFloat(params[0]))
		case "grayscale":
			im = effect.Grayscale(im)
		case "median":
			im = effect.Median(im,lib.ParseSafeFloat(params[0]))
		case "sepia":
			im = effect.Sepia(im)
		case "sharpen":
			im = effect.Sharpen(im)
		case "sobel":
			im = effect.Sobel(im)
		case "unsharpmask":
			if len(params) == 2 {
				im = effect.UnsharpMask(im, lib.ParseSafeFloat(params[0]), lib.ParseSafeFloat(params[1]))
			}
		case "crop":
			if len(params) == 4 {
				im = transform.Crop(im, image.Rect(lib.ParseSafeInt(params[0]), lib.ParseSafeInt(params[1]), lib.ParseSafeInt(params[2]), lib.ParseSafeInt(params[3])))
			}
		case "invert":
			im = effect.Invert(im)
		case "flipv":
			im = transform.FlipV(im)
		case "fliph":
			im = transform.FlipH(im)
		case "removebackground":
			im = RemoveBackground(im,lib.ParseSafeInt(params[0]))
		case "floodfill":
			if len(params) == 4 {
				c, err := ParseColor(params[2])
				if err != nil {
					im = paint.FloodFill(im, image.Point{lib.ParseSafeInt(params[0]), lib.ParseSafeInt(params[1])}, c, uint8(lib.ParseSafeInt(params[3])))
				}
			}

		}
	}

	return im
}

func RemoveBackground(img image.Image,sensitivity int) image.Image{
	result := paint.FloodFill(img, image.Point{0, 0}, color.RGBA{0, 0, 0, 0}, uint8(sensitivity))
	result = paint.FloodFill(result, image.Point{img.Bounds().Max.X, 0}, color.RGBA{0, 0, 0, 0}, uint8(sensitivity))
	result = paint.FloodFill(result, image.Point{0, img.Bounds().Max.Y}, color.RGBA{0, 0, 0, 0}, uint8(sensitivity))
	result = paint.FloodFill(result, image.Point{img.Bounds().Max.X, img.Bounds().Max.Y}, color.RGBA{0, 0, 0, 0}, uint8(sensitivity))
	return result
}
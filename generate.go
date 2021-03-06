package composer

import (
	"fmt"
	"github.com/chai2010/webp"
	"github.com/disintegration/imaging"
	"github.com/nfnt/resize"
	"gopkg.in/fogleman/gg.v1"
	"image"
	"image/jpeg"
	"io"
	"net/http"
	"strings"
)

var imageCache = map[string]*image.Image{}

func (o *Image) Create() error {
	o.Context = gg.NewContext(o.Width, o.Height)
	for _, item := range o.Objects {
		err := item.DrawTo(o.Context)
		if err != nil {
			return err
		}
	}

	return nil
}

func (o *Object) DrawTo(context *gg.Context) error {
	var err error
	if o.Type == IMAGE {
		err = o.drawImageTo(context)
	} else if o.Type == TEXT {
		err = o.drawTextTo(context)
	} else if o.Type == RECT {
		err = o.drawRectTo(context)
	} else if o.Type == OVAL {
		err = o.drawOvalTo(context)
	}

	return err
}

func (o *Image) SaveJPG(path string, quality int) error {
	if o.Context == nil {
		return fmt.Errorf("empty context")
	}
	return imaging.Save(o.Context.Image(), path, imaging.JPEGQuality(quality))
}

func (o *Image) SavePNG(path string) error {
	if o.Context == nil {
		return fmt.Errorf("empty context")
	}
	return o.Context.SavePNG(path)
}

func (o *Image) Rotate(angle float64) {
	o.Context.Rotate(angle)
}

func (o *Image) EncodeJPG(w io.Writer, j *jpeg.Options) error {
	return imaging.Encode(w, o.Context.Image(), imaging.JPEG)

}

func (o *Image) EncodePNG(w io.Writer) error {
	return o.Context.EncodePNG(w)
}

func (o *Image) SaveWebP(path string, lossless, exact bool, quality float32) error {
	if o.Context == nil {
		return fmt.Errorf("empty context")
	}
	return webp.Save(path, o.Context.Image(), &webp.Options{
		Lossless: lossless,
		Quality:  quality,
		Exact:    exact,
	})

}

func (o *Image) Image() image.Image {
	return o.Context.Image()
}

func (o *Image) Scale(x, y float64) {
	o.Context.Scale(x, y)
}

func (o *Image) Resize(w, h int) {
	im := resize.Resize(uint(w), uint(h), o.Context.Image(), resize.Lanczos3)
	o.Context = gg.NewContext(w, h)
	o.Context.DrawImage(im, 0, 0)
}

func (o *Object) drawOvalTo(context *gg.Context) error {
	context.DrawEllipse(float64(o.Left), float64(o.Top), float64(o.Width), float64(o.Height))
	c, err := ParseColor(o.Color)
	if err != nil {
		return err
	}
	context.SetColor(c)
	context.Fill()
	return nil
}

func (o *Object) drawRectTo(context *gg.Context) error {
	context.DrawRectangle(float64(o.Left), float64(o.Top), float64(o.Width), float64(o.Height))
	c, err := ParseColor(o.Color)
	if err != nil {
		return err
	}
	context.SetColor(c)
	context.Fill()
	return nil
}

func (o *Object) drawImageTo(context *gg.Context) error {
	var im *image.Image
	if o.Cache {
		if cached, ok := imageCache[o.Value]; ok {
			im = cached
		}
	}
	if im == nil {
		if strings.HasPrefix(o.Value, "http") {
			res, err := http.Get(o.Value)
			if err != nil || res.StatusCode != 200 {
				return err
			}
			defer res.Body.Close()
			img, _, err := image.Decode(res.Body)
			if err != nil {
				return err
			}
			im = &img
		} else {
			img, err := gg.LoadImage(o.Value)
			if err != nil {
				return err
			}
			im = &img
		}
		if o.Cache {
			imageCache[o.Value] = im
		}
	}

	if o.Width != 0 && o.Height != 0 {
		img := resize.Resize(uint(o.Width), uint(o.Height), *im, resize.Lanczos3)
		im = &img
	} else if o.Width != 0 || o.Height != 0 {
		img := resize.Thumbnail(uint(o.Width), uint(o.Height), *im, resize.Lanczos3)
		im = &img
	}
	if o.Effect != "" {
		img := ApplyEffect(o.Effect, *im)
		im = &img
	}
	if o.HAlign == CENTER {
		o.Left -= (*im).Bounds().Max.X / 2
	} else if o.HAlign == RIGHT {
		o.Left -= (*im).Bounds().Max.X
	}
	if o.VAlign == MIDDLE {
		o.Top -= (*im).Bounds().Max.Y / 2
	} else if o.VAlign == BOTTOM {
		o.Top -= (*im).Bounds().Max.Y
	}
	context.DrawImage(*im, o.Left, o.Top)
	return nil
}

func (o *Object) drawTextTo(context *gg.Context) error {
	context.LoadFontFace(o.Font, float64(o.FontSize))
	c, err := ParseColor(o.Color)
	if err != nil {
		return err
	}
	context.SetColor(c)

	var tw, th float64
	if o.WordWrap {
		tw, th = context.MeasureMultilineString(o.Value, o.LineSpacing)
	} else {
		tw, th = context.MeasureString(o.Value)
	}
	top := float64(o.Top)
	left := float64(o.Left)
	var align = gg.AlignLeft
	if o.VAlign == MIDDLE {
		top -= th / 2
	} else if o.VAlign == BOTTOM {
		top -= th
	}
	if o.HAlign == CENTER {
		align = gg.AlignCenter
	} else if o.HAlign == RIGHT {
		align = gg.AlignRight
		left += tw
	}
	if o.WordWrap {
		context.DrawStringWrapped(o.Value, left, top, 0.5, 0.5, float64(o.Width), o.LineSpacing, align)
	} else {
		context.DrawString(o.Value, left, top)
	}
	return nil
}

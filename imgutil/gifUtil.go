package imgutil

import (
	"gitee.com/dk83/goutils/zlog"
	"github.com/disintegration/gift"
	"image"
	"image/color"
	"image/draw"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"
	"time"
)

//https://github.com/ajmadsen/jingleping-go/blob/e8a803e9af3a969c1fa6ae9f199a0b7e418583d7/images.go

// maskNonTransparent wraps an existing image and is useful for generating an
// alpha mask of the non-black and non-transparent pixels of the image.
type maskNonTransparent struct {
	image.Image
}

// Bounds returns the bounds of the wrapped image, since the mask should be the
// same size.
func (i maskNonTransparent) Bounds() image.Rectangle {
	return i.Image.Bounds()
}

// ColorModel is always AlphaModel since we are generating an alpha mask.
func (i maskNonTransparent) ColorModel() color.Model {
	return color.AlphaModel
}

// At returns a fully opaque pixel if the wrapped image has a non-black and non-
// transparent pixel at the same location. It ignores black pixels since black
// pixels are "invisible" on the lighted display board.
func (i maskNonTransparent) At(x, y int) color.Color {
	_, _, _, a := i.Image.At(x, y).RGBA()
	// transparent or black pixel
	if a != 0 {
		//return color.Alpha{255}
	}
	return color.Alpha{0}
}

// copyRGBA copies an RGBA image. It was supposed to be an optimization for
// copying the canvas for the current frame specific to RGBA, but a generic
// draw implementation was used instead for specifically typed images.
func copyRGBA(im *image.RGBA) *image.RGBA {
	second := image.NewRGBA(im.Bounds())
	draw.Draw(second, im.Bounds(), im, image.Point{0, 0}, draw.Src)
	return second
}

func MakeImageArray(g *gif.GIF) []image.Image {
	gifLen := len(g.Image)
	explodedGif := make([]image.Image, gifLen)
	lastNpd := 0 // The last non-previous disposal Image index
	imageBounds := image.Rect(0, 0, g.Config.Width, g.Config.Height)

	drawOverBg := func(img image.Image) *image.RGBA {
		newimg := image.NewRGBA(imageBounds)
		var c color.Color
		if g.BackgroundIndex >= 0 && g.Config.ColorModel != nil && len(g.Config.ColorModel.(color.Palette)) > 0 {
			c = g.Config.ColorModel.(color.Palette)[g.BackgroundIndex]
		} else {
			c = color.Transparent
		}
		draw.Draw(newimg, imageBounds, &image.Uniform{c}, imageBounds.Min, draw.Src)
		b := img.Bounds()
		draw.Draw(newimg, b, img, b.Min, draw.Over)
		return newimg
	}

	drawOverImg := func(lower, higher image.Image) *image.RGBA {
		newimg := image.NewRGBA(imageBounds)
		b := higher.Bounds()
		draw.Draw(newimg, imageBounds, lower, image.ZP, draw.Src)
		draw.Draw(newimg, b, higher, b.Min, draw.Over)
		return newimg
	}

	for i, v := range g.Image {
		if i == 0 {
			explodedGif[i] = drawOverBg(v)
			continue
		}

		switch g.Disposal[i-1] {
		case gif.DisposalBackground:
			explodedGif[i] = drawOverBg(v)
			lastNpd = i
		case gif.DisposalPrevious:
			explodedGif[i] = drawOverImg(explodedGif[lastNpd], v)
		default:
			explodedGif[i] = drawOverImg(explodedGif[i-1], v)
			lastNpd = i
		}
	}

	return explodedGif

}

func CutGif(srcGif *gif.GIF, destFile string) error {
	d, err := os.Create(destFile)
	if err != nil {
		zlog.Error(err)
		return err
	}
	defer d.Close()
	srcGif.Delay = srcGif.Delay[:maxFrameSize]
	srcGif.Image = srcGif.Image[:maxFrameSize]
	srcGif.Disposal = srcGif.Disposal[:maxFrameSize]
	err = gif.EncodeAll(d, srcGif)
	if err != nil {
		zlog.Error(err)
		return err
	}
	return nil
}

func ScaleToSize(pic image.Image, maxSize int) (*image.RGBA, float64, error) {
	g := gift.New()
	size := float64(maxSize)
	scale := float64(0)
	if float64(pic.Bounds().Size().X)/float64(pic.Bounds().Size().Y) > 1 {
		scale = size / float64(pic.Bounds().Size().X)
		g = gift.New(gift.Resize(int(size), int(float64(pic.Bounds().Size().Y)*scale), gift.LinearResampling))
	} else {
		scale = size / float64(pic.Bounds().Size().Y)
		g = gift.New(gift.Resize(int(float64(pic.Bounds().Size().X)*scale), int(size), gift.LinearResampling))
	}

	rotateImg := image.NewRGBA(g.Bounds(pic.Bounds()))
	g.Draw(rotateImg, pic)

	return rotateImg, scale, nil
}

func ScaleFileToUV(file string, w int, h int) (string, error) {
	f, err := os.Open(file)
	if err != nil {
		return file, err
	}

	defer func() {
		f.Close()
	}()
	pic, _, err := image.Decode(f)
	if err != nil {
		return file, err
	}

	img := ScaleToUV(pic, w, h)
	scale := float64(w) / float64(h)
	ext := strings.ToLower(filepath.Ext(file))

	if float64(pic.Bounds().Size().X)/float64(pic.Bounds().Size().Y) != scale {
		ext = ".png"
	}

	dst := filepath.Join(filepath.Dir(file), time.Now().Format("20060102150405")+ext)
	if ext == ".png" {
		f1, err := os.Create(dst)
		if err != nil {
			return file, err
		}

		defer func() { f1.Close() }()

		err = png.Encode(f1, img)
	} else {
		f1, err := os.Create(dst)
		if err != nil {
			return file, err
		}

		defer func() { f1.Close() }()

		err = jpeg.Encode(f1, img, &jpeg.Options{Quality: 100})
	}

	return dst, err
}

func ScaleToUV(pic image.Image, w int, h int) image.Image {
	var img *image.RGBA
	scale := float64(w) / float64(h)

	img = image.NewRGBA(image.Rect(0, 0, w, h))

	if float64(pic.Bounds().Size().X)/float64(pic.Bounds().Size().Y) > scale {
		//宽更长   填充Y
		g := gift.New(gift.Resize(w, int(float64(w)*float64(pic.Bounds().Size().Y)/float64(pic.Bounds().Size().X)), gift.LinearResampling))

		g.DrawAt(img, pic, image.Pt(0, int((float64(img.Bounds().Size().Y)-float64(g.Bounds(pic.Bounds()).Size().Y))/2)), gift.CopyOperator)
	} else {

		g := gift.New(gift.Resize(int(float64(h)*float64(pic.Bounds().Size().X)/float64(pic.Bounds().Size().Y)), h, gift.LinearResampling))

		g.DrawAt(img, pic, image.Pt(int((float64(img.Bounds().Size().X)-float64(g.Bounds(pic.Bounds()).Size().X))/2), 0), gift.CopyOperator)
	}

	return img
}

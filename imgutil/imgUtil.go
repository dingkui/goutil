package imgutil

import (
	"fmt"
	"gitee.com/dk83/goutils/zlog"
	"github.com/disintegration/gift"
	"image"
	"image/gif"
	_ "image/jpeg"
	"image/png"
	"math"
	"os"
)

const maxLineNum = 6
const maxSize = 339 //(2048 - 2) / 6 - 2
const maxFrameSize = maxLineNum * maxLineNum

//图片合并成一个列表
//一行7个 多行
func Images2Plane(images *[]*image.RGBA) *image.RGBA {

	w := ((*images)[0].Bounds().Size().X + 4) * len(*images)
	h := (*images)[0].Bounds().Size().Y

	iv := math.Min(float64(len(*images)), maxLineNum)
	ih := math.Ceil(float64(len(*images)) / maxLineNum)
	w = ((*images)[0].Bounds().Size().X + 4) * int(iv)
	h = ((*images)[0].Bounds().Size().Y + 4) * int(ih)

	img := image.NewRGBA(image.Rect(0, 0, w, h))

	for j := 0; j < int(ih); j++ {
		for i := 0; i < int(iv); i++ {
			if i+j*maxLineNum >= len(*images) {
				break
			}
			im := (*images)[i+j*maxLineNum]
			g := gift.New()
			g.DrawAt(img, im, image.Pt(i*((*images)[0].Bounds().Size().X+4)+2, j*((*images)[0].Bounds().Size().Y+4)+2), gift.CopyOperator)
		}
	}

	return img
}

func ConvertGif(src, dst string, width, height int) ([]int, error) {
	f, err := os.Open(src)
	if err != nil {
		return nil, err
	}

	defer func() {
		f.Close()
	}()

	gifs, err := gif.DecodeAll(f)
	if err != nil {
		return nil, err
	}

	if len(gifs.Image) > maxFrameSize {
		gifs.Delay = gifs.Delay[:maxFrameSize]
		gifs.Image = gifs.Image[:maxFrameSize]
		gifs.Disposal = gifs.Disposal[:maxFrameSize]
	}

	pics := make([]*image.RGBA, 0)

	frames := MakeImageArray(gifs)
	for i, value := range frames {
		//zlog.Debug(value.At(0,0))
		g := gift.New()
		d := image.NewRGBA(image.Rect(0, 0, gifs.Config.Width, gifs.Config.Height))
		g.DrawAt(d, value, value.Bounds().Min, gift.CopyOperator)

		img2 := ScaleToUV(d, width, height)

		d = img2.(*image.RGBA)

		imgRGBA, _, err := ScaleToSize(d, maxSize)
		if err != nil {
			return nil, err
		}

		ff, err := os.Create(fmt.Sprintf("temp/%d.png", i))
		if err != nil {
			zlog.Error(err)
		}

		defer func() { ff.Close() }()
		png.Encode(ff, imgRGBA)

		pics = append(pics, imgRGBA)
	}

	rgbs := Images2Plane(&pics)

	f2, err := os.Create(dst)
	if err != nil {
		return nil, err
	}

	defer func() { f2.Close() }()

	err = png.Encode(f2, rgbs)

	return gifs.Delay, err
}

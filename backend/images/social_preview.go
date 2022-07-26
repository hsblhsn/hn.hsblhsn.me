package images

import (
	"bytes"
	"embed"
	"github.com/disintegration/imaging"
	"github.com/fogleman/gg"
	"github.com/golang/freetype/truetype"
	"github.com/pkg/errors"
	"golang.org/x/image/font"
	"image"
	"image/color"
	"path/filepath"
)

const (
	SocialPreviewWidth  = 1200
	SocialPreviewHeight = 628
)

type SocialPreviewGenerator struct {
	background   image.Image
	titleFont    font.Face
	brandingFont font.Face
}

func (sp *SocialPreviewGenerator) Generate(title string) (image.Image, error) {
	maximumTitleLength := 120
	if len(title) > maximumTitleLength {
		ellipsis := "..."
		title = title[:maximumTitleLength-len(ellipsis)]
		title += ellipsis
	}
	dc := gg.NewContext(SocialPreviewWidth, SocialPreviewHeight)
	dc.DrawImage(sp.background, 0, 0)

	// add background rectangle
	margin := 20.0
	x := margin
	y := margin
	w := float64(dc.Width()) - (2.0 * margin)
	h := float64(dc.Height()) - (2.0 * margin)
	dc.SetColor(color.RGBA{A: 204})
	dc.DrawRectangle(x, y, w, h)
	dc.Fill()

	// add branding logo
	dc.SetFontFace(sp.brandingFont)
	dc.SetColor(color.White)
	s := "Hackernews"
	marginX := 50.0
	marginY := 30.0
	textWidth, textHeight := dc.MeasureString(s)
	x = float64(dc.Width()) - textWidth - marginX
	y = float64(dc.Height()) - textHeight - marginY
	dc.DrawString(s, x, y)

	// add branding website
	dc.SetFontFace(sp.brandingFont)
	r, g, b, _ := color.White.RGBA()
	mutedColor := color.RGBA{
		R: uint8(r),
		G: uint8(g),
		B: uint8(b),
		A: uint8(200),
	}
	dc.SetColor(mutedColor)
	marginY = 30
	s = "https://hn.hsblhsn.me/"
	_, textHeight = dc.MeasureString(s)
	x = 70
	y = float64(dc.Height()) - textHeight - marginY
	dc.DrawString(s, x, y)

	// add title
	dc.SetFontFace(sp.titleFont)
	textRightMargin := 60.0
	textTopMargin := 60.0
	x = textRightMargin
	y = textTopMargin
	maxWidth := float64(dc.Width()) - textRightMargin - textRightMargin
	dc.SetColor(color.Black)
	dc.DrawStringWrapped(title, x+1, y+1, 0, 0, maxWidth, 1.5, gg.AlignLeft)
	dc.SetColor(color.White)
	dc.DrawStringWrapped(title, x, y, 0, 0, maxWidth, 1.5, gg.AlignLeft)
	return dc.Image(), nil
}

//go:embed resources/*
var resources embed.FS

func readResource(path string) []byte {
	b, err := resources.ReadFile(filepath.Join("resources", path))
	if err != nil {
		panic(err)
	}
	return b
}

func newSocialPreviewGenerator() (*SocialPreviewGenerator, error) {
	imgBytes := readResource("background_image.jpg")
	img, _, err := image.Decode(bytes.NewReader(imgBytes))
	if err != nil {
		return nil, errors.Wrap(err, "images: could not decode image")
	}
	img = imaging.Fill(img, SocialPreviewWidth, SocialPreviewHeight, imaging.Center, imaging.Lanczos)

	titleFont, err := loadFontFace(readResource("title_font.ttf"), 60)
	if err != nil {
		return nil, errors.Wrap(err, "images: could not load title font")
	}
	brandingFont, err := loadFontFace(readResource("default_font.ttf"), 30)
	if err != nil {
		return nil, errors.Wrap(err, "images: could not load branding font")
	}
	return &SocialPreviewGenerator{
		background:   img,
		titleFont:    titleFont,
		brandingFont: brandingFont,
	}, nil
}

func loadFontFace(fontBytes []byte, points float64) (font.Face, error) {
	f, err := truetype.Parse(fontBytes)
	if err != nil {
		return nil, err
	}
	face := truetype.NewFace(f, &truetype.Options{
		Size: points,
	})
	return face, nil
}

package main

import (
	"strings"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

var (
	LineHeight, LetterSpace int32 = int32(TXDEF + 4), 1

	txSurf                           *sdl.Surface
	fontcharsDEFAULT, fontcharsFONT1 []FONTCH

	standardCharacters = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789:;<=>?!#$%&'()*+,-./@[]^_`{|}~'\"' "

	FontDefault, Font1 *ttf.Font

	TX8, TX10, TX12, TX14, TX18, TX24, TX30, TX36, TX48, TX60, TX72, TX100, TX150, TX200 = 8, 10, 12, 14, 18, 24, 30, 36, 48, 60, 72, 100, 150, 200

	TXDEF = TX14
)

type FONTCH struct {
	character string
	tex       *sdl.Texture
	rec       sdl.Rect
}

func fontsInit() {
	ttf.Init()
	FontDefault, Err = ttf.OpenFont("fonts/Rubik-Medium.ttf", TXDEF)
	if Err != nil {
		Errlist = append(Errlist, Err)
	}
	Font1, Err = ttf.OpenFont("fonts/Rancho-Regular.ttf", TX24)
	if Err != nil {
		Errlist = append(Errlist, Err)
	}
}

func TextHereXYܛTEXT(txt string, x, y int32) {
	t := strings.Split(txt, "")
	for i := 0; i < len(t); i++ {
		for j := 0; j < len(fontcharsDEFAULT); j++ {
			if t[i] == fontcharsDEFAULT[j].character {
				Rndr.Copy(fontcharsDEFAULT[j].tex, &fontcharsDEFAULT[j].rec, &sdl.Rect{x, y, fontcharsDEFAULT[j].rec.W, fontcharsDEFAULT[j].rec.H})
				x += fontcharsDEFAULT[j].rec.W + LetterSpace
				break
			}
		}
	}
}

func TextHereܛTEXT(txt string, point sdl.Point) {
	t := strings.Split(txt, "")
	x := point.X
	y := point.Y
	//spc := int32(2)
	for i := 0; i < len(t); i++ {
		for j := 0; j < len(fontcharsDEFAULT); j++ {
			if t[i] == fontcharsDEFAULT[j].character {
				Rndr.Copy(fontcharsDEFAULT[j].tex, &fontcharsDEFAULT[j].rec, &sdl.Rect{x, y, fontcharsDEFAULT[j].rec.W, fontcharsDEFAULT[j].rec.H})
				x += fontcharsDEFAULT[j].rec.W + LetterSpace
				break
			}
		}
	}
}

func TextHereCenterܛTEXT(txt string, point sdl.Point, fontNum int) {
	t := strings.Split(txt, "")
	x := point.X
	y := point.Y
	spc := int32(2)
	width := int32(0)
	height := int32(0)
	for i := 0; i < len(t); i++ {
		for j := 0; j < len(fontcharsDEFAULT); j++ {
			if t[i] == fontcharsDEFAULT[j].character {
				width += fontcharsDEFAULT[j].rec.W + spc
				if height == 0 {
					height = fontcharsDEFAULT[j].rec.H / 2
				}
				break
			}
		}
	}
	x -= width / 2
	y -= height
	for i := 0; i < len(t); i++ {
		for j := 0; j < len(fontcharsDEFAULT); j++ {
			if t[i] == fontcharsDEFAULT[j].character {
				Rndr.Copy(fontcharsDEFAULT[j].tex, &fontcharsDEFAULT[j].rec, &sdl.Rect{x, y, fontcharsDEFAULT[j].rec.W, fontcharsDEFAULT[j].rec.H})
				x += fontcharsDEFAULT[j].rec.W + spc
				break
			}
		}
	}
}
func TextHereCenterColorܛTEXT(txt string, point sdl.Point, color []uint8, fontNum int) {

	var fontSelected []FONTCH

	switch fontNum {
	case 0: //FONT DEFAULT
		fontSelected = fontcharsDEFAULT
	case 1: //FONT1
		fontSelected = fontcharsFONT1
	}

	t := strings.Split(txt, "")
	x := point.X
	y := point.Y
	spc := int32(2)
	width := int32(0)
	height := int32(0)
	for i := 0; i < len(t); i++ {
		for j := 0; j < len(fontSelected); j++ {
			if t[i] == fontSelected[j].character {
				width += fontSelected[j].rec.W + spc
				if height == 0 {
					height = fontSelected[j].rec.H / 2
				}
				break
			}
		}
	}
	x -= width / 2
	y -= height
	for i := 0; i < len(t); i++ {
		for j := 0; j < len(fontSelected); j++ {
			if t[i] == fontSelected[j].character {
				fontSelected[j].tex.SetColorMod(color[0], color[1], color[2])
				fontSelected[j].tex.SetAlphaMod(color[3])
				Rndr.Copy(fontSelected[j].tex, &fontSelected[j].rec, &sdl.Rect{x, y, fontSelected[j].rec.W, fontSelected[j].rec.H})
				x += fontSelected[j].rec.W + spc
				break
			}
		}
	}
	//fontcharsDEFAULT = fontTexRevert(fontcharsDEFAULT)
}
func fontTexRevert(f []FONTCH) []FONTCH {
	for i := 0; i < len(f); i++ {
		fontcharsDEFAULT[i].tex.SetColorMod(255, 255, 255)
		fontcharsDEFAULT[i].tex.SetAlphaMod(255)
	}
	return f
}
func fontTextureSheet() {
	t := strings.Split(standardCharacters, "")
	for i := 0; i < len(t); i++ {
		fontcharsDEFAULT = append(fontcharsDEFAULT, fontCharCreate(t[i], 0))
		fontcharsFONT1 = append(fontcharsFONT1, fontCharCreate(t[i], 1))
	}
}

func fontCharCreate(singleCharacter string, fontNum int) FONTCH {

	var w, h int

	switch fontNum {
	case 0: //FONT DEFAULT
		txSurf, _ = FontDefault.RenderUTF8Blended(singleCharacter, WHITEǁ3())
		defer txSurf.Free()
		w, h, _ = FontDefault.SizeUTF8(singleCharacter)
	case 1: //FONT1
		txSurf, _ = Font1.RenderUTF8Blended(singleCharacter, WHITEǁ3())
		defer txSurf.Free()
		w, h, _ = Font1.SizeUTF8(singleCharacter)
	}

	f := FONTCH{}
	f.tex, _ = Rndr.CreateTextureFromSurface(txSurf)
	f.rec = sdl.Rect{0, 0, int32(w), int32(h)}
	f.character = singleCharacter
	return f
}

/*
// GRID TEXT GRID TEXT GRID TEXT GRID TEXT GRID TEXT GRID TEXT GRID TEXT
func DrܛGridNumsONLY(grid GridIso) {
	for i := 0; i < len(grid.recs); i++ {
		TxtHere(fmt.Sprint(i), grid.recs[i].cntr)
	}
}

// DRAW TEXT DRAW TEXT DRAW TEXT DRAW TEXT DRAW TEXT DRAW TEXT
func DrܟText(txt string, x, y int32) {
	color := WHITEǁ3()
	if ChkܢRGBAColorsMatch(WinBackgroundColor, ColorSDLtoSlice(color)) {
		color = BLACKǁ3()
	}
	t, _ := FontDefault.RenderUTF8Blended(txt, color)
	r := sdl.Rect{x, y, t.W, t.H}
	t.Blit(&t.ClipRect, Surface, &r)
	t.Free()
}
func DrܟTextColor(txt string, x, y int32, color sdl.Color) {
	t, _ := FontDefault.RenderUTF8Blended(txt, color)
	r := sdl.Rect{x, y, t.W, t.H}
	t.Blit(&t.ClipRect, Surface, &r)
	t.Free()
}
func DrܟTextCenterPoint(txt string, point sdl.Point) {
	color := WHITEǁ3()
	if ChkܢRGBAColorsMatch(WinBackgroundColor, ColorSDLtoSlice(color)) {
		color = BLACKǁ3()
	}
	w, h := TextSize(txt)
	x := point.X - w/2
	y := point.Y - h/2
	t, _ := FontDefault.RenderUTF8Blended(txt, color)
	r := sdl.Rect{x, y, t.W, t.H}
	t.Blit(&t.ClipRect, Surface, &r)
	t.Free()
}


func TextSize(txt string) (W, H int32) {
	t, _ := FontDefault.RenderUTF8Blended(txt, sdl.Color{R: 255, G: 255, B: 255, A: 255})
	W = t.W
	H = t.H
	t.Free()
	return W, H
}

func WriteFPS(x, y float32, c sdl.Color) {

		TxtSurface, Err = FontDefault.RenderUTF8Blended("FPS "+fmt.Sprintf("%.2f", 60), c)
		TxtRec.X = int32(x)
		TxtRec.Y = int32(y)
		TxtRec.W = TxtSurface.W
		TxtRec.H = TxtSurface.H
		TxtSurface.Blit(&TxtSurface.ClipRect, Surface, &TxtRec)
		Win.UpdateSurface()


}
*/

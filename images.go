package main

import (
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

var (
	imLoadSurf *sdl.Surface
	texImLoad  *sdl.Texture

	closeWinRec sdl.Rect

	loadTex *sdl.Texture

	anim8TILETEST TILESHEET

	etcIM []IM

	cursorSize      = int32(32)
	closeWinRecSize = int32(32)

	play1Anim, play2Anim, play3Anim ANIM
)

type IM struct {
	rec       sdl.Rect
	tex       *sdl.Texture
	tileSheet TILESHEET
}
type TILESHEET struct {
	tileIM                []IM
	tex                   *sdl.Texture
	name                  string
	widthTile, heightTile int32
	numTiles              int
}
type ANIM struct {
	im                 IM
	frames, framecount int
	xStart             int32
}
type ANIMSHEET8DIR struct {
	name                  string
	tex                   *sdl.Texture
	widthTile, heightTile int32

	up, rightUp, right, rightDown, down, leftDown, left, leftUp []ANIM8DIR
}
type ANIM8DIR struct {
	rec                  sdl.Rect
	xStart               int32
	numTiles, frameCount int
	animSheet            ANIMSHEET8DIR
}

func imagesܛMAKE() {

	//GROUND TILES
	loadTex = ImageToTextureܛLOAD("im/grassTiles.png")
	grassTILES = TileSheetܛMAKE(loadTex, 0, 0, 18, 18, 0, 0, 17, 1, "grass tiles")

	//TERRAIN
	var rockRecs []sdl.Rect
	rockRecs = append(rockRecs, sdl.Rect{0, 0, 63, 48}, sdl.Rect{64, 0, 61, 46}, sdl.Rect{127, 0, 34, 51}, sdl.Rect{162, 1, 28, 63}, sdl.Rect{0, 49, 57, 30}, sdl.Rect{0, 80, 24, 20}, sdl.Rect{63, 45, 60, 55}, sdl.Rect{135, 54, 22, 24}, sdl.Rect{161, 69, 24, 25})
	loadTex = ImageToTextureܛLOAD("im/rocks.png")
	rocksTERRAIN = TileSheetRecListܛMAKE(loadTex, rockRecs, "rocks")

	//ETC
	loadTex = ImageToTextureܛLOAD("im/etc.png")
	ii := IM{}
	ii.tex = loadTex
	ii.rec = sdl.Rect{0, 0, 12, 12}
	etcIM = append(etcIM, ii) //0 CURSOR
	ii.rec = sdl.Rect{12, 0, 16, 16}
	etcIM = append(etcIM, ii) //1 CLOSE WINDOW
	ii.rec = sdl.Rect{30, 0, 16, 16}
	etcIM = append(etcIM, ii) //2 SHIELD

	//CARD BG
	loadTex = ImageToTextureܛLOAD("im/cardBG.png")
	ii = IM{}
	ii.tex = loadTex
	ii.rec = sdl.Rect{0, 0, 565, 800}
	c := CARD{}
	c.im = ii
	c.nm = "Card Text"
	cards = append(cards, c)
	makeCards()

	//PLAYERS
	loadTex = ImageToTextureܛLOAD("im/players.png")
	ii = IM{}
	ii.tex = loadTex
	ii.rec = sdl.Rect{0, 0, 16, 16}
	play1Anim.im = ii
	play1Anim.frames = 2
	play1Anim.xStart = ii.rec.X

	makePlayers()
}

func ImageToTextureܛLOAD(path string) *sdl.Texture {
	imLoadSurf, Err = img.Load(path)
	if Err != nil {
		Errlist = append(Errlist, Err)
	}

	texImLoad, Err = Rndr.CreateTextureFromSurface(imLoadSurf)
	if Err != nil {
		Errlist = append(Errlist, Err)
	}
	return texImLoad
}

func TileSheetܛMAKE(tex *sdl.Texture, xStart, yStart, width, height, spaceWidth, spaceHeight int32, columns, rows int, name string) TILESHEET {

	t := TILESHEET{}
	var imgs []IM

	a := columns * rows
	t.numTiles = a
	t.widthTile = width
	t.heightTile = height
	t.name = name
	t.tex = tex

	x := xStart
	y := yStart
	c := 0
	i := IM{}
	for a > 0 {
		i.rec = sdl.Rect{x, y, width, height}
		imgs = append(imgs, i)
		x += width + spaceWidth
		c++
		a--
		if c == columns {
			c = 0
			y += height + spaceHeight
			x = xStart
		}
	}

	t.tileIM = imgs

	for i := 0; i < len(t.tileIM); i++ {
		t.tileIM[i].tileSheet = t
	}

	return t
}
func TileSheetRecListܛMAKE(tex *sdl.Texture, recList []sdl.Rect, name string) TILESHEET {

	t := TILESHEET{}
	var imgs []IM
	t.numTiles = len(recList)
	t.name = name
	t.tex = tex
	newIm := IM{}
	for i := 0; i < len(recList); i++ {
		newIm.rec = recList[i]
		imgs = append(imgs, newIm)
	}
	t.tileIM = imgs
	for i := 0; i < len(t.tileIM); i++ {
		t.tileIM[i].tileSheet = t
	}

	return t
}

// 8 DIRECTION ANIMATIONS
func Anim8DirSheetܛMAKE(tex *sdl.Texture, xStart, yStart, width, height, spaceWidth, spaceHeight int32, numTilesPerAnim int, name string, directionOrder []int) ANIMSHEET8DIR {
	a := ANIMSHEET8DIR{}
	a.name = name
	a.tex = tex
	a.widthTile = width
	a.heightTile = height

	columns := numTilesPerAnim * 8

	x := xStart
	y := yStart

	var a1, a2, a3, a4, a5, a6, a7, a8 []ANIM8DIR
	/*
		DIRECTION CLOCKWISE >>
		1 UP / 2 rightUp / 3 RIGHT / 4 RIGHTDOWN / 5 DOWN / 6 leftDown / 7 LEFT / 8 LEFTUP
	*/

	c2 := 1
	c3 := 0
	totalTiles := 8
	for totalTiles > 0 {

		b := ANIM8DIR{}
		b.xStart = x
		b.numTiles = numTilesPerAnim
		b.rec = sdl.Rect{x, y, width, height}

		switch c2 {
		case 1:
			a1 = append(a1, b)
		case 2:
			a2 = append(a2, b)
		case 3:
			a3 = append(a3, b)
		case 4:
			a4 = append(a4, b)
		case 5:
			a5 = append(a5, b)
		case 6:
			a6 = append(a6, b)
		case 7:
			a7 = append(a7, b)
		case 8:
			a8 = append(a8, b)
		}
		x += (width + spaceWidth) * int32(numTilesPerAnim)
		c3 += numTilesPerAnim
		c2++
		if c3 == columns {
			x = xStart
			y += height + spaceHeight
			c3 = 0
		}

		totalTiles--
	}

	for i := 0; i < len(directionOrder); i++ {
		switch directionOrder[i] {
		case 1:
			switch i {
			case 0:
				a.up = a1
			case 1:
				a.up = a2
			case 2:
				a.up = a3
			case 3:
				a.up = a4
			case 4:
				a.up = a5
			case 5:
				a.up = a6
			case 6:
				a.up = a7
			case 7:
				a.up = a8
			}
		case 2:
			switch i {
			case 0:
				a.rightUp = a1
			case 1:
				a.rightUp = a2
			case 2:
				a.rightUp = a3
			case 3:
				a.rightUp = a4
			case 4:
				a.rightUp = a5
			case 5:
				a.rightUp = a6
			case 6:
				a.rightUp = a7
			case 7:
				a.rightUp = a8
			}
		case 3:
			switch i {
			case 0:
				a.right = a1
			case 1:
				a.right = a2
			case 2:
				a.right = a3
			case 3:
				a.right = a4
			case 4:
				a.right = a5
			case 5:
				a.right = a6
			case 6:
				a.right = a7
			case 7:
				a.right = a8
			}
		case 4:
			switch i {
			case 0:
				a.rightDown = a1
			case 1:
				a.rightDown = a2
			case 2:
				a.rightDown = a3
			case 3:
				a.rightDown = a4
			case 4:
				a.rightDown = a5
			case 5:
				a.rightDown = a6
			case 6:
				a.rightDown = a7
			case 7:
				a.rightDown = a8
			}
		case 5:
			switch i {
			case 0:
				a.down = a1
			case 1:
				a.down = a2
			case 2:
				a.down = a3
			case 3:
				a.down = a4
			case 4:
				a.down = a5
			case 5:
				a.down = a6
			case 6:
				a.down = a7
			case 7:
				a.down = a8
			}
		case 6:
			switch i {
			case 0:
				a.leftDown = a1
			case 1:
				a.leftDown = a2
			case 2:
				a.leftDown = a3
			case 3:
				a.leftDown = a4
			case 4:
				a.leftDown = a5
			case 5:
				a.leftDown = a6
			case 6:
				a.leftDown = a7
			case 7:
				a.leftDown = a8
			}
		case 7:
			switch i {
			case 0:
				a.left = a1
			case 1:
				a.left = a2
			case 2:
				a.left = a3
			case 3:
				a.left = a4
			case 4:
				a.left = a5
			case 5:
				a.left = a6
			case 6:
				a.left = a7
			case 7:
				a.left = a8
			}
		case 8:
			switch i {
			case 0:
				a.leftUp = a1
			case 1:
				a.leftUp = a2
			case 2:
				a.leftUp = a3
			case 3:
				a.leftUp = a4
			case 4:
				a.leftUp = a5
			case 5:
				a.leftUp = a6
			case 6:
				a.leftUp = a7
			case 7:
				a.leftUp = a8
			}
		}
	}

	for i := 0; i < len(a.up); i++ {
		a.up[i].animSheet = a
	}
	for i := 0; i < len(a.rightUp); i++ {
		a.up[i].animSheet = a
	}
	for i := 0; i < len(a.right); i++ {
		a.up[i].animSheet = a
	}
	for i := 0; i < len(a.rightDown); i++ {
		a.up[i].animSheet = a
	}
	for i := 0; i < len(a.down); i++ {
		a.up[i].animSheet = a
	}
	for i := 0; i < len(a.leftDown); i++ {
		a.up[i].animSheet = a
	}
	for i := 0; i < len(a.left); i++ {
		a.up[i].animSheet = a
	}
	for i := 0; i < len(a.leftUp); i++ {
		a.up[i].animSheet = a
	}

	return a
}

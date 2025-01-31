package main

import (
	"fmt"
	"time"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

var (
	Win          *sdl.Window
	Rndr         *sdl.Renderer
	Err          error
	Errlist      []error
	ONOFF        = true
	WinW, WinH   int32
	Surface      *sdl.Surface
	Texture      *sdl.Texture
	ScreenBorder sdl.Rect

	WinBackgroundColor []uint8

	//SCREEN CENTER POINTS
	CNTR, CNTRTOP, CNTRBOT, CNTRLEFT, CNTRRIGHT sdl.Point
	//UNITS
	UNIT, UNIT2, UNIT3, UNIT4, UNIT5, UNIT6, UNIT7, UNIT8, UNIT9, UNIT10, UNIT11, UNIT12, UNIT13, UNIT14, UNIT15, UNIT16, UNIT17, UNIT18, UNIT19, UNIT20, UNITHALF, UNITQUARTER int32

	//FPS TIMERS
	nowDelta, lastDelta uint64
	Delta               float32
	//FPS
	targetFPS          float64
	frameDelay         = float64(1000 / targetFPS) // Delay in milliseconds
	frameStart         time.Time
	frameTime          float64
	FPS                float64
	frameCount, frames int
	startTime          = time.Now()
)

func restart() {
	terrain = nil
	lev = GRID{}

	levelܛMAKE()
}

func initial() {
	//FONTS
	fontsInit()
	fontTextureSheet()
	//IMAGES
	imagesܛMAKE()
	//LEVEL
	levelܛMAKE()
}

func WIN(width, height, baseUnitSize int32, winTitle string, backgroundcolor []uint8, setFPS int) {

	Win, Err = sdl.CreateWindow(winTitle, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		width, height, sdl.WINDOW_SHOWN|sdl.WINDOW_BORDERLESS)
	if Err != nil {
		Errlist = append(Errlist, Err)
	}
	Rndr, Err = sdl.CreateRenderer(Win, -1, sdl.RENDERER_ACCELERATED|sdl.RENDERER_PRESENTVSYNC)
	Rndr.SetDrawBlendMode(sdl.BLENDMODE_BLEND)
	if Err != nil {
		Errlist = append(Errlist, Err)
	}

	Rndr.SetDrawColor(backgroundcolor[0], backgroundcolor[1], backgroundcolor[2], backgroundcolor[3])
	Rndr.Clear()

	//CREATE TRANSPARENT SURFACE WINDOW SIZE
	if sdl.BYTEORDER == sdl.BIG_ENDIAN {
		Surface, Err = sdl.CreateRGBSurface(sdl.SWSURFACE, int32(WinW), int32(WinH), 32, 0xFF000000, 0x00FF0000, 0x0000FF00, 0x000000FF)
	} else {
		Surface, Err = sdl.CreateRGBSurface(sdl.SWSURFACE, int32(WinW), int32(WinH), 32, 0x000000FF, 0x0000FF00, 0x00FF0000, 0xFF000000)
	}
	if Err != nil {
		Errlist = append(Errlist, Err)
	}

	//FPS
	targetFPS = float64(setFPS)
	frameDelay = 1000 / targetFPS // Delay in milliseconds
	//BACKGROUND COLOR
	WinBackgroundColor = backgroundcolor
	//SCREEN BORDER REC
	ScreenBorder = sdl.Rect{0, 0, width, height}
	//SIZES
	sizes(width, height, baseUnitSize) //MAKE CENTERS/UNITS/SIZES
	//INITIAL
	initial()

}

func CORE() {

	defer Win.Destroy()
	defer Rndr.Destroy()
	defer Texture.Destroy()

	defer FontDefault.Close()
	defer Font1.Close()
	defer ttf.Quit()

	defer loadTex.Destroy()
	defer texImLoad.Destroy()

	defer Surface.Free()
	defer imLoadSurf.Free()

	if !ONOFF {

		if len(Errlist) > 0 {
			for i := 0; i < len(Errlist); i++ {
				fmt.Println(Errlist[i])
			}
		}
		if len(fontcharsDEFAULT) > 0 {
			for i := 0; i < len(fontcharsDEFAULT); i++ {
				fontcharsDEFAULT[i].tex.Destroy()
			}
		}

		Font1.Close()
		FontDefault.Close()
		ttf.Quit()

		Texture.Destroy()
		loadTex.Destroy()
		texImLoad.Destroy()

		Rndr.Destroy()
		Win.Destroy()

		sdl.Quit()

	}
}

func sizes(w, h, unit int32) {

	WinW = w
	WinH = h

	CNTR = sdl.Point{w / 2, h / 2}
	CNTRTOP = CNTR
	CNTRTOP.Y = 0
	CNTRBOT = CNTR
	CNTRBOT.Y = h
	CNTRLEFT = CNTR
	CNTRLEFT.X = 0
	CNTRRIGHT = CNTR
	CNTRRIGHT.X = w

	UNIT = unit

	UNIT2, UNIT3, UNIT4, UNIT5, UNIT6, UNIT7, UNIT8, UNIT9, UNIT10, UNIT11, UNIT12, UNIT13, UNIT14, UNIT15, UNIT16, UNIT17, UNIT18, UNIT19, UNIT20 = UNIT*2, UNIT*3, UNIT*4, UNIT*5, UNIT*6, UNIT*7, UNIT*8, UNIT*9, UNIT*10, UNIT*11, UNIT*12, UNIT*13, UNIT*14, UNIT*15, UNIT*16, UNIT*17, UNIT*18, UNIT*19, UNIT*20

	UNITHALF = UNIT / 2
	UNITQUARTER = UNIT / 4

	//RECS
	debugBGrec = sdl.Rect{0, 0, 300, WinH}
	spc := int32(12)
	closeWinRec = sdl.Rect{WinW - (closeWinRecSize + spc), spc, closeWinRecSize, closeWinRecSize}
}

func getDelta() {
	tickT := sdl.GetTicks64()
	Delta = float32(tickT-lastDelta) * 0.001
	lastDelta = tickT
}

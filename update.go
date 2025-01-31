package main

import (
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

var (
	//TIME
	Seconds int

	//MOUSE
	mouseXY                sdl.Point
	mouseL, mouseR, mouseM bool

	mouseLclickPause, mouseRclickPause, mouseMclickPause int

	//KEYS
	keyExit, keyF1, keyF2, keyUp, keyDown, keyRight, keyLeft bool

	//TIMERS
	opacFade1      = uint8(100)
	opacFade1onOff bool
)

func UPDATE() {
	sdl.ShowCursor(sdl.DISABLE)
	MOUSE()
	CLOSEWIN()
	if debugOn {
		DEBUG()
	}
	EVENTS()

	//CURSOR
	cursorRec := sdl.Rect{mouseXY.X, mouseXY.Y, cursorSize, cursorSize}
	etcIM[0].tex.SetColorMod(MAGENTAǁ3().R, MAGENTAǁ3().G, MAGENTAǁ3().B)
	etcIM[0].tex.SetAlphaMod(150)
	Rndr.Copy(etcIM[0].tex, &etcIM[0].rec, &cursorRec)
	etcIMrevert()
	//DRAW SURFACE
	Texture, _ = Rndr.CreateTextureFromSurface(Surface)
	Rndr.Copy(Texture, nil, &sdl.Rect{0, 0, int32(WinW), int32(WinH)})
	Rndr.Present()
	Surface.FillRect(&sdl.Rect{0, 0, Surface.W, Surface.H}, 0) //CLEAR SURFACE
	Texture.Destroy()
	Rndr.SetDrawColor(WinBackgroundColor[0], WinBackgroundColor[1], WinBackgroundColor[2], WinBackgroundColor[3])
	Rndr.Clear()

	AFTER()

}
func etcIMrevert() {
	etcIM[0].tex.SetColorMod(255, 255, 255)
	etcIM[0].tex.SetAlphaMod(255)
}
func CLOSEWIN() {
	if PointInRec(mouseXY, closeWinRec) {
		etcIM[1].tex.SetColorMod(255, 0, 0)
		Rndr.Copy(etcIM[1].tex, &etcIM[1].rec, &closeWinRec)
		etcIMrevert()
		if mouseL {
			ONOFF = false
		}
	} else {
		Rndr.Copy(etcIM[1].tex, &etcIM[1].rec, &closeWinRec)
	}

}
func B4DRAW() {
	frameStart = time.Now()
}
func AFTER() {

	//FPS
	frames++
	frameCount++

	frameTime = float64(time.Since(frameStart).Milliseconds())
	if frameTime < frameDelay {
		sdl.Delay(uint32(frameDelay - frameTime)) // Wait to achieve target FPS
	}
	if time.Since(startTime).Seconds() >= 1 {
		FPS = float64(frameCount) / time.Since(startTime).Seconds()
		frameCount = 0
		startTime = time.Now()
	}

	TIMERS()

}
func TIMERS() {
	if frameCount == 0 {
		Seconds++
	}
	if debugTimer > Seconds {
		debugInfoܛDISPLAY()
	}

	if mouseLclickPause > 0 {
		mouseLclickPause--
	}
	if mouseRclickPause > 0 {
		mouseRclickPause--
	}
	if mouseMclickPause > 0 {
		mouseMclickPause--
	}

	//COLOR FADE
	if opacFade1onOff {
		if opacFade1 > 50 {
			opacFade1 -= 4
		} else {
			opacFade1onOff = false
		}
	} else {
		if opacFade1 < 150 {
			opacFade1 += 4
		} else {
			opacFade1onOff = true
		}
	}
}

// EVENTS
func EVENTS() {

	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch t := event.(type) {
		case *sdl.KeyboardEvent:
			if event.GetType() == sdl.KEYDOWN {
				switch t.Keysym.Scancode {
				case sdl.SCANCODE_ESCAPE:
					keyExit = true
				case sdl.SCANCODE_F1:
					keyF1 = true
				case sdl.SCANCODE_F2:
					keyF2 = true
				case sdl.SCANCODE_A, sdl.SCANCODE_LEFT:
					keyLeft = true

				case sdl.SCANCODE_W, sdl.SCANCODE_UP:
					keyUp = true

				case sdl.SCANCODE_D, sdl.SCANCODE_RIGHT:
					keyRight = true

				case sdl.SCANCODE_S, sdl.SCANCODE_DOWN:
					keyDown = true

				}
			} else if event.GetType() == sdl.KEYUP {
				switch t.Keysym.Scancode {
				case sdl.SCANCODE_A, sdl.SCANCODE_LEFT:
					keyLeft = false

				case sdl.SCANCODE_W, sdl.SCANCODE_UP:
					keyUp = false

				case sdl.SCANCODE_D, sdl.SCANCODE_RIGHT:
					keyRight = false

				case sdl.SCANCODE_S, sdl.SCANCODE_DOWN:
					keyDown = false

				}
			}
		}
	}
	KEYS()

}

// INPUT
func MOUSE() {

	x, y, click := sdl.GetMouseState()
	mouseXY.X, mouseXY.Y = x, y

	if click == sdl.ButtonLMask() && mouseLclickPause == 0 {
		mouseLclickPause = int(FPS) / 3
		mouseL = true
	} else {
		mouseL = false
	}

	if click == sdl.ButtonRMask() && mouseRclickPause == 0 {
		mouseRclickPause = int(FPS) / 3
		mouseR = true
	} else {
		mouseR = false
	}

	if click == sdl.ButtonMMask() && mouseRclickPause == 0 {
		mouseMclickPause = int(FPS) / 3
		mouseM = true
	} else {
		mouseM = false
	}

}
func KEYS() {
	if keyExit { //ESCAPE
		ONOFF, keyExit = false, false
	}
	if keyF1 { //F1
		debugOn = !debugOn
		keyF1 = false
	}
	if keyF2 { //F2
		restart()
		keyF2 = false
	}

}

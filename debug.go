package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

var (
	debugOn    bool
	debugInfo  []string
	debugTimer int
	debugBGrec sdl.Rect

	intDEBUG int
)

// OBJ OBJ OBJ OBJ OBJ OBJ OBJ OBJ OBJ OBJ OBJ OBJ OBJ OBJ OBJ OBJ OBJ OBJ OBJ OBJ
func ObjCollideܛCHECK(o1, o2 OBJ) bool {
	return o1.rec.X < o2.rec.X+o2.rec.W && o1.rec.X+o1.rec.W > o2.rec.X && o1.rec.Y < o2.rec.Y+o2.rec.H && o1.rec.Y+o1.rec.H > o2.rec.Y
}
func ObjAddToListܛCHECK(o OBJ, o2 []OBJ) bool {
	canadd := true
	if len(o2) > 0 {
		for i := 0; i < len(o2); i++ {
			if ObjCollideܛCHECK(o, o2[i]) {
				canadd = false
			}
			if !canadd {
				break
			}
		}
	}
	return canadd
}

// DEBUG DEBUG DEBUG DEBUG DEBUG DEBUG DEBUG DEBUG DEBUG DEBUG DEBUG DEBUG DEBUG
func DEBUG() {
	RecNormalAlphaܛDRAW(debugBGrec, MAROON(), 150)
	var x, y int32 = 10, 10
	TextHereXYܛTEXT("FPS "+fmt.Sprintf("%.0f", FPS)+" seconds "+fmt.Sprint(Seconds), x, y)
	y += LineHeight
	TextHereXYܛTEXT("mouseX "+fmt.Sprint(mouseXY.X)+" mouseY "+fmt.Sprint(mouseXY.X), x, y)
	y += LineHeight
	TextHereXYܛTEXT("keyLeft "+fmt.Sprint(keyLeft)+" keyRight "+fmt.Sprint(keyRight), x, y)
	y += LineHeight
	TextHereXYܛTEXT("keyUp "+fmt.Sprint(keyUp)+" keyDown "+fmt.Sprint(keyDown), x, y)
	y += LineHeight
	TextHereXYܛTEXT("intDEBUG "+fmt.Sprint(intDEBUG)+" len(cards) "+fmt.Sprint(len(cards)), x, y)
	y += LineHeight
	TextHereXYܛTEXT("pl1.bloknum "+fmt.Sprint(pl1.bloknum)+" len(cards) "+fmt.Sprint(len(cards)), x, y)
	y += LineHeight

}
func debugInfoܛDISPLAY() {
	TextHereXYܛTEXT(debugInfo[len(debugInfo)-1], CNTR.X, 10)
}
func debugInfoܛADD(t string) {
	debugTimer = Seconds + 3
	debugInfo = append(debugInfo, t)
}

// CHECKS CHECKS CHECKS CHECKS CHECKS CHECKS CHECKS CHECKS CHECKS CHECKS CHECKS
func MouseBlockܛCHECK(g GRID) int {
	num := 0
	for i := 0; i < len(g.bloks); i++ {
		for j := 0; j < len(g.bloks[i].recs); j++ {
			if PointInTriangleܛCHECK(mouseXY, g.bloks[i].recs[0].triangles[0]) || PointInTriangleܛCHECK(mouseXY, g.bloks[i].recs[0].triangles[1]) {
				num = i
				break
			}
		}
	}
	return num
}
func MouseRecܛCHECK(r sdl.Rect) bool {
	return PointInRec(mouseXY, r)
}
func PointInTriangleܛCHECK(p sdl.Point, t TRI) bool {
	as_x := p.X - t.vertPoints[0].X
	as_y := p.Y - t.vertPoints[0].Y
	s_ab := (t.vertPoints[1].X-t.vertPoints[0].X)*as_y-(t.vertPoints[1].Y-t.vertPoints[0].Y)*as_x > 0
	if (t.vertPoints[2].X-t.vertPoints[0].X)*as_y-(t.vertPoints[2].Y-t.vertPoints[0].Y)*as_x > 0 == s_ab {
		return false
	}
	if (t.vertPoints[2].X-t.vertPoints[1].X)*(p.Y-t.vertPoints[1].Y)-(t.vertPoints[2].Y-t.vertPoints[1].Y)*(p.X-t.vertPoints[1].X) > 0 != s_ab {
		return false
	}
	return true
}
func ColorsMatchܛCHECK(color1, color2 []uint8) bool {
	if len(color1) != len(color2) {
		return false
	}
	for i := range color1 {
		if color1[i] != color2[i] {
			return false
		}
	}
	return true
}

/*

func ColorsMatchܛCHECK(color1, color2 []uint8) bool {
	if len(color1) != len(color2) {
		return false
	}
	for i := range color1 {
		if color1[i] != color2[i] {
			return false
		}
	}
	return true
}


func ChkܢPointInRectangle(p sdl.Point, r sdl.Rect) bool {
	return p.X >= r.X && p.X <= r.X+r.W && p.Y >= r.Y && p.Y <= r.Y+r.H
}

func ChkܢGridBlokMouse(g GridIso) int {
	num := 0
	for i := 0; i < len(g.recs); i++ {
		if ChkܢPointInTriangle(mouseXY, g.recs[i].triangles[0]) || ChkܢPointInTriangle(mouseXY, g.recs[i].triangles[1]) {
			num = i
			break
		}
	}
	return num
}
func ChkܢGridBlokMouseDrawColor(g GridIso, color sdl.Color) int {
	num := 0
	for i := 0; i < len(g.recs); i++ {
		if ChkܢPointInTriangle(mouseXY, g.recs[i].triangles[0]) || ChkܢPointInTriangle(mouseXY, g.recs[i].triangles[1]) {
			DrܛIsoRecSolidColor(g.recs[i], color)
			num = i
			break
		}
	}
	return num
}

func ChkܢPointInTriangle(p sdl.Point, t Tri) bool {
	as_x := p.X - t.vertPoints[0].X
	as_y := p.Y - t.vertPoints[0].Y
	s_ab := (t.vertPoints[1].X-t.vertPoints[0].X)*as_y-(t.vertPoints[1].Y-t.vertPoints[0].Y)*as_x > 0
	if (t.vertPoints[2].X-t.vertPoints[0].X)*as_y-(t.vertPoints[2].Y-t.vertPoints[0].Y)*as_x > 0 == s_ab {
		return false
	}
	if (t.vertPoints[2].X-t.vertPoints[1].X)*(p.Y-t.vertPoints[1].Y)-(t.vertPoints[2].Y-t.vertPoints[1].Y)*(p.X-t.vertPoints[1].X) > 0 != s_ab {
		return false
	}
	return true
}
func ChkܢRGBAColorsMatch(color1, color2 []uint8) bool {
	if len(color1) != len(color2) {
		return false
	}
	for i := range color1 {
		if color1[i] != color2[i] {
			return false
		}
	}
	return true
}
*/

package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

// PLAYERS PLAYERS PLAYERS PLAYERS PLAYERS PLAYERS PLAYERS PLAYERS PLAYERS
func PlayersܛDRAW() {
	num := 0
	for i := 0; i < len(lev.bloks); i++ {
		if lev.bloks[i].num == pl1.bloknum {
			num = i
			break
		}
	}
	pl1.zIndex = lev.bloks[num].zIndex
	play1Anim = AnimTopBlokܛDRAW(play1Anim, lev.bloks[num].cntr.X, lev.bloks[num].cntr.Y, 3, 15)
}

// ANIM ANIM ANIM ANIM ANIM ANIM ANIM ANIM ANIM ANIM ANIM ANIM ANIM
func AnimܛDRAW(a ANIM, xCnt, yCnt, zoom int32, spd int) ANIM {
	w := a.im.rec.W * zoom
	h := a.im.rec.H * zoom
	r := sdl.Rect{xCnt - w/2, yCnt - h/2, w, h}
	Rndr.Copy(a.im.tex, &a.im.rec, &r)
	if frames%spd == 0 {
		a.framecount++
		a.im.rec.X += a.im.rec.W
		if a.framecount == a.frames {
			a.im.rec.X = a.xStart
			a.framecount = 0
		}
	}
	return a
}
func AnimTopBlokܛDRAW(a ANIM, xCnt, yCnt, zoom int32, spd int) ANIM {
	w := a.im.rec.W * zoom
	h := a.im.rec.H * zoom
	r := sdl.Rect{xCnt - w/2, yCnt - h/2, w, h}
	r.Y -= h / 2
	Rndr.Copy(a.im.tex, &a.im.rec, &r)
	if frames%spd == 0 {
		a.framecount++
		a.im.rec.X += a.im.rec.W
		if a.framecount == a.frames {
			a.im.rec.X = a.xStart
			a.framecount = 0
		}
	}
	return a
}

// IM IM IM IM IM IM IM IM IM IM IM IM IM IM IM IM IM IM IM IM IM IM IM
func IMܛDRAW(i IM, x, y, zoom int32) {
	r := sdl.Rect{x, y, i.rec.W * zoom, i.rec.H * zoom}
	Rndr.Copy(i.tex, &i.rec, &r)
}
func IMShadowܛDRAW(i IM, x, y, zoom int32) {
	i.tex.SetColorMod(70, 70, 70)
	i.tex.SetAlphaMod(120)
	r := sdl.Rect{x, y, i.rec.W * zoom, i.rec.H * zoom}
	r2 := r
	r2.X -= 7
	r2.Y += 7
	Rndr.Copy(i.tex, &i.rec, &r2)
	i.tex.SetColorMod(255, 255, 255)
	i.tex.SetAlphaMod(255)
	Rndr.Copy(i.tex, &i.rec, &r)
}

// TILESHEET TILESHEET TILESHEET TILESHEET TILESHEET TILESHEET TILESHEET
func TileSheetܛDRAW(tileSheet TILESHEET, x, y, space, zoom int32) {

	ox := x
	for i := 0; i < len(tileSheet.tileIM); i++ {
		drec := sdl.Rect{x, y, tileSheet.tileIM[i].rec.W * zoom, tileSheet.tileIM[i].rec.H * zoom}
		Rndr.Copy(tileSheet.tex, &tileSheet.tileIM[i].rec, &drec)
		TextHereXYܛTEXT(fmt.Sprint(i), x, y)
		x += tileSheet.tileIM[i].rec.W*zoom + space
		if x > WinW-(tileSheet.tileIM[i].rec.W*zoom+space) {
			x = ox
			y += tileSheet.tileIM[i].rec.H*zoom + space
		}
	}
	TextHereXYܛTEXT(tileSheet.name, x, y+tileSheet.heightTile/2)

}

// CARDS CARDS CARDS CARDS CARDS CARDS CARDS CARDS CARDS CARDS CARDS CARDS
func CardListܛDRAW(c []CARD, x, y, spc, zoom int32, zoomNEG bool) {
	var newW, newH int32
	if zoomNEG {
		newW = c[0].im.rec.W / zoom
		newH = c[0].im.rec.H / zoom
	} else {
		newW = c[0].im.rec.W * zoom
		newH = c[0].im.rec.H * zoom
	}

	for i := 0; i < len(c); i++ {
		cardRec := sdl.Rect{x, y, newW, newH}
		if MouseRecܛCHECK(cardRec) {
			if mouseL && mouseLclickPause == 0 {
				c[i].selected = !c[i].selected
			}

			Rndr.SetDrawColor(colorOpacity(DARKMAGENTA(), opacFade1))
			Rndr.FillRect(&cardRec)
		}

		Rndr.Copy(c[i].im.tex, &c[i].im.rec, &cardRec)

		//NAME TEXT
		cnt := sdl.Point{x + newW/2, y + newH/2}
		txcnt := cnt
		txcnt.Y -= (newH / 2) - (newH / 8)
		TextHereCenterColorܛTEXT(c[i].nm, txcnt, BLACK(), 1)
		//MAIN ICON
		siz := c[i].iconMAIN.rec.W * 2
		c[i].recIcon = sdl.Rect{x + 20, y + 50, siz, siz}
		Rndr.Copy(c[i].iconMAIN.tex, &c[i].iconMAIN.rec, &c[i].recIcon)
		c[i].recIcon.Y += c[i].recIcon.H + 4
		Rndr.Copy(c[i].iconMAIN.tex, &c[i].iconMAIN.rec, &c[i].recIcon)
		c[i].recIcon.Y += c[i].recIcon.H + 4
		Rndr.Copy(c[i].iconMAIN.tex, &c[i].iconMAIN.rec, &c[i].recIcon)
		x += newW + spc
	}
}
func CardSingleܛDRAW(c CARD, x, y, zoom int32, zoomNEG bool) {
	var newW, newH int32
	if zoomNEG {
		newW = c.im.rec.W / zoom
		newH = c.im.rec.H / zoom
	} else {
		newW = c.im.rec.W * zoom
		newH = c.im.rec.H * zoom
	}
	Rndr.Copy(c.im.tex, &c.im.rec, &sdl.Rect{x, y, newW, newH})
}

// OBJ OBJ OBJ OBJ OBJ OBJ OBJ OBJ OBJ OBJ OBJ OBJ OBJ OBJ OBJ OBJ OBJ OBJ
func ObjܛDRAW(o []OBJ) {
	for i := 0; i < len(o); i++ {
		Rndr.Copy(o[i].im.tileSheet.tex, &o[i].im.rec, &o[i].rec)
	}
}
func ObjLinesܛDRAW(o []OBJ, color []uint8) {
	Rndr.SetDrawColor(color[0], color[1], color[2], color[3])
	for i := 0; i < len(o); i++ {
		Rndr.DrawRect(&o[i].rec)
	}
}

// GRID GRID GRID GRID GRID GRID GRID GRID GRID GRID GRID GRID GRID GRID
func GridBlokCentersܛDRAW(grid GRID) {
	for i := 0; i < len(grid.bloks); i++ {
		Rndr.SetDrawColor(BLUEǁ2())
		r := sdl.Rect{grid.bloks[i].cntr.X - 4, grid.bloks[i].cntr.Y - 4, 8, 8}
		Rndr.FillRect(&r)
	}
}
func GridܛDRAW(grid GRID) {
	for i := 0; i < len(grid.bloks); i++ {
		GridRecܛDRAW(grid.bloks[i])
	}
}
func GridLinesOnlyܛDRAW(grid GRID, color []uint8) {
	b := visibleBloks(grid)
	for i := 0; i < len(b); i++ {
		if !b[i].onoff {
			GridRecLinesܛDRAW(b[i], color)
		}
	}
}
func GridTilesܛDRAW(grid GRID) {
	b := visibleBloks(grid)
	for i := 0; i < len(b); i++ {
		if !b[i].onoff {
			Rndr.Copy(b[i].tileIM.tileSheet.tex, &b[i].tileIM.rec, &b[i].dRec)
		}
	}
}
func GridTilesMouseBlokܛDRAW(grid GRID) {
	b := visibleBloks(grid)
	mouseNum := grid.bloks[MouseBlockܛCHECK(grid)].num
	for i := 0; i < len(b); i++ {
		if !b[i].onoff {
			Rndr.Copy(b[i].tileIM.tileSheet.tex, &b[i].tileIM.rec, &b[i].dRec)
		}
		if b[i].num == mouseNum {
			GridRecLinesܛDRAW(b[i], MAGENTA())
		}
	}
}
func visibleBloks(grid GRID) []BLOK {
	var b []BLOK
	for i := 0; i < len(grid.bloks); i++ {
		if PointInRec(grid.bloks[i].floorCorners[0], ScreenBorder) || PointInRec(grid.bloks[i].floorCorners[1], ScreenBorder) || PointInRec(grid.bloks[i].floorCorners[2], ScreenBorder) || PointInRec(grid.bloks[i].floorCorners[3], ScreenBorder) {
			b = append(b, grid.bloks[i])
		}
	}
	return b
}

// RECS RECS RECS RECS RECS RECS RECS RECS RECS RECS RECS RECS RECS RECS
func GridRecܛDRAW(blok BLOK) {
	TrianglesܛDRAW(blok.recs[0].triangles)
	LineܛDRAW(blok.recs[0].triangles[0].vertPoints[0], blok.recs[0].triangles[0].vertPoints[1], RED())
	LineܛDRAW(blok.recs[0].triangles[0].vertPoints[0], blok.recs[0].triangles[1].vertPoints[1], RED())
	TextHereCenterܛTEXT(fmt.Sprint(blok.num), blok.cntr, 0)
}
func GridRecLinesܛDRAW(blok BLOK, color []uint8) {
	LineܛDRAW(blok.recs[0].triangles[0].vertPoints[0], blok.recs[0].triangles[0].vertPoints[1], color)
	LineܛDRAW(blok.recs[0].triangles[0].vertPoints[0], blok.recs[0].triangles[1].vertPoints[1], color)
	TextHereCenterܛTEXT(fmt.Sprint(blok.num), blok.cntr, 0)
}
func GridRecܛDRAWNoOutlineܛDRAW(rec REC) {
	TrianglesܛDRAW(rec.triangles)
}

// RECS NORMAL RECS NORMAL RECS NORMAL RECS NORMAL RECS NORMAL RECS NORMAL RECS NORMAL
func RecNormalAlphaܛDRAW(rec sdl.Rect, color []uint8, opacity uint8) {
	color[3] = opacity
	Rndr.SetDrawColor(color[0], color[1], color[2], color[3])
	Rndr.FillRect(&rec)
}

// BLOKS BLOKS BLOKS BLOKS BLOKS BLOKS BLOKS BLOKS BLOKS BLOKS BLOKS BLOKS
func BlokܛDRAW(blok BLOK) {
	TrianglesܛDRAW(blok.recs[1].triangles)
	TrianglesܛDRAW(blok.recs[2].triangles)
	TrianglesܛDRAW(blok.recs[3].triangles)
	BlokOutlineܛDRAW(blok)
}
func BlokNoOutlineܛDRAW(blok BLOK) {
	TrianglesܛDRAW(blok.recs[1].triangles)
	TrianglesܛDRAW(blok.recs[2].triangles)
	TrianglesܛDRAW(blok.recs[3].triangles)
}

func BlokOutlineܛDRAW(blok BLOK) {
	LineܛDRAW(blok.frontPoints[0], blok.frontPoints[1], blok.colorLines)
	LineܛDRAW(blok.frontPoints[0], blok.frontPoints[2], blok.colorLines)
	LineܛDRAW(blok.frontPoints[3], blok.frontPoints[4], blok.colorLines)
	LineܛDRAW(blok.frontPoints[4], blok.frontPoints[5], blok.colorLines)
	LineܛDRAW(blok.frontPoints[5], blok.frontPoints[6], blok.colorLines)
	LineܛDRAW(blok.frontPoints[3], blok.frontPoints[6], blok.colorLines)
	LineܛDRAW(blok.frontPoints[0], blok.frontPoints[3], blok.colorLines)
	LineܛDRAW(blok.frontPoints[1], blok.frontPoints[4], blok.colorLines)
	LineܛDRAW(blok.frontPoints[2], blok.frontPoints[6], blok.colorLines)
}
func BlokOutlineOnlyColorܛDRAW(blok BLOK, color []uint8) {
	LineܛDRAW(blok.frontPoints[0], blok.frontPoints[1], color)
	LineܛDRAW(blok.frontPoints[0], blok.frontPoints[2], color)
	LineܛDRAW(blok.frontPoints[3], blok.frontPoints[4], color)
	LineܛDRAW(blok.frontPoints[4], blok.frontPoints[5], color)
	LineܛDRAW(blok.frontPoints[5], blok.frontPoints[6], color)
	LineܛDRAW(blok.frontPoints[3], blok.frontPoints[6], color)
	LineܛDRAW(blok.frontPoints[0], blok.frontPoints[3], color)
	LineܛDRAW(blok.frontPoints[1], blok.frontPoints[4], color)
	LineܛDRAW(blok.frontPoints[2], blok.frontPoints[6], color)
}

// TRIANGLES  TRIANGLES TRIANGLES TRIANGLES TRIANGLES TRIANGLES TRIANGLES
func TrianglesܛDRAW(triangles []TRI) {
	var v []sdl.Vertex
	for i := 0; i < len(triangles); i++ {
		v = append(v, triangles[i].vertices...)
	}
	Rndr.RenderGeometry(nil, v, nil)
}
func TrianglesColorܛDRAW(triangles []TRI, color sdl.Color) {
	var v []sdl.Vertex
	for i := 0; i < len(triangles); i++ {
		for j := 0; j < len(triangles[i].vertices); j++ {
			triangles[i].vertices[j].Color = color
		}

		v = append(v, triangles[i].vertices...)
	}
	Rndr.RenderGeometry(nil, v, nil)
}

// LINES LINES LINES LINES LINES LINES LINES LINES LINES LINES LINES LINES
func LinesܛDRAW(points []sdl.Point, color []uint8) {
	Rndr.SetDrawColor(color[0], color[1], color[2], color[3])
	Rndr.DrawLines(points)
}
func LineXYܛDRAW(x1, y1, x2, y2 int32, color []uint8) {
	Rndr.SetDrawColor(color[0], color[1], color[2], color[3])
	Rndr.DrawLine(x1, y1, x2, y2)
}
func LineܛDRAW(p1, p2 sdl.Point, color []uint8) {
	Rndr.SetDrawColor(color[0], color[1], color[2], color[3])
	Rndr.DrawLine(p1.X, p1.Y, p2.X, p2.Y)
}

// ANIMS ANIMS ANIMS ANIMS ANIMS ANIMS ANIMS ANIMS ANIMS ANIMS ANIMS ANIMS
/*
func PlayerܛDRAW(pl PLAYR, idleSpeed, runSpeed int) {

	//DRAW

	switch pl.state {
	case 1: //RUN
		switch pl.direc {
		case 1: //UP
			Rndr.Copy(pl.runAnim.tex, &pl.runAnim.up[0].rec, &pl.drawRec)
		case 2: //RIGHTUP
			Rndr.Copy(pl.runAnim.tex, &pl.runAnim.rightUp[0].rec, &pl.drawRec)
		case 3: //RIGHT
			Rndr.Copy(pl.runAnim.tex, &pl.runAnim.right[0].rec, &pl.drawRec)
		case 4: //RIGHTDOWN
			Rndr.Copy(pl.runAnim.tex, &pl.runAnim.rightDown[0].rec, &pl.drawRec)
		case 5: //DOWN
			Rndr.Copy(pl.runAnim.tex, &pl.runAnim.down[0].rec, &pl.drawRec)
		case 6: //LEFTDOWN
			Rndr.Copy(pl.runAnim.tex, &pl.runAnim.leftDown[0].rec, &pl.drawRec)
		case 7: //LEFT
			Rndr.Copy(pl.runAnim.tex, &pl.runAnim.left[0].rec, &pl.drawRec)
		case 8: //LEFTUP
			Rndr.Copy(pl.runAnim.tex, &pl.runAnim.leftUp[0].rec, &pl.drawRec)
		}
	case 0: //IDLE
		switch pl.direc {
		case 1: //UP
			Rndr.Copy(pl.idleAnim.tex, &pl.idleAnim.up[0].rec, &pl.drawRec)
		case 2: //RIGHTUP
			Rndr.Copy(pl.idleAnim.tex, &pl.idleAnim.rightUp[0].rec, &pl.drawRec)
		case 3: //RIGHT
			Rndr.Copy(pl.idleAnim.tex, &pl.idleAnim.right[0].rec, &pl.drawRec)
		case 4: //RIGHTDOWN
			Rndr.Copy(pl.idleAnim.tex, &pl.idleAnim.rightDown[0].rec, &pl.drawRec)
		case 5: //DOWN
			Rndr.Copy(pl.idleAnim.tex, &pl.idleAnim.down[0].rec, &pl.drawRec)
		case 6: //LEFTDOWN
			Rndr.Copy(pl.idleAnim.tex, &pl.idleAnim.leftDown[0].rec, &pl.drawRec)
		case 7: //LEFT
			Rndr.Copy(pl.idleAnim.tex, &pl.idleAnim.left[0].rec, &pl.drawRec)
		case 8: //LEFTUP
			Rndr.Copy(pl.idleAnim.tex, &pl.idleAnim.leftUp[0].rec, &pl.drawRec)
		}
	}

	//UP ANIMS
	switch pl.state {
	case 1: //RUN
		if frameCount%idleSpeed == 0 {
			switch pl.direc {
			case 1: //UP
				pl.runAnim.up[0].rec.X += pl.runAnim.up[0].rec.W
				pl.runAnim.up[0].frameCount++
				if pl.runAnim.up[0].frameCount >= pl.runAnim.up[0].numTiles {
					pl.runAnim.up[0].rec.X = pl.runAnim.up[0].xStart
					pl.runAnim.up[0].frameCount = 0
				}
			case 2: //RIGHTUP
				pl.runAnim.rightUp[0].rec.X += pl.runAnim.rightUp[0].rec.W
				pl.runAnim.rightUp[0].frameCount++
				if pl.runAnim.rightUp[0].frameCount >= pl.runAnim.rightUp[0].numTiles {
					pl.runAnim.rightUp[0].rec.X = pl.runAnim.rightUp[0].xStart
					pl.runAnim.rightUp[0].frameCount = 0
				}
			case 3: //RIGHT
				pl.runAnim.right[0].rec.X += pl.runAnim.right[0].rec.W
				pl.runAnim.right[0].frameCount++
				if pl.runAnim.right[0].frameCount >= pl.runAnim.right[0].numTiles {
					pl.runAnim.right[0].rec.X = pl.runAnim.right[0].xStart
					pl.runAnim.right[0].frameCount = 0
				}
			case 4: //RIGHTDOWN
				pl.runAnim.rightDown[0].rec.X += pl.runAnim.rightDown[0].rec.W
				pl.runAnim.rightDown[0].frameCount++
				if pl.runAnim.rightDown[0].frameCount >= pl.runAnim.rightDown[0].numTiles {
					pl.runAnim.rightDown[0].rec.X = pl.runAnim.rightDown[0].xStart
					pl.runAnim.rightDown[0].frameCount = 0
				}
			case 5: //DOWN
				pl.runAnim.down[0].rec.X += pl.runAnim.down[0].rec.W
				pl.runAnim.down[0].frameCount++
				if pl.runAnim.down[0].frameCount >= pl.runAnim.down[0].numTiles {
					pl.runAnim.down[0].rec.X = pl.runAnim.down[0].xStart
					pl.runAnim.down[0].frameCount = 0
				}
			case 6: //LEFTDOWN
				pl.runAnim.leftDown[0].rec.X += pl.runAnim.leftDown[0].rec.W
				pl.runAnim.leftDown[0].frameCount++
				if pl.runAnim.leftDown[0].frameCount >= pl.runAnim.leftDown[0].numTiles {
					pl.runAnim.leftDown[0].rec.X = pl.runAnim.leftDown[0].xStart
					pl.runAnim.leftDown[0].frameCount = 0
				}
			case 7: //LEFT
				pl.runAnim.left[0].rec.X += pl.runAnim.left[0].rec.W
				pl.runAnim.left[0].frameCount++
				if pl.runAnim.left[0].frameCount >= pl.runAnim.left[0].numTiles {
					pl.runAnim.left[0].rec.X = pl.runAnim.left[0].xStart
					pl.runAnim.left[0].frameCount = 0
				}
			case 8: //LEFTUP
				pl.runAnim.leftUp[0].rec.X += pl.runAnim.leftUp[0].rec.W
				pl.runAnim.leftUp[0].frameCount++
				if pl.runAnim.leftUp[0].frameCount >= pl.runAnim.leftUp[0].numTiles {
					pl.runAnim.leftUp[0].rec.X = pl.runAnim.leftUp[0].xStart
					pl.runAnim.leftUp[0].frameCount = 0
				}
			}
		}
	case 0: //IDLE
		if frameCount%idleSpeed == 0 {
			switch pl.direc {
			case 1: //UP
				pl.idleAnim.up[0].rec.X += pl.idleAnim.up[0].rec.W
				pl.idleAnim.up[0].frameCount++
				if pl.idleAnim.up[0].frameCount >= pl.idleAnim.up[0].numTiles {
					pl.idleAnim.up[0].rec.X = pl.idleAnim.up[0].xStart
					pl.idleAnim.up[0].frameCount = 0
				}
			case 2: //RIGHTUP
				pl.idleAnim.rightUp[0].rec.X += pl.idleAnim.rightUp[0].rec.W
				pl.idleAnim.rightUp[0].frameCount++
				if pl.idleAnim.rightUp[0].frameCount >= pl.idleAnim.rightUp[0].numTiles {
					pl.idleAnim.rightUp[0].rec.X = pl.idleAnim.rightUp[0].xStart
					pl.idleAnim.rightUp[0].frameCount = 0
				}
			case 3: //RIGHT
				pl.idleAnim.right[0].rec.X += pl.idleAnim.right[0].rec.W
				pl.idleAnim.right[0].frameCount++
				if pl.idleAnim.right[0].frameCount >= pl.idleAnim.right[0].numTiles {
					pl.idleAnim.right[0].rec.X = pl.idleAnim.right[0].xStart
					pl.idleAnim.right[0].frameCount = 0
				}
			case 4: //RIGHTDOWN
				pl.idleAnim.rightDown[0].rec.X += pl.idleAnim.rightDown[0].rec.W
				pl.idleAnim.rightDown[0].frameCount++
				if pl.idleAnim.rightDown[0].frameCount >= pl.idleAnim.rightDown[0].numTiles {
					pl.idleAnim.rightDown[0].rec.X = pl.idleAnim.rightDown[0].xStart
					pl.idleAnim.rightDown[0].frameCount = 0
				}
			case 5: //DOWN
				pl.idleAnim.down[0].rec.X += pl.idleAnim.down[0].rec.W
				pl.idleAnim.down[0].frameCount++
				if pl.idleAnim.down[0].frameCount >= pl.idleAnim.down[0].numTiles {
					pl.idleAnim.down[0].rec.X = pl.idleAnim.down[0].xStart
					pl.idleAnim.down[0].frameCount = 0
				}
			case 6: //LEFTDOWN
				pl.idleAnim.leftDown[0].rec.X += pl.idleAnim.leftDown[0].rec.W
				pl.idleAnim.leftDown[0].frameCount++
				if pl.idleAnim.leftDown[0].frameCount >= pl.idleAnim.leftDown[0].numTiles {
					pl.idleAnim.leftDown[0].rec.X = pl.idleAnim.leftDown[0].xStart
					pl.idleAnim.leftDown[0].frameCount = 0
				}
			case 7: //LEFT
				pl.idleAnim.left[0].rec.X += pl.idleAnim.left[0].rec.W
				pl.idleAnim.left[0].frameCount++
				if pl.idleAnim.left[0].frameCount >= pl.idleAnim.left[0].numTiles {
					pl.idleAnim.left[0].rec.X = pl.idleAnim.left[0].xStart
					pl.idleAnim.left[0].frameCount = 0
				}
			case 8: //LEFTUP
				pl.idleAnim.leftUp[0].rec.X += pl.idleAnim.leftUp[0].rec.W
				pl.idleAnim.leftUp[0].frameCount++
				if pl.idleAnim.leftUp[0].frameCount >= pl.idleAnim.leftUp[0].numTiles {
					pl.idleAnim.leftUp[0].rec.X = pl.idleAnim.leftUp[0].xStart
					pl.idleAnim.leftUp[0].frameCount = 0
				}
			}
		}
	}

}
*/

//GRIDS GRIDS GRIDS GRIDS GRIDS GRIDS GRIDS GRIDS GRIDS GRIDS GRIDS GRIDS
/**
func GridܛDRAW(grid Grid) {

	for i := 0; i < len(grid.bloks); i++ {
		IsoRecSolidܛDRAW(grid.bloks[i].recs[0])
	}

}
func GridBloksܛDRAW(grid Grid) {

	for i := 0; i < len(grid.bloks); i++ {
		BlokܛDRAW(grid.bloks[i])
	}

}

//BLOKS
func DrܛIsoBlokShaded(blok BlokIso, darkTop, darkRight bool, shadowAmount uint8) {
	TrianglesSolidܛDRAW(blok.visibleTriangles)

	if shadowAmount > 255 {
		shadowAmount = 255
	}
	if shadowAmount < 0 {
		shadowAmount = 0
	}

	if darkTop && darkRight {
		newTri := blok.triTop
		for i := 0; i < len(newTri); i++ {
			for j := 0; j < len(newTri[i].vertices); j++ {
				newTri[i].vertices[j].Color = BLACKǁ3()
				newTri[i].vertices[j].Color.A = shadowAmount
			}
		}
		TrianglesSolidܛDRAW(newTri)
		shadowAmount -= shadowAmount / 4
		newTri = blok.triRightFront
		for i := 0; i < len(newTri); i++ {
			for j := 0; j < len(newTri[i].vertices); j++ {
				newTri[i].vertices[j].Color = BLACKǁ3()
				newTri[i].vertices[j].Color.A = shadowAmount
			}
		}
		TrianglesSolidܛDRAW(newTri)
	} else if darkTop && !darkRight {
		newTri := blok.triTop
		for i := 0; i < len(newTri); i++ {
			for j := 0; j < len(newTri[i].vertices); j++ {
				newTri[i].vertices[j].Color = BLACKǁ3()
				newTri[i].vertices[j].Color.A = shadowAmount
			}
		}
		TrianglesSolidܛDRAW(newTri)
		shadowAmount -= shadowAmount / 4
		newTri = blok.triLeftFront
		for i := 0; i < len(newTri); i++ {
			for j := 0; j < len(newTri[i].vertices); j++ {
				newTri[i].vertices[j].Color = BLACKǁ3()
				newTri[i].vertices[j].Color.A = shadowAmount
			}
		}
		TrianglesSolidܛDRAW(newTri)
	} else if !darkTop && !darkRight {
		newTri := blok.triLeftFront
		for i := 0; i < len(newTri); i++ {
			for j := 0; j < len(newTri[i].vertices); j++ {
				newTri[i].vertices[j].Color = BLACKǁ3()
				newTri[i].vertices[j].Color.A = shadowAmount
			}
		}
		TrianglesSolidܛDRAW(newTri)
		shadowAmount -= shadowAmount / 4
		newTri = blok.triRightFront
		for i := 0; i < len(newTri); i++ {
			for j := 0; j < len(newTri[i].vertices); j++ {
				newTri[i].vertices[j].Color = BLACKǁ3()
				newTri[i].vertices[j].Color.A = shadowAmount
			}
		}
		TrianglesSolidܛDRAW(newTri)
	} else if !darkTop && darkRight {
		newTri := blok.triRightFront
		for i := 0; i < len(newTri); i++ {
			for j := 0; j < len(newTri[i].vertices); j++ {
				newTri[i].vertices[j].Color = BLACKǁ3()
				newTri[i].vertices[j].Color.A = shadowAmount
			}
		}
		TrianglesSolidܛDRAW(newTri)
		shadowAmount -= shadowAmount / 4
		newTri = blok.triLeftFront
		for i := 0; i < len(newTri); i++ {
			for j := 0; j < len(newTri[i].vertices); j++ {
				newTri[i].vertices[j].Color = BLACKǁ3()
				newTri[i].vertices[j].Color.A = shadowAmount
			}
		}
		TrianglesSolidܛDRAW(newTri)
	}

	color := BLACK()
	if ColorsMatchܛCHECK(color, blok.colorSidesRGBA) {
		color = WHITE()
	}
	LineܛDRAW(blok.pointsBottom[0], blok.pointsTop[0], color)
	LineܛDRAW(blok.pointsBottom[1], blok.pointsTop[1], color)
	LineܛDRAW(blok.pointsBottom[3], blok.pointsTop[3], color)

	LineܛDRAW(blok.pointsBottom[0], blok.pointsBottom[1], color)
	LineܛDRAW(blok.pointsBottom[0], blok.pointsBottom[3], color)

	IsoRecOutlineܛDRAW(blok.pointsTop, color)
}

func BlokܛDRAW(blok Blok) {
	TrianglesSolidܛDRAW(blok.visibleTriangles)
	color := BLACK()
	if ColorsMatchܛCHECK(color, blok.colorSidesRGBA) {
		color = WHITE()
	}
	LineܛDRAW(blok.pointsBottom[0], blok.pointsTop[0], color)
	LineܛDRAW(blok.pointsBottom[1], blok.pointsTop[1], color)
	LineܛDRAW(blok.pointsBottom[3], blok.pointsTop[3], color)

	LineܛDRAW(blok.pointsBottom[0], blok.pointsBottom[1], color)
	LineܛDRAW(blok.pointsBottom[0], blok.pointsBottom[3], color)

	IsoRecOutlineܛDRAW(blok.pointsTop, color)
}

//ISORECS ISORECS ISORECS ISORECS ISORECS ISORECS ISORECS ISORECS ISORECS
func IsoRecSolidܛDRAW(rec RecIso) {
	TriangleSolidܛDRAW(rec.triangles[0])
	TriangleSolidܛDRAW(rec.triangles[1])
	IsoRecOutlineܛDRAW(rec.points, WHITE())
}
func IsoRecSolidNoOutlineܛDRAW(rec RecIso) {
	TriangleSolidܛDRAW(rec.triangles[0])
	TriangleSolidܛDRAW(rec.triangles[1])
}
func IsoRecOutlineܛDRAW(points []sdl.Point, color []uint8) {
	LineܛDRAW(points[0], points[1], color)
	LineܛDRAW(points[1], points[2], color)
	LineܛDRAW(points[2], points[3], color)
	LineܛDRAW(points[3], points[0], color)
}



//TRIANGLES  TRIANGLES TRIANGLES TRIANGLES TRIANGLES TRIANGLES TRIANGLES
func TrianglesSolidܛDRAW(triangles []Tri) {
	var v []sdl.Vertex
	for i := 0; i < len(triangles); i++ {
		v = append(v, triangles[i].vertices...)
	}
	Rndr.RenderGeometry(nil, v, nil)
}
func TriangleSolidܛDRAW(tri Tri) {
	Rndr.RenderGeometry(nil, tri.vertices, nil)
}

//GRID GRID GRID GRID GRID GRID GRID GRID GRID GRID GRID GRID GRID
/*func DrܛGridRecCentersONLY(grid GridIso, color []uint8) {
	for i := 0; i < len(grid.recs); i++ {
		//DrܛRecCenterSolid(grid.recs[i].cntr, 4, 4, color)
	}
}
func DrܛGrid(grid GridIso) {
	DrܛGridRecsLines(grid)
	DrܛGridPoints(grid)
}
func DrܛGridPoints(grid GridIso) {
	Rndr.SetDrawColor(grid.colorPoints[0], grid.colorPoints[1], grid.colorPoints[2], grid.colorPoints[3])
	Rndr.DrawPoints(grid.points)
}
func DrܛGridRecs(grid GridIso) {
	for i := 0; i < len(grid.recs); i++ {
		DrܛTrianglesSolid(grid.recs[i].triangles)
	}
}
func DrܛGridRecsLines(grid GridIso) {
	for i := 0; i < len(grid.recs); i++ {
		DrܛTrianglesSolid(grid.recs[i].triangles)
		DrܛIsoRecLinesPoints(grid.recs[i].points, grid.colorLines)
	}
}
func DrܛGridLines(grid GridIso) {
	for i := 0; i < len(grid.recs); i++ {
		DrܛIsoRecLinesPoints(grid.recs[i].points, grid.colorLines)
	}
}

// POINTS POINTS POINTS POINTS POINTS POINTS POINTS POINTS POINTS
func DrܛPoints(points []sdl.Point, color []uint8) {
	Rndr.SetDrawColor(color[0], color[1], color[2], color[3])
	Rndr.DrawPoints(points)
}
func DrܛPoint(point sdl.Point, color []uint8) {
	Rndr.SetDrawColor(color[0], color[1], color[2], color[3])
	Rndr.DrawPoint(point.X, point.Y)
}

// LINES LINES LINES LINES LINES LINES LINES LINES LINES LINES
func DrܛLineXY(x1, y1, x2, y2 int32, color []uint8) {
	Rndr.SetDrawColor(color[0], color[1], color[2], color[3])
	Rndr.DrawLine(x1, y1, x2, y2)
}
func DrܛLinePoints(p1, p2 sdl.Point, color []uint8) {
	Rndr.SetDrawColor(color[0], color[1], color[2], color[3])
	Rndr.DrawLine(p1.X, p1.Y, p2.X, p2.Y)
}
func DrܛLinesPoints(points []sdl.Point, color []uint8) {
	Rndr.SetDrawColor(color[0], color[1], color[2], color[3])
	Rndr.DrawLines(points)
}

// ISOMETRIC ISOMETRIC ISOMETRIC ISOMETRIC ISOMETRIC ISOMETRIC
func DrܛIsoBlok(blok BlokIso) {
	DrܛTrianglesSolid(blok.visibleTriangles)
	color := BLACK()
	if ChkܢRGBAColorsMatch(color, blok.colorSidesRGBA) {
		color = WHITE()
	}
	DrܛLinePoints(blok.pointsBottom[0], blok.pointsTop[0], color)
	DrܛLinePoints(blok.pointsBottom[1], blok.pointsTop[1], color)
	DrܛLinePoints(blok.pointsBottom[3], blok.pointsTop[3], color)

	DrܛLinePoints(blok.pointsBottom[0], blok.pointsBottom[1], color)
	DrܛLinePoints(blok.pointsBottom[0], blok.pointsBottom[3], color)

	DrܛIsoRecLinesPoints(blok.pointsTop, color)
}
func DrܛIsoBlokShaded(blok BlokIso, darkTop, darkRight bool, shadowAmount uint8) {
	DrܛTrianglesSolid(blok.visibleTriangles)

	if shadowAmount > 255 {
		shadowAmount = 255
	}
	if shadowAmount < 0 {
		shadowAmount = 0
	}

	if darkTop && darkRight {
		newTri := blok.triTop
		for i := 0; i < len(newTri); i++ {
			for j := 0; j < len(newTri[i].vertices); j++ {
				newTri[i].vertices[j].Color = BLACKǁ3()
				newTri[i].vertices[j].Color.A = shadowAmount
			}
		}
		DrܛTrianglesSolid(newTri)
		shadowAmount -= shadowAmount / 4
		newTri = blok.triRightFront
		for i := 0; i < len(newTri); i++ {
			for j := 0; j < len(newTri[i].vertices); j++ {
				newTri[i].vertices[j].Color = BLACKǁ3()
				newTri[i].vertices[j].Color.A = shadowAmount
			}
		}
		DrܛTrianglesSolid(newTri)
	} else if darkTop && !darkRight {
		newTri := blok.triTop
		for i := 0; i < len(newTri); i++ {
			for j := 0; j < len(newTri[i].vertices); j++ {
				newTri[i].vertices[j].Color = BLACKǁ3()
				newTri[i].vertices[j].Color.A = shadowAmount
			}
		}
		DrܛTrianglesSolid(newTri)
		shadowAmount -= shadowAmount / 4
		newTri = blok.triLeftFront
		for i := 0; i < len(newTri); i++ {
			for j := 0; j < len(newTri[i].vertices); j++ {
				newTri[i].vertices[j].Color = BLACKǁ3()
				newTri[i].vertices[j].Color.A = shadowAmount
			}
		}
		DrܛTrianglesSolid(newTri)
	} else if !darkTop && !darkRight {
		newTri := blok.triLeftFront
		for i := 0; i < len(newTri); i++ {
			for j := 0; j < len(newTri[i].vertices); j++ {
				newTri[i].vertices[j].Color = BLACKǁ3()
				newTri[i].vertices[j].Color.A = shadowAmount
			}
		}
		DrܛTrianglesSolid(newTri)
		shadowAmount -= shadowAmount / 4
		newTri = blok.triRightFront
		for i := 0; i < len(newTri); i++ {
			for j := 0; j < len(newTri[i].vertices); j++ {
				newTri[i].vertices[j].Color = BLACKǁ3()
				newTri[i].vertices[j].Color.A = shadowAmount
			}
		}
		DrܛTrianglesSolid(newTri)
	} else if !darkTop && darkRight {
		newTri := blok.triRightFront
		for i := 0; i < len(newTri); i++ {
			for j := 0; j < len(newTri[i].vertices); j++ {
				newTri[i].vertices[j].Color = BLACKǁ3()
				newTri[i].vertices[j].Color.A = shadowAmount
			}
		}
		DrܛTrianglesSolid(newTri)
		shadowAmount -= shadowAmount / 4
		newTri = blok.triLeftFront
		for i := 0; i < len(newTri); i++ {
			for j := 0; j < len(newTri[i].vertices); j++ {
				newTri[i].vertices[j].Color = BLACKǁ3()
				newTri[i].vertices[j].Color.A = shadowAmount
			}
		}
		DrܛTrianglesSolid(newTri)
	}

	color := BLACK()
	if ChkܢRGBAColorsMatch(color, blok.colorSidesRGBA) {
		color = WHITE()
	}
	DrܛLinePoints(blok.pointsBottom[0], blok.pointsTop[0], color)
	DrܛLinePoints(blok.pointsBottom[1], blok.pointsTop[1], color)
	DrܛLinePoints(blok.pointsBottom[3], blok.pointsTop[3], color)

	DrܛLinePoints(blok.pointsBottom[0], blok.pointsBottom[1], color)
	DrܛLinePoints(blok.pointsBottom[0], blok.pointsBottom[3], color)

	DrܛIsoRecLinesPoints(blok.pointsTop, color)
}
func DrܛIsoBlokOutlineColor(blok BlokIso, outlineColor []uint8) {
	DrܛTrianglesSolid(blok.visibleTriangles)
	DrܛLinePoints(blok.pointsBottom[0], blok.pointsTop[0], outlineColor)
	DrܛLinePoints(blok.pointsBottom[1], blok.pointsTop[1], outlineColor)
	DrܛLinePoints(blok.pointsBottom[3], blok.pointsTop[3], outlineColor)

	DrܛLinePoints(blok.pointsBottom[0], blok.pointsBottom[1], outlineColor)
	DrܛLinePoints(blok.pointsBottom[0], blok.pointsBottom[3], outlineColor)

	DrܛIsoRecLinesPoints(blok.pointsTop, outlineColor)
}
func DrܛIsoBlokNoOutline(blok BlokIso) {
	DrܛTrianglesSolid(blok.visibleTriangles)
}
func DrܛIsoRecLinesPoints(points []sdl.Point, color []uint8) {
	DrܛLinePoints(points[0], points[1], color)
	DrܛLinePoints(points[1], points[2], color)
	DrܛLinePoints(points[2], points[3], color)
	DrܛLinePoints(points[3], points[0], color)
}
func DrܛIsoBlockLines(midbotcorner sdl.Point, w, h int32, color []uint8) {
	i := IsoPointsXY(midbotcorner.X, midbotcorner.Y, w)
	DrܛIsoRecLinesPoints(i, color)
	DrܛLineXY(i[0].X, i[0].Y, i[0].X, i[0].Y-h, color)
	DrܛLineXY(i[1].X, i[1].Y, i[1].X, i[1].Y-h, color)
	DrܛLineXY(i[2].X, i[2].Y, i[2].X, i[2].Y-h, color)
	DrܛLineXY(i[3].X, i[3].Y, i[3].X, i[3].Y-h, color)
	midbotcorner.Y -= h
	i = IsoPointsXY(midbotcorner.X, midbotcorner.Y, w)
	DrܛIsoRecLinesPoints(i, color)
}
func DrܛIsoBlockSolidOutline(midbotcorner sdl.Point, w, h int32, color []uint8) {
	i := IsoPointsXY(midbotcorner.X, midbotcorner.Y, w)
	midbotcorner.Y -= h
	i2 := IsoPointsXY(midbotcorner.X, midbotcorner.Y, w)
	triPoints := []sdl.Point{i2[0], i2[1], i2[2]}
	DrܛTriangleSolidPoints(triPoints, ColorSliceToSDLColor(color))
	triPoints = []sdl.Point{i2[2], i2[3], i2[0]}
	DrܛTriangleSolidPoints(triPoints, ColorSliceToSDLColor(color))
	triPoints = []sdl.Point{i[0], i[1], i2[1]}
	DrܛTriangleSolidPoints(triPoints, ColorSliceToSDLColor(color))
	triPoints = []sdl.Point{i[0], i[3], i2[3]}
	DrܛTriangleSolidPoints(triPoints, ColorSliceToSDLColor(color))
	triPoints = []sdl.Point{i[0], i2[0], i2[1]}
	DrܛTriangleSolidPoints(triPoints, ColorSliceToSDLColor(color))
	DrܛTriangleSolidPoints(triPoints, ColorSliceToSDLColor(color))
	triPoints = []sdl.Point{i[0], i2[0], i2[3]}
	DrܛTriangleSolidPoints(triPoints, ColorSliceToSDLColor(color))
	DrܛLinePoints(i[0], i2[0], BLACK())
	DrܛLinePoints(i[1], i2[1], BLACK())
	DrܛLinePoints(i[3], i2[3], BLACK())
	DrܛLinePoints(i[0], i[1], BLACK())
	DrܛLinePoints(i[0], i[3], BLACK())
	DrܛIsoRecLinesPoints(i2, BLACK())
}
func DrܛIsoBlockSolidOutlineColor(midbotcorner sdl.Point, w, h int32, color, colorOutline []uint8) {
	i := IsoPointsXY(midbotcorner.X, midbotcorner.Y, w)
	midbotcorner.Y -= h
	i2 := IsoPointsXY(midbotcorner.X, midbotcorner.Y, w)
	triPoints := []sdl.Point{i2[0], i2[1], i2[2]}
	DrܛTriangleSolidPoints(triPoints, ColorSliceToSDLColor(color))
	triPoints = []sdl.Point{i2[2], i2[3], i2[0]}
	DrܛTriangleSolidPoints(triPoints, ColorSliceToSDLColor(color))
	triPoints = []sdl.Point{i[0], i[1], i2[1]}
	DrܛTriangleSolidPoints(triPoints, ColorSliceToSDLColor(color))
	triPoints = []sdl.Point{i[0], i[3], i2[3]}
	DrܛTriangleSolidPoints(triPoints, ColorSliceToSDLColor(color))
	triPoints = []sdl.Point{i[0], i2[0], i2[1]}
	DrܛTriangleSolidPoints(triPoints, ColorSliceToSDLColor(color))
	DrܛTriangleSolidPoints(triPoints, ColorSliceToSDLColor(color))
	triPoints = []sdl.Point{i[0], i2[0], i2[3]}
	DrܛTriangleSolidPoints(triPoints, ColorSliceToSDLColor(color))
	DrܛLinePoints(i[0], i2[0], colorOutline)
	DrܛLinePoints(i[1], i2[1], colorOutline)
	DrܛLinePoints(i[3], i2[3], colorOutline)
	DrܛLinePoints(i[0], i[1], colorOutline)
	DrܛLinePoints(i[0], i[3], colorOutline)
	DrܛIsoRecLinesPoints(i2, colorOutline)
}

// SHAPES SHAPES SHAPES SHAPES SHAPES SHAPES SHAPES SHAPES SHAPES SHAPES
func DrܛRecCenterSolid(cntr sdl.Point, w, h int32, color []uint8) {
	r := sdl.Rect{cntr.X - w/2, cntr.Y - h/2, w, h}
	Rndr.SetDrawColor(color[0], color[1], color[2], color[3])
	Rndr.FillRect(&r)
}
func DrܛRecSolid(r sdl.Rect, color []uint8) {
	Rndr.SetDrawColor(color[0], color[1], color[2], color[3])
	Rndr.FillRect(&r)
}
func DrܛIsoRecSolidColor(r RecIso, color sdl.Color) {
	r.triangles[0].vertices[0].Color = color
	r.triangles[0].vertices[1].Color = color
	r.triangles[0].vertices[2].Color = color
	r.triangles[1].vertices[0].Color = color
	r.triangles[1].vertices[1].Color = color
	r.triangles[1].vertices[2].Color = color

	DrܛTrianglesSolid(r.triangles)
}
func DrܛTriangleSolid(tri Tri) {
	Rndr.RenderGeometry(nil, tri.vertices, nil)
}
func DrܛTrianglesSolid(triangles []Tri) {
	var v []sdl.Vertex
	for i := 0; i < len(triangles); i++ {
		v = append(v, triangles[i].vertices...)
	}
	Rndr.RenderGeometry(nil, v, nil)
}
func DrܛTriangleSolidPoints(points []sdl.Point, color sdl.Color) {
	v1 := sdl.Vertex{Point2FPoint(points[0]), color, sdl.FPoint{0, 0}}
	v2 := sdl.Vertex{Point2FPoint(points[1]), color, sdl.FPoint{0, 0}}
	v3 := sdl.Vertex{Point2FPoint(points[2]), color, sdl.FPoint{0, 0}}
	vertices := []sdl.Vertex{v1, v2, v3}
	Rndr.RenderGeometry(nil, vertices, nil)
}
func DrܛTriangleSolidXY(x1, y1, x2, y2, x3, y3 int32, color sdl.Color) {
	v1 := sdl.Vertex{sdl.FPoint{float32(x1), float32(y1)}, color, sdl.FPoint{0, 0}}
	v2 := sdl.Vertex{sdl.FPoint{float32(x2), float32(y2)}, color, sdl.FPoint{0, 0}}
	v3 := sdl.Vertex{sdl.FPoint{float32(x3), float32(y3)}, color, sdl.FPoint{0, 0}}
	vertices := []sdl.Vertex{v1, v2, v3}
	Rndr.RenderGeometry(nil, vertices, nil)
}

*/

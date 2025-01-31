package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

var (
	lev  GRID
	levW = 24

	grassTILES, rocksTERRAIN TILESHEET

	terrain []OBJ

	cards []CARD

	pl1, pl2, pl3 PLAYER
)

type PLAYER struct {
	bloknum int
	nm      string
	onoff   bool
	zIndex  int
}

type OBJ struct {
	im   IM
	rec  sdl.Rect
	cntr sdl.Point
}
type CARD struct {
	im, iconMAIN IM
	rec, recIcon sdl.Rect
	nm           string
	selected     bool
}

func makePlayers() {
	pl1.bloknum = 325
	pl1.nm = "Player 1"
	pl1.onoff = true

}
func makeCards() {

	num := 10

	for i := 0; i < num; i++ {
		cards = append(cards, cards[0])
	}

	for i := 0; i < len(cards); i++ {
		switch Roll6() {
		case 1:
			cards[i].nm = "Block"
			cards[i].iconMAIN = etcIM[2]
		case 2:
			cards[i].nm = "Guard"
			cards[i].iconMAIN = etcIM[2]
		case 3:
			cards[i].nm = "Throw"
			cards[i].iconMAIN = etcIM[2]
		case 4:
			cards[i].nm = "Swing"
			cards[i].iconMAIN = etcIM[2]
		case 5:
			cards[i].nm = "Evade"
			cards[i].iconMAIN = etcIM[2]
		case 6:
			cards[i].nm = "Scavenge"
			cards[i].iconMAIN = etcIM[2]
		}
	}

}

func levelܛMAKE() {

	lev = GridTopCenterOffsetByܛMAKE(levW, levW, UNIT2, UNIT2, UNIT2, ORANGE(), true)
	GridAutoTileRandom(lev, grassTILES)
	lev = GridEdgeDeform(lev)
	lev = addTerrain(lev)

}

func addTerrain(grid GRID) GRID {

	multi := int32(2)
	ob := OBJ{}

	for i := 0; i < len(grid.bloks); i++ {
		if !grid.bloks[i].onoff {
			if Roll12() == 12 {
				ob.im = rocksTERRAIN.tileIM[RINT(0, len(rocksTERRAIN.tileIM))]
				ob.rec = sdl.Rect{0, 0, ob.im.rec.W * multi, ob.im.rec.H * multi}
				ob.cntr = grid.bloks[i].cntr
				ob.cntr.Y -= ob.im.rec.H / 2
				ob.rec.X = ob.cntr.X - ob.rec.W/2
				ob.rec.Y = ob.cntr.Y - ob.rec.H/2
				if ObjAddToListܛCHECK(ob, terrain) {
					terrain = append(terrain, ob)
				}
			}
		}
	}

	return grid
}

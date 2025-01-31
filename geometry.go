package main

import (
	"sort"

	"github.com/veandco/go-sdl2/sdl"
)

var (
	gridList []GRID
)

type TRI struct {
	vertices   []sdl.Vertex
	vertPoints []sdl.Point
	color      sdl.Color
	colorRGBA  []uint8
}

type GRID struct {
	bloks       []BLOK
	frontCorner sdl.Point
}
type BLOK struct {
	frontCorner, cntr         sdl.Point
	dRec                      sdl.Rect
	recs, shadowRecs          []REC
	color                     []uint8
	colorLines                []uint8
	width, height             int32
	frontPoints, floorCorners []sdl.Point
	zIndex, num               int
	tileIM                    IM
	onoff, visible            bool
}
type REC struct {
	triangles []TRI
	color     []uint8
	side      int //SIDE OF BLOK >> 0 BOTTOM / 1 TOP / 2 LEFT / 3 RIGHT

}

// TRIANGLES
func TrianglePoints(p1, p2, p3 sdl.Point, color []uint8) TRI {
	t := TRI{}
	t.colorRGBA = color
	t.color = ColorSliceToSDLColor(color)
	t.vertPoints = []sdl.Point{p1, p2, p3}
	v1 := sdl.Vertex{Point2FPoint(p1), t.color, sdl.FPoint{0, 0}}
	v2 := sdl.Vertex{Point2FPoint(p2), t.color, sdl.FPoint{0, 0}}
	v3 := sdl.Vertex{Point2FPoint(p3), t.color, sdl.FPoint{0, 0}}
	t.vertices = []sdl.Vertex{v1, v2, v3}
	return t
}
func Triangle(p []sdl.Point, color []uint8) TRI {
	t := TRI{}
	t.colorRGBA = color
	t.color = ColorSliceToSDLColor(color)
	t.vertPoints = []sdl.Point{p[0], p[1], p[2]}
	v1 := sdl.Vertex{Point2FPoint(p[0]), t.color, sdl.FPoint{0, 0}}
	v2 := sdl.Vertex{Point2FPoint(p[1]), t.color, sdl.FPoint{0, 0}}
	v3 := sdl.Vertex{Point2FPoint(p[2]), t.color, sdl.FPoint{0, 0}}
	t.vertices = []sdl.Vertex{v1, v2, v3}
	return t
}

// GRID GRID GRID GRID GRID GRID GRID GRID GRID GRID GRID GRID
func GridEdgeDeform(grid GRID) GRID {

	for i := 0; i < len(lev.bloks); i++ {
		flip := false
		if lev.bloks[i].num < 25 && lev.bloks[i].num > 1 {
			flip = true
		} else if lev.bloks[i].num%27 == 0 && lev.bloks[i].num > 27 && lev.bloks[i].num < 675 {
			flip = true
		} else if lev.bloks[i].num > 703 && lev.bloks[i].num < 727 {
			flip = true
		} else if lev.bloks[i].num > 53 && (lev.bloks[i].num-26)%27 == 0 && lev.bloks[i].num < 701 {
			flip = true
		}
		if flip {
			if Roll12() > 7 {
				lev.bloks[i].onoff = true
			}
		}
	}

	return grid
}
func GridܛMAKE(frontCorner sdl.Point, widthNumBloks, heightNumBloks int, widthBlok, heightBlok int32, color []uint8, recOrSquareImages bool) GRID {
	grid := GRID{}

	x := frontCorner.X
	ox := x
	y := frontCorner.Y
	oy := y
	a := widthNumBloks * heightNumBloks
	c := 0
	c2 := 0
	zin := 0
	oz := zin

	for a > 0 {
		v2 := sdl.Point{x, y}
		b := BlokܛMAKE(v2, widthBlok, heightBlok, color, recOrSquareImages)
		b.zIndex = zin
		b.num = c2
		grid.bloks = append(grid.bloks, b)
		x -= widthBlok
		y -= widthBlok / 2
		zin++
		a--
		c++
		c2++
		if c == widthNumBloks {
			c = 0
			x = ox
			x += widthBlok
			ox = x
			y = oy
			y -= widthBlok / 2
			oy = y
			zin = oz
			zin++
			oz = zin
		}
	}
	grid = GridZindexܛSORT(grid)
	return grid
}

func GridCenteredܛMAKE(widthNumBloks, heightNumBloks int, widthBlok, heightBlok int32, color []uint8, recOrSquareImages bool) GRID {
	height := int32(heightNumBloks) * widthBlok
	frontcorner := CNTR
	frontcorner.Y += height / 2
	grid := GridܛMAKE(frontcorner, widthNumBloks, heightNumBloks, widthBlok, heightBlok, color, recOrSquareImages)
	return grid
}
func GridTopCenterܛMAKE(widthNumBloks, heightNumBloks int, widthBlok, heightBlok int32, color []uint8, recOrSquareImages bool) GRID {
	height := int32(heightNumBloks) * widthBlok
	frontcorner := CNTR
	frontcorner.Y = height
	grid := GridܛMAKE(frontcorner, widthNumBloks, heightNumBloks, widthBlok, heightBlok, color, recOrSquareImages)
	return grid
}
func GridTopCenterOffsetByܛMAKE(widthNumBloks, heightNumBloks int, widthBlok, heightBlok, offsetTop int32, color []uint8, recOrSquareImages bool) GRID {
	height := int32(heightNumBloks) * widthBlok
	frontcorner := CNTR
	frontcorner.Y = height + offsetTop
	grid := GridܛMAKE(frontcorner, widthNumBloks, heightNumBloks, widthBlok, heightBlok, color, recOrSquareImages)
	return grid
}
func GridAutoTile(grid GRID, tileIM IM) {
	for i := 0; i < len(grid.bloks); i++ {
		grid.bloks[i].tileIM = tileIM
	}
}
func GridAutoTileRandom(grid GRID, tileSheet TILESHEET) {
	for i := 0; i < len(grid.bloks); i++ {
		grid.bloks[i].tileIM = tileSheet.tileIM[RINT(0, len(tileSheet.tileIM))]
	}
}

func GridTileSingleTile(grid GRID, tileSheet TILESHEET, tileSheetNum int) {
	for i := 0; i < len(grid.bloks); i++ {
		grid.bloks[i].tileIM = tileSheet.tileIM[tileSheetNum]
	}
}

// SORT
func GridZindexܛSORT(grid GRID) GRID {
	sort.Slice(grid.bloks, func(i, j int) bool { return grid.bloks[i].zIndex > grid.bloks[j].zIndex })
	return grid
}

// BLOKS BLOKS BLOKS BLOKS BLOKS BLOKS BLOKS BLOKS BLOKS BLOKS BLOKS
func BlokܛMAKE(frontCorner sdl.Point, w, h int32, color []uint8, recOrSquareImages bool) BLOK {
	b := BLOK{}
	if recOrSquareImages {
		b.dRec = sdl.Rect{frontCorner.X - w, frontCorner.Y - w, w * 2, w * 2}
	} else {
		b.dRec = sdl.Rect{frontCorner.X - w, frontCorner.Y - w, w * 2, w * 2}
	}
	b.frontCorner = frontCorner
	b.cntr = frontCorner
	b.cntr.Y -= w / 2
	b.width = w
	b.height = h
	b.color = color
	if ColorsMatchܛCHECK(color, BLACK()) {
		b.colorLines = WHITE()
	} else {
		b.colorLines = BLACK()
	}
	r := REC{}
	r.color = color
	r.side = 0
	points := RecPointsܛMAKE(frontCorner, w, h, 0)
	b.frontPoints = append(b.frontPoints, points[0], points[1], points[3])
	b.floorCorners = append(b.frontPoints, points[0], points[1], points[3])
	v2 := points[0]
	v2.Y -= w
	b.floorCorners = append(b.floorCorners, v2)
	r.triangles = append(r.triangles, TrianglePoints(points[0], points[1], points[2], color))
	r.triangles = append(r.triangles, TrianglePoints(points[2], points[3], points[0], color))
	b.recs = append(b.recs, r)
	b.shadowRecs = append(b.shadowRecs, r)
	r = REC{}
	r.color = color
	r.side = 1
	points = RecPointsܛMAKE(frontCorner, w, h, 1)
	b.frontPoints = append(b.frontPoints, points...)
	r.triangles = append(r.triangles, TrianglePoints(points[0], points[1], points[2], color))
	r.triangles = append(r.triangles, TrianglePoints(points[2], points[3], points[0], color))
	b.recs = append(b.recs, r)
	b.shadowRecs = append(b.shadowRecs, r)
	r = REC{}
	r.color = color
	r.side = 2
	points = RecPointsܛMAKE(frontCorner, w, h, 2)
	r.triangles = append(r.triangles, TrianglePoints(points[0], points[1], points[2], color))
	r.triangles = append(r.triangles, TrianglePoints(points[2], points[3], points[0], color))
	b.recs = append(b.recs, r)
	b.shadowRecs = append(b.shadowRecs, r)
	r = REC{}
	r.color = color
	r.side = 3
	points = RecPointsܛMAKE(frontCorner, w, h, 3)
	r.triangles = append(r.triangles, TrianglePoints(points[0], points[1], points[2], color))
	r.triangles = append(r.triangles, TrianglePoints(points[2], points[3], points[0], color))
	b.recs = append(b.recs, r)
	b.shadowRecs = append(b.shadowRecs, r)
	return b
}
func BlokRandomColorܛMAKE(frontCorner sdl.Point, w, h int32, recOrSquareImages bool) BLOK {
	color := COLORǁRANDOM()
	b := BlokܛMAKE(frontCorner, w, h, color, recOrSquareImages)
	return b
}

// RECS RECS RECS RECS RECS RECS RECS RECS RECS RECS RECS RECS RECS RECS RECS
func RecColorܛCHANGE(r REC, color sdl.Color) REC {
	for i := 0; i < len(r.triangles); i++ {
		for j := 0; j < len(r.triangles[i].vertices); j++ {
			r.triangles[i].vertices[j].Color = color
		}
	}
	return r
}

func RecܛCOPY(r REC) REC {
	q := REC{}
	q = r
	return q
}

// POINTS POINTS POINTS POINTS POINTS POINTS POINTS POINTS POINTS
func RecPointsܛMAKE(point sdl.Point, w, h int32, side int) []sdl.Point {
	var p []sdl.Point

	if side > 3 || side < 0 {
		debugInfoܛADD("Invalid Blok Side >> 0 Bottom / 1 Top / 2 Left / 3 Right ")
		p = append(p, CNTR, CNTR, CNTR, CNTR)
	} else {
		switch side {
		case 0:
			p = append(p, point)
			point.X -= w
			point.Y -= w / 2
			p = append(p, point)
			point.X += w
			point.Y -= w / 2
			p = append(p, point)
			point.X += w
			point.Y += w / 2
			p = append(p, point)
		case 1:
			point.Y -= w
			p = append(p, point)
			point.X -= w
			point.Y -= w / 2
			p = append(p, point)
			point.X += w
			point.Y -= w / 2
			p = append(p, point)
			point.X += w
			point.Y += w / 2
			p = append(p, point)
		case 2:
			p = append(p, point)
			point.X -= w
			point.Y -= w / 2
			p = append(p, point)
			point.Y -= h
			p = append(p, point)
			point.X += w
			point.Y += w / 2
			p = append(p, point)
		case 3:
			p = append(p, point)
			point.X += w
			point.Y -= w / 2
			p = append(p, point)
			point.Y -= h
			p = append(p, point)
			point.X -= w
			point.Y += w / 2
			p = append(p, point)

		}

	}
	return p
}

/*

type Blok struct {
	num, zIndex                    int
	recs                           []RecIso
	colorSidesRGBA, colorLinesRGBA []uint8
	colorSides, colorLines         sdl.Color
	cntr, nearestFloorCorner       sdl.Point
	solid                          bool

	pointsTop, pointsBottom []sdl.Point

	widthTotal, heightTotal, widthSide, heightSide int32

	frontVisibleVertices, backHiddenVertices []sdl.Vertex

	visibleTriangles, triLeftFront, triLeftBack, triRightFront, triRightBack, triTop, triBottom []Tri

	layer, number int

	name string
}
type Grid struct {
	bloks                                          []Blok
	widthTotal, lengthTotal, blokWidth, blokHeight int32
	cntr, nearestFloorCorner                       sdl.Point
	numBloksTotal                                  int
	colorSidesRGBA, colorLinesRGBA                 []uint8
	colorSides, colorLines                         sdl.Color
}

type RecIso struct {
	points            []sdl.Point
	triangles         []Tri
	color, colorLines []uint8
}

type BlokIso struct {
	pointsTop, pointsBottom        []sdl.Point
	colorSidesRGBA, colorLinesRGBA []uint8
	colorSides, colorLines         sdl.Color
	nearestFloorCorner             sdl.Point

	widthTotal, heightTotal, widthSide, heightSide int32

	frontVisibleVertices, backHiddenVertices []sdl.Vertex

	visibleTriangles, triLeftFront, triLeftBack, triRightFront, triRightBack, triTop, triBottom []Tri

	layer, number int

	name string
}

/*
	type GridIso struct {
		recs []RecIso

		points, linePoints []sdl.Point

		widthTotal, heightTotal, widthSide int32

		layer, num, widthNumBloks, heightNumBloks, bloksNumArea int

		name string

		color, colorPoints, colorLines []uint8

		midBotCorner sdl.Point

		visibleGridUpdated bool
	}
*/

/*
// SORT
func BloksZindexܛSORT(z Grid) Grid {
	sort.Slice(z.bloks, func(i, j int) bool { return z.bloks[i].zIndex > z.bloks[j].zIndex })
	return z
}

// GRID
func GridCenteredRandomBloksRandomColorܛMAKE(widthNumBloks, lengthNumBloks int, widthBlok, heightBlok int32, color []uint8) Grid {
	lengthTotal := widthBlok * int32(widthNumBloks)
	nearestCorner := CNTR
	nearestCorner.Y += lengthTotal / 2
	zz := GridܛMAKE(nearestCorner, widthNumBloks, lengthNumBloks, widthBlok, heightBlok, color)
	for i := 0; i < len(zz.bloks); i++ {
		if FlipCoin() {
			zz.bloks[i].solid = true
		}
	}
	zz = BloksZindexܛSORT(zz)
	for i := 0; i < len(zz.bloks); i++ {
		zz.bloks[i] = BlokRandomColor(zz.bloks[i])
	}
	return zz

}
func GridCenteredRandomBloksܛMAKE(widthNumBloks, lengthNumBloks int, widthBlok, heightBlok int32, color []uint8) Grid {
	lengthTotal := widthBlok * int32(widthNumBloks)
	nearestCorner := CNTR
	nearestCorner.Y += lengthTotal / 2
	zz := GridܛMAKE(nearestCorner, widthNumBloks, lengthNumBloks, widthBlok, heightBlok, color)
	for i := 0; i < len(zz.bloks); i++ {
		if FlipCoin() {
			zz.bloks[i].solid = true
		}
	}
	zz = BloksZindexܛSORT(zz)
	return zz

}
func GridCenteredܛMAKE(widthNumBloks, lengthNumBloks int, widthBlok, heightBlok int32, color []uint8) Grid {
	lengthTotal := widthBlok * int32(widthNumBloks)
	nearestCorner := CNTR
	nearestCorner.Y += lengthTotal / 2
	zz := GridܛMAKE(nearestCorner, widthNumBloks, lengthNumBloks, widthBlok, heightBlok, color)
	zz = BloksZindexܛSORT(zz)
	return zz
}
func GridܛMAKE(nearestCorner sdl.Point, widthNumBloks, lengthNumBloks int, widthBlok, heightBlok int32, color []uint8) Grid {
	zz := Grid{}

	zz.blokWidth = widthBlok
	zz.blokHeight = heightBlok

	zz.colorSidesRGBA = color
	zz.colorSides = ColorSliceToSDLColor(color)
	if ColorsMatchܛCHECK(BLACK(), color) {
		zz.colorLinesRGBA = WHITE()
		zz.colorLines = WHITEǁ3()
	} else {
		zz.colorLinesRGBA = BLACK()
		zz.colorLines = BLACKǁ3()
	}

	a := widthNumBloks * lengthNumBloks

	zz.numBloksTotal = a
	zz.widthTotal = (widthBlok * 2) * int32(widthNumBloks)
	zz.lengthTotal = widthBlok * int32(widthNumBloks)

	zz.nearestFloorCorner = nearestCorner

	c := 0
	c2 := 0
	zin := 0
	prevZin := zin
	v2 := nearestCorner
	prevV2 := v2

	for a > 0 {

		b := BlokܛMAKE(nearestCorner, widthBlok, heightBlok, color)
		b.num = c2

		b.colorSidesRGBA = zz.colorSidesRGBA
		b.colorSides = zz.colorSides
		b.colorLinesRGBA = zz.colorLinesRGBA
		b.colorLines = zz.colorLines

		b.zIndex = zin

		ip := IsoPointsܛMAKE(v2, widthBlok)
		r := RecIso{}
		r.points = ip
		r.color = color
		r.colorLines = b.colorLinesRGBA
		t := TrianglePoints(ip[0], ip[1], ip[2], color)
		r.triangles = append(r.triangles, t)
		t = TrianglePoints(ip[2], ip[3], ip[0], color)
		r.triangles = append(r.triangles, t)

		ip = BlokSidePointsܛMAKE(v2, widthBlok, heightBlok, false)
		t = TrianglePoints(ip[0], ip[1], ip[2], color)
		r.triangles = append(r.triangles, t)
		t = TrianglePoints(ip[2], ip[3], ip[0], color)
		r.triangles = append(r.triangles, t)
		ip = BlokSidePointsܛMAKE(v2, widthBlok, heightBlok, true)
		t = TrianglePoints(ip[0], ip[1], ip[2], color)
		r.triangles = append(r.triangles, t)
		t = TrianglePoints(ip[2], ip[3], ip[0], color)
		r.triangles = append(r.triangles, t)

		b.recs = append(b.recs, r)

		zz.bloks = append(zz.bloks, b)

		a--
		c++
		c2++
		zin++
		v2.X -= widthBlok
		v2.Y -= widthBlok / 2
		if c == widthNumBloks {
			c = 0
			v2 = prevV2
			v2.X += widthBlok
			v2.Y -= widthBlok / 2
			prevV2 = v2

			zin = prevZin
			zin++
			prevZin = zin
		}
	}

	return zz
}

// POINTS
func BlokSidePointsܛMAKE(nearestCorner sdl.Point, w, h int32, leftRight bool) []sdl.Point {
	var points []sdl.Point
	points = append(points, nearestCorner)
	v2 := nearestCorner
	if leftRight {
		v2.X += w
		v2.Y -= w / 2
		points = append(points, v2)
		v2.Y -= h
		points = append(points, v2)
		v2.X -= w
		v2.Y += w / 2
		points = append(points, v2)
	} else {
		v2.X -= w
		v2.Y -= w / 2
		points = append(points, v2)
		v2.Y -= h
		points = append(points, v2)
		v2.X += w
		v2.Y += w / 2
		points = append(points, v2)
	}
	return points
}
func IsoPointsXYܛMAKE(xNearestCorner, yNearestCorner, w int32) []sdl.Point {
	p1 := sdl.Point{xNearestCorner, yNearestCorner}
	p2 := p1
	p2.X += w
	p2.Y -= w / 2
	p3 := p2
	p3.X -= w * 2
	p4 := p1
	p4.Y -= w
	points := []sdl.Point{p1, p3, p4, p2}
	return points
}

func IsoPointsܛMAKE(nearestCorner sdl.Point, w int32) []sdl.Point {
	p2 := nearestCorner
	p2.X += w
	p2.Y -= w / 2
	p3 := p2
	p3.X -= w * 2
	p4 := nearestCorner
	p4.Y -= w
	points := []sdl.Point{nearestCorner, p3, p4, p2}
	return points
}

// BLOKS
func BlokChangeColor(b Blok) Blok {

	return b
}
func BlokRandomColor(b Blok) Blok {
	newcolor := COLORǁRANDOMǁ3()
	for i := 0; i < len(b.recs); i++ {
		for j := 0; j < len(b.recs[i].triangles); j++ {
			for k := 0; k < len(b.recs[i].triangles[j].vertices); k++ {
				b.recs[i].triangles[j].vertices[k].Color = newcolor
			}
		}
	}

	for i := 0; i < len(b.visibleTriangles); i++ {
		for j := 0; j < len(b.visibleTriangles[i].vertices); j++ {
			b.visibleTriangles[i].color = newcolor

		}
	}

	return b
}
func BlokܛMAKE(nearestCorner sdl.Point, w, h int32, color []uint8) Blok {
	b := Blok{}
	b.cntr = nearestCorner
	b.cntr.Y -= w / 2
	b.nearestFloorCorner = nearestCorner
	b.widthTotal = w * 2
	b.heightTotal = h * 2
	b.widthSide = w
	b.heightSide = h
	b.colorSidesRGBA = color
	b.colorSides = ColorSliceToSDLColor(color)
	b.pointsBottom = IsoPointsXYܛMAKE(nearestCorner.X, nearestCorner.Y, w)
	nearestCorner.Y -= h
	b.pointsTop = IsoPointsXYܛMAKE(nearestCorner.X, nearestCorner.Y, w)

	t := TrianglePoints(b.pointsBottom[0], b.pointsBottom[1], b.pointsBottom[2], color)
	b.triBottom = append(b.triBottom, t)
	t = TrianglePoints(b.pointsBottom[2], b.pointsBottom[3], b.pointsBottom[0], color)
	b.triBottom = append(b.triBottom, t)

	t = TrianglePoints(b.pointsTop[0], b.pointsTop[1], b.pointsTop[2], color)
	b.triTop = append(b.triTop, t)
	t = TrianglePoints(b.pointsTop[2], b.pointsTop[3], b.pointsTop[0], color)
	b.triTop = append(b.triTop, t)

	t = TrianglePoints(b.pointsBottom[0], b.pointsTop[1], b.pointsBottom[1], color)
	b.triLeftFront = append(b.triLeftFront, t)
	t = TrianglePoints(b.pointsBottom[0], b.pointsTop[0], b.pointsTop[1], color)
	b.triLeftFront = append(b.triLeftFront, t)

	t = TrianglePoints(b.pointsBottom[0], b.pointsTop[3], b.pointsBottom[3], color)
	b.triRightFront = append(b.triRightFront, t)
	t = TrianglePoints(b.pointsBottom[0], b.pointsTop[0], b.pointsTop[3], color)
	b.triRightFront = append(b.triRightFront, t)

	t = TrianglePoints(b.pointsBottom[1], b.pointsTop[2], b.pointsBottom[2], color)
	b.triLeftBack = append(b.triLeftBack, t)
	t = TrianglePoints(b.pointsBottom[1], b.pointsTop[1], b.pointsTop[2], color)
	b.triLeftBack = append(b.triLeftBack, t)

	t = TrianglePoints(b.pointsBottom[3], b.pointsTop[2], b.pointsBottom[2], color)
	b.triRightBack = append(b.triRightBack, t)
	t = TrianglePoints(b.pointsBottom[3], b.pointsTop[3], b.pointsTop[2], color)
	b.triRightBack = append(b.triRightBack, t)

	b.visibleTriangles = append(b.visibleTriangles, b.triTop...)
	b.visibleTriangles = append(b.visibleTriangles, b.triLeftFront...)
	b.visibleTriangles = append(b.visibleTriangles, b.triRightFront...)

	return b
}

/*
	func IsoGridBottomCenterScreen(widthNumBloks, heightNumBloks int, widthBlokSide int32, color []uint8) GridIso {
		v2 := CNTR
		v2.Y = WinH
		g := IsoGrid(v2, widthNumBloks, heightNumBloks, widthBlokSide, color)
		return g
	}

	func IsoGridTopCenterScreen(widthNumBloks, heightNumBloks int, widthBlokSide int32, color []uint8) GridIso {
		heightTotal := widthBlokSide * int32(widthNumBloks)
		v2 := CNTR
		v2.Y = heightTotal
		g := IsoGrid(v2, widthNumBloks, heightNumBloks, widthBlokSide, color)
		return g
	}

	func IsoGridLeftCenterScreen(widthNumBloks, heightNumBloks int, widthBlokSide int32, color []uint8) GridIso {
		widthTotal := (widthBlokSide * 2) * int32(widthNumBloks)
		heightTotal := widthBlokSide * int32(widthNumBloks)
		v2 := CNTR
		v2.Y += heightTotal / 2
		v2.X = widthTotal / 2
		g := IsoGrid(v2, widthNumBloks, heightNumBloks, widthBlokSide, color)
		return g
	}

	func IsoGridRightCenterScreen(widthNumBloks, heightNumBloks int, widthBlokSide int32, color []uint8) GridIso {
		widthTotal := (widthBlokSide * 2) * int32(widthNumBloks)
		heightTotal := widthBlokSide * int32(widthNumBloks)
		v2 := CNTR
		v2.Y += heightTotal / 2
		v2.X = WinW - widthTotal/2
		g := IsoGrid(v2, widthNumBloks, heightNumBloks, widthBlokSide, color)
		return g
	}

	func IsoGridCenterScreen(widthNumBloks, heightNumBloks int, widthBlokSide int32, color []uint8) GridIso {
		heightTotal := widthBlokSide * int32(widthNumBloks)
		v2 := CNTR
		v2.Y += heightTotal / 2
		g := IsoGrid(v2, widthNumBloks, heightNumBloks, widthBlokSide, color)
		return g
	}

	func IsoGridColors(midBotCorner sdl.Point, widthNumBloks, heightNumBloks int, widthBlokSide int32, colorRecs, colorLines, colorPoints []uint8) GridIso {
		g := IsoGrid(midBotCorner, widthNumBloks, heightNumBloks, widthBlokSide, colorRecs)
		g.colorLines = colorLines
		g.colorPoints = colorPoints
		return g
	}

	func IsoGrid(midBotCorner sdl.Point, widthNumBloks, heightNumBloks int, widthBlokSide int32, color []uint8) GridIso {
		g := GridIso{}
		g.num = len(GridList)
		g.widthTotal = (widthBlokSide * 2) * int32(widthNumBloks)
		g.heightTotal = widthBlokSide * int32(widthNumBloks)
		g.midBotCorner = midBotCorner
		g.bloksNumArea = widthNumBloks * heightNumBloks
		g.widthSide = widthBlokSide
		g.color = color
		g.colorLines = BLACK()
		g.colorPoints = RED()
		if ChkܢRGBAColorsMatch(color, BLACK()) {
			g.colorLines = WHITE()
		}

		v2 := midBotCorner
		prevX := v2.X
		prevY := v2.Y
		a := g.bloksNumArea
		c := 0
		c2 := 0
		z := 0
		prevZ := 0
		for a > 0 {
			p := IsoPoints(v2, widthBlokSide)
			g.points = append(g.points, p...)
			r := RecIso{}
			r.points = p
			r.color = color
			//r.number = c2
			//r.zIndex = z
			t := TrianglePoints(p[0], p[1], p[2], color)
			r.triangles = append(r.triangles, t)
			t = TrianglePoints(p[2], p[3], p[0], color)
			r.triangles = append(r.triangles, t)
			//r.cntr.X = p[0].X
			//r.cntr.Y = p[1].Y
			g.recs = append(g.recs, r)

			v2.X -= widthBlokSide
			v2.Y -= widthBlokSide / 2

			c++
			c2++
			a--
			if c == widthNumBloks {
				c = 0
				z = prevZ
				z++
				prevZ = z
				v2.X = prevX
				v2.X += widthBlokSide
				prevX = v2.X
				v2.Y = prevY
				v2.Y -= widthBlokSide / 2
				prevY = v2.Y
			}
		}

		GridList = append(GridList, g)
		return g

}

*/

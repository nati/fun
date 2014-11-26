package go3D

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type Canvas interface {
	DrawPoint(*Point2D)
	DrawLine(*Point2D)
	Render3D(*Point3D, *PolygonGroup)
	Clear()
	Flush()
}

type TerminalCanvas struct {
	Width, Height int
	Char          string
	writer        *bufio.Writer
}

func (t *TerminalCanvas) DrawPoint(p *Point2D) {
	x := int(p.x) + t.Width/2
	y := int(p.y) + t.Height/2
	if x < 0 || t.Width < x || y < 0 || t.Height < y {
		return
	}
	t.writer.WriteString(fmt.Sprintf("\033[%d;%dH%s", y, x, t.Char))
}

func (t *TerminalCanvas) DrawLine(p1 Point2D, p2 Point2D) {
	steep := math.Abs(p2.y-p1.y) > math.Abs(p2.x-p1.x)
	if steep {
		p1.x, p1.y = p1.y, p1.x
		p2.x, p2.y = p2.y, p2.x
	}
	if p1.x > p2.x {
		p1.x, p2.x = p2.x, p1.x
		p1.y, p2.y = p2.y, p1.y
	}
	dx := p2.x - p1.x
	dy := math.Abs(p2.y - p1.y)
	err := dx / 2.0
	y := p1.y
	var ystep float64
	if p1.y < p2.y {
		ystep = 1.0
	} else {
		ystep = -1.0
	}
	for x := p1.x; x < p2.x; x++ {
		if steep {
			t.DrawPoint(NewPoint2D(y, x))
		} else {
			t.DrawPoint(NewPoint2D(x, y))
		}
		err = err - dy
		if err < 0 {
			y = y + ystep
			err = err + dx
		}
	}
}

func (t *TerminalCanvas) Clear() {
	t.writer.WriteString("\033[2J")
}

func (t *TerminalCanvas) Flush() {
	t.writer.Flush()
}

func NewTerminalCanvas(width, height int, char string) *TerminalCanvas {
	canvas := &TerminalCanvas{
		Width:  width,
		Height: height,
		Char:   char,
	}
	canvas.writer = bufio.NewWriter(os.Stdout)
	return canvas
}

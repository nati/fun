package go3D

import (
	"fmt"
	"math"
)

type Point3D struct {
	x float64
	y float64
	z float64
}

type Point2D struct {
	x float64
	y float64
}

type PolygonGroup struct {
	name          string
	polygons      []*Polygon
	vertexs       []*Point3D
	textures      []*Point2D
	normalVectors []*Point3D
	Point         *Point3D
	Rotate        *Point3D
	Center        *Point3D
}

func (p *PolygonGroup) String() string {
	response := fmt.Sprintf("PolygonGroup: %s \n", p.name)
	response += "\tVertex ["
	for _, vertex := range p.vertexs {
		response += vertex.String() + ", "
	}
	response += "]\n"
	response += "\tNormalVectors ["
	for _, nv := range p.normalVectors {
		response += nv.String() + ", "
	}
	response += "]\n"
	response += "\tTexures ["
	for _, t := range p.textures {
		response += t.String() + ", "
	}
	response += "]\n"
	response += "\tPolygons ["
	for _, p := range p.polygons {
		response += p.String() + ", "
	}
	response += "]\n"
	return response
}

type Polygon struct {
	elements []*Element
}

func (p *Polygon) String() string {
	response := ""
	for _, element := range p.elements {
		response += element.String() + " "
	}
	return response
}

func NewPolygon() *Polygon {
	return &Polygon{}
}

type Element struct {
	vertex       int
	texture      int
	normalVector int
}

func NewPolygonGroup(name string) *PolygonGroup {
	return &PolygonGroup{
		name:   name,
		Point:  NewPoint3D(0.0, 0.0, 0.0),
		Center: NewPoint3D(0.0, 0.0, 0.0),
		Rotate: NewPoint3D(0.0, 0.0, 0.0),
	}
}

func (p *PolygonGroup) SetPoint(x, y, z float64) {
	p.Point.x = x
	p.Point.y = y
	p.Point.z = z
}

func (p *PolygonGroup) SetRotate(x, y, z float64) {
	p.Rotate.x = x
	p.Rotate.y = y
	p.Rotate.z = z
}

func (p *PolygonGroup) SetCenter(x, y, z float64) {
	p.Center.x = x
	p.Center.y = y
	p.Center.z = z
}

func NewPoint2D(x, y float64) *Point2D {
	return &Point2D{
		x: x,
		y: y,
	}
}

func (p *Point2D) String() string {
	return fmt.Sprintf("(%f, %f)", p.x, p.y)
}

func NewPoint3D(x, y, z float64) *Point3D {
	return &Point3D{
		x: x,
		y: y,
		z: z,
	}
}

func (p *Point3D) String() string {
	return fmt.Sprintf("(%f, %f, %f)", p.x, p.y, p.z)
}

func NewElement(vertex, texture, normalVector int) *Element {
	return &Element{
		vertex:       vertex,
		texture:      texture,
		normalVector: normalVector,
	}
}

func (e *Element) String() string {
	return fmt.Sprintf("(%d, %d, %d)", e.vertex, e.texture, e.normalVector)
}

func xRotateMatrix(a float64) [3][3]float64 {
	return [3][3]float64{
		{1.0, 0.0, 0.0},
		{0.0, math.Cos(a), -math.Sin(a)},
		{0.0, math.Sin(a), math.Cos(a)},
	}
}

func yRotateMatrix(a float64) [3][3]float64 {
	return [3][3]float64{
		{math.Cos(a), 0.0, math.Sin(a)},
		{0.0, 1.0, 0.0},
		{-math.Sin(a), 0.0, math.Cos(a)},
	}
}

func (p *Point3D) rotate(m [3][3]float64) *Point3D {
	return NewPoint3D(
		m[0][0]*p.x+m[0][1]*p.y+m[0][2]*p.z,
		m[1][0]*p.x+m[1][1]*p.y+m[1][2]*p.z,
		m[2][0]*p.x+m[2][1]*p.y+m[2][2]*p.z)
}

func (p1 *Point3D) add(p2 *Point3D) *Point3D {
	return NewPoint3D(
		p1.x+p2.x,
		p1.y+p2.y,
		p1.z+p2.z,
	)
}

func (p1 *Point3D) sub(p2 *Point3D) *Point3D {
	return NewPoint3D(
		p1.x-p2.x,
		p1.y-p2.y,
		p1.z-p2.z,
	)
}

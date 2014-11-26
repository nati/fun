package go3D

func (p *PolygonGroup) Projection(camera, p3d *Point3D) *Point2D {
	p3d = p3d.sub(p.Center)
	p3d = p3d.rotate(xRotateMatrix(p.Rotate.x))
	p3d = p3d.rotate(yRotateMatrix(p.Rotate.y))
	p3d = p3d.add(p.Point)
	distance := p.Point.z - camera.z
	if distance == 0.0 {
		return NewPoint2D(p3d.x, p3d.y)
	} else if distance > 0.0 {
		return NewPoint2D(p3d.x/distance, p3d.y/distance)
	}
	return NewPoint2D(0, 0)
}

func (c *TerminalCanvas) Render3D(camera *Point3D, polygonGroup *PolygonGroup) {
	for _, polygon := range polygonGroup.polygons {
		elements := polygon.elements
		if len(elements) < 3 {
			continue
		}
		for i := 0; i < len(elements); i++ {
			p3d1 := polygonGroup.vertexs[elements[i].vertex-1]
			p3d2 := polygonGroup.vertexs[elements[(i+1)%len(elements)].vertex-1]
			p2d1 := polygonGroup.Projection(camera, p3d1)
			p2d2 := polygonGroup.Projection(camera, p3d2)
			c.DrawLine(*p2d1, *p2d2)
			//c.DrawPoint(p2d1)
		}
	}
}

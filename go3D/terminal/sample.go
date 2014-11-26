package main

import (
	"fmt"
	"github.com/nati/fun/go3D"
	"math"
	"time"
)

func main() {
	//canvas := go3D.NewTerminalCanvas(200, 40, "█")
	canvas := go3D.NewTerminalCanvas(200, 60, "*")
	polygonGroups, err := go3D.ReadObjFile("test.obj")
	if err != nil {
		fmt.Println(err)
		return
	}
	i := 0.0
	for {
		camera := go3D.NewPoint3D(0, 0, -3)
		canvas.Clear()
		for _, polygonGroup := range polygonGroups {
			polygonGroup.SetRotate(i, i, 0)
			polygonGroup.SetCenter(0, 0, 0)
			canvas.Render3D(camera, polygonGroup)
		}
		canvas.Flush()
		time.Sleep(100 * time.Millisecond)
		i += math.Pi / 30.0
	}
}

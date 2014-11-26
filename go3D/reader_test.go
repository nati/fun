package go3D

import (
	"fmt"
	"testing"
)

func TestReadObjFile(t *testing.T) {
	polygons, err := ReadObjFile("./cube.obj")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	fmt.Println(polygons)
}

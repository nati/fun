package go3D

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadObjFile(fileName string) ([]*PolygonGroup, error) {
	var polygonGroups []*PolygonGroup
	var fp *os.File
	fp, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer fp.Close()

	var polygonGroup *PolygonGroup

	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		line := scanner.Text()
		row := strings.Fields(line)
		if len(row) < 2 {
			continue
		}
		typeCode := row[0]
		switch typeCode {
		case "g":
			polygonGroup = NewPolygonGroup(row[1])
			polygonGroups = append(polygonGroups, polygonGroup)
		case "v":
			x, err := strconv.ParseFloat(row[1], 64)
			if err != nil {
				return nil, err
			}
			y, err := strconv.ParseFloat(row[2], 64)
			if err != nil {
				return nil, err
			}
			z, err := strconv.ParseFloat(row[3], 64)
			if err != nil {
				return nil, err
			}
			polygonGroup.vertexs = append(polygonGroup.vertexs, NewPoint3D(x, y, z))
		case "vt":
			x, err := strconv.ParseFloat(row[1], 64)
			if err != nil {
				return nil, err
			}
			y, err := strconv.ParseFloat(row[2], 64)
			if err != nil {
				return nil, err
			}
			polygonGroup.textures = append(polygonGroup.textures, NewPoint2D(x, y))
		case "vn":
			x, err := strconv.ParseFloat(row[1], 64)
			if err != nil {
				return nil, err
			}
			y, err := strconv.ParseFloat(row[2], 64)
			if err != nil {
				return nil, err
			}
			z, err := strconv.ParseFloat(row[3], 64)
			if err != nil {
				return nil, err
			}
			polygonGroup.normalVectors = append(polygonGroup.normalVectors, NewPoint3D(x, y, z))
		case "f":
			polygon := NewPolygon()
			for _, value := range row[1:] {
				elementData := strings.Split(value, "/")
				vertex, err := strconv.Atoi(elementData[0])
				if err != nil {
					vertex = 0
				}
				texture, err := strconv.Atoi(elementData[1])
				if err != nil {
					texture = 0
				}
				normalVertex, err := strconv.Atoi(elementData[2])
				if err != nil {
					normalVertex = 0
				}
				polygon.elements = append(polygon.elements, NewElement(vertex, texture, normalVertex))
			}
			polygonGroup.polygons = append(polygonGroup.polygons, polygon)
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	fmt.Printf("test %s", polygonGroups)
	return polygonGroups, nil
}

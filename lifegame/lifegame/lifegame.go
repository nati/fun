package lifegame

import (
	"bytes"
	"image"
	"image/color"
	"io/ioutil"
	"math/rand"
	"strconv"
	"sync"
)

//Map uint64erface represents map
type Map interface {
	Get(x, y uint64) bool
	Set(x, y uint64, flag bool)
	Width() uint64
	Height() uint64
}

//ArrayMap is array based impl for Map
type ArrayMap struct {
	width  uint64
	height uint64
	BitMap [][]bool
}

//NewArrayMap creates ArrayMap
func NewArrayMap(width, height int) Map {
	width += 2
	height += 2
	var bitmap = make([][]bool, height)
	for i := 0; i < height; i++ {
		bitmap[i] = make([]bool, width)
	}
	return &ArrayMap{
		width:  uint64(width),
		height: uint64(height),
		BitMap: bitmap,
	}
}

//Get flag
func (m *ArrayMap) Get(x, y uint64) bool {
	return m.BitMap[y][x]
}

//Set flag
func (m *ArrayMap) Set(x, y uint64, flag bool) {
	m.BitMap[y][x] = flag
}

//Width return width
func (m *ArrayMap) Width() uint64 {
	return m.width - 2
}

//Height return height
func (m *ArrayMap) Height() uint64 {
	return m.height - 2
}

//BitMap is bitmap based impl for Map
type BitMap struct {
	width  uint64
	height uint64
	bitmap []uint64
}

const rank = 64

//NewBitMap creates ArrayMap
func NewBitMap(width, height int) Map {
	width += 2
	height += 2
	max := width * height
	var bitmap = make([]uint64, max/rank+1)
	return &BitMap{
		width:  uint64(width),
		height: uint64(height),
		bitmap: bitmap,
	}
}

//Get flag
func (m *BitMap) Get(x, y uint64) bool {
	index := x + y*m.width
	listIndex := index / rank
	bitIndex := index - listIndex*rank
	bit := m.bitmap[listIndex]
	return bit&(1<<bitIndex) != 0
}

//Set flag
func (m *BitMap) Set(x, y uint64, flag bool) {
	index := x + y*m.width
	listIndex := index / rank
	bitIndex := index - listIndex*rank
	if flag {
		m.bitmap[listIndex] |= (1 << bitIndex)
	} else {
		m.bitmap[listIndex] &= ^(1 << bitIndex)
	}
}

//Width return width
func (m *BitMap) Width() uint64 {
	return m.width - 2
}

//Height return height
func (m *BitMap) Height() uint64 {
	return m.height - 2
}

//NextState compute next state
func NextState(x, y uint64, m Map) bool {
	count := 0
	if m.Get(x-1, y-1) {
		count++
	}
	if m.Get(x-1, y) {
		count++
	}
	if m.Get(x-1, y+1) {
		count++
	}
	if m.Get(x, y-1) {
		count++
	}
	if m.Get(x, y+1) {
		count++
	}
	if m.Get(x+1, y-1) {
		count++
	}
	if m.Get(x+1, y) {
		count++
	}
	if m.Get(x+1, y+1) {
		count++
	}
	if m.Get(x, y) {
		if count == 2 || count == 3 {
			return true
		}
	} else {
		if count == 3 {
			return true
		}
	}
	return false
}

//Next compute next state for next map
func Next(currentState, nextState Map) {
	var x, y uint64
	for x = 1; x <= currentState.Width(); x++ {
		for y = 1; y <= currentState.Height(); y++ {
			nextState.Set(x, y, NextState(x, y, currentState))
		}
	}
}

//NextParallel compute next state for next map
func NextParallel(currentState, nextState Map) {
	var wg sync.WaitGroup
	var x, y uint64
	for x = 1; x <= currentState.Width(); x++ {
		wg.Add(1)
		go func(x uint64) {
			for y = 1; y <= currentState.Height(); y++ {
				nextState.Set(x, y, NextState(x, y, currentState))
			}
			wg.Done()
		}(x)
	}
	wg.Wait()
}

//RandomPoints set num of true
func RandomPoints(m Map, num int) {
	for i := 0; i < num; i++ {
		x := rand.Intn(int(m.Width()))
		y := rand.Intn(int(m.Height()))
		m.Set(uint64(x), uint64(y), true)
	}
}

//DrawMap write map to canvas
func DrawMap(rgba *image.RGBA, m Map) {
	var x, y uint64
	for x = 1; x <= m.Width(); x++ {
		for y = 1; y <= m.Height(); y++ {
			if m.Get(x, y) {
				rgba.Set(int(x), int(y), color.RGBA{0, 128, 0, 255})
			}
		}
	}
}

func readRLEToken(data []byte, bitMap Map) {
	ignore := false
	var x, y uint64
	x = 1
	y = 1
	var buffer bytes.Buffer
	for i := 0; i < len(data); i++ {
		c := data[i]
		switch c {
		case '#', 'x':
			ignore = true
		case '\n':
			ignore = false
		case '!':
			return
		case '$':
			y++
			x = 1
			buffer.Reset()
			continue
		case 'o', 'b':
			if !ignore {
				flag := false
				if c == 'o' {
					flag = true
				}
				num := buffer.String()
				buffer.Reset()
				count := 1
				if num != "" {
					count, _ = strconv.Atoi(num)
				}
				for i := 0; i < count; i++ {
					bitMap.Set(x, y, flag)
					x++
				}
			}
		default:
			if !ignore {
				buffer.WriteByte(c)
			}
			continue
		}
	}
}

//ReadRLE reads RLE file format
func ReadRLE(filename string, bitMap Map) error {
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	readRLEToken(contents, bitMap)
	return nil
}

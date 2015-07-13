package main

import (
	"fmt"
	"github.com/google/gxui"
	"github.com/google/gxui/drivers/gl"
	"github.com/google/gxui/samples/flags"
	"github.com/google/gxui/themes/dark"
	"github.com/nati/fun/lifegame/lifegame"
	"image"
	"math/rand"
	"sync"
	"time"
)

func appMain(driver gxui.Driver) {
	theme := dark.CreateTheme(driver)

	wSize := 500
	hSize := 250
	window := theme.CreateWindow(1000, 500, "LifeGame")
	window.SetScale(flags.DefaultScaleFactor)

	canvasImage := theme.CreateImage()
	window.AddChild(canvasImage)
	ticker := time.NewTicker(time.Millisecond * 30)

	noise := 0
	addNoise := false
	maps := [2]lifegame.Map{
		lifegame.NewBitMap(wSize, hSize),
		lifegame.NewBitMap(wSize, hSize),
	}
	//maps := [2]lifegame.Map{
	//	lifegame.NewArrayMap(size, size),
	//	lifegame.NewArrayMap(size, size),
	//}

	fmt.Println("Start lifegame")
	lifegame.ReadRLE("./broken-lines.rle", maps[0])

	go func() {
		i := 0
		for _ = range ticker.C {
			currentState := maps[i%2]
			nextState := maps[(i+1)%2]
			var wg sync.WaitGroup
			wg.Add(1)
			go func() {
				start := time.Now()
				//lifegame.NextParallel(currentState, nextState)
				lifegame.Next(currentState, nextState)
				elasped := time.Since(start)
				fmt.Printf("elasped %s \n", elasped)
				if addNoise {
					if i < 50 {
						lifegame.RandomPoints(nextState, rand.Intn(noise*100))
					} else {
						lifegame.RandomPoints(nextState, rand.Intn(noise+1))
					}
				}
				wg.Done()
			}()
			wg.Add(1)
			go func() {
				driver.Call(func() {
					rgba := image.NewRGBA(image.Rect(0, 0, wSize, hSize))
					lifegame.DrawMap(rgba, currentState)
					texture := driver.CreateTexture(rgba, 0.5)
					canvasImage.SetTexture(texture)
					wg.Done()
				})
			}()
			wg.Wait()
			i++

		}
	}()

	window.OnClose(driver.Terminate)
}

func main() {
	gl.StartDriver(appMain)
}

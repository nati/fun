package main

import (
	"flag"
	"fmt"
	"github.com/google/gxui"
	"github.com/google/gxui/drivers/gl"
	"github.com/google/gxui/samples/flags"
	"github.com/google/gxui/themes/dark"
	"github.com/nati/fun/lifegame/lifegame"
	"image"
	"os"
	"time"
)

func appMain(driver gxui.Driver) {
	args := flag.Args()
	if len(args) != 1 {
		fmt.Print("usage: lifegame rle_file_path\n")
		os.Exit(1)
	}
	file := args[0]

	theme := dark.CreateTheme(driver)

	wSize := 500
	hSize := 500
	window := theme.CreateWindow(1000, 1000, "LifeGame")
	window.SetScale(flags.DefaultScaleFactor)

	canvasImage := theme.CreateImage()
	window.AddChild(canvasImage)
	ticker := time.NewTicker(time.Millisecond * 30)

	maps := [2]lifegame.Map{
		lifegame.NewBitMap(wSize, hSize),
		lifegame.NewBitMap(wSize, hSize),
	}

	fmt.Println("Start lifegame")
	lifegame.ReadRLE(file, maps[0])

	go func() {
		i := 0
		for _ = range ticker.C {
			currentState := maps[i%2]
			nextState := maps[(i+1)%2]
			start := time.Now()
			lifegame.Next(currentState, nextState)
			elasped := time.Since(start)
			fmt.Printf("elasped %s \n", elasped)
			driver.Call(func() {
				rgba := image.NewRGBA(image.Rect(0, 0, wSize, hSize))
				lifegame.DrawMap(rgba, currentState)
				texture := driver.CreateTexture(rgba, 0.5)
				canvasImage.SetTexture(texture)
			})
			i++
		}
	}()

	window.OnClose(driver.Terminate)
}

func main() {
	gl.StartDriver(appMain)
}

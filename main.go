package main

import (
	"fmt"
	"math/rand"
	"time"
	//        "image/color"
	//        "log"
	//        "math"

	//        "github.com/golang/freetype/truetype"
	//        "golang.org/x/image/font"

	"github.com/hajimehoshi/ebiten"
	//        "github.com/hajimehoshi/ebiten/audio"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	//        "github.com/hajimehoshi/ebiten/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/inpututil"
	//        "github.com/hajimehoshi/ebiten/text"
)

var gopher [10]int

var (
	keys = []ebiten.Key{
		ebiten.Key0,
		ebiten.Key1,
		ebiten.Key2,
		ebiten.Key3,
		ebiten.Key4,
		ebiten.Key5,
		ebiten.Key6,
		ebiten.Key7,
		ebiten.Key8,
		ebiten.Key9,
	}
)

var gopherUp = 0
var gopherX = 0

func update(screen *ebiten.Image) error {
	if gopherUp == 0 {
		seed := rand.NewSource(time.Now().UnixNano())
		rnd := rand.New(seed)
		gopherX = rnd.Intn(10)
		gopherUp = 100
	} else {
		gopherUp--
	}

	ebitenutil.DebugPrint(screen, fmt.Sprintf("FPS: %0.2f Gopher[%d] is here...", ebiten.CurrentFPS(), gopherX))
	for i, key := range keys {
		if inpututil.IsKeyJustPressed(key) {
			ebitenutil.DebugPrint(screen, fmt.Sprintf("FPS: %0.2f\n%d\nkey %d pressed!", ebiten.CurrentFPS(), i, key))
			if i == gopherX {
				ebitenutil.DebugPrint(screen, fmt.Sprintf("\n\n\n\nOUCH!!!"))
			} else {
				ebitenutil.DebugPrint(screen, fmt.Sprintf("\n\n\n\n\nYOU MISSED ME!!!"))
			}
		}
	}

	return nil
}

func main() {
	ebiten.Run(update, 320, 240, 2, "Hello world!")
}

package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"log"
)

type Ship struct {
	image      *ebiten.Image
	gameObject GameObject
}

func NewShip(screenWidth, screenHeight int) *Ship {
	path := "D:\\IDEA Project\\goDemo\\src\\alienGame\\ship.png"
	img, _, err := ebitenutil.NewImageFromFile(path)
	if err != nil {
		log.Fatal(err)
	}

	width, height := img.Size()
	ship := &Ship{
		image: img,
		gameObject: GameObject{
			width:  width,
			height: height,
			x:      float64(screenWidth-width) / 2,
			y:      float64(screenHeight - height),
		},
	}

	return ship
}

func (ship *Ship) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(ship.gameObject.x, ship.gameObject.y)
	screen.DrawImage(ship.image, op)
}

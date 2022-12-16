package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"log"
)

type Alien struct {
	image       *ebiten.Image
	gameObject  GameObject
	speedFactor float64
}

func NewAlien(cfg *Config) *Alien {
	path := "D:\\IDEA Project\\goDemo\\src\\alienGame\\alien.png"
	img, _, err := ebitenutil.NewImageFromFile(path)
	if err != nil {
		log.Fatal(err)
	}

	width, height := img.Size()
	return &Alien{
		image: img,
		gameObject: GameObject{
			width:  width / 2,
			height: height / 2,
			x:      0,
			y:      0,
		},
		speedFactor: cfg.AlienSpeedFactor,
	}
}

func (alien *Alien) outOfScreen(cfg *Config) bool {
	if alien.gameObject.height > cfg.ScreenHeight/2 {
		return true
	}
	return false
}

func (g *Game) addAlien(alien *Alien) {
	g.aliens[alien] = struct{}{}
}

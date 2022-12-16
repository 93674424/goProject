package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image"
)

type Bullet struct {
	image       *ebiten.Image
	gameObject  GameObject
	speedFactor float64
}

func (bullet *Bullet) outOfScreen() bool {
	return bullet.gameObject.y < -float64(bullet.gameObject.height)
}

func (g *Game) addBullet(bullet *Bullet) {
	g.bullets[bullet] = struct{}{}
}

func NewBullet(cfg *Config, ship *Ship) *Bullet {
	rect := image.Rect(0, 0, cfg.BulletWidth, cfg.BulletHeight)
	img := ebiten.NewImageWithOptions(rect, nil)
	img.Fill(cfg.BulletColor)

	return &Bullet{
		image: img,
		gameObject: GameObject{
			width:  cfg.BulletWidth,
			height: cfg.BulletHeight,
			x:      ship.gameObject.x + float64(ship.gameObject.width-cfg.BulletWidth)/2,
			y:      float64(cfg.ScreenHeight/2 - ship.gameObject.height - cfg.BulletHeight),
		},
		speedFactor: cfg.BulletSpeedFactor,
	}
}

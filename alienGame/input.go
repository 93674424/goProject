package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"time"
)

type Input struct {
	msg            string
	lastBulletTime time.Time
}

func (i *Input) IsKeyPressed() bool {
	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		return true
	}
	return false
}

func (i *Input) IsKeyEnter() bool {
	if ebiten.IsKeyPressed(ebiten.KeyEnter) {
		return true
	}
	return false
}

func (i *Input) Update(ship *Ship, cfg *Config, g *Game) {
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		fmt.Println("←←←←←←←←←←←←←←←←←←")
		i.msg = "left pressed"
		ship.gameObject.x -= cfg.ShipSpeedFactor
		if ship.gameObject.x < -float64(ship.gameObject.width)/2 {
			ship.gameObject.x = -float64(ship.gameObject.width) / 2
		}
	} else if ebiten.IsKeyPressed(ebiten.KeyRight) {
		fmt.Println("→→→→→→→→→→→→→→→→→→")
		i.msg = "right pressed"
		ship.gameObject.x += cfg.ShipSpeedFactor
		if ship.gameObject.x > float64(cfg.ScreenWidth/2)-float64(ship.gameObject.width)/2 {
			ship.gameObject.x = float64(cfg.ScreenWidth/2) - float64(ship.gameObject.width)/2
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		fmt.Println("------------------")
		i.msg = "space pressed"
		if len(g.bullets) < g.cfg.MaxBulletNum &&
			time.Now().Sub(i.lastBulletTime).Milliseconds() > g.cfg.BulletInterval {
			bullet := NewBullet(g.cfg, g.ship)
			g.addBullet(bullet)
			i.lastBulletTime = time.Now()
		}
	}
}

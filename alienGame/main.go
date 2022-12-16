package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"image/color"
	"log"
)

func (g *Game) CheckCollision() {
	for alien := range g.aliens {
		for bullet := range g.bullets {
			if CheckCollision(&alien.gameObject, &bullet.gameObject) {
				delete(g.aliens, alien)
				delete(g.bullets, bullet)
			}
		}
	}
}

func (g *Game) Update() error {
	switch g.mode {
	case ModeTitle:
		if g.input.IsKeyPressed() {
			g.mode = ModeGame
		}
	case ModeGame:
		for bullet := range g.bullets {
			bullet.gameObject.y -= bullet.speedFactor
		}
		for alien := range g.aliens {
			alien.gameObject.y += alien.speedFactor
		}
		g.input.Update(g.ship, g.cfg, g)

		g.CheckCollision()

		for bullet := range g.bullets {
			if bullet.outOfScreen() {
				delete(g.bullets, bullet)
			}
		}
		for alien := range g.aliens {
			if alien.outOfScreen(g.cfg) {
				g.failCount++
				delete(g.aliens, alien)
				continue
			}

			if CheckCollision(&alien.gameObject, &g.ship.gameObject) {
				g.failCount++
				delete(g.aliens, alien)
				continue
			}
		}

		if g.failCount >= 3 {
			g.overMsg = "Game Over!"
		} else if len(g.aliens) == 0 {
			g.overMsg = "You Win!"
		}

		if len(g.overMsg) > 0 {
			g.mode = ModeOver
			g.aliens = make(map[*Alien]struct{})
			g.bullets = make(map[*Bullet]struct{})
		}
	case ModeOver:
		if g.input.IsKeyEnter() {
			g.init()
			g.failCount = 0
			g.overMsg = ""
			g.mode = ModeTitle
		}
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(g.cfg.BgColor)

	var titleTexts []string
	var texts []string
	switch g.mode {
	case ModeTitle:
		titleTexts = []string{"ALIEN INVASION"}
		texts = []string{"", "", "", "", "", "", "", "PRESS SPACE KEY"}
	case ModeGame:
		g.ship.Draw(screen)
		for bullet := range g.bullets {
			bullet.Draw(screen)
		}
		for alien := range g.aliens {
			alien.Draw(screen)
		}
	case ModeOver:
		//texts = []string{"", "GAME OVER!"}
		texts = []string{"", g.overMsg}
	}

	for i, l := range titleTexts {
		x := (g.cfg.ScreenWidth/2 - len(l)*g.cfg.TitleFontSize) / 2
		text.Draw(screen, l, titleArcadeFont, x, (i+4)*g.cfg.TitleFontSize, color.White)
	}
	for i, l := range texts {
		x := (g.cfg.ScreenWidth/2 - len(l)*g.cfg.FontSize) / 2
		text.Draw(screen, l, arcadeFont, x, (i+4)*g.cfg.FontSize, color.White)
	}
	//显示
	ebitenutil.DebugPrint(screen, g.input.msg)
}

func (ship *Ship) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(ship.gameObject.x, ship.gameObject.y)
	screen.DrawImage(ship.image, op)
}

func (bullet *Bullet) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(bullet.gameObject.x, bullet.gameObject.y)
	screen.DrawImage(bullet.image, op)
}

func (alien *Alien) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(alien.gameObject.x, alien.gameObject.y)
	screen.DrawImage(alien.image, op)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	outsideWidth, outsideHeight = g.cfg.ScreenWidth/2, g.cfg.ScreenHeight/2
	return outsideWidth, outsideHeight
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("外星人入侵")
	game := NewGame()
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}

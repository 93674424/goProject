package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"log"
)

const (
	ModeTitle int = iota
	ModeGame
	ModeOver
)

var (
	titleArcadeFont font.Face
	arcadeFont      font.Face
)

type Game struct {
	mode      int
	input     *Input
	ship      *Ship
	cfg       *Config
	bullets   map[*Bullet]struct{}
	aliens    map[*Alien]struct{}
	failCount int
	overMsg   string
}

type GameObject struct {
	width  int
	height int
	x      float64
	y      float64
}

type Entity interface {
	Width() int
	Height() int
	X() float64
	Y() float64
}

func (gameObj *GameObject) Width() int {
	return gameObj.width
}

func (gameObj *GameObject) Height() int {
	return gameObj.height
}

func (gameObj *GameObject) X() float64 {
	return gameObj.x
}

func (gameObj *GameObject) Y() float64 {
	return gameObj.y
}

func (g *Game) addBullet(bullet *Bullet) {
	g.bullets[bullet] = struct{}{}
}

func (g *Game) addAlien(alien *Alien) {
	g.aliens[alien] = struct{}{}
}

func (g *Game) init() {
	g.CreateFonts()
	g.CreateAliens()
}

func (g *Game) CreateFonts() {
	tt, err := opentype.Parse(fonts.PressStart2P_ttf)
	if err != nil {
		log.Fatal(err)
	}
	const dpi = 72
	titleArcadeFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    float64(g.cfg.TitleFontSize),
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}
	arcadeFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    float64(g.cfg.FontSize),
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}
	_, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    float64(g.cfg.SmallFontSize),
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}
}

func (g *Game) CreateAliens() {
	alien := NewAlien(g.cfg)

	availableSpaceX := g.cfg.ScreenWidth/2 - 2*alien.gameObject.width
	numAliens := availableSpaceX / (2 * alien.gameObject.width)

	for row := 0; row < 2; row++ {
		for i := 0; i < numAliens; i++ {
			alien = NewAlien(g.cfg)
			alien.gameObject.x = float64(alien.gameObject.width + 2*alien.gameObject.width*i)
			alien.gameObject.y = float64(alien.gameObject.height*row) * 2
			g.addAlien(alien)
		}
	}
}

func NewGame() *Game {
	cfg := loadConfig()
	ebiten.SetWindowSize(cfg.ScreenWidth, cfg.ScreenHeight)
	ebiten.SetWindowTitle(cfg.Title)

	g := &Game{
		input:   &Input{msg: "Hello, World!"},
		ship:    NewShip(cfg.ScreenWidth/2, cfg.ScreenHeight/2),
		cfg:     cfg,
		bullets: make(map[*Bullet]struct{}),
		aliens:  make(map[*Alien]struct{}),
	}

	g.init()
	return g
}

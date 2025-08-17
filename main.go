package main

import (
	_ "image/jpeg"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type game struct {
	bg        *ebiten.Image
	person    *ebiten.Image
	cat       *ebiten.Image
	messageBG *ebiten.Image
}

func newGame() (*game, error) {
	g := &game{}

	// 画像を読み込む
	img, _, err := ebitenutil.NewImageFromFile("haikei.jpg")
	if err != nil {
		return nil, err
	}
	g.bg = img
	img, _, err = ebitenutil.NewImageFromFile("person.png")
	if err != nil {
		return nil, err
	}
	g.person = img

	img, _, err = ebitenutil.NewImageFromFile("cat.png")
	if err != nil {
		return nil, err
	}
	g.cat = img

	img, _, err = ebitenutil.NewImageFromFile("message_haikei.png")
	if err != nil {
		return nil, err
	}
	g.messageBG = img
	w, h := g.bg.Bounds().Dx(), g.bg.Bounds().Dy()
	ebiten.SetWindowSize(w, h) // ← ウィンドウを画像サイズに
	// もし拡大縮小なしの固定内部解像度にしたいなら：
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeDisabled)

	return g, nil
}

func (g *game) Update() error {
	return nil
}

func (g *game) Draw(screen *ebiten.Image) {
	// 画像を描画
	op := &ebiten.DrawImageOptions{}
	screen.DrawImage(g.bg, op)

	op = &ebiten.DrawImageOptions{}
	op.GeoM.Scale(-1, 1)
	op.GeoM.Translate(1280, 0)
	screen.DrawImage(g.person, op)

	op = &ebiten.DrawImageOptions{}
	// op.GeoM.Scale(0.5, 0.5)
	// op.GeoM.Translate(560, 0)
	screen.DrawImage(g.cat, op)

	op = &ebiten.DrawImageOptions{}

	op.GeoM.Translate(0, 450)
	screen.DrawImage(g.messageBG, op)
}

func (g *game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func main() {
	ebiten.SetWindowTitle("ノベルゲーム")
	ebiten.SetWindowSize(1280, 720)
	g, err := newGame()
	if err != nil {
		panic(err)
	}
	if err := ebiten.RunGame(g); err != nil {
		panic(err)
	}
}

package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"image"
	"time"
	"log"
)

const (
	scale        = 3 // Fator de escala para a imagem
	screenWidth  = 16*20*scale
	screenHeight = 16*12*scale
)

var (
	logicTicker *time.Ticker
	frameTicker *time.Ticker
	animTicker *time.Ticker
	charactersImage *ebiten.Image
	tilesImage *ebiten.Image
	//grassTile *ebiten.Image
	smiley Character
	scout Character
	king Character
	snake Character
	grassPack TilePack
	dirtPack TilePack
	tiles [] *TilePack
	characters[] *Character
	//grassTilemap [][]int
)

type TilePack struct {
	image *ebiten.Image
	tilemap [][2]int
}

type Character struct {
	sheet *ebiten.Image
	index int
	indexRange [2]int
	image *ebiten.Image
	position [2]int
	invert int
}

type Game struct{}

func (g *Game) Update() error {
	select{
	case<-logicTicker.C:
		//select{
		//case<-frameTicker.C:
		if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) {
			scout.position[1]--
		}
		if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) {
			scout.position[1]++
		}
		if inpututil.IsKeyJustPressed(ebiten.KeyArrowLeft) {
			scout.invert = -1
			scout.position[0]--
		}
		if inpututil.IsKeyJustPressed(ebiten.KeyArrowRight) {
			scout.invert = 1
			scout.position[0]++
		}
		//default:
		//}

		var x0, y0 int

		select{
		case <-animTicker.C:
			for _, currentCharacter := range characters{
				currentCharacter.index++
			}
		default:
		}

		for _, currentCharacter := range characters {
			if currentCharacter.position[0] < 0{currentCharacter.position[0] = 0}
			if currentCharacter.position[0] > 19{currentCharacter.position[0] = 19}
			if currentCharacter.position[1] < 1{currentCharacter.position[1] = 1}
			if currentCharacter.position[1] > 11{currentCharacter.position[1] = 11}
			if currentCharacter.index > currentCharacter.indexRange[1]{
				currentCharacter.index = currentCharacter.indexRange[0]
			}
		}

		x0, y0 = scout.index*32, 0
		rect := image.Rect(x0, y0, x0+32, y0+32)
		scout.image = ebiten.NewImageFromImage(scout.sheet.SubImage(rect))

		x0, y0 = smiley.index*32, 0
		rect = image.Rect(x0, y0, x0+32, y0+32)
		smiley.image = ebiten.NewImageFromImage(smiley.sheet.SubImage(rect))

		x0, y0 = king.index*32, 0
		rect = image.Rect(x0, y0, x0+32, y0+32)
		king.image = ebiten.NewImageFromImage(king.sheet.SubImage(rect))

		x0, y0 = snake.index*32, 0
		rect = image.Rect(x0, y0, x0+32, y0+32)
		snake.image = ebiten.NewImageFromImage(snake.sheet.SubImage(rect))
	default:
	}
	return nil
}

/*
func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		scout.position[1]--
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		scout.position[1]++
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		scout.invert = -1
		scout.position[0]--
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		scout.invert = 1
		scout.position[0]++
	}

	select{
	case<-logicTicker.C:
		select{
		case <-animTicker.C:
			for _, currentCharacter := range characters{
				currentCharacter.index++
			}
		default:
		}

		for _, currentCharacter := range characters {
			if currentCharacter.position[0] < 0{currentCharacter.position[0] = 0}
			if currentCharacter.position[0] > 19*16{currentCharacter.position[0] = 19*16}
			if currentCharacter.position[1] < 1*16{currentCharacter.position[1] = 1*16}
			if currentCharacter.position[1] > 11*16{currentCharacter.position[1] = 11*16}
			if currentCharacter.index > currentCharacter.indexRange[1]{
				currentCharacter.index = currentCharacter.indexRange[0]
			}
		}

		var x0, y0 int
		x0, y0 = scout.index*32, 0
		rect := image.Rect(x0, y0, x0+32, y0+32)
		scout.image = ebiten.NewImageFromImage(scout.sheet.SubImage(rect))

		x0, y0 = smiley.index*32, 0
		rect = image.Rect(x0, y0, x0+32, y0+32)
		smiley.image = ebiten.NewImageFromImage(smiley.sheet.SubImage(rect))

		x0, y0 = king.index*32, 0
		rect = image.Rect(x0, y0, x0+32, y0+32)
		king.image = ebiten.NewImageFromImage(king.sheet.SubImage(rect))

		x0, y0 = snake.index*32, 0
		rect = image.Rect(x0, y0, x0+32, y0+32)
		snake.image = ebiten.NewImageFromImage(snake.sheet.SubImage(rect))
	default:
	}
	return nil
}*/

func (g *Game) Draw(screen *ebiten.Image) {
    // Desenha a imagem completa na tela
	for _, i := range tiles{
		for _, pos := range i.tilemap {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(pos[0]*16), float64(pos[1]*16))
			op.GeoM.Scale(scale, scale)
			screen.DrawImage(i.image, op)
		}
	}

	for _, currentCharacter := range characters {
		options := &ebiten.DrawImageOptions{}
		x, y := float64(currentCharacter.position[0]*16)-8, float64(currentCharacter.position[1]*16-16)
		options.GeoM.Scale(float64(currentCharacter.invert), 1)
		if currentCharacter.invert == -1{options.GeoM.Translate(32, 0)}
		options.GeoM.Translate(x, y)
		options.GeoM.Scale(scale, scale)
		screen.DrawImage(currentCharacter.image, options)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
    // Always prints "240 240" on a screen with 125% UI scaling, no matter if the window is 300x300 or 301x301.
    // Therefore there is a loss of information here.
    // fmt.Println(outsideWidth, outsideHeight)

    s := ebiten.DeviceScaleFactor()
    return int(float64(outsideWidth) * s), int(float64(outsideHeight) * s) // Will result in 300x300 with 125% UI scaling.
}

func main() {
    ebiten.SetWindowSize(screenWidth, screenHeight) // Define o tamanho da janela com escala
    ebiten.SetWindowTitle("my game")

	scout = Character{
		index: 0,
		invert: 1,
		indexRange: [2]int{0, 3},
		position: [2]int{0, 2},
	}
	
	smiley = Character{
		index: 0,
		invert: -1,
		indexRange: [2]int{0, 3},
		position: [2]int{4, 2},
	}

	king = Character{
		index: 0,
		invert: -1,
		indexRange: [2]int{0, 3},
		position: [2]int{7, 6},
	}
	
	snake = Character{
		index: 0,
		invert: 1,
		indexRange: [2]int{1, 3},
		position: [2]int{2, 3},
	}
	
	characters = []*Character{
		&scout,
		&smiley,
		&king,
		&snake,
	}

	var err error
	// Carrega o spritesheet dos personagens
	// a função ebitenutil.NewImageFromFile importa a imagem no arquivo especificado no primeiro argumento
	// `charactersImage`, `_` e `err` recebem os valores de retorno da função.
	// _ (underline) serve para ignorar o segundo valor
	charactersImage, _, err = ebitenutil.NewImageFromFile(
		"../../assets/image/characters.png")
	if err != nil {
		log.Fatal(err)
	}

	rect := image.Rect(0, 32*2, 32*23, 32*3)
	scout.sheet = ebiten.NewImageFromImage(charactersImage.SubImage(rect))

	rect = image.Rect(0, 32*0, 32*23, 32*1)
	smiley.sheet = ebiten.NewImageFromImage(charactersImage.SubImage(rect))

	rect = image.Rect(0, 32*1, 32*23, 32*2)
	king.sheet = ebiten.NewImageFromImage(charactersImage.SubImage(rect))
	
	rect = image.Rect(0, 32*3, 32*23, 32*4)
	snake.sheet = ebiten.NewImageFromImage(charactersImage.SubImage(rect))

	tilesImage, _, err = ebitenutil.NewImageFromFile("../../assets/image/sheet.png")
	rect = image.Rect(16*8, 0, 16*9, 16)

	grassPack = TilePack{
		image: ebiten.NewImageFromImage(
			tilesImage.SubImage(image.Rect(16*8, 0, 16*9, 16))),
		tilemap: [][2]int{
			{0,10},{1,10},{2,10},{3,10},{4,10},{5,10},{6,10},{7,10},{8,10},{9,10},{10,10},
		},
	}
	dirtPack = TilePack{
		image: ebiten.NewImageFromImage(
			tilesImage.SubImage(image.Rect(16*7, 16, 16*8, 16*2))),
		tilemap: [][2]int{
			{0,11},{1,11},{2,11},{3,11},{4,11},{5,11},{6,11},{7,11},{8,11},{9,11},{10,11},
		},
	}

	tiles = []*TilePack{&grassPack, &dirtPack}

	logicTicker = time.NewTicker(time.Second/60)
	defer logicTicker.Stop()
	frameTicker = time.NewTicker(time.Second/4)
	defer frameTicker.Stop()
	animTicker = time.NewTicker(time.Second /8)
	defer animTicker.Stop()

    game := &Game{}
    // Inicia o loop do jogo
    if err := ebiten.RunGame(game); err != nil {
        log.Fatal(err)
    }
}

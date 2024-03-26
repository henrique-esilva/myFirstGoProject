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
	grassTile *ebiten.Image
	smiley Character
	scout Character
	king Character
	snake Character
	characters[]*Character
	grassTilemap [][]int
	kingTilemap [][]int

)

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

func (g *Game) Draw(screen *ebiten.Image) {
    // Desenha a imagem completa na tela
    
	for _, pos := range grassTilemap {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(pos[0]*16), float64(pos[1]*16))
		op.GeoM.Scale(scale, scale)
		screen.DrawImage(grassTile, op)
	}

	for _, pos := range kingTilemap {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(pos[0]*16-8), float64(pos[1]*16-16))
		op.GeoM.Scale(scale, scale)
		screen.DrawImage(smiley.image, op)
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

	/*
	options := &ebiten.DrawImageOptions{}

	x, y := float64(scout.position[0]*16)-8, float64(scout.position[1]*16-16)
	options.GeoM.Scale(float64(scout.invert), 1)
	if scout.invert == -1{options.GeoM.Translate(32, 0)}
	options.GeoM.Translate(x, y)
	options.GeoM.Scale(scale, scale)
	screen.DrawImage(scout.image, options)

	options = &ebiten.DrawImageOptions{}

	x, y = float64(king.position[0]*16)-8, float64(king.position[1]*16-16)
	options.GeoM.Scale(float64(king.invert), 1)
	if king.invert == -1{options.GeoM.Translate(32, 0)}
	options.GeoM.Translate(x, y)
	options.GeoM.Scale(scale, scale)
	screen.DrawImage(king.image, options)

	options = &ebiten.DrawImageOptions{}

	x, y = float64(snake.position[0]*16-8), float64(snake.position[1]*16-16)
	options.GeoM.Scale(float64(snake.invert), 1)
	if snake.invert == -1{options.GeoM.Translate(32, 0)}
	options.GeoM.Translate(x, y)
	options.GeoM.Scale(scale, scale)
	screen.DrawImage(snake.image, options)
	*/
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
	kingTilemap = [][]int{{0,5},{1,5},{2,5}}

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
	grassTile = ebiten.NewImageFromImage(tilesImage.SubImage(rect))
	grassTilemap = [][]int{{0,10},{1,10},{2,10},{3,10},{4,10}}

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

package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	//"encoding/json"
	"image"
	"time"
	"log"
	"fmt"
)

const (
	scale        = 2 // Fator de escala para a imagem
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
	tiles [137] *ebiten.Image
	tilemap []int
	mapWidth int
	mapHeight int
	characters[] *Character
)

type Tilemap struct {
	layers []struct {
		data []int `json:"data"`
	} `json:"layers"`
}

type Tilemap2 struct {
	data []int `json:"data"`
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
	//for _, j := range tilemap.layers{ for index, i := range j.data{
	/*for y, i := range tilemap{
		for x, j := range i{
			if j != 0 {
				ind := j-1
				pos := [2]int{x, y}
				op := &ebiten.DrawImageOptions{}
				op.GeoM.Translate(float64(pos[0]*16), float64(pos[1]*16))
				op.GeoM.Scale(scale, scale)
				screen.DrawImage(tiles[ind], op)
			}
		}
	}*/

	for ind, imgind := range tilemap{
		if imgind != 0 {
			corr := imgind-1
			pos := [2]int{ind%20, ind/20}
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(pos[0]*16), float64(pos[1]*16))
			op.GeoM.Scale(scale, scale)
			screen.DrawImage(tiles[corr], op)
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

func resumeRect (x, y int) image.Rectangle {
	x, y = 16*(x), 16*(y)
	a := x+16
	b := y+16
	return image.Rect(x, y, a, b)
}

func main() {
    ebiten.SetWindowSize(screenWidth, screenHeight) // Define o tamanho da janela com escala
    ebiten.SetWindowTitle("my game")

	scout = Character{
		index: 0,
		invert: 1,
		indexRange: [2]int{0, 3},
		position: [2]int{1, 9},
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
		position: [2]int{8, 9},
	}
	
	snake = Character{
		index: 0,
		invert: -1,
		indexRange: [2]int{1, 3},
		position: [2]int{12, 7},
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

	// tiles com 136 de comprimento
	/*for x:=0;x<16;x++{
		for y:=0;y<7;y++{
			tiles[x][y]=ebiten.NewImageFromImage(tilesImage.SubImage(resumeRect(x, y)))
		}
	}*/
	for x := 0; x < 17; x++ {
		for y:=0;y<8;y++{
			tiles[y*17+x]=ebiten.NewImageFromImage(tilesImage.SubImage(resumeRect(x, y)))
		}
	}
	fmt.Println("%+v\n", tiles)

	tilemap=[]int{
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 49, 50, 51, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 11, 12, 13, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 49, 50, 0, 0, 28, 25, 30, 0, 0, 0, 0, 0, 0, 0,
	0, 11, 12, 13, 0, 11, 12, 12, 13, 0, 28, 29, 30, 0, 0, 0, 0, 0, 0, 0,
	0, 28, 29, 30, 0, 28, 29, 25, 30, 0, 28, 29, 30, 0, 0, 0, 0, 0, 0, 0}

	// index 11 at 170th position
	// 10, 8

	//var tilemap Tilemap2
	/*err = json.Unmarshal([]byte(myJson), &tilemap)
    if err != nil {
        fmt.Println("Erro ao decodificar JSON:", err)
        return
    }*/

	fmt.Printf("%+v\n", tilemap)

	// tiles = [10][7]TilePack{}

	// for x:=0; x<10; x++{
	// 	for y:=0; y<7; y++{
	// 		tiles[x][y] = TilePack{
	// 			image:ebiten.NewImageFromImage(tilesImage.SubImage(resumeRect(x, y))),
	// 		}
	// 	}
	// }

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

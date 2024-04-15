package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	//"github.com/hajimehoshi/ebiten/v2/inpututil"
	"encoding/json"
	"io/ioutil"
	"image"
	"time"
	"log"
	"fmt"
)

const (
	scale        = 2 // Fator de escala para a imagem
	screenWidth  = 320*scale
	screenHeight = 192*scale
)

var (
	logicTicker *time.Ticker
	frameTicker *time.Ticker
	animTicker *time.Ticker
	fallTicker *time.Ticker
	charactersImage *ebiten.Image
	tilesImage *ebiten.Image
	//grassTile *ebiten.Image
	smiley Character
	scout Character
	king Character
	snake Character
	tiles [137*2] *ebiten.Image
	colisionTerrain [20][13]int
	tilemap Tilemap
	mapWidth int
	mapHeight int
	characters[] *Character
)

//type JsonLayers struct {
//	Data []int `json:"data"`
//}

type Tilemap struct {
	Layers []struct {
		Data []int `json:"data"`
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
	jump int
	landed int
}

type Game struct{}

func (g *Game) Update() error {
	select{
	case<-logicTicker.C:
		//select{
		//case<-frameTicker.C:
		/*if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		inpututil.IsKeyJustPressed(ebiten.)
			scout.position[1]++
		}*/
		limitx, limity := 19, 11
		if scout.landed == 1 {//position[1] >= limity{
			if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
				scout.landed = 0
				scout.jump = 3
			}
		}
		// if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		// 	scout.invert = -1
		// 	scout.position[0]--
		// }
		// if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		// 	scout.invert = 1
		// 	scout.position[0]++
		// }
		//default:
		//}

		var x0, y0 int

		select{
		case <-animTicker.C:
			for _, currentCharacter := range characters{
				currentCharacter.index++
				if currentCharacter.index > currentCharacter.indexRange[1]{
					currentCharacter.index = currentCharacter.indexRange[0]
				}
			}
		default:
		}

		select{
		case <-fallTicker.C:
			if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
				scout.invert = -1
				scout.position[0]--
			}
			if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
				scout.invert = 1
				scout.position[0]++
			}
			for _, currentCharacter := range characters {
				if currentCharacter.position[0] < 0{currentCharacter.position[0] = 0}
				if currentCharacter.position[0] > limitx{currentCharacter.position[0] = limitx}
				if currentCharacter.position[1] < 1{currentCharacter.position[1] = 1}
				if currentCharacter.position[1] > limity{currentCharacter.position[1] = limity}

				if currentCharacter.jump > 0 {
					currentCharacter.position[1]--
					currentCharacter.jump--
				} else {
					if currentCharacter.position[1] <= limity {
						currentCharacter.position[1]++
						if colisionTerrain[currentCharacter.position[0]][currentCharacter.position[1]] == 1 || currentCharacter.position[1] > limity {
							currentCharacter.position[1]--
							currentCharacter.landed = 1
						} else {currentCharacter.landed = 0}
					} else {currentCharacter.landed = 1}
				}
			}
		default:
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

	/*for ind, imgind := range tilemap{
		if imgind != 0 {
			corr := imgind-1
			pos := [2]int{ind%20, ind/20}
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(pos[0]*16), float64(pos[1]*16))
			op.GeoM.Scale(scale, scale)
			screen.DrawImage(tiles[corr], op)
		}
	}*/

	for _, layer := range tilemap.Layers {
		for ind, imgind := range layer.Data{
			if imgind != 0 {
				corr := imgind//-1
				pos := [2]int{ind%20, ind/20}
				op := &ebiten.DrawImageOptions{}
				op.GeoM.Translate(float64(pos[0]*16), float64(pos[1]*16))
				op.GeoM.Scale(scale, scale)
				screen.DrawImage(tiles[corr], op)
			}
		}
	}

	for _, currentCharacter := range characters {
		options := &ebiten.DrawImageOptions{}
		//x, y := float64(currentCharacter.position[0]*16)-8, float64(currentCharacter.position[1]*16-16)
		x, y := float64(currentCharacter.position[0]*16)-8, float64(currentCharacter.position[1]*16)-16
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

    //s := ebiten.DeviceScaleFactor()
    return int(float64(outsideWidth)), int(float64(outsideHeight)) // Will result in 300x300 with 125% UI scaling.
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
		position: [2]int{1, 0},
	}
	
	smiley = Character{
		index: 0,
		invert: -1,
		indexRange: [2]int{0, 3},
		position: [2]int{4*16, 2*16},
	}

	king = Character{
		index: 0,
		invert: -1,
		indexRange: [2]int{0, 3},
		position: [2]int{8*16, 9*16},
	}
	
	snake = Character{
		index: 0,
		invert: -1,
		indexRange: [2]int{1, 3},
		position: [2]int{12*16, 7*16},
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

	// an array tiles holds all tile's images
	for x := 0; x < 17; x++ {
		for y:=0;y<8;y++{
			tiles[y*17+x+1]=ebiten.NewImageFromImage(tilesImage.SubImage(resumeRect(x, y)))
		}
	}

	tilesImage, _, err = ebitenutil.NewImageFromFile("../../assets/image/sheet2.png")
	for x := 0; x < 17; x++ {
		for y:=0;y<8;y++{
			tiles[y*17+x+137]=ebiten.NewImageFromImage(tilesImage.SubImage(resumeRect(x, y)))
		}
	}
	//rect = image.Rect(16*8, 0, 16*9, 16)

	// import game map from archives
	myJson, err := ioutil.ReadFile("../../map/meuMapa.json")
	if err != nil {
		fmt.Println("Erro ao ler o arquivo:", err)
		return
	}
	fmt.Println(string(myJson))

	var tt Tilemap

	err = json.Unmarshal(myJson, &tt)
	if err != nil {
		fmt.Println("Erro ao decodificar o JSON:", err)
		return
	}

	tilemap = tt

	colisionTerrain[1][10] = 1
	colisionTerrain[2][10] = 1
	colisionTerrain[3][10] = 1
	colisionTerrain[1][11] = 1
	colisionTerrain[2][11] = 1
	colisionTerrain[3][11] = 1

	colisionTerrain[5][10] = 1
	colisionTerrain[6][10] = 1
	colisionTerrain[7][10] = 1
	colisionTerrain[8][10] = 1
	colisionTerrain[5][11] = 1
	colisionTerrain[6][11] = 1
	colisionTerrain[7][11] = 1
	colisionTerrain[8][11] = 1

	colisionTerrain[10][11] = 1
	colisionTerrain[11][11] = 1
	colisionTerrain[12][11] = 1
	colisionTerrain[10][10] = 1
	colisionTerrain[11][10] = 1
	colisionTerrain[12][10] = 1
	colisionTerrain[10][ 9] = 1
	colisionTerrain[11][ 9] = 1
	colisionTerrain[12][ 9] = 1
	colisionTerrain[10][ 8] = 1
	colisionTerrain[11][ 8] = 1
	colisionTerrain[12][ 8] = 1

	colisionTerrain[16][ 8] = 1
	colisionTerrain[17][ 8] = 1
	colisionTerrain[18][ 8] = 1
	colisionTerrain[19][ 8] = 1

	colisionTerrain[16][ 9] = 1
	colisionTerrain[17][ 9] = 1
	colisionTerrain[18][ 9] = 1
	colisionTerrain[19][ 9] = 1

	colisionTerrain[16][10] = 1
	colisionTerrain[17][10] = 1
	colisionTerrain[18][10] = 1
	colisionTerrain[19][10] = 1

	colisionTerrain[16][11] = 1
	colisionTerrain[17][11] = 1
	colisionTerrain[18][11] = 1
	colisionTerrain[19][11] = 1

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
	fallTicker = time.NewTicker(time.Second/2)
	defer fallTicker.Stop()
	frameTicker = time.NewTicker(time.Second/4)
	defer frameTicker.Stop()
	animTicker = time.NewTicker(time.Second /2)
	defer animTicker.Stop()

    game := &Game{}
    // Inicia o loop do jogo
    if err := ebiten.RunGame(game); err != nil {
        log.Fatal(err)
    }
}

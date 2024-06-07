package main

import (
	"fmt"
	"image"
	"image/color"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

var (
	screenWidth                          = 300
	screenHeight                         = 300
	screenWidth2x                        = screenWidth * 2
	screenHeight2x                       = screenHeight * 2
	bgColor                   color.RGBA = color.RGBA{255, 133, 69, 0}
	TXT                       string     = "Orange"
	count                     int
	CountPlusWindows          int
	CountNegativeWindows      int
	inGame                    bool
	QuesNumber                int
	score                     int
	rand1                     int    = rand.Intn(15)
	rand2                     int    = rand.Intn(15)
	correctAnswer             int    = rand1 * rand2
	wrongAnswer               int    = rand1*rand2 + rand.Intn(20-1+1) + 1
	showQues                  string = fmt.Sprint(rand1, "*", rand2, " = ?")
	showCorrectAnswerForQues1 string = fmt.Sprint("1- ", correctAnswer)
	showWrongAnswerForQues1   string = fmt.Sprint("2- ", wrongAnswer)
	showWrongAnswerForQues2   string = fmt.Sprint("1- ", wrongAnswer)
	showCorrectAnswerForQues2 string = fmt.Sprint("2- ", correctAnswer)
	falseInQues1TrueInQues2   bool
	randomQues                int = rand.Intn(2-1+1) + 1
)

type Game struct{}

func (g *Game) Update() error {
	if !inGame {
		if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
			switch count {
			case 0:
				bgColor = color.RGBA{R: 255, G: 0, B: 0, A: 255}
				TXT = "Red"
				count++
			case 1:
				bgColor = color.RGBA{0, 255, 0, 0}
				TXT = "Green"
				count++
			case 2:
				bgColor = color.RGBA{0, 0, 255, 0}
				TXT = "Blue"
				count++
			case 3:
				bgColor = color.RGBA{48, 0, 92, 0}
				TXT = "Purple"
				count++
			default:
				bgColor = color.RGBA{255, 133, 69, 0}
				TXT = "Orange"
				count = 0

			}
		}
		if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) {
			if CountPlusWindows < 23 {
				CountPlusWindows++
				screenWidth2x += 10
				screenHeight2x += 10
				ebiten.SetWindowSize(screenWidth2x, screenHeight2x)
			}
		}
		if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) {
			if CountNegativeWindows < 23 {
				CountPlusWindows--
				CountNegativeWindows++
				screenWidth2x -= 10
				screenHeight2x -= 10
				ebiten.SetWindowSize(screenWidth2x, screenHeight2x)
			}
		}
		if inpututil.IsKeyJustPressed(ebiten.KeyZ) {
			fmt.Println("Started")
			inGame = true
		}
	} else {
		switch QuesNumber {
		default:
			if !falseInQues1TrueInQues2 {
				checkCorrectOrWrongAnswerQues1()
			} else {
				checkCorrectOrWrongAnswerQues2()
			}
		} //INGAME
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	if !inGame {
		currentTPS := int(ebiten.CurrentTPS())
		tps := fmt.Sprint("TPS: ", currentTPS)
		ColorBG := fmt.Sprint("Press Enter For Change Color\n Color:", TXT)

		screen.Fill(bgColor)
		Line(screen)
		ebitenutil.DebugPrint(screen, tps)
		ebitenutil.DebugPrintAt(screen, "Press Down For Smaller\nPress Up For Bigger", 0, 20)
		ebitenutil.DebugPrintAt(screen, ColorBG, 0, 60)
		ebitenutil.DebugPrintAt(screen, "Press Z For Start", 0, 100)

	} else { //INGAME
		currentTPS := int(ebiten.CurrentTPS())
		tps := fmt.Sprint("TPS: ", currentTPS)
		scoreGame := fmt.Sprint("Score: ", score)

		screen.Fill(bgColor)
		ebitenutil.DebugPrintAt(screen, tps, 10, 11)
		ebitenutil.DebugPrintAt(screen, scoreGame, 233, 11)
		Line(screen)
		gameQues(screen)
	}

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenWidth
}

func Line(screen *ebiten.Image) {
	if !inGame {
		ebitenutil.DrawRect(screen, 0, 270, 300, 8, color.RGBA{210, 255, 50, 5})
	} else {
		ebitenutil.DrawRect(screen, 0, 290, 300, 7, color.RGBA{210, 255, 50, 5})
		ebitenutil.DrawRect(screen, 290, 0, 7, 300, color.RGBA{210, 255, 50, 5})
		ebitenutil.DrawRect(screen, 0, 5, 300, 7, color.RGBA{210, 255, 50, 5})
		ebitenutil.DrawRect(screen, 5, 0, 7, 300, color.RGBA{210, 255, 50, 5})
	}
}

func gameQues(screen *ebiten.Image) {
	switch QuesNumber {
	default:
		randomQuesFunc(screen)
	}
}

func createQues1(screen *ebiten.Image) { // correct answer on 1
	falseInQues1TrueInQues2 = false

	ebitenutil.DebugPrintAt(screen, showQues, 114, 90)
	ebitenutil.DebugPrintAt(screen, showCorrectAnswerForQues1, 96, 120)
	ebitenutil.DebugPrintAt(screen, showWrongAnswerForQues1, 148, 120)
	ebitenutil.DebugPrintAt(screen, "Enter Key of Your Answer", 76, 146)
	ebitenutil.DrawRect(screen, 60, 80, 175, 3, color.Black)
	ebitenutil.DrawRect(screen, 60, 180, 175, 3, color.Black)
	ebitenutil.DrawRect(screen, 60, 80, 3, 100, color.Black)
	ebitenutil.DrawRect(screen, 235, 80, 3, 103, color.Black)
}

func createQues2(screen *ebiten.Image) { // correct answer on 2
	falseInQues1TrueInQues2 = true

	ebitenutil.DebugPrintAt(screen, showQues, 114, 90)
	ebitenutil.DebugPrintAt(screen, showWrongAnswerForQues2, 96, 120)
	ebitenutil.DebugPrintAt(screen, showCorrectAnswerForQues2, 148, 120)
	ebitenutil.DebugPrintAt(screen, "Enter Key of Your Answer", 76, 146)
	ebitenutil.DrawRect(screen, 60, 80, 175, 3, color.Black)
	ebitenutil.DrawRect(screen, 60, 180, 175, 3, color.Black)
	ebitenutil.DrawRect(screen, 60, 80, 3, 100, color.Black)
	ebitenutil.DrawRect(screen, 235, 80, 3, 103, color.Black)

}

func checkCorrectOrWrongAnswerQues1() {
	if inpututil.IsKeyJustPressed(ebiten.Key1) {
		QuesNumber++
		score++
		createNewRandQues1()
	}
	if inpututil.IsKeyJustPressed(ebiten.Key2) {
		QuesNumber++
		if score != 0 {
			score--
		}
		createNewRandQues1()
	}
}

func checkCorrectOrWrongAnswerQues2() {
	if inpututil.IsKeyJustPressed(ebiten.Key1) {
		QuesNumber++
		if score != 0 {
			score--
		}
		createNewRandQues2()
	}
	if inpututil.IsKeyJustPressed(ebiten.Key2) {
		QuesNumber++
		score++
		createNewRandQues2()
	}
}

func createNewRandQues1() {
	rand1 = rand.Intn(15)
	rand2 = rand.Intn(15)
	correctAnswer = rand1 * rand2
	wrongAnswer = rand1*rand2 + rand.Intn(20-4+1) + 4
	showQues = fmt.Sprint(rand1, "*", rand2, " = ?")
	showCorrectAnswerForQues1 = fmt.Sprint("1- ", correctAnswer, "")
	showWrongAnswerForQues1 = fmt.Sprint("2- ", wrongAnswer, "")
	randomQues = rand.Intn(2-1+1) + 1
	fmt.Println(wrongAnswer, "wrong")
	fmt.Println(correctAnswer, "corect")
	showWrongAnswerForQues2 = fmt.Sprint("1- ", wrongAnswer, "")
	showCorrectAnswerForQues2 = fmt.Sprint("2- ", correctAnswer, "")
}

func createNewRandQues2() {
	rand1 = rand.Intn(15)
	rand2 = rand.Intn(15)
	correctAnswer = rand1 * rand2
	wrongAnswer = rand1*rand2 + rand.Intn(20-4+1) + 4
	showQues = fmt.Sprint(rand1, "*", rand2, " = ?")
	showWrongAnswerForQues2 = fmt.Sprint("1- ", wrongAnswer, "")
	showCorrectAnswerForQues2 = fmt.Sprint("2- ", correctAnswer, "")
	randomQues = rand.Intn(2-1+1) + 1
	fmt.Println(wrongAnswer, "wrong")
	fmt.Println(correctAnswer, "corect")
	showWrongAnswerForQues2 = fmt.Sprint("1- ", wrongAnswer, "")
	showCorrectAnswerForQues2 = fmt.Sprint("2- ", correctAnswer, "")
}

func randomQuesFunc(screen *ebiten.Image) {
	if randomQues == 1 {
		createQues1(screen)
	} else {
		createQues2(screen)
	}
}

func main() {
	icon := ebiten.NewImage(16, 16)
	icon.Fill(color.RGBA{50, 210, 30, 0})
	ebiten.SetWindowResizable(true)
	ebiten.SetCursorMode(ebiten.CursorModeHidden)
	ebiten.SetWindowIcon([]image.Image{icon})
	ebiten.SetWindowSize(screenWidth2x, screenHeight2x)
	ebiten.SetWindowTitle("Math Game")
	err := ebiten.RunGame(&Game{})
	if err != nil {
		panic(err)
	}
}

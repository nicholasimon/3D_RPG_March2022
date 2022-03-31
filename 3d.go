package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var (

	//CORE CORE CORE CORE CORE CORE CORE CORE CORE CORE

	//DEV
	dev, dev2, dev3 bool

	//GRID
	grid                  []xgrid
	gridw                 = 150
	grida                 = gridw * gridw
	boundboxsize          = float32(0.04)
	drawblok, drawbloknxt int
	drawsizew             = 30
	drawsizea             = drawsizew * drawsizew
	//TILES
	basetile = float32(16)
	multi    = float32(3)
	tilesize = basetile * multi
	//SCREEN
	scrwf32, scrhf32, frames32 float32
	scrw, scrh                 int32
	scrhint, scrwint, frames   int

	//CAMERAS
	cam2d     rl.Camera2D
	cam3d     rl.Camera
	cam3dorig rl.Vector3

	xcam, ycam, zcam float32
	//INP
	mousev2   rl.Vector2
	mouseblok int
	selpoint  rl.Vector3
	selblok   int
	//TXT
	txts   = int32(10)
	txtm   = int32(20)
	txtl   = int32(40)
	txtdef = txtm

	//TIMERS
	clickpause int32

	fps = int32(60)

	fadeblink  = float32(0.2)
	fadeblink2 = float32(0.2)

	fadeblinkon, fadeblink2on, onoff1, onoff2, onoff3, onoff6, onoff10, onoff15, onoff30, onoff60 bool

	//BLANKS
	blankv3 = rl.NewVector3(777777777777, 777777777777, 777777777777)

	//IMGS
	imgs     rl.Texture2D
	tileimgs []rl.Rectangle
)

//MARK: STRUCTS
type xgrid struct {
}
type xmodel struct {
	pos rl.Vector3

	model rl.Model
}

//MARK: CAMS CAMS CAMS CAMS CAMS CAMS CAMS CAMS CAMS CAMS CAMS CAMS CAMS CAMS CAMS CAMS CAMS CAMS
func camthreed() { //MARK: camthreed

	rl.BeginMode3D(cam3d)

	rl.EndMode3D()
}

func nocam() { //MARK: nocam

	//closewin
	if closewinloc(1860, 32, brightred(), brightyellow()) {
		rl.CloseWindow()
	}
}
func devui() { //MARK: devui

	siderec := rl.NewRectangle(0, 0, 300, 1080)

	rl.DrawRectangleRec(siderec, rl.Fade(rl.Green, 0.5))

	x := int32(siderec.X + 10)
	y := int32(10)

	txt := "gridw"
	rl.DrawText(txt, x, y, txts, rl.White)
	x += rl.MeasureText(txt, txts) + txts
	txt = fmt.Sprint(gridw)
	rl.DrawText(txt, x, y, txts, rl.White)
	x = int32(siderec.X + 10)
	y += txts

	txt = "grida"
	rl.DrawText(txt, x, y, txts, rl.White)
	x += rl.MeasureText(txt, txts) + txts
	txt = fmt.Sprint(grida)
	rl.DrawText(txt, x, y, txts, rl.White)
	x = int32(siderec.X + 10)
	y += txts

}

//MARK: UPDATE UPDATE UPDATE UPDATE UPDATE UPDATE UPDATE UPDATE UPDATE UPDATE UPDATE UPDATE UPDATE

func upplayer() { //MARK: upplayer

}
func upcams() { //MARK: upcams

	//cam3d.Target.X = player.pos.X
	//	cam3d.Target.Z = player.pos.Z
	//	cam3d.Position = rl.NewVector3(player.pos.X+xcam, player.pos.Y+ycam, player.pos.Z+zcam)

}
func update() { //MARK: update

	inp()
	timers()
	upplayer()
	upcams()

}

// MARK: FUNCTIONS FUNCTIONS FUNCTIONS FUNCTIONS FUNCTIONS FUNCTIONS FUNCTIONS FUNCTIONS FUNCTIONS

//MARK: MAKE MAKE MAKE MAKE MAKE MAKE MAKE MAKE MAKE MAKE MAKE MAKE MAKE MAKE MAKE MAKE MAKE MAKE

func makegrid() { //MARK: makegrid

}
func makeplayer() { //MARK: makeplayer

}
func makeimgs() { //MARK: makeimgs

	x := float32(11)
	y := float32(4)

	for {

		tileimgs = append(tileimgs, rl.NewRectangle(x, y, 16, 16))

		x += 24
		if x > 278 {
			x = 11
			y += 24
		}
		if y > 295 {
			break
		}
	}

	x = float32(309)
	y = float32(44)

}

//MARK:  CORE CORE CORE CORE CORE CORE CORE CORE CORE CORE CORE CORE CORE CORE CORE CORE CORE CORE

func raylib() { //MARK: raylib

	rl.InitWindow(scrw, scrh, "GAME TITLE")
	rl.SetExitKey(rl.KeyEnd) // key to end the game and close window
	rl.HideCursor()

	imgs = rl.LoadTexture("imgs.png") // load images
	makeimgs()

	initial()

	rl.SetTargetFPS(fps)

	for !rl.WindowShouldClose() {

		frames++
		mousev2 = rl.GetMousePosition()

		rl.BeginDrawing()

		rl.ClearBackground(rl.Black)

		camthreed()

		nocam()
		if dev {
			devui()
		}

		//cursor
		cursorv1 := rl.NewVector2(mousev2.X+20, mousev2.Y+5)
		cursorv2 := rl.NewVector2(mousev2.X+5, mousev2.Y+20)
		rl.DrawTriangle(cursorv1, mousev2, cursorv2, rl.Fade(rl.Magenta, fadeblink))
		rl.DrawTriangleLines(cursorv1, mousev2, cursorv2, rl.Black)

		rl.EndDrawing()
		update()
	}

	rl.CloseWindow()

}
func initial() { //MARK: initial

	selpoint = blankv3

	//	makemap()

}
func main() { //MARK: main
	rand.Seed(time.Now().UnixNano())     // random numbers
	rl.SetTraceLog(rl.LogError)          // hides info window
	rl.SetConfigFlags(rl.FlagMsaa4xHint) // enable 4X anti-aliasing

	scr(0)
	cams()
	raylib()

}
func cams() { //MARK: update
	cam3d.Position = rl.NewVector3(45, 35, 60)
	cam3d.Target = rl.NewVector3(0.0, 0.0, 0.0)
	cam3d.Up = rl.NewVector3(0.0, 1.0, 0.0)
	cam3d.Fovy = 30.0
	//cam3d.Projection = rl.CameraProjection(rl.CameraPerspective)
	cam3dorig = cam3d.Position

}

func inp() { //MARK: inp

	//cams

	if rl.IsKeyPressed(rl.KeyKpAdd) {

	}

	//dev
	if rl.IsKeyPressed(rl.KeyF1) {
		if dev {
			dev = false
		} else {
			dev = true
		}

	}
	if rl.IsKeyPressed(rl.KeyF2) {
		if dev2 {
			dev2 = false
		} else {
			dev2 = true
		}
	}
	if rl.IsKeyPressed(rl.KeyF3) {
		if dev3 {
			dev3 = false
		} else {
			dev3 = true
		}
	}

}
func scr(num int) { //MARK: scr
	switch num {
	case 0:
		scrh = int32(rl.GetScreenHeight())
		scrw = int32(rl.GetScreenWidth())
	case 1:
		scrh = int32(rl.GetScreenHeight())
		scrw = int32(rl.GetScreenWidth())

	}

	scrhf32 = float32(1080)
	scrwf32 = float32(1920)
	scrhint = int(scrhf32)
	scrwint = int(scrwf32)

}
func closewinloc(x, y float32, col, col2 rl.Color) bool { //MARK: closewinloc

	close := false

	rec := rl.NewRectangle(x, y, 32, 32)
	if rl.CheckCollisionPointRec(mousev2, rec) {
		rl.DrawRectangleRec(rec, col2)
		if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
			close = true
		}
	} else {
		rl.DrawRectangleRec(rec, col)
	}
	return close

}
func timers() { //MARK: timers
	if clickpause != 0 {
		clickpause--
		if clickpause < 0 {
			clickpause = 0
		}
	}

	if frames%1 == 0 {
		if onoff1 {
			onoff1 = false
		} else {
			onoff1 = true
		}
	}

	if frames%2 == 0 {
		if onoff2 {
			onoff2 = false
		} else {
			onoff2 = true
		}
	}
	if frames%3 == 0 {
		if onoff3 {
			onoff3 = false
		} else {
			onoff3 = true
		}
	}
	if frames%6 == 0 {
		if onoff6 {
			onoff6 = false
		} else {
			onoff6 = true
		}
	}
	if frames%10 == 0 {
		if onoff10 {
			onoff10 = false
		} else {
			onoff10 = true
		}
	}
	if frames%15 == 0 {
		if onoff15 {
			onoff15 = false
		} else {
			onoff15 = true
		}
	}
	if frames%30 == 0 {
		if onoff30 {
			onoff30 = false
		} else {
			onoff30 = true
		}
	}
	if frames%60 == 0 {
		if onoff60 {
			onoff60 = false
		} else {
			onoff60 = true
		}
	}
	if fadeblinkon {
		if fadeblink > 0.4 {
			fadeblink -= 0.05
		} else {
			fadeblinkon = false
		}
	} else {
		if fadeblink < 0.9 {
			fadeblink += 0.05
		} else {
			fadeblinkon = true
		}
	}
}

// MARK: colors
// https://www.rapidtables.com/web/color/RGB_Color.html
func darkred() rl.Color {
	color := rl.NewColor(55, 0, 0, 255)
	return color
}
func semidarkred() rl.Color {
	color := rl.NewColor(70, 0, 0, 255)
	return color
}
func brightorange() rl.Color {
	color := rl.NewColor(253, 95, 0, 255)
	return color
}
func brightred() rl.Color {
	color := rl.NewColor(230, 0, 0, 255)
	return color
}
func randomgrey() rl.Color {
	color := rl.NewColor(uint8(rInt(160, 193)), uint8(rInt(160, 193)), uint8(rInt(160, 193)), uint8(rInt(0, 255)))
	return color
}
func randombluelight() rl.Color {
	color := rl.NewColor(uint8(rInt(0, 180)), uint8(rInt(120, 256)), uint8(rInt(120, 256)), 255)
	return color
}
func randombluedark() rl.Color {
	color := rl.NewColor(0, 0, uint8(rInt(120, 250)), 255)
	return color
}
func randomyellow() rl.Color {
	color := rl.NewColor(255, uint8(rInt(150, 256)), 0, 255)
	return color
}
func randomorange() rl.Color {
	color := rl.NewColor(uint8(rInt(250, 256)), uint8(rInt(60, 210)), 0, 255)
	return color
}
func randomred() rl.Color {
	color := rl.NewColor(uint8(rInt(128, 256)), uint8(rInt(0, 129)), uint8(rInt(0, 129)), 255)
	return color
}
func randomgreen() rl.Color {
	color := rl.NewColor(uint8(rInt(0, 170)), uint8(rInt(100, 256)), uint8(rInt(0, 50)), 255)
	return color
}
func randomcolor() rl.Color {
	color := rl.NewColor(uint8(rInt(0, 256)), uint8(rInt(0, 256)), uint8(rInt(0, 256)), 255)
	return color
}
func brightyellow() rl.Color {
	color := rl.NewColor(uint8(255), uint8(255), uint8(0), 255)
	return color
}
func brightbrown() rl.Color {
	color := rl.NewColor(uint8(218), uint8(165), uint8(32), 255)
	return color
}
func brightgrey() rl.Color {
	color := rl.NewColor(uint8(212), uint8(212), uint8(213), 255)
	return color
}

// MARK: random numbers
func rInt(min, max int) int {
	return rand.Intn(max-min) + min
}
func rInt32(min, max int) int32 {
	i := rand.Intn(max-min) + min
	return int32(i)
}
func rFloat32(min, max float32) float32 {
	return (rand.Float32() * (max - min)) + min
}
func flipcoin() bool {
	var b bool
	a := rInt(0, 10001)
	if a < 5000 {
		b = true
	}
	return b
}
func rolldice() int {
	a := rInt(1, 7)
	return a
}

//MARK: other functions
func lastdigits(num int) int {
	number := num % 1e2 //change 1e2 to 1e3 to 1e4 etc for more digit places
	return number
}
func firstdigits(num int) int {
	number := num / 1e3 //change 1e2 to 1e3 to 1e4 etc for more digit places
	return number
}
func timehere(x, y float32) {
	currentTime := time.Now()
	txtlen := rl.MeasureText(currentTime.Format("15:04"), txtdef)
	x -= float32(txtlen + txtdef)
	rl.DrawText(currentTime.Format("15:04"), int32(x), int32(y), txtdef, rl.White)
}
func getabs(num float32) float32 {
	return float32(math.Abs(float64(num)))
}
func absdiff32(num, num2 float32) float32 {

	diff := float32(0)
	if num == num2 {
		diff = 0
	} else if num == 0 || num2 == 0 {
		if num == 0 {
			diff = float32(math.Abs(float64(num2)))
		} else {
			diff = float32(math.Abs(float64(num)))
		}
	} else if num > 0 && num2 > 0 {
		if num > num2 {
			diff = num - num2
		} else {
			diff = num2 - num
		}
	} else if num > 0 && num2 < 0 || num < 0 && num2 > 0 {

		if num > 0 {
			diff = num + float32(math.Abs(float64(num2)))
		} else {
			diff = num2 + float32(math.Abs(float64(num)))
		}

	} else if num < 0 && num2 < 0 {
		num = float32(math.Abs(float64(num)))
		num2 = float32(math.Abs(float64(num2)))
		if num > num2 {
			diff = num - num2
		} else {
			diff = num2 - num
		}
	}

	return diff

}
func angle2points(start, destination rl.Vector2) float32 { //make sure destination vector is vec2
	angle := float32(math.Atan2(float64(destination.Y-start.Y), float64(destination.X-start.X)))*(180/math.Pi) + 90
	//change +30 (addition value at end) to angle to compensate for polygon rotation difference
	return angle

}

/*
func remblok(s []xblok, index int) []xblok { //remove struct from a slice
	return append(s[:index], s[index+1:]...)
}
*/
func remstring(s []string, index int) []string { //remove string from a slice
	return append(s[:index], s[index+1:]...)
}
func diagsquare(sidelength float32) float32 {
	return sidelength * float32(math.Sqrt(2))
}
func circlearea(radius float32) float32 {
	return math.Pi * radius * radius
}

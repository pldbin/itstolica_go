package gui

// import (
// 	"fyne.io/fyne/v2/app"
// 	"fyne.io/fyne/v2/container"
// 	"fyne.io/fyne/v2/widget"
// )
import (
	"fmt"

	"github.com/pldbin/prog1/internal/server"
)

var I = 9

func FuncAddT() {
	k := server.AddT(I)
	fmt.Println(k)

}

// func ShowApp() {
// 	a := app.New()
// 	w := a.NewWindow("Hello")

// 	hello := widget.NewLabel("Hello Fyne!")
// 	w.SetContent(container.NewVBox(
// 		hello,
// 		widget.NewButton("Hi!", func() {
// 			hello.SetText("Welcome :)")
// 		}),
// 	))

// 	w.ShowAndRun()
// }

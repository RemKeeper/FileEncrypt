package main

import (
	_ "embed"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

//go:embed FileEncrypt.svg
var icon []byte

var (
	Windows fyne.Window
)

func main() {
	App := app.New()
	App.SetIcon(fyne.NewStaticResource("FileEncrypt.svg", icon))
	Windows = App.NewWindow("文件加密")
	Windows.Resize(fyne.NewSize(600, 600))
	Windows.SetContent(MainUI())
	Windows.ShowAndRun()
}

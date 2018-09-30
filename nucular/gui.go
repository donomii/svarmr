package main

import (
	
	"github.com/donomii/svarmrgo"
	
"github.com/donomii/nucular/rect"
 //"image"
 "github.com/donomii/nucular"
 "image/color"
 nstyle "github.com/donomii/nucular/style"
 "github.com/donomii/glim"
// "github.com/disintegration/imaging"
// "image/draw"

)

func handleMessage(m svarmrgo.Message) []svarmrgo.Message {
	out := []svarmrgo.Message{}
	switch m.Selector {
	case "reveal-yourself":
		out = append(out, svarmrgo.Message{Selector: "announce", Arg: "nucular-gui"})
	case "start-gui":
		OpenWin()
		out = append(out, svarmrgo.Message{Selector: "announce", Arg: "Gui Started"})
	}
	return out
}

func OpenWin() {
	wnd := nucular.NewMasterWindow(0, "MyWindow", updatefn)
	var theme nstyle.Theme = nstyle.DarkTheme
	const scaling = 1.8
	wnd.SetStyle(nstyle.FromTheme(theme, scaling))
	wnd.Main()

}


func main() {
	conn := svarmrgo.CliConnect()
	svarmrgo.HandleInputLoop(conn, handleMessage)
		
}



func updatefn(w *nucular.Window) {
	w.Row(30).Dynamic(1)
	w.Label("Dynamic fixed column layout with generated position and size (LayoutRowDynamic):", "LC")
	w.Row(30).Dynamic(1)
	w.LabelColored("Hello", "LC", color.RGBA{255,255,255,255})
	img, _ := glim.DrawStringRGBA(9.6, color.RGBA{255,255,255,255}, "Hello again", "f1.ttf")
	newH := img.Bounds().Max.Y
	w.Row(newH).Dynamic(1)
	w.Image(img)
	img2, W, H := glim.GFormatToImage(img, nil, 0, 0)
	img2 = glim.MakeTransparent(img2, color.RGBA{0,0,0,0})
	img3 := glim.Rotate270(W, H, img2)
	img4 := glim.ImageToGFormatRGBA(H, W, img3)
	img5 := img4
	w.Image(img5)
	w.Cmds().DrawImage(rect.Rect{50, 100, 200, 200}, img5)
	}

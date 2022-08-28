package main

import (
	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
	"github.com/fogleman/gg"
)

func main() {
	const S = 1024 / 2
	dc := gg.NewContext(S, S)
	dc.SetRGB(1, 1, 1)
	dc.Clear()
	dc.SetRGB(0, 0, 0)
	//dc.SetFontFace(font.Face)

	//-- /home/zaki/.local/share/fonts/Rakkas-Regular.ttf
	//-- /home/zaki/.local/share/fonts/CourierPrime-Regular.ttf
	//-- /home/zaki/.local/share/fonts/Cairo-VariableFont_wght.ttf
	//-- /usr/share/fonts/truetype/noto/NotoNaskhArabicUI-Regular.ttf

	if err := dc.LoadFontFace("/home/zaki/.local/share/fonts/Cairo-VariableFont_wght.ttf", 20); err != nil {
		panic(err)
	}

	qrCode, _ := qr.Encode("Hello World", qr.M, qr.Auto)
	qrCode, _ = barcode.Scale(qrCode, 100, 100)

	//text := garabic.Shape("السلام عليكم 2 كيف الحال؟")
	text := "Hi i'm zaki i'm a developer! Hi i'm zaki i'm a developer!Hi i'm zaki i'm a developer!Hi i'm zaki i'm a developer!"

	dc.DrawStringWrapped(text, S/2, S/2, 0.5, 0.5, 200.0, 2.0, gg.AlignCenter)
	dc.DrawImageAnchored(qrCode, S/2, S-(qrCode.Bounds().Dy()/2)-10, 0.5, 0.5)
	dc.DrawRoundedRectangle(5, 5, S-10, S-10, 10)
	dc.SetLineWidth(2)
	dc.Stroke()
	dc.SavePNG("out.png")
}

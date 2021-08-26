package main

import (
	"FHITabule/calendarparser"
	"github.com/webview/webview"
)

func main() {
	calendarparser.Initialization()

	debug := true
	w := webview.New(debug)
	defer w.Destroy()
	w.SetTitle("Minimal webview example")
	w.SetSize(800, 600, webview.HintNone)
	w.Navigate("https://imhd.sk/ba/online-zastavkova-tabula?theme=white&zoom=67&st=66")
	w.Run()
}

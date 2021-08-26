package main

import (
	"FHITabule/calendarparser"
	"github.com/asticode/go-astikit"
	"github.com/asticode/go-astilectron"
	"log"
	"os"
	"time"
)

func main() {
	calendarparser.Initialization()

	// Make HTTP GET request
	//response, err := http.Get("https://imhd.sk/ba/online-zastavkova-tabula?theme=white&zoom=67&st=66")
	var a, _ = astilectron.New(log.New(os.Stderr, "", 0), astilectron.Options{
		AppName:            "<your app name>",
		AppIconDefaultPath: "", // If path is relative, it must be relative to the data directory
		AppIconDarwinPath:  "", // Same here
		VersionAstilectron: "",
		VersionElectron:    "",
	})
	defer a.Close()

	// Start astilectron
	a.Start()
	var w, _ = a.NewWindow("https://imhd.sk/ba/online-zastavkova-tabula?theme=white&zoom=67&st=66", &astilectron.WindowOptions{
		Center: astikit.BoolPtr(true),
		Height: astikit.IntPtr(600),
		Width:  astikit.IntPtr(600),
	})
	w.Create()

	time.Sleep(10 * time.Second)
	w.ExecuteJavaScript("window.location.href = 'https://google.com';")

	// Blocking pattern
	a.Wait()

}

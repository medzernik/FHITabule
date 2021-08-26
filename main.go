package main

import (
	"FHITabule/calendarparser"
	"fmt"
	"github.com/asticode/go-astikit"
	"github.com/asticode/go-astilectron"
	"log"
	"os"
	"time"
)

func main() {
	calendarparser.Initialization()

	var runElectron = false

	if runElectron == true {
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
			Center:      astikit.BoolPtr(true),
			Height:      astikit.IntPtr(600),
			Width:       astikit.IntPtr(600),
			Minimizable: astikit.BoolPtr(false),
			Fullscreen:  astikit.BoolPtr(true),
		})
		w.Create()

		//TODO: google API cesta na hopu MHD + auto
		var link_imhd string = "https://imhd.sk/ba/online-zastavkova-tabula?theme=white&zoom=67&st=66"

		//{{$imhd = $link_imhd}}

		for {
			err := w.ExecuteJavaScript("window.location.href = " + link_imhd + ";")
			if err != nil {
				fmt.Println("ERROR:", err)
			}
			time.Sleep(10 * time.Second)
		}

		// Blocking pattern
		a.Wait()
	}

}

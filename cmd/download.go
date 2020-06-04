package cmd

import (
	"fmt"

	"github.com/Tylerholland12/scdl-api/soundcloud"
)

func downloadSong(args []string) {
	url := args[0]

	if Artwork == true {
		// album art image
		fmt.Println("lmao")
	}
	// song name
	soundcloud.ExtractSong(url)

}

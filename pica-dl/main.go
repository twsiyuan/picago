// Copyright 2014 Tamás Gulácsi. All rights reserved.
// Use of this source code is governed by an Apache 2.0
// license that can be found in the LICENSE file.

// pica-dl implements a simple Picasa Web downloader.
package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/tgulacsi/picago"
)

// See https://developers.google.com/accounts/docs/OAuth2InstalledApp .
func main() {
	flagID := flag.String("id", os.Getenv("CLIENT_ID"), "application client ID")
	flagSecret := flag.String("secret", os.Getenv("CLIENT_SECRET"), "application client secret")

	flag.Parse()
	userid := flag.Arg(0)

	pica := picago.NewClient(*flagID, *flagSecret)
	albums, err := picago.GetAlbums(pica, userid)
	if err != nil {
		log.Fatalf("error listing albums: %v", albums)
	}

	for _, album := range albums {
		log.Printf("downloading album %s.", album)
		photos, err := picago.GetPhotos(client, userid, album.ID)
		if err != nil {
			log.Printf("error listing photos of %s: %v", album.ID, err)
			continue
		}
		log.Printf("album %s contains %d photos.", album.ID, len(photos))
		for _, photo := range photos {
			log.Printf("Photo: %s", photo)
		}
	}
}
// main.go
package main

import (
	"github.com/jessevdk/go-flags"
	"os"
	"log"
	"strings"
	"bufio"
	"encoding/base64"
	"github.com/beplus/makesplash/app/helper/image-helper"
	"fmt"
)

var version = "master"

var args struct {
	Filename      string `short:"f" long:"file" description:"filename to make assets"`
	OnlyLandscape bool `short:"l" long:"only-landscape" description:"only landscape splash will be generated"`
	OnlyPortrait  bool `short:"p" long:"only-portrait" description:"only portrait splash will be generated"`
	Version       bool `short:"v" long:"version" description:"Show version."`
}

func main() {
	_, err := flags.ParseArgs(&args, os.Args)
	if err == nil {
		if args.Version {
			fmt.Printf(" Version: %v \n License: MIT \n Copyright (c) 2017 BePlus s.r.o. (https://be.plus/makeicon) \n", version)
		} else {

			if args.Filename == "" {
				log.Fatal("No file specified. Provide the path to the file using -f or --file.")
			}

			orientation := ""
			if args.OnlyLandscape {
				orientation = "landscape"
			}
			if args.OnlyPortrait {
				orientation = "portrait"
			}
			save(args.Filename, orientation)
		}
	}
}

func save(file string, orientation string) {

	// todo windows \
	s := strings.Split(file, "/")
	filename := s[len(s)-1]
	// prepare json body

	filenameArray := strings.Split(filename, ".")
	name, extension := filenameArray[0], filenameArray[1]

	myImage, err := image_helper.NewMyImageFromBase64(getImageBase64(file), name, extension)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Processing splash... It could few seconds...")

	err = myImage.Upload(orientation)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println("Splash saved to folder 'AppSplash'.")
	}
}

func getImageBase64(filename string) string {
	imgFile, err := os.Open(filename)
	if err != nil {
		log.Fatal("Open file fail")
	}
	defer imgFile.Close()

	// create a new buffer base on file size
	fInfo, _ := imgFile.Stat()
	var size int64 = fInfo.Size()
	buf := make([]byte, size)

	// read file content into buffer
	fReader := bufio.NewReader(imgFile)
	fReader.Read(buf)

	// convert the buffer bytes to base64 string - use buf.Bytes() for new image
	return base64.StdEncoding.EncodeToString(buf)
}

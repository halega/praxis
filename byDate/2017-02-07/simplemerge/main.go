// Copyright (c) 2017 Stanislav Kalashnikov. All rights reserved.
//
// SimpleMerge scans current directory for *.xml files and merge
// them into a single file, excluding duplicate entries.
package main

import (
	"encoding/xml"
	"flag"
	"log"
	"os"
	"path/filepath"
)

var outFilename = flag.String("o", "out.xml", "output filename")

type fileElem struct {
	MD5  string `xml:"md5"`
	Name string `xml:"name"`
}

func main() {
	inFilenames, err := filepath.Glob("*.xml")
	if err != nil {
		log.Fatal(err)
	}
	out, err := os.Create(*outFilename)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()
	encoder := xml.NewEncoder(out)
	for _, filename := range inFilenames {
		if filename == *outFilename {
			continue
		}
		f, err := os.Open(filename)
		if err != nil {
			log.Fatal(err)
		}
		decoder := xml.NewDecoder(f)
		var elem fileElem
		err = decoder.Decode(&elem)
		if err != nil {
			encoder.Encode(&elem)
		}
	}
}

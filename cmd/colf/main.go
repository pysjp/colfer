package main

import (
	"flag"
	"log"

	colfer "github.com/pascaldekloe/colfer/go"
)

func main() {
	log.SetFlags(0)

	flag.Parse()
	switch len(flag.Args()) {
	default:
		log.Fatal("Too many arguments")
	case 0:
		log.Fatal("Please specify the destination platform as an argument")
	case 1:
		if p := flag.Arg(0); p != "go" {
			log.Fatalf(`Unsupported destination platform: %s
For now, "go" is the only option`, p)
		}
	}

	objects, err := colfer.ReadDefs()
	if err != nil {
		log.Fatal(err)
	}
	if len(objects) == 0 {
		log.Fatal(`Colfer definitons not found (file extension ".colf")`)
	}

	for _, o := range objects {
		if err := o.Generate(); err != nil {
			log.Fatal(err)
		}
	}
}
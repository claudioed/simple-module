package main

import (
	"flag"
	"github.com/dustinkirkland/golang-petname"
	"github.com/fatih/color"
	"math/rand"
	"time"
)

var (
	words = flag.Int("words", 2, "The number of words in the pet name")
	separator = flag.String("separator", "-", "The separator between words in the pet name")
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}


func main() {
	flag.Parse()
	rand.Seed(time.Now().UnixNano())
	color.Cyan(petname.Generate(*words, *separator))
}

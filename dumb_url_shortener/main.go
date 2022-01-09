package main

import (
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	a := App{}
	a.Initialize()
	a.Run()
}

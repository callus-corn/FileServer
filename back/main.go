package main

import (
	"os"
)

func main() {
	os.Mkdir("home/yamazaki",0777)
	os.Create("home/yamazaki/test.txt")
	for {
		
	}
}

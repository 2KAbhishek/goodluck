package main

import (
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"runtime"
)

func randInt(min, max int) int {
	return min + rand.Intn(max-min)
}

func main() {
	var files []string
	var fortunePath string

	switch osType := runtime.GOOS; osType {
	case "linux":
		fortunePath = "/usr/share/fortune"
	case "darwin":
		fortunePath = "/usr/local/Cellar/fortune/9708/share/games/fortunes"
	default:
		fortunePath = os.Getenv("FORTUNE_PATH")
	}

	err := filepath.Walk(fortunePath, func(path string, f os.FileInfo, err error) error {
		if err != nil {
			log.Fatal(err)
		}
		if filepath.Ext(path) == ".dat" {
			return nil
		}
		files = append(files, path)
		return nil
	})
	if err != nil {
		panic(err)
	}
}

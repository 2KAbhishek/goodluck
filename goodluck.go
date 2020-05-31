package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"
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

	rand.Seed(time.Now().UnixNano())
	i := randInt(1, len(files))
	randomFile := files[i]

	file, err := os.Open(randomFile)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	b, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}

	quotes := string(b)

	quotesSlice := strings.Split(quotes, "%")
	j := randInt(1, len(quotesSlice))

	fmt.Print(quotesSlice[j])

}

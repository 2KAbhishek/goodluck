package main

import "math/rand"

func randInt(min, max int) int {
	return min + rand.Intn(max-min)
}

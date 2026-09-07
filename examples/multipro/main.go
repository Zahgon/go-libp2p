package main

import (
	"log"
	"math/rand"
)

func main() {
	rnd := rand.New(rand.NewSource(666))

	port1 := rnd.Intn(100) + 10000
	port2 := port1 + 1

	done := make(chan bool, 1)

	h1 := makeRandomNode(port1, done)
	h2 := makeRandomNode(port2, done)

	log.Printf("This is a conversation between %s and %s\n", h1.ID(), h2.ID())

	run(h1, h2, done)
}

func makeRandomNode(port int, done chan bool) *Node { _ = "STUB: not implemented"; return nil }

func run(h1, h2 *Node, done <-chan bool) { _ = "STUB: not implemented"; return }

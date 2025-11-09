package pkg

import (
	"math/rand"
	"time"
)

var baseLetters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

type StringGenerator interface {
	Generate(len int) string
}

type StringGeneratorImpl struct {
	rnd     *rand.Rand
	letters []rune
}

func NewRandomGenerator() *StringGeneratorImpl {

	return &StringGeneratorImpl{
		rnd:     rand.New(rand.NewSource(time.Now().UnixNano())),
		letters: baseLetters,
	}
}

func (g *StringGeneratorImpl) Generate(l int) string {
	var b = make([]rune, l)
	for i := range b {
		b[i] = g.letters[g.rnd.Intn(len(g.letters))]
	}
	return string(b)
}

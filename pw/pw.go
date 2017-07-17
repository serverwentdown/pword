package pw

//go:generate go run words_generate.go 

import (
	"log"

	"strings"
	"crypto/rand"
	"math/big"
)

type Strength int

const (
	StrengthOnline Strength = 4
	StrengthOffline Strength = 6
	StrengthCrypto Strength = 8
)

type Mode int

const (
	ModeShort Mode = iota
	ModeLong
)

type Generator struct {
	Mode
	Strength
}

func NewGenerator(m Mode, s Strength) Generator {
	return Generator{
		Mode: m,
		Strength: s,
	}
}

func (g Generator) Next() string {
	n := int(g.Strength)
	s := []string{}
	for i := 0; i < n; i++ {
		s = append(s, g.GenerateWord());
	}
	return strings.Join(s, " ")
}

func (g Generator) GenerateWord() string {
	var wordsList []string
	if g.Mode == ModeShort {
		wordsList = WordsShort
	}
	if g.Mode == ModeLong {
		wordsList = WordsLong
	}

	c := len(wordsList)
	i, err := rand.Int(rand.Reader, big.NewInt(int64(c)))
	if err != nil {
		log.Fatal(err)
	}
	return wordsList[i.Int64()]
}

package pw

//go:generate go run words_generate.go 

import (
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

func (g Generator) Next() (string, error) {
	n := int(g.Strength)
	s := []string{}
	for i := 0; i < n; i++ {
		w, err := g.GenerateWord()
		if err != nil {
			return "", err
		}
		s = append(s, w)
	}
	return strings.Join(s, " "), nil
}

func (g Generator) GenerateWord() (string, error) {
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
		return "", err
	}
	return wordsList[i.Int64()], nil
}

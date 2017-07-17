package main

import (
	"fmt"

	"github.com/serverwentdown/pword/pw"
)

func main() {
	g := pw.NewGenerator(pw.ModeShort, pw.StrengthOnline)
	fmt.Println(g.Next())
}

package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/serverwentdown/pword/pw"
	
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()

	app.Name = "pword"
	app.Usage = "generate secure passwords"
	app.Version = "0.1.0"

	app.Flags = []cli.Flag {
		cli.StringFlag{
			Name: "count, c",
			Value: "auto",
			Usage: "Generates `NUM` passwords for you to choose from",
		},
		cli.BoolFlag{
			Name: "1",
			Usage: "Equivalent to --count 1",
		},
		cli.BoolFlag{
			Name: "stronger",
			Usage: "Chooses from a list of 7,776 words instead",
		},
	}

	app.Commands = []cli.Command{
		{
			Name: "online",
			Usage: "Generates passwords for use on websites",
			Action: func(c *cli.Context) error {
				return generate(c, pw.StrengthOnline)
			},
		},
		{
			Name: "offline",
			Usage: "Generates passwords for use offline (laptops, encrypted drives)",
			Action: func(c *cli.Context) error {
				return generate(c, pw.StrengthOffline)
			},
		},
		{
			Name: "crypto",
			Usage: "Generates extremely secure passwords",
			Action: func(c *cli.Context) error {
				return generate(c, pw.StrengthCrypto)
			},
		},
		{
			Name: "recall",
			Usage: "Utility with autocomplete to help you recall passwords",
			Action: func(c *cli.Context) error {
				return recall(c)
			},
		},
	}

	app.Run(os.Args)
}

func generate(c *cli.Context, strength pw.Strength) error {
	count := int(strength)
	countString := c.GlobalString("count")
	countParsed, err := strconv.ParseUint(countString, 10, 64)
	if err == nil {
		count = int(countParsed)
	}
	if c.GlobalBool("1") {
		count = 1
	}

	mode := pw.ModeShort
	if c.GlobalBool("stronger") {
		mode = pw.ModeLong
	}
	
	g := pw.NewGenerator(mode, strength)

	for i := 0; i < count; i++ {
		pw, err := g.Next()
		if err != nil {
			return err
		}
		fmt.Println(pw)
	}

	return nil
}

func recall(c *cli.Context) error {
	return nil
}

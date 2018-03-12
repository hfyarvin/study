package main

import (
	"fmt"
	"log"
	"os"
	"sort"

	"github.com/urfave/cli"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func main() {
	app := cli.NewApp()
	app.Name = "boom"
	app.Usage = "make an explosive entrance"
	app.Action = GetArgs
	app.Flags = GetFlags()
	app.Commands = GetCommands()

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	err := app.Run(os.Args)
	check(err)
}

func Action1(c *cli.Context) error {
	fmt.Println("boom! I say!")
	return nil
}

func GetArgs(c *cli.Context) error {
	fmt.Printf("Hello %q", c.Args().Get(0))

	name := "Nefertiti"
	if c.NArg() > 0 {
		name = c.Args().Get(0)
	}

	if c.String("lang") == "spanish" {
		fmt.Println("Hola", name)
	} else {
		fmt.Println("Hello", name)
	}

	return nil
}

func GetFlags() []cli.Flag {
	var flags []cli.Flag
	flag1 := cli.StringFlag{
		Name:  "lang",
		Value: "english",
		Usage: "language for the greeting",
	}

	flag2 := cli.StringFlag{
		Name:  "config, c",
		Usage: "Load configuratuin from `FILE`",
	}

	flags = append(flags, flag1, flag2)

	return flags
}

func GetCommands() []cli.Command {
	var commands []cli.Command
	command1 := cli.Command{
		Name:    "complete",
		Aliases: []string{"c"},
		Usage:   "complete a task on the list",
		Action:  Action1,
	}

	command2 := cli.Command{
		Name:    "add",
		Aliases: []string{"a"},
		Usage:   "add a task to the list",
		Action: func(c *cli.Context) error {
			return nil
		},
	}

	commands = append(commands, command1, command2)
	return commands
}

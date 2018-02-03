package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

func main() {
	run()
}

func run1() {
	app := cli.NewApp()
	app.Name = "boom"
	app.Usage = "make ..."
	app.Action = appAction1 //函数参数
	app.Run(os.Args)
}

//设置aciton
func appAction1(c *cli.Context) error {
	fmt.Println("Boom I Say!")
	return nil
}

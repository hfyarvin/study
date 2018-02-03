package main

import (
	"fmt"
	"os"
	"sort"

	"github.com/urfave/cli"
	// "github.com/urfave/cli/altsrc"
)

var (
	language = "lala"
	Revision = "fafafaf"
)

func main() {
	run1()
}

func run1() {
	//--version -v
	cli.VersionPrinter = func(c *cli.Context) {
		fmt.Printf("version=%s revision=%s\n", c.App.Version, Revision)
	}

	app := cli.NewApp()
	app.Name = "boom"
	app.Usage = "make ..."
	app.Version = "1.0.0"
	app.Flags = getFlags() //获取flag列表
	// app.Commands = getCommands()
	app.Commands = getCommands()
	app.Action = hello //函数参数

	//排序
	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))
	app.Run(os.Args)
}

//设置aciton
func getParam(c *cli.Context) error {
	n := 0
	fmt.Printf("param No.%d: %q\n", n+1, c.Args().Get(n))
	return nil
}

func GetParamCommands() (list []cli.Command) {
	for i := 0; i < 5; i++ {
		c := cli.Command{
			Name:    fmt.Sprintf("param-%d", i+1),
			Aliases: []string{fmt.Sprintf("p%d", i+1)},
			Usage:   fmt.Sprintf("get param No.%d way", i+1),
			Action:  getParam,
		}
		list = append(list, c)
	}
	return list
}

func getFlags() []cli.Flag {
	var list []cli.Flag

	f := cli.StringFlag{
		Name:        "lang,l", //并设置了简称
		Value:       "eng",    //default value
		Usage:       "language for the greeting",
		Destination: &language, //绑定参数
		EnvVar:      "LEGACY_COMPAT_LANG,APP_LANG,LANG",
		// FilePath: "",
	}

	f2 := cli.StringFlag{
		Name:  "config,c",
		Usage: "Load configuration from `FILE`",
	}

	list = append(list, f, f2)
	return list
}

func getCommands() (list []cli.Command) {
	c1 := cli.Command{
		Name:         "hello",
		Aliases:      []string{"hey"},
		Usage:        "Say Hello",
		Action:       hello,
		Category:     "Normal", //分类
		BashComplete: helloComplete,
	}
	c2 := cli.Command{
		Name:        "param",
		Aliases:     []string{"p"},
		Usage:       "Get first param!",
		Subcommands: GetParamCommands(),
		Category:    "Params", //分类
	}

	list = append(list, c1, c2)
	return list
}

func hello(c *cli.Context) error {
	name := "na"
	fmt.Println(c.NArg())
	if c.NArg() > 0 { //表示有参数
		name = c.Args().Get(0)
	}
	//show customer error
	if language == "showerr" {
		return cli.NewExitError("it is a custom error", 233)
	}

	if language == "spanlish" {
		fmt.Println("Hola ", name)
	} else {
		fmt.Println("Hello", name)
	}

	return nil
}

func helloComplete(c *cli.Context) {
	tasks := []string{"cook", "clean", "laundry", "eat", "sleep", "code"}
	// if c.NArg() > 1 {
	// 	return
	// }
	for _, t := range tasks {
		fmt.Println(t)
	}
}

package main

import (
	"fmt"
	"github.com/urfave/cli"
	"os"
)

func test_cli()  {
	var language string
	var recursive bool
	app := cli.NewApp()
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name : "lang ,l",
			Value : "english",
			Usage : "select language",
			Destination : &language,
		},
		cli.BoolFlag{
			Name : "recursive,r",
			Usage : "recursive for the greeting",
			Destination : &recursive,
		},
	}
	app.Action = func (c *cli.Context) error {
		var cmd string
		//c.NArg 参数的数量
		if c.NArg() > 0 {
			cmd = c.Args()[0]
			fmt.Println("cmd is",cmd)
		}
		fmt.Println(recursive)
		fmt.Println(language)
		return nil
	}
	app.Run(os.Args)
}

func main() {
	app := cli.NewApp()
	app.Name = "greet"
	app.Usage = "fight the loneliness!"
	app.Action = func(c *cli.Context) error {
		fmt.Println("Hello friend")
		return nil
	}
	// 执行Action里面函数的代码
	// app.Run(os.Args)

	test_cli()
}

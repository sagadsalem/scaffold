package main

import (
	"fmt"
	"github.com/sagadsalem/scaffold/scaffold"
	"github.com/urfave/cli"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

func main() {
	app := cli.NewApp()
	app.Version = "1.0.0-rc"
	app.Usage = "Generate scaffold project layout for Go."
	app.Commands = []cli.Command{
		{
			Name:    "init",
			Aliases: []string{"i"},
			Usage:   "Generate scaffold project layout",
			Action: func(c *cli.Context) error {

				_, b, _, _ := runtime.Caller(0)
				basepath := filepath.Dir(b)

				println("files added to the current directory " + basepath)
				err := scaffold.New(false).Generate(basepath)
				//fmt.Printf("error:%+v\n", err)
				if err == nil {
					fmt.Println("Success Created. Please execute `make up` to start service.")
				}

				return err
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

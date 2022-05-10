package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/urfave/cli"
)

func main() {
	var request string
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "c",
				Value:    request,
				Required: true,
				Usage:    "Update the configuration file",
			},
		},
		Action: func(c *cli.Context) error {
			//{"configuration": [{"name": "engine","value": "mysql"},];}
			var confModel ConfigurationModel
			err := json.Unmarshal([]byte(c.String("c")), &confModel)
			if err != nil {
				log.Fatal(err.Error())
			}
			for _, v := range confModel.Configuration {
				f, err := os.OpenFile("/home/postgres/postgres.conf", os.O_APPEND|os.O_WRONLY, 0644)
				if err != nil {
					log.Fatal(err.Error())
				}
				defer f.Close()
				_, err = f.WriteString(v.Name + " = " + v.Value + "\n")
				if err != nil {
					log.Fatal(err.Error())
				}
				err = f.Close()
				if err != nil {
					log.Fatal(err.Error())
				}
			}
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

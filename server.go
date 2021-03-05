package main

import (
	"os"
	"fmt"
	"log"
	"net/http"
	"strings"
	"github.com/urfave/cli/v2"
)

func ip(w http.ResponseWriter, req *http.Request) {
	parts := strings.Split(req.RemoteAddr, ":")
	fmt.Fprintf(w, "%v", parts[0])
}

func main() {
	app := &cli.App{
		Name: "Simple ip server",
		Usage: "Simple server to discover public IP address",
		Commands: []*cli.Command{
			&cli.Command{
				Name: "server",
				Action: func(c *cli.Context) error {
					log.Println("Starting server on port:", c.String("port"))
					http.HandleFunc("/ip", ip)
					http.ListenAndServe(":" + c.String("port"), nil)
					return nil
				},
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name: "port",
						Value: "8080",
					},
				},
			},
			&cli.Command{
				Name: "greet",
				Action: func(c *cli.Context) error {
					fmt.Println("Hello friend!")
					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
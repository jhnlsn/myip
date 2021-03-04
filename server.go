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
	// (&cli.App{}).Run(os.Args)

	app := &cli.App{
		Name: "Simple ip server",
		Usage: "fight the loneliness!",
		// Action: func(c *cli.Context) error {
		// 	fmt.Println("Hello friend!")
		// 	return nil
		// },
		Commands: []*cli.Command{
			&cli.Command{
				Name: "server",
				Action: func(c *cli.Context) error {
					http.HandleFunc("/ip", ip)
					http.ListenAndServe(":8080", nil)
					return nil
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
		// Name: "server",
		// Usage: "Start IP response server",
		// Action: func(c *cli.Context) error {
		// 	http.HandleFunc("/ip", ip)
		// 	http.ListenAndServe(":8080", nil)
		// }
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
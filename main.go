package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/urfave/cli/v2"
)

func main() {
	//Flags variables
	var urlContains, refererContains, loggerIs string
	var requestHeaders cli.StringSlice
	var statusCode, first, last int
	var count, infoLogs, errorLogs, reverse bool

	app := &cli.App{
		Name:            "caddylogs",
		Usage:           "Caddy Logs Filter",
		Version:         "0.1",
		HideHelpCommand: true,
		ArgsUsage:       "<log file path>",
		Authors: []*cli.Author{
			{
				Name:  "Ali Alhajji",
				Email: "AliAlhajji1@hotmail.com",
			},
		},

		Description: "A CLI tool that filters JSON access logs emitted by Caddy Server (https://caddyserver.com)",
	}

	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:        UrlContains,
			Aliases:     []string{"u"},
			Usage:       "Search for text in URLs",
			Destination: &urlContains,
		},

		&cli.StringFlag{
			Name:        LoggerIs,
			Aliases:     []string{"g"},
			Usage:       "Search for logs in a specific logger (must be exact logger name)",
			Destination: &loggerIs,
		},

		&cli.StringFlag{
			Name:        RefererContains,
			Aliases:     []string{"r"},
			Usage:       "Search for text in Request's Referer field",
			Destination: &refererContains,
		},

		&cli.BoolFlag{
			Name:        InfoLogs,
			Aliases:     []string{"i"},
			Usage:       "Info logs only",
			Destination: &infoLogs,
		},

		&cli.BoolFlag{
			Name:        ErrorLogs,
			Aliases:     []string{"e"},
			Usage:       "Error logs only",
			Destination: &errorLogs,
		},

		&cli.StringSliceFlag{
			Name:        Header,
			Aliases:     []string{"x"},
			Usage:       "Filter logs based on request headers. format: --header <key=value>. You can specify multiple headers: --header key1=value1 --header key2=value2",
			Destination: &requestHeaders,
		},

		&cli.BoolFlag{
			Name:        Count,
			Aliases:     []string{"c"},
			Usage:       "Print only the number of logs that match the filters.",
			Value:       false,
			Destination: &count,
		},

		&cli.IntFlag{
			Name:        StatusCode,
			Aliases:     []string{"s"},
			Usage:       "Filter logs based on the status code of the response.",
			Destination: &statusCode,
		},

		&cli.IntFlag{
			Name:        First,
			Aliases:     []string{"f"},
			Usage:       "Get first n records of the current filtered logs.",
			Destination: &first,
		},

		&cli.IntFlag{
			Name:        Last,
			Aliases:     []string{"l"},
			Usage:       "Get last n records of the current filtered logs.",
			Destination: &last,
		},

		&cli.BoolFlag{
			Name:        Reverse,
			Usage:       "Reverse the results (recent first).",
			Destination: &reverse,
		},
	}

	app.Action = func(c *cli.Context) error {
		//Exit if the log file path is not provided
		if c.Args().Len() != 1 {
			cli.ShowAppHelp(c)
			return nil
		}

		//Exit if the log file cannot be opened
		lm, err := NewLogsManager(c.Args().First())
		if err != nil {
			fmt.Println("could not open logs file:", err.Error())
			return nil
		}

		if infoLogs && errorLogs {
			fmt.Println("--info and --error cannot be both set to TRUE")
			return nil
		}

		if loggerIs != "" {
			lm.LoggerIs(loggerIs)
		}

		if urlContains != "" {
			lm.UrlContains(urlContains)
		}

		if refererContains != "" {
			lm.RefererContains(refererContains)
		}

		if statusCode != 0 {
			lm.StatusCode(statusCode)
		}

		if infoLogs {
			lm.InfoLogs()
		}

		if errorLogs {
			lm.ErrorLogs()
		}

		if len(requestHeaders.Value()) != 0 {
			for _, val := range requestHeaders.Value() {
				pair := strings.Split(val, "=")
				if len(pair) != 2 {
					fmt.Printf("incorrect format of header %s. Headers must be  formatted as: key=value.", val)
					return nil
				}

				lm.RequestHeaderIs(pair[0], pair[1])
			}
		}

		if first != 0 {
			lm.First(first)
		}

		if last != 0 {
			lm.Last(last)
		}

		//If --count flag is specified, print the number of logs and exit
		if count {
			fmt.Println(lm.GetLogsCount())
			return nil
		}

		if reverse {
			lm.Reverse()
		}

		//Print the logs in the terminal
		for i, l := range lm.GetLogs() {
			fmt.Printf("[%d] %s\n", (i + 1), l.CommonLog)
		}

		return nil
	}

	app.Run(os.Args)
}

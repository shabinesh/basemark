package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
	"log/slog"
	"os"
	"os/exec"
	"strings"
	"time"
)

func runBenchmark(command string, acceptanceTime float64, variancePercent float64, iterations int) ([]float64, bool) {
	acceptableUpperLimit := acceptanceTime * (1 + variancePercent/100)

	results := make([]float64, iterations)

	for i := 0; i < iterations; i++ {
		start := time.Now()
		cmd := exec.Command("sh", "-c", command)
		err := cmd.Run()
		if err != nil {
			fmt.Printf("Error executing command: %v\n", err)
			return nil, false
		}
		executionTime := time.Since(start).Seconds() * 1000 // Convert to milliseconds
		results[i] = executionTime
	}

	passFail := true
	for _, result := range results {
		if result > acceptableUpperLimit {
			passFail = false
			break
		}
	}

	return results, passFail
}

func main() {

	app := &cli.App{
		Name:  "basemark",
		Usage: "CLI benchmarking tool for CI",
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:     "iterations",
				Aliases:  []string{"n"},
				Required: true,
				Usage:    "number of iterations to run",
			},
			&cli.Float64Flag{
				Name:     "acceptance-time",
				Aliases:  []string{"t"},
				Required: true,
				Usage:    "Acceptance time in milli seconds",
			},
			&cli.Float64Flag{
				Name:     "variance",
				Aliases:  []string{"v"},
				Required: true,
				Usage:    "acceptable variance from acceptance-time in percent",
			},
		},
		Action: func(c *cli.Context) error {
			iter := c.Int("iterations")
			acceptanceTime := c.Float64("acceptance-time")
			variancePercent := c.Float64("variance")

			if acceptanceTime <= 99 {
				slog.Warn("Acceptance time too low, did you mean in milliseconds")
			}

			cmd := strings.Join(c.Args().Slice(), " ")
			slog.Info("Running...",
				slog.Float64("acceptanceTime(milliseconds)", acceptanceTime),
				slog.Float64("variance(percent)", variancePercent),
				slog.Int("iterations", iter))

			results, pass := runBenchmark(cmd, acceptanceTime, variancePercent, iter)

			for i, result := range results {
				fmt.Printf("Run %d: %.2f ms\n", i+1, result)
			}

			if pass {
				fmt.Println("Test Result: PASS")
				os.Exit(0)
			} else {
				fmt.Println("Test Result: FAIL")
				os.Exit(1)
			}

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	if err := mainE(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func mainE() error {
	if len(os.Args) > 2 {
		return fmt.Errorf("expected only a single argument, got %d", len(os.Args))
	}

	var input string
	if len(os.Args) == 2 {
		input = os.Args[1]
	}

	if input == "now" || input == "" {
		fmt.Println(time.Now().Unix())
		return nil
	}

	in, err := strconv.Atoi(input)
	if err != nil {
		return fmt.Errorf("error converting input to int: %w", err)
	}
	t := time.Unix(int64(in), 0)
	delta := time.Since(t)
	if delta < 0 {
		delta = time.Until(t)
	}
	if delta.Hours() > 24*365*20 {
		t = time.Unix(int64(in/1000), 0)
	}
	fmt.Println(t)
	return nil
}

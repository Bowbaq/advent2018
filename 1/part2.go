package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

func main() {
	input, err := ioutil.ReadFile("./input")
	if err != nil {
		log.Fatal(errors.Wrap(err, "couldn't read input"))
	}

	var (
		currentFrequency int64
		knownFrequencies = map[int64]bool{
			0: true,
		}

		lines = strings.Split(strings.TrimSpace(string(input)), "\n")
	)

	for {
		for _, line := range lines {
			delta, err := strconv.ParseInt(line, 10, 64)
			if err != nil {
				log.Fatal(errors.Wrap(err, "couldn't parse input"))
			}

			currentFrequency += delta
			if _, seen := knownFrequencies[currentFrequency]; seen {
				fmt.Printf("Result: %d\n", currentFrequency)
				return
			}

			knownFrequencies[currentFrequency] = true
		}
	}
}

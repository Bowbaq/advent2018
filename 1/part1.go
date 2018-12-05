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

	var frequency int64
	for _, line := range strings.Split(strings.TrimSpace(string(input)), "\n") {
		delta, err := strconv.ParseInt(line, 10, 64)
		if err != nil {
			log.Fatal(errors.Wrap(err, "couldn't parse input"))
		}

		frequency += delta
	}

	fmt.Printf("Result: %d\n", frequency)
}

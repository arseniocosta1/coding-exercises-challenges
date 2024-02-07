package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"

	"gopl-solutions/ch2/2.2/conv/distconv"
)

func parseInput(input string) (distconv.Unit, error) {
	re := regexp.MustCompile(`^(\d+)(.*)$`)

	matches := re.FindStringSubmatch(input)

	if len(matches) != 3 {
		panic("Invalid input")
	}

	num, err := strconv.ParseFloat(matches[1], 64)
	if err != nil {
		panic(err)
	}
	unit := matches[2]

	switch unit {
	case "ft", "m", "mi", "pc":
		return distconv.NewUnit(num, unit)
	default:
		log.Panicf("Unknown unit: %s", unit)
	}

	return nil, nil
}

type Killometer float64

func (k Killometer) Ratio() float64 {
	return 1000
}

func (k Killometer) ToMeter() distconv.Meter {
	return distconv.Meter(k * 1000)
}

func (k Killometer) FromMeter(meter distconv.Meter) distconv.Unit {
	return Killometer(meter / 1000)
}

func (k Killometer) String() string {
	return fmt.Sprintf("%gkm", k)
}

type Millimeter float64

func (m Millimeter) Ratio() float64 {
	return 0.001
}

func (m Millimeter) ToMeter() distconv.Meter {
	return distconv.Meter(m / 1000)
}

func (m Millimeter) FromMeter(meter distconv.Meter) distconv.Unit {
	return Millimeter(meter * 1000)
}

func (m Millimeter) String() string {
	return fmt.Sprintf("%gmm", m)
}

func main() {

	// Register a new KM Unit factory
	distconv.RegisterUnitFactory("km", func(value float64) distconv.Unit {
		return Killometer(value)
	})

	distconv.RegisterUnitFactory("mm", func(value float64) distconv.Unit {
		return Millimeter(value)
	})

	args := os.Args[1:]
	var nums []string

	if len(args) == 0 {
		scanner := bufio.NewScanner(os.Stdin)

		for scanner.Scan() {
			arg := scanner.Text()
			nums = append(nums, arg)
		}
	}

	for _, arg := range args {
		nums = append(nums, arg)
	}

	for _, num := range nums {
		input, err := parseInput(num)

		if err != nil {
			panic(err)
		}

		var sb strings.Builder

		for _, unit := range distconv.AllUnitConversions(input) {
			if input == unit {
				sb.WriteString(fmt.Sprintf("\033[38;2;200;0;0;48;2;122;122;122m%-26s\033[0m", input))
				continue
			}
			sb.WriteString(fmt.Sprintf("\u001B[38;2;125;125;125m%-26s\u001B[0m", unit))
		}

		fmt.Printf("%s\n", sb.String())
	}

	convert, err := distconv.Convert(distconv.Feet(1), "m")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", convert)

}

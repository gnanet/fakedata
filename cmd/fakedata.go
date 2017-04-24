package main

import (
	"fmt"
	"math/rand"
	"os"
	"sort"
	"time"

	"github.com/lucapette/fakedata/pkg/fakedata"
	flag "github.com/spf13/pflag"
)

var usage = `
  Usage: fakedata field1 [[field2] [field3]]
    [--tick d]
    [--max n]
    [--generators]
    [--format]

    fakedata --help

  Options:
    --generators    list all available generators
    --max n         generate data up to n [default: 10]
    --tick d        generate data every d [default: 10ms]
    --help          show help information
    --format f      generate data in f format [options: csv|tab, default: " "]
`

var generatorsFlag = flag.Bool("generators", false, "list all the generators")
var tickFlag = flag.Duration("tick", 10*time.Millisecond, "generate data every d milliseconds")
var maxFlag = flag.Int("max", 10, "generate up to n rows")
var helpFlag = flag.Bool("help", false, "print usage")
var formatFlag = flag.String("format", "", "Output format")

func main() {
	if *helpFlag {
		fmt.Printf(usage)
		os.Exit(0)
	}

	if *generatorsFlag {
		generators := fakedata.List()
		sort.Strings(generators)

		for _, name := range generators {
			fmt.Printf("%s\n", name)
		}
		os.Exit(0)
	}

	if len(flag.Args()) == 0 {
		fmt.Printf(usage)
		os.Exit(0)
	}

	rand.Seed(time.Now().UnixNano())

	tick := time.Tick(*tickFlag)

	total := 0

	for range tick {
		fmt.Print(fakedata.GenerateRow(flag.Args(), *formatFlag))

		if total++; total == *maxFlag {
			return
		}
	}
}

func init() {
	flag.Parse()
}
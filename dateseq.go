package dateseq

import (
	"flag"
	"fmt"
	"io"
	"log"
	"strconv"
	"time"
)

const cmdName = "dateseq"
const defaultTimeFormat = "20060102"

var now = time.Now

// Run the dateseq
func Run(argv []string, outStream, errStream io.Writer) error {
	log.SetOutput(errStream)
	fs := flag.NewFlagSet(
		fmt.Sprintf("%s (v%s rev:%s)", cmdName, version, revision), flag.ContinueOnError)
	fs.SetOutput(errStream)
	ver := fs.Bool("version", false, "display version")
	if err := fs.Parse(argv); err != nil {
		return err
	}
	if *ver {
		return printVersion(outStream)
	}
	switch len(argv) {
	case 0:
		printUsage(outStream)
	case 1:
		d, err := strconv.Atoi(argv[0])
		if err != nil {
			return fmt.Errorf("invalid argument: integer expected, but passed %s", argv[0])
		}
		startDate := now()
		for i := 0; i < d; i++ {
			log.Println(startDate.AddDate(0, 0, i).Format(defaultTimeFormat))
		}
	case 2:
		d, err := strconv.Atoi(argv[0])
		if err != nil {
			return fmt.Errorf("invalid argument: integer expected, but passed %s", argv[0])
		}
		startDate, err := time.Parse(defaultTimeFormat, argv[1])
		if err != nil {
			return fmt.Errorf("invalid argument: date expected (yyyyMMdd), but passed %s", argv[1])
		}
		for i := 0; i < d; i++ {
			log.Println(startDate.AddDate(0, 0, i).Format(defaultTimeFormat))
		}
	}
	return nil
}

func printVersion(out io.Writer) error {
	_, err := fmt.Fprintf(out, "%s v%s (rev:%s)\n", cmdName, version, revision)
	return err
}

func printUsage(out io.Writer) error {
	_, err := fmt.Fprint(out, "usage: dateseq number_of_days [start_day]\n")
	return err
}

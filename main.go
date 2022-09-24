package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"

	"github.com/adhocore/gronx"
	"github.com/docopt/docopt-go"
)

func app(r io.Reader, startTime, endTime time.Time, showBlanks bool) {
	var cronjobL []cronjob
	scanner := bufio.NewScanner(r)

	// read in all lines from the crontab and create a 'cronjob' struct
	// for each line.
	for scanner.Scan() {
		var l = scanner.Text()
		l = strings.TrimSpace(l)
		// skip blank lines and also lines that start with a hash (#)
		if len(l) > 0 && l[0:1] != "#" {
			cj, err := NewCronjob(l)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
			} else {
				cronjobL = append(cronjobL, *cj)
			}
		}
	}

	// iterate through the time window and, for each minute, interate through
	// the cronjobs to see which ones are "due"
	gron := gronx.New()
	t := startTime
	for t.Before(endTime) || t.Equal(endTime) {
		matchCount := 0
		for _, cronjob := range cronjobL {
			match, _ := gron.IsDue(cronjob.timespec, t)
			if match {
				fmt.Println(t.Format("2006-01-02 15:04:05"), cronjob.command)
				matchCount++
			}
		}

		if showBlanks && matchCount == 0 {
			fmt.Println(t.Format("2006-01-02 15:04:05"))
		}

		t = t.Add(time.Minute * 1)
	}

}

func main() {
	usage := `A command-line tool that parses a crontab and reports all cronjobs that are due to run between the two specified timestamps.

Usage:
  crongap [-f <crontab>] [-b] <starttime> <endtime>
  crongap -h | --help
  crongap --version

Options:
  -f, --crontab <f>  The crontab file to be parsed [default: -]
  -b, --blanks       Output blank lines for times when no jobs are due
  <starttime>        The start of the time window (format YYYY-MM-DDHH:mm)
  <endtime>          The end of the time window (format YYYY-MM-DDHH:mm)
`

	opts, _ := docopt.ParseArgs(usage, nil, "https://github.com/alasdairmorris/crongap v0.0.1")

	startTimeStr, _ := opts.String("<starttime>")
	startTime, err := time.Parse("2006-01-0215:04", startTimeStr)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Unable to parse start time", startTimeStr)
		os.Exit(1)
	}

	endTimeStr, _ := opts.String("<endtime>")
	endTime, err := time.Parse("2006-01-0215:04", endTimeStr)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Unable to parse end time", endTimeStr)
		os.Exit(1)
	}

	if endTime.Before(startTime) {
		fmt.Fprintln(os.Stderr, "End time must be later than start time")
		os.Exit(1)
	}

	showBlanks, _ := opts.Bool("--blanks")

	if filepath, _ := opts.String("--crontab"); filepath == "-" {
		app(os.Stdin, startTime, endTime, showBlanks)
	} else {
		r, err := os.Open(filepath)
		if err != nil {
			log.Fatal(err)
		}
		defer r.Close()
		app(r, startTime, endTime, showBlanks)
	}

}

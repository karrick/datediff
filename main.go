package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	var fromDate, toDate time.Time
	var err error

	switch len(os.Args) {
	case 1:
		fmt.Fprintf(os.Stderr, "usage: datediff fromDate [toDate]\n")
		os.Exit(2)
	case 2:
		if fromDate, err = parseDate(os.Args[1]); err != nil {
			panic(err)
		}
		toDate = time.Now()
	case 3:
		if fromDate, err = parseDate(os.Args[1]); err != nil {
			panic(err)
		}
		if toDate, err = parseDate(os.Args[2]); err != nil {
			panic(err)
		}
	}

	days := toDate.Sub(fromDate).Hours() / 24

	var result []string
	if years := int(days / 365); years > 0 {
		result = []string{fmt.Sprintf("%d years", years)}
		days = days - float64(365*years)
	}

	if months := int(days / 30); months > 0 {
		result = append(result, fmt.Sprintf("%d months", months))
		days = days - float64(30*months)
	}

	result = append(result, fmt.Sprintf("%d days", int(days)))

	fmt.Println(strings.Join(result, ", "))
}

func parseDate(someDate string) (time.Time, error) {
	return time.Parse(time.RFC3339, someDate+"T00:00:00Z")
}

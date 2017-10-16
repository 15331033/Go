package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type StringArray []string

func (s *StringArray) Set(value string) error {
	*s = append(*s, value)
	return nil
}
func (s *StringArray) String() string {
	return fmt.Sprint([]string(*s))
}
func (s *StringArray) Size() int {
	return len(*s)
}

func (s *StringArray) Top() (string, error) {
	if s.Size() > 0 {
		var temp = s.String()
		temp = strings.TrimSuffix(temp, "[")
		temp = strings.TrimSuffix(temp, "]")
		return temp, nil
	}
	return "", fmt.Errorf("empty error")
}

var (
	flagSet   = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	startPage = flagSet.Uint("s", 0, "tart_page, must be specified and greter than 0")
	endPage   = flagSet.Uint("e", 0, "end_page, must be specified and greter than 0")
	pageLen   = flagSet.String("l", "", "page_size, exclusive to -f, must be postive,default is 72")
	pageType  = flagSet.Bool("f", false, "use \\f to divide pages, exclusive to -l")
	printDest = StringArray{}
	pLen      uint
	pDest     string
	inputFile string
)

func init() {
	flagSet.Var(&printDest, "d", "destination")
}

func process_args() string {
	flagSet.Parse(os.Args[1:])
	if *startPage == 0 {
		return "missing parameter s or s is not positive."
	}
	if *endPage == 0 {
		return "missing parameter e or s is not positive."
	}
	if *startPage > *endPage {
		return "invalid parameters e and s,s must be smaller than e."
	}
	if *pageLen != "" && *pageType {
		return "invalid parameters l and f,-l and -f shouldn't be used together."
	}
	if *pageLen != "" {
		temp, err1 := strconv.ParseUint(*pageLen, 10, 64)
		pLen = uint(temp)
		if err1 != nil {
			return "invalid parameter l,you must input positive number for l."
		}
		if pLen <= 0 {
			return "invalid parameter l,l must be a positive number."
		}
	}
	if printDest.Size() > 0 {
		temp, _ := printDest.Top()
		if temp == "" {
			return "invalid parameter d,d must not be empty."
		}
		pDest = temp
	}
	if *pageLen == "" && !(*pageType) {
		pLen = 72
	}
	if len(flagSet.Args()) != 0 {
		inputFile = flagSet.Args()[0]
	}
	return ""
}

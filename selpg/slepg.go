package main

import (
	"bufio"
	"io"
	"os"
	"os/exec"
	"fmt"
)

const INBUFSIZ = 16 * 1024

var (
	err         error
	programname string
	stdin       io.WriteCloser
)

func writeIntoPrintInByLine(reader *bufio.Reader, startPage uint, endPage uint, pLen uint) error {
	startLine := (startPage-1)*pLen + 1
	line_ctr := uint(1)
	endLine := (endPage - startPage + 1) * pLen
	for {
		if line_ctr > endLine {
			break
		}
		line, err := reader.ReadString('\n')
		if line_ctr < startLine {
			if err != nil {
				return err
			} else {
				line_ctr++
				continue
			}
		}
		if err == nil {
			stdin.Write([]byte(line))
			line_ctr++
			continue
		} else if err == io.EOF {
			stdin.Write([]byte(line))
			line_ctr++
			if line_ctr < endLine {
				return err
			}
			return nil
		} else {
			return err
		}
	}
	return nil
}
func writeIntoPrintInByF(reader *bufio.Reader, startPage uint, endPage uint) error {
	page_ctr := uint(1)
	for {
		if page_ctr > endPage {
			break
		}
		line, err := reader.ReadString('\f')
		if page_ctr < startPage {
			if err != nil {
				return err
			} else {
				page_ctr++
				continue
			}
		}
		if err == nil {
			stdin.Write([]byte(line))
			page_ctr++
			continue
		} else if err == io.EOF {
			stdin.Write([]byte(line))
			if page_ctr < endPage {
				return err
			}
			return nil
		} else {
			return err
		}
	}
	return nil
}

func main() {
	programname = os.Args[0]
	errmsg := process_args()
	if errmsg != "" {
		fmt.Fprintf(os.Stderr, "%s : ", programname)
		fmt.Fprintln(os.Stderr, errmsg)
		os.Exit(2)
	}
	var printer *exec.Cmd
	var inputStream = os.Stdin
	var printErr error
	if inputFile != "" {
		inputStream, err = os.Open(inputFile)
		defer inputStream.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s : ", programname)
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}
	myReader := bufio.NewReader(inputStream)
	stdin = os.Stdout
	if printDest.Size() > 0 && pDest != "" {
		//printer = exec.Command("lp", "-d", pDest)
		printer = exec.Command("cat", "-n")
		stdin, printErr = printer.StdinPipe()
		if printErr != nil {
			fmt.Fprintf(os.Stderr, "%s : ", programname)
			fmt.Fprintln(os.Stderr, printErr)
		}
		if *pageType {
			err = writeIntoPrintInByF(myReader, *startPage, *endPage)
		} else {
			err = writeIntoPrintInByLine(myReader, *startPage, *endPage, pLen)
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s : ", programname)
			fmt.Fprintln(os.Stderr, err)
		}
		stdin.Close()
		printer.Stdout = os.Stdout
		printer.Start()
	}
	if *pageType {
		err = writeIntoPrintInByF(myReader, *startPage, *endPage)
	} else {
		err = writeIntoPrintInByLine(myReader, *startPage, *endPage, pLen)
	}
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s : ", programname)
		fmt.Fprintln(os.Stderr, err)
	}
}

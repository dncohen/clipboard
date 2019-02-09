package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"github.com/atotto/clipboard"
)

func main() {
	timeout := flag.Duration("t", 0, "Erase clipboard after timeout.  Durations are specified like \"20s\" or \"2h45m\".  0 (default) means never erase.")
	verbose := flag.Bool("v", false, "Verbose output.")
	trim := flag.Bool("trim", false, "Trim whitespace.")
	flag.Parse()

	in, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	var out string
	if *trim {
		out = strings.TrimSpace(string(in))
	} else {
		out = string(in)
	}

	if err := clipboard.WriteAll(out); err != nil {
		panic(err)
	}
	if *verbose {
		fmt.Println(out, []byte(out))
		fmt.Printf("wrote %d bytes to clipboard\n", len(out))
	}

	if timeout != nil && *timeout > 0 {
		<-time.After(*timeout)
		text, err := clipboard.ReadAll()
		if err != nil {
			os.Exit(1)
		}
		if text == out {
			err = clipboard.WriteAll("")
		}
	}
	if err != nil {
		os.Exit(1)
	}
}

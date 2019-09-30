package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Panic(err)
	}

	cmd := exec.Command("hunspell")
	cmd.Stdin = file

	out, err := cmd.Output()
	if err != nil {
		log.Panic(err)
	}
	outs := string(out)

	if !valid(outs) {
		for _, v := range strings.Split(outs, "\n") {
			b, _ := regexp.MatchString(`^&`, v)
			if b {
				regex := *regexp.MustCompile(`^&\s([^\s]+)\s\d+\s\d+:\s(.+)$`)
				res := regex.FindAllStringSubmatch(v, -1)
				for i := range res {
					fmt.Printf("Misspelling: '%s'. Suggestions: '%s'\n", res[i][1], res[i][2])
				}
			}
		}
		fmt.Println("Use `git commit --amend` to amend misspelling")
	}
}

func valid(s string) (b bool) {
	b, _ = regexp.MatchString(`&`, s)
	return !b
}

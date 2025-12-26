package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strings"
)

type Text struct {
}

func main() {

	file, err := os.Open("cave.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	b, _ := io.ReadAll(file)

	text := string(b)

	textArr := strings.Fields(text)

	sort.StringSlice(textArr).Sort()

	uniq := []string{}
	trimmed := []string{}
	uniqTrimmed := []string{}
	buf := ""
	bufSpec := ""

	for _, v := range textArr {
		trimmedV := strings.Trim(v, "!@#$%^&*():;\"'\\,.?")

		if buf != trimmedV {
			uniq = append(uniq, trimmedV)

			buf = trimmedV

			if len(trimmedV) < len(v) {
				trimmedSpec := strings.ReplaceAll(v, trimmedV, "")
				trimmed = append(trimmed, trimmedSpec)
			}
		}
	}

	sort.StringSlice(trimmed).Sort()

	for _, v := range trimmed {
		if bufSpec != v {
			bufSpec = v
			uniqTrimmed = append(uniqTrimmed, v)
		}
	}

	words := strings.Join(uniq, "\n")

	fmt.Println(words)
	fmt.Println(len(uniq))
	fmt.Println(uniqTrimmed)
	fmt.Println(len(uniqTrimmed))
}

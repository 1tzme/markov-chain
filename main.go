package main

import (
	"fmt"
	"os"
	"strings"

	"markov-chain/markov"
)

func main() {
	wordLimit, prefixLen, startPref := markov.ParseFlags()

	stat, _ := os.Stdin.Stat()
	if (stat.Mode()&os.ModeCharDevice) != 0 && stat.Size() == 0 {
		fmt.Fprintln(os.Stderr, "Error: no input text")
		os.Exit(1)
	}
	if wordLimit < prefixLen {
		fmt.Println("-w should be equal or more than prefix length")
		os.Exit(1)
	}

	chain := markov.NewChain(prefixLen)
	chain.Build(os.Stdin)

	if chain.Size() == 0 {
		fmt.Println("Not enough length to build chain")
		os.Exit(1)
	}
	if wordLimit > chain.Size() {
		fmt.Println("Word limit exceeds imput size")
		os.Exit(1)
	}

	start := make(markov.Prefix, prefixLen)
	if startPref != "" {
		trimmed := strings.Join(strings.Fields(startPref), " ")
		words := strings.Fields(trimmed)
		if len(words) != prefixLen {
			fmt.Println("Starting prefix len mismatch")
			os.Exit(1)
		}
		copy(start, words)

		if !chain.HasPrefix(start) {
			fmt.Println("Prefix not found or can not be continued")
			os.Exit(1)
		}
	}
	if startPref == "" {
		wordLimit += 2
	}
	fmt.Println(chain.Generate(wordLimit, start))
}

package main

import (
	"flag"
	"fmt"
)

func main() {
	repoPtr := flag.String("repo", "", "Repo name")
	commitMsg := flag.String("com", "", "Commit message")
	flag.Parse() // Parses the flags

	fmt.Printf("Using command: git -C %v add . ; git -C %v commit -m %v && git -C %v push", *repoPtr, *repoPtr, *commitMsg, *repoPtr)
}

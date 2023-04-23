package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-git/go-git/v5"
)

func main() {

	start := time.Now()

	dir := os.Args[1]
	repo, err := git.PlainOpen(dir)
	if err != nil {
		fmt.Printf("Unable to open cloned repository directory %s\n", err)
		os.Exit(-1)
	}

	// Get current commit ref
	headCommitRef, err := repo.Head()
	if err != nil {
		fmt.Printf("Unable to fetch head for the cloned repository %s\n", err)
		os.Exit(-1)
	}

	// Get previous commit ref
	previousCommitHash, err := repo.ResolveRevision("HEAD~1")
	if err != nil {
		fmt.Printf("Unable to fetch head for the cloned repository %s\n", err)
		os.Exit(-1)
	}

	filesFromCurrentCommit, err := getFilesForCommit(repo, headCommitRef.Hash())
	if err != nil {
		fmt.Printf("Could not fetch files for current commit %s\n", err)
		os.Exit(-1)
	}

	filesFromPreviousCommit, err := getFilesForCommit(repo, *previousCommitHash)
	if err != nil {
		fmt.Printf("Could not fetch files for previous commit %s\n", err)
		os.Exit(-1)
	}

	// Print diff
	diffInFiles := findChangedFiles(filesFromPreviousCommit, filesFromCurrentCommit)
	fmt.Printf("DIFF %+v\n", diffInFiles)
	elapsed := time.Since(start)
	log.Printf("Execution time %s\n", elapsed)
}

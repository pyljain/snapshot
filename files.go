package main

import (
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
)

func getFilesForCommit(repo *git.Repository, commit plumbing.Hash) (map[string]string, error) {
	c, err := repo.CommitObject(commit)
	if err != nil {
		return nil, err
	}

	tree, err := c.Tree()
	if err != nil {
		return nil, err
	}

	result := make(map[string]string)

	tree.Files().ForEach(func(f *object.File) error {
		result[f.Name] = f.Hash.String()
		return nil
	})

	return result, nil
}

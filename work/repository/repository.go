package repository

import (
	"os"
	"errors"

	"path"
	"gopkg.in/src-d/go-git.v4"
)

func GetRepository(dir string) (*git.Repository, error) {
	if repo, err := git.PlainOpen(dir); err == nil {
		return repo, nil
	} else {
		parent := path.Dir(dir)

		if _, err := os.Stat(parent); os.IsExist(err) {
			return GetRepository(parent)
		} else {
			err := errors.New("Repository is not found.")
			return nil, err
		}
	}
}
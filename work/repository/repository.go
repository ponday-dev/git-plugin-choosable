package repository

import (
	"os"
	"path"
	"errors"

	"gopkg.in/AlecAivazis/survey.v1"
	"gopkg.in/src-d/go-git.v4"
    "gopkg.in/src-d/go-git.v4/plumbing"
)

func Repository(dir string) (*git.Repository, error) {
	if repo, err := git.PlainOpen(dir); err == nil {
		return repo, nil
	} else {
		parent := path.Dir(dir)

		if _, err := os.Stat(parent); os.IsNotExist(err) {
			err := errors.New("Repository is not found.")
			return nil, err
		} else {
			return Repository(parent)
		}
	}
}

func Branches(repo *git.Repository) ([]string, map[string]*plumbing.Reference, error) {
    references, _ := repo.Branches()
	branches := []string{}
	branchesMap := map[string]*plumbing.Reference{}

    references.ForEach(func(ref *plumbing.Reference) error {
		branch := string([]rune(ref.Strings()[0])[11:])

		branches = append(branches, branch)
		branchesMap[branch] = ref

		return nil
    })

	return branches, branchesMap, nil
}

func SelectBranch(branches []string) string {
	var branch string

	prompt := &survey.Select{
        Message: "Select Branch",
        Options: branches,
    }
	survey.AskOne(prompt, &branch, nil)
	
	return branch
}

package main

import (
    "os"

    "main/repository"
)

func main() {
    cwd, _ := os.Getwd()
    repo, err := repository.Repository(cwd)
    if err != nil {
        panic(err)
    }

    branches, branchesMap, _ := repository.Branches(repo)
	branch := repository.SelectBranch(branches)
	
	repo.Storer.RemoveReference(branchesMap[branch].Name())
}

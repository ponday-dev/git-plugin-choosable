package main

import (
    "os"
    "os/exec"

    "main/repository"
)

func main() {
    cwd, _ := os.Getwd()
    repo, err := repository.Repository(cwd)
    if err != nil {
        panic(err)
    }

    branches, _, _ := repository.Branches(repo)
    branch := repository.SelectBranch(branches)

    exec.Command("git", "checkout", branch).Run()
}

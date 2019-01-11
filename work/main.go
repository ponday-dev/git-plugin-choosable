package main

import (
    "fmt"
    "os"

    "gopkg.in/src-d/go-git.v4"
    "gopkg.in/src-d/go-git.v4/plumbing"
    "gopkg.in/AlecAivazis/survey.v1"

    "main/repository"
)

func main() {
    cwd, _ := os.Getwd()
    r, err := repository.GetRepository(cwd)
    if err != nil {
        panic(err)
    }
    prefix := "refs/heads/"

    references, _ := r.Branches()
    branches := []string{}
    branchesMap := map[string]plumbing.ReferenceName{}
    references.ForEach(func(ref *plumbing.Reference) error {
        strings := ref.Strings()
        branches = append(branches, string([]rune(strings[0])[len(prefix):]))
        branchesMap[strings[0]] = ref.Name()
        return nil
    })

    var branch string
    prompt := &survey.Select{
        Message: "Select Branch",
        Options: branches,
    }
    survey.AskOne(prompt, &branch, nil)

    w, _ := r.Worktree()
    options := &git.CheckoutOptions{}
    options.Branch = branchesMap[prefix + branch]
    options.Create = false
    options.Force = false
    w.Checkout(options)
}

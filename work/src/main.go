package main

import (
    "fmt"
    "os"
    "path"
    "errors"

    "gopkg.in/src-d/go-git.v4"
    "gopkg.in/src-d/go-git.v4/plumbing"
    "gopkg.in/AlecAivazis/survey.v1"
)

func main() {
    cwd, _ := os.Getwd()
    fmt.Println(cwd)
    r, err := getRepository(cwd)
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

    fmt.Println(branch)
}

func getRepository(dir string) (*git.Repository, error) {
    if r, err := git.PlainOpen(dir); err == nil{
        return r, nil
    } else {
        parent := path.Dir(dir)
        if _, err := os.Stat(parent); os.IsExist(err){
            return getRepository(parent)
        } else {
            err := errors.New("Repository is not found.")
            return nil, err
        }
    }
}

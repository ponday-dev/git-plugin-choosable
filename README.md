# git-plugin-choosable

This is a git plugin that allows you to select targets from the branch list for `git checkout` and `git branch -d`.

## Installation

```bash
brew tap ponday-dev/git-plugin-choosable
brew install git-plugin-choosable
```

## How to use

### Toggle current branch

When you want to toggle current branch with this plugin, please use `git switch` instead of `git checkout`.

```bash
git switch
```

### Delete branch

When you want to delete branch with this plugin, please use `git remove` instead of `git branch -d`

```bash
git remove
```

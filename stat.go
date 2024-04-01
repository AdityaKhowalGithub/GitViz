package main

import (
	"fmt"
	"sort"
	"time"
	"path/filepath"
    "github.com/go-git/go-git"
   "gopkg.in/src-d/go-git.v5/plumbing/object"
)
const outOfRange = 99999
const daysInLastSixMonths = 183
const weeksInLastSixMonths = 26
func stats(email string){
    commits := processRepos(email)
    printStats(commits)
    
}


func processRepos(email string) map[int]int{
    filepath := getDotFilePath()

    repos := filetoSlice(filepath)

    commits := make(map[int]int, daysInLastSixMonths)

    for i := daysInLastSixMonths; i > 0; i--{
        commits[i] = 0
    }

    for _, i := range repos{
        commits := fillCommits(email, i, commits)
    }
    return commits
}

func fillCommits(email string, path string, commits map[int]int) map[int]int{
    repo, err := git.PlainOpen(path)
    if err != nil{
        panic(err)
    }

    ref, err := repo.Head()
    if err != nil {
        panic(err)
    }

        // get the commits history starting from HEAD
    iterator, err := repo.Log(&git.LogOptions{From: ref.Hash()})
    if err != nil {
        panic(err)
    }
    // iterate the commits
    offset := calcOffset()
    err = iterator.ForEach(func(c *object.Commit) error {
        daysAgo := countDaysSinceDate(c.Author.When) + offset

        if c.Author.Email != email {
            return nil
        }

        if daysAgo != outOfRange {
            commits[daysAgo]++
        }

        return nil
    })
    if err != nil {
        panic(err)
    }

    return commits
}


func printStats(commits map[int]int){

}



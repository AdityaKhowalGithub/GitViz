package main

import (
	"path/filepath"

	"github.com/go-git/go-git/v5"
)

func stats(email string){
    commits := processRepos(email)
    printStats(commits)
    
}


func processRepos(email string) map[int]int{
    filepath := getDotFilePath()

    repos := filetoSlice(filepath)

    daysInMap := daysInLastSixMonths



}




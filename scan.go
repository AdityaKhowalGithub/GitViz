package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func scan(folder string){
    fmt.Printf("Found folders:\n\n")

    //scan all .git folders -> list of repos founds
    //make a list/slice with all folder paths with .git
    //store paths in a file


    repos := scanFolderRecursive(folder)

    filepath := getDotFilePath()

    save(filepath, repos)

    fmt.Printf("\n\nAdded!\n\n")



}


func scanFolderRecursive(folder string){

}

func scanFolder(folders []string, folder string) []string{
    folder = strings.TrimSuffix(folder, "/")

    f, err := os.Open(folder)

    if err != nil{
        log.Fatal(err)
    }

    files, err := f.Readdir(-1)
    f.Close()

    if err != nil {
        log.Fatal(err)
    }


    var path string

    for _, file := range files{

        if file.IsDir(){
            path = folder + "/" + file.Name()
            if file.Name() == ".git"{
                path = strings.TrimSuffix(path, "/.git")
                fmt.Println(path)
                folders = append(folders, path)
                continue
            }
            if file.Name() == "Vendors" || file.Name() == "node_modules"{
                continue

            }
            folders = scanFolder(folders, path)
        }

    }
    return folders

}




package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "os/user"
    "strings"
    "io"
)

func scan(folder string){
    fmt.Printf("Found folders:\n\n")

    //scan all .git folders -> list of repos founds
    //make a list/slice with all folder paths with .git
    //store paths in a file


    repos := scanFolder(make([]string, 0),folder)

    filepath := getDotFilePath()

    save(filepath, repos)

    fmt.Printf("\n\nAdded!\n\n")



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


func getDotFilePath() string{

    usr, err := user.Current()
    if err != nil{
        log.Fatal(err)
    }

    dotFile := usr.HomeDir+ "/.gogitlocalstats"

    return dotFile

}

func save(filepath string, repo []string) {
    oldRepos := filetoSlice(filepath)
    repos := joinSlices(repo, oldRepos)
    SlicetoFile(repos, filepath)
}


func filetoSlice(filepath string) []string{
    f := openFile(filepath)
    defer f.Close()
    var lines []string

    scanner := bufio.NewScanner(f)
    for scanner.Scan(){
        lines = append(lines, scanner.Text()) 
    }
    if err := scanner.Err(); err != nil {
        if err != io.EOF {
            panic(err)
        }
    }
    return lines
}
func openFile(filePath string) *os.File {
    f, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, 0755)
    if err != nil {
        if os.IsNotExist(err) {
            // file does not exist
            _, err = os.Create(filePath)
            if err != nil {
                panic(err)
            }
        } else {
            // other error
            panic(err)
        }
    }

    return f
}

func SlicetoFile(repos []string, filepath string){
    content := strings.Join(repos, "\n")
    os.WriteFile(filepath, []byte(content), 0755) 

}

func sliceContains(slice []string, value string) bool {
    for _, v := range slice {
        if v == value {
            return true
        }
    }
    return false
}

func joinSlices(neww []string, old []string) []string{
    for _, i := range neww{
        if !sliceContains(old, i){
            old = append(old, i)
        }
    } 

    return old
}

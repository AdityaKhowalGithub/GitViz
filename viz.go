package main

import ("flag")

//takes in a path
//crawls path and all subfolders
//looks for git repos
func scan(path string){

}


//generates git contribution graph
func stats(email string){

}



func main(){

    var folder string
    var email string

    flag.StringVar(&folder, "add","", "add a folder to scan for repos")
    flag.StringVar(&email, "email","your@email.com","the email to scan")

    flag.Parse()

    if folder != ""{
        scan(folder)
        return
    }

    stats(email)

}

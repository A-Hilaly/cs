package main

import "fmt"

func Hey() {
    fmt.Println("hello")
    return
}

func main() {
    Hey()
    //fmt.Println("Hello")
    /*
    wm := &WalkMan{StartPath : "/Users/ial-ah/Github/gogp"}
    err := wm.ValidateStartingPath()
    if err != nil {
        panic(err)
    }
    err = wm.Walk()
    if err != nil {
        panic(err)
    }
    list, err := wm.GetLanguageFiles("Golang")
    if err != nil {
        panic(err)
    }
    fmt.Println(list)
    fmt.Println(wm.GetOccuringLanguages())
    */
}

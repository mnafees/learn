/**
 * https://github.com/gophercises/quiz
 * Exercise #1: Quiz Game, Part 1
 *
 * 2017 Mohammed Nafees        hello <at> mnafees <dot> me
 */
package main

import (
    "fmt"
    "encoding/csv"
    "io/ioutil"
    "bytes"
    "io"
)

func main() {
    dat, err := ioutil.ReadFile("problems.csv")
    if err != nil {
        panic(err)
    }

    r := csv.NewReader(bytes.NewReader(dat))
    var score uint
    for {
        qa, err := r.Read()
        if err == io.EOF {
            break
        }
        if err != nil {
            panic(err)
        }

        fmt.Print(qa[0] + ": ")
        var input string
        fmt.Scanln(&input)
        if input == qa[1] {
            score++
        }
    }
    fmt.Println("\nYour score: ", score)
}

package main

import (
    "fmt"
)

func main() {
    dayMonths := map[string]int{
    "Jan": 31,
    "Feb": 28,
    "Mar": 31,
    "Apr": 30,
    "May": 31,
    "Jun": 30,
    "Jul": 31,
    "Aug": 31,
    "Sep": 30,
    "Oct": 31,
    "Nov": 30,
    "Dec": 31,
  }

    has28 := 0
    
    for _, days := range dayMonths {
        if days == 28 {
           has28 += 1
         }
    }

    fmt.Printf("%d months have 28 days\n", has28)

//    days, ok := dayMonths["Jan"] //this does a bool check to see if there's a match in the map, if not, says no, otherwise prints the number of days
//    if !ok {
//        fmt.Printf("Can't get days for January")
//    } else {
//      fmt.Printf("%d\n", days)
//    }
}


// Note - if you access an element of a map that doesn't exist, you will get a "zero" response (0 for an int)

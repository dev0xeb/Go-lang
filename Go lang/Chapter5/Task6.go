package main

import "fmt"

func main() {
    for day := 1; day <= 12; day++ {
        var dayString string
        switch day {
        case 1:
            dayString = "first"
        case 2:
            dayString = "second"
        case 3:
            dayString = "third"
        case 4:
            dayString = "fourth"
        case 5:
            dayString = "fifth"
        case 6:
            dayString = "sixth"
        case 7:
            dayString = "seventh"
        case 8:
            dayString = "eighth"
        case 9:
            dayString = "ninth"
        case 10:
            dayString = "tenth"
        case 11:
            dayString = "eleventh"
        case 12:
            dayString = "twelfth"
        }

        fmt.Printf("On the %s day of Christmas, my true love gave to me:\n", dayString)

        switch day {
        case 12:
            fmt.Println("twelve Drummers Drumming")
        fallthrough
        case 11:
            fmt.Println("eleven Pipers Piping")
        fallthrough
        case 10:
            fmt.Println("ten Lords a Leaping")
        fallthrough
        case 9:
            fmt.Println("nine Ladies Dancing")
        fallthrough
        case 8:
            fmt.Println("eight Maids a Milking")
        fallthrough
        case 7:
            fmt.Println("seven Swans a Swimming")
        fallthrough
        case 6:
            fmt.Println("six Geese a Laying")
        fallthrough
        case 5:
            fmt.Println("five Golden Rings")
        fallthrough
        case 4:
            fmt.Println("four Calling Birds")
        fallthrough
        case 3:
            fmt.Println("three French Hens")
        fallthrough
        case 2:
            fmt.Println("two Turtle Doves")
        fallthrough
        case 1:
            fmt.Println("a Partridge in a Pear Tree")
        }

        fmt.Println()
    }
}

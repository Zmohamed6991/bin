package main

import (
    "fmt"
    "os"
    "strconv"
)

// mapping the numbers to words using slice's index for tens and units
var tens = []string{
    "", "", "twenty", "thirty", "forty", "fifty", "sixty", "seventy", "eighty", "ninety",
}

var units = []string{
    "zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine",
    "ten", "eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen",
    "seventeen", "eighteen", "nineteen",
}

func main() {
    if len(os.Args) < 2 {
        fmt.Println("./bin/numbers-to-words <number>")
        os.Exit(1)
    }

    input := os.Args[1]
    num, err := strconv.Atoi(input)
    if err != nil {
        fmt.Println("Please enter a valid integer.")
        os.Exit(1)
    }

    result := numberToWords(num)
    fmt.Println(result)
}

func numberToWords(n int) string {
    if n == 0 {
        return "zero"
    } else if n < 1000 {
        return convertHundreds(n)
    } else if n < 100000 {
        return convertThousands(n)
    }
    return "number out of range"
}

func convertHundreds(n int) string {
    if n == 0 {
        return ""
    } else if n < 20 {
        return units[n]
    } else if n < 100 {
        if n%10 == 0 {
            return tens[n/10]
        }
        return tens[n/10] + "-" + units[n%10]
    } else if n < 1000 {
        remainder := n % 100
        if remainder == 0 {
            return units[n/100] + " hundred"
        }
        return units[n/100] + " hundred and " + convertHundreds(remainder)
    }
    return ""
}

func convertThousands(n int) string {
    if n < 1000 {
        return convertHundreds(n)
    }

    thousands := n / 1000
    remainder := n % 1000

    num := convertHundreds(thousands) + " thousand"

    if remainder == 0 {
        return num
    }

    if remainder < 100 {
        num += " and " + convertHundreds(remainder)
    } else {
        num += ", " + convertHundreds(remainder)
    }

    return num
}
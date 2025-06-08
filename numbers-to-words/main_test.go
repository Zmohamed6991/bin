package main

import (
    "fmt"
    "testing"
)


func TestNumberToWords(t *testing.T) {
    tests := []struct {
        input    int
        expected string
    }{
        {0, "zero"},
        {7, "seven"},
        {13, "thirteen"},
        {20, "twenty"},
        {25, "twenty-five"},
        {50, "fifty"},
        {100, "one hundred"},
        {105, "one hundred and five"},
        {999, "nine hundred and ninety-nine"},
        {1000, "one thousand"},
        {1010, "one thousand and ten"},
        {1100, "one thousand, one hundred"},
        {99999, "ninety-nine thousand, nine hundred and ninety-nine"},
        {100000, "number out of range"},
    }

    for _, test := range tests {
        // Use the input value as the subtest name
        t.Run(fmt.Sprintf("Input=%d", test.input), func(t *testing.T) {
            result := numberToWords(test.input)
            if result != test.expected {
                t.Errorf("expected %q but got %q", test.expected, result)
            }
        })
    }
}

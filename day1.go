package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)
func DoDay1() {
    filepath := "day1mock.txt"
    file, err := os.Open(filepath)
    if err != nil {
        fmt.Println("File reading error", err)
    }
    defer file.Close()

    var result int64

    var arr1 []int
    var arr2 []int
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        wordBreakDown := strings.Fields(scanner.Text())
        if i, err := strconv.Atoi(wordBreakDown[0]); err == nil{
            arr1 = append(arr1, i)
        }
        if i, err := strconv.Atoi(wordBreakDown[1]); err == nil{
            arr2 = append(arr2, i)
        }
    }

    sort.Slice(arr1, func(i, j int) bool {
        return arr1[i] < arr1[j]
    })

    sort.Slice(arr2, func(i, j int) bool {
        return arr2[i] < arr2[j]
    })

    for i := 0; i < len(arr1); i++ {
        result += int64(math.Abs(float64(arr1[i]) - float64(arr2[i])))
    }

    fmt.Println(result)
}

func count[T any](slice []T, f func(T) bool) int {
    count := 0
    for _, s := range slice {
        if f(s) {
            count++
        }
    }
    return count
}

func DoDay1Continue() {
    filepath := "day1.txt"
    file, err := os.Open(filepath)
    if err != nil {
        fmt.Println("File reading error", err)
    }
    defer file.Close()

    var result int

    var arr1 []string
    var arr2 []string
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        wordBreakDown := strings.Fields(scanner.Text())
        arr1 = append(arr1, wordBreakDown[0])
        arr2 = append(arr2, wordBreakDown[1])
    }

    for i := 0; i < len(arr1); i++ {
        r := count(arr2, func(s string) bool { return arr1[i] == s })
        if WhyDoINeedToAlwaysWriteThisStatementIfIJustWantToStringToIntThisSimpleOperationRequireALotOfCode, err := strconv.Atoi(arr1[i]); err == nil {
            result += WhyDoINeedToAlwaysWriteThisStatementIfIJustWantToStringToIntThisSimpleOperationRequireALotOfCode * r     
        }
    }
    
    fmt.Println(result)
}

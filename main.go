package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func contains(target string, list []string) bool {
	for _, r := range list {
		if target == r {
			return true
		}
	}
	return false
}

func nextLine() string {
	sc.Scan()
	return sc.Text()
}

func nextInt() int {
	sc.Scan()

	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}
func nextF64() float64 {
	sc.Scan()

	i, e := strconv.ParseFloat(sc.Text(), 64)
	if e != nil {
		panic(e)
	}
	return i
}

func main() {
	sc.Split(bufio.ScanWords) // スペース区切りの場合使用
	string := nextLine()
	int := nextInt()
	fmt.Print(string, int)
	sort.SliceStable(words, func(i, j int) bool { return words[i] < words[j] })

}

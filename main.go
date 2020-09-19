package main

import (
	"bufio"
	"fmt"
	"os"
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

func main() {
	sc.Split(bufio.ScanWords) // スペース区切りの場合使用
	string := nextLine()
	int := nextInt()

	fmt.Print(string, int)

}

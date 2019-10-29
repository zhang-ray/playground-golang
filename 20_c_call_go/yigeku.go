/// 编译生成 so 的命令：
/// go build -o yigeku.so -buildmode=c-shared yigeku.go
/// 参考： https://medium.com/learning-the-go-programming-language/calling-go-functions-from-other-languages-4c7d8bcc69bf

package main

import (
	"C"
	"fmt"
	"math"
	"sort"
	"sync"
)

var count int
var myMutex sync.Mutex

//export Sub
func Sub(a, b int) int { return a - b }

//export Sine
func Sine(x float64) float64 { return math.Sin(x) }

//export Sort
func Sort(vals []int) { sort.Ints(vals) }

//export Log
func Log(msg string) int {
	myMutex.Lock()
	defer myMutex.Unlock()
	fmt.Println(msg)
	count++
	return count
}

func main() {}



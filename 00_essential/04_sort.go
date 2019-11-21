package main
func main() {
	// 数组的三种定义方式
	// 编译器计算 var a 的长度
	var a = [...]int{8,4,5,2,1,9,0,7,3,6}
    // var a [10]int = [10]int{10, 25, 32, 11, 6, 36, 18, 22, 5, 7}
    // var a [10]int = [...]int{10, 25, 32, 11, 6, 36, 18, 22, 5, 7}

	print("Array(before sort): ")
	for i:=0; i < len(a); i++ {
		print(a[i]," ")
	}
	print("\n")

	for i:=0; i < len(a); i++ {
		for j:=0; j < len(a)- i - 1; j++{
			if a[j] > a[j+1] {
				// 数据交换
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}

	print("Array(after sort): ")
	for i:=0; i < len(a); i++ {
		print(a[i]," ")
	}
	print("\n")
}
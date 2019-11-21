package main

func main() {
	// a 是 array
	var a = [...]int32{1,2,3,4,5}
	// print(a, "\n") // 不能输出 array

	// b 是 slice，寄生在 a 中
	var b0 = a[:]
	var b1 = a[0:]
	var b2 = a[:3]
	var b3 = a[1:3]
	println()
	print(&b0, "\n")
	print(b0, "\n")
	print(b1, "\n")
	print(b2, "\n")
	print(b3, "\n")
	// slice 是一个二元组：（str （*uint8）， len（int64））
}
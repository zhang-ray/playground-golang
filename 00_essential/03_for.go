package main
func main() {
	b:=0 // declare with initial value
	while:=0 // Golang 没有 while 关键字
	for i :=1; i < 10; i++ {
		/* 
		 * if b is not used, it could report error
		 * 如果一个变量没有被使用，算是编译错误
		 */
		while+=i
		b += 3
		for j := 1; j < i+1; j++ {
			print("*")
		}
		print("\n")
	}
}
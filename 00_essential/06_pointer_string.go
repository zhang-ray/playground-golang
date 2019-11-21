package main

// slice 是一个二元组：（str （*uint8）， len（int64））
// string 也是这样
func mod(b *string) {
	*b = "modified"
}

func main() {
	var s string="hello"
	println(s)
	mod(&s)
	println(s)
}
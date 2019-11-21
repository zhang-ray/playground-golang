package main

/*struct的定义如下：
 * type 结构体名 struct{}，其中定义的变量不要var，但是仍然是倒序。
 */
type S struct {
	age  int
	name string
}

type S1 struct {
	S
}

func main() {
	/*
	 * 匿名变量的访问：
	 * 在S1中有一个匿名变量S，
	 * 对于S中的变量的访问可以直接写a.name如13行所示；
	 * 当然也可以写成a.S.name(其变量明就是S)；
	 * 如果S1中又定义了name，写全可以访问到S中的变量了；
	 */

	var a S1 = S1{S{10, "tom"}}
	// var a S1=S1{S{name:"tom"}}
	println(a.age)
	println(a.name)

}

// package main

// import "fmt"

// type Vertex struct {
//     X int
//     Y int
// }

// func main() {
//     v := Vertex{1, 2}
//     v.X = 4
//     fmt.Println(v.X)
// }

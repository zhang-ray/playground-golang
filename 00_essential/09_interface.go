package main

/*
 http://blog.chinaunix.net/uid-20104447-id-3778623.html
 interface是一个定义了一系列方法的集合。interface不需要明确标明，
 只要一个类型的被作为 receiver 实现了所有的interface的方法，
 编译器就能自动识别该类型实现了该interface；
 如下面的代码：
 S1实现了myface；无论是送入S1的指针，还是对象，都能实现对setName的调用，
 如果setName的receiver是*S;则20行会报错。
 这就应证了golang手册中的一句话，
 T的方法集包含receiver为T 和*T的所有方法，
 而*T的方法集只包含receiver为*T的方法；
 （更通俗的表达方法时，当参数（receiver是T）时，
 调用该方法的对象既可以时T,也可以时*T； 
 当receiver为*T时，调用时的参数只能时*T；
  所以当interface作为参数时，具体是传入T还是*T。
  需要看该对象的具体实现，如果该对象的receiver只有*T，
  则参数只能传入*T，否则可以任选T或者*T；
   当然只有当receiver为*T，且参数传入*T时，
   才可能真正修改到T对象的成员变量；
*/

type myface interface {
    setName()
}
type S1 struct {
    name string
}

func (s S1) setName() {
    s.name = "hello"
}
func test(a myface) {
    a.setName()
}

func main() {
    var a, b S1
    test(&a)
    test(b)
    println(a.name)
    println(b.name)
}
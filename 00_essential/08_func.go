package main

type iface interface {
    setName()
}
type S1 struct {
    name string
}

func (s *S1) setName() {
    s.name = "hello"
}

func main() {
    var a S1
    (*S1).setName(&a)
    // a.setName()      // 等价写法
    println(a.name)
}
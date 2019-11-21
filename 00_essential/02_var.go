package main

func main() {
	var i     int
	var j8    int8
	var j16   int16
	var j32u   uint32
	var j64u   uint64
	// var j128 int128
	var k *int
	i+=1
	i++
	j16++;j32u++;j64u++;
	k=&i
	print(i, "\t");
	print(j8, "\n");
	print(k, "\n");
	// print(i*k, "\n");
	// k = &j //error
	print(k, "\n")
}
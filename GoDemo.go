package main

import "fmt"


func iotaTest() { ///in Python:def iotaTest
	fmt.Println("  ")
	fmt.Println("  ")
	fmt.Println("  ")
	fmt.Println("###################################################################")
	fmt.Println("Function 'iotaTest' Runing")
	const(
		a = iota
		b
		c
		d = "Hello.Com Is FatDev Ch-Company"
		e = iota
		f = "Ahahahahahahaha!!?!?!??!?!"
		g = iota
		h
		i
		j
		k
		l
		m = "End 'iota' Range"
	)
	fmt.Println(a,b,c,d,e,f,g,h,i,j,k,l,m)
	fmt.Println("  ")
	fmt.Println("  ")
	fmt.Println("  ")
}
func secondFunc() {
	const high = 20
	const low = 10
	fmt.Println("  ")
	fmt.Println("  ")
	fmt.Println("  ")

	fmt.Println("###################################################################")
	fmt.Println("Function 'secondFunc' Runing")
	fmt.Println("建筑物A高度为",high)
	fmt.Println("建筑物B高度为",low)
	fmt.Println("两个建筑物高度相差",high-low)
}
func main() {
	fmt.Println("Go Basic Demo v1.0rc1")
	fmt.Println("  ")
	fmt.Println("  ")
	fmt.Println("  ")
	fmt.Println("###################################################################")
	fmt.Println("Function 'main' Runing")
	fmt.Println("Hello, 世界")
	fmt.Println("这是我的第一个程序!!")
	var identifier = "Valuerible Access!!"
	var i1,i2,i3 = "V1 Good! ","V2 Good! ","V3 Good! "
	fmt.Println(identifier)
	fmt.Println(i1+i2+i3)
	fmt.Println("i1 Variable Memory Diercty:",&i1)
	secondFunc()
	iotaTest()
	fmt.Println(" ")
	fmt.Println(" ")
	fmt.Println(" ")
	fmt.Println(" ")
	fmt.Println(" ")
	fmt.Println("###################################################################")
	fmt.Println("Go Basic Test/Learn Demo by FatManDJ in FatDev")
	fmt.Println("Opensource With GPLv3,FSOv8 and FPCOAv1 Agreement")
	fmt.Println("FatDev Studio Copyright")
}

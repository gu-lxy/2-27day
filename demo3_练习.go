package main

import "fmt"

func main() {
	a := 8
	b := 10
	c := 12

	//1、关系运算符的结果是bool：true、false
	//2、使用&&和||对整个表达式进行分割  a < b、b/2 > 3、a * 2 > b、c != 12
	//3、分别计算结果：true、true、true、false
	//4、true && true && true || false
	//5、true && true || false
	//6、true || false =》 flase
	result1 := a < b && b/2 > 3 && a * 2 > b || c != 12
	fmt.Println("运算结果是：",result1)

	result2 := a < b && b/2 > 3 && a * 2 < b || c != 12
	fmt.Println("result2运算结果是:",result2)
}

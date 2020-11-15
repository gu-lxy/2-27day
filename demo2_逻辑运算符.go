package main

import "fmt"

func main() {
	/**
	   命题（数学）：
	      1、今天天气阴天
	      2、我们在网络上进行直播学习
	   逻辑命题：
	      1、今天天气阴天，并且我们在网络上进行直播学习
	      2、今天天气阴天，我们没有在网络上进行直播学习
	      3、今天天气不是阴天，我们在网络上进行直播学习
	      4、...

	     数学里面的命题符号：
	          最基本的操作符号：
	               “与”操作： ^
	               “或”操作： V
	               “非”操作： ¬

	   逻辑运算符：与、或、非
	        与：&& （并且、同时的意思）
	           格式：操作符的两边都是布尔类型bool    a && b  , 那么a是bool类型，b也是bool类型
	           规则：a和b都是true的时候，&&的结果是true；a和b两种其中有一个false，那么&&的结果就是false
	           规律：全真则真，一假则假

	        或：|| （或者、或者a、或者b）
	          格式：操作符的两边是表达式，表达式的结果是bool， a || b， a是bool类型，b也是bool类型
	          规则：a和b，其中只要有一个为true，则结果就是true；只有a和b都是false的时候，结果才是false
	          规律：一真则真，全假则假
	        非：!
	           格式：对操作表达式或者操作数进行取反
	           true --> false
	           false --> true
	 */
	b1 := false // bool类型
	b2 := false
	b3 := true
	b4 := true

	result := b1 && b2 // false && false
	fmt.Println(result)// 结果：false

	result1 := b3 && b2 // true && false
	fmt.Println(result1) // 结果是false

	result2 := b3 && b4
	fmt.Println(result2) // 结果是true

	//一个练习
	num1 := 8
	num2 := 10
	// num1 >= num2 && num1 < num2  // 表达式
	//1、通过&& 把整个表达式分成左右两部分 左边:num1 >= num2 右边:num1 < num2
	//2、分别计算左边和右边的结果： 左边是false  右边是true
	//3、false && true  得到结果: false
	result3 := num1>=num2 && num1 < num2 // 计算结果
	fmt.Println(result3) //结果?

	//1、通过&&分成3部分
	//2、计算每一个部分的结果 true、false、false
	result4 := num1 != num2 && num1 > num2 && num1 >= num2 // 结果
	fmt.Println(result4)

	//
	a := 3
	b := 4

	// 1、通过||符号将整个表达式划分为3个部分 分别是 a < b 、a == b、a > b
	// 2、分别计算每一个表达式的值： true、false、false
	result5 := a < b || a == b || a > b // 结果
	fmt.Println(result5) // 结果: true


    //逻辑非
    result6 := !(a < b) //  !true --> false
    fmt.Println(result6) // 结果 ? false






}

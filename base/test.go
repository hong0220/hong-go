package base

import (
	"fmt"
	"unsafe"
)

// 常量
const (
	a = "abc"
	b = len(a)
	c = unsafe.Sizeof(a)
)

// 一个可以返回多个值的函数
func testFunc() (int, int, string) {
	fmt.Println("testFunc")

	a, b, c := 1, 2, "str"
	return a, b, c
}

func testFor() {
	fmt.Println("testFor")

	var a = 0
	for a := 0; a < 10; a++ {
		fmt.Println(a)
	}
	fmt.Println(a)

	// range
	nums := []int{1, 2, 3}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println(sum)
}

func testIf() {
	fmt.Println("testIf")

	var number = 1
	if number := 4; number < 100 {
		fmt.Println(number)
	}
	fmt.Println(number)
}

func testSelect() {
	fmt.Println("testSelect")

	var c1, c2, c3 chan int //定义通信通道
	var i1, i2 int
	select {
	case i1 = <-c1:
		fmt.Println("接收消息：received ", i1, " from c1\n")
	case c2 <- i2:
		fmt.Println("发送消息：sent ", i2, " to c2\n")
	case i3, ok := <-c3:
		if ok {
			fmt.Println("接收消息：received ", i3, " from c3\n")
		} else {
			fmt.Println("c3 is closed - 消息关闭\n")
		}
	default:
		fmt.Println("no communication - 未连接 \n")
	}
}

func testPointer() {
	fmt.Println("testPointer")

	var i int
	var ptr *int
	var pptr **int
	i = 10
	ptr = &i    //接收i的地址值
	pptr = &ptr //接收ptr的地址值
	fmt.Println("%d\n", i)
	fmt.Println("%d\n", *ptr)
	fmt.Println("%d\n", **pptr)
}

type dosomething interface {
	work() int
	study() int
}

type Student struct {
	name string
	age  int
}

func (stu Student) work() int {
	fmt.Println("name:%s,age:%d -- work\n", stu.name, stu.age)
	return 3
}

func (stu Student) study() {
	fmt.Println("name:%s,age:%d -- study\n", stu.name, stu.age)
}

func testStruct() {
	fmt.Println("testStruct")

	s1 := new(Student)
	s1.name = "zhang san"
	s1.age = 23
	i := s1.work // 作为变量使用
	a := i()
	fmt.Println(a)

	s2 := new(Student)
	s2.name = "Li si"
	s2.age = 23
	s2.study()
}

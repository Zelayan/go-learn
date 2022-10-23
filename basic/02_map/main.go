package main

import "fmt"

type Student struct {
	Name string
}

var list map[string]Student

func main() {

	list = make(map[string]Student)

	student := Student{"Aceld"}

	list["student"] = student
	// cannot assign to struct field list["student"].Name in mapcompilerUnaddressableFieldAssign
	// list["student"].Name = "LDB"
	// list["student"] = student,是一个值拷贝过程
	fmt.Println(list["student"])
}

func foreachMap() {
	type student struct {
		Name string
		Age  int
	}
	//定义map
    m := make(map[string]*student)

    //定义student数组
    stus := []student{
        {Name: "zhou", Age: 24},
        {Name: "li", Age: 23},
        {Name: "wang", Age: 22},
    }

    //将数组依次添加到map中
    for _, stu := range stus {
        m[stu.Name] = &stu
		// 这种方式，所有的value 都指向最后一个，因为这里stu是值拷贝
    }
	for i := range stus {
		m[stus[i].Name]= &stus[i]
	}

    //打印map
    for k,v := range m {
        fmt.Println(k ,"=>", v.Name)
    }
}
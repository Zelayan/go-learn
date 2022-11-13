package go_reflect

import (
	"fmt"
	"reflect"
)

type User struct {
	Name string
}

func IndirectDemo() {
	val1 := []int{1, 2, 3, 4}

	var val2 reflect.Value = reflect.ValueOf(&val1)
	// ptr
	fmt.Println("&val2 type:", val2.Kind())

	// using the function
	indirectI := reflect.Indirect(val2)

	val1 = append(val1, 5)
	fmt.Println("indirectI  type:", indirectI.Kind())
	fmt.Println("indirectI  value:", indirectI)
}

// typeOfDemo 获取type类型
func typeOfDemo() {
	typeI := reflect.TypeOf(1)
	fmt.Println(typeI)
	type2 := reflect.TypeOf("hello")
	fmt.Println(type2)

	of := reflect.TypeOf(&User{Name: "1`3"})
	fmt.Println(of)
	fmt.Println(of.Kind())
	fmt.Println(of.Elem())
	fmt.Println(of.Elem().Kind())
}

// valueOfDemo get value
func valueOfDemo() {
	iValue := reflect.ValueOf(1)
	sValue := reflect.ValueOf("s")
	userPtrValue := reflect.ValueOf(&User{Name: "ze"})
	fmt.Println(iValue)
	fmt.Println(sValue)
	fmt.Println(userPtrValue)

	// value to type
	iType := iValue.Type()
	sType := sValue.Type()
	fmt.Println(iType.Kind() == reflect.Int, iValue.Kind() == reflect.Int, iType.Kind() == iValue.Kind())
	fmt.Println(sType.Kind() == reflect.String, sValue.Kind() == reflect.String, sType.Kind() == sValue.Kind())

	//指针Value和非指针Value互相转换
	userValue := userPtrValue.Elem()                    //Elem() 指针Value转为非指针Value
	fmt.Println(userValue.Kind(), userPtrValue.Kind())  //struct ptr
	userPtrValue3 := userValue.Addr()                   //Addr() 非指针Value转为指针Value
	fmt.Println(userValue.Kind(), userPtrValue3.Kind()) //struct ptr
}

// NewDemo 利用反射创建struct
func NewDemo() {
	// 利用反射创建struct
	user := reflect.TypeOf(User{})
	value := reflect.New(user) //根据reflect.Type创建一个对象，得到该对象的指针，再根据指针提到reflect.Value
	value.Elem().FieldByName("Name").SetString("123123")
	u := value.Interface().(*User)
	fmt.Println(u)
}

// NewSliceDemo 利用反射创建 slice
func NewSliceDemo() {
	var slice []User
	sliceType := reflect.TypeOf(slice)
	makeSlice := reflect.MakeSlice(sliceType, 1, 3)
	makeSlice.Index(0).Set(reflect.ValueOf(User{Name: "xxx"}))

	users := makeSlice.Interface().([]User)
	fmt.Println(users[0].Name)
}

func NewMapDemo() {
	var userMap map[int]*User
	mapType := reflect.TypeOf(userMap)
	makeMap := reflect.MakeMap(mapType)

	user := &User{Name: "zee"}
	key := reflect.ValueOf(1)
	makeMap.SetMapIndex(key, reflect.ValueOf(user))
	makeMap.MapIndex(key).Elem().FieldByName("Name").SetString("haahhah")
	userMap = makeMap.Interface().(map[int]*User)
	fmt.Println(userMap[1])
}

package main

import (
	"fmt"
	"strings"
)

type person struct {
	name string
	age  int
}

func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

//함수에서 타입을 반드시 명시해주어야 한다.
func multiply(a int, b int) int {
	return a * b
}

func multipleArgument(words ...string) {
	fmt.Println(words)
}

// naked return : return할 변수를 미리 정해둔다.
func nakedReturn(name string) (lengt int, uppercase string) {
	//defer : 함수가 다 끝났을때 실행되는것
	defer fmt.Println("I'm done")

	lengt = len(name)
	uppercase = strings.ToUpper(name)

	//무언가를 return해주긴 해야한다.
	return
}

func superAdd(numbers ...int) int {
	total := 0
	//range : array에 loop를 적용가능하게 하는것(index, array[index])
	for index, number := range numbers {
		total += number
		fmt.Println(index, number)
	}

	return total
}

func canIDrink(age int) bool {
	//if문에서 조건을 정의하기전에 변수를 정의하여 조금 더 보기 쉽게 표현이 가능하다.
	if koreanAge := age + 2; koreanAge < 18 {
		return false
	}
	return true
}

func switchDrink(age int) bool {
	//switch case에서도 if와 같이 변수를 먼저 정의 가능
	switch age {
	case 10:
		return false
	case 18:
		return true
	}
	return false
}

func main() {
	//name := "~~~" 이거랑 var name string = "~~~" 이 둘이 같은거임.
	//len, upperName := lenAndUpper("dddd")

	//  _ : 리턴값을 무시
	// flen, _ := lenAndUpper("ffff")
	// fmt.Println(flen)

	multipleArgument("ddd", "aaa", "fff", "Ccc", "sss", "qqq")

	total := superAdd(1, 2, 3, 4, 5, 6)
	fmt.Println("total = ", total)

	fmt.Println(canIDrink(16))
	fmt.Println(switchDrink(16))

	//go에서 메모리에 대한 접근을 하고싶을때는 주소값으로 &를 사용, 주소값 참조는 *사용(c언어와 같다)
	a := 2
	b := &a
	fmt.Println(&a, b)

	//array정의 방법 : 변수명 := [길이]타입{내용}
	names := [3]string{"ddd", "aaa", "ccc"}
	fmt.Println(names)

	//slice : 길이가 없는 array
	//append함수 : append(slice, value) slice에 value를 추가한다.
	nmaess := []string{"aaa", "bbb", "ccc"}
	nmaess = append(nmaess, "ddd")
	fmt.Println(nmaess)

	//map : map[key자료형]value자료형{값}
	test_map := map[string]string{"name": "ddd", "age": "11"}
	fmt.Println(test_map)

	//구조체를 정의하려면 맨 위에 type 구조체명 struct{}로 정의
	//순서대로 값을 적어도 되지만 밑의 방법을 더 선호
	js := person{name: "jsheo", age: 26}
	fmt.Println(js)
}

package main

import (
	"fmt"
	"strings"
)

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

func main() {
	//name := "~~~" 이거랑 var name string = "~~~" 이 둘이 같은거임.
	//len, upperName := lenAndUpper("dddd")

	//  _ : 리턴값을 무시
	// flen, _ := lenAndUpper("ffff")
	// fmt.Println(flen)

	multipleArgument("ddd", "aaa", "fff", "Ccc", "sss", "qqq")

}

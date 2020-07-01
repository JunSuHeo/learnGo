# Learning Go on nomad coder

## package and main   
Go 언어는 package 단위로 모듈화를 진행   
또한 package name이 main일 경우, 컴파일러가 이를 인식하여 실행프로그램을 만든다.   
   
따라서 단순한 모듈을 만들고자 할때 package를 main으로 만들어서는 안된다.   
컴파일이 필요 없다면 main을 package name으로 정하지 않으면 된다.   
   
Go 컴파일러는 main을 찾아 실행시키기 때문에 실행 가능하게 만드려면 main으로 만들어야 한다.   
이때 func main을 만들어 줘서 여기서 부터 실행할 것이라고 정해준다.   
      
```
//package를 main으로 설정
package main

//main 함수 정의
func main() {

}
```
***
      
## import
Go 언어에서 import는 vscode상에서 자동으로 진행해준다.   
   
```
import (
	"fmt"
	"strings"
)
```
      
***
## func & package
Go 언어에서 함수를 정의할때 **func 키워드** 를 사용하여 함수를 정의   
다른 package에서 있는 함수를 사용하기 위해서는   
해당 함수 이름의 첫번째 글자가 대문자여야 한다.   
   
만약 함수 이름이 대문자가 아니라면,   
해당 함수는 private으로 취급되어 다른 package에서 사용이 불가능하다.   
   

***
## func
Go 언어는 변수를 정의할때 
	[변수명][자료형]
순으로 오게 되는데, 함수에서도 마찬가지이다.   
   
함수를 정의할때
	func [함수명]([argument][argument type]) [return type]{}
순으로 이루어 진다.(함수 타입을 반드시 명시해줘야 한다.)   
   
   
Go 언어는 return value를 여러개 지정할 수 있다.   
```
func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}
```   
	(int, string)
을 통해 return할 값이 int자료형과 string 자료형이 있다는 것을 알린다.   
   

##### naked return
```
// naked return : return할 변수를 미리 정해둔다.
func nakedReturn(name string) (lengt int, uppercase string) {
	lengt = len(name)
	uppercase = strings.ToUpper(name)

	//무언가를 return해주긴 해야한다.
	return
}
```
lengt라는 변수명과 uppercase라는 변수명을 return값으로 지정해놓고   
그 안에 값을 집어 넣은 다음, return을 사용하면 두개의 값이 return된다.   
   
##### defer
defer : 함수가 다 끝났을때 실행이 되는것   
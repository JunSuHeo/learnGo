# Learning Go on nomad coder

## package and main   
Go 언어는 package 단위로 모듈화를 진행   
또한 package name이 main일 경우, 컴파일러가 이를 인식하여 실행프로그램을 만든다.   
   
따라서 단순한 모듈을 만들고자 할때 package를 main으로 만들어서는 안된다.   
컴파일이 필요 없다면 main을 package name으로 정하지 않으면 된다.   
   
Go 컴파일러는 main을 찾아 실행시키기 때문에 실행 가능하게 만드려면 main으로 만들어야 한다.   
이때 func main을 만들어 줘서 여기서 부터 실행할 것이라고 정해준다.   

```
package main

func main() {

}
```


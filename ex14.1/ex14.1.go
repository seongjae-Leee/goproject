// 포인터 : 메모리 주소를 값(숫자)으로 갖는 타입.

/*
var a int
var p *int  <- p라는 변수가 int 타입의 메모리 주소를 값으로 갖는 것. 즉 int형 포인터
p = &a  <- a의 메모리 주소를 포인터 변수 p에 대입하는 것

*p = 20  <- p라는 메모리 공간에 20을 넣으라고 한건데 20은 int고 p는 *int 타입이기 때문에 에러가 날 것
*/

/*
여러 포인터 변수가 하나의 변수를 가르킬 수 있다. 즉 여러 포인트가 같은 메모리 공간을 가리킬 수 있다.
var a int
var b int
var c int 일 때,
a = 100
b = a
c = a 이면
a, b, c는 그 공간의 값이 100이다.


var p1 *int
var p2 *int
var p3 *int
p1 = &a
p2 = p1
p3 = p1
이렇게 되면 p1, p2, p3는 a라는 같은(하나의) 메모리 공간을 가르킨다.
*/

package main

import "fmt"

func main() {
	var a int = 500
	var p *int

	p = &a

	// &p는 포인터 주소를 나타냄
	fmt.Printf("p의 값: %p\n", p)
	fmt.Printf("p가 가르키는 메모리의 값: %d\n", *p) // <- 포인터 앞에 *을 붙이면 그 공간의 값을 나타냄. 즉 a값
	*p = 100                               // p가 가르키는 공간의 값을 100으로 바꿔라, 즉 a 값을 100으로 바꿔라
	fmt.Printf("a의 값 : %d\n", a)
}

/*
결과값 :
p의 값: 0xc000012078     <- a의 메모리 주소값(cf) 16진수기 때문에 a~f도 숫자를 가르키기 위해 사용됨)
p가 가르키는 메모리의 값: 500     <- p가 가르키는 공간의 값 = a이기 때문에 500
a의 값 : 100     <-
*/

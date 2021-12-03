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


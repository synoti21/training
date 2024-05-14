package main

import "fmt"

const (
	ADD      = "+"
	DELETE   = "-"
	MULTIPLE = "*"
	DIVIDE   = "/"
)

func main() {
	var a, b int
	var operator string

	fmt.Printf("첫 번째 숫자를 입력하세요: ")
	fmt.Scanf("%d", &a)
	fmt.Printf("두 번째 숫자를 입력하세요: ")
	fmt.Scanf("%d", &b)
	fmt.Printf("연산자를 입력하세요 (+, -, *, /): ")
	fmt.Scanf("%s", &operator)

	if operator != ADD && operator != DELETE && operator != MULTIPLE && operator != DIVIDE {
		fmt.Println("에러: 유효하지 않은 연산자입니다.")
		return
	}

	switch operator {
	case ADD:
		fmt.Printf("결과: %d\n", a+b)
	case DELETE:
		fmt.Printf("결과: %d\n", a-b)
	case MULTIPLE:
		fmt.Printf("결과: %d\n", a*b)
	case DIVIDE:
		if b == 0 {
			fmt.Println("에러: 0으로 나눌 수 없습니다.")
			return
		}
		fmt.Printf("결과: %f\n", float64(a)/float64(b))
	}
}

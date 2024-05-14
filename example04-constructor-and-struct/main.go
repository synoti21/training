// make student admin system

package main

import "fmt"

type student struct {
	name string
	age  int
}

func (s *student) setName(name string) {
	s.name = name
}

func (s *student) setAge(age int) {
	s.age = age
}
func (s *student) getName() string {
	return s.name
}

func (s *student) getAge() int {
	return s.age
}

func main() {
	var tn string
	var ta int

	st := student{
		name: "default",
		age:  0,
	}

	fmt.Print("학생의 이름을 입력하세요: ")
	fmt.Scanln(&tn)
	fmt.Print("학생의 나이를 입력하세요: ")
	fmt.Scanln(&ta)

	st.setName(tn)
	st.setAge(ta)

	fmt.Printf("학생 정보:\n이름: %s\n나이: %d", st.getName(), st.getAge())
}

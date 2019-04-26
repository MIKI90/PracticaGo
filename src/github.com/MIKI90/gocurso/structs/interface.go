package structs

import "fmt"

//interface para PlatziCourse y PlatziCareer
type Platzi interface {
	Subscribe(name string)
}

func defertest() {
	fmt.Println("funcion InterfaceTest a terminado")
}

func InterfaceTest() {
	defer defertest()
	// platziCourse := PlatziCourse{Name: "Go", Slug: "go", Skills: []string{"frontend", "3"}}
	// fmt.Println(platziCourse)

	// platziCourse1 := new(PlatziCourse)
	// platziCourse1.Name = "MIKI"
	// platziCourse1.Slug = "go"
	// platziCourse1.Skills = []string{"backend", "44"}
	// fmt.Println(platziCourse1)
	// platziCourse.Subscribe("MIKI")
	platziCourse := PlatziCourse{Name: "Go", Slug: "go", Skills: []string{"frontend", "3"}}
	platziCarrer := new(PlatziCareer)
	platziCarrer.Name = "GoCareer"
	platziCarrer.Slug = "go"
	callSubscribe(platziCourse, "MIKI")
	callSubscribe(platziCarrer, "MIKI")
}

func callSubscribe(p Platzi, n string) {
	p.Subscribe(n)
}

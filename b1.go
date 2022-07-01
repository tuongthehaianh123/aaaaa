package main

import (
	"fmt"
	"os"
)

type Student struct {
	StuId          string
	StuName        string
	StuAge         string
	Stuphonenumber string
	StuGrade       string
}

func NewStudent(id, name, age, phone, grade string) *Student {
	newStu := new(Student)
	newStu.StuId = id
	newStu.StuName = name
	newStu.StuAge = age
	newStu.Stuphonenumber = phone
	newStu.StuGrade = grade

	return newStu
}

var AllStu = make(map[string]Student, 100)

func Menu() {
	fmt.Println("------------------Student Information Management System------------------")
	fmt.Println("------------------1, add student----------------------")
	fmt.Println("------------------2, modify student information-------------------")
	fmt.Println("------------------3, find students----------------------")
	fmt.Println("------------------4, student information display-------------------")
	fmt.Println("------------------0, exit the system----------------------")
	fmt.Println("-------------------------------------------------")
}

func AddStu() {
	var id string
	var name string
	var age string
	var phone string
	var grade string

	fmt.Print("Please enter id: ")
	fmt.Scanf("%s\n", &id)
	fmt.Print("Please enter your name: ")
	fmt.Scanf("%s\n", &name)
	fmt.Print("Please enter your age: ")
	fmt.Scanf("%s\n", &age)
	fmt.Print("Please enter the phone: ")
	fmt.Scanf("%f\n", &phone)
	fmt.Print("Please output grade: ")
	fmt.Scanf("%s\n", &grade)

	newStu := NewStudent(id, name, age, phone, grade)
	fmt.Println(newStu)

	AllStu[newStu.StuId] = *newStu
	fmt.Println("Added successfully!")
}
func SearchStu(id string) bool {
	value, ok := AllStu[id]
	if ok {
		fmt.Println("The student was found!")
		fmt.Printf("Student ID: %s Name: %s Age: %s Phone: %s Grade: %.2f\n\n",
			value.StuId, value.StuName, value.StuAge, value.Stuphonenumber, value.StuGrade)
		return ok
	} else {
		fmt.Println("There is no such student!")
		return false
	}
}
func ModifyStu() {
	var id string
	fmt.Print("Please output the student ID of the student you are looking for:")
	fmt.Scanf("%s\n", &id)
	res := SearchStu(id)
	if res {
		var id string
		var name string
		var age string
		var phone string
		var grade string

		fmt.Print("Please enter id: ")
		fmt.Scanf("%s\n", &id)
		fmt.Print("Please enter your name: ")
		fmt.Scanf("%s\n", &name)
		fmt.Print("Please enter your age: ")
		fmt.Scanf("%s\n", &age)
		fmt.Print("Please enter your phone: ")
		fmt.Scanf("%s\n", &phone)
		fmt.Print("Please output grade: ")
		fmt.Scanf("%s\n", &grade)
		newStu := NewStudent(id, name, age, phone, grade)
		AllStu[id] = *newStu
		fmt.Println("Modified successfully!")
	}
}

func ShowStu() {
	for _, value := range AllStu {
		fmt.Printf("Student ID: %s Name: %s Age: %s Phone: %s Grade:  %.2f\n\n",
			value.StuId, value.StuName, value.StuAge, value.Stuphonenumber, value.StuGrade)
	}
}
func main() {
	for {
		Menu()
		var opt int
		var id string
		fmt.Println("Please enter your operation")
		fmt.Scanf("%d\n", &opt)
		switch opt {
		case 1:
			AddStu()
		case 2:
			ModifyStu()
		case 3:
			fmt.Print("Please output the student ID of the student you are looking for:")
			fmt.Scanf("%s\n", &id)
			SearchStu(id)
		case 4:
			ShowStu()
		case 0:
			os.Exit(0)
		default:
			fmt.Println("Input error! Please re-enter")

		}
	}

}

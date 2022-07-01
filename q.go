package main

import (
	"fmt"
	"os"
)

//How is the student information management system
//Student class
type Student struct {
	StuId    string
	StuName  string
	StuGrade string
	StuScore string
}

//Student constructor
func NewStudent(id string, name string, grade string, score string) *Student {
	newStu := new(Student)
	newStu.StuId = id
	newStu.StuName = name
	newStu.StuGrade = grade
	newStu.StuScore = score
	return newStu
}

//Define student global variables
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
	var grade string
	var score string
	fmt.Print("Please enter id: ")
	fmt.Scanf("%s\n", &id)
	fmt.Print("Please enter your name: ")
	fmt.Scanf("%s\n", &name)
	fmt.Print("Please output grade: ")
	fmt.Scanf("%s\n", &grade)
	fmt.Print("Please enter the score: ")
	fmt.Scanf("%f\n", &score)
	newStu := NewStudent(id, name, grade, score)
	fmt.Println(newStu)
	//All students join the student object
	AllStu[newStu.StuId] = *newStu
	fmt.Println("Added successfully!")
}
func SearchStu(id string) bool {
	value, ok := AllStu[id]
	if ok {
		fmt.Println("The student was found!")
		fmt.Printf("Student ID: %s Name: %s Grade: %s Score: %.2f\n\n",
			value.StuId, value.StuName, value.StuGrade, value.StuScore)
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
		var grade string
		var score string
		fmt.Print("Please enter id: ")
		fmt.Scanf("%s\n", &id)
		fmt.Print("Please enter your name: ")
		fmt.Scanf("%s\n", &name)
		fmt.Print("Please output grade: ")
		fmt.Scanf("%s\n", &grade)
		fmt.Print("Please enter the score: ")
		fmt.Scanf("%f\n", &score)
		newStu := NewStudent(id, name, grade, score)
		AllStu[id] = *newStu
		fmt.Println("Modified successfully!")
	}
}

func ShowStu() {
	for _, value := range AllStu {
		fmt.Printf("Student ID: %s Name: %s Grade: %s Score: %.2f\n\n",
			value.StuId, value.StuName, value.StuGrade, value.StuScore)
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

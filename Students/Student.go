package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

type Student struct {
	StuId          string
	StuName        string
	StuAge         string
	Stuphonenumber string
	StuGrade       string
}

func main() {
	Menu()
	for {

		var opt int
		fmt.Println("Please enter your operation")
		fmt.Scanf("%d\n", &opt)
		switch opt {
		case 1:
			AddStu()
		case 2:
			readfile()

		// case 3:
		// 	fmt.Print("Please output the student ID of the student you are looking for:")
		// 	fmt.Scanf("%s\n", &id)
		// 	SearchStu(id)
		case 4:
			//Remove()
			// ShowStu()
		case 0:
			os.Exit(0)
		default:
			fmt.Println("Input error! Please re-enter")

		}

	}

}

var AllStu = make(map[string]Student, 100)

func Menu() {
	fmt.Println("------------------Student Information Management System------------------")
	fmt.Println("------------------1, add student----------------------")
	fmt.Println("------------------2, Readfile-------------------")
	fmt.Println("------------------3, find students----------------------")
	fmt.Println("------------------4, student information display-------------------")
	fmt.Println("------------------0, exit the system----------------------")
	fmt.Println("-------------------------------------------------")
}
func readfile() {
	fmt.Printf("----------------All information student----------------" + "\n")
	stream, err := ioutil.ReadFile("data.txt")
	if err != nil {
		log.Fatal(err)

	}

	readString := string(stream)

	var list []string = strings.Split(readString, "\n") // slice string \n
	var list1 []Student
	// fmt.Println(readString)

	for i := 0; i <= len(list)-1; i++ {
		var list2 []string = strings.Split(list[i], ",")

		obj := Student{strconv.Itoa(i + 1), list2[1], list2[2], list2[3], list2[4]}
		list1 = append(list1, obj)

	}
	// fmt.Println(list1)
	fmt.Println("ID" + " " + "Name" + " " + "Age" + "  " + "Phone" + "    " + "Grade")
	for i := 0; i <= len(list1)-1; i++ {

		fmt.Println(list1[i].StuId + " " + list1[i].StuName + " " + list1[i].StuAge + " " + list1[i].Stuphonenumber + " " + list1[i].StuGrade)

	}
	fmt.Println("Please output the student ID of the student you are looking for:")
	var id string

	fmt.Scanln(&id)
	for d := 0; d <= len(list1)-1; d++ {

		if list1[d].StuId == id {
			fmt.Println("---------Here your result is:------------")
			fmt.Println(list1[d].StuId + " " + list1[d].StuName + " " + list1[d].StuAge + " " + list1[d].Stuphonenumber + " " + list1[d].StuGrade)
			break
		}
		if d == len(list1)-1 {
			fmt.Println("------------not found in the students list---------------")
		}
	}
	fmt.Println("---------Update imformation the student--------------")
	var ID string

	fmt.Scanln(&ID)

	for d := 0; d <= len(list1)-1; d++ {

		if list1[d].StuId == ID {
			var name string
			var age string
			var phone string
			var grade string

			fmt.Println("Please enter your name: ")
			fmt.Scanf("%s\n", &name)
			fmt.Println("Please enter your age: ")
			fmt.Scanf("%s\n", &age)
			fmt.Println("Please enter the phone: ")
			fmt.Scanf("%s\n", &phone)
			fmt.Println("Please enter the grade: ")
			fmt.Scanf("%s\n", &grade)
			list1[d].StuName = name
			list1[d].StuAge = age
			list1[d].Stuphonenumber = phone
			list1[d].StuGrade = grade
			obj := Student{strconv.Itoa(len(list1) + 1), name, age, phone, grade}
			list1 = append(list1, obj)

			Remove(ID, list1)
			break
		}
		if d == len(list1)-1 {
			fmt.Println("--------Not found in the students list-----------")
		}
	}

	fmt.Println("-----------Delete Student---------")
	var xoaid string

	fmt.Scanln(&xoaid)
	for d := 0; d <= len(list1)-1; d++ {

		if list1[d].StuId == xoaid {

			Remove(xoaid, list1)
			break
			fmt.Println("delete successful")
		}
		if d == len(list1)-1 {
			fmt.Println("not found in the students ")
		}

	}

	fmt.Println("----------Students Garde=>80")

	for d := 0; d <= len(list1)-1; d++ {
		k, _ := strconv.Atoi(list1[d].StuGrade)
		if k >= 80 {
			fmt.Println(list1[d].StuId + " " + list1[d].StuName + " " + list1[d].StuAge + " " + list1[d].Stuphonenumber + " " + list1[d].StuGrade)

		}

	}
}
func AddStu() {
	var id string
	var name string
	var age string
	var phone string
	var grade string
	fmt.Printf("----------------Create student----------------")
	fmt.Println("Please enter id: ")
	fmt.Scanf("%s\n", &id)
	fmt.Println("Please enter your name: ")
	fmt.Scanf("%s\n", &name)
	fmt.Println("Please enter your age: ")
	fmt.Scanf("%s\n", &age)
	fmt.Println("Please enter the phone: ")
	fmt.Scanf("%s\n", &phone)
	fmt.Println("Please enter the grade: ")
	fmt.Scanf("%s\n", &grade)
	// fmt.Println("nhap grade")
	// fmt.Scanln(&g)

	var list1 []Student
	obj := Student{id, name, age, phone, grade}
	list1 = append(list1, obj)

	fmt.Println(list1)
	fmt.Println("Added successfully!")
	f, err := os.OpenFile("data.txt", os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		log.Fatal(err)
	}
	f.WriteString("\n" + id + "," + name + "," + age + "," + phone + "," + grade)

	fmt.Println("Successful create")

}
func Remove(id string, arr []Student) {
	e := os.Remove("data.txt")
	if e != nil {
		log.Fatal(e)
	}
	emptyFile, err := os.Create("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	emptyFile.Close()
	f, err := os.OpenFile("data.txt", os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < len(arr); i++ {

		if arr[i].StuId != id {
			f.WriteString(arr[i].StuId + "," + arr[i].StuName + "," + arr[i].StuAge + "," + arr[i].Stuphonenumber + "," + arr[i].StuGrade)
			if i != len(arr)-1 {
				f.WriteString("\n")
			}
		}

	}
	f.Close()
}

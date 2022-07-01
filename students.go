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
// func NewStudent(id, name, age, phone, grade string) *Student {
// 	newStu := new(Student)
// 	newStu.StuId = id
// 	newStu.StuName = name
// 	newStu.StuAge = age
// 	newStu.Stuphonenumber = phone
// 	newStu.StuGrade = grade

// 	return newStu
// }

func main() {

	Menu()
	for {
		
		var opt int
		fmt.Println("Please enter your operation")
		fmt.Scanf("%d\n", &opt)
		switch opt {
		case 1:
		//	AddStu()
		case 2:
			readfile()
			
		// case 3:
		// 	fmt.Print("Please output the student ID of the student you are looking for:")
		// 	fmt.Scanf("%s\n", &StuId)
		// 	SearchStu(StuId)
		case 4:
			Remove()
		case 5:
		//	ModifyStu()
			
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
	fmt.Println("------------------5, update information display-------------------")
	fmt.Println("------------------0, exit the system----------------------")
	fmt.Println("-------------------------------------------------")
}
func readfile() {
	fmt.Printf("----------------read file student----------------"+"\n")
	stream, err := ioutil.ReadFile("data.txt")
	if err != nil {
		log.Fatal(err)

	}

	readString := string(stream)

	var list []string = strings.Split(readString, "\n")// slice string \n
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
	
		if list1[d].StuId==id{
			fmt.Println(list1[d].StuId + " " + list1[d].StuName + " " + list1[d].StuAge + " " + list1[d].Stuphonenumber + " " + list1[d].StuGrade)
			break
		}
		if d==len(list1)-1{
			fmt.Println("not found in the students list")
		}
	}
}


// func SearchStu(StuId string) bool {
// 	value, ok := AllStu[StuId]
// 	if ok {
// 		fmt.Println("The student was found!")
// 		fmt.Printf("Student ID: %s Name: %s Age: %s Phone: %s Grade: %s \n\n",
// 			value.StuId, value.StuName, value.StuAge, value.Stuphonenumber, value.StuGrade)
// 		return ok
// 	} else {
// 		fmt.Println("There is no such student!")
// 		return false
// 	}
// }


func Remove(){
	// Removing file from the directory
    // Using Remove() function
	e := os.Remove("data.txt")
    if e != nil {
        log.Fatal(e)
    }
}

// func ModifyStu() {
// 	var id string
// 	fmt.Print("Please output the student ID of the student you are looking for:")
// 	fmt.Scanf("%s\n", &id)
// 	res := SearchStu(id)
// 	if res {
// 		var id string
// 		var name string
// 		var age string
// 		var phone string
// 		var grade string

// 		fmt.Print("Please enter id: ")
// 		fmt.Scanf("%s\n", &id)
// 		fmt.Print("Please enter your name: ")
// 		fmt.Scanf("%s\n", &name)
// 		fmt.Print("Please enter your age: ")
// 		fmt.Scanf("%s\n", &age)
// 		fmt.Print("Please enter your phone: ")
// 		fmt.Scanf("%s\n", &phone)
// 		fmt.Print("Please output grade: ")
// 		fmt.Scanf("%s\n", &grade)
// 		newStu := NewStudent(id, name, age, phone, grade)
// 		AllStu[id] = *newStu
// 		fmt.Println("Modified successfully!")
// 	}
// }

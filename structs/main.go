package main

import "fmt"

type contactInfo struct {
	email string
	pinCode int
}

type person struct {
	fstName string
	lstName string
	//contact contactInfo
	contactInfo
}

func main() {
	// alexa := person{fstName: "Alexa", lstName: "Pearson"}
	// var alexa person
	// //assign empty string when no values are added 
	
	// alexa.fstName = "Alexa"
	// alexa.lstName = "Pearson"
	// fmt.Println(alexa)
	// fmt.Printf("%+v", alexa)

	jim := person {
		fstName: "Jim",
		lstName: "Carry",
		//contact: contactInfo { //one can avoid field name for embedde structs.
		  contactInfo: contactInfo{	
			email: "jim@gmail.com",
			pinCode: 751002,

		},//unlike other langs there is a "," after every line even the last line
	}

	// jimPointer := &jim
	// jimPointer.updateName("Jimmy")
    //jim.updateName("Jimmy")
	//fmt.Printf("%+v", jim)

	//shortcut for structs
	jim.updateName("Jimmy")
	jim.print()
}

// func (p person) updateName(newFirstName string) {
// 	//Go passes parameters by value into a function
// 	p.fstName = newFirstName
// }

//using pointers
func (ptrToPrsn *person) updateName(newFirstName string) {// Go allows to call a fucntion with receiver to 
	//a certain type with either the type or the pointer to the type.
	//*person is a type desc
	//Go passes parameters by value into a function
	(*ptrToPrsn).fstName = newFirstName //*ptrToPrsn is an operator
}

//using receiver function
func (p person) print() {
  fmt.Printf("%+v", p)
}
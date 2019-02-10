package main

import "fmt"
import "strconv"

func starter(){
	fmt.Println("Hello World, This is my first try @golang")
}

// go function experimenting on switch case and multiple return types
func get_me_a_quiz_question(option int) (int,string){
	quiz := make(map[string]string)
	quiz["cricket"] = "Who won the 2011 worldcup"
	quiz["USA"] = "Who is the president of USA?"
	quiz["war"] = "How many days did the word war 2 lasted?"

	switch option {
	case 1:
		return option,quiz["cricket"]
	case 2:
		return option,quiz["USA"]
	case 3:
		return option,quiz["war"]
	}
	return option,"No Valid option given"
}

// addition operation done by a simple go function
func addition(a , b int) int {
	fmt.Println("Adding two numbers",a,b)
	return a + b
}


func slice_me(args []string) []string{
	i :=0
	for i<10{
		args[i] = "ID:"+strconv.Itoa(i)
		i =i+1
	}
	return args
}

// Go needs a main function to kick start things
func main(){
	starter()

	addition_answer := addition(5,9)
	fmt.Println("Answer = ",addition_answer)

	fmt.Println(get_me_a_quiz_question(2))

	students_id := make([]string, 10)
	slice_me(students_id)
	fmt.Println(students_id)

}
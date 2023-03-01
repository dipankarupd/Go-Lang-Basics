package main
import "fmt"

func main()  {


	var name string = "Dipankar"
	fmt.Println(name)

	var age int32 = 20
	fmt.Println(age)


	dipankar := "Hello World"
	fmt.Println(dipankar)

	num := 100
	fmt.Printf("%v , %T", num, num)

	fmt.Println()
	
	var s1 , s2 string = "Abiskar","Yadav"
	fmt.Println(s1,s2)

	// also supports these types of the variables: 

	var (

		username = "Dipankar"
		userage = 20
		gpa = 4.85
		description = "Hello Coders, Happy coding. I love this language very much"
	)

	fmt.Println(username, userage, gpa, description)
}
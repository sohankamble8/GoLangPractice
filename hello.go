package main

import (
	"fmt"
)
 
func main(){
		fmt.Println("Hello Sohan.. Welcome to Go..!!")

		names := Names{"Suresh", "Sagar", "Sumant"}

		fmt.Println(names[:2])
		fmt.Println(names[0])

		// Conveting Names in to Byte array to write into the file
		fmt.Println(names.toString(","))
		fmt.Println()
		// fmt.Println(names.toString(" "))
		// fmt.Println()
		// fmt.Println(names.toString("\n"))

		// Save to file
		names.saveToFile("names", " ")

		// Read from file
		fmt.Println(readNamesFromFile("names").toString(","))
		// os.Remove("names")

		// Read file from place where file not exist - Error scenario
		// fmt.Println(readNamesFromFile("name").toString(","))

		// Shuffle Names
		fmt.Println(names.shuffle().toString(","))

		// names.print()
		// fmt.Println(os.Getwd())
		// fmt.Println(os.Getpid())
		// fmt.Println(os.Executable())

}

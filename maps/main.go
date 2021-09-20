package main

import "fmt"

func main() {
    // to declare a map
	//var colors map[string]string


  colors := map[string]string{
 	  "red": "#FF0000",
  	  "green": "#4hr98u1",
 	  "blue": "#9c0081",
    }

   //  colors := make(map[string]string)

   //  colors["white"] = "#ffffff"

   //to delete a map
   // delete(colors, "white")

   //fmt.Println(colors)
  PrintMap(colors)
}

func PrintMap(c map[string]string) {
    for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow",
		"http://golang.org",
		"http://amazon.com",
	}

	//making a channel for interacting between go routines
	c := make(chan string)

	// adding go keyword for parallelism

	for _, link := range links {
        go checkLink(link, c)
	}
     
	//for looping exactly as the no.of links
	// for i := 0; i < len(links); i++ {
    //     fmt.Println(<-c)
	// }

  //for infinite looping
 //   for l := range c{
 // 	  go checkLink(l, c)
  //     }

  //To provide sleep duration in between declare a funcn literal

  for l := range c {
	  go func(link string) { //the link here refers to l in the bottom args so that we never access the same vars in diift. go rountines.
		  time.Sleep(5 * time.Second)
		  checkLink(link, c)
	    }(l)
    }
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, " might be down")
		c <- link
		return
	}

	fmt.Println(link, " is up")
	c <- link
}

package main

import (
	//"flag"
	"fmt"
	"github.com/martinvelez/super_snitch/moss"
)

var (
	userId int
	language = "c"
	version = "0.0.1"
	showVersion bool
)


/*
func init() {
	flag.IntVar(&userId, "User ID", 0, "Your MOSS User ID")
	flag.StringVar(&language, "Programming Language", "c", "Indicate a programming language")
  flag.BoolVar(&showVersion, "version", false, "View the version of this application")
	flag.Parse()
}
*/


func main(){
	fmt.Println("MOSS CLI")
	//fmt.Println("args:", flag.Args())
	moss.Send()	
}


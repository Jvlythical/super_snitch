package main

import (
	"net"
	"fmt"
	"io/ioutil"
	"os"
	"log"
	"strings"
	"strconv"
)



func main() {

	query := ""
  // connect to this socket
  conn, err := net.Dial("tcp", "moss.stanford.edu:7690")
  //conn, err := net.Dial("tcp", "127.0.0.1:3333")
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("Connected to %s",conn.RemoteAddr().String())
	defer conn.Close()
	// MOSS_USER_ID is an environment variable.
	userId := os.Getenv("MOSS_USER_ID")
	message := fmt.Sprintf("moss %d\n", userId)
	conn.Write([]byte(message)
	conn.Write([]byte("directory 0\n"))
	conn.Write([]byte("X 0\n"))
	conn.Write([]byte("maxmatches 10\n"))
	conn.Write([]byte("show 250\n"))
	conn.Write([]byte("language java\n"))
	buff := make([]byte, 1024)
	n, _ := conn.Read(buff)
	message := string(buff[:n]) 
	log.Printf("Receive: %s", message)
	message = strings.TrimSpace(message)
	if message == "no" {
		conn.Write([]byte("end\n"))
		conn.Close()
		log.Fatalln("Unrecognized language java")
		os.Exit(1)
	}

	var files = make([]string, 3)
	files[0] = "H1.java"
	files[1] = "H2.java"
	files[2] = "H3.java"
	for i := 1; i < 4; i++ {
		//uploadFile
		fname := files[i-1]
		log.Printf("Uploading %s ...", fname)
		dat, err := ioutil.ReadFile(fname)
		check(err)
		size := len(dat)
		m1 := fmt.Sprintf("file %d %s %d %s\n", i, "java", size, fname)
		log.Printf(m1)
		conn.Write([]byte(m1))
		conn.Write(dat)
		log.Printf("done.\n")

	}

	conn.Write([]byte("query 0 " + query + " \n"))
	log.Println("Query submitted.  Waiting for the server's response.")

	// listen for reply
	buff = make([]byte, 1024)
	n, err = conn.Read(buff)
	message = string(buff[:n])
	// close connection
	conn.Write([]byte("end\n"))

	log.Println("n = " + strconv.Itoa(n))
	log.Println(message)
	conn.Close()
}

package moss

/*
https://github.com/soachishti/moss.py/blob/master/mosspy/moss.py
*/

import (
	"fmt"
	"net"
	"os"
	"log"
	"io/ioutil"
	"strings"
	"strconv"
)

const server string = "moss.stanford.edu"
const port = 7690


func uploadFile(conn net.Conn, id int, fname string, lang string) {
	//uploadFile
	log.Printf("Uploading %s ...", fname)
	dat, err := ioutil.ReadFile(fname)
	check(err)
	size := len(dat)
	m1 := fmt.Sprintf("file %d %s %d %s\n", id, lang, size, fname)
	log.Printf(m1)
	conn.Write([]byte(m1))
	conn.Write(dat)
	log.Printf("done.\n")
}


func Send () {
	var userId = os.Getenv('MOSS_USER_ID')
	var optL = "c"
	var optM = 10
	var optD = 0
	var optX = 0
	var optC = ""
	var optN = 250

	optL = "java"
  // connect to this socket
  conn, err := net.Dial("tcp", server + ":" + strconv.Itoa(port))
  //conn, err := net.Dial("tcp", "127.0.0.1:3333")
	check(err)
	log.Printf("Connected to %s",conn.RemoteAddr().String())
	defer conn.Close()

	var message = fmt.Sprintf("moss %d\n", userId)
	log.Printf(message)
	conn.Write([]byte(message))

	message = fmt.Sprintf("directory %d\n", optD)
	log.Printf(message)
	conn.Write([]byte(message))

	message = fmt.Sprintf("X %d\n", optX)
	log.Printf(message)
	conn.Write([]byte(message))

	message = fmt.Sprintf("maxmatches %d\n", optM)
	log.Printf(message)
	conn.Write([]byte(message))

	message = fmt.Sprintf("show %d\n", optN)
	log.Printf(message)
	conn.Write([]byte(message))

	message = fmt.Sprintf("language %s\n", optL)
	log.Printf(message)
	conn.Write([]byte(message))

	//supported language?
	buff := make([]byte, 1024)
	n, err := conn.Read(buff)
	check(err)
	log.Printf("n = %d", n)
	message = string(buff[:n]) 
	message = strings.TrimSpace(message)
	log.Printf("Receive: %s", message)
	if message == "no" {
		conn.Write([]byte("end\n"))
		conn.Close()
		log.Fatalln(fmt.Sprintf("Unrecognized language %s",optL))
		os.Exit(1)
	}

	//TODO: upload base files

	//upload files
	var files = make([]string, 3)
	files[0] = "H1.java"
	files[1] = "H2.java"
	files[2] = "H3.java"
	for i := 1; i < 4; i++ {
		fname := files[i-1]
		uploadFile(conn, i, fname, optL)
	}

	conn.Write([]byte(fmt.Sprintf("query 0 %s\n", optC)))
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


func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}


func check(e error) {
	if e != nil {
		panic(e)
	}
}


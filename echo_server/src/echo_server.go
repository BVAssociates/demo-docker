package main

import (
        "fmt"
        "net"
        "bufio"
        "log"
        "os"
)

func echo(s net.Conn, i int, content string) {
        defer s.Close();

        fmt.Printf("%d: %v <-> %v\n", i, s.LocalAddr(), s.RemoteAddr())
        message := []byte(fmt.Sprintf("Hello, I am %s\n", content))
        s.Write(message)
        fmt.Printf("%d: closed\n", i)
}

func main() {

	if len(os.Args) != 2 {
            log.Fatal(fmt.Sprintf("USAGE: %s /path/to/file", os.Args[0]))
        }

        file, errfile := os.Open(os.Args[1])
	if errfile != nil {
            log.Fatal(errfile)
        }

	bf := bufio.NewReader(file)
	content_bytes, isPrefix, errbf := bf.ReadLine()
        if errbf != nil {
             log.Fatal(errbf)
        }
        if isPrefix {
        } 

	content := string(content_bytes)

        fmt.Printf("I will answer %s to all requests", content)

        l, e := net.Listen("tcp", ":12321")
        if e != nil {
            log.Fatal(e)
        }
        fmt.Printf("%v: listenning to 12321\n", l.Addr())
        for i := 0; e == nil; i++ {
                var s net.Conn
                s, e = l.Accept()
                go echo(s, i, content)
        }
}

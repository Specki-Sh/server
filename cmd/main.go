package main

import (
	"/pkg/server/"
	"net"
	"os"
)

func main() {
	host := "0.0.0.0"
	port := "9999"

	if err := execute(host, port); err != nil {
		os.Exit(1)
	}
}

func execute(host string, port string) (err error) {
	srv := server.NewServer(net.JoinHostPort(host, port))
	srv.Register("/", func(conn net.Conn) {
		//TODO:
	})
	srv.Resiter("/about", func(conn net.Conn) {
		//TODO:
	})

}

// 	listener, err := net.Listen("tcp", net.JoinHostPort(host, port))
// 	if err != nil {
// 		log.Print(err)
// 		return err
// 	}
// 	defer func() {
// 		if cerr := listener.Close(); cerr != nil {
// 			if err == nil {
// 				err = cerr
// 				return
// 			}
// 			log.Print(cerr)
// 		}
// 	}()

// 	for {
// 		conn, err := listener.Accept()
// 		if err != nil {
// 			log.Print(err)
// 			//Идём давать мёд следующему
// 			continue
// 		}

// 		err = handle(conn)
// 		if err != nil {
// 			log.Print(err)
// 			continue
// 		}
// 	}

// 	//return
// }

// func handle(conn net.Conn) (err error) {
// 	defer func() {
// 		if cerr := conn.Close(); cerr != nil {
// 			if err == nil {
// 				err = cerr
// 				return
// 			}
// 			log.Print(err)
// 		}
// 	}()
// 	// TODO: handle connection
// 	conn.Write([]byte("Hello!\r\n"))

// 	buf := make([]byte, 4096)
// 	for {
// 		n, err := conn.Read(buf)
// 		if err == io.EOF {
// 			log.Printf("%s", buf[:n])
// 			return nil
// 		}
// 		if err != nil {
// 			return err
// 		}
// 		log.Printf("%s", buf[:n])
// 	}
// }

package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	// Create a tcp address
	addr, err := net.Listen("tcp", ":5000")
	if err != nil {
		panic(err)
	}

	defer addr.Close()
	for {
		con, err := addr.Accept()
		fmt.Println("Conexão estabelecida")
		if err != nil {
			panic(err)
		}
		go func(c net.Conn) {
			data, _ := bufio.NewReader(con).ReadString('\n')
			fmt.Println(data)
			con.Write([]byte("Sua Msg foi recebinda com sucesso\n"))
			con.Close()
			fmt.Println("Conexão encerrada")
		}(con)
	}
}

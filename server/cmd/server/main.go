package main

import (
	"bufio"
	"errors"
	"net"
	"os"
	"strings"
	"time"

	"github.com/henri-debug/ccvp-protocol/internal/message"
	request "github.com/henri-debug/ccvp-protocol/internal/message/request"
	"github.com/henri-debug/ccvp-protocol/internal/serverutils"
	logs "github.com/henri-debug/ccvp-protocol/pkg/log"
	"github.com/henri-debug/ccvp-protocol/pkg/validator"
)

const (
	serverHost         = "localhost"
	serverPort         = "1028"
	serverType         = "tcp"
	maxConnections int = 3
)

var (
	actualConnections int
)

func main() {

	logs.InfoLog("Starting " + serverType + " Server On " + serverHost + " In Port " + serverPort)

	actualConnections = 0

	l, err := net.Listen(serverType, serverHost+":"+serverPort)
	if err != nil {
		logs.ErrorLog("Error in Listening: " + err.Error())
		os.Exit(1)
	}

	defer l.Close()

	for {
		if actualConnections < maxConnections {
			connection, err := l.Accept()

			if err != nil {
				logs.ErrorLog("Error in Connection: " + err.Error())
				return
			}

			logs.ConnectionLog(connection.RemoteAddr().String())
			generateConnectionLog(connection.RemoteAddr().String())

			go persistConnection(connection)
			actualConnections++

			if actualConnections == maxConnections {
				logs.InfoLog("Maximum connections reached, the next ones will be in the queue") // Irá ficar na fila caso o numero esteja no máximo
			}
		}
	}
}

func generateConnectionLog(client string) {
	path := "../../logs"

	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			logs.Log("Error creating log DIR")
		}
	}

	f, err := os.OpenFile(path+"/LOG["+(time.Now().Format("2006-01-02"))+"].txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)

	if err != nil {
		logs.Log("Error Creating Log File")
	}

	defer f.Close()

	_, err = f.WriteString("[CONNECTION]| " + client + " | " + time.Now().Format(time.UnixDate) + "\n")

	if err == nil {
		logs.Log("Connected " + client)
	} else {
		logs.Log("Error Writing in Log")
	}
}

func persistConnection(conn net.Conn) {
	var clientRequest request.Request

	client := conn.RemoteAddr().String()

	readBuffer, _ := bufio.NewReader(conn).ReadBytes('\n')
	request := string(readBuffer)

	clientData := strings.Split(request, " ")

	if len(clientData) != 4 { // STT 102 Wrong Request
		conn.Write([]byte("102 00 00000000000"))
		logs.RequestErrorLog("102", client)
	} else {
		clientRequest.TOR = clientData[0]
		clientRequest.Length = clientData[1]
		clientRequest.TOD = clientData[2]
		clientRequest.NOD = clientData[3]

		clientRequest.NOD = clientRequest.NOD[:len(clientRequest.NOD)-1] // Retira o /n

		typeTor := message.CheckTOR(clientRequest.TOR)

		if typeTor == 0 {
			logs.RequestLog(clientRequest.TOR+" "+clientRequest.Length+" "+clientRequest.TOD+" "+clientRequest.NOD, client)

			typeTod := message.CheckTOD(clientRequest.TOD)

			validateTod(typeTod, clientRequest, conn, client)

		} else if typeTor == 1 { // ALLOWED TYPES
			logs.RequestLog(clientRequest.TOR+" "+clientRequest.Length+" "+clientRequest.TOD+" "+clientRequest.NOD, client)
			serverutils.SendMessage(conn, clientRequest, client, "303")
		} else {
			conn.Write([]byte("102 00 00000000000"))
			logs.RequestErrorLog("102", client)
		}

	}
	conn.Close()
	logs.ClosedLog(client)
	actualConnections--
}

func validateTod(typeTod int, clientRequest request.Request, conn net.Conn, client string) {
	if typeTod == 1 { // CNPJ

		validLength := message.CheckLength(clientRequest.Length, typeTod)

		validateLength(validLength, conn, clientRequest, client)

	} else if typeTod == 0 { // CPF

		validLength := message.CheckLength(clientRequest.Length, typeTod)

		validateLength(validLength, conn, clientRequest, client)

	} else {
		conn.Write([]byte("101 00 00000000000"))
		logs.RequestErrorLog("101", client)
	}
}

func validateLength(validLength bool, conn net.Conn, clientRequest request.Request, client string) {
	if validLength {
		if validator.ValidateCPF(clientRequest.NOD) {
			serverutils.SendMessage(conn, clientRequest, client, "301")
		} else {
			serverutils.SendMessage(conn, clientRequest, client, "302")
		}
	} else {
		serverutils.SendMessage(conn, clientRequest, client, "103")
	}
}

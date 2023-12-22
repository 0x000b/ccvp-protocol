package serverutils

import (
	"net"
	"strconv"

	request "github.com/henri-debug/ccvp-protocol/internal/message/request"
	logs "github.com/henri-debug/ccvp-protocol/pkg/log"
)

func makeResponse(request request.Request, stt string) string {
	switch stt {
	case "101":
		return (stt + " " + request.Length + " " + request.NOD)
	case "102":
		return (stt + " " + request.Length + " " + request.NOD)
	case "103":
		return (stt + " " + request.Length + " " + request.NOD)
	case "301":
		return (stt + " " + request.Length + " " + request.NOD)
	case "302":
		return (stt + " " + request.Length + " " + request.NOD)
	default:
		return (stt + " " + strconv.Itoa(len("CPF;CNPJ")) + " " + "CPF;CNPJ")
	}
}

func SendMessage(conn net.Conn, clientRequest request.Request, client string, stt string) {
	conn.Write([]byte(makeResponse(clientRequest, stt)))

	if stt == "303" {
		logs.TypesResponseLog(stt, client)
	} else {
		logs.ValidateResponseLog(stt, client)
	}

}

import socket

HOST = "127.0.0.1"
PORT = 1028

def summaryResponse(response):
    match response[0]:
        case "101":
            print("101 Wrong Type Of Document")
            print("Document Length: ",response[1])
            print("Document: ",response[2])
        case "102":
            print("102 Wrong Request")
            print("Document Length: ",response[1])
            print("Document: ",response[2])
        case "103":
            print("103 Wrong Number Sequence")
            print("Document Length: ",response[1])
            print("Document: ",response[2])
        case "301":
            print("301 Accepted")
            print("Document Length: ",response[1])
            print("Document: ",response[2])
        case "302":
            print("302 Rejected")
            print("Document Length: ",response[1])
            print("Document: ",response[2])
        case "303":
            print("303 Received")
            print("Types Length: ",response[1])
            print("Allowed Types: ", response[2].split(";"))
    print("\n")


def sendRequest(HOST, PORT, message):
    with socket.socket(socket.AF_INET, socket.SOCK_STREAM) as s:
        s.connect((HOST, PORT))
        s.sendall(bytes(message,'utf-8'))
        response = s.recv(1024)
        response = response.decode('utf-8')
    summaryResponse(response.split(" "))

while 1:
    tor = input("TOR: ")

    if(int(tor) == 0):
        length = input("Nod Length: ")
        tod = input("TOD: ")
        nod = input("NOD: ")
        request = tor + " " + length + " " + tod + " " + nod + "\n" 
        print("\n")
        sendRequest(HOST,PORT,request)
        
    elif (int(tor) == 1):
        request = tor + " " + "00" + " " + "0" + " " + "00000000000" + "\n"
        print("\n")
        sendRequest(HOST,PORT,request)

    else:
        length = input("Nod Length: ")
        tod = input("TOD: ")
        nod = input("NOD: ")
        request = tor + " " + length + " " + tod + " " + nod + "\n"
        print("\n")
        sendRequest(HOST,PORT,request)
    

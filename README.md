# Basic-GoLang-HTTP-Server



How to Run:
1-Save the code to a file, for example main.go.
2-Open a terminal and run:
go run main.go

Test it using curl:
GET /ping:
curl http://localhost:8080/ping
Output:
{"message":"pong"}
POST /echo:

curl -X POST -H "Content-Type: application/json" -d '{"name":"Payam","lang":"Go"}' http://localhost:8080/echo
Output:

{"name":"Payam","lang":"Go"}

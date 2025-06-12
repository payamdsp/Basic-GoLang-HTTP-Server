# Basic-GoLang-HTTP-Server



## 📦 How to Run:
```bash
1-Save the code to a file, for example main.go.
2-Open a terminal and run:
go run main.go
```

## ✅ Test it using curl:
GET /ping:
```bash
curl http://localhost:8080/ping
```
Output:
```bash

{"message":"pong"}
```
POST /echo:
```bash
curl -X POST -H "Content-Type: application/json" -d '{"name":"Payam","lang":"Go"}' http://localhost:8080/echo
```
Output:
```bash
{"name":"Payam","lang":"Go"}
```

# Experimenting with GO ernvironment.
---
- Running go project
```shell
go run main.go
```
---
### Creating go mod file:

`go mod init github.com/username/repo`

and import it like 

`"github.com/username/repo/module"`

---

### Adding request body: 
- Create a struct for body (type)
```go
type ReqBody struct {
	Field1 string `json:"field1"`
	Field2 string `json:"field2"`
}
```
- Create a body with that struct

```go
reqBody := &ReqBody{
		Field1: "value1",
		Field2: "value2",
	}
```
- Marshal the body
```go
reqBodyBytes, err := json.Marshal(reqBody)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
```
- Create a buffer for marshalled body
```go
	req, err := http.NewRequest("POST", "url.com", bytes.NewBuffer(reqBodyBytes))

```
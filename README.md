# UDP terminal server

Run this server 
```sh
go run .
```

Then use `w1` as your `io.Writer` in your code e.g. `fmt.Fprintln(w1, "T1 : ", i)`

```go
	w1, err := net.Dial("udp", "127.0.0.1:8080")
	if err != nil {
		log.Fatal(err)
	}
```	
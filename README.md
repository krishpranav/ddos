# ddos
A simple ddos written in go

[![forthebadge](https://forthebadge.com/images/badges/made-with-go.svg)](https://forthebadge.com)

example of using:

```golang
func main() {
	workers := 100
	d, err := ddos.New("http://127.0.0.1:80", workers)
	if err != nil {
		panic(err)
	}
	d.Run()
	time.Sleep(time.Second)
	d.Stop()
	fmt.Println("DDoS attack server: http://127.0.0.1:80")
	// Output: DDoS attack server: http://127.0.0.1:80
}
```


# Installation:
- git clone https://github.com/krishpranav/ddos
- cd ddos
- go build ddos.go
- ./ddos

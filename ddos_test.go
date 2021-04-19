package ddos_test

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"testing"
	"time"

	ddos "github.com/krishpranav/ddos"
	freeport "github.com/Konstantin8105/FreePort"

)

func TestNewDDoS(t *testing.T) {
	d, err := ddos.New("http://127.0.0.1", 1)
	if err != nil {
		t.Error("Cannot create a new ddos structure. Error = ", err)
	}
	if d == nil {
		t.Error("Cannot create a new ddos structure")
	}
}

func TestDDoS(t *testing.T) {
	port, err := freeport.Get()
	if err != nil {
		t.Errorf("Cannot found free tcp port. Error = %v", err)
	}
	createServer(port, t)

	url := "http://127.0.0.1:" + strconv.Itoa(port)
	d, err := ddos.New(url, 1000)
	if err != nil {
		t.Error("Cannot create a new ddos structure")
	}
	d.Run()
	time.Sleep(time.Second)
	d.Stop()
	success, amount := d.Result()
	if success == 0 || amount == 0 {
		t.Errorf("Negative result of DDoS attack.\n"+
			"Success requests = %v.\n"+
			"Amount requests = %v", success, amount)
	}
	t.Logf("Statistic: %d %d", success, amount)
}

// create a simple go server
func createServer(port int, t *testing.T) {
	go func() {
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Hi there %s!", r.URL.Path[1:])
		})
		if err := http.ListenAndServe(":"+strconv.Itoa(port), nil); err != nil {
			t.Fatalf("Server is down. %v", err)
		}
	}()
}

func TestWorkers(t *testing.T) {
	_, err := ddos.New("127.0.0.1", 0)
	if err == nil {
		t.Error("Cannot create a new ddos structure")
	}
}


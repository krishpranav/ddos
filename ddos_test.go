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
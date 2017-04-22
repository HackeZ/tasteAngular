package controller

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"testing"

	"GutHub/srv/protocol"

	"github.com/valyala/fasthttp"
)

const (
	localhost = "http://127.0.0.1:"
)

func TestGetGutList(t *testing.T) {

	ctx := &fasthttp.RequestCtx{}

	GetGutList(ctx)

}

func TestAddGut(t *testing.T) {
	ctx := &fasthttp.RequestCtx{}

	h := &fasthttp.RequestHeader{}
	h.SetMethod("POST")

	gut := protocol.Gut{
		Title:       "test1",
		Description: "desc test1",
		Ingredients: []protocol.Ingredient{
			{
				Amount:         "1",
				AmountUnits:    "2",
				IngredientName: "3",
			},
			{
				Amount:         "4",
				AmountUnits:    "5",
				IngredientName: "6",
			},
		},
		Instructions: "inst test1",
	}

	w := &bytes.Buffer{}
	encoder := json.NewEncoder(w)
	err := encoder.Encode(gut)
	if err != nil {
		t.Fatal("Unexpected error when Encode a gut:", err)
	}
	bw := bufio.NewWriter(w)
	err = h.Write(bw)
	if err != nil {
		t.Fatalf("Unexpected error when writing request header: %s", err)
	}
	if err := bw.Flush(); err != nil {
		t.Fatalf("Unexpected error when flushing request header: %s", err)
	}

	ctx.Request.Header = *h

	t.Logf("ctx: %+v", ctx)

	AddGut(ctx)
}

func TestGetGutList2(t *testing.T) {
	port := 8088
	defer startServerOnPort(t, port, GetGutList).Close()

	// your tests here for client connecting to the given port
	resp, err := http.Get(localhost + "8088/")
	if err != nil {
		t.Error("cannot new get list of gut:", err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Log("reading response failed:", err)
	}

	t.Log(string(body))
}

func TestAddGut2(t *testing.T) {
	port := 8089
	defer startServerOnPort(t, port, AddGut).Close()
}

func startServerOnPort(t *testing.T, port int, h fasthttp.RequestHandler) io.Closer {
	ln, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		t.Fatalf("cannot start tcp server on port %d: %s", port, err)
	}
	go fasthttp.Serve(ln, h)
	return ln
}

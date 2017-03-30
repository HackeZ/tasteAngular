package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"sync"

	"GutHub/srv/protocol"
	fasthp "github.com/valyala/fasthttp"
)

var (
	// GutList save all guts
	gutList []*protocol.Gut
	// gutID identity of gut
	gutID int64
	// Mutex for make sure ID unqueness
	mutex *sync.Mutex
)

func init() {
	gutList = make([]*protocol.Gut, 0, 256)
	id = len(gutList)
	mutex = &sync.Mutex{}
}

// GetGutList return list of guts
func GetGutList(ctx *fasthp.RequestCtx) {
	var returnData protocol.ReturnData
	returnData.Code = -1
	log.Println("got GetGutList req.")

	var data bytes.Buffer
	result, err := json.Marshal(gutList)
	if err != nil {
		returnData.Message = "get gut list failed, error:" + err.Error
		log.Fatalln(returnData.Message)
		result, _ = json.Marshal(returnData)

		ctx.SetStatusCode(408)
		ctx.Write(result)
		return
	}
	returnData.Code = 0
	returnData.Message = "get list of guts success."
	returnData.Data = string(result)
	result, _ = json.Marshal(returnData)

	ctx.SetContentType("application/json")
	ctx.SetStatusCode(200)
	ctx.Write(result) // ignore handle error
	return
}

// AddGut append a new gut to list.
func AddGut(ctx *fasthp.RequestCtx) {
	var returnData protocol.ReturnData
	returnData.Code = -1
	log.Println("got AddGut req.")

	if !ctx.IsPost {
		returnData.Message = "can not handle AddGut do not in POST method."
		log.Fatalln(returnData.Message)
		result, _ := json.Marshal(returnData)

		ctx.SetStatusCode(408)
		ctx.Write(result)
		return
	}

	gutByte := ctx.FormValue("gut")
	var gut protocol.Gut
	err := json.Unmarshal(gutByte, &gut)
	if err != nil {
		returnData.Message = "unmarshal data failed, error:" + err.Error()
		log.Fatalln(returnData.Message)
		result, _ := json.Marshal(returnData)

		ctx.SetStatusCode(408)
		ctx.Write(result)
		return
	}

	mutex.Lock()
	id := gutID
	gutID++
	mutex.Unlock()

	gutList = append(gutList, func() *protocol.Gut {
		gut.ID = id
		return &gut
	}())

	returnData.Code = 0
	returnData.Message = "add gut to list success"
	log.Println(returnData.Message)
	result, _ = json.Marshal(returnData)

	ctx.SetContentType("application/json")
	ctx.SetStatusCode(200)
	ctx.Write(result) // ignore handle error
	return
}

func ModifyGut(ctx *fasthp.RequestCtx) {
	log.Println("got ModifyGut req.")
}

func DeleteGut(ctx *fasthp.RequestCtx) {
	log.Println("got DeleteGut req.")
}

package controller

import (
	"encoding/json"
	"log"
	"sync"

	"GutHub/srv/protocol"
	fasthp "github.com/valyala/fasthttp"
)

var (
	// GutList save all guts
	gutList []*protocol.Gut
	// commonID identity of gut
	commonID int
	// Mutex for make sure ID uniqueness
	mutex *sync.Mutex
)

func init() {
	gutList = make([]*protocol.Gut, 0, 256)
	commonID = len(gutList) + 1
	mutex = &sync.Mutex{}
}

// GetGutList return list of guts
func GetGutList(ctx *fasthp.RequestCtx) {
	var returnData protocol.ReturnData
	returnData.Code = -1
	log.Println("got GetGutList req.")

	result, err := json.Marshal(gutList)
	if err != nil {
		returnData.Message = "get gut list failed, error:" + err.Error()
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

	if !ctx.IsPost() {
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
	id := commonID
	commonID++
	mutex.Unlock()

	gutList = append(gutList, func() *protocol.Gut {
		gut.ID = id
		return &gut
	}())

	returnData.Code = 0
	returnData.Message = "add gut to list success"
	log.Println(returnData.Message)
	result, _ := json.Marshal(returnData)

	ctx.SetContentType("application/json")
	ctx.SetStatusCode(200)
	ctx.Write(result) // ignore handle error
	return
}

func ModifyGut(ctx *fasthp.RequestCtx) {
	var returnData protocol.ReturnData
	returnData.Code = -1
	log.Println("got ModifyGut req.")

	gutByte := ctx.FormValue("gut")
	var gutChange protocol.Gut
	err := json.Unmarshal(gutByte, &gutChange)
	if err != nil {
		returnData.Message = "unmarshal data failed, error:" + err.Error()
		log.Fatalln(returnData.Message)
		result, _ := json.Marshal(returnData)

		ctx.SetStatusCode(408)
		ctx.Write(result)
		return
	}

	for _, gut := range gutList {
		if gut.ID == gutChange.ID {
			gut = &gutChange
			returnData.Code = 0
			returnData.Message = "modify gut success"
			log.Println(returnData.Message)
			result, _ := json.Marshal(returnData)

			ctx.SetStatusCode(200)
			ctx.Write(result)
			return
		}
	}

	returnData.Message = "can not find this gut in list"
	log.Fatalln(returnData.Message)
	result, _ := json.Marshal(returnData)

	ctx.SetStatusCode(408)
	ctx.Write(result)
	return
}

func DeleteGut(ctx *fasthp.RequestCtx) {
	var returnData protocol.ReturnData
	returnData.Code = -1
	log.Println("got DeleteGut req.")

	gutByte := ctx.FormValue("gut")
	var id int
	err := json.Unmarshal(gutByte, &id)
	if err != nil {
		returnData.Message = "unmarshal data failed, error:" + err.Error()
		log.Fatalln(returnData.Message)
		result, _ := json.Marshal(returnData)

		ctx.SetStatusCode(408)
		ctx.Write(result)
		return
	}

	if id <= 0 {
		returnData.Message = "validate id"
		log.Fatalln(returnData.Message)
		result, _ := json.Marshal(returnData)

		ctx.SetStatusCode(408)
		ctx.Write(result)
		return
	}

	for index, gut := range gutList {
		if gut.ID == id {
			gutList = append(gutList[:index], gutList[index+1:]...)
			returnData.Code = 0
			returnData.Message = "delete gut success"
			log.Println(returnData.Message)
			result, _ := json.Marshal(returnData)

			ctx.SetStatusCode(200)
			ctx.Write(result)
			return
		}
	}

	returnData.Message = "can not find this gut in list"
	log.Fatalln(returnData.Message)
	result, _ := json.Marshal(returnData)

	ctx.SetStatusCode(408)
	ctx.Write(result)
	return
}

package controller

import (
	fasthp "github.com/valyala/fasthttp"
)

func GetGutList(ctx *fasthp.RequestCtx) {
	log.Println("got GetGutList req.")
}

func AddGut(ctx *fasthp.RequestCtx) {
	log.Println("got AddGut req.")
}

func ModifyGut(ctx *fasthp.RequestCtx) {
	log.Println("got ModifyGut req.")
}

func DeleteGut(ctx *fasthp.RequestCtx) {
	log.Println("got DeleteGut req.")
}

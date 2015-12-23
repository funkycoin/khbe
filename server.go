package main

import (
	"github.com/conformal/btcjson"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"./modules/setting"
	"strconv"
	"time"
)

type blockViewModel struct {
	Date  time.Time
	Block interface{}
}

func main() {
	m := martini.Classic()

	// render html templates from templates directory
	m.Use(render.Renderer(render.Options{
		Layout: "layout",
	}))

	setting.NewConfigContext()

	m.Get("/", func(r render.Render) {
		id := 1
		cmd, err := btcjson.NewGetInfoCmd(id)
		if err != nil {
			r.HTML(500, "error", err)
		}
		reply, err := btcjson.RpcSend(setting.RpcUser, setting.RpcPassword, setting.RpcHost, cmd)
		if err != nil {
			r.HTML(500, "error", err)
		}
		if reply.Result != nil {
			if data, ok := reply.Result.(*btcjson.InfoResult); ok {
				r.HTML(200, "index", data)
			}
		}
	})

	m.Get("/block/:hash", func(params martini.Params, r render.Render) {
		id := 1
		hash := params["hash"]
		cmd, err := btcjson.NewGetBlockCmd(id, hash)
		if err != nil {
			r.HTML(500, "error", err)
		}
		reply, err := btcjson.RpcSend(setting.RpcUser, setting.RpcPassword, setting.RpcHost, cmd)
		if err != nil {
			r.HTML(500, "error", err)
		}
		if reply.Result != nil {
			if data, ok := reply.Result.(*btcjson.BlockResult); ok {
				viewmodel := blockViewModel{Block: data, Date: time.Unix(data.Time, 0)}
				r.HTML(200, "block", viewmodel)
			}
		}
	})

	m.Get("/tx/:txid", func(params martini.Params, r render.Render) {
		id := 1
		hash := params["txid"]
		cmd, err := btcjson.NewGetRawTransactionCmd(id, hash)
		if err != nil {
			r.HTML(500, "error", err)
		}
		reply, err := btcjson.RpcSend(setting.RpcUser, setting.RpcPassword, setting.RpcHost, cmd)
		if err != nil {
			r.HTML(500, "error", err)
		}
		if reply.Result != nil {
			// r.HTML(200, "tx", reply.Result)
			if str, ok := reply.Result.(string); ok {
				cmd, err := btcjson.NewDecodeRawTransactionCmd(id, str)
				if err != nil {
					r.HTML(500, "error", err)
				}
				reply, err := btcjson.RpcSend(setting.RpcUser, setting.RpcPassword, setting.RpcHost, cmd)
				if err != nil {
					r.HTML(500, "error", err)
				}

				if data, ok := reply.Result.(*btcjson.TxRawDecodeResult); ok {
					r.HTML(200, "tx", data)
				} else {
					r.HTML(500, "error", "Could not parse raw transaction")
				}
			} else {
				r.HTML(500, "error", "Could not decode raw transaction")
			}
		}
	})

	m.Get("/blockByHeight/:height", func(params martini.Params, r render.Render) {
		id := 1
		height, err := strconv.ParseInt(params["height"], 10, 64)
		cmd, err := btcjson.NewGetBlockHashCmd(id, height)
		if err != nil {
			r.HTML(500, "error", err)
		}
		reply, err := btcjson.RpcSend(setting.RpcUser, setting.RpcPassword, setting.RpcHost, cmd)
		if err != nil {
			r.HTML(500, "error", err)
		}
		if reply.Result != nil {
			if hash, ok := reply.Result.(string); ok {
				r.Redirect("/block/"+hash, 302)
			}
		}
	})

	m.Run()
}

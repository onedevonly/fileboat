package main

import (
	"os"
	"github.com/valyala/fasthttp"
)

func main() {
	if _, err := os.Stat("fileboat.conf"); os.IsNotExist(err) {
		os.WriteFile("fileboat.conf", []byte(
			"PORT=6000\nROOT=./files\nCOOKIE_NAME=fb_auth\nSITE_TITLE=Fileboat\nBANNER=Fileboat\nSERVER_INFO=True\nCUSTOM_CSS=\nCUSTOM_HTML=\nCUSTOM_JS=\nPASSWORD=",
		), 0644)
	}
	if _, err := os.Stat(root); os.IsNotExist(err) {
		os.MkdirAll(root, 0755)
	}
	loadConfig()
	println("Fileboat is running on port " + port)
	fasthttp.ListenAndServe(":" + port, handler)
}
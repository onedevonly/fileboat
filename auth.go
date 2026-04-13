package main

import (
	"bytes"
	"crypto/subtle"
	"os"
	"strings"
	"time"

	"github.com/valyala/fasthttp"
)

func loadConfig() {
	data, err := os.ReadFile("fileboat.conf")
	if err != nil {
		return
	}

	for _, l := range bytes.Split(data, []byte("\n")) {
		line := strings.TrimSpace(string(l))
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}

		k, v := strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1])

		switch k {
		case "SITE_TITLE":
			if v != "" {
				title = v
			}
		case "BANNER":
			if v != "" {
				banner = v
			}
		case "SERVER_INFO":
			showStats = strings.ToLower(v) == "true"
		case "PORT":
			if v != "" {
				port = v
			}
		case "ROOT":
			if v != "" {
				root = v
			}
		case "COOKIE_NAME":
			if v != "" {
				cookieName = v
			}
		case "CUSTOM_CSS":
			customCSS = v
		case "CUSTOM_HTML":
			customHTML = v
		case "PASSWORD":
			password = v
		}
	}
}

func checkCookie(ctx *fasthttp.RequestCtx) bool {
	if password == "" {
		return true
	}
	return string(ctx.Request.Header.Cookie(cookieName)) == password
}

func auth(ctx *fasthttp.RequestCtx) bool {
	if password == "" || checkCookie(ctx) {
		return true
	}
	ctx.SetStatusCode(401)
	ctx.SetContentType("text/html")
	ctx.SetBody(pwForm())
	return false
}

func login(ctx *fasthttp.RequestCtx) {
	if password == "" {
		ctx.Redirect("/", 302)
		return
	}

	if subtle.ConstantTimeCompare([]byte(string(ctx.FormValue("password"))), []byte(password)) != 1 {
		ctx.SetStatusCode(401)
		ctx.SetContentType("text/html")
		ctx.SetBody(pwForm())
		return
	}

	cookie := fasthttp.Cookie{}
	cookie.SetKey(cookieName)
	cookie.SetValue(password)
	cookie.SetPath("/")
	cookie.SetExpire(time.Now().Add(time.Hour * 24 * 7))
	ctx.Response.Header.SetCookie(&cookie)

	ctx.Redirect("/", 302)
}
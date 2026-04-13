package main

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/valyala/fasthttp"
)

func pretty(b int64) string {
	units := []string{"B", "KB", "MB", "GB", "TB", "PB", "EB", "ZB", "YB", "BB"}
	v := float64(b)
	i := 0
	for v >= 1024 && i < len(units)-1 {
		v /= 1024
		i++
	}
	return fmt.Sprintf("%.2f %s", v, units[i])
}

func listDir(path string) ([]os.FileInfo, error) {
	entries, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	var dirs, files []os.FileInfo
	for _, entry := range entries {
		info, _ := entry.Info()
		if entry.IsDir() {
			dirs = append(dirs, info)
		} else {
			files = append(files, info)
		}
	}

	sort.Slice(dirs, func(i, j int) bool {
		return strings.ToLower(dirs[i].Name()) < strings.ToLower(dirs[j].Name())
	})
	sort.Slice(files, func(i, j int) bool {
		return strings.ToLower(files[i].Name()) < strings.ToLower(files[j].Name())
	})

	return append(dirs, files...), nil
}

func pwForm() []byte {
	var b bytes.Buffer
	b.WriteString(`<!DOCTYPE html><html id="html"><head id="head">
<meta charset="utf-8">
<title>` + title + `</title>
<link href="https://fonts.googleapis.com/icon?family=Material+Icons" rel="stylesheet">
<style>
body{margin:0;font-family:Arial;background:#0f1115;color:#eaeaea;display:flex;justify-content:center;align-items:center;min-height:100vh}
#l-c{max-width:400px;width:90%;background:#151922;border:1px solid #2a2f3a;border-radius:10px;padding:40px;text-align:center}
#l-h{font-size:26px;font-weight:600;margin-bottom:24px}
#l-f{display:flex;flex-direction:column;gap:16px}
#l-i{width:100%;padding:14px;background:#11151c;border:1px solid #2a2f3a;border-radius:6px;color:#fff;font-size:14px;box-sizing:border-box}
#l-s{width:100%;padding:14px;background:#7aa2f7;border:none;border-radius:6px;color:#fff;font-size:14px;cursor:pointer;font-weight:600}
#l-s:hover{background:#8ab4ff}
` + customCSS + `
</style>
</head><body>
<div id="l-c">
<div id="l-h">` + title + `</div>
<form id="l-f" action="/login" method="POST">
<input id="l-i" type="password" name="password" placeholder="Enter password" required>
<input id="l-s" type="submit" value="Login">
</form>
</div>
<script>
` + customJS + `
</script>
</body></html>`)
	return b.Bytes()
}

func render(files []os.FileInfo, path string) []byte {
	var b bytes.Buffer

	b.WriteString(`<!DOCTYPE html><html id="html"><head id="head">
<meta charset="utf-8">
<meta name="viewport" content="width=device-width, initial-scale=1">
<title>` + title + `</title>
<link href="https://fonts.googleapis.com/icon?family=Material+Icons" rel="stylesheet">
<style id="base-styles">
*{box-sizing:border-box}
body{margin:0;font-family:Arial,sans-serif;background:#0f1115;color:#eaeaea}
#h{text-align:center;padding:28px;font-size:26px;font-weight:600;background:#151922;border-bottom:1px solid #2a2f3a}
#c{width:92%;max-width:1100px;margin:0 auto 40px}
#ch{margin-top:20px}
#co{background:#151922;border:1px solid #2a2f3a;border-radius:10px;overflow:hidden;margin-top:20px}
#s{width:100%;padding:14px;background:#11151c;border:none;color:#fff;font-size:14px;border-bottom:1px solid #2a2f3a;outline:none}
#s::placeholder{color:#666}
table{width:100%;border-collapse:collapse}
#t{background:#1b2130;color:#aaa;font-size:12px;text-transform:uppercase}
th{padding:12px;text-align:left;font-weight:600}
.r{border-bottom:1px solid #222838;transition:background .15s}
.r:hover{background:#1b2233}
.r:last-child{border-bottom:none}
td{padding:12px}
a{color:#eaeaea;text-decoration:none;display:flex;align-items:center;gap:8px}
a:hover{color:#7aa2f7}
.ico{font-size:20px;color:#7aa2f7;vertical-align:middle}
.sz,.dt{color:#888;font-size:13px}
#f{text-align:center;padding:18px;font-size:12px;color:#777}
#f a{color:#ccc;text-decoration:none;display:inline;}
#f a:hover{color:#7aa2f7}
` + customCSS + `
</style>
<script>
function s(){var q=document.getElementById("s").value.toLowerCase();var r=document.getElementsByClassName("r");for(var i=0;i<r.length;i++)r[i].style.display=r[i].textContent.toLowerCase().includes(q)?"":"none"}
` + customJS + `
</script>
</head><body id="body">
<div id="h">` + banner + `</div>
<div id="c">
<div id="ch">` + customHTML + `</div>`)

	if path != "/" {
		b.WriteString(`<div id="pr" class="r"><a id="pl" href="../"><span class="material-icons ico">arrow_upward</span><span>..</span></a><td></td><td></td></div>`)
	}

	b.WriteString(`<div id="co">
<input id="s" onkeyup="s()" placeholder="Search files...">
<table>
<thead id="thead"><tr><th>Name</th><th>Size</th><th>Created/Modified</th></tr></thead>
<tbody id="tbody">`)

	for i, f := range files {
		name := f.Name()
		isDir := f.IsDir()
		size := "-"
		if !isDir {
			size = pretty(f.Size())
		}
		link := name
		if isDir {
			link = name + "/"
		}
		safeName := strings.ReplaceAll(strings.ReplaceAll(name, "<", "&lt;"), ">", "&gt;")

		b.WriteString(fmt.Sprintf(`<tr id="row-%d" class="r"><td id="name-cell-%d"><a id="link-%d" href="`+link+`"><span id="icon-%d" class="material-icons ico">`+icon(isDir, name)+`</span><span id="name-%d">`+safeName+`</span></a></td><td id="size-cell-%d" class="sz">`+size+`</td><td id="date-cell-%d" class="dt">`+f.ModTime().Format("2006-01-02 15:04")+`</td></tr>`, i, i, i, i, i, i, i))
	}

	b.WriteString(`</tbody></table></div></div>`)

	if showStats {
		b.WriteString(`<div id="f">Served using <a href="https://github.com/onedevonly/fileboat" target="_blank">fileboat</a> · Uptime ` + time.Since(startTime).Round(time.Second).String() + `</div>`)
	} else {
		b.WriteString(`<div id="f">Served using <a href="https://github.com/onedevonly/fileboat" target="_blank">fileboat</a></div>`)
	}

	b.WriteString(`</body></html>`)
	return b.Bytes()
}

func handler(ctx *fasthttp.RequestCtx) {
	path := string(ctx.Path())

	if path == "/login" {
		login(ctx)
		return
	}

	if !auth(ctx) {
		return
	}

	full := filepath.Join(root, path)

	if info, err := os.Stat(full); err == nil && !info.IsDir() {
		data, err := os.ReadFile(full)
		if err != nil {
			ctx.SetStatusCode(500)
			return
		}
		ctx.SetContentType("application/octet-stream")
		ctx.Write(data)
		return
	}

	files, err := listDir(full)
	if err != nil {
		ctx.SetStatusCode(404)
		return
	}

	ctx.SetContentType("text/html")
	ctx.Write(render(files, path))
}
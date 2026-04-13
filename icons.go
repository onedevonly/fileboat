package main

import (
	"path/filepath"
	"strings"
)

func icon(isDir bool, name string) string {
	if isDir {
		return "folder"
	}

	e := strings.ToLower(filepath.Ext(name))

	switch e {
	case ".exe", ".msi", ".bat", ".cmd", ".ps1", ".sh", ".vb", ".vbs":
		return "terminal"
	case ".go", ".rs", ".py", ".js", ".ts", ".jsx", ".tsx", ".java", ".c", ".cpp", ".h", ".hpp", ".cs", ".rb", ".php", ".lua", ".nim", ".zig":
		return "code"
	case ".html", ".htm", ".xhtml", ".xml", ".svg", ".xaml", ".uxml", ".jade", ".pug", ".aspx", ".jsp":
		return "html"
	case ".css", ".scss", ".sass", ".less":
		return "css"
	case ".json", ".toml", ".yml", ".yaml", ".plist", ".properties", ".env":
		return "data_object"
	case ".png", ".jpg", ".jpeg", ".gif", ".webp", ".bmp", ".ico", ".tiff", ".tif", ".heic", ".heif", ".avif":
		return "image"
	case ".mp4", ".mkv", ".mov", ".avi", ".webm", ".flv", ".wmv", ".m4v", ".mpg", ".mpeg", ".3gp", ".ogv":
		return "movie"
	case ".mp3", ".wav", ".flac", ".ogg", ".m4a", ".aac", ".wma", ".opus", ".mid", ".midi":
		return "music_note"
	case ".zip", ".rar", ".7z", ".tar", ".gz", ".bz2", ".xz", ".tgz", ".lz", ".lz4", ".zst", ".cab", ".squashfs", ".deb", ".rpm", ".pkg":
		return "archive"
	case ".pdf":
		return "picture_as_pdf"
	case ".txt", ".log", ".md", ".nfo", ".readme", ".changelog", ".license", ".gitignore", ".gitattributes", ".editorconfig":
		return "description"
	case ".doc", ".docx", ".odt", ".rtf", ".gdoc", ".gsheet", ".gslides":
		return "article"
	case ".xls", ".xlsx", ".csv", ".tsv":
		return "table_chart"
	case ".img", ".vhd", ".vhdx", ".qcow2", ".vmdk", ".vdi", ".hdd", ".ova", ".ovf", ".raw":
		return "disc_full"
	case ".apk", ".aab", ".ipa", ".appx", ".msix", ".app", ".xapk", ".jar", ".apks":
		return "smartphone"
	case ".dll", ".so", ".dylib", ".a", ".lib", ".o", ".obj", ".drv", ".sys":
		return "extension"
	case ".ttf", ".otf", ".woff", ".woff2", ".eot", ".fon", ".fnt":
		return "font_download"
	case ".psd", ".psb":
		return "brush"
	case ".sql", ".db", ".sqlite", ".sqlite3", ".duckdb":
		return "storage"
	case ".ini", ".cfg", ".conf", ".reg", ".inf", ".service", ".config", ".settings":
		return "settings"
	case ".bak", ".backup", ".old", ".orig", ".swp", ".tmp", ".cache", ".pyc", ".pyo":
		return "restore"
	default:
		return "draft"
	}
}
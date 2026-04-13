@echo off

rem Building:
rem go build -o fileboat.exe

rem Building with optimizations for smaller file size:
rem go build -ldflags="-s -w" -trimpath -o fileboat.exe

rem Building for debugging:
rem go build -gcflags="all=-N -l" -o fileboat.exe

rem Compressing:
rem upx --best fileboat.exe

rem Find the best compression method (takes a while):
rem upx --ultra-brute fileboat.exe

go build -ldflags="-s -w" -trimpath -o fileboat.exe
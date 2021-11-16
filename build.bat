@echo off
cmd /c clean
xcopy src\resources bin\resources\

go build -o bin\screenshot_debug.exe go_screenshot/src/main
cd ./bin
start ./screenshot_debug.exe
cd ../
go build -o bin\screenshot_rel.exe -ldflags="-H windowsgui" go_screenshot/src/main

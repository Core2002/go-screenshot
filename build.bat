cmd /c clean
xcopy src\resources bin\

go build -o bin\screenshot_debug.exe go_screenshot/src/main
go build -o bin\screenshot_rel.exe -ldflags="-H windowsgui" go_screenshot/src/main

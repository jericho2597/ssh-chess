@echo off
echo Building and running the Chess Engine Server...

rem Set GOARCH to amd64 for 64-bit build
set GOARCH=amd64

rem Build the Go application
go build -o server.exe ./server/server.go

if %errorlevel% neq 0 (
    echo Build failed. Please check your Go code for errors.
    pause
    exit /b %errorlevel%
)

echo Build successful. Running the application...

rem Run the application
start /B server.exe

echo Server is now running. Press any key to stop the server.
pause >nul

rem Kill the application
taskkill /F /IM server.exe >nul 2>&1
echo Server stopped.
pause
if (Get-Command "go") {
    # Run the Go program
    go run ./ghalerts.go > example.html
    Start-Process example.html
} else {
    Write-Host "Go is not installed."
}

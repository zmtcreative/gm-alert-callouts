if (Get-Command "go") {
    # Run the Go program
    go run ./showalerts.go > example.html
    Start-Process example.html
} else {
    Write-Host "Go is not installed."
}

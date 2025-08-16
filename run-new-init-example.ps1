#!/usr/bin/env pwsh

# Run the new initialization example and open the output in the default browser

Write-Host "Running new AlertCallouts initialization example..." -ForegroundColor Green

# Change to the new-init directory
Set-Location examples\new-init

# Run the Go program
try {
    go run .

    if ($LASTEXITCODE -eq 0) {
        Write-Host "Example generated successfully!" -ForegroundColor Green

        # Open the generated HTML file in the default browser
        if (Test-Path "new-init-example.html") {
            Write-Host "Opening new-init-example.html in default browser..." -ForegroundColor Cyan
            Start-Process "new-init-example.html"
        } else {
            Write-Host "Error: new-init-example.html was not created" -ForegroundColor Red
        }
    } else {
        Write-Host "Error: Go program failed with exit code $LASTEXITCODE" -ForegroundColor Red
    }
} catch {
    Write-Host "Error running example: $_" -ForegroundColor Red
} finally {
    # Return to the original directory
    Set-Location ..\..
}

Write-Host "Done!" -ForegroundColor Green

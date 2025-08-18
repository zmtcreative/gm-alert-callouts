[CmdletBinding()]
param(
    [Parameter(Mandatory = $false, HelpMessage = "View the example in a browser.")]
    [switch]$View,
    [Parameter(Mandatory = $false, HelpMessage = "Path to a markdown file to render. If not specified, the default sample will be used.")]
    [string]$File
)

if ( ($File) -and (-not (Test-Path -Path $File -PathType Leaf -ErrorAction SilentlyContinue)) ) {
    $File = $null
    Write-Host -ForegroundColor Yellow "No file specified, using default sample in the showalerts.go."
}

if (Get-Command "go") {
    # Run the Go program
    if ($null -eq $File) {
        go run ./showalerts.go > example.html
    }
    else {
        go run ./showalerts.go -f $File > example.html
    }

    if ($View) {
        Start-Process example.html
    }
    else {
        Write-Host "Example generated: example.html"
        Write-Host ""
        Write-Host "You can view the example in a browser by rerunning the script with the -View switch."
        Write-Host "The HTML output contains a 'refresh' directive to refresh the page every 10 seconds,"
        Write-Host "so you only need to use -View once to open the page in the browser if you're experimenting"
        Write-Host "and you can leave it open to see changes automatically."
        Write-Host ""
    }
} else {
    Write-Host "Go is not installed."
}

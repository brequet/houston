function Kill-Processes([string]$processName) {
    $processes = Get-Process | Where-Object {$_.Name -eq $processName}
    Write-Host "Killing all $processName processes... ($($processes.Count) processes)" -ForegroundColor DarkYellow
    $processes | Stop-Process -Force
}

# Kill all Python processes
Kill-Processes "Python"

# Kill all Firefox processes
Kill-Processes "firefox"

Kill-Processes "geckodriver"


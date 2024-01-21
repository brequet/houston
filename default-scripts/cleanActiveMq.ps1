# Define the list of directories to be emptied
$directories = @("$env:BIOMERIEUX_DATA\horus\logs", "$env:BIOMERIEUX_DATA\horus\activemq")

# Iterate over each directory
foreach ($dir in $directories) {
    # Check if the directory exists
    if (Test-Path $dir) {
        # Remove all items in the directory
        Get-ChildItem -Path $dir -Recurse | Remove-Item -Force -Recurse -ErrorAction SilentlyContinue
        Write-Output "All files have been removed from the directory: $dir"
    }
    else {
        Write-Output "Directory $dir does not exist."
    }
}


Write-Output "Truncate activemq_msgs table"
psql -U postgres -d horus_db -c "TRUNCATE TABLE public.activemq_msgs;"

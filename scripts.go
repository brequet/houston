package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

const (
	POWERSHELL_CMD     = `C:\Windows\System32\WindowsPowerShell\v1.0\powershell.exe`
	SCRIPT_FOLDER_PATH = "./default-scripts"
)

type FileDetails struct {
	Name      string    `json:"name"`
	Type      string    `json:"type"`
	ModTime   time.Time `json:"modTime"`
	Extension string    `json:"extension"`
}

func (fd FileDetails) String() string {
	return fmt.Sprintf("Name: %s, Type: %s, Extension: %s, Last Modified: %v", fd.Name, fd.Type, fd.Extension, fd.ModTime)
}

func ExecuteScriptByName(scriptName string) (string, error) {
	scriptPath := fmt.Sprintf("%s/%s", SCRIPT_FOLDER_PATH, scriptName)
	// Check if script path exists
	if _, err := os.Stat(scriptPath); err != nil {
		if os.IsNotExist(err) {
			return "", fmt.Errorf("Script file does not exist: %s\n", scriptPath)
		} else {
			return "", fmt.Errorf("Failed to check if script file exists: %s\n", err)
		}
	}

	log.Printf("Running script: '%s'\n", scriptPath)
	cmd := exec.Command(POWERSHELL_CMD, "-ExecutionPolicy", "Bypass", "-File", scriptPath)

	out, outErr := cmd.CombinedOutput()
	return string(out), outErr
}

func GetAllScripts() ([]FileDetails, error) {
	return getFilesInFolder(SCRIPT_FOLDER_PATH)
}

func getFilesInFolder(folderPath string) ([]FileDetails, error) {
	files, err := os.ReadDir(folderPath)
	if err != nil {
		return nil, err
	}

	var result []FileDetails
	for _, file := range files {
		fileInfo, err := file.Info()
		if err != nil {
			log.Fatal(err)
		}

		var fileType string
		if fileInfo.IsDir() {
			// ignore directories
			continue
		}
		fileType = "File"

		extension := filepath.Ext(fileInfo.Name())

		result = append(result, FileDetails{
			Name:      fileInfo.Name(),
			Type:      fileType,
			ModTime:   fileInfo.ModTime(),
			Extension: extension,
		})
	}

	return result, nil
}

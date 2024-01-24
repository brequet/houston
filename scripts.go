package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

const (
	POWERSHELL_CMD     = `C:\Windows\System32\WindowsPowerShell\v1.0\powershell.exe`
	SCRIPT_FOLDER_PATH = "./default-scripts"
)

type CommandData struct {
	Name      string
	Extension string
	Parameter string
}

var CommandMap = map[string]CommandData{
	"ps1": {Name: "powershell", Extension: "ps1", Parameter: "-ExecutionPolicy Bypass -File"},
	"py":  {Name: "python", Extension: "py", Parameter: ""},
}

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

	extension := getFileExtension(scriptName)
	command, ok := CommandMap[extension]
	if !ok {
		return "", fmt.Errorf("Script type not handled: '%s'\n", getFileExtension(scriptName))
	}
	args := constructArgs(command.Parameter, scriptPath)

	log.Printf("Running script: %s %s\n", command.Name, args)
	cmd := exec.Command(command.Name, args...)

	out, outErr := cmd.CombinedOutput()
	return string(out), outErr
}

func constructArgs(commmandParameter, targetFile string) []string {
	args := strings.Fields(commmandParameter)
	args = append(args, targetFile)
	return args
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

func getFileExtension(filename string) string {
	ext := filepath.Ext(filename)
	if ext == "" {
		fmt.Println("The file has no extension")
		return ""
	}
	return ext[1:] // skip the dot
}

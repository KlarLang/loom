package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

const klangBinSubstr = ".klang/bin"

func installCommand() {
	// Verifica se está rodando com sudo (apenas Unix)
	if runtime.GOOS != "windows" && os.Geteuid() == 0 {
		fmt.Println("❌ ERROR: Do not run 'loom install' with sudo or su!")
		fmt.Println("   This will install Klang in the root user's home directory.")
		fmt.Println("   Run without sudo:")
		fmt.Println("     loom install")
		return
	}

	contentToInstall := []string{}

	if len(os.Args) <= 2 {
		contentToInstall = append(contentToInstall, "klang")
	} else {
		contentToInstall = os.Args[2:]
	}

	for _, item := range contentToInstall {
		if item != "klang" {
			fmt.Printf("Warning: loom does not support %s yet. Installing klang only.\n", item)
		}
	}

	currentUser, err := user.Current()
	if err != nil {
		fmt.Printf("Error getting user home path: %v\n", err)
		return
	}

	homeUserPath := currentUser.HomeDir
	klangBasePath := filepath.Join(homeUserPath, ".klang")

	paths := []string{
		klangBasePath,
		filepath.Join(klangBasePath, "bin"),
		filepath.Join(klangBasePath, "version"),
		filepath.Join(klangBasePath, "active"),
	}

	for _, path := range paths {
		if err = os.MkdirAll(path, os.ModePerm); err != nil {
			fmt.Printf("Permission error while creating '%s': %v\n", path, err)
			return
		}
	}

	klangJarFullPath := filepath.Join(klangBasePath, "active", "klang.jar")

	// Cria script apropriado para o OS
	if err := createExecutableScript(klangBasePath, klangJarFullPath); err != nil {
		fmt.Printf("Error creating executable script: %v\n", err)
		return
	}

	// Adiciona ao PATH (diferente para Windows e Unix)
	if err := addToPath(klangBasePath, homeUserPath); err != nil {
		fmt.Printf("Warning: %v\n", err)
	}

	klangJarUrl, err := getLatestKlangJarURL()
	if err != nil {
		log.Fatalf("Error determining the latest download URL: %v", err)
		return
	}

	klangJarPath := filepath.Join(paths[3], "klang.jar")

	fmt.Printf("Downloading %s to %s...\n", klangJarUrl, klangJarPath)

	if err := downloadFile(klangJarPath, klangJarUrl); err != nil {
		log.Fatalf("Error downloading the file: %v", err)
		return
	}

	fmt.Println("Download complete!")
	fmt.Println("\n=============================================")
	fmt.Println("Klang installed successfully!")

	if runtime.GOOS == "windows" {
		fmt.Println("Restart your terminal or add to PATH manually:")
		fmt.Printf("  %s\\bin\n", klangBasePath)
	} else {
		shellConfigPath, _ := determineShellConfigPath()
		fmt.Println("Restart your terminal or run:")
		fmt.Printf("  source %s\n", shellConfigPath)
	}

	fmt.Println("Then verify installation with:")
	fmt.Println("  kc --version")
	fmt.Println("=============================================")
}

func createExecutableScript(klangBasePath, klangJarFullPath string) error {
	if runtime.GOOS == "windows" {
		// Windows: cria .bat
		kcPath := filepath.Join(klangBasePath, "bin", "kc.bat")
		kcContent := []byte(fmt.Sprintf("@echo off\r\njava -jar \"%s\" %%*", klangJarFullPath))
		return makeFile(kcContent, kcPath)
	} else {
		// Unix: cria script shell
		kcPath := filepath.Join(klangBasePath, "bin", "kc")
		kcContent := []byte(fmt.Sprintf("#!/bin/sh\njava -jar \"%s\" \"$@\"", klangJarFullPath))
		if err := makeFile(kcContent, kcPath); err != nil {
			return err
		}
		return os.Chmod(kcPath, 0755)
	}
}

func addToPath(klangBasePath, homeUserPath string) error {
	if runtime.GOOS == "windows" {
		return addToPathWindows(klangBasePath)
	}
	return addToPathUnix(klangBasePath, homeUserPath)
}

func addToPathWindows(klangBasePath string) error {
	klangBinPath := filepath.Join(klangBasePath, "bin")

	// Cria script PowerShell
	psScriptPath := filepath.Join(klangBasePath, "add-to-path.ps1")
	psContent := fmt.Sprintf(`# Add Klang to PATH
$klangBinPath = "%s"
$currentPath = [Environment]::GetEnvironmentVariable("Path", "User")
if ($currentPath -notlike "*$klangBinPath*") {
    [Environment]::SetEnvironmentVariable("Path", "$currentPath;$klangBinPath", "User")
    Write-Host "✔ Added to PATH successfully!" -ForegroundColor Green
} else {
    Write-Host "✔ Already in PATH" -ForegroundColor Green
}`, klangBinPath)

	if err := os.WriteFile(psScriptPath, []byte(psContent), 0644); err != nil {
		return fmt.Errorf("failed to create PowerShell script: %w", err)
	}

	// EXECUTA o script PowerShell automaticamente
	fmt.Println("\nAdding to PATH...")
	cmd := exec.Command("powershell", "-ExecutionPolicy", "Bypass", "-File", psScriptPath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Println("\n⚠️  Failed to add to PATH automatically. You can:")
		fmt.Println("1. Run this PowerShell script as Administrator:")
		fmt.Printf("   %s\n", psScriptPath)
		fmt.Println("\n2. Or add manually:")
		fmt.Println("   Press Win + X → System → Advanced → Environment Variables")
		fmt.Printf("   Add: %s\n", klangBinPath)
		return nil
	}

	return nil
}

func addToPathUnix(klangBasePath, homeUserPath string) error {
	shellConfigPath, err := determineShellConfigPath()
	if err != nil {
		return fmt.Errorf("could not determine shell config file: %w", err)
	}

	fmt.Printf("Shell determined. Editing file: %s\n", shellConfigPath)

	expandedConfigPath := strings.Replace(shellConfigPath, "~", homeUserPath, 1)
	klangBinPath := filepath.Join(klangBasePath, "bin")
	klangBinPathLine := fmt.Sprintf("export PATH=\"%s:$PATH\"", klangBinPath)

	if found, err := fileContains(expandedConfigPath, klangBinSubstr); err == nil && !found {
		if err := appendLine(expandedConfigPath, klangBinPathLine); err != nil {
			return fmt.Errorf("failed to add PATH to shell config file: %w", err)
		}
		fmt.Println("\nAdded ~/.klang/bin to your PATH.")
	}

	return nil
}

type GitHubRelease struct {
	TagName string `json:"tag_name"`
	Assets  []struct {
		Name               string `json:"name"`
		BrowserDownloadURL string `json:"browser_download_url"`
	} `json:"assets"`
}

func getLatestKlangJarURL() (string, error) {
	const apiURL = "https://api.github.com/repos/KlangLang/Klang/releases"

	client := &http.Client{Timeout: 10 * time.Second}

	resp, err := client.Get(apiURL)
	if err != nil {
		return "", fmt.Errorf("failed to reach GitHub API: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("GitHub API returned status: %s", resp.Status)
	}

	var releases []GitHubRelease

	if err := json.NewDecoder(resp.Body).Decode(&releases); err != nil {
		return "", fmt.Errorf("failed to parse GitHub API response: %w", err)
	}

	if len(releases) == 0 {
		return "", fmt.Errorf("no releases found")
	}

	release := releases[0]

	for _, asset := range release.Assets {
		if asset.Name == "klang.jar" {
			fmt.Printf("Found latest version: %s\n", release.TagName)
			return asset.BrowserDownloadURL, nil
		}
	}

	return "", fmt.Errorf("klang.jar not found in release %s", release.TagName)
}

func determineShellConfigPath() (string, error) {
	shellPath := os.Getenv("SHELL")
	if shellPath == "" {
		return "", fmt.Errorf("SHELL variable not defined")
	}

	shellName := filepath.Base(shellPath)
	switch shellName {
	case "bash":
		return "~/.bashrc", nil
	case "zsh":
		return "~/.zshrc", nil
	case "fish":
		return "~/.config/fish/config.fish", nil
	default:
		return "~/.profile", nil
	}
}

func appendLine(path, content string) error {
	line := content + "\n"
	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("error opening file '%s': %w", path, err)
	}
	defer f.Close()

	if _, err := f.WriteString(line); err != nil {
		return fmt.Errorf("error writing to '%s': %w", path, err)
	}
	return nil
}

func makeFile(content []byte, path string) error {
	if err := os.WriteFile(path, content, 0644); err != nil {
		fmt.Printf("Permission error while creating '%s': %v\n", path, err)
		return err
	}
	return nil
}

func fileContains(path, substring string) (bool, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return false, fmt.Errorf("error reading '%s': %w", path, err)
	}
	return strings.Contains(string(data), substring), nil
}

func downloadFile(filepath string, url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("download failed: HTTP status %s", resp.Status)
	}

	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}

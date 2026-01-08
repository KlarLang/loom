@echo off
setlocal enabledelayedexpansion

REM ------------------------------
REM Detect OS e ARCH
REM ------------------------------

for /f "tokens=2 delims==" %%a in ('wmic os get osarchitecture /value ^| find "="') do set ARCH=%%a

REM Normalizar arquitetura
if "%ARCH%"=="64-bit" set ARCH=x86_64
if "%ARCH%"=="32-bit" set ARCH=i386
if "%ARCH%"=="ARM64-based PC" set ARCH=arm64

echo Detected Architecture: %ARCH%

REM ------------------------------
REM Define version
REM ------------------------------
set VERSION=%1
if "%VERSION%"=="" set VERSION=v0.1.5

set REPO=KlangLang/loom
set FILE=loom_Windows_%ARCH%.zip
set URL=https://github.com/%REPO%/releases/download/%VERSION%/%FILE%

echo Downloading Loom %VERSION%
echo URL: %URL%
curl -L "%URL%" -o "%FILE%"

if not exist "%FILE%" (
    echo ❌ Failed to download %FILE%
    exit /b 1
)

REM ------------------------------
REM Extract ZIP
REM ------------------------------
echo Extracting...
tar -xf "%FILE%"

if errorlevel 1 (
    echo ❌ Failed to extract %FILE%
    exit /b 1
)

REM ------------------------------
REM Locate loom.exe
REM ------------------------------

set LOOMBIN=
for /r %%f in (loom.exe) do (
    set LOOMBIN=%%f
)

if "%LOOMBIN%"=="" (
    echo ❌ loom.exe not found inside archive.
    exit /b 1
)

echo Found binary: %LOOMBIN%

REM ------------------------------
REM Install bin
REM ------------------------------

set TARGET=%USERPROFILE%\bin

if not exist "%TARGET%" (
    mkdir "%TARGET%"
)

echo Installing to %TARGET%\loom.exe
copy /y "%LOOMBIN%" "%TARGET%\loom.exe" >nul

if errorlevel 1 (
    echo ❌ Failed to copy loom.exe
    exit /b 1
)

REM ------------------------------
REM Cleanup
REM ------------------------------
echo Cleaning up...
del /q "%FILE%" >nul 2>&1
rmdir /s /q "%~dp0loom_Windows_%ARCH%" >nul 2>&1

REM ------------------------------
REM Add to PATH using PowerShell
REM ------------------------------
echo.
echo Adding to PATH...

powershell -Command "$target = '%TARGET%'; $currentPath = [Environment]::GetEnvironmentVariable('Path', 'User'); if ($currentPath -notlike \"*$target*\") { [Environment]::SetEnvironmentVariable('Path', \"$currentPath;$target\", 'User'); Write-Host '✔ Added to PATH!' -ForegroundColor Green } else { Write-Host '✔ Already in PATH' -ForegroundColor Green }"

echo.
echo =============================================
echo ✔ Loom installed successfully!
echo.
echo ⚠ IMPORTANT: Close this terminal and open a NEW one
echo Then run: loom --version
echo =============================================

endlocal
@echo off
setlocal enabledelayedexpansion

REM ------------------------------
REM Detect ARCH (Windows 11 compatible)
REM ------------------------------

set ARCH=

REM Método 1: PROCESSOR_ARCHITECTURE
if "%PROCESSOR_ARCHITECTURE%"=="AMD64" set ARCH=x86_64
if "%PROCESSOR_ARCHITECTURE%"=="x86" set ARCH=i386
if "%PROCESSOR_ARCHITECTURE%"=="ARM64" set ARCH=arm64

REM Método 2: Fallback usando PowerShell se ARCH ainda estiver vazio
if "%ARCH%"=="" (
    for /f "delims=" %%i in ('powershell -Command "[System.Environment]::Is64BitOperatingSystem"') do set IS64BIT=%%i
    if "!IS64BIT!"=="True" (
        set ARCH=x86_64
    ) else (
        set ARCH=i386
    )
)

REM Método 3: Último fallback
if "%ARCH%"=="" set ARCH=x86_64

echo Detected -^> Windows-%ARCH%

REM ------------------------------
REM Get version
REM ------------------------------
set REPO=KlangLang/loom

if "%1"=="" (
    echo Fetching latest version from GitHub...
    
    set TMPJSON=%TEMP%\loom_releases.json
    
    curl -s "https://api.github.com/repos/%REPO%/releases/latest" -o "!TMPJSON!"
    
    for /f "delims=" %%i in ('powershell -Command "(Get-Content '!TMPJSON!' | ConvertFrom-Json).tag_name"') do set VERSION=%%i
    
    del "!TMPJSON!" 2>nul
    
    if "!VERSION!"=="" (
        echo ❌ Failed to fetch latest version, using fallback v0.9.0
        set VERSION=v0.9.0
    ) else (
        echo Found latest version: !VERSION!
    )
) else (
    set VERSION=%1
    echo Using specified version: !VERSION!
)

set FILE=loom_Windows_!ARCH!.zip
set URL=https://github.com/!REPO!/releases/download/!VERSION!/!FILE!

echo Downloading: !URL!

REM ------------------------------
REM Create temp directory
REM ------------------------------
set TMPDIR=%TEMP%\loom_install_!RANDOM!
mkdir "!TMPDIR!"

REM ------------------------------
REM Download
REM ------------------------------
curl -L --fail "!URL!" -o "!TMPDIR!\!FILE!"

if not exist "!TMPDIR!\!FILE!" (
    echo ❌ Failed to download !FILE!
    echo.
    echo Possible reasons:
    echo - Version !VERSION! does not exist
    echo - No release for Windows-!ARCH!
    echo - Network error
    echo.
    echo Check releases at: https://github.com/!REPO!/releases
    rmdir /s /q "!TMPDIR!"
    exit /b 1
)

REM ------------------------------
REM Extract ZIP
REM ------------------------------
echo Extracting...
tar -xf "!TMPDIR!\!FILE!" -C "!TMPDIR!"

if errorlevel 1 (
    echo ❌ Failed to extract !FILE!
    rmdir /s /q "!TMPDIR!"
    exit /b 1
)

REM ------------------------------
REM Locate loom.exe
REM ------------------------------
set LOOMBIN=
for /f "delims=" %%f in ('dir /s /b "!TMPDIR!\loom.exe" 2^>nul') do (
    set LOOMBIN=%%f
    goto :found
)
:found

if "!LOOMBIN!"=="" (
    echo ❌ loom.exe not found inside archive.
    echo Contents:
    dir /s /b "!TMPDIR!"
    rmdir /s /q "!TMPDIR!"
    exit /b 1
)

echo Found binary: !LOOMBIN!

REM ------------------------------
REM Install to %USERPROFILE%\bin
REM ------------------------------
set TARGET=%USERPROFILE%\bin

if not exist "!TARGET!" (
    mkdir "!TARGET!"
)

echo Installing to !TARGET!\loom.exe
copy /y "!LOOMBIN!" "!TARGET!\loom.exe" >nul

if errorlevel 1 (
    echo ❌ Failed to copy loom.exe
    echo Source: !LOOMBIN!
    echo Target: !TARGET!\loom.exe
    rmdir /s /q "!TMPDIR!"
    exit /b 1
)

REM ------------------------------
REM Cleanup
REM ------------------------------
echo Cleaning up...
rmdir /s /q "!TMPDIR!"

REM ------------------------------
REM Add to PATH
REM ------------------------------
echo.
echo Adding to PATH...

powershell -Command "$target = '!TARGET!'; $currentPath = [Environment]::GetEnvironmentVariable('Path', 'User'); if ($currentPath -notlike \"*$target*\") { [Environment]::SetEnvironmentVariable('Path', \"$currentPath;$target\", 'User'); Write-Host '✔ Added to PATH!' -ForegroundColor Green; Write-Host '' } else { Write-Host '✔ Already in PATH' -ForegroundColor Green; Write-Host '' }"

set PATH=!PATH!;!TARGET!

echo.
echo =============================================
echo ✔ Loom !VERSION! installed successfully!
echo.
echo You can now run: loom --version
echo (Or restart your terminal for system-wide access)
echo =============================================

endlocal
@echo off
setlocal enabledelayedexpansion

echo ===============================
echo ZenSec Batch Processor
echo ===============================
echo.

set /p folder="Enter the full path to the folder to process: "
if not exist "%folder%" (
    echo Folder does not exist.
    pause
    exit /b 1
)

set /p keyfile="Enter the full path to your keyfile: "
if not exist "%keyfile%" (
    echo Keyfile does not exist.
    pause
    exit /b 1
)

set /p mode="Do you want to Encrypt (E) or Decrypt (D)? [E/D]: "

if /i "%mode%"=="E" (
    echo.
    echo Starting ENCRYPTION process...
    for /R "%folder%" %%F in (*) do (
        if not "%%~xF"==".enc" (
            echo Encrypting: "%%F"
            zensec.exe -encrypt -file "%%F" -keyfile "%keyfile%"
        )
    )
) else if /i "%mode%"=="D" (
    echo.
    echo Starting DECRYPTION process...
    for /R "%folder%" %%F in (*.enc) do (
        echo Decrypting: "%%F"
        zensec.exe -decrypt -file "%%F" -keyfile "%keyfile%"
    )
) else (
    echo Invalid choice.
)

echo.
echo Process complete.
pause

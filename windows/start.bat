@echo off
set "ScriptFile=script.ps1"
set "ScriptPath=%~dp0%ScriptFile%"
PowerShell.exe -NoProfile -ExecutionPolicy Bypass -Command "& {Start-Process PowerShell.exe -ArgumentList '-NoProfile -ExecutionPolicy Bypass -File ""%ScriptPath%""' -Verb RunAs}"
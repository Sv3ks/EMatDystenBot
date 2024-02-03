$URL = "https://contera.dk/tabel/dysten/start.aspx"
$SCRIPT1 = "document.getElementById{(}'ctl00_ContentPlaceHolder1_ImageButton_PlusS'{)}.click{(}{)}"
$SCRIPT2 = "document.getElementById{(}'Link_Start2'{)}.click{(}{)}"

Start-Process chrome -ArgumentList "--remote-debugging-port=9222 $URL" -PassThru | Out-Null
Start-Sleep -Seconds 1

Add-Type -AssemblyName System.Windows.Forms
[System.Windows.Forms.SendKeys]::SendWait("{F12}")

Start-Sleep -Seconds 1

[System.Windows.Forms.SendKeys]::SendWait($SCRIPT1 + "{ENTER}")
Start-Sleep -Milliseconds 500
[System.Windows.Forms.SendKeys]::SendWait($SCRIPT2 + "{ENTER}")
Start-Sleep -Seconds 4

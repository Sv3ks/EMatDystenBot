$URL = "https://emat.dk/"
#$URL2 = "https://contera.dk/tabel/dysten/start.aspx"
$SCRIPT1 = "ctl00_ContentPlaceHolder1_ImageButton_PlusS.click{(}{)}"
$SCRIPT2 = "Link_Start2.click{(}{)}"
$SCRIPT3 = "for{(}let index=0;index<12;index{+}{+}{)}{{}Textbox_BrugersvarA.value=eval{(}Label_Spg.innerHTML{)};img3.click{(}{)};{}}" # Pure script: for (let index = 0; index < 12; index++) {Textbox_BrugersvarA.value=eval(Label_Spg.innerHTML);img3.click();}
$SCRIPT4 = "img4.click{(}{)}"
$SCRIPT5 = "ctl00_ContentPlaceHolder1_ImageButton_VSelvStart.click{(}{)}"
Add-Type -AssemblyName System.Windows.Forms # Allows the script to use keyboard

function StartNewExercise {
	[System.Windows.Forms.SendKeys]::SendWait($SCRIPT1 + "{ENTER}")
	Start-Sleep -Milliseconds 500
	[System.Windows.Forms.SendKeys]::SendWait($SCRIPT2 + "{ENTER}")
}

function SolveExercise {
	[System.Windows.Forms.SendKeys]::SendWait($SCRIPT3 + "{ENTER}")
	Start-Sleep -Milliseconds 500
	[System.Windows.Forms.SendKeys]::SendWait($SCRIPT4 + "{ENTER}")
	Start-Sleep -Milliseconds 500
	[System.Windows.Forms.SendKeys]::SendWait($SCRIPT5 + "{ENTER}")
}

Start-Process chrome -ArgumentList "$URL" -PassThru | Out-Null # Open Website

Write-Output "Please go through login process and return to harvesting site. (Press any key when done)"
$null = $Host.UI.RawUI.ReadKey('NoEcho,IncludeKeyDown');
Write-Output "Return to browser. Beginning in"
Write-Output "5"
Start-Sleep -Seconds 1
Write-Output "4"
Start-Sleep -Seconds 1
Write-Output "3"
Start-Sleep -Seconds 1
Write-Output "2"
Start-Sleep -Seconds 1
Write-Output "1"
Start-Sleep -Seconds 1
Write-Output "Beginning exercise harvesting. You can stop at any moment by closing this window."
[System.Windows.Forms.SendKeys]::SendWait("{F12}") # Open Console
while($true) {
	Start-Sleep -Seconds 1
	StartNewExercise
	Start-Sleep -Seconds 4 # Wait for exercise to begin
	SolveExercise
}
Start-Sleep -Seconds 1
$null = $Host.UI.RawUI.ReadKey('NoEcho,IncludeKeyDown');
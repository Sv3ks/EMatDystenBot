$VERSION = "v1.1.0"
$URL = "https://emat.dk/"
$SCRIPT1 = "ctl00_ContentPlaceHolder1_ImageButton_RegnehierarkietS.click()"
$SCRIPT2 = "Link_Start2.click()"
$SCRIPT3 = "for(let index=0;index<13;index++){Textbox_BrugersvarA.value=FacitA[ValgtSpg];img3.click();}"
$SCRIPT4 = "img4.click()"
$SCRIPT5 = "ctl00_ContentPlaceHolder1_ImageButton_VSelvStart.click()"
Add-Type -AssemblyName System.Windows.Forms # Allows the script to use keyboard

function StartNewExercise {
	Set-Clipboard $SCRIPT1
	[System.Windows.Forms.SendKeys]::SendWait("^(v){ENTER}")
	Start-Sleep -Milliseconds 250
	Set-Clipboard $SCRIPT2
	[System.Windows.Forms.SendKeys]::SendWait("^(v){ENTER}")
}

function SolveExercise {
	Set-Clipboard $SCRIPT3
	[System.Windows.Forms.SendKeys]::SendWait("^(v){ENTER}")
	Start-Sleep -Milliseconds 250
	Set-Clipboard $SCRIPT4
	[System.Windows.Forms.SendKeys]::SendWait("^(v){ENTER}")
	Start-Sleep -Milliseconds 250
	Set-Clipboard $SCRIPT5
	[System.Windows.Forms.SendKeys]::SendWait("^(v){ENTER}")
}

function CheckForUpdates {
	Write-Output "Checking for updates..."
	$HTML = Invoke-WebRequest -Uri "https://sv3ks.github.io/EMatDystenBot/web/releases/latest/"
	$LATEST = $HTML.ParsedHtml.body.document.getElementById('version').innerHTML
	if ($VERSION.Equals($LATEST)) {
		Write-Output "EMatDystenBot is up to date."
	} else {
		Write-Output "EMatDystenBot has a new update. (Press any key to start anyways)"
		$null = $Host.UI.RawUI.ReadKey('NoEcho,IncludeKeyDown');
	}
}

CheckForUpdates

Start-Process chrome -ArgumentList "$URL" -PassThru | Out-Null # Open Website

Write-Output "Please go through login process and return to harvesting site. (Press any key when done)"
$null = $Host.UI.RawUI.ReadKey('NoEcho,IncludeKeyDown');
Write-Output "Return to browser. Beginning in"
Write-Output "3"
Start-Sleep -Seconds 1
Write-Output "2"
Start-Sleep -Seconds 1
Write-Output "1"
Start-Sleep -Seconds 1
Write-Output "Beginning exercise harvesting. You can stop at any moment by closing this window."
[System.Windows.Forms.SendKeys]::SendWait("{F12}") # Open Console
$i = 0
while($true) {
	Start-Sleep -Milliseconds 250
	StartNewExercise
	SolveExercise
	$genpoints = ++$i*12*4 # exercises completed * questions pr exercise * points pr question
	Clear-Host 
	Write-Output "Generated Points: $genpoints"
}
Start-Sleep -Seconds 1
$null = $Host.UI.RawUI.ReadKey('NoEcho,IncludeKeyDown');
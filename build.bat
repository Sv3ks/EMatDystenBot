@echo off

cls
echo Installing requirements...

pip install -r requirements.txt

echo Building source...

pyinstaller --log-level ERROR --clean --onefile --name EMatDystenBot ./src/main.py

echo Build done
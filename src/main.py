from templates import templates
import settings
import ui

fileName = "settings.json"
load = {}

try:
	load = settings.Load(fileName)
except:
	load = templates["settings"]

load = ui.Start(load)

settings.Save(load,fileName,4)
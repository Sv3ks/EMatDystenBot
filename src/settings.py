import json

def Load():
	print("load settings")

def Save(settings):
	file = open('settings.json', 'w')
	json.dump(settings, file)
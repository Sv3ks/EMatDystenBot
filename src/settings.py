import json

def Load(fileName):
	file = open(fileName)
	return json.load(file)

def Save(settings, fileName, indent):
	file = open(fileName, 'w')
	json.dump(settings, file, indent=indent)
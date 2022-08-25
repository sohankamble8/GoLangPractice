package main

import (
	"os"
	"testing"
)

func TestSaveToFileAndReadFile(t *testing.T){
	os.Remove("_demoTesting")
	names := Names{"Suresh", "Sagar", "Sumant"}

	names.saveToFile("_demoTesting", " ")

	loadedNameList := readNamesFromFile("_demoTesting")

	if len(loadedNameList) != 3 {
		t.Errorf("Expected 3 names but got %v", len(loadedNameList))
	}
	os.Remove("_demoTesting")

}

func TestShuffleNames(t *testing.T){
	os.Remove("_demoTesting")
	names := Names{"Suresh", "Sagar", "Sumant"}
	shuffeledNames := names.shuffle()
	shuffeledNames.saveToFile("_demoTesting", " ")

	loadedNameList := readNamesFromFile("_demoTesting")

	if loadedNameList[0] != shuffeledNames[0] {
		t.Errorf("Expected 3 names but got %v", len(loadedNameList))
	}
	os.Remove("_demoTesting")

}
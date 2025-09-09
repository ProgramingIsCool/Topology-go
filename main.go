package main

import (
	"fmt"
	"kiliev/topology_demo_package/demo_packages/topology"
	"log"
)

func main() {
	// Load from TXT
	data, err := topology.LoadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	tree, err := topology.BuildTopologyFromData(data)
	if err != nil {
		log.Fatal(err)
	}

	// --- Export ---
	//jsonBytes, _ := tree.MarshalToJSON() Да го преобразявам в JSON за следваща функция
	xmlBytes, _ := tree.MarshalToXML()

	//fmt.Println("JSON:\n", string(jsonBytes))
	fmt.Println("XML:\n", string(xmlBytes))

	// --- Import back from JSON ---
	//newTreeFromJSON, err := topology.UnmarshalFromJSON(jsonBytes)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println("Loaded back from JSON:", newTreeFromJSON)

	// --- Import back from XML ---
	/*
		newTreeFromXML, err := topology.UnmarshalFromXML(xmlBytes)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Loaded back from XML:", newTreeFromXML)
	*/

}

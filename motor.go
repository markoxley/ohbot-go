package ohbot

import (
	"log"
	"strconv"

	"github.com/beevik/etree"
)

type Motor struct {
	Pos, Min, Max, Rest float64
	Rev, Attached       bool
	Type                string
}

func NewMotor() *Motor {
	return &Motor{
		Pos:      11,
		Min:      0,
		Max:      0,
		Rest:     0,
		Rev:      false,
		Attached: false,
		Type:     "",
	}
}

func loadMotorDefs() {
	tree := etree.NewDocument()
	if err := tree.ReadFromFile(ohbotMotorDefFile); err != nil {
		log.Fatalf("Unable to read motor definitions file: %s", err.Error())
	}

	root := tree.SelectElement("Motors")
	for _, element := range root.SelectElements("Motor") {
		indexStr := element.SelectAttrValue("Motor", "")
		index, err := strconv.Atoi(indexStr)
		if err != nil {
			log.Fatal("Error reading motor min def")
		}
		min, err := strconv.ParseFloat(element.SelectAttrValue("Min", ""), 64)
		if err != nil {
			log.Fatalf("Error reading motor min def: %s", err.Error())
		}
		max, err := strconv.ParseFloat(element.SelectAttrValue("Max", ""), 64)
		if err != nil {
			log.Fatalf("Error reading motor max def: %s", err.Error())
		}
		rest, err := strconv.ParseFloat(element.SelectAttrValue("RestPosition", ""), 64)
		if err != nil {
			log.Fatalf("Error reading motor rest def: %s", err.Error())
		}
		rev := element.SelectAttrValue("Reverse", "") == "True"

		motors[index].Min = min / 1000 * 180
		motors[index].Max = max / 1000 * 180
		motors[index].Rest = rest
		motors[index].Pos = rest
		motors[index].Rev = rev
	}
}

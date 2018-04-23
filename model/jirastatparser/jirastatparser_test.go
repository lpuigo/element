package jirastatparser

import (
	"testing"
	"encoding/json"
	"os"
)

const StatFile1 = `test/extract Tempo 2018-04-05.csv`



func TestJiraStat_SpentHourBy(t *testing.T) {
	js := NewJiraStat()
	err := js.LoadFromFile(StatFile1)
	if err != nil {
		t.Fatalf("could not LoadFromFile(%s):%s", StatFile1, err.Error())
	}

	jsns, err := js.CreateJiraStatNodes(1, 14)
	if err != nil {
		t.Fatalf("could not LoadFromFile(%s):%s", StatFile1, err.Error())
	}

	je := json.NewEncoder(os.Stdout)
	je.SetIndent("","\t")
	je.Encode(jsns)
}

package jirastatparser

import (
	ris "github.com/lpuig/novagile/src/server/manager/recordindexedset"
	"sort"
	"strconv"
	"strings"
	jn "github.com/lpuig/element/model/jirastatnode"
)

type JiraStat struct {
	Stats *ris.RecordIndexedSet
}

func NewJiraStat() *JiraStat {
	indexes := []ris.IndexDesc{}
	indexes = append(indexes, ris.NewIndexDesc("TeamAuthorWeeks", "Team", "Author", "StartWeek"))
	indexes = append(indexes, ris.NewIndexDesc("Weeks", "StartWeek"))
	res := &JiraStat{Stats: ris.NewRecordIndexedSet(indexes...)}

	return res
}

func (js *JiraStat) LoadFromFile(file string) error {
	return js.Stats.AddCSVDataFromFile(file)
}

func (js *JiraStat) SpentHourBy(indexname string) (keys []string, values []float64, err error) {
	cs, e := js.Stats.GetRecordColNumByName("Hours")
	if e != nil {
		return nil, nil, e
	}
	colTimeSpent := cs[0]
	keys = js.Stats.GetIndexKeys(indexname)
	sort.Strings(keys)
	values = make([]float64, len(keys))
	var val float64
	for i, key := range keys {
		recs := js.Stats.GetRecordsByIndexKey(indexname, key)
		val = 0.0
		for _, rec := range recs {
			if v, err := strconv.ParseFloat(rec[colTimeSpent], 64); err != nil {
				return nil, nil, err
			} else {
				val += v
			}
		}
		values[i] = val
	}
	return
}

func (js *JiraStat) CreateJiraStatNodes(minW, maxW int) (jsns []*jn.JiraStatNode, err error) {
	weeks := js.Stats.GetIndexKeys("Weeks")
	weekrange := map[string]int{}
	for _, w := range weeks {
		weekstr := strings.TrimLeft(w, "!")
		weekn, e := strconv.ParseInt(weekstr, 10, 32)
		weeknum := int(weekn)
		if e != nil {
			continue
		}
		if weeknum < minW || weeknum > maxW {
			continue
		}
		weekrange[weekstr] = int(weeknum) - minW
	}

	nbWeeks := maxW - minW + 1

	keys, hours, err := js.SpentHourBy("TeamAuthorWeeks")
	if err != nil {
		return
	}

	ota := ""
	var jsn *jn.JiraStatNode

	for i, key := range keys {
		cols := strings.Split(key, "!")[1:]
		numweek, found := weekrange[cols[2]]
		if !found {
			continue
		}
		ta := cols[0] + "-" + cols[1]
		if ta != ota {
			jsn = jn.NewBEJiraStatNode(cols[0], cols[1], nbWeeks)
			jsns = append(jsns, jsn)
			ota = ta
		}
		jsn.HourLogs[numweek] = hours[i]
	}
	return
}

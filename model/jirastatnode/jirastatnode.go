package jirastatnode

import "github.com/gopherjs/gopherjs/js"

type JiraStatNode struct {
	*js.Object

	Team     string    `json:"team"      js:"team"`
	Author   string    `json:"author"    js:"author"`
	HourLogs []float64 `json:"hour_logs" js:"hour_logs"`
}

func NewBEJiraStatNode(t, a string, nbweek int) *JiraStatNode {
	return &JiraStatNode{Team:t, Author:a, HourLogs:make([]float64,nbweek)}
}



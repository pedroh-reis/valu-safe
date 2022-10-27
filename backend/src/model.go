package main

import "github.com/pedroh-reis/valu-safe/backend/src/postgres"

type ChangeStateInput struct {
	Id string `json:"id"`
}

type GetStatisticsInput struct {
	Id        string `json:"id"`
	Timeframe string `json:"timeframe"`
	Value     int    `json:"value"`
}

type GetStatisticsResult struct {
	TimesUnlocked int                         `json:"timesunlocked"`
	History       []postgres.GetHistoryResult `json:"history"`
}

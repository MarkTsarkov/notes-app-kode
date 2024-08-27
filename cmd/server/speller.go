package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type SpellCheckResponse struct {
	Code int    `json:"code"`
	Pos  int    `json:"pos"`
	Row  int    `json:"row"`
	Col  int    `json:"col"`
	Len  int    `json:"len"`
	Word string `json:"word"`
	S    []string `json:"s"`
}

func checkSpelling(note string) ([]SpellCheckResponse, error) {
	
	query := url.QueryEscape(note)
	apiURL := fmt.Sprintf("https://speller.yandex.net/services/spellservice.json/checkText?text=%s", query)

	resp, err := http.Get(apiURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var spellCheckResponse []SpellCheckResponse
	if err := json.NewDecoder(resp.Body).Decode(&spellCheckResponse); err != nil {
		return nil, err
	}

	return spellCheckResponse, nil
}

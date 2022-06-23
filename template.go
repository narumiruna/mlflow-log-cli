package main

import (
	"encoding/json"
	"io/ioutil"
	"math"
	"time"
)

type Param struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Metric struct {
	Key       string  `json:"key"`
	Value     float64 `json:"value"`
	Timestamp int64   `json:"timestamp"` // unix timestamp in milliseconds
	Step      int64   `json:"step"`
}

type Tag struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type MLflowData struct {
	Params  []Param  `json:"params,omitempty"`
	Metrics []Metric `json:"metrics,omitempty"`
	Tags    []Tag    `json:"tags,omitempty"`
}

func main() {
	params := []Param{
		{Key: "batch_size", Value: "32"},
		{Key: "lr", Value: "0.001"},
	}

	var metrics []Metric
	for i := 0; i < 10; i++ {
		metric := Metric{
			Key:       "loss",
			Value:     -math.Log(float64(i) + 1),
			Timestamp: time.Now().Add(time.Duration(i) * time.Second).UnixMilli(),
			Step:      int64(i),
		}
		metrics = append(metrics, metric)
	}

	tags := []Tag{
		{Key: "task", Value: "classification"},
		{Key: "type", Value: "supervised"},
	}

	data := MLflowData{
		Params:  params,
		Metrics: metrics,
		Tags:    tags,
	}

	bytes, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile("template.json", bytes, 0644)
	if err != nil {
		panic(err)
	}

}

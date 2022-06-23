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

type MLflowLog struct {
	Params  []Param  `json:"params,omitempty"`
	Metrics []Metric `json:"metrics,omitempty"`
	Tags    []Tag    `json:"tags,omitempty"`
}

func main() {
	params := []Param{
		{Key: "key1", Value: "value1"},
		{Key: "key2", Value: "value2"},
	}

	var metrics []Metric
	for i := 0; i < 100; i++ {
		metric := Metric{
			Key:       "log_x",
			Value:     math.Log(float64(i) + 1),
			Timestamp: time.Now().UnixMilli(),
			Step:      int64(i),
		}
		metrics = append(metrics, metric)
	}

	tags := []Tag{
		{Key: "key1", Value: "value1"},
		{Key: "key2", Value: "value2"},
	}

	data := MLflowLog{
		Params:  params,
		Metrics: metrics,
		Tags:    tags,
	}

	bytes, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile("examples/data.json", bytes, 0644)
	if err != nil {
		panic(err)
	}

}

package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// https://prometheus.io/docs/tutorials/instrumenting_http_server_in_go/
var gauge = promauto.NewGauge(prometheus.GaugeOpts{
	Name: "one_label_instance_details",
	ConstLabels: prometheus.Labels{
		"one_label_instance_name":   "onelabel",
		"one_label_instance_status": "ok",
	},
})

func main() {
	fmt.Println("Server started. Listening on port 2112")
	go processMetrics()
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/10", generateTenDifferentLabelValues)
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe("localhost:2112", nil)
}

func homeHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello, this is the home page!")
}

func processMetrics() {
	for {
		randValue := rand.Float64() * 100.0
		gauge.Set(randValue)
		time.Sleep(1 * time.Second)
	}
}

func generateTenDifferentLabelValues(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, `
	Hello, 
	Generating 10 different values for instance_name label!
	You should now see ten_labels_instance_details with count 10 in top 10 series
	1 series for each label	
`)

	for i := 0; i < 10; i++ {
		instanceName := fmt.Sprintf("tenlabel%d", i)
		gauge := promauto.NewGauge(prometheus.GaugeOpts{
			Name: "ten_label_values_instance_details",
			ConstLabels: prometheus.Labels{
				"ten_label_values_instance_name":   instanceName,
				"ten_label_values_instance_status": "ok",
			},
		})
		gauge.Set(77.77)
	}
}
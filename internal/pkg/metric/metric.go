package metric

import "github.com/prometheus/client_golang/prometheus"

var GetRequestsCounter = prometheus.NewCounter(prometheus.CounterOpts{
	Name: "get_requests_total",
	Help: "The total number of GET requests",
})

func init() {
	prometheus.MustRegister(GetRequestsCounter)
}

package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/common/log"
	"github.com/prometheus/common/version"
)

var (
	showVersion   = flag.Bool("version", false, "Print version information.")
	accessToken   = flag.String("config.token", "", "papertrail account token.")
	listenAddress = flag.String("web.listen-address", ":9098", "The address to listen on for HTTP requests.")
	metricsPath   = flag.String("web.telemetry-path", "/metrics", "Path under which to expose metrics.")
)

type LogTransfer struct {
	HardLimit   int64
	PlanLimit   int64
	Used        int64
	UsedPercent float64
}

type Measurement struct {
	LogTransfer *LogTransfer
	Success     int
	Duration    float64
}

func getPapertrail() (*LogTransfer, error) {
	return &LogTransfer{
		HardLimit:   0,
		PlanLimit:   0,
		Used:        0,
		UsedPercent: 10.0,
	}, nil
}

func getMeasurement() *Measurement {
	start := time.Now()
	success := 0
	logTransfer, err := getPapertrail()
	duration := time.Since(start).Seconds()

	if err == nil {
		log.Debugf("OK to call papertrail api (after %fs).", duration)
		success = 1
	} else {
		log.Infof("ERROR: %s (failed after %fs).", err, duration)
	}

	return &Measurement{
		LogTransfer: logTransfer,
		Duration:    duration,
		Success:     success,
	}
}

func init() {
	prometheus.MustRegister(version.NewCollector("papertrail_exporter"))
}

func main() {
	flag.Parse()

	if *showVersion {
		fmt.Fprintln(os.Stdout, version.Print("papertrail_exporter"))
		os.Exit(0)
	}

	log.Infoln("Starting papertrail-exporter", version.Info())

	http.HandleFunc(*metricsPath, func(w http.ResponseWriter, r *http.Request) {
		measurement := getMeasurement()
		fmt.Fprintf(w, "papertrail_log_transfer_hard_limit{} %d\n", measurement.LogTransfer.HardLimit)
		fmt.Fprintf(w, "papertrail_log_transfer_plan_limit{} %d\n", measurement.LogTransfer.PlanLimit)
		fmt.Fprintf(w, "papertrail_log_transfer_used{} %d\n", measurement.LogTransfer.Used)
		fmt.Fprintf(w, "papertrail_log_transfer_used_parcent{} %f\n", measurement.LogTransfer.UsedPercent)
		fmt.Fprintf(w, "papertrail_duration_seconds{} %f\n", measurement.Duration)
		fmt.Fprintf(w, "papertrail_success{} %d\n", measurement.Success)
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`<html>
			<head><title>Script Exporter</title></head>
			<body>
			<h1>Papertrail Exporter</h1>
			<p><a href="` + *metricsPath + `">Metrics</a></p>
			</body>
			</html>`))
	})

	log.Infoln("Listening on", *listenAddress)

	if err := http.ListenAndServe(*listenAddress, nil); err != nil {
		log.Fatalf("Error starting HTTP server: %s", err)
	}
}

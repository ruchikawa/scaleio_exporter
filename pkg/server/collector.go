package server

import (
	"bytes"
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

type capacityMetrics struct {
	CapacityInUseInKb     float64 `json:"capacityInUseInKb"`
	ThinCapacityInUseInKb float64 `json:"thinCapacityInUseInKb"`
	MaxCapacityInKb       float64 `json:"maxCapacityInKb"`
}

var (
	capacityInUseInKb = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "capacityInUseInKb",
			Help: "capacityInUseInKb help",
		},
	)

	thinCapacityInUseInKb = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "thinCapacityInUseInKb",
			Help: "thinCapacityInUseInKb help",
		},
	)

	maxCapacityInKb = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "maxCapacityInKb",
			Help: "maxCapacityInKb help",
		},
	)
)

func init() {
	prometheus.MustRegister(capacityInUseInKb)
	prometheus.MustRegister(thinCapacityInUseInKb)
	prometheus.MustRegister(maxCapacityInKb)
}

func getStats(args *Args) {

	urlTarget := "https://" + args.IPAddr + "/api/types/System/instances/action/querySelectedStatistics"

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: args.Insecure,
			},
		},
	}

	query := `{"properties":["maxCapacityInKb", "capacityInUseInKb", "thinCapacityInUseInKb"]}`

	for {

		token := getToken(args)

		var p capacityMetrics
		req, err := http.NewRequest("POST", urlTarget, bytes.NewBuffer([]byte(query)))
		if err != nil {
			log.Fatal(err)
		}

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Basic "+basicAuth("", string(token)))

		resp, err := client.Do(req)
		if err != nil {
			log.Fatal(err)
		}

		body, error := ioutil.ReadAll(resp.Body)
		if error != nil {
			log.Fatal(error)
		}

		log.Println(string(body))
		err = json.Unmarshal(body, &p)
		if err != nil {
			log.Println(err)
		}

		capacityInUseInKb.Set(p.CapacityInUseInKb)
		thinCapacityInUseInKb.Set(p.ThinCapacityInUseInKb)
		maxCapacityInKb.Set(p.MaxCapacityInKb)

		time.Sleep(time.Duration(args.Refresh) * time.Second)
	}
}

func getToken(args *Args) (token []byte) {
	urlTarget := "https://" + args.IPAddr + "/api/login"
	username := args.Username
	password := args.Password
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: args.Insecure,
			},
		},
	}
	req, err := http.NewRequest("GET", urlTarget, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.SetBasicAuth(username, password)

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	token, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	return
}

// maybe this function can be deleted.
func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(strings.Replace(auth, "\"", "", -1)))
}

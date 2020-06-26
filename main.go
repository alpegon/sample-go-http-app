package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

var responseTime, port int
var colour, version string

func getIP(r *http.Request) string {
	forwarded := r.Header.Get("X-FORWARDED-FOR")
	if forwarded != "" {
		return forwarded
	}
	return r.RemoteAddr
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	time.Sleep(time.Duration(responseTime) * time.Second)
	fmt.Fprintf(w, "/demo\n/health\n/ip\n/version")
	fmt.Printf("%v - /\n", getIP(r))
}

func demoHandler(w http.ResponseWriter, r *http.Request) {
	time.Sleep(time.Duration(responseTime) * time.Second)
	fmt.Fprintf(w, fmt.Sprintf("<html><body><h1 style=\"color:%v\">%v</h1></body></html>", colour, version))
	fmt.Printf("%v - /demo\n", getIP(r))
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	time.Sleep(time.Duration(responseTime) * time.Second)
	fmt.Fprintf(w, "OK")
	fmt.Printf("%v - /healthz\n", getIP(r))
}

func ipHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Content-Type", "application/json")
	resp, _ := json.Marshal(map[string]string{
		"ip": getIP(r),
	})
	fmt.Printf("%v - /ip\n", getIP(r))
	w.Write(resp)
}

func versionHandler(w http.ResponseWriter, r *http.Request) {
	time.Sleep(time.Duration(responseTime) * time.Second)
	fmt.Fprintf(w, fmt.Sprintf("%v", version))
	fmt.Printf("%v - /version\n", getIP(r))
}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", rootHandler)
	router.HandleFunc("/demo", demoHandler)
	router.HandleFunc("/healthz", healthCheckHandler)
	router.HandleFunc("/ip", ipHandler)
	router.HandleFunc("/version", versionHandler)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), router))
}

func getIntEnvVar(envVar string) int {
	val, err := strconv.ParseInt(os.Getenv(envVar), 10, 0)
	if err != nil {
		fmt.Printf("Error parsing %v\n", envVar)
		fmt.Println(err)
		os.Exit(1)
	}
	return int(val)
}

func getStringEnvVar(envVar string) string {
	val := os.Getenv(envVar)
	if val == "" {
		fmt.Printf("Empty value for %v\n", envVar)
		os.Exit(1)
	}
	return val
}
func main() {
	colour = getStringEnvVar("COLOUR")
	loadTime := getIntEnvVar("LOAD_TIME")
	responseTime = getIntEnvVar("RESPONSE_TIME")
	port = getIntEnvVar("PORT")
	version = getStringEnvVar("VERSION")
	fmt.Printf("COLOUR=%v\nLOAD_TIME=%v\nRESPONSE_TIME=%v\nPORT=%v\nVERSION=%v\n", colour, loadTime, responseTime, port, version)
	time.Sleep(time.Duration(loadTime) * time.Second)
	handleRequests()
}

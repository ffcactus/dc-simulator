package main

import (
	"dc-simulator/dto"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
)

var ip = ""

func init() {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		log.Fatal(err)
	}
	for _, address := range addrs {
		// check the address type and if it is not a loopback the display it
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ip = ipnet.IP.String()
				return
			}
		}
	}
	log.Fatal("can not get IP address")
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func getProcessors(w http.ResponseWriter, r *http.Request) {
	processors := []dto.Processor{}
	for i := 1; i < 3; i++ {
		processors = append(processors, dto.Processor{
			Location:     fmt.Sprintf("Socket %d", i),
			Model:        "Intel Xeon Gold 8804",
			SerialNumber: fmt.Sprintf("SN_%s", ip),
			Manufacture:  "Intel",
			PartNumber:   fmt.Sprintf("Part_%s", ip),
		})
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(dto.ProcessorResponse{
		Items: processors,
	})
	w.WriteHeader(http.StatusOK)
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/processors", getProcessors)
	log.Fatal(http.ListenAndServe(":80", nil))
}

func main() {
	handleRequests()
}

// GetLocalIP returns the non loopback local IP of the host
func GetLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, address := range addrs {
		// check the address type and if it is not a loopback the display it
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}

package main

import (
	"dc-simulator/dto"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
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

// getProcessors is the handler for /processors.
func getProcessors(w http.ResponseWriter, r *http.Request) {
	items := []dto.Processor{}
	for i := 1; i < 3; i++ {
		items = append(items, dto.Processor{
			Location:     fmt.Sprintf("Socket %d", i),
			Model:        "Intel Xeon Gold 8804",
			SerialNumber: fmt.Sprintf("PROCESSOR_SN_%d_%s", i, ip),
			Manufacture:  "Intel",
			PartNumber:   fmt.Sprintf("MEM_Part_%d_%s", i, ip),
		})
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(dto.ProcessorResponse{
		Items: items,
	})
	w.WriteHeader(http.StatusOK)
}

// getMemory is the handler for /memory.
func getMemory(w http.ResponseWriter, r *http.Request) {
	items := []dto.Memory{}
	for i := 1; i < 9; i++ {
		items = append(items, dto.Memory{
			Location:     fmt.Sprintf("Slot %d", i),
			Model:        "DDR4",
			SerialNumber: fmt.Sprintf("Memory_SN_%d_%s", i, ip),
			Manufacture:  "Samsung",
			PartNumber:   fmt.Sprintf("Memory_Part_%d_%s", i, ip),
			Capacity:     8,
		})
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(dto.MemoryResponse{
		Items: items,
	})
	w.WriteHeader(http.StatusOK)
}

// getDrives is the handler for /drives.
func getDrives(w http.ResponseWriter, r *http.Request) {
	items := []dto.Drive{}
	for i := 1; i < 9; i++ {
		items = append(items, dto.Drive{
			Location:     fmt.Sprintf("Slot %d", i),
			Model:        "HDD SAS",
			SerialNumber: fmt.Sprintf("Drive_SN_%d_%s", i, ip),
			Manufacture:  "Western Digital",
			PartNumber:   fmt.Sprintf("Drive_Part_%d_%s", i, ip),
			Capacity:     2048,
		})
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(dto.DriveResponse{
		Items: items,
	})
	w.WriteHeader(http.StatusOK)
}

// getPowerSupplies is the handler for /power-supplies.
func getPowerSupplies(w http.ResponseWriter, r *http.Request) {
	items := []dto.PowerSupply{}
	for i := 1; i < 5; i++ {
		items = append(items, dto.PowerSupply{
			Location:     fmt.Sprintf("Slot %d", i),
			Model:        "PS1000",
			SerialNumber: fmt.Sprintf("PowerSupply_SN_%d_%s", i, ip),
			Manufacture:  "Huawei",
			PartNumber:   fmt.Sprintf("PowerSupply_Part_%d_%s", i, ip),
		})
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(dto.PowerSupplyResponse{
		Items: items,
	})
	w.WriteHeader(http.StatusOK)
}

// getFans is the handler for /fans.
func getFans(w http.ResponseWriter, r *http.Request) {
	items := []dto.Fan{}
	for i := 1; i < 5; i++ {
		items = append(items, dto.Fan{
			Location:     fmt.Sprintf("Slot %d", i),
			Model:        "F110",
			SerialNumber: fmt.Sprintf("Fan_SN_%d_%s", i, ip),
			Manufacture:  "Huawei",
			PartNumber:   fmt.Sprintf("Fan_Part_%d_%s", i, ip),
		})
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(dto.FanResponse{
		Items: items,
	})
	w.WriteHeader(http.StatusOK)
}

// getCards is the handler for /cards.
func getCards(w http.ResponseWriter, r *http.Request) {
	items := []dto.Card{}
	items = append(items, dto.Card{
		Location:     fmt.Sprintf("Slot 1"),
		Model:        "LSI1288",
		SerialNumber: fmt.Sprintf("Card_SN_1_%s", ip),
		Manufacture:  "Huawei",
		PartNumber:   fmt.Sprintf("Card_Part_1_%s", ip),
	})
	items = append(items, dto.Card{
		Location:     fmt.Sprintf("Slot 2"),
		Model:        "Mazz220",
		SerialNumber: fmt.Sprintf("Card_SN_2_%s", ip),
		Manufacture:  "Huawei",
		PartNumber:   fmt.Sprintf("Card_Part_2_%s", ip),
	})

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(dto.CardResponse{
		Items: items,
	})
	w.WriteHeader(http.StatusOK)
}

func handleRequests(port string) {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/processors", getProcessors)
	http.HandleFunc("/memory", getMemory)
	http.HandleFunc("/drives", getDrives)
	http.HandleFunc("/power-supplies", getPowerSupplies)
	http.HandleFunc("/fans", getFans)
	http.HandleFunc("/cards", getCards)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Printf("Usage: executable port")
		os.Exit(-1)
	}
	handleRequests(args[1])
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

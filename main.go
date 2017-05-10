package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"runtime"
	"time"

	"github.com/nicholasjackson/http-info/templates"
)

type stats struct {
	Arch    string
	OS      string
	NumCPUs int
	Network []network
	AWS     aws
}

type network struct {
	Name       string
	IPAdresses []string
}

type aws struct {
	PrivateIP string
	PublicIP  string
}

func main() {
	http.DefaultServeMux.HandleFunc("/", infoHandler)

	log.Println("Starting server on port 9091")
	http.ListenAndServe(":9091", http.DefaultServeMux)
}

func infoHandler(rw http.ResponseWriter, r *http.Request) {
	t, err := template.New("stats").Parse(templates.StatsTemplate)
	if err != nil {
		log.Println(err)
		rw.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(rw, "Error")
		return
	}

	err = t.Execute(rw, getStats())
	if err != nil {
		log.Println(err)
	}
}

func getStats() stats {
	return stats{
		Arch:    runtime.GOARCH,
		OS:      runtime.GOOS,
		NumCPUs: runtime.NumCPU(),
		Network: getNetwork(),
		AWS:     getAWS(),
	}
}

func getNetwork() []network {
	networks := make([]network, 0)

	i, _ := net.Interfaces()
	for _, ifc := range i {
		net := network{Name: ifc.Name}

		addrs, _ := ifc.Addrs()
		addrsString := make([]string, 0)

		for _, addr := range addrs {
			addrsString = append(addrsString, addr.String())
		}

		net.IPAdresses = addrsString
		networks = append(networks, net)
	}

	return networks
}

func getAWS() aws {
	http.DefaultClient.Timeout = 1000 * time.Millisecond
	data := aws{}

	resp, err := http.Get("http://169.254.169.254/latest/meta-data/local-ipv4")
	if err == nil {
		defer resp.Body.Close()
		ip, _ := ioutil.ReadAll(resp.Body)
		data.PrivateIP = string(ip)
	}

	resp, err = http.Get("http://169.254.169.254/latest/meta-data/public-ipv4")
	if err == nil {
		defer resp.Body.Close()
		ip, _ := ioutil.ReadAll(resp.Body)
		data.PublicIP = string(ip)
	}

	return data
}

package main

import (
	"flag"
	"fmt"
	"github.com/atotto/clipboard"
	"golang.org/x/crypto/ssh/terminal"
)

var print = fmt.Println

type Service struct {
	Id, Name, Password string
}

func listServices(services []Service) {
	print("Available services:")
	for _, i := range services {
		print(i.Id, "-", i.Name)
	}
	print("*********")

}

func getServicePassword(serviceId *string, services []Service) string {
	for _, i := range services {
		if *serviceId == i.Id {
			if checkMasterPassword() {
				clipboard.WriteAll(i.Password)
				return "The password for " + i.Name + " is in your clipboard"
			}
			return "Invalid master password"
		}
	}
	listServices(services)
	return *serviceId + " was not found\n*********"
}

func getServices() []Service {
	var services = []Service{}
	services = append(services, Service{"1", "Netflix", "myNetflixPassword"})
	services = append(services, Service{"2", "Spotify", "mySpotifyPassword"})
	return services
}

func checkMasterPassword() bool {
	print("Enter master password: ")
	password, err := terminal.ReadPassword(0)
	masterPassword := "12345"
	if err == nil && string(password) == masterPassword {
		return true
	}
	return false
}

func main() {
	print("**Password manager**")
	service := flag.String("service", "0", "Service ID. 0 to display all")
	flag.Parse()
	message := getServicePassword(service, getServices())
	print(message)
}

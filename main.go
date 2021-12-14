package main

import (
	"flag"
	"fmt"
	"github.com/atotto/clipboard"
	"golang.org/x/crypto/ssh/terminal"
	"strconv"
	"time"
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

func setClipboard(service Service, ttl *int) {
	clipboard.WriteAll(service.Password)
	print("The password for " + service.Name + " is in your clipboard, but will disapear in " + strconv.Itoa(*ttl) + " seconds!")
	time.Sleep(time.Duration(*ttl) * time.Second)
	clipboard.WriteAll("")
	print("Your clipboard has been erased")
}

func getServicePassword(serviceId *string, services []Service, ttl *int) bool {
	for _, i := range services {
		if *serviceId == i.Id {
			if checkMasterPassword() {
				setClipboard(i, ttl)
				return true
			}
			print("Invalid master password")
			return false
		}
	}
	print(*serviceId + " was not found\n*********")
	listServices(services)
	return false
}

func checkMasterPassword() bool {
	print("Enter master password: ")
	password, err := terminal.ReadPassword(0)
	if err == nil && string(password) == masterPassword {
		return true
	}
	return false
}

func main() {
	print("**Password manager**")
	service := flag.String("s", "", "Service ID")
	ttl := flag.Int("t", 5, "Time in seconds the clipboard with the password will be valid for")

	flag.Parse()
	if *service == "" {
		listServices(getServices())
	} else {
		getServicePassword(service, getServices(), ttl)
	}

}

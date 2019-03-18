package main

import (
	"crypto/tls"
	"errors"
	"flag"
	"fmt"
	"syscall"

	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"

	"golang.org/x/crypto/ssh/terminal"
)

type (
	// Credentials struct is the UCP credentials used for login
	Credentials struct {
		Username string `json:"username,omitempty"`
		Password string `json:"password,omitempty"`
		Token    string `json:"token,omitempty"`
	}
	// LoginResponse struct is the reply to the login attempt, it contains the session token
	LoginResponse struct {
		Token string `json:"auth_token,omitempty"`
	}
	// DockerTrustedRegistry struct is the DTR config stored in UCP
	DockerTrustedRegistry struct {
		HostAddress              string `json:"hostAddress"`
		ServiceID                string `json:"serviceID"`
		CABundle                 string `json:"caBundle"`
		BatchScanningDataEnabled bool   `json:"batchScanningDataEnabled"`
	}
	// Registries struct is a list of DockerTrustedRegistry configs
	Registries struct {
		Registries []DockerTrustedRegistry
	}
)

func main() {
	help := flag.Bool("-h", false, "print help menu")
	address := flag.String("a", "", "UCP address")
	username := flag.String("u", "", "UCP username")
	scanToggle := flag.String("s", "disable", "enable or disable UCP Scanning Data endpoint")
	flag.Parse()
	if *help {
		flag.PrintDefaults()
		return
	}
	if *username == "" {
		fmt.Println("Error: UCP Username is not provided")
		flag.PrintDefaults()
		return
	}
	if *address == "" {
		fmt.Println("Error: UCP Address is not provided")
		flag.PrintDefaults()
		return
	}
	if *scanToggle != "enable" && *scanToggle != "disable" {
		fmt.Printf("Error: -s cannot be %s ; choose enable or disable \n", *scanToggle)
		flag.PrintDefaults()
		return
	}
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	token, err := getUCPToken(*username, *address, client)
	if err != nil {
		fmt.Println(err)
		return
	}

	originalDTRConfigs, err := getDTRConfigs(*address, token, client)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("***** Please save this output; Original config prior to %s Scanning Data endpoint *****\n\n\n", *scanToggle)
	printPrettyJson(originalDTRConfigs)

	err = toggleScanFlag(*address, token, client, originalDTRConfigs, *scanToggle == "disable")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("\n\n\n ***** Successfuly %sd UCP Scanning Data endpoint. New Config ***** \n\n\n", *scanToggle)
	newDTRConfigs, err := getDTRConfigs(*address, token, client)
	if err != nil {
		fmt.Println(err)
		return
	}
	printPrettyJson(newDTRConfigs)
	return
}

func printPrettyJson (registries Registries) {
	src, err := json.Marshal(registries)
	if err != nil {
		fmt.Println(err)
		return
	}
	dst := &bytes.Buffer{}
	if err := json.Indent(dst, src, "", "  "); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(dst.String())
	return
}

func getUCPToken(username string, adress string, client *http.Client) (string, error) {
	var password string
	fmt.Printf("Password for UCP user %s: ", username)
	bytePassword, err := terminal.ReadPassword(int(syscall.Stdin))
	if err != nil {
		return "", fmt.Errorf("Failed to read password: %v", err)
	}
	fmt.Println("")
	password = string(bytePassword)
	if password == "" {
		return "", errors.New("Password is empty")
	}

	creds := Credentials{
		Username: username,
		Password: password,
	}
	reqJSON, err := json.Marshal(creds)
	if err != nil {
		return "", err
	}
	endpoint := &url.URL{
		Scheme: "https",
		Host:   adress,
		Path:   path.Join("/", "auth", "login"),
	}
	resp, err := client.Post(endpoint.String(), "application/json", bytes.NewBuffer(reqJSON))
	if err != nil {
		return "", err
	}
	body, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if resp.StatusCode == 200 {
		var loginResp LoginResponse
		if err := json.Unmarshal(body, &loginResp); err != nil {
			return "", err
		}
		return loginResp.Token, nil
	}
	return "", errors.New("Failed to get UCP session token")
}

func toggleScanFlag(address string, token string, client *http.Client, registries Registries, disable bool) error {
	bearer := "Bearer " + token
	endpoint := &url.URL{
		Scheme: "https",
		Host:   address,
		Path:   path.Join("/", "api", "ucp", "config", "dtr"),
	}

	for _, dtrconfig := range registries.Registries {
		dtrconfig.BatchScanningDataEnabled = !disable
		reqJSON, err := json.Marshal(dtrconfig)
		if err != nil {
			return err
		}
		req, err := http.NewRequest("POST", endpoint.String(), bytes.NewBuffer(reqJSON))
		req.Header.Add("Authorization", bearer)
		req.Header.Add("accept", "application/json")

		resp, err := client.Do(req)
		if err != nil {
			return err
		}
		if resp.StatusCode != 200 {
			fmt.Printf("Successfuly disabled scan to %s", dtrconfig.HostAddress)
		}
	}
	return nil
}

func getDTRConfigs(address string, token string, client *http.Client) (Registries, error) {
	bearer := "Bearer " + token
	endpoint := &url.URL{
		Scheme: "https",
		Host:   address,
		Path:   path.Join("/", "api", "ucp", "config", "dtr"),
	}
	req, err := http.NewRequest("GET", endpoint.String(), nil)
	req.Header.Add("Authorization", bearer)
	req.Header.Add("accept", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return Registries{}, err
	}
	body, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	var registries Registries
	if resp.StatusCode == 200 {
		if err := json.Unmarshal(body, &registries); err != nil {
			return Registries{}, err
		}
	}
	if len(registries.Registries) == 0 {
		return Registries{}, errors.New("Could not find any stored DTR config")
	}
	return registries, nil
}

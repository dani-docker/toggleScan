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
	Credentials struct {
		Username string `json:"username,omitempty"`
		Password string `json:"password,omitempty"`
		Token    string `json:"token,omitempty"`
	}
	LoginResponse struct {
		Token string `json:"auth_token,omitempty"`
	}

	DockerTrustedRegistry struct {
		HostAddress              string `json:"hostAddress"`
		ServiceID                string `json:"serviceID"`
		CABundle                 string `json:"caBundle"`
		BatchScanningDataEnabled bool   `json:"batchScanningDataEnabled"`
	}
	Registries struct {
		Registries []DockerTrustedRegistry
	}
)

func main() {
	address := flag.String("a", "", "UCP Address")
	username := flag.String("u", "", "username")
	flag.Parse()
	if *username == "" {
		fmt.Println("Error: UCP Username is not provided")
		return
	}
	if *address == "" {
		fmt.Println("Error: UCP Address is not provided")
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

	dst := &bytes.Buffer{}
	fmt.Println("***** Please save this output; Original config \n\n\n *****")
	src, err := json.Marshal(originalDTRConfigs)
	if err != nil {
		fmt.Println(err)
		return
	}
	if err := json.Indent(dst, src, "", "  "); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(dst.String())

	err = disableScanFlag(*address, token, client, originalDTRConfigs)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("***** Successfuly disabled scan to all DTR instances. New Config *****")
	newdst := &bytes.Buffer{}
	newDTRConfigs, err := getDTRConfigs(*address, token, client)
	newsrc, err := json.Marshal(newDTRConfigs)
	if err != nil {
		fmt.Println(err)
		return
	}
	if err := json.Indent(newdst, newsrc, "", "  "); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(newdst.String())
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

func disableScanFlag(address string, token string, client *http.Client, registries Registries) error {
	bearer := "Bearer " + token
	endpoint := &url.URL{
		Scheme: "https",
		Host:   address,
		Path:   path.Join("/", "api", "ucp", "config", "dtr"),
	}

	for _, dtrconfig := range registries.Registries {
		dtrconfig.BatchScanningDataEnabled = false
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

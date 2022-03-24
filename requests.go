package requests

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func Get(url string, timeout time.Duration) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	client := &http.Client{}
	if timeout != 0 {
		client.Timeout = timeout
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("remote server occer error and http code is %d", resp.StatusCode)
	}
	response, _ := ioutil.ReadAll(resp.Body)
	return response, nil
}

func GetInsecure(url string, timeout time.Duration) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	//跳过证书验证
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{
		Transport: tr,
	}
	if timeout != 0 {
		client.Timeout = timeout
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("remote server occer error and http code is %d", resp.StatusCode)
	}
	response, _ := ioutil.ReadAll(resp.Body)
	return response, nil
}

func GetWithHeader(url string, header map[string]string, timeout time.Duration) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	for key, value := range header {
		req.Header.Set(key, value)
	}
	client := &http.Client{}
	if timeout != 0 {
		client.Timeout = timeout
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("remote server occer error and http code is %d", resp.StatusCode)
	}
	response, _ := ioutil.ReadAll(resp.Body)
	return response, nil
}

func GetInsecureWithHeader(url string, header map[string]string, timeout time.Duration) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	for key, value := range header {
		req.Header.Set(key, value)
	}
	//跳过证书验证
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{
		Transport: tr,
	}
	if timeout != 0 {
		client.Timeout = timeout
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("remote server occer error and http code is %d", resp.StatusCode)
	}
	response, _ := ioutil.ReadAll(resp.Body)
	return response, nil
}

// Post {timeout=time.Second * 3}
func Post(url string, body []byte, header map[string]string, timeout time.Duration) ([]byte, error) {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	for key, value := range header {
		req.Header.Set(key, value)
	}
	client := &http.Client{}
	if timeout != 0 {
		client.Timeout = timeout
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	response, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("remote server occer error and http code is %d, response: %v", resp.StatusCode, string(response))
	}
	//response, _ := ioutil.ReadAll(resp.Body)
	return response, nil
}

func PostInsecure(url string, body []byte, header map[string]string, timeout time.Duration) ([]byte, error) {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	for key, value := range header {
		req.Header.Set(key, value)
	}
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{
		Transport: tr,
	}
	if timeout != 0 {
		client.Timeout = timeout
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	response, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("remote server occer error and http code is %d, response: %v", resp.StatusCode, string(response))
	}
	//response, _ := ioutil.ReadAll(resp.Body)
	return response, nil
}

func Put(url string, body []byte, header map[string]string, timeout time.Duration) ([]byte, error) {
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	for key, value := range header {
		req.Header.Set(key, value)
	}
	client := &http.Client{}
	if timeout != 0 {
		client.Timeout = timeout
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	response, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("remote server occer error and http code is %d, response: %v", resp.StatusCode, string(response))
	}
	return response, nil
}

func Delete(url string, timeout time.Duration) ([]byte, error) {
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}
	client := &http.Client{}
	if timeout != 0 {
		client.Timeout = timeout
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("remote server occer error and http code is %d", resp.StatusCode)
	}
	response, _ := ioutil.ReadAll(resp.Body)
	return response, nil
}


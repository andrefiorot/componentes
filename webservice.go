package componentes

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	CAMINHO = "pedidos"
)

func BasicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

func Wenvio(post interface{}, endpoint string, identificador string, token string) ([]byte, error) {
	var data []byte
	client := &http.Client{}
	jsonValue, _ := json.Marshal(post)
	//	gravar_arquivo(jsonValue, caminho+"envio_"+identificador+".json")

	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(jsonValue))
	req.Header.Add("Authorization", token)

	req.Header.Add("Content-Type", "application/json; charset=UTF-8")
	response, err := client.Do(req)
	fmt.Println("Response ", identificador, response)
	defer response.Body.Close()

	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		fmt.Println("Non-OK HTTP status:", response.StatusCode)
		err = errors.New("Status Code Failed")
		//	data, err = ioutil.ReadAll(response.Body)
		return data, err

	}

	data, err = ioutil.ReadAll(response.Body)
	return data, err

}

// Realiza um Post Recebendo diretamenteo o []byte, ao inves do struct
func WPost(post []byte, endpoint string, identificador string, token string, login string, senha string) ([]byte, error) {
	var data []byte
	client := &http.Client{}

	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(post))

	// Se o token não estiver preenchido usa autenticação basic
	if len(token) > 0 {
		req.Header.Add("Authorization", token)
	} else {
		req.SetBasicAuth(login, senha)
	}

	req.Header.Add("Content-Type", "application/json; charset=UTF-8")

	response, err := client.Do(req)
	fmt.Println("Response ", identificador, response)
	defer response.Body.Close()

	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		fmt.Println("Non-OK HTTP status:", response.StatusCode)
		err = errors.New("Status Code Failed: " + response.Status)
		//		data, err = ioutil.ReadAll(response.Body)
		return data, err

	}

	data, err = ioutil.ReadAll(response.Body)
	return data, err

}

// Request de requisição POST, GET
func WRequest(post interface{}, metodo string, endpoint string, token string, login string, senha string, timeout time.Duration) ([]byte, error) {

	var data []byte
	client := &http.Client{Timeout: timeout}
	jsonValue, _ := json.Marshal(post)
	req, err := http.NewRequest(metodo, endpoint, bytes.NewBuffer(jsonValue))

	// Se o token não estiver preenchido usa autenticação basic
	if len(token) > 0 {
		req.Header.Add("Authorization", token)
	} else {
		req.SetBasicAuth(login, senha)
	}

	req.Header.Add("Content-Type", "application/json; charset=UTF-8")

	response, err := client.Do(req)
	defer response.Body.Close()
	if err != nil {
		fmt.Println("Erro ao realizar o request ", err)
		return nil, err

	}

	if response.StatusCode != http.StatusOK {
		fmt.Println("Non-OK HTTP status:", response.StatusCode)
		err = errors.New("Status Code Failed: " + response.Status)

		return nil, err

	}

	data, err = ioutil.ReadAll(response.Body)
	return data, err

}

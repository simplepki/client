package client

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/spf13/viper"
	"github.com/simplepki/client/config"
	"github.com/simplepki/client/tls"

)

func New() Client {
	config.Load()
	log.Print("account: ", viper.Get("account"))
	log.Print("id: ", viper.Get("id"))
	log.Print("chain: ", viper.Get("chain"))
	log.Print("endpoint: ", viper.Get("endpoint"))
	return Client{}
}
type Client struct {}

func (c Client) NewCertPair() error {
	cert := tls.NewCert(viper.GetString("account"), viper.GetString("chain"), viper.GetString("id"))
	//log.Println(string(cert.CSRRequest(viper.GetString("token"))))
	log.Println("Client sending CSR to: ", viper.GetString("endpoint"))
	certRequest, err := http.NewRequest("GET", viper.GetString("endpoint")+"/sign_csr", bytes.NewBuffer(cert.CSRRequest(viper.GetString("token"))))
	if err != nil {
		return err
	}
	certRequest.Header.Set("Content-Type", "application/json")
	httpClient := newHTTPClient()
	log.Println("client sending CSR")
	response, err := httpClient.Do(certRequest)
	if err != nil {
		log.Println("Client.Do error: ", err.Error())
		return err
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println("Client response error: ", err.Error())
		return nil
	}

	log.Println(string(body))

	return nil
}

func NewTLSConfig(){}


func newHTTPClient() *http.Client {
	c := &http.Client{
		Timeout: 1 * time.Minute,
	}

	return c
}


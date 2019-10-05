package tls

import (
	"bytes"
	"crypto/x509/pkix"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/simplepki/core/keypair"
)

type Certificate struct {
	Id      string
	Path    string
	Url     string
	KeyPair keypair.KeyPair
}

type jsonCSR struct {
	Id   string `json:"id"`
	Path string `json:"path"`
	CSR  string `json:"csr"`
}

type jsonSignedCert struct{}

func NewCert(path, id, url string) *Certificate {
	//only in memory at the moment
	kp := keypair.NewKeyPair("memory")

	newCert := &Certificate{
		Id:      id,
		Path:    path,
		Url:     url,
		KeyPair: kp,
	}

	newCert.sendCSR()

	return newCert
}

func (c *Certificate) base64EncodedCSR() string {
	var spiffePath string

	if strings.Contains(c.Path, "spiffe://") {
		spiffePath = c.Path
	} else {
		spiffePath = fmt.Sprintf("spiffe://%s", c.Path)
	}

	pkixName := pkix.Name{
		CommonName: fmt.Sprintf("%s/%s", spiffePath, c.Id),
	}

	csr := c.KeyPair.CreateCSR(pkixName, []string{})
	log.Println("got csr: ", csr.Raw)
	log.Printf("ecoding certificate of length: %v\n", len(csr.Raw))
	b64KP := base64.StdEncoding.EncodeToString(csr.Raw)

	return b64KP
}

func (c *Certificate) toJson() []byte {
	jsonStruct := jsonCSR{
		Id:   c.Id,
		Path: c.Path,
		CSR:  c.base64EncodedCSR(),
	}

	jsonBytes, err := json.Marshal(jsonStruct)
	if err != nil {
		log.Fatal(err)
	}

	return jsonBytes
}

func (c *Certificate) sendCSR() {
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/csr", c.Url), bytes.NewBuffer(c.toJson()))
	if err != nil {
		log.Fatal(err.Error())
	}

	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err.Error())
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Printf("recieved response: %#v\n", string(body))
}

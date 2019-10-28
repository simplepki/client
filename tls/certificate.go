package tls

import (
	//"bytes"
	"crypto/x509/pkix"
	"encoding/base64"
	"encoding/json"
	"fmt"
	//"io/ioutil"
	"log"
	//"net/http"
	"strings"

	"github.com/simplepki/core/keypair"
)

type Certificate struct {
	Id           string
	Intermediate string
	Account string
	KeyPair      keypair.KeyPair
}

type jsonCSR struct {
	InterName string `json:"intermediate_name"`
	CertName  string `json:"cert_name"`
	Account   string `json:"account"`
	CSR       string `json:"csr"`
}

type jsonSignedCert struct{}

func NewCert(account, intermediate, id string) *Certificate {
	//only in memory at the moment
	kp := keypair.NewKeyPair("memory")

	var intermediateString string
	if strings.Contains(intermediate, "spiffe://") {
		intermediateString = intermediate
	} else {
		intermediateString = fmt.Sprintf("spiffe://%s", intermediate)
	}

	newCert := &Certificate{
		Account: account,
		Id:           id,
		Intermediate: intermediateString,
		KeyPair:      kp,
	}

	return newCert
}

func (c *Certificate) base64EncodedCSR() string {

	pkixName := pkix.Name{
		CommonName: fmt.Sprintf("%s/%s", c.Intermediate, c.Id),
	}

	csr := c.KeyPair.CreateCSR(pkixName, []string{})
	log.Println("got csr: ", csr.Raw)
	log.Printf("ecoding certificate of length: %v\n", len(csr.Raw))
	b64KP := base64.StdEncoding.EncodeToString(csr.Raw)

	return b64KP
}

func (c *Certificate) Json() []byte {
	jsonStruct := jsonCSR{
		CertName:   c.Id,
		InterName: c.Intermediate,
		Account: c.Account,
		CSR:  c.base64EncodedCSR(),
	}

	jsonBytes, err := json.Marshal(jsonStruct)
	if err != nil {
		log.Fatal(err)
	}

	return jsonBytes
}

/*func sendCSR(json ) {
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
}*/

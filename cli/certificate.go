package main

import (
	"log"

	"github.com/simplepki/client/tls"
	"github.com/spf13/cobra"
)

var intermediate string
var id string
var url string
var account string

func init() {
	rootCmd.AddCommand(certCmd)
	certCmd.AddCommand(newCertCmd)
	newCertCmd.PersistentFlags().StringVarP(&intermediate, "intermediate-ca", "c", "test-ca/test-inter", "path to request certificate from")
	newCertCmd.PersistentFlags().StringVarP(&id, "id", "i", "test-client", "id of certificate to request")
	newCertCmd.PersistentFlags().StringVarP(&url, "url", "u", "", "url to send request to")
	newCertCmd.PersistentFlags().StringVarP(&account, "acount", "a", "test", "account to add certificate to")
}

var certCmd = &cobra.Command{
	Use:   "cert",
	Short: "new, renew, close out, and list certs",
}

var newCertCmd = &cobra.Command{
	Use:   "new",
	Short: "generate new certificate",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("new cert cmd")
		cert := tls.NewCert(account, intermediate, id)
		log.Printf("cert: %#v\n", cert)
		log.Println("json:\n", string(cert.Json()))

	},
}

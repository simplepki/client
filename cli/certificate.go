package main

import (
	"log"

	//"github.com/simplepki/client/tls"
	"github.com/simplepki/client"
	_ "github.com/simplepki/client/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	rootCmd.AddCommand(certCmd)
	certCmd.AddCommand(newCertCmd)
	newCertCmd.PersistentFlags().StringP("intermediate-chain", "c", "test-ca/test-inter", "path to request certificate from")
	newCertCmd.PersistentFlags().StringP("id", "i", "test-client", "id of certificate to request")
	newCertCmd.PersistentFlags().StringP("endpoint", "e", "", "url to send request to")
	newCertCmd.PersistentFlags().StringP("account", "a", "test", "account to add certificate to")
	newCertCmd.PersistentFlags().StringP("token", "t", "", "token for authn/z to simple pki service")
	viper.BindPFlag("account", newCertCmd.PersistentFlags().Lookup("account"))
	viper.BindPFlag("id", newCertCmd.PersistentFlags().Lookup("id"))
	viper.BindPFlag("chain", newCertCmd.PersistentFlags().Lookup("intermediate-chain"))
	viper.BindPFlag("endpoint", newCertCmd.PersistentFlags().Lookup("endpoint"))
	viper.BindPFlag("token", newCertCmd.PersistentFlags().Lookup("token"))
}

var certCmd = &cobra.Command{
	Use:   "cert",
	Short: "new, renew, close out, and list certs",
}

var newCertCmd = &cobra.Command{
	Use:   "new",
	Short: "generate new certificate",
	Run: func(cmd *cobra.Command, args []string) {
		c := client.New()
		err := c.NewCertPair()
		if err != nil {
			log.Println(err.Error)
		}
		//log.Println("new cert cmd")
		//cert := tls.NewCert(account, intermediate, id)
		//log.Println("json:\n", string(cert.Json()))
	},
}

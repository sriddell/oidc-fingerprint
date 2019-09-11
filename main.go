package main

import (
	"fmt"
	"crypto/tls"
	"crypto/sha1"
	"os"
	"net/url"
)

// getIssuerCAThumbprint obtains thumbprint of root CA by connecting to the
// OIDC issuer and parsing certificates
func main() {
	config := &tls.Config{InsecureSkipVerify: false}
	u, err := url.Parse(os.Args[1])
	if err != nil {
		fmt.Println(err)
	}
	conn, err := tls.Dial("tcp", u.Host + ":443", config)
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()

	cs := conn.ConnectionState()
	if numCerts := len(cs.PeerCertificates); numCerts >= 1 {
		root := cs.PeerCertificates[numCerts-1]
		thumbprint := fmt.Sprintf("{\"thumbprint\":\"%x\"}", sha1.Sum(root.Raw))
		fmt.Println(thumbprint)
	}
}

package domaincheck

// This package contains the domain check related functions.

import (
	"crypto/tls"
	"time"

	"github.com/akosgarai/projectregister/pkg/model"
)

// HasSSL checks if the domain has SSL certificate.
func HasSSL(domain *model.Domain) bool {
	// retry the check if the domain hasn't got SSL, to prevent the false negative result
	maxAttempts := 3
	attempt := 0
	for attempt < maxAttempts {
		if hasSSL(domain.Name) {
			return true
		}
		attempt++
		// wait attempt*2 seconds before the next try
		time.Sleep(time.Duration(attempt*2) * time.Second)
	}
	return false
}

// hasSSL checks if the domain has SSL certificate.
func hasSSL(domain string) bool {
	// Check if the domain has an SSL certificate
	conn, err := tls.Dial("tcp", domain+":443", nil)
	if err != nil {
		return false
	}
	defer conn.Close()
	// Check whether the SSL certificate and the hostname match
	err = conn.VerifyHostname(domain)
	if err != nil {
		return false
	}
	// check the expiration date of the certificate
	certs := conn.ConnectionState().PeerCertificates
	now := time.Now()
	expirationDate := certs[0].NotAfter
	if now.After(expirationDate) {
		return false
	}

	return true
}

/*
Copyright © 2026 Vic Fernandez III <@cyberphor>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package emass

import (
	"crypto/tls"
	"net/http"

	"github.com/deathlabs/emu/models"
)

func getTLSCertificates(publicKeyPath string, privateKeyPath string) ([]tls.Certificate, error) {
	var (
		certificate  tls.Certificate
		certificates []tls.Certificate
		err          error
	)

	certificate, err = tls.LoadX509KeyPair(publicKeyPath, privateKeyPath)
	if err != nil {
		return nil, err
	}

	certificates = []tls.Certificate{
		certificate,
	}

	return certificates, nil
}

func getTLSConfig(publicKeyPath string, privateKeyPath string) (*tls.Config, error) {
	var (
		certificates []tls.Certificate
		config       *tls.Config
		err          error
	)

	certificates, err = getTLSCertificates(publicKeyPath, privateKeyPath)
	if err != nil {
		return nil, err
	}

	config = &tls.Config{
		Certificates: certificates,
		MinVersion:   tls.VersionTLS12,
	}

	return config, nil
}

func getTransport(publicKeyPath string, privateKeyPath string) (*http.Transport, error) {
	var (
		config    *tls.Config
		err       error
		transport *http.Transport
	)

	config, err = getTLSConfig(publicKeyPath, privateKeyPath)
	if err != nil {
		return nil, err
	}

	transport = &http.Transport{
		TLSClientConfig: config,
	}

	return transport, nil
}

func GetClient(profile models.ConfigProfile) (*http.Client, error) {
	var (
		client    *http.Client
		err       error
		transport *http.Transport
	)

	transport, err = getTransport(profile.PublicKeyPath, profile.PrivateKeyPath)
	if err != nil {
		return nil, err
	}

	client = &http.Client{
		Transport: transport,
	}

	return client, nil
}

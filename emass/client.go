package emass

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/deathlabs/emu/v4/config"
	"github.com/deathlabs/emu/v4/models"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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
		Certificates:  certificates,
		MinVersion:    tls.VersionTLS12,
		Renegotiation: tls.RenegotiateOnceAsClient,
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

func getHTTPClient(profile models.ConfigProfile) (*http.Client, error) {
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

func SetupClient(cmd *cobra.Command, args []string) error {
	var err error

	// Check output format specified.
	if config.OutputFormat != "json" && config.OutputFormat != "yaml" {
		return fmt.Errorf("invalid output format \"%s\"", config.OutputFormat)
	}

	// Set the config filepath to the default if one is not provided.
	if config.Filename != config.DefaultConfigFilePath {
		viper.SetConfigType("yaml")
		viper.SetConfigFile(config.Filename)
	} else {
		viper.SetConfigFile(filepath.Join(".", config.DefaultConfigFilePath))
	}

	// Copy the config into memory.
	err = viper.ReadInConfig()
	if err != nil {
		return err
	}

	// Unmarshal the config into a Config struct.
	err = viper.Unmarshal(&config.Data)
	if err != nil {
		return err
	}

	// Resolve the profiles for each system in the config.
	config.Data.ResolveProfilesToSystems()

	return nil
}

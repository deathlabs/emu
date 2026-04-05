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
package models

import (
	"errors"
	"fmt"
)

type Profile struct {
	Name           string `json:"name" yaml:"name"`
	APIKey         string `json:"apiKey" yaml:"apiKey"`
	PublicKeyPath  string `json:"publicKeyPath" yaml:"publicKeyPath"`
	PrivateKeyPath string `json:"privateKeyPath" yaml:"privateKeyPath"`
}

type System struct {
	Name        string  `json:"name" yaml:"name"`
	ID          int     `json:"id" yaml:"id"`
	ProfileName string  `json:"profileName" yaml:"profileName"`
	Profile     Profile `json:"profile" yaml:"profile"`
}

type Config struct {
	URL      string    `json:"url" yaml:"url"`
	Profiles []Profile `json:"profiles" yaml:"profiles"`
	Systems  []System  `json:"systems" yaml:"systems"`
	Settings struct {
		Output struct {
			Format string `json:"format" yaml:"format"`
		} `json:"output" yaml:"output"`
	} `json:"settings" yaml:"settings"`
}

func (config *Config) Validate() error {
	if config.URL == "" {
		return errors.New("a URL to eMASS was not provided")
	}
	return nil
}

func (config *Config) FindProfile(name string) (*Profile, error) {
	for i := range config.Profiles {
		if config.Profiles[i].Name == name {
			return &config.Profiles[i], nil
		}
	}
	return nil, fmt.Errorf("profile %q not found", name)
}

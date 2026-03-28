package types

type Credential struct {
	APIKey      string `json:"api_key"`
	Certifcates struct {
		PublicKeyPath  string `json:"public_key_path"`
		PrivateKeyPath string `json:"private_key_path"`
	} `json:"certificates"`
}

type System struct {
	Credential Credential `json:"credential"`
}

type Config struct {
	EMASS struct {
		URL         string                `json:"url"`
		Credentials map[string]Credential `json:"credentials"`
		Systems     map[string]System     `json:"systems"`
	} `json:"emass"`
	EMU struct {
		Output struct {
			Format string `json:"format"`
		} `json:"output"`
	} `json:"emu"`
}

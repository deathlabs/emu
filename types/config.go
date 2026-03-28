package types

type Config struct {
	EMASS struct {
		User struct {
			UID int `json:"uid"`
		} `json:"user"`
		API struct {
			URL string `json:"url"`
		} `json:"api"`
	} `json:"emass"`
}

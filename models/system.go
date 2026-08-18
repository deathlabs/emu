package models

type System struct {
	ID            int           `mapstructure:"id" json:"id" yaml:"id"`
	Name          string        `mapstructure:"name" json:"name" yaml:"name"`
	ConfigProfile ConfigProfile `mapstructure:"-" json:"profile" yaml:"profile"`
}

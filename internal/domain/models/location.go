package models

type Location struct {
	Region string `yaml:"region" json:"region"`
	County string `yaml:"country" json:"country"`
	City   string `yaml:"city" json:"city"`
}

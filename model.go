package main

type ConfigurationModel struct {
	Configuration []ConfigurationRecord `json:"configuration"`
}

type ConfigurationRecord struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

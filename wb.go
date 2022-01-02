package wb

import (
	"encoding/json"
	"io/ioutil"
)

type Bank struct {
	Name              string   `json:"name"`
	Address           string   `json:"address"`
	Businessname      string   `json:"business_name"`
	BranchCode        string   `json:"branch_code"`
	Code              string   `json:"code"`
	Description       string   `json:"description"`
	SwiftCode         string   `json:"swift_code"`
	StateProvince     string   `json:"state_province"`
	StateProvinceCode string   `json:"state_province_code"`
	City              string   `json:"city"`
	Website           string   `json:"website"`
	Branches          []string `json:"branches"`
}

func GetNigerianBanks() ([]Bank, error) {
	var banks []Bank
	banksJson, err := ioutil.ReadFile("./data/nigeria.json")
	if err != nil {
		return banks, err
	}

	if err := json.Unmarshal(banksJson, &banks); err != nil {
		return banks, err
	}

	return banks, nil

}

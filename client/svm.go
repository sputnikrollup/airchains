package client

import (
	"airchains/utils"
	"airchains/utils/svm"
	"fmt"
	"strings"

	"github.com/manifoldco/promptui"
)

func SVMInitialiazation() {
	data := map[string]interface{}{
		"daInfo":        svmDAType(),
		"sequencerInfo": svmSequencerType(),
	}
	utils.CreateFolderAndJSONFile(data, "config", "solana_config.json")
	daType := data["daInfo"].(map[string]interface{})["daSelected"].(string)
	daTypetoLower := strings.ToLower(daType)
	utils.ENVSetup("svm-sequencer-node", "8899", daTypetoLower)

}

func svmDAType() map[string]interface{} {
	var chainInfoKeyValue map[string]interface{}

	prompt := promptui.Select{
		Label: "Select your DA type : ",
		Items: []string{"Celestia", "Avail"},
		Templates: &promptui.SelectTemplates{
			Label:    "{{ . | cyan }}",
			Active:   "\U000025B6 {{ if eq . \"Eigen Layer\" }}{{ . | red | cyan }} (Coming Soon){{ else }}{{ . | green }}{{ end }}",
			Inactive: "  {{ . | black }}",
			Selected: "\U000025B6 {{ . | cyan | cyan }}",
		},
	}
	_, daTech, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return svmDAType()
	}

	fmt.Printf("You selected %s\n", daTech)
	utils.DaRPC()
	utils.SettlementRPC()

	chainInfoKeyValue = map[string]interface{}{
		"daSelected": daTech,
	}

	utils.DelayLoader()
	return chainInfoKeyValue
}

func svmSequencerType() map[string]interface{} {
	var sequencerKeyValue map[string]interface{}
	prompt := promptui.Select{
		Label: "Select your sequencer type : ",
		Items: []string{"Air Sequencer"},
		Templates: &promptui.SelectTemplates{
			Label:    "{{ . | cyan }}",
			Active:   "\U000025B6 {{ if eq . \"Espresso Sequencer\" }}{{ . | red | cyan }} (Coming Soon){{ else }}{{ . | green }}{{ end }}",
			Inactive: "  {{ . | black }}",
			Selected: "\U000025B6 {{ . | cyan | cyan }}",
		},
	}
	_, sequencerValue, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		svmSequencerType()
	}
	svm.SolanaSequencerClone()
	sequencerKeyValue = map[string]interface{}{
		"sequencerType": sequencerValue,
	}
	utils.DelayLoader()
	return sequencerKeyValue
}

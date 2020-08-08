package op

import "encoding/json"

func ListItems(vault string) ([]Item, error) {
	flags := make(map[string]string)

	if vault != "" {
		flags["vault"] = vault
	}

	out, err := executor.RunOp("op list items", flags)

	if err != nil {
		return nil, err
	}

	var items []Item
	errJson := json.Unmarshal([]byte(out), &items)

	if errJson != nil {
		return nil, errJson
	}

	return items, nil
}

func ListTemplates() ([]Template, error) {
	out, err := executor.RunOp("op list templates", nil)

	if err != nil {
		return nil, err
	}

	var templates []Template
	errJson := json.Unmarshal([]byte(out), &templates)

	if errJson != nil {
		return nil, errJson
	}

	return templates, nil
}

func ListVaults() ([]Vault, error) {
	out, err := executor.RunOp("op list vaults", nil)

	if err != nil {
		return nil, err
	}

	var vaults []Vault
	errJson := json.Unmarshal([]byte(out), &vaults)

	if errJson != nil {
		return nil, errJson
	}

	return vaults, nil
}

package op

import "encoding/json"

// ListItems : gets a list of items in a vault without details
func (o Op) ListItems(vault string) ([]Item, error) {
	flags := make(map[string]string)

	if vault != "" {
		flags["vault"] = vault
	}

	out, err := o.executor.RunOp("op list items", flags)

	if err != nil {
		return nil, err
	}

	var items []Item
	errJSON := json.Unmarshal([]byte(out), &items)

	if errJSON != nil {
		return nil, errJSON
	}

	return items, nil
}

// ListTemplates : gets a list of templates
func (o Op) ListTemplates() ([]Template, error) {
	out, err := o.executor.RunOp("op list templates", nil)

	if err != nil {
		return nil, err
	}

	var templates []Template
	errJSON := json.Unmarshal([]byte(out), &templates)

	if errJSON != nil {
		return nil, errJSON
	}

	return templates, nil
}

// ListVaults : gets a list of vaults
func (o Op) ListVaults() ([]Vault, error) {
	out, err := o.executor.RunOp("op list vaults", nil)

	if err != nil {
		return nil, err
	}

	var vaults []Vault
	errJSON := json.Unmarshal([]byte(out), &vaults)

	if errJSON != nil {
		return nil, errJSON
	}

	return vaults, nil
}

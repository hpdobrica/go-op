package op

import (
	"encoding/json"
	"fmt"
)

// GetItem : Gets an item with details
func (o Op) GetItem(itemID string) (ItemWithDetails, error) {

	out, err := o.executor.RunOp(fmt.Sprintf("op get item %s", itemID), nil)

	var item ItemWithDetails

	if err != nil {
		return item, err
	}

	errJSON := json.Unmarshal([]byte(out), &item)

	if errJSON != nil {
		return item, errJSON
	}

	return item, nil
}

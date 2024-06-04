package gologix_tests

import (
	"testing"

	"github.com/treavorj/gologix"
)

func TestSubList(t *testing.T) {

	tc := getTestConfig()
	client := gologix.NewClient(tc.PLC_Address)
	err := client.Connect()
	if err != nil {
		t.Error(err)
		return
	}
	defer func() {
		err := client.Disconnect()
		if err != nil {
			t.Errorf("problem disconnecting. %v", err)
		}
	}()

	_, err = client.ListSubTags("Program:gologix_tests", 1, nil)
	if err != nil {
		t.Error(err)
		return
	}

}

func compare_array_order(have, want []int) bool {
	if len(have) != len(want) {
		return false
	}

	for i := range have {
		if have[i] != want[i] {
			return false
		}
	}
	return true
}

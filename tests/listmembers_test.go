package gologix_tests

import (
	"testing"

	"github.com/treavorj/gologix"
)

func TestMembersList(t *testing.T) {

	tc := getTestConfig()
	client := gologix.NewClient(tc.PlcAddress)
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

	_, err = client.ListMembers(1658)
	if err != nil {
		t.Error(err)
	}

}

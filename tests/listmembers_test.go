package gologix_tests

import (
	"testing"

	"github.com/danomagnum/gologix"
)

func TestMembersList(t *testing.T) {
	t.Skip("controller specific test")

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

package main

import (
	"bytes"
	"log"
	"os"
	"time"

	"github.com/treavorj/gologix"
)

// Demo program for readng an INT tag named "TestInt" in the controller.
func main() {
	var err error

	// setup the client.  If you need a different path you'll have to set that.
	client := gologix.NewClient("192.168.2.244")
	client.Path = &bytes.Buffer{}

	// for example, to have a controller on slot 1 instead of 0 you could do this
	// client.Path, err = gologix.Serialize(gologix.CIPPort{PortNo: 1}, gologix.CIPAddress(1))
	// or this
	// client.Path, err = gologix.ParsePath("1,1")

	// connect using parameters in the client struct
	err = client.Connect()
	if err != nil {
		log.Printf("Error opening client. %v", err)
		os.Exit(1)
	}
	// setup a deffered disconnect.  If you don't disconnect you might have trouble reconnecting because
	// you won't have sent the close forward open.  You'll have to wait for the CIP connection to time out
	// if that happens (about a minute)
	defer client.Disconnect()

	// define a variable with a type that matches the tag you want to read.  In this case it is an INT so
	// int16 or uint16 will work.
	var tag1 int32
	// call the read function.
	// note that tag names are case insensitive.
	// also note that for atomic types and structs you need to use a pointer.
	// for slices you don't use a pointer.

	//err = client.Read("testint", &tag1)
	err = client.Read("inputs[0]", &tag1)
	if err != nil {
		log.Printf("error reading testint. %v", err)
		os.Exit(1)
	}

	t := time.Second * 287
	for {
		time.Sleep(t)
		//err = client.Read("testint", &tag1)
		err = client.Read("inputs[0]", &tag1)
		if err != nil {
			log.Printf("error reading testint after %v seconds: %v", t, err)
			break
		}
		// do whatever you want with the value
		log.Printf("after %v seconds tag1 has value %d", t, tag1)
		t += time.Second * 1

	}

}

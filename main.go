package main

import (
	"fmt"
	"os"
	// "reflect"
	"github.com/sc-app/pkg/storing"
	"github.com/sc-app/pkg/web"
	"github.com/sc-app/pkg/web/controllers"
	"github.com/sc-app/pkg/handles"
	"github.com/sc-app/internal"
	"github.com/sc-app/blockchain"
)

func main() {

	var myStorage storing.UserStorage
	myStorage.Init("session_test")
	sessionObj := &handles.Session {
		Handle: "11.test",
		Ip: "156.106.193.160",
		Port: 8011,
		SessionId: "he15rxhkp1561irkno8gpw9hb",
	}

	// // err := sessionObj.Init()
	// // if err != nil {
	// // 	panic(err)
	// // }

	// boostrapper := internal.Boostrapper { 	
	// 	ConfigFile	 :	"cmd/app/user_t.json",
	// 	Storage  :	&myStorage,
	// 	HandleSession : sessionObj,
	// }

	// boostrapper.Init()
	// fmt.Printf("%+v", sessionObj)

	// //fmt.Println(myStorage.Topics)
	// //myStorage.Print()
	
	// fmt.Printf("\n-----------------------------------------------\n")	
	// /* Move to a Handlession type method */
	// tmp := *sessionObj
	// fmt.Printf("Interface to Handle System ready.\n Prefix: %q, Ip: %q, Port: %d, sessionId: %q\n\n",
	// 	tmp.Handle, tmp.Ip, tmp.Port, tmp.SessionId)
	 
	fmt.Printf("\n-----------------------------------------------\n")	
	
	// Definition of the Fabric SDK properties
	fSetup := blockchain.FabricSetup{
		// Network parameters
		OrdererID: "orderer.hf.chainhero.io",

		// Channel parameters
		ChannelID:     "chainhero",
		ChannelConfig: os.Getenv("GOPATH") + "/src/github.com/sc-app/fixtures/artifacts/chainhero.channel.tx",

		// Chaincode parameters
		ChainCodeID:     "heroes-service",
		ChaincodeGoPath: os.Getenv("GOPATH"),
	//	ChaincodePath:   "github.com/sc-app/chaincode/",
		OrgAdmin:        "Admin",
		OrgName:         "org1",
		ConfigFile:      "config.yaml",

		// User parameters
		UserName: "User1",
	}
	
	boostrapper := internal.Boostrapper { 	
		ConfigFile	 :	"cmd/app/user_t.json",
		Storage  :	&myStorage,
		HandleSession : sessionObj,
		Fabric	: &fSetup,
	}

	boostrapper.Init()
	myStorage.Print()
	fmt.Println(myStorage.Topics, "\n")

	err := fSetup.Initialize()
	if err != nil {
		fmt.Printf("Unable to initialize the Fabric SDK: %v\n", err)
		return
	}
	// Close SDK
	defer fSetup.CloseSDK()


	// Install and instantiate the chaincode
	err = fSetup.InstallAndInstantiateCC(myStorage.UserBytes())

	if err != nil {
		fmt.Printf("Unable to install and instantiate the chaincode: %v\n", err)
		return
	}

	appServ := &controllers.Application {
		&myStorage,
		sessionObj,
		&fSetup,
	}
	
}
	// fmt.Printf("%+v\n", myStorage)














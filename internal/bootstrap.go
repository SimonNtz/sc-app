package internal

import (
	"fmt"
	"encoding/json"
	"github.com/sc-app/pkg/handles"
	"github.com/sc-app/pkg/users"
	"github.com/sc-app/pkg/topics"
	"errors"
	"github.com/sc-app/pkg/storing"
	"os"
	"io/ioutil"
	"github.com/mitchellh/mapstructure"
	"github.com/sc-app/blockchain"
	"github.com/Gohandler/pkg/handle/helper"
	"encoding/base64"
)

type Boostrapper struct {
	ConfigFile	string
	ConfigData map[string]interface{}
	Storage		*storing.UserStorage
	HandleSession *handles.Session
	Fabric        *blockchain.FabricSetup
}

func (bs *Boostrapper) Init() {
	bs.ConfigData = make(map[string]interface{})
	bs.ConfigData = getJSON(bs.ConfigFile)
	/* USE METHODS */
	bs.HandleSession.Handle = bs.ConfigData["prefix"].(string)
	bs.HandleSession.Key = bs.ConfigData["keypriv"].(string)
	bs.HandleSession.Keypub = bs.ConfigData["keypub"].(string)

	bs.Users()
	bs.Topics()
	bs.ChainCode()
}

func (bs *Boostrapper) Users() {
	userList := bs.ConfigData["users"].([]interface{})
	if len(userList) == 1 {
		hdlRep := bs.HandleSession.GetHandle(userList[0].(map[string]string)["id"])
		vlistVal, err := getVLIST(hdlRep["values"].([]interface{}))
		if err != nil {
			panic(err)
		}

		for _, v := range vlistVal {
			vmap := v.(map[string]interface{})
			u := users.User{
				Id: vmap["handle"].(string),
				Consents: make(map[string][]string),
			}
			bs.Storage.Add(&u)
		}
	} else {
		for _, v := range userList{
			u := users.User{
				Id: v.(map[string]interface{})["id"].(string),
				Consents: make(map[string][]string),
			}
			bs.Storage.Add(&u)
		}
	}
}

func (bs *Boostrapper) Topics() {
	var topicsValues []interface{}
	var topicIdx []string
	topicList := bs.ConfigData["topics"].([]interface{})

	if len(topicList) == 1 {
		topicsRep 	  := bs.HandleSession.GetHandle(topicList[0].(map[string]interface{})["id"].(string))
		topicsId, err := getVLIST(topicsRep["values"].([]interface{}))
		if err == nil {
			for _, v := range topicsId {
				vmap := v.(map[string]interface{})
				tmp := bs.HandleSession.GetHandle(fmt.Sprintf("%s?index=%d", vmap["handle"], int(vmap["index"].(float64))))
				topicIdx = append(topicIdx, vmap["handle"].(string))
				topicsValues = append(topicsValues, tmp["values"].([]interface{})[0])
			}
		} else {
			topicsValues = topicsRep["values"].([]interface{})
		}
	}

	for k, v := range topicsValues {
		tmpL1 := v.(map[string] interface{})
		tmpL2 := tmpL1["data"].(map[string] interface{})
		tmpL3 := tmpL2["value"].(string)

		valueMap := make(map[string]string)		
		err := json.Unmarshal([]byte(tmpL3), &valueMap)

		if err != nil {
				panic(err)
		}
		var hdl string
		idx := int(tmpL1["index"].(float64))
		if len(topicIdx) > 0 {
			hdl = topicIdx[k]
		} else {
			hdl = topicList[0].(map[string]string)["id"]
		}

		var opt topics.Option
		err = mapstructure.Decode(bs.ConfigData["options"], &opt)
		if err != nil {
			panic(err)
		}

		bs.Storage.AddTopic(topics.Topic {
			Id: fmt.Sprintf("%d:%s", idx, hdl),
			// Id: int(tmpL1["index"].(float64)),
			Lead : valueMap["title"],
			Corpus: valueMap["text"],
			OptsModel : opt,
		})
	}
} 

func getVLIST(hdlValues []interface{}) ([]interface{}, error) {
	var err error
	var tmp []interface{}
	for _, v := range hdlValues{
		rec := v.(map[string] interface{})
		if hdltype:= rec["type"].(string); hdltype == "HS_VLIST"{
			tmp = rec["data"].(map[string]interface{})["value"].([]interface{})
		} else { 
			err =  errors.New("NO HS_VLIST FOUND IN CONFIG FILE.")
		} 
	}
	return tmp, err	
}

func getJSON(jsonSRC string) map[string]interface{}{
	jsonFile, err := os.Open(jsonSRC)
	
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error from reading: %s\n", err)
	}
	
	var hdlList map[string] interface{}
	
	defer jsonFile.Close()
	
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error from reading: %s\n", err)
	}
	
	err = json.Unmarshal(byteValue, &hdlList)
	if err != nil {
		panic(err)
	}
	return hdlList
}

/*Move cc path to config file*/
func (bs *Boostrapper)ChainCode(){
	hdl := bs.HandleSession.GetHandle(bs.ConfigData["chaincode"].(string))
	sign :=  helper.GetHandleSign(hdl) 
	if ! helper.VerifyHandle(foo , bs.HandleSession.Keypub) {
		panic("Chaincode signature check failed")
	}
	cc64 := helper.GetHandleType(hdl, "dlt.cc")
	cc, err := base64.StdEncoding.DecodeString(cc64["Value"].(string))
	if err != nil {
		panic(err)
	}
	bs.Fabric.ChaincodeCode = cc
} 
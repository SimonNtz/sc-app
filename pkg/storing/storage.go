package storing

import (
	um "github.com/sc-app/pkg/users"
	tm "github.com/sc-app/pkg/topics"
	"fmt"
	"encoding/json"
	"reflect"
)

type UserStorage struct {
	Id       		string
	UserRepo 		map[string]*um.User //map[string] um.User
	Topics	     	[]tm.Topic
}


// SetName receives a pointer to Foo so it can modify it. 

func (us *UserStorage) Add(u *um.User) {
	us.UserRepo[u.Id] = u
	// fmt.Printf("%+v\n", *us)
}

func (us *UserStorage) AddSignature(id string, sign string) {
	us.UserRepo[id].Signature = sign
}

func (us *UserStorage) AddConsent(id string, consents map[string][]string) {
	us.UserRepo[id].Consents = consents
}

func (us *UserStorage) AddTopic(topic tm.Topic) {
	us.Topics = append(us.Topics, topic)
	// fmt.Printf("%+v\n", *us)
}


func (us UserStorage) Get(id string) um.User {
	if usr := *us.UserRepo[id]; usr.Id != "" {
		return usr
	} else {
		fmt.Printf("User with ID: %s does not exists.\n", id)
		return um.User{}
	}
}

func (us *UserStorage) Init(id string) {
	us.Id = id
	us.UserRepo = make(map[string] *um.User)
	// us.Topics = []Topic{Topic{Id: 555, Lead: "bgh", Corpus: "rrl"}}
}

func (us UserStorage) Print() {
	fmt.Printf("Users registered for: %s \n \n", us.Id)
	for _, val := range us.UserRepo {
		b, err := json.MarshalIndent(val, "", "  ")
		if err == nil {
				fmt.Println(string(b))
		}
		// fmt.Printf("PID: %s, CONSENTS: %v\n SIGNATURE: %s\n", usrVal.Id, usrVal.Consents, usrVal.Signature)
	}
}

func (us UserStorage) UserBytes() [][]byte{
	val := reflect.ValueOf(us.UserRepo).MapKeys()
	
	var usrBytes [][]byte
    for _, v := range(val) {
        usrBytes = append(usrBytes, []byte(v.String()))
	}
	return usrBytes
}


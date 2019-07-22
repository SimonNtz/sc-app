package users

import (
	// "encoding/json"
)

type User struct {
	Id			 	string `json:"pid"`
	Signature		string `json:"signature"`
	Consents       	map[string][]string `json:"consents"`
}

// func (u User) Consent {
// 	// u.send(Subject)
// }

// func (u User) Sign string {
// 	// sign or hash lets see if the get call is not too expensive user info
// }

// func (u User) IsSigned string {
// 	// CHECK IF USER.SIGNATURE not nil
// }
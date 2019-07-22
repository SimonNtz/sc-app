package handles


import (
	"fmt"
	// "encoding/json"
	"github.com/Gohandler/pkg/handle/helper"
	"github.com/Gohandler/pkg/handle/util"
	"errors"
	"strconv"
)

type Session struct {
	Handle		string
	Ip			string
	Port		int
	SessionId 	string
	Key			string
	KeyPub 		string
}


/* Update from a single entry to an array as an input*/

func FillHandle(data map[string] string) []byte {
	var hdl_json string
/*Use Gohandle type*/
	switch index := data["index"]; index {
	case "120":
		hdl_json = fmt.Sprintf(`[{"index": 120, "type": %s,"data": {"value":%q, "format":"string"}},
							   {"index": 130, "type": "STATUS", "data": {"value":"authentified", "format":"string"}}]`,
								data["type"], data["value"])

	default:
		index, _ := strconv.Atoi(data["index"]) 
		hdl_json = fmt.Sprintf(`[{"index": %d, "type": %q, "data": {"value":%q, "format":%q}}]`,
								 index, data["type"], data["value"], data["format"])
	}	
	
	return []byte(hdl_json)
}

func (session Session) SignHandle(handle string) string{
	rep := util.ResolveLHS(handle, session.Ip, session.Port)
	sign, _ := helper.SignHandle("11.test", handle, rep, session.Key)
	hdl := map[string]string{ 
		"index": "400",  
		"type": "HS_SIGNATURE",
		"value": string(sign),
		"format": "string",
	}
	url := fmt.Sprintf("%s?index=various", handle)
	util.PutLHS(url, FillHandle(hdl), session.Ip, session.Port, session.SessionId)
	return string(sign)
	
	// hdl2 := map[string]string{ 
	// 	"index": "401",  
	// 	"type": "10320/sig.digest",
	// 	"value": string(digests),
	// }
}

// func VerifyHandle(){}

func (session Session) PostHandle(handle string, data map[string]string) {
	util.PutLHS(handle, FillHandle(data), session.Ip, session.Port, session.SessionId)
}

/*Verify reponseCode here or in utils*/
func (session Session) GetHandle(handle string) map[string]interface{} {
	rep := util.ResolveLHS(handle , session.Ip ,  session.Port)
	return rep
}

func (session *Session) Init() error {
   srv, sess :=  helper.GetSessionId(session.Handle)
   cnonce, sign := helper.ChallengeSite(session.Key, sess)
   if ! util.PostAuthLHS(sess.Id, cnonce, sign) {
	  return errors.New("Authentication to handle system failed.")
   }

   session.Ip 		 = srv.Address
   session.Port 	 = srv.Interfaces.Port
   session.SessionId = sess.Id
   
   return nil
// sess := model.Session{Id:"ftjwcwc6qsbg1mbko146l02bu", Nonce: "6Gd/ry4WCw5+YOGh6tkySA=="}
//"C:/Users/noetzlin/Documents/DOA/GO/admpriv.pem"
//    cnonce, sign := helper.ChallengeSite(key, sess)
//    util.PutLHS(fmt.Sprintf("11.test/EVENTXXXX/USERS/%s", id), hdl, site, sess.Id)
}



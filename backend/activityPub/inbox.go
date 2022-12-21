package activityPub

import (
	"net/http"

	"github.com/gorilla/mux"
)

// /{username}/inbox
// /{username}/inbox/
func Inbox(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username, ok := vars["username"]
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
	}
	_ = username
	// pk, err := utils.ReadKeyFromFile("privatekey.pem")
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }
	// fmt.Println(pk)
	// key := utils.GenerateKey("")
	// color.HiMagenta(fmt.Sprintf("%s", *key))
	// keyBytes := utils.MarshalPublicKey(key)
	// color.HiMagenta(string(keyBytes))
	// return

	// msg := "-----BEGIN PUBLIC KEY-----\nMIICIjANBgkqhkiG9w0BAQEFAAOCAg8AMIICCgKCAgEAvHKTk7Nm8TFF2RUX+MwQ\n69hPKqcKE4JcRVx8eJW5ApSlQs4pu58wQO7yD+QvYIb4IKKOajQhKaaf3vCR0SWm\nwHLVtFn2Iwzm22VFR897eIKCN8DFGT4Nkq3/BiX8ZXb4NbtX5hry/Bj7b2MbDgfX\nptcoafxjofHUkqyZ5VcwgT+ctPjqAaiAHYmnxq4HPoxX5quuvNAe8Gl2Ij+ErE5v\n8Z/jXTMkOQJ0TBPCCSB/PTF0SSPRlwbyZVxXx/dTS/P5ewBO5KPJk9Ii/TDXSDFk\nJKmXxVd7kUIszAQ4qvg97mIi/L8TQf4g6D4Cj0EpbtRZnpv59vvoN7AcrgmKo0CI\n6J08wen9pGpiBp/gXs83mx6YY63SIqLlWEGpEgRZnrpe4nW24N0G1puiTSAiG3o2\n+loXh3J7b4fbcC/UR0pxLLtnhcYXhxUEyfb0d8DV8KjC4GlX9eiDmx0P5fSsUPdc\nOfHeKCKMLnI1FYbTEn+Ed7dYG4RwkXgzJPpigwHIWO2oqaGqWGt2UeOTr+HzbTqW\nOKw+tk1M/3X451DjWTmoGzLPDdoakB8pFyQ8YrrUcpUpLn9n4Susr2deek3UwfgK\nl23OwSttCkw4vs6+p063X3g0NwhWvZYDS/hJDJNbs33LHErfsiUidAmz8yEMsY0X\nKtdeL3dq3dvXYKwyHQ59zJcCAwEAAQ==\n-----END PUBLIC KEY-----\n"
	// block, _ := pem.Decode([]byte(msg))
	// color.HiMagenta(string(block.Type))
	// // color.HiMagenta(string(block.Bytes))
	// publicKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	// if err != nil {
	// 	color.Red(err.Error())
	// }
	// color.HiMagenta(fmt.Sprintf("%s", publicKey))
	// // color.HiMagenta(fmt.Sprintf("%T", publicKey)) // *rsa.PublicKey
}

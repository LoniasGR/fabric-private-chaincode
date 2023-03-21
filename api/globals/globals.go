package globals

import "github.com/hyperledger/fabric-private-chaincode/api/pkg"

var Config *pkg.Config
var Secret = []byte("secret")

const Userkey = "user"
const AppName = "fabric-private-chaincode"

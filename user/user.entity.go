package user

type User struct {
	Username         string `json:"username"`
	UserID           int
	Address          string
	Secret           string
	ReceivingAddress string
}

type Secret struct {
	Secret string `json:"secret"`
}

type MnemonicResponse struct {
	Secret string `json:"secret"`
	Error  string `json:"error"`
}

type Address struct {
	Success bool    `json:"success"`
	Account Account `json:"account"`
}

type Account struct {
	Address              string `json:"address"`
	UnconfirmedBalance   string `json:"unconfirmedBalance"`
	Balance              string `json:"balance"`
	PublicKey            string `json:"publicKey"`
	UnconfirmedSignature int    `json:"unconfirmedSignature"`
	SecondSignature      int    `json:"secondSignature"`
	SecondPublicKey      string `json:"secondPublicKey"`
	MultiSignatures      string `json:"multisignatures"`
	UMultiSignatures     string `json:"u_multisignatures"`
}

func GetUser(username string) User {
	var user User
	user.Username = username
	return user
}

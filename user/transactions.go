package user

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type SendTransactionRequest struct {
	Secret      string `json:"secret"`      // "Secret key of account",
	Amount      int    `json:"amount"`      // Amount of transaction * 10^8. Example: to send 1.1234 LISK, use 112340000 as amount */,
	RecipientID string `json:"recipientId"` // "Recipient of transaction. Address or username.",
}

type SendTransactionResponse struct {
	Secret       string `json:"secret"`       // "Secret key of account",
	Amount       string `json:"amount"`       // Amount of transaction * 10^8. Example: to send 1.1234 LISK, use 112340000 as amount */,
	RecipientID  string `json:"recipientId"`  // "Recipient of transaction. Address or username.",
	PublicKey    string `json:"publicKey"`    // "Public key of sender account, to verify secret passphrase in wallet. Optional, only for UI",
	SecondSecret string `json:"secondSecret"` // "Secret key from second transaction, required if user uses second signature"
}

func (u *User) SendLisk(amount string, toUsername string) error {
	sendingUser := User{Username: toUsername}
	if !sendingUser.UserExists() {
		return errors.New("Error getting user")
	}
	amountInt, _ := strconv.Atoi(amount)

	sendingUser.GetUserData()
	request := SendTransactionRequest{
		Secret:      u.Secret,
		Amount:      amountInt,
		RecipientID: u.Address,
	}

	var resp SendTransactionResponse
	transactionURL := "https://testnet.lisk.io/api/transactions"
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(request)

	client := &http.Client{}
	req, err := http.NewRequest("PUT", transactionURL, b)
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	if err != nil {
		log.Fatal(err)
	}

	response, err := client.Do(req)
	if err != nil {
		// handle error
		log.Fatal(err)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	responseAsString := string(responseData)
	err = json.Unmarshal([]byte(responseAsString), &resp)
	if err != nil {
		return errors.New("Error sending transaction")
	}
	fmt.Println(responseAsString)
	return nil
}

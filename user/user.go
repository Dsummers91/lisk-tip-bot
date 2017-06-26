package user

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"bytes"

	_ "github.com/lib/pq" //Postgres
)

func GenerateAddress() (Address, MnemonicResponse, error) {
	var address Address
	var mnemonicResponse MnemonicResponse
	mnemonicURL := "http://127.0.0.1:7200/account/new"

	response, err := http.Get(mnemonicURL)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	responseAsString := string(responseData)

	err = json.Unmarshal([]byte(responseAsString), &mnemonicResponse)
	if err != nil {
		return address, mnemonicResponse, err
	}

	accountURL := "https://testnet.lisk.io/api/accounts/open"
	secret := Secret{Secret: mnemonicResponse.Secret}
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(secret)
	response, err = http.Post(accountURL, "application/json; charset=utf-8", b)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	responseData, err = ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	responseAsString = string(responseData)
	fmt.Println(responseAsString)
	err = json.Unmarshal([]byte(responseAsString), &address)
	if err != nil {
		return address, mnemonicResponse, err
	}
	return address, mnemonicResponse, nil
}

func (u *User) GetUserData() {
	db, _ := sql.Open("postgres", "user=lisktiptest password=password dbname=postgres sslmode=disable")
	rows, _ := db.Query("SELECT username, address, secret, COALESCE(receiving_address, '')  FROM users_test WHERE username = $1", u.Username)
	for rows.Next() {
		err := rows.Scan(&u.Username, &u.Address, &u.Secret, &u.ReceivingAddress)
		if err != nil {
			log.Fatalf("Scan: %v", err)
		}
		return
	}
	return
}

func (u *User) UserExists() bool {
	db, _ := sql.Open("postgres", "user=lisktiptest password=password dbname=postgres sslmode=disable")
	rows, err := db.Query("SELECT username FROM users_test WHERE username = $1", u.Username)
	if err != nil {
		return u.CreateUser()
	}
	var r User
	for rows.Next() {
		err := rows.Scan(&r.Username)
		if err != nil {
			log.Fatalf("Scan: %v", err)
		}
		return true
	}
	return u.CreateUser()
}

func (u *User) CreateUser() bool {
	db, err := sql.Open("postgres", "user=lisktiptest password=password dbname=postgres sslmode=disable")
	account, mnemonic, _ := GenerateAddress()

	var redditUsername string
	err = db.QueryRow(`INSERT INTO users_test(username, address, secret)
		VALUES($1, $2, $3) RETURNING username`, u.Username, account.Account.Address, mnemonic.Secret).Scan(&redditUsername)

	if err != nil {
		log.Fatal(err)
	}
	return true
}

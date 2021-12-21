/*
SPDX-License-Identifier: Apache-2.0
*/

package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"crypto/sha256"
	"crypto"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// SmartContract provides functions for managing coinaccount
type SmartContract struct {
	contractapi.Contract
}

// Car describes basic details of coinaccount
type CoinAccount struct {
	Account   string `json:"account"`
	Balance   float64 `json:"balance"`
	Publickey string `json:"publickey"`
}


// InitLedger adds a base set of coinaccounts to the ledger
func (s *SmartContract) InitLedger(ctx contractapi.TransactionContextInterface) error {
	coinaccounts := []CoinAccount{
		CoinAccount{Account: "test1", Balance: 100, Publickey: "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAtHXGVArjsMMKHoxDOpAlsX2GHcJY2OfLUWrKB5bu8iSrlK4iqml7TtsEbrWRqgdTKbZ65yAqPeDBAsWoo5ZMtsNXKf3OsmOace2moAHkDGDUvEOSDYcU0akKnCUD98dVcHyyEg8VyGwXlDFYwCoRGhj42uy9xdNV5XNdBW/5+ZLgKG/iJ36bI+FdhhdKAqXYf6ikPetU2jglXV7/k4V4CA1kK9omgCpBj3o46RE3AmTstu4I82NxnhsEM3n0rpzIvz6CMbQMof2oQDmvHAqT2fHML6EF7p3nfyRrf9z8w52vQitIs5X0Nve1cmsmhbUThm9+clu3XplYk1cERPJ3nwIDAQAB"},
		CoinAccount{Account: "test2", Balance: 200, Publickey: "MIIBITANBgkqhkiG9w0BAQEFAAOCAQ4AMIIBCQKCAQB3Ubf3TEJDPtaJZdupbzIq7e4hZbkrjsrPVUtDdfMGEYaxtj60Idcj6MuVUjCmGXOYkk9LjFmJpS3pxaxmbIY6yxTpDsTL6kUcH5F/OJT7aDkiRQuMff6vvBFT2Zji0nV/LkC/exgNPU/ogceMiVaTRowjUnwIy3/obuqwCOphapZkruLbLUO/+hELfKKMVyoPY0jh6FTWxFx0tOCXGvOFKkhrAjyH/ZJkcI9bP0CliNEu7ttGpGgXKSrXCXIinSJgwgQhTCXR97dLJkLwf5cFaC+8j0Zr9ghvcYwLPYbQmaLjY+We/otplTgPslHNdf5gVNMrW5oK6BRsfGMcE6g/AgMBAAE="},
		CoinAccount{Account: "test3", Balance: 300, Publickey: "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA3XEVLekATXrNhnNnQRVgBkbw+Sv4rQyLyvHA6S3yxWBQFP3bvrsaCGjIO52OUuvOi8fIHIAU9gD0YrglE1ZDsgForFhtt4+N3/XI9J5pObj9kGz+Y/4Rh61MuimrOokESZUMEj6Q8IdtCVMFo18BxCZIa0CkXs28VIQL50OpDWmio9Xg4tmaBH2JX50WAQRvtagF2QS3JEQvJvQAvap+3KgWh/3332emzrVzOxjS35tEVov2aNha1uZgoc4+YULa3ljTx8igeS9Qj34suBIPmwUmo7qOVdemzmrfIg2RorBA7DauBUF+vY67Js31dZU/LEgrvuj27sd99JqOX6vntQIDAQAB"},
		CoinAccount{Account: "test4", Balance: 400, Publickey: "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAlBt1zMfBpu+M5JmCza0WPSIE7T7J3ShWzQuIx1xI2lH4VEtXw9PFEq6rbL9WDQHMh8IfwIJDzZmMW8xSIiHqBNAz5dSgs9tKe7C6mLSaVkjtQ/W3ln3Z6ufCWWjwPFxsboKYU9SrNJuubAXCWQ/AYN8dd75os31f9badXkcb3BRxjsdSqGxGjlIm7r8EWQMJeEvdv6aLkSew6bweLiebkZrQiaHbjETN+aLJM5e6DG4Hld2Ya577//F1Dsf73RyrYzK6AzXUlgK0znr6OjkXrJCAbiCDUfQHsf8MRP4nVAfrCmbfn4TUnN4Usap4D3JpcAcIYM5AfINxUsTQx6nEOQIDAQAB"},
		CoinAccount{Account: "test5", Balance: 500, Publickey: "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAzpB9ZPgGFiNv1K6jVqRMWJBV/3iYzIRB46EZ4pRxzPt5SFGuTDc3hfawx7WbsWwZyTwoBRUyPnL1iODLRn8Weo6Xy9X5hdMMmAoiZoIssyIhNsZhTCzteBz8CVTbSDUX00yhQjJZVkrTBkBHFllzljiY1+Ovpp0hlAAzmPs9kr0TPr61lCBU4D9IA4ibL8WqCunvt3eaOFNKXGy0sYwe9bXoqyNDWxil7etZEJLv/9igLYWuFGQvlAPabLvu7vi3V6DsLLli2SFVQZ+IHJxD7dI49rL1TqDWSU2+FepXn0ynFEGiUQ2Kd6o/RoHMTyxQG9eQK7ykCKM520rgZbtNbQIDAQAB"},
	}

	for _, coinaccount := range coinaccounts {
		coinaccountAsBytes, _ := json.Marshal(coinaccount)
		err := ctx.GetStub().PutState(coinaccount.Account, coinaccountAsBytes)

		if err != nil {
			return fmt.Errorf("Failed to put to world state. %s", err.Error())
		}
	}

	return nil
}

// CreateCoinAccount adds a new coinaccount to the world state with given details
func (s *SmartContract) CreateCoinAccount(ctx contractapi.TransactionContextInterface, account string, balance float64, publickey string) error {
	coinaccount := CoinAccount{
		Account:   account,
		Balance:  balance,
		Publickey: publickey,
	}

	coinaccountAsBytes, _ := json.Marshal(coinaccount)

	return ctx.GetStub().PutState(account, coinaccountAsBytes)
}

// QueryCar returns the car stored in the world state with given id
func (s *SmartContract) QueryCoinAccount(ctx contractapi.TransactionContextInterface, account string) (*CoinAccount, error) {
	coinaccountAsBytes, err := ctx.GetStub().GetState(account)

	if err != nil {
		return nil, fmt.Errorf("Failed to read from world state. %s", err.Error())
	}

	if coinaccountAsBytes == nil {
		return nil, fmt.Errorf("%s does not exist", account)
	}

	coinaccount := new(CoinAccount)
	_ = json.Unmarshal(coinaccountAsBytes, coinaccount)

	return coinaccount, nil
}

func RsaVerySignWithSha256Base64(originalData, signData, pubKey string) error{
	sign, err := base64.StdEncoding.DecodeString(signData)
	if err != nil {
		return err
	}
	public, _ := base64.StdEncoding.DecodeString(pubKey)
	pub, err := x509.ParsePKIXPublicKey(public)
	if err != nil {
		return err
	}
	hash := sha256.New()
	hash.Write([]byte(originalData))
	return rsa.VerifyPKCS1v15(pub.(*rsa.PublicKey), crypto.SHA256, hash.Sum(nil), sign)
}

// QueryCar returns the car stored in the world state with given id
func (s *SmartContract) GetBalance(ctx contractapi.TransactionContextInterface, id int, timestamp string, account string, key string, signature string) (float64, error) {
	coinaccountAsBytes, err := ctx.GetStub().GetState(account)

	if err != nil {
		return -1, fmt.Errorf("Failed to read from world state. %s", err.Error())
	}

	if coinaccountAsBytes == nil {
		return -1, fmt.Errorf("%s does not exist", account)
	}

	coinaccount := new(CoinAccount)
	_ = json.Unmarshal(coinaccountAsBytes, coinaccount)

	var verifyError error
	verifyError = RsaVerySignWithSha256Base64(strconv.Itoa(id)+"+"+timestamp+"+account+"+account+"+"+key, signature, coinaccount.Publickey)
	if verifyError != nil {
		return -1, fmt.Errorf("Failed to verify signature. %s,%s,%s,%s", verifyError.Error(), strconv.Itoa(id)+"+"+timestamp+"+account+"+account+"+"+key, signature, coinaccount.Publickey)
	}

	return coinaccount.Balance, nil
}

// QueryCar returns the car stored in the world state with given id
func (s *SmartContract) QuerySend(ctx contractapi.TransactionContextInterface, id int, timestamp string, from string, to string, amount float64, key string, signature string) (float64, error) {
	if amount <= 0 {
		return -1, fmt.Errorf("Failed, amount error")
	}

	coinaccountAsBytes, err := ctx.GetStub().GetState(from)

	if err != nil {
		return -1, fmt.Errorf("Failed to read from world state. %s", err.Error())
	}

	if coinaccountAsBytes == nil {
		return -1, fmt.Errorf("%s does not exist", from)
	}

	coinaccount := new(CoinAccount)
	_ = json.Unmarshal(coinaccountAsBytes, coinaccount)

	cointoAsBytes, err := ctx.GetStub().GetState(to)

	if err != nil {
		return -1, fmt.Errorf("Failed to read from world state. %s", err.Error())
	}

	if cointoAsBytes == nil {
		return -1, fmt.Errorf("%s does not exist", to)
	}

	cointo := new(CoinAccount)
	_ = json.Unmarshal(cointoAsBytes, cointo)

	var verifyError error
	verifyError = RsaVerySignWithSha256Base64(strconv.Itoa(id)+"+"+timestamp+"+from+"+from+"+to+"+to+"+amount+"+fmt.Sprint(amount)+"+"+key, signature, coinaccount.Publickey)
	if verifyError != nil {
		return -1, fmt.Errorf("Failed to verify signature. %s,%s,%s,%s", verifyError.Error(), strconv.Itoa(id)+"+"+timestamp+"+from+"+from+"+to+"+to+"+amount+"+fmt.Sprint(amount)+"+"+key, signature, coinaccount.Publickey)
	}

	if coinaccount.Balance < amount {
		return -1, fmt.Errorf("Failed to send, balance is not enough.")
	}

	coinaccount.Balance = coinaccount.Balance - amount
	coinaccountAsBytesNew, _ := json.Marshal(coinaccount)
	err = ctx.GetStub().PutState(coinaccount.Account, coinaccountAsBytesNew)

		if err != nil {
			return -1, fmt.Errorf("Failed to put to world state. %s", err.Error())
		}
	cointo.Balance = cointo.Balance + amount
	cointoAsBytesNew, _ := json.Marshal(cointo)
	err = ctx.GetStub().PutState(cointo.Account, cointoAsBytesNew)

		if err != nil {
			return -1, fmt.Errorf("Failed to put to world state. %s", err.Error())
		}

	return coinaccount.Balance, nil
}

// QueryAllCars returns all cars found in world state
func (s *SmartContract) QueryAllCoinAccounts(ctx contractapi.TransactionContextInterface) ([]*CoinAccount, error) {
	resultsIterator, err := ctx.GetStub().GetStateByRange("", "")
        if err != nil {
		return nil, err
        }
	defer resultsIterator.Close()

        var accounts []*CoinAccount
        for resultsIterator.HasNext() {
                queryResponse, err := resultsIterator.Next()
                if err != nil {
			return nil, err
		}

		var account CoinAccount
		err = json.Unmarshal(queryResponse.Value, &account)
		if err != nil {
			return nil, err
		}
		accounts = append(accounts, &account)
	}
	return accounts, nil
}

// ChangeCarOwner updates the owner field of car with given id in world state
//func (s *SmartContract) ChangeCarOwner(ctx contractapi.TransactionContextInterface, carNumber string, newOwner string) error {
//	car, err := s.QueryCar(ctx, carNumber)
//
//	if err != nil {
//		return err
//	}
//
//	car.Owner = newOwner
//
//	carAsBytes, _ := json.Marshal(car)
//
//	return ctx.GetStub().PutState(carNumber, carAsBytes)
//}

func main() {

	chaincode, err := contractapi.NewChaincode(new(SmartContract))

	if err != nil {
		fmt.Printf("Error create fabcoin chaincode: %s", err.Error())
		return
	}

	if err := chaincode.Start(); err != nil {
		fmt.Printf("Error starting fabcoin chaincode: %s", err.Error())
	}
}

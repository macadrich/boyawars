package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"net/http"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	biddingwars "gitlab.com/macadrich/boyawars/pkg/boyawars"
)

const (
	energiURL       = "your-test-network"
	contractAddress = "smart-contract-address"
	chainID         = 49797
)

type Account struct {
	Name  string `json:"account"`
	Hash  string `json:"hash"`
	Value int64  `json:"value"`
}

var client *ethclient.Client
var biddingWarsContract *biddingwars.Biddingwars

func init() {
	var err error
	client, err = ethclient.Dial(energiURL)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	biddingWarsContract, err = biddingwars.NewBiddingwars(common.HexToAddress(contractAddress), client)
	if err != nil {
		log.Fatalf("Failed to instantiate a smart contract: %v", err)
	}
}

func authorizeAccount(w http.ResponseWriter, account Account) (*bind.TransactOpts, error) {
	vprivateKey, err := crypto.HexToECDSA(account.Hash)
	if err != nil {
		log.Fatalf("Invalid private key: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return nil, err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(vprivateKey, big.NewInt(chainID))
	if err != nil {
		log.Fatalf("Failed to create authorized transactor: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return nil, err
	}

	return auth, nil
}

// Handler
func startGameHandler(w http.ResponseWriter, r *http.Request) {
	var account Account

	err := json.NewDecoder(r.Body).Decode(&account)
	if err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	// Create an authorized transactor
	auth, err := authorizeAccount(w, account)
	if err != nil {
		log.Fatalf("Failed to create authorized transactor: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Send transaction to start the game
	tx, err := biddingWarsContract.StartGame(auth)
	if err != nil {
		log.Printf("Failed to start the game: %v", err)
		http.Error(w, "Failed to start the game", http.StatusInternalServerError)
		return
	}
	fmt.Println("Gamer started..")
	fmt.Fprintf(w, "Game started; transaction hash: %s", tx.Hash().Hex())
}

func bidHandler(w http.ResponseWriter, r *http.Request) {
	var account Account

	// Decode the request body
	err := json.NewDecoder(r.Body).Decode(&account)
	if err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	// Create an authorized transactor
	auth, err := authorizeAccount(w, account)
	if err != nil {
		log.Fatalf("Failed to create authorized transactor: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	auth.Value = big.NewInt(account.Value)

	// Send transaction to bid
	tx, err := biddingWarsContract.Bid(auth)
	if err != nil {
		log.Printf("Failed to bid in game: %v", err)
		http.Error(w, "Failed to bid in game", http.StatusInternalServerError)
		return
	}
	fmt.Println("Gamer Bidding..")
	fmt.Fprintf(w, "Bid started; transaction hash: %s", tx.Hash().Hex())
}

func endGameHandler(w http.ResponseWriter, r *http.Request) {
	var account Account

	// Decode the request body
	err := json.NewDecoder(r.Body).Decode(&account)
	if err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	// Create an authorized transactor
	auth, err := authorizeAccount(w, account)
	if err != nil {
		log.Fatalf("Failed to create authorized transactor: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Send transaction to end the game
	tx, err := biddingWarsContract.EndGame(auth)
	if err != nil {
		log.Printf("Failed to end the game: %v", err)
		http.Error(w, "Failed to end the game", http.StatusInternalServerError)
		return
	}
	fmt.Println("Game Ended")
	fmt.Fprintf(w, "Game Ended; transaction hash: %s", tx.Hash().Hex())
}

func main() {
	http.HandleFunc("/start", startGameHandler)
	http.HandleFunc("/bid", bidHandler)
	http.HandleFunc("/end", endGameHandler)

	fmt.Println("Server started at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

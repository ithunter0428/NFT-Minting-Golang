package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/rs/zerolog/log"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

const RinkebyExplorer = "https://rinkeby.etherscan.io/tx/"
const ethBlockchainName = "eth"

func CreateInstance(endpointUrl string, walletPk string, smartContractAddress string, chainId *big.Int) (instance *Store, txOptions *bind.TransactOpts, err error) {

	client, err := ethclient.Dial(endpointUrl)
	if err != nil {
		log.Error().Err(err)
		return
	}

	privateKey, err := crypto.HexToECDSA(walletPk)
	if err != nil {
		log.Error().Err(err)
		return
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Error().Err(err).Msg("Cannot assert type: publicKey is not of type *ecdsa.PublicKey")
		return
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Error().Err(err)
		return
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Error().Err(err)
		return
	}

	// NewKeyedTransactorWithChainID
	if chainId.Int64() != -1 {
		txOptions, err = bind.NewKeyedTransactorWithChainID(privateKey, chainId)
		if err != nil {
			log.Error().Err(err)
			return
		}
	} else {
		txOptions = bind.NewKeyedTransactor(privateKey)
	}

	txOptions.Nonce = big.NewInt(int64(nonce))
	txOptions.Value = big.NewInt(0)     // in wei
	txOptions.GasLimit = uint64(300000) // in units
	txOptions.GasPrice = gasPrice

	address := common.HexToAddress(smartContractAddress)
	instance, err = NewStore(address, client)
	if err != nil {
		log.Error().Err(err)
		return
	}
	return instance, txOptions, nil
}

func MintNft(instance *Store, txOptions *bind.TransactOpts, metadataURI string) (txHash string, err error) {

	tx, err := instance.MintToken(txOptions, metadataURI)
	if err != nil {
		log.Error().Err(err)
		return
	}
	txHash = tx.Hash().Hex()
	fmt.Printf("tx sent: %s", txHash)

	return txHash, nil
}

func MintEthNft(name string, description string, imageCid string) (txHash string, err error) {
	// Make json and upload
	nftJson := NftJson{Name: name, Description: description, Image: imageCid}
	metadataURI, err := UploadJsonToIpfs(nftJson)
	if err != nil {
		log.Error().Err(err).Msg("Error uploading json to Ipfs")
		return
	}

	// Create tx
	instance, txOptions, err := CreateInstance(C.EthEndpointUrl, C.EthDeployWalletPk, C.EthNftContractAddress, big.NewInt(-1))
	if err != nil {
		log.Error().Err(err)
		return
	}

	txHash, err = MintNft(instance, txOptions, metadataURI)
	if err != nil {
		log.Error().Err(err)
		return
	}

	txHash = RinkebyExplorer + txHash
	return
}

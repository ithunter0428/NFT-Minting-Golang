package main

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type cfg struct {
	EthDeployWalletPk           string `envconfig:"ETH_DEPLOY_WALLET_PK"`
	EthDeployWalletAddress      string `envconfig:"ETH_DEPLOY_WALLET_ADDRESS"`
	EthEndpointUrl              string `envconfig:"ETH_ENDPOINT_URL"`
	EthNftContractAddress       string `envconfig:"ETH_NFT_CONTRACT_ADDRESS"`
	EthNftMarketContractAddress string `envconfig:"ETH_NFT_MARKET_CONTRACT_ADDRESS"`
	NftStorageKey               string `envconfig:"NFT_STORAGE_KEY"`
}

var C = new(cfg)

func Init() {
	_ = godotenv.Overload(".env", ".env.local")
	_ = envconfig.Process("", C)

	C.EthDeployWalletPk = "d5bd5848ebe290405812422144f9e32e16589ba268853ed24d7df44a16970bc6"
	C.EthDeployWalletAddress = "0x7a9B4F815C02ff69743E9a3e665A591AE4bB372c"
	C.EthEndpointUrl = "https://rinkeby.infura.io/v3/14dc2121c0be4716983144938da1b9e6"
	C.EthNftContractAddress = "0x0ED06150f3Bb1E164d0065fAa4EAbC4843659Ae8"
	C.NftStorageKey = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJkaWQ6ZXRocjoweGQzNjY3MEMzNTI3OEJhZDBkZDk3MDU3Y0M5NEM2NDJkOTY2MDkxQzEiLCJpc3MiOiJuZnQtc3RvcmFnZSIsImlhdCI6MTY1ODY0OTAwNDgxNiwibmFtZSI6IlBhc3Npb25lZXIifQ.UFk5yzBR0FCLciu1zQR5-N1kQYq_AzMJyRtZIWvJXqM"
}

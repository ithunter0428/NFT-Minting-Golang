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

	C.EthDeployWalletPk = "a09d63baffb90ad6926e7caaddc437f9193eed759ecef20740b3270e3797cc08"
	C.EthDeployWalletAddress = "0x3C9684637FAD39E2D230F4E8867DEE89F9689C16"
	C.EthEndpointUrl = "https://rinkeby.infura.io/v3/14dc2121c0be4716983144938da1b9e6"
	C.EthNftContractAddress = "0x0ED06150f3Bb1E164d0065fAa4EAbC4843659Ae8"
	C.NftStorageKey = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJkaWQ6ZXRocjoweGQzNjY3MEMzNTI3OEJhZDBkZDk3MDU3Y0M5NEM2NDJkOTY2MDkxQzEiLCJpc3MiOiJuZnQtc3RvcmFnZSIsImlhdCI6MTY1ODY0OTAwNDgxNiwibmFtZSI6IlBhc3Npb25lZXIifQ.UFk5yzBR0FCLciu1zQR5-N1kQYq_AzMJyRtZIWvJXqM"
}

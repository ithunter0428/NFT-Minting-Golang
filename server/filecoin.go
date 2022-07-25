package main

// https://github.com/nftstorage/go-client/blob/main/docs/NFTStorageAPI.md#store

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"

	"github.com/rs/zerolog/log"
)

type NeoStorageResponse struct {
	Ok    bool
	Value ValueField
}

type ValueField struct {
	Cid     string
	Created string
}

type NftJson struct {
	Name        string
	Description string
	Image       string
}

func UploadJsonToIpfs(nftJson NftJson) (cid string, err error) {
	b, err := json.Marshal(nftJson)
	if err != nil {
		log.Error().Err(err).Msg("Error converting to json")
		return
	}
	cid, err = postToNftStorage(bytes.NewReader(b))
	if err != nil {
		log.Error().Err(err).Msg("Error uploading json to Ipfs")
		return
	}
	return
}

func UploadFileToIpfs(fileUrl string) (cid string, err error) {
	dir, _ := os.Getwd()
	localFilePath := fmt.Sprintf("%s%s%s", path.Dir(dir), "/images/", fileUrl)
	fmt.Fprintf(os.Stdout, "Image File Path: %v\n", localFilePath)

	// Upload localFilePath
	f, err := os.Open(localFilePath)
	if err != nil {
		return
	}
	defer f.Close()
	cid, err = postToNftStorage(f)
	if err != nil {
		log.Error().Err(err).Msg("Error while posting to NFT storage")
		return "", err
	}
	return
}

func postToNftStorage(body io.Reader) (cid string, err error) {
	req, err := http.NewRequest("POST", "https://api.nft.storage/upload", body)
	if err != nil {
		return
	}

	fmt.Fprintf(os.Stdout, "NftStorageKey: %v\n", C.NftStorageKey)
	req.Header.Set("Authorization", "Bearer "+C.NftStorageKey)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}

	defer resp.Body.Close()
	fmt.Fprintf(os.Stdout, "NFT Storage response: %v\n", resp.StatusCode)
	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return "", err
		}
		var response NeoStorageResponse
		json.Unmarshal(bodyBytes, &response)
		return response.Value.Cid, nil
	}
	return
}

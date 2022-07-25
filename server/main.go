package main

import (
	"errors"
	"fmt"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

type CreateNFTRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Image       string `json:"image"`
}

func getCreateNftRequest(ctx *gin.Context) (createNftRequest CreateNFTRequest, err error) {
	if err = ctx.BindJSON(&createNftRequest); err != nil {
		log.Error().Err(err).Msg("Failed to encode CreateNFTRequest")
		return
	}

	if len(createNftRequest.Name) == 0 {
		createNftRequest.Name = "Digital Verse"
	}
	if len(createNftRequest.Description) == 0 {
		createNftRequest.Description = "Celebrity video"
	}
	if len(createNftRequest.Image) == 0 {
		log.Error().Msg("image not specified")
		return createNftRequest, errors.New("image not specified")
	}
	return createNftRequest, nil
}

func main() {
	Init()

	r := gin.Default()

	r.Use(cors.Default())

	r.Static("/images", "./images")

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("/create_eth_nft", func(ctx *gin.Context) {
		createReq, err := getCreateNftRequest(ctx)
		if err != nil {
			log.Error().Err(err).Msg("Failed to get params from request")
			return
		}

		ipfsCid, err := UploadFileToIpfs(createReq.Image)
		if err != nil {
			log.Error().Err(err).Msg("Failed to upload file to IPFS")
			return
		}

		txHash, err := MintEthNft(createReq.Name, createReq.Description, ipfsCid)
		if err != nil {
			log.Error().Err(err).Msg("Failed to mint NFT")
			return
		}

		ctx.JSON(200, gin.H{
			"tx_hash": txHash,
			"url":     RinkebyExplorer,
			"fileUrl": "https://ipfs.io/ipfs/" + ipfsCid,
			"error":   err,
		})
	})

	r.POST("/upload_file", func(c *gin.Context) {
		file, err := c.FormFile("image")
		if err != nil {
			log.Error().Err(err).Msg("image upload error")
			c.JSON(500, gin.H{
				"message": "Server error",
				"data":    nil,
			})
		}

		// save image to ./images dir
		uniqueId := uuid.New()
		filename := strings.Replace(uniqueId.String(), "-", "", -1)
		fileExt := strings.Split(file.Filename, ".")[1]
		image := fmt.Sprintf("%s.%s", filename, fileExt)

		err = c.SaveUploadedFile(file, fmt.Sprintf("./images/%s", image))
		if err != nil {
			log.Error().Err(err).Msg("image save error")
			c.JSON(500, gin.H{
				"message": "Server error",
				"data":    nil,
			})
		}

		// create meta data and send to client
		imageUrl := fmt.Sprintf("http://localhost:8080/images/%s", image)
		data := map[string]interface{}{
			"imageName": image,
			"imageUrl":  imageUrl,
			"header":    file.Header,
			"size":      file.Size,
		}
		c.JSON(201, gin.H{
			"message": "Image uploaded successfully",
			"data":    data,
		})
	})

	r.Run()
}

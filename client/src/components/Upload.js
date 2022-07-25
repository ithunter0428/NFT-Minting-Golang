import { Button, Heading, VStack, Image, HStack, Tag, Input, Link, Text } from "@chakra-ui/react";
import React from "react";
import { useRef } from "react";
import useUpload from "../hooks/useUpload";

function Upload() {
  const imageRef = useRef(null);
  const {
    loading,
    image,
    name,
    description,
    uploadedImage,
    mintedNFT,

    handleMintNFT,
    handleChangeImage,
    handleUploadImage,
    handleChangeName,
    handleChangeDescription
  } = useUpload();
  return (
    <>
      <input
        style={{ display: "none" }}
        type="file"
        accept="image/*"
        ref={imageRef}
        onChange={handleChangeImage}
      />
      <VStack>
        <Heading>NFT Minting</Heading>
        <Button
          onClick={() => imageRef.current.click()}
          colorScheme="blue"
          size="lg"
        >
          Select Image
        </Button>
      </VStack>

      {image && (
        <VStack my="4">
          <Image
            src={URL.createObjectURL(image)}
            width="300px"
            height="300px"
            alt="selected image..."
          />
          <Button
            onClick={handleUploadImage}
            variant="outline"
            colorScheme="green"
            isLoading={loading}
          >
            Upload
          </Button>
        </VStack>
      )}

      {mintedNFT && (
        <VStack my="4">
          <Text>
            File Url: &nbsp;
            <Link color='teal.500' href={mintedNFT?.fileUrl} isExternal>
              {mintedNFT?.fileUrl}
            </Link>
          </Text>
          <Text>
            Transaction: &nbsp;
            <Link color='teal.500' href={mintedNFT?.tx_hash} isExternal>
              {mintedNFT?.tx_hash}
            </Link>
          </Text>
        </VStack>
      )}

      {uploadedImage && (
        <VStack my="4">
          <Image
            src={uploadedImage?.imageUrl}
            width="300px"
            height="300px"
            alt={uploadedImage?.imageName}
          />

          <Input
            value={name}
            onChange={(e) => handleChangeName(e.target.value)}
            width="300px"
            placeholder='Name'
            _placeholder={{ opacity: 1, color: 'gray.500' }}
          />

          <Input
            value={description}
            onChange={(e) => handleChangeDescription(e.target.value)}
            width="300px"
            placeholder='Description'
            _placeholder={{ opacity: 1, color: 'gray.500' }}
          />

          <HStack>
            <Tag variant="outline" colorScheme="blackAlpha">
              ~ {Math.floor(uploadedImage?.size / 1024)} Kb
            </Tag>
            <Button
              variant="solid"
              colorScheme="red"
              onClick={handleMintNFT}
              isLoading={loading}
            >
              Mint
            </Button>
          </HStack>
        </VStack>
      )}
    </>
  );
}

export default Upload;

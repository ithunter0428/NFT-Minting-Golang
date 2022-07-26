import { useState } from "react";
import axios from "axios";
import { useToast } from "@chakra-ui/react";

const useUpload = () => {
  const [image, setImage] = useState(null);
  const [name, setName] = useState(null);
  const [description, setDescription] = useState(null);
  const [loading, setLoading] = useState(false);

  const [uploadedImage, setUploadedImage] = useState(null);
  const [mintedNFT, setMintedNFT] = useState(null);

  const toast = useToast();

  const handleChangeImage = (e) => {
    setImage(e.target.files[0]);
    setMintedNFT(null);
  };

  const handleUploadImage = async () => {
    try {
      setLoading(true);
      const formData = new FormData();
      formData.append("image", image);
      const res = await axios.post("/upload_file", formData);
      if (res.data.data) {
        console.log(res.data);
        setUploadedImage(res.data.data);
        toast({
          title: "Image Uploaded",
          description: res.data.message,
          status: "success",
          duration: 4000,
          isClosable: true,
        });
      }
    } catch (error) {
      console.log(error);
    } finally {
      setImage(null);
      setLoading(false);
    }
  };

  const handleMintNFT = async () => {
    try {
      setLoading(true);

      const res = await axios.post(`/create_eth_nft`, {
        Name: name,
        Description: description,
        Image: uploadedImage.imageName,
      });
      if (res.data) {
        console.log(res.data);
        setMintedNFT(res.data);

        setUploadedImage(null);
        setName("");
        setDescription("");

        toast({
          title: "NFT Minted",
          description: res.data.message,
          status: "success",
          duration: 4000,
          isClosable: true,
        });
      }
    } catch (error) {
      console.log(error);
      toast({
        title: "Error",
        description: error?.response?.data?.error,
        status: "error",
        duration: 4000,
        isClosable: true,
      });
    } finally {
      setLoading(false);
    }
  };

  const handleChangeName = (val) => {
    setName(val);
  };

  const handleChangeDescription = (val) => {
    setDescription(val);
  };

  return {
    image,
    name,
    description,
    uploadedImage,
    mintedNFT,
    loading,
    handleChangeImage,
    handleUploadImage,
    handleMintNFT,
    handleChangeName,
    handleChangeDescription,
  };
};

export default useUpload;

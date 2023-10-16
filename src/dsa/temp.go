package dsa

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type FileMetaData struct {
	Hash     string   `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	Size_    string   `protobuf:"bytes,2,opt,name=size,proto3" json:"size,omitempty"`
	Encoding string   `protobuf:"bytes,3,opt,name=encoding,proto3" json:"encoding,omitempty"`
	TxnHash  []string `protobuf:"bytes,4,rep,name=txnHash,proto3" json:"txnHash,omitempty"`
}

type MsgHostManifest struct {
	Creator      string                  `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Version      string                  `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	HashFunction string                  `protobuf:"bytes,3,opt,name=hashFunction,proto3" json:"hashFunction,omitempty"`
	MetaData     map[string]FileMetaData `json:"metaData"`
}

type MsgStorage struct {
	Creator string            `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Version string            `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	Storage map[string]string `protobuf:"bytes,3,rep,name=storage,proto3" json:"data,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func computeSHA1(data []byte) string {
	hasher := sha256.New()

	hasher.Write(data)
	return hex.EncodeToString(hasher.Sum(nil))
}

func ProcessFolder(folderPath string) (MsgHostManifest, MsgStorage, error) {
	metaData := make(map[string]FileMetaData)
	storage := make(map[string]string)

	err := filepath.Walk(folderPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			// Read the file content
			content, err := ioutil.ReadFile(path)
			if err != nil {
				return err
			}

			// Compute SHA1 hash for the content
			hash := computeSHA1(content)

			// Use relative path instead of absolute
			relativePath := strings.TrimPrefix(path, folderPath)

			// Populate FileMetaData for the file
			metaData[relativePath] = FileMetaData{
				Hash:     hash,
				Size_:    fmt.Sprintf("%d", len(content)),
				Encoding: "utf-8",
				TxnHash:  []string{""},
			}

			// Store the base64 encoded content for simplicity
			encodedContent := base64.StdEncoding.EncodeToString(content)
			storage[relativePath] = encodedContent
		}
		return nil
	})

	if err != nil {
		return MsgHostManifest{}, MsgStorage{}, err
	}

	msgHostManifest := MsgHostManifest{
		Creator:      "nil",
		Version:      "v1.0",
		HashFunction: "SHA1",
		MetaData:     metaData,
	}

	msgStorage := MsgStorage{
		Creator: "nil",
		Version: "v1.0",
		Storage: storage,
	}

	return msgHostManifest, msgStorage, nil
}

func temp_main() {
	// Parse folder path from command-line arguments
	folderPath := flag.String("folder", "", "path to folder")
	flag.Parse()

	if *folderPath == "" {
		fmt.Println("Please provide a folder path using the -folder flag")
		os.Exit(1)
	}

	msgHostManifest, msgStorage, err := ProcessFolder(*folderPath)
	if err != nil {
		fmt.Println("Error processing folder:", err)
		os.Exit(1)
	}

	// Print the populated structs
	jsonData, err := json.MarshalIndent(msgHostManifest, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling to JSON:", err)
		return
	}
	fmt.Println(string(jsonData))

	jsonData2, err := json.MarshalIndent(msgStorage, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling to JSON:", err)
		return
	}
	fmt.Println(string(jsonData2))
}

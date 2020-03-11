package main

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
	"github.com/gin-gonic/gin"
	"net/http"
	"encoding/json"
	"fmt"
)

// 首先我们需要定义一个结构体，用来存放块

type Block struct{
	Index int     // 整个块在链中的位置
	Timestamp string      // 块生成的时间戳
	BPM int          // 心跳
	Hash string         // 这个块通过 SHA256 算法生成的散列值
	PreHash string   // 前个块通过SHA256算法生成的散列值
}


// 其次，我们需要定义一个结构表示整条链，最简单的即为一个Block的Slice
var Blockchain []Block


func calculateHash(block Block) string {
	// 用来计算给定数据的SHA256散列值
	record := string(block.Index) + block.Timestamp + string(block.BPM) + block.PreHash
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}


func generateBlock(oldBlock Block, BPM int) (Block, error) {
	// 创建一个新块
	var newBlock Block
	t := time.Now()
	newBlock.Index = oldBlock.Index + 1
	newBlock.Timestamp = t.String()
	newBlock.BPM = BPM
	newBlock.PreHash = oldBlock.Hash
	newBlock.Hash = calculateHash(newBlock)
	Blockchain = append(Blockchain, newBlock)
	return newBlock, nil
}

func isBlockValid(newBlock, oldBlock Block) bool {
	// 校验块 判断一个块是否有被篡改。
	if oldBlock.Index != newBlock.Index - 1 {
		return false
	}
	if oldBlock.Hash != newBlock.PreHash {
		return false
	}
	if calculateHash(newBlock) != newBlock.Hash {
		return false
	}
	return true
}


func replaceChain(newBlocks []Block) {
	if len(newBlocks) > len(Blockchain) {
		Blockchain = newBlocks
	}
}


func getChain (c *gin.Context){
	data, _ := json.Marshal(Blockchain)
	fmt.Printf("%s\n", data)
	c.String(http.StatusOK, "%s", data)
}

func insertChain (c *gin.Context){
	generateBlock(Blockchain[(len(Blockchain)-1)], 55)
	c.String(http.StatusOK, "%s", len(Blockchain))
}

func main() {
	t := time.Now()
	genesisBlock := Block{0, t.String(), 0, "", ""}
	Blockchain = append(Blockchain, genesisBlock)
	router := gin.Default()
	router.GET("/", getChain)
	router.GET("/add", insertChain)
	router.Run(":8888")
}

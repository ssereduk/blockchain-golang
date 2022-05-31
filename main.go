package main

import "fmt"

//создание блока

//описание стрктури бока

type Block struct {
	Timestamp  int64
	Data []byte
	PrevBlockHash []byte
	Hash []byte
}

//соеденим все поля блока а результат захеэшируем SHA-256

func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)

	b.Hash = hash[:]
}

//создаем конструктор для нашег блока
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}}
	block.SetHash()
	return block
} 
//Блок Готов!

//------------------------
//первый блокчейн

type Blockchain struct {
     blocks []*Block
}

//возможность добавлять блоки в него

func (bc *Blockchain) AddBlock(data string) { 
      prevBlock := bc.blocks[len(bc.blocks)-1]
	  newBlock := NewBlock(data, prevBlock.Hash)
	  bc.blocks = append(bc.blocks, newBlock)
}

//создаем генезис блок, первый в цепочке

func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}

//создаем функцию создающую блокчейн с генезис блоком

func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{NewGenesisBlock()}}
}

//проверка работает ли корректно блокчейн

func main() {
	bc := NewBlockchain()

	bc.AddBlock("Send 1 ETH to Ivan")
	bc.AddBlock("Send 2 more ETH  to Ivan")

	for _, block := range bc.blocks {
		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println()
	}
	
	fmt.Println("Го ниндзя")
}
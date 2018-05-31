package main

import (
    "os"
    "fmt"
    "context"
    "encoding/json"
    "path/filepath"
    "io/ioutil"
    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/ethclient"
    "log"
)



type Config struct {
    Node string
    Port string
    Walletaddress string
    Privatekey string
}

type Allocation struct {
    UserId string
    Address string
    Amount int64
}

func LoadConfiguration() Config {
    var config Config
    dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
    configFile, _ := os.Open(dir+"/conf.json")
    defer configFile.Close()
    configBytes, _ := ioutil.ReadAll(configFile)
    json.Unmarshal(configBytes, &config)

    return config
}



func main() {
    config := LoadConfiguration()
    // Create an IPC based RPC connection to a remote Ethereum node
    conn, err := ethclient.Dial(config.Port)
    if err != nil {
        log.Fatalf("Failed to connect to the Ethereum client: %v", err)
    }
    client := conn

    println(client)
    // Get node and contract access details
    fmt.Println("Node :" ,config.Node)
    fmt.Println("Port :" ,config.Port)
    fmt.Println("Wallet :" ,config.Walletaddress)
    fmt.Println("Key :" ,config.Privatekey)

    // Contract Address for the AEN token
    // contract, err := NewAENToken(common.HexToAddress("0x0BEf619cF38cF0c22967289b8419720fBd1Db9f7"),conn)
    // if err != nil {
    //     log.Fatalf("Unable to Instatiate contract: %v", err)
    // }

    // // Get wallet amount Address
    // amt, _ := contract.BalanceOf(&bind.CallOpts{}, common.HexToAddress("0x0Cb539400641F8949d458b2d5a9e8A3C8bD488f2") )
    // fmt.Println(amt)

    // ctx := context.Background()
    // tx, pending, _ := conn.TransactionByHash(ctx, common.HexToHash("0xbfd84f3c895e130df9a66b63a07f02fb2703df5d6f7290d0bd57ce124f3419f7"))
    // if !pending {
    //     fmt.Println(tx)
    // }

    ctx := context.Background()
    tx, pending, _ := conn.TransactionByHash(ctx, common.HexToHash("0x30999361906753dbf60f39b32d3c8fadeb07d2c0f1188a32ba1849daac0385a8"))
    if !pending {
       fmt.Println(tx)
    }
}
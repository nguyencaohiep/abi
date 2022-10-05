package Crawl

import (
	"context"
	constant "event_logs/Constant"
	"fmt"
	"log"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/ethclient"
)

type LogIncreaseLiquidity struct {
	TokenId   *big.Int
	Liquidity *big.Int
	Amount0   *big.Int
	Amount1   *big.Int
}

type LogTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
}

// LogApproval ..
type LogApproval struct {
	TokenOwner common.Address
	Spender    common.Address
	TokenId    *big.Int
}

func CrawlUni() {
	client, err := ethclient.Dial(constant.URLINFURA)
	if err != nil {
		log.Fatal(err)
	}

	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(15645903),
		ToBlock:   big.NewInt(15645903),
	}

	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}

	contractAbi, err := abi.JSON(strings.NewReader(string(constant.ABIUNI)))
	if err != nil {
		log.Fatal(err)
	}
	for _, value := range contractAbi.Events {
		fmt.Println("inputs : ", value.Inputs)
		fmt.Println("ID : ", value.ID)
		fmt.Println("Name : ", value.Name)
		fmt.Println("RawName : ", value.RawName)
		fmt.Println("Anonymous : ", value.Anonymous)
		fmt.Println("Sig : ", value.Sig)
		fmt.Println("String : ", value.String())
	}
	// logIncreaseLiquiditySig := []byte(contractAbi.Events["IncreaseLiquidity"].Sig)
	// LogApprovalSig := []byte("Approval(address,address,uint256)")
	// logIncreaseLiquiditySigHash := crypto.Keccak256Hash(logIncreaseLiquiditySig)
	// logApprovalSigHash := crypto.Keccak256Hash(LogApprovalSig)

	for _, vLog := range logs {
		// switch vLog.Topics[0].Hex() {
		// case logIncreaseLiquiditySigHash.Hex():
		fmt.Println()
		fmt.Printf("Log Name: IncreaseLiquidity\n")
		fmt.Println("Txh ", vLog.TxHash.Hex(), vLog.TxHash)
		fmt.Println("LogIndex ", vLog.Index)

		var logIncreaseLiquidity LogIncreaseLiquidity

		lenNonIndex := len(contractAbi.Events["IncreaseLiquidity"].Inputs.NonIndexed())
		tmp := make([]interface{}, lenNonIndex)

		err := contractAbi.UnpackIntoInterface(&tmp, "IncreaseLiquidity", vLog.Data)
		if err != nil {
			log.Fatal(err)
		}

		// transferEvent.From = common.HexToAddress(vLog.Topics[1].Hex())
		// transferEvent.To = common.HexToAddress(vLog.Topics[2].Hex())
		logIncreaseLiquidity.TokenId, _ = math.ParseBig256(vLog.Topics[1].Hex())
		// fmt.Printf("From: %s\n", transferEvent.From.Hex())
		// fmt.Printf("To: %s\n", transferEvent.To.Hex())
		fmt.Printf("TokenId: %v\n", logIncreaseLiquidity.TokenId)
		fmt.Println("data : ", tmp)
		// case logApprovalSigHash.Hex():
		fmt.Println()
		fmt.Printf("Log Name: Approval\n")
		fmt.Println("Txh ", vLog.TxHash.Hex())
		fmt.Println("LogIndex ", vLog.Index)

		var approvalEvent LogApproval
		lenNonIndex = len(contractAbi.Events["Approval"].Inputs.NonIndexed())
		tmp = make([]interface{}, lenNonIndex)

		err = contractAbi.UnpackIntoInterface(&tmp, "Approval", vLog.Data)
		if err != nil {
			log.Fatal(err)
		}

		// approvalEvent.TokenOwner = common.HexToAddress(vLog.Topics[1].Hex())
		// approvalEvent.Spender = common.HexToAddress(vLog.Topics[2].Hex())
		// approvalEvent.TokenId, _ = math.ParseBig256(vLog.Topics[3].Hex())

		fmt.Printf("Token Owner: %s\n", approvalEvent.TokenOwner.Hex())
		// fmt.Printf("Spender: %s\n", approvalEvent.Spender.Hex())
		fmt.Printf("TokenId: %v\n", approvalEvent.TokenId)
		fmt.Println("data : ", tmp)
	}
	// }
}

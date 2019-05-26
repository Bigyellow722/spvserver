package main

import (
	"errors"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
	"strconv"
)

type SimpleChaincode struct {
}

func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface) peer.Response {

	return shim.Success(nil)
}

func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	fun, args := stub.GetFunctionAndParameters()

	var result string
	var err error
	if fun == "set" {
		result, err = set(stub, args)
	} else if fun == "get" {
		result, err = get(stub, args)
	} else if fun == "transfer" {
		result, err = transfer(stub, args)
	} else {
		result, err = "", errors.New("不支持的服务")
	}
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success([]byte(result))
}

func set(stub shim.ChaincodeStubInterface, args []string) (string, error) {

	if len(args) != 3 {
		return "", fmt.Errorf("给定的参数错误")
	}

	err := stub.PutState(args[0], []byte(args[1]))
	if err != nil {
		return "", fmt.Errorf(err.Error())
	}

	err = stub.SetEvent(args[2], []byte{})
	if err != nil {
		return "", fmt.Errorf(err.Error())
	}

	return string(args[0]), nil

}

func transfer(stub shim.ChaincodeStubInterface, args []string) (string, error) {

	if len(args) != 4 {
		return "", fmt.Errorf("给定的参数错误")
	}

	valuename1, err := stub.GetState(args[0])
	if err != nil {
		return "", fmt.Errorf(err.Error())
	}

	valuename2, err := stub.GetState(args[1])
	if err != nil {
		return "", fmt.Errorf(err.Error())
	}

	name1int, err := strconv.ParseInt(string(valuename1), 10, 64)
	name2int, err := strconv.ParseInt(string(valuename2), 10, 64)

	value, err := strconv.ParseInt(args[2], 10, 64)

	name1int = name1int - value
	name2int = name2int + value

	name1str := strconv.FormatInt(name1int, 10)
	name2str := strconv.FormatInt(name2int, 10)

	err = stub.PutState(args[0], []byte(name1str))
	if err != nil {
		return "", fmt.Errorf(err.Error())
	}

	err = stub.PutState(args[1], []byte(name2str))
	if err != nil {
		return "", fmt.Errorf(err.Error())
	}

	err = stub.SetEvent(args[3], []byte{})
	if err != nil {
		return "", fmt.Errorf(err.Error())
	}

	return string(args[2]), nil

}

func get(stub shim.ChaincodeStubInterface, args []string) (string, error) {
	if len(args) != 1 {
		return "", fmt.Errorf("给定的参数错误")
	}
	result, err := stub.GetState(args[0])
	if err != nil {
		return "", fmt.Errorf("获取数据发生错误")
	}
	if result == nil {
		return "", fmt.Errorf("根据 %s 没有获取到相应的数据", args[0])
	}
	return string(result), nil

}

func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("启动SimpleChaincode时发生错误: %s", err)
	}
}

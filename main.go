package main

import (
	"fmt"
	"github.com/kongyixueyuan.com/kongyixueyuan/sdkInit"
	"github.com/kongyixueyuan.com/kongyixueyuan/service"
	"os"
)

const (
	configFile  = "config.yaml"
	initialized = false
	SimpleCC    = "simplecc"
)

func main() {

	initInfo := &sdkInit.InitInfo{

		ChannelID:     "kevinkongyixueyuan",
		ChannelConfig: os.Getenv("GOPATH") + "/src/github.com/kongyixueyuan.com/kongyixueyuan/fixtures/artifacts/channel.tx",

		OrgAdmin:       "Admin",
		OrgName:        "Org1",
		OrdererOrgName: "orderer.kevin.kongyixueyuan.com",

		ChaincodeID:     SimpleCC,
		ChaincodeGoPath: os.Getenv("GOPATH"),
		ChaincodePath:   "github.com/kongyixueyuan.com/kongyixueyuan/chaincode/",
		UserName:        "User1",
	}

	sdk, err := sdkInit.SetupSDK(configFile, initialized)
	if err != nil {
		fmt.Printf(err.Error())
		return
	}

	defer sdk.Close()

	err = sdkInit.CreateChannel(sdk, initInfo)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	channelClient, err := sdkInit.InstallAndInstantiateCC(sdk, initInfo)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(channelClient)

	serviceSetup := service.ServiceSetup{
		ChaincodeID: SimpleCC,
		Client:      channelClient,
	}

	fmt.Println("创建两个用户test1,test2")

	msg1, err := serviceSetup.SetInfo("test1", "10000")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(msg1)
	}

	msg2, err := serviceSetup.SetInfo("test2", "0")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(msg2)
	}

	fmt.Println("显示两个用户的余额")

	msg3, err := serviceSetup.GetInfo("test1")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(msg3)
	}
	msg4, err := serviceSetup.GetInfo("test2")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(msg4)
	}

	fmt.Println("执行转账操作")
	msg5, err := serviceSetup.Transfer("test1", "test2", "100")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(msg5)
	}

	fmt.Println("显示两个用户的余额")

	msg6, err := serviceSetup.GetInfo("test1")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(msg6)
	}
	msg7, err := serviceSetup.GetInfo("test2")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(msg7)
	}
	/*serviceSetup := service.ServiceSetup{
		ChaincodeID:SimpleCC,
		Client:channelClient,
	}

	view.Spvservice(serviceSetup)*/
}

package service

import "github.com/hyperledger/fabric-sdk-go/pkg/client/channel"

func (t *ServiceSetup) SetInfo(name string, num string) (string, error) {
	eventID := "eventSetInfo"
	reg, notifiler := regitserEvent(t.Client, t.ChaincodeID, eventID)
	defer t.Client.UnregisterChaincodeEvent(reg)

	req := channel.Request{ChaincodeID: t.ChaincodeID, Fcn: "set", Args: [][]byte{[]byte(name), []byte(num), []byte(eventID)}}
	resp, err := t.Client.Execute(req)
	if err != nil {
		return "", err
	}

	err = eventResult(notifiler, eventID)
	if err != nil {
		return "", nil
	}

	return string(resp.TransactionID), nil
}

func (t *ServiceSetup) Transfer(name1 string, name2 string, num string) (string, error) {
	eventID := "eventTransfer"
	reg, notifiler := regitserEvent(t.Client, t.ChaincodeID, eventID)
	defer t.Client.UnregisterChaincodeEvent(reg)

	req := channel.Request{ChaincodeID: t.ChaincodeID, Fcn: "transfer", Args: [][]byte{[]byte(name1), []byte(name2), []byte(num), []byte(eventID)}}
	resp, err := t.Client.Execute(req)

	if err != nil {
		return "", err
	}

	err = eventResult(notifiler, eventID)
	if err != nil {
		return "", nil
	}

	return string(resp.Payload), nil
}

func (t *ServiceSetup) GetInfo(name string) (string, error) {
	req := channel.Request{ChaincodeID: t.ChaincodeID, Fcn: "get", Args: [][]byte{[]byte(name)}}
	resp, err := t.Client.Query(req)
	if err != nil {
		return "", nil
	}
	return string(resp.Payload), nil
}

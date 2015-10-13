// File generated by github.com/ebuchman/rpc-gen

package core_client

import (
	"fmt"
	acm "github.com/eris-ltd/toadserver/Godeps/_workspace/src/github.com/tendermint/tendermint/account"
	ctypes "github.com/eris-ltd/toadserver/Godeps/_workspace/src/github.com/tendermint/tendermint/rpc/core/types"
	rpctypes "github.com/eris-ltd/toadserver/Godeps/_workspace/src/github.com/tendermint/tendermint/rpc/types"
	"github.com/eris-ltd/toadserver/Godeps/_workspace/src/github.com/tendermint/tendermint/types"
	"io/ioutil"
	"net/http"
)

type Client interface {
	BlockchainInfo(minHeight int, maxHeight int) (*ctypes.ResultBlockchainInfo, error)
	BroadcastTx(tx types.Tx) (*ctypes.ResultBroadcastTx, error)
	Call(fromAddress []byte, toAddress []byte, data []byte) (*ctypes.ResultCall, error)
	CallCode(fromAddress []byte, code []byte, data []byte) (*ctypes.ResultCall, error)
	DumpConsensusState() (*ctypes.ResultDumpConsensusState, error)
	DumpStorage(address []byte) (*ctypes.ResultDumpStorage, error)
	GenPrivAccount() (*ctypes.ResultGenPrivAccount, error)
	Genesis() (*ctypes.ResultGenesis, error)
	GetAccount(address []byte) (*ctypes.ResultGetAccount, error)
	GetBlock(height int) (*ctypes.ResultGetBlock, error)
	GetName(name string) (*ctypes.ResultGetName, error)
	GetStorage(address []byte, key []byte) (*ctypes.ResultGetStorage, error)
	ListAccounts() (*ctypes.ResultListAccounts, error)
	ListNames() (*ctypes.ResultListNames, error)
	ListUnconfirmedTxs() (*ctypes.ResultListUnconfirmedTxs, error)
	ListValidators() (*ctypes.ResultListValidators, error)
	NetInfo() (*ctypes.ResultNetInfo, error)
	SignTx(tx types.Tx, privAccounts []*acm.PrivAccount) (*ctypes.ResultSignTx, error)
	Status() (*ctypes.ResultStatus, error)
}

func (c *ClientHTTP) BlockchainInfo(minHeight int, maxHeight int) (*ctypes.ResultBlockchainInfo, error) {
	values, err := argsToURLValues([]string{"minHeight", "maxHeight"}, minHeight, maxHeight)
	if err != nil {
		return nil, err
	}
	resp, err := http.PostForm(c.addr+reverseFuncMap["BlockchainInfo"], values)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	response, err := unmarshalCheckResponse(body)
	if err != nil {
		return nil, err
	}
	if response.Result == nil {
		return nil, nil
	}
	result, ok := response.Result.(*ctypes.ResultBlockchainInfo)
	if !ok {
		return nil, fmt.Errorf("response result was wrong type")
	}
	return result, nil
}

func (c *ClientHTTP) BroadcastTx(tx types.Tx) (*ctypes.ResultBroadcastTx, error) {
	values, err := argsToURLValues([]string{"tx"}, tx)
	if err != nil {
		return nil, err
	}
	resp, err := http.PostForm(c.addr+reverseFuncMap["BroadcastTx"], values)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	response, err := unmarshalCheckResponse(body)
	if err != nil {
		return nil, err
	}
	if response.Result == nil {
		return nil, nil
	}
	result, ok := response.Result.(*ctypes.ResultBroadcastTx)
	if !ok {
		return nil, fmt.Errorf("response result was wrong type")
	}
	return result, nil
}

func (c *ClientHTTP) Call(fromAddress []byte, toAddress []byte, data []byte) (*ctypes.ResultCall, error) {
	values, err := argsToURLValues([]string{"fromAddress", "toAddress", "data"}, fromAddress, toAddress, data)
	if err != nil {
		return nil, err
	}
	resp, err := http.PostForm(c.addr+reverseFuncMap["Call"], values)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	response, err := unmarshalCheckResponse(body)
	if err != nil {
		return nil, err
	}
	if response.Result == nil {
		return nil, nil
	}
	result, ok := response.Result.(*ctypes.ResultCall)
	if !ok {
		return nil, fmt.Errorf("response result was wrong type")
	}
	return result, nil
}

func (c *ClientHTTP) CallCode(fromAddress []byte, code []byte, data []byte) (*ctypes.ResultCall, error) {
	values, err := argsToURLValues([]string{"fromAddress", "code", "data"}, fromAddress, code, data)
	if err != nil {
		return nil, err
	}
	resp, err := http.PostForm(c.addr+reverseFuncMap["CallCode"], values)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	response, err := unmarshalCheckResponse(body)
	if err != nil {
		return nil, err
	}
	if response.Result == nil {
		return nil, nil
	}
	result, ok := response.Result.(*ctypes.ResultCall)
	if !ok {
		return nil, fmt.Errorf("response result was wrong type")
	}
	return result, nil
}

func (c *ClientHTTP) DumpConsensusState() (*ctypes.ResultDumpConsensusState, error) {
	values, err := argsToURLValues(nil)
	if err != nil {
		return nil, err
	}
	resp, err := http.PostForm(c.addr+reverseFuncMap["DumpConsensusState"], values)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	response, err := unmarshalCheckResponse(body)
	if err != nil {
		return nil, err
	}
	if response.Result == nil {
		return nil, nil
	}
	result, ok := response.Result.(*ctypes.ResultDumpConsensusState)
	if !ok {
		return nil, fmt.Errorf("response result was wrong type")
	}
	return result, nil
}

func (c *ClientHTTP) DumpStorage(address []byte) (*ctypes.ResultDumpStorage, error) {
	values, err := argsToURLValues([]string{"address"}, address)
	if err != nil {
		return nil, err
	}
	resp, err := http.PostForm(c.addr+reverseFuncMap["DumpStorage"], values)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	response, err := unmarshalCheckResponse(body)
	if err != nil {
		return nil, err
	}
	if response.Result == nil {
		return nil, nil
	}
	result, ok := response.Result.(*ctypes.ResultDumpStorage)
	if !ok {
		return nil, fmt.Errorf("response result was wrong type")
	}
	return result, nil
}

func (c *ClientHTTP) GenPrivAccount() (*ctypes.ResultGenPrivAccount, error) {
	values, err := argsToURLValues(nil)
	if err != nil {
		return nil, err
	}
	resp, err := http.PostForm(c.addr+reverseFuncMap["GenPrivAccount"], values)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	response, err := unmarshalCheckResponse(body)
	if err != nil {
		return nil, err
	}
	if response.Result == nil {
		return nil, nil
	}
	result, ok := response.Result.(*ctypes.ResultGenPrivAccount)
	if !ok {
		return nil, fmt.Errorf("response result was wrong type")
	}
	return result, nil
}

func (c *ClientHTTP) Genesis() (*ctypes.ResultGenesis, error) {
	values, err := argsToURLValues(nil)
	if err != nil {
		return nil, err
	}
	resp, err := http.PostForm(c.addr+reverseFuncMap["Genesis"], values)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	response, err := unmarshalCheckResponse(body)
	if err != nil {
		return nil, err
	}
	if response.Result == nil {
		return nil, nil
	}
	result, ok := response.Result.(*ctypes.ResultGenesis)
	if !ok {
		return nil, fmt.Errorf("response result was wrong type")
	}
	return result, nil
}

func (c *ClientHTTP) GetAccount(address []byte) (*ctypes.ResultGetAccount, error) {
	values, err := argsToURLValues([]string{"address"}, address)
	if err != nil {
		return nil, err
	}
	resp, err := http.PostForm(c.addr+reverseFuncMap["GetAccount"], values)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	response, err := unmarshalCheckResponse(body)
	if err != nil {
		return nil, err
	}
	if response.Result == nil {
		return nil, nil
	}
	result, ok := response.Result.(*ctypes.ResultGetAccount)
	if !ok {
		return nil, fmt.Errorf("response result was wrong type")
	}
	return result, nil
}

func (c *ClientHTTP) GetBlock(height int) (*ctypes.ResultGetBlock, error) {
	values, err := argsToURLValues([]string{"height"}, height)
	if err != nil {
		return nil, err
	}
	resp, err := http.PostForm(c.addr+reverseFuncMap["GetBlock"], values)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	response, err := unmarshalCheckResponse(body)
	if err != nil {
		return nil, err
	}
	if response.Result == nil {
		return nil, nil
	}
	result, ok := response.Result.(*ctypes.ResultGetBlock)
	if !ok {
		return nil, fmt.Errorf("response result was wrong type")
	}
	return result, nil
}

func (c *ClientHTTP) GetName(name string) (*ctypes.ResultGetName, error) {
	values, err := argsToURLValues([]string{"name"}, name)
	if err != nil {
		return nil, err
	}
	resp, err := http.PostForm(c.addr+reverseFuncMap["GetName"], values)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	response, err := unmarshalCheckResponse(body)
	if err != nil {
		return nil, err
	}
	if response.Result == nil {
		return nil, nil
	}
	result, ok := response.Result.(*ctypes.ResultGetName)
	if !ok {
		return nil, fmt.Errorf("response result was wrong type")
	}
	return result, nil
}

func (c *ClientHTTP) GetStorage(address []byte, key []byte) (*ctypes.ResultGetStorage, error) {
	values, err := argsToURLValues([]string{"address", "key"}, address, key)
	if err != nil {
		return nil, err
	}
	resp, err := http.PostForm(c.addr+reverseFuncMap["GetStorage"], values)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	response, err := unmarshalCheckResponse(body)
	if err != nil {
		return nil, err
	}
	if response.Result == nil {
		return nil, nil
	}
	result, ok := response.Result.(*ctypes.ResultGetStorage)
	if !ok {
		return nil, fmt.Errorf("response result was wrong type")
	}
	return result, nil
}

func (c *ClientHTTP) ListAccounts() (*ctypes.ResultListAccounts, error) {
	values, err := argsToURLValues(nil)
	if err != nil {
		return nil, err
	}
	resp, err := http.PostForm(c.addr+reverseFuncMap["ListAccounts"], values)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	response, err := unmarshalCheckResponse(body)
	if err != nil {
		return nil, err
	}
	if response.Result == nil {
		return nil, nil
	}
	result, ok := response.Result.(*ctypes.ResultListAccounts)
	if !ok {
		return nil, fmt.Errorf("response result was wrong type")
	}
	return result, nil
}

func (c *ClientHTTP) ListNames() (*ctypes.ResultListNames, error) {
	values, err := argsToURLValues(nil)
	if err != nil {
		return nil, err
	}
	resp, err := http.PostForm(c.addr+reverseFuncMap["ListNames"], values)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	response, err := unmarshalCheckResponse(body)
	if err != nil {
		return nil, err
	}
	if response.Result == nil {
		return nil, nil
	}
	result, ok := response.Result.(*ctypes.ResultListNames)
	if !ok {
		return nil, fmt.Errorf("response result was wrong type")
	}
	return result, nil
}

func (c *ClientHTTP) ListUnconfirmedTxs() (*ctypes.ResultListUnconfirmedTxs, error) {
	values, err := argsToURLValues(nil)
	if err != nil {
		return nil, err
	}
	resp, err := http.PostForm(c.addr+reverseFuncMap["ListUnconfirmedTxs"], values)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	response, err := unmarshalCheckResponse(body)
	if err != nil {
		return nil, err
	}
	if response.Result == nil {
		return nil, nil
	}
	result, ok := response.Result.(*ctypes.ResultListUnconfirmedTxs)
	if !ok {
		return nil, fmt.Errorf("response result was wrong type")
	}
	return result, nil
}

func (c *ClientHTTP) ListValidators() (*ctypes.ResultListValidators, error) {
	values, err := argsToURLValues(nil)
	if err != nil {
		return nil, err
	}
	resp, err := http.PostForm(c.addr+reverseFuncMap["ListValidators"], values)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	response, err := unmarshalCheckResponse(body)
	if err != nil {
		return nil, err
	}
	if response.Result == nil {
		return nil, nil
	}
	result, ok := response.Result.(*ctypes.ResultListValidators)
	if !ok {
		return nil, fmt.Errorf("response result was wrong type")
	}
	return result, nil
}

func (c *ClientHTTP) NetInfo() (*ctypes.ResultNetInfo, error) {
	values, err := argsToURLValues(nil)
	if err != nil {
		return nil, err
	}
	resp, err := http.PostForm(c.addr+reverseFuncMap["NetInfo"], values)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	response, err := unmarshalCheckResponse(body)
	if err != nil {
		return nil, err
	}
	if response.Result == nil {
		return nil, nil
	}
	result, ok := response.Result.(*ctypes.ResultNetInfo)
	if !ok {
		return nil, fmt.Errorf("response result was wrong type")
	}
	return result, nil
}

func (c *ClientHTTP) SignTx(tx types.Tx, privAccounts []*acm.PrivAccount) (*ctypes.ResultSignTx, error) {
	values, err := argsToURLValues([]string{"tx", "privAccounts"}, tx, privAccounts)
	if err != nil {
		return nil, err
	}
	resp, err := http.PostForm(c.addr+reverseFuncMap["SignTx"], values)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	response, err := unmarshalCheckResponse(body)
	if err != nil {
		return nil, err
	}
	if response.Result == nil {
		return nil, nil
	}
	result, ok := response.Result.(*ctypes.ResultSignTx)
	if !ok {
		return nil, fmt.Errorf("response result was wrong type")
	}
	return result, nil
}

func (c *ClientHTTP) Status() (*ctypes.ResultStatus, error) {
	values, err := argsToURLValues(nil)
	if err != nil {
		return nil, err
	}
	resp, err := http.PostForm(c.addr+reverseFuncMap["Status"], values)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	response, err := unmarshalCheckResponse(body)
	if err != nil {
		return nil, err
	}
	if response.Result == nil {
		return nil, nil
	}
	result, ok := response.Result.(*ctypes.ResultStatus)
	if !ok {
		return nil, fmt.Errorf("response result was wrong type")
	}
	return result, nil
}

func (c *ClientJSON) BlockchainInfo(minHeight int, maxHeight int) (*ctypes.ResultBlockchainInfo, error) {
	request := rpctypes.RPCRequest{
		JSONRPC: "2.0",
		Method:  reverseFuncMap["BlockchainInfo"],
		Params:  []interface{}{minHeight, maxHeight},
		ID:      "",
	}
	body, err := c.RequestResponse(request)
	if err != nil {
		return nil, err
	}
	response, err := unmarshalCheckResponse(body)
	if err != nil {
		return nil, err
	}
	if response.Result == nil {
		return nil, nil
	}
	result, ok := response.Result.(*ctypes.ResultBlockchainInfo)
	if !ok {
		return nil, fmt.Errorf("response result was wrong type")
	}
	return result, nil
}

func (c *ClientJSON) BroadcastTx(tx types.Tx) (*ctypes.ResultBroadcastTx, error) {
	request := rpctypes.RPCRequest{
		JSONRPC: "2.0",
		Method:  reverseFuncMap["BroadcastTx"],
		Params:  []interface{}{tx},
		ID:      "",
	}
	body, err := c.RequestResponse(request)
	if err != nil {
		return nil, err
	}
	response, err := unmarshalCheckResponse(body)
	if err != nil {
		return nil, err
	}
	if response.Result == nil {
		return nil, nil
	}
	result, ok := response.Result.(*ctypes.ResultBroadcastTx)
	if !ok {
		return nil, fmt.Errorf("response result was wrong type")
	}
	return result, nil
}

func (c *ClientJSON) Call(fromAddress []byte, toAddress []byte, data []byte) (*ctypes.ResultCall, error) {
	request := rpctypes.RPCRequest{
		JSONRPC: "2.0",
		Method:  reverseFuncMap["Call"],
		Params:  []interface{}{fromAddress, toAddress, data},
		ID:      "",
	}
	body, err := c.RequestResponse(request)
	if err != nil {
		return nil, err
	}
	response, err := unmarshalCheckResponse(body)
	if err != nil {
		return nil, err
	}
	if response.Result == nil {
		return nil, nil
	}
	result, ok := response.Result.(*ctypes.ResultCall)
	if !ok {
		return nil, fmt.Errorf("response result was wrong type")
	}
	return result, nil
}

func (c *ClientJSON) CallCode(fromAddress []byte, code []byte, data []byte) (*ctypes.ResultCall, error) {
	request := rpctypes.RPCRequest{
		JSONRPC: "2.0",
		Method:  reverseFuncMap["CallCode"],
		Params:  []interface{}{fromAddress, code, data},
		ID:      "",
	}
	body, err := c.RequestResponse(request)
	if err != nil {
		return nil, err
	}
	response, err := unmarshalCheckResponse(body)
	if err != nil {
		return nil, err
	}
	if response.Result == nil {
		return nil, nil
	}
	result, ok := response.Result.(*ctypes.ResultCall)
	if !ok {
		return nil, fmt.Errorf("response result was wrong type")
	}
	return result, nil
}

func (c *ClientJSON) DumpConsensusState() (*ctypes.ResultDumpConsensusState, error) {
	request := rpctypes.RPCRequest{
		JSONRPC: "2.0",
		Method:  reverseFuncMap["DumpConsensusState"],
		Params:  []interface{}{},
		ID:      "",
	}
	body, err := c.RequestResponse(request)
	if err != nil {
		return nil, err
	}
	response, err := unmarshalCheckResponse(body)
	if err != nil {
		return nil, err
	}
	if response.Result == nil {
		return nil, nil
	}
	result, ok := response.Result.(*ctypes.ResultDumpConsensusState)
	if !ok {
		return nil, fmt.Errorf("response result was wrong type")
	}
	return result, nil
}

func (c *ClientJSON) DumpStorage(address []byte) (*ctypes.ResultDumpStorage, error) {
	request := rpctypes.RPCRequest{
		JSONRPC: "2.0",
		Method:  reverseFuncMap["DumpStorage"],
		Params:  []interface{}{address},
		ID:      "",
	}
	body, err := c.RequestResponse(request)
	if err != nil {
		return nil, err
	}
	response, err := unmarshalCheckResponse(body)
	if err != nil {
		return nil, err
	}
	if response.Result == nil {
		return nil, nil
	}
	result, ok := response.Result.(*ctypes.ResultDumpStorage)
	if !ok {
		return nil, fmt.Errorf("response result was wrong type")
	}
	return result, nil
}

func (c *ClientJSON) GenPrivAccount() (*ctypes.ResultGenPrivAccount, error) {
	request := rpctypes.RPCRequest{
		JSONRPC: "2.0",
		Method:  reverseFuncMap["GenPrivAccount"],
		Params:  []interface{}{},
		ID:      "",
	}
	body, err := c.RequestResponse(request)
	if err != nil {
		return nil, err
	}
	response, err := unmarshalCheckResponse(body)
	if err != nil {
		return nil, err
	}
	if response.Result == nil {
		return nil, nil
	}
	result, ok := response.Result.(*ctypes.ResultGenPrivAccount)
	if !ok {
		return nil, fmt.Errorf("response result was wrong type")
	}
	return result, nil
}

func (c *ClientJSON) Genesis() (*ctypes.ResultGenesis, error) {
	request := rpctypes.RPCRequest{
		JSONRPC: "2.0",
		Method:  reverseFuncMap["Genesis"],
		Params:  []interface{}{},
		ID:      "",
	}
	body, err := c.RequestResponse(request)
	if err != nil {
		return nil, err
	}
	response, err := unmarshalCheckResponse(body)
	if err != nil {
		return nil, err
	}
	if response.Result == nil {
		return nil, nil
	}
	result, ok := response.Result.(*ctypes.ResultGenesis)
	if !ok {
		return nil, fmt.Errorf("response result was wrong type")
	}
	return result, nil
}

func (c *ClientJSON) GetAccount(address []byte) (*ctypes.ResultGetAccount, error) {
	request := rpctypes.RPCRequest{
		JSONRPC: "2.0",
		Method:  reverseFuncMap["GetAccount"],
		Params:  []interface{}{address},
		ID:      "",
	}
	body, err := c.RequestResponse(request)
	if err != nil {
		return nil, err
	}
	response, err := unmarshalCheckResponse(body)
	if err != nil {
		return nil, err
	}
	if response.Result == nil {
		return nil, nil
	}
	result, ok := response.Result.(*ctypes.ResultGetAccount)
	if !ok {
		return nil, fmt.Errorf("response result was wrong type")
	}
	return result, nil
}

func (c *ClientJSON) GetBlock(height int) (*ctypes.ResultGetBlock, error) {
	request := rpctypes.RPCRequest{
		JSONRPC: "2.0",
		Method:  reverseFuncMap["GetBlock"],
		Params:  []interface{}{height},
		ID:      "",
	}
	body, err := c.RequestResponse(request)
	if err != nil {
		return nil, err
	}
	response, err := unmarshalCheckResponse(body)
	if err != nil {
		return nil, err
	}
	if response.Result == nil {
		return nil, nil
	}
	result, ok := response.Result.(*ctypes.ResultGetBlock)
	if !ok {
		return nil, fmt.Errorf("response result was wrong type")
	}
	return result, nil
}

func (c *ClientJSON) GetName(name string) (*ctypes.ResultGetName, error) {
	request := rpctypes.RPCRequest{
		JSONRPC: "2.0",
		Method:  reverseFuncMap["GetName"],
		Params:  []interface{}{name},
		ID:      "",
	}
	body, err := c.RequestResponse(request)
	if err != nil {
		return nil, err
	}
	response, err := unmarshalCheckResponse(body)
	if err != nil {
		return nil, err
	}
	if response.Result == nil {
		return nil, nil
	}
	result, ok := response.Result.(*ctypes.ResultGetName)
	if !ok {
		return nil, fmt.Errorf("response result was wrong type")
	}
	return result, nil
}

func (c *ClientJSON) GetStorage(address []byte, key []byte) (*ctypes.ResultGetStorage, error) {
	request := rpctypes.RPCRequest{
		JSONRPC: "2.0",
		Method:  reverseFuncMap["GetStorage"],
		Params:  []interface{}{address, key},
		ID:      "",
	}
	body, err := c.RequestResponse(request)
	if err != nil {
		return nil, err
	}
	response, err := unmarshalCheckResponse(body)
	if err != nil {
		return nil, err
	}
	if response.Result == nil {
		return nil, nil
	}
	result, ok := response.Result.(*ctypes.ResultGetStorage)
	if !ok {
		return nil, fmt.Errorf("response result was wrong type")
	}
	return result, nil
}

func (c *ClientJSON) ListAccounts() (*ctypes.ResultListAccounts, error) {
	request := rpctypes.RPCRequest{
		JSONRPC: "2.0",
		Method:  reverseFuncMap["ListAccounts"],
		Params:  []interface{}{},
		ID:      "",
	}
	body, err := c.RequestResponse(request)
	if err != nil {
		return nil, err
	}
	response, err := unmarshalCheckResponse(body)
	if err != nil {
		return nil, err
	}
	if response.Result == nil {
		return nil, nil
	}
	result, ok := response.Result.(*ctypes.ResultListAccounts)
	if !ok {
		return nil, fmt.Errorf("response result was wrong type")
	}
	return result, nil
}

func (c *ClientJSON) ListNames() (*ctypes.ResultListNames, error) {
	request := rpctypes.RPCRequest{
		JSONRPC: "2.0",
		Method:  reverseFuncMap["ListNames"],
		Params:  []interface{}{},
		ID:      "",
	}
	body, err := c.RequestResponse(request)
	if err != nil {
		return nil, err
	}
	response, err := unmarshalCheckResponse(body)
	if err != nil {
		return nil, err
	}
	if response.Result == nil {
		return nil, nil
	}
	result, ok := response.Result.(*ctypes.ResultListNames)
	if !ok {
		return nil, fmt.Errorf("response result was wrong type")
	}
	return result, nil
}

func (c *ClientJSON) ListUnconfirmedTxs() (*ctypes.ResultListUnconfirmedTxs, error) {
	request := rpctypes.RPCRequest{
		JSONRPC: "2.0",
		Method:  reverseFuncMap["ListUnconfirmedTxs"],
		Params:  []interface{}{},
		ID:      "",
	}
	body, err := c.RequestResponse(request)
	if err != nil {
		return nil, err
	}
	response, err := unmarshalCheckResponse(body)
	if err != nil {
		return nil, err
	}
	if response.Result == nil {
		return nil, nil
	}
	result, ok := response.Result.(*ctypes.ResultListUnconfirmedTxs)
	if !ok {
		return nil, fmt.Errorf("response result was wrong type")
	}
	return result, nil
}

func (c *ClientJSON) ListValidators() (*ctypes.ResultListValidators, error) {
	request := rpctypes.RPCRequest{
		JSONRPC: "2.0",
		Method:  reverseFuncMap["ListValidators"],
		Params:  []interface{}{},
		ID:      "",
	}
	body, err := c.RequestResponse(request)
	if err != nil {
		return nil, err
	}
	response, err := unmarshalCheckResponse(body)
	if err != nil {
		return nil, err
	}
	if response.Result == nil {
		return nil, nil
	}
	result, ok := response.Result.(*ctypes.ResultListValidators)
	if !ok {
		return nil, fmt.Errorf("response result was wrong type")
	}
	return result, nil
}

func (c *ClientJSON) NetInfo() (*ctypes.ResultNetInfo, error) {
	request := rpctypes.RPCRequest{
		JSONRPC: "2.0",
		Method:  reverseFuncMap["NetInfo"],
		Params:  []interface{}{},
		ID:      "",
	}
	body, err := c.RequestResponse(request)
	if err != nil {
		return nil, err
	}
	response, err := unmarshalCheckResponse(body)
	if err != nil {
		return nil, err
	}
	if response.Result == nil {
		return nil, nil
	}
	result, ok := response.Result.(*ctypes.ResultNetInfo)
	if !ok {
		return nil, fmt.Errorf("response result was wrong type")
	}
	return result, nil
}

func (c *ClientJSON) SignTx(tx types.Tx, privAccounts []*acm.PrivAccount) (*ctypes.ResultSignTx, error) {
	request := rpctypes.RPCRequest{
		JSONRPC: "2.0",
		Method:  reverseFuncMap["SignTx"],
		Params:  []interface{}{tx, privAccounts},
		ID:      "",
	}
	body, err := c.RequestResponse(request)
	if err != nil {
		return nil, err
	}
	response, err := unmarshalCheckResponse(body)
	if err != nil {
		return nil, err
	}
	if response.Result == nil {
		return nil, nil
	}
	result, ok := response.Result.(*ctypes.ResultSignTx)
	if !ok {
		return nil, fmt.Errorf("response result was wrong type")
	}
	return result, nil
}

func (c *ClientJSON) Status() (*ctypes.ResultStatus, error) {
	request := rpctypes.RPCRequest{
		JSONRPC: "2.0",
		Method:  reverseFuncMap["Status"],
		Params:  []interface{}{},
		ID:      "",
	}
	body, err := c.RequestResponse(request)
	if err != nil {
		return nil, err
	}
	response, err := unmarshalCheckResponse(body)
	if err != nil {
		return nil, err
	}
	if response.Result == nil {
		return nil, nil
	}
	result, ok := response.Result.(*ctypes.ResultStatus)
	if !ok {
		return nil, fmt.Errorf("response result was wrong type")
	}
	return result, nil
}

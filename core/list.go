package core

import (
	"os"
	"fmt"

	cclient "github.com/tendermint/tendermint/rpc/core_client"
//	log "github.com/Sirupsen/logrus"
)

// need to think of the name spacing carefullly
func ListAllTheNames() ([]string, error) {
	//GET chain URL
	url := os.Getenv("MINTX_NODE_ADDR")

	c := cclient.NewClient(url, "HTTP")

	res, err := c.ListNames()
	if err != nil {
		return nil, fmt.Errorf("error calling c.LN(): %v", err)
	}

	names := make([]string, len(res.Names))

	for i, name := range res.Names{
		names[i] = name.Name
	}
	return names, nil
}

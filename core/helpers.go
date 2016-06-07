package core

import (
	"fmt"
	"net/http"
	"strings"

	log "github.com/Sirupsen/logrus"
)

func CacheHashAll(ipAddrs, hash string) error {

	//IPaddrs, _ := getTheNames() -> TODO use some decentralized DNS
	IPs := strings.Split(ipAddrs, ",")
	for _, ip := range IPs {
		endpoint := UrlHandler(ip, "11113", "/cacheHash?hash=", hash)
		log.WithField("=>", endpoint).Warn("Pinning hash to remote IPFS/toadserver node:")
		_, err := http.Post(endpoint, "", nil)
		if err != nil {
			// don't return an err, there are other IPs to hit
			log.WithField("=>", endpoint).Warn("error making post request to:")
			log.Error(err)
			continue
		}
	}
	return nil
}

func UrlHandler(host, port, endpoint, arg string) string {
	if arg == "" {
		return fmt.Sprintf("http://%s:%s%s", host, port, endpoint)
	} else {
		return fmt.Sprintf("http://%s:%s%s%s", host, port, endpoint, arg)
	}
}

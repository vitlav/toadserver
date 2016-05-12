package main

import (
	log "github.com/Sirupsen/logrus"
	logger "github.com/eris-ltd/common/go/log"

	"github.com/spf13/cobra"
)

var (
	DefaultChainAddr = "http://0.0.0.0:46657"
	ChainFlag        string

	DefaultToadHost = "localhost"
	ToadHost        string

	DefaultToadPort = "11113"
	ToadPort        string
)

func main() {
	BuildToadserverCommand()
	Toadserver.PersistentPreRun = before
	//Toadserver.PersistentPostRun = after
	Toadserver.Execute()
}

var Toadserver = &cobra.Command{
	Use:   "toadserver",
	Short: "",
	Long:  "",
	Run:   func(cmd *cobra.Command, args []string) { cmd.Help() },
}

func BuildToadserverCommand() {
	Toadserver.AddCommand(startCmd)
	Toadserver.AddCommand(putCmd)
	Toadserver.AddCommand(getCmd)

	addToadserverFlags()
}

func addToadserverFlags() {
	startCmd.Flags().StringVarP(&ToadHost, "host", "", DefaultToadHost, "specify the host for toadserver to run on")
	startCmd.Flags().StringVarP(&ToadPort, "port", "", DefaultToadPort, "specify the port for toadserver to run on")

	putCmd.Flags().StringVarP(&ChainFlag, "node-addr", "", DefaultChainAddr, "specify the chain to use")
	putCmd.Flags().StringVarP(&ToadHost, "host", "", DefaultToadHost, "specify the host")
	putCmd.Flags().StringVarP(&ToadPort, "port", "", DefaultToadPort, "specify the port")

	getCmd.Flags().StringVarP(&ChainFlag, "node-addr", "", DefaultChainAddr, "specify the chain to use")
	getCmd.Flags().StringVarP(&ToadHost, "host", "", DefaultToadHost, "specify the host")
	getCmd.Flags().StringVarP(&ToadPort, "port", "", DefaultToadPort, "specify the port that toadserver was started on")
}

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "",
	Long:  "",
	Run:   startServer,
}

var putCmd = &cobra.Command{
	Use:   "put",
	Short: "",
	Long:  "",
	Run:   putFiles,
}

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "",
	Long:  "",
	Run:   getFiles,
}

func before(cmd *cobra.Command, args []string) {
	log.SetFormatter(logger.ErisFormatter{})
}

package main

import (
	log "github.com/Sirupsen/logrus"
	logger "github.com/eris-ltd/common/go/log"
	"github.com/eris-ltd/toadserver/version"

	"github.com/spf13/cobra"
)

const VERSION = version.VERSION

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

// TODO helpers
var Toadserver = &cobra.Command{
	Use:   "toadserver",
	Short: "A simple IPFS & chain based download server.",
	Long: `A simple IPFS & chain base download server.
The toadserver indexes ipfs hashes on the name registry of
a running eris-db chain.` + "\nVersion:\n" + VERSION,
	Run: func(cmd *cobra.Command, args []string) { cmd.Help() },
}

func BuildToadserverCommand() {
	Toadserver.AddCommand(startCmd)
	Toadserver.AddCommand(putCmd)
	Toadserver.AddCommand(getCmd)
	Toadserver.AddCommand(lsCmd)

	addToadserverFlags()
}

// TODO deduplicate flags; add persistence?
// expose any os.Getenv as a flag.
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
	Short: "Start the toadserver.",
	Long:  "Start the toadserver.",
	Run:   startServer,
}

var putCmd = &cobra.Command{
	Use:   "put FILE",
	Short: "Add a file to the toadserver.",
	Long:  "Add a file to the toadserver.",
	Run:   putFiles,
}

var getCmd = &cobra.Command{
	Use:   "get FILE",
	Short: "Retrieve a file from the toadserver.",
	Long:  "Retrieve a file from the toadserver.",
	Run:   getFiles,
}

var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List files registered with the toadserver.",
	Long:  "List files registered with the toadserver.",
	Run:   listFiles,
}

func before(cmd *cobra.Command, args []string) {
	log.SetFormatter(logger.ErisFormatter{})
}

package main

import (
	log "github.com/Sirupsen/logrus"
	logger "github.com/eris-ltd/common/go/log"

	"github.com/spf13/cobra"
)

func main() {
	BuildToadserverCommand()
	//Toadserver.PersistenPreRun = before
	//Toadserver.PersistenPostRun = after
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

package cmd

import (
	as "github.com/nghiatrandev/cob-as-prt/as"
	asdb "github.com/nghiatrandev/cob-as-prt/dataservice"
	"github.com/spf13/cobra"
	"log"
)

var cmdServer = &cobra.Command{
	Use:   "server",
	Short: "Launch the HTTP server",
	Run: func(cmd *cobra.Command, args []string) {

		client, err := as.NewClient(asCfg)
		if err != nil {
			log.Fatal(err)
		}

		asdb.NewDataServiceAerospike(client, asCfg.Namespace)
	},
}

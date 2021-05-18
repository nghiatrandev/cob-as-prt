package cmd

import (
	as "github.com/nghiatrandev/cob-as-prt/as"
	asdb "github.com/nghiatrandev/cob-as-prt/dataservice"
	"github.com/nghiatrandev/cob-as-prt/factory"
	"github.com/spf13/cobra"
	"log"
)

var cmdServer = &cobra.Command{
	Use:   "server",
	Short: "Launch the HTTP server",
	Run: func(cmd *cobra.Command, args []string) {

		f := factory.NewDefaultFactory(cfg)
		hander := f.BuildHandler()
		hander.Routes()

		asClient, err := as.NewClient(cfg.AsConf)
		if err != nil {
			log.Fatal(err)
		}

		asdb.NewDataServiceAerospike(asClient, cfg.AsConf.Namespace)
	},
}

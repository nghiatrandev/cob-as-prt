package cmd

import (
	"github.com/nghiatrandev/sample_project/factory"
	"github.com/spf13/cobra"
	"log"
	"net/http"
	"strconv"
)

var cmdServer = &cobra.Command{
	Use:   "server",
	Short: "Launch the HTTP server",
	Run: func(cmd *cobra.Command, args []string) {

		f := factory.NewDefaultFactory(cfg)
		handler := f.BuildHandler()
		mux := handler.Routes()

		//asClient, err := as.NewClient(cfg.AsConf)
		//if err != nil {
		//	log.Fatal(err)
		//}
		//
		//asdb.NewDataServiceAerospike(asClient, cfg.AsConf.Namespace)

		log.Println("Start server")
		http.ListenAndServe(":"+strconv.Itoa(cfg.HttpConfig.Port), mux)

	},
}

package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"path"

	"github.com/RedisLabs/cf-redislabs-broker/redislabs"
	"github.com/RedisLabs/cf-redislabs-broker/redislabs/config"
	"github.com/RedisLabs/cf-redislabs-broker/redislabs/instancebinders"
	"github.com/RedisLabs/cf-redislabs-broker/redislabs/instancecreators"
	"github.com/RedisLabs/cf-redislabs-broker/redislabs/persisters"
	"github.com/pivotal-cf/brokerapi"
	"github.com/pivotal-golang/lager"
)

var (
	localPersisterPath string
	brokerStateRoot    string
	brokerConfigPath   string
)

func init() {
	flag.StringVar(&brokerConfigPath, "c", "", "Configuration File")
	flag.StringVar(&brokerStateRoot, "s", os.Getenv("HOME"), "State Root Folder")

	flag.Parse()

	localPersisterPath = path.Join(brokerStateRoot, ".redislabs-broker", "state.json")
}

func main() {
	brokerLogger := lager.NewLogger("redislabs-service-broker")
	brokerLogger.RegisterSink(lager.NewWriterSink(os.Stdout, lager.DEBUG))
	brokerLogger.RegisterSink(lager.NewWriterSink(os.Stderr, lager.ERROR))

	if brokerConfigPath == "" {
		brokerLogger.Error("No config file specified", nil)
		os.Exit(1)
	}

	brokerLogger.Info("Using config file: " + brokerConfigPath)

	conf, err := config.LoadFromFile(brokerConfigPath)
	if err != nil {
		brokerLogger.Error("Failed to load the config file", err, lager.Data{
			"broker-config-path": brokerConfigPath,
		})
		os.Exit(1)
	}

	serviceBroker := redislabs.NewServiceBroker(
		instancecreators.NewDefault(conf, brokerLogger),
		instancebinders.NewDefault(conf, brokerLogger),
		persisters.NewLocalPersister(localPersisterPath),
		conf,
		brokerLogger,
	)

	brokerUser := os.Getenv("BROKER_USERNAME")
	if brokerUser == "" {
		brokerLogger.Error("no broker username given", nil)
		os.Exit(1)
	}
	brokerPass := os.Getenv("BROKER_PASSWORD")
	if brokerPass == "" {
		brokerLogger.Error("no broker password given", nil)
		os.Exit(1)
	}

	credentials := brokerapi.BrokerCredentials{
		Username: brokerUser,
		Password: brokerPass,
	}

	brokerAPI := brokerapi.New(serviceBroker, brokerLogger, credentials)
	http.Handle("/", brokerAPI)
	brokerLogger.Info("Listening for requests", lager.Data{
		"port": conf.ServiceBroker.Port,
	})
	err = http.ListenAndServe(fmt.Sprintf(":%d", conf.ServiceBroker.Port), nil)
	if err != nil {
		brokerLogger.Error("Failed to start the server", err)
	}
}

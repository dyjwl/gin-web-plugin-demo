/*
Copyright © 2022 NAME HERE <zhaosir_1993@163.com>
*/
package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/dyjwl/gin-web-plugin-demo/cmd/crontab"
	gindemo "github.com/dyjwl/gin-web-plugin-demo/cmd/gin-demo"
	"github.com/dyjwl/gin-web-plugin-demo/configs"
	"github.com/dyjwl/gin-web-plugin-demo/internal/store"
	"github.com/dyjwl/gin-web-plugin-demo/internal/store/mysql"
	"github.com/dyjwl/gin-web-plugin-demo/internal/store/pg"
	"github.com/dyjwl/gin-web-plugin-demo/internal/store/sqlite"
	"github.com/dyjwl/gin-web-plugin-demo/pkg/cache"
	"github.com/dyjwl/gin-web-plugin-demo/pkg/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// rootCmd represents the base command when called without any subcommands
var (
	// Used for flags.
	cfgFile     string
	userLicense string

	rootCmd = &cobra.Command{
		Use:   "gin-web-plugin-demo",
		Short: "A demo of use gin to develop a web application",
		Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
		// Uncomment the following line if your bare application
		// has an action associated with it:
		Run: func(cmd *cobra.Command, args []string) {
			log.InitLog()
			cache.SyncRedis = cache.NewClient(configs.Config.Redis)
			cache.InitRedSync(cache.SyncRedis)
			log.Info("config", zap.Any("content", configs.Config))
			switch configs.Config.Database.Dialect {
			case "mysql":
				storeIns, _ := mysql.GetMysqlFactoryOr(&configs.Config.Database)
				store.SetClient(storeIns)
			case "postgres":
				storeIns, _ := pg.GetPgFactoryOr(&configs.Config.Database)
				store.SetClient(storeIns)
			case "sqlite":
				storeIns, _ := sqlite.GetSqliteFactoryOr(&configs.Config.Database)
				store.SetClient(storeIns)
			}

			crontab.Run()
			log.Info("app run,config: ", zap.Any("config", configs.Config))
			gindemo.StartServer()
		},
	}
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {

	rootCmd.PersistentFlags().StringVar(&cfgFile,
		"config", "", "config file (default is $HOME/.cobra.yaml)")
	cobra.OnInitialize(initConfig)
	// rootCmd.AddCommand(addCmd)
	// rootCmd.AddCommand(initCmd)
}

func initConfig() {

	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)
		// Search config in home directory with name ".cobra" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".cobra")
	}
	viper.SetEnvPrefix("GENDEMO")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Error: failed to read configuration file(%s): %v\n", cfgFile, err)
		os.Exit(1)
	}
	if err := viper.Unmarshal(&configs.Config); err != nil {
		panic("unmarsh config error")
	}
}

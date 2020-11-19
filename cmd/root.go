/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/varluffy/gindemo/internal/app"
	"github.com/varluffy/gindemo/internal/app/initialize"
	"github.com/varluffy/gindemo/pkg/logger"
	"github.com/varluffy/gindemo/pkg/util"
	"os"

	"github.com/spf13/viper"
)

var (
	cfgFile string
	// 可以通过编译的方式指定版本号：go build -ldflags "-X main.VERSION=x.x.x"
	Version = "1.0.0"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gindemo",
	Short: "gin 框架demo",
	Long:  "gin 框架demo",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
	Run: func(cmd *cobra.Command, args []string) {
		initialize.InitLogger()
		logger.SetVersion(Version)
		logger.SetTraceIDFunc(util.NewTraceID)
		ctx := logger.NewTraceIDContext(context.Background(), util.NewTraceID())
		logger.Info(ctx, "root cmd run")
		err := app.Run(ctx)
		if err != nil {
			panic(err)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "config/dev.yaml", "config file")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	viper.SetConfigFile(cfgFile)
	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
		return
	}
	panic("viper.ReadInConfig error:" + cfgFile)
}

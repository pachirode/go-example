package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
	"github.com/spf13/viper"
)

func base() {
	var rootCmd = &cobra.Command{
		Use:   "base",
		Short: "Base corba command test",
		Long:  "Base corba command test built with cobra",
		Run: func(cmd *cobra.Command, args []string) {
			println("Base")
		},
	}

	addChildCommand(rootCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Printf("%v", err)
		os.Exit(1)
	}

}

func addChildCommand(rootCmd *cobra.Command) {
	var childCmd = &cobra.Command{
		Use:   "child",
		Short: "Child cobra",
		Long:  "Child cobra command build with cobra",
		Run: func(cmd *cobra.Command, args []string) {
			println("child")
		},
		Args: cobra.MaximumNArgs(10),
	}

	rootCmd.AddCommand(childCmd)
}

func addFlags(rootCmd *cobra.Command) {
	//持久标志
	var Verbose bool
	rootCmd.PersistentFlags().BoolVar(&Verbose, "verbose", false, "verbose")

	// 本地标志
	var local bool
	rootCmd.Flags().BoolVar(&local, "local", false, "local")

	// 必选标志
	var must string
	rootCmd.Flags().StringVarP(&must, "must", "r", "", "(required)")
	rootCmd.MarkFlagRequired("must")

}

func setViper(rootCmd *cobra.Command) {
	var tmp string
	rootCmd.PersistentFlags().StringVar(&tmp, "tmp", "tmp", "tmp")
	viper.BindPFlag("tmp", rootCmd.PersistentFlags().Lookup("tmp"))
}

var cfgFile string

func initViper(rootCmd *cobra.Command) {
	cobra.OnInitialize(initConfig)
	rootCmd.Flags().StringVarP(&cfgFile, "config", "c", "", "config file")
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.AddConfigPath(".")
		viper.SetConfigName(".cobra")
	}

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Can't read config:", err)
		os.Exit(1)
	}
}

func setHelpMsg(rootCmd *cobra.Command) {
	// 控制 help 命令
	rootCmd.SetHelpCommand(&cobra.Command{
		Use:    "help",
		Short:  "Custom help command",
		Hidden: true,
		Run: func(cmd *cobra.Command, args []string) {
			println("Custom help command")
		},
	})

	// 控制 -h/--help
	rootCmd.SetHelpFunc(func(command *cobra.Command, strings []string) {
		println(strings)
	})
}

func genDocs(rootCmd *cobra.Command) {
	file, err := os.Create("./docs/command.md")
	if err != nil {
		os.Exit(1)
	}
	defer file.Close()
	if err = doc.GenMarkdown(rootCmd, file); err != nil {
		os.Exit(1)
	}
}

func main() {
	base()
}

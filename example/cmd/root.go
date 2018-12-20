package cmd

import (
        "flag"
        "fmt"
        "github.com/go-kit/kit/log"
        "github.com/spf13/pflag"
        "github.com/spf13/viper"
        "os"
        "github.com/spf13/cobra"
)

func init() {Init()}

var (
        defaultConfig *viper.Viper
        logger  *log.Logger
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
        Use:   "example",
        Short: "cookiecutter.project_short_description",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
        if err := rootCmd.Execute(); err != nil {
                fmt.Println(err)
                os.Exit(1)
        }
}

func GetConfig() *viper.Viper {
        return defaultConfig
}

func GetLogger() *log.Logger {
        return logger
}

func Init() {
        {
                logger = log.NewLogfmtLogger(os.Stdout)
                logger = log.With(logger, "ts", log.DefaultTimestampUTC)
                logger = log.With(logger, "caller", log.DefaultCaller)
        }
        {
                defaultConfig = viper.New()
                defaultConfig.AutomaticEnv()
                defaultConfig.AddConfigPath(os.Getenv("$HOME")) // name of config file (without extension)
                defaultConfig.AddConfigPath(".")
                defaultConfig.SetEnvPrefix("cookiecutter.app_name")
                defaultConfig.SetDefault("viper_config_name", "cookiecutter.viper_config_name")
                defaultConfig.SetDefault("full_name", "cookiecutter.full_name")
                defaultConfig.SetDefault("github_username", "cookiecutter.github_username")
                defaultConfig.SetDefault("app_name", "cookiecutter.app_name")
                defaultConfig.SetDefault("project_short_description", "cookiecutter.project_short_description")
                defaultConfig.SetDefault("docker_hub_username", "cookiecutter.docker_hub_username")
                defaultConfig.SetDefault("docker_image", "cookiecutter.docker_image")
                defaultConfig.SetDefault("docker_build_image_version", "cookiecutter.docker_build_image_version")
                defaultConfig.SetDefault("json_logs", "cookiecutter.json_logs")
                defaultConfig.SetDefault("log_level", "cookiecutter.log_level")
                defaultConfig.SetDefault("grpc_port", "cookiecutter.grpc_port")
                defaultConfig.SetDefault("http_port", "cookiecutter.http_port")
        }

        // If a config file is found, read it in.
        if err := defaultConfig.ReadInConfig(); err != nil {
                logger.Info("failed to read config file, writing defaults...")
                if err := defaultConfig.WriteConfigAs("cookiecutter.viper_config_name"+".yaml"); err != nil {
                        logger.Fatal("failed to write config")
                        os.Exit(1)
                }

        } else {
                logger.Info("Using config file-->", defaultConfig.ConfigFileUsed())
                if err := defaultConfig.WriteConfig(); err != nil {
                        logger.Fatal("failed to write config file")
                        os.Exit(1)
                }
        }
}
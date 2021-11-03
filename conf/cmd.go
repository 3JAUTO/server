package conf

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	ConfPath string
	Env      string
	ServerID int32

	Addr      string
	LogFile   string
	LogLevel  string
	DBConnStr string
	DBName    string
	IsDebug   bool
)

func InitFlags(cmd *cobra.Command) {
	flags := cmd.PersistentFlags()
	flags.StringVar(&Env, "env", "", "server environment. e.g. dev|local|qas|prd")
	flags.Int32Var(&ServerID, "serverid", 0, "server id")
	flags.StringVar(&Addr, "addr", "", "http server listen address")
	flags.StringVar(&LogFile, "logfile", "", "log output file. stdout|stderr|file")
	flags.StringVar(&LogLevel, "loglevel", "", "log output level")
}

func Init(cmd *cobra.Command) error {

	if err := initConfigFile(); err != nil {
		return err
	}
	flags := cmd.PersistentFlags()
	if Addr == "" {
		Addr = envConf.Addr
	}
	if LogFile == "" {
		LogFile = envConf.LogFile
	}
	if LogLevel == "" {
		LogLevel = envConf.LogLevel
	}
	dbConnStr := fmt.Sprintf("%s:%s@tcp(%s)/?parseTime=true&loc=Local", envConf.DBUser,
		envConf.DBPasswd, envConf.DBUrl)
	DBName = fmt.Sprintf("%s_%s_%d", envConf.DBName, Env, ServerID)
	flags.StringVar(&DBConnStr, "db", dbConnStr, "database connection string")
	flags.BoolVar(&IsDebug, "debug", true, "if debug")
	return nil
}

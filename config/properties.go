package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

func ConfigureProperties() {
	SetupConfigSource("config", "properties")

	if err := viper.ReadInConfig(); err != nil {
		log.Panic("Error cargando propiedades")
	}
}

func SetupConfigSource(defaultConfigFolderPath string, configFileToUse string) {
	viper.AddConfigPath(defaultConfigFolderPath)
	viper.AutomaticEnv()
	viper.SetConfigType("yaml")

	viper.SetConfigName(configFileToUse)
}

func GetPresenter(weekday string) string {
	return viper.GetString(fmt.Sprintf("days.%s.presenter", weekday))
}

func GetAlternate(weekday string) string {
	return viper.GetString(fmt.Sprintf("days.%s.alternate", weekday))
}

func GetJiraUrl() string {
	return viper.GetString("jira-url")
}

func GetMeetUrl() string {
	return viper.GetString("meet-url")
}

func GetWebhookUrl() string {
	return viper.GetString("webhook-url")
}

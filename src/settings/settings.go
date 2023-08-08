package settings

import (
	"encoding/json"
	"example/raylib-game/src/mines"
	"os"
)

// The settings filepath
const settingsFilepath string = "data/settings.json"

// The settings structure
type Settings struct {
	SettingsPath string `json:"settings_path"`
	Theme        string `json:"theme"`
	Width        int    `json:"width"`
	Height       int    `json:"height"`
	Bombs        int    `json:"bombs"`
}

// The default settings for the app
var defaultSettings = Settings{
	SettingsPath: "data/settings.json",
	Theme:        "",
	Width:        30,
	Height:       16,
	Bombs:        50,
}

// Load the settings from a file
func (settings *Settings) LoadFromFile(defaultTheme string) error {
	data, err := os.ReadFile(settingsFilepath)
	if err != nil {
		defaultSettings.Theme = defaultTheme
		*settings = defaultSettings
		data, err := json.MarshalIndent(settings, "  ", "")
		if err != nil {
			return err
		}

		if err := os.WriteFile(settingsFilepath, data, 0644); err != nil {
			return err
		}
	}

	return json.Unmarshal(data, settings)
}

// Write the changed settigs into the file
func (settings *Settings) WriteToFile(newSettings Settings) error {
	// Try to check if the settings are actually valid
	_, err := mines.GenerateBoard(newSettings.Width, newSettings.Height, newSettings.Bombs)
	if err != nil {
		return err
	}

	// Marshall the json data and write it to the file
	data, err := json.MarshalIndent(newSettings, "", "")
	if err != nil {
		return err
	}

	if err = os.WriteFile(settingsFilepath, data, 0644); err != nil {
		return err
	}

	// Make the newsettings be the new settings
	*settings = newSettings
	return nil
}

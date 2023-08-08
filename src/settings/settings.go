package settings

import (
	"encoding/json"
	"example/raylib-game/src/mines"
	"os"
	"path/filepath"
)

// The settings structure
type Settings struct {
	Theme        string `json:"theme"`
	ThemePath    string `json:"-"`
	SettingsPath string `json:"-"`
	Width        int    `json:"width"`
	Height       int    `json:"height"`
	Bombs        int    `json:"bombs"`
}

// The default settings for the app
var defaultSettings = Settings{
	Theme:  "",
	Width:  30,
	Height: 16,
	Bombs:  50,
}

// Load the settings from a file
func (settings *Settings) LoadFromFile(defaultTheme string) error {
	cfgDir, err := os.UserConfigDir()
	if err != nil {
		return err
	}

	path := filepath.Join(cfgDir, "gomines", "settings.json")
	settings.SettingsPath = path
	settings.ThemePath = filepath.Join(cfgDir, "gomines", "themes")
	data, err := os.ReadFile(path)
	if err == nil {
		return json.Unmarshal(data, settings)
	}

	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return err
	}

	defaultSettings.Theme = defaultTheme
	*settings = defaultSettings
	data, err = json.MarshalIndent(settings, "", "  ")
	if err != nil {
		return err
	}

	if err := os.WriteFile(path, data, 0644); err != nil {
		return err
	}

	return nil
}

// Write the changed settigs into the file
func (settings *Settings) WriteToFile(newSettings Settings) error {
	// Try to check if the settings are actually valid
	if _, err := mines.GenerateBoard(newSettings.Width, newSettings.Height, newSettings.Bombs); err != nil {
		return err
	}

	// Marshall the json data and write it to the file
	data, err := json.MarshalIndent(newSettings, "", "  ")
	if err != nil {
		return err
	}

	cfgDir, err := os.UserConfigDir()
	if err != nil {
		return err
	}

	path := filepath.Join(cfgDir, "gomines", "settings.json")
	settings.SettingsPath = path
	if err = os.WriteFile(path, data, 0644); err != nil {
		return err
	}

	// Make the newsettings be the new settings
	newSettings.ThemePath = settings.ThemePath
	newSettings.SettingsPath = settings.SettingsPath
	*settings = newSettings
	return nil
}

package settings

import (
	"encoding/json"
	"errors"
	"example/raylib-game/src/mines"
	"io/ioutil"
	"os"
)

// The settings filepath
const filepath string = "data/settings.json"

// The settings structure
type Settings struct {
	SettingsPath string `json:"settings_path"`
	Theme        string `json:"theme"`
	Width        int    `json:"width"`
	Height       int    `json:"height"`
	Bombs        int    `json:"bombs"`
}

// The default settings for the app
var DefaultSettings Settings = Settings{
	SettingsPath: "data/settings.json",
	Theme:        "hello_kitty",
	Width:        30,
	Height:       16,
	Bombs:        50,
}

// Load the settings from a file
func (settings *Settings) LoadFromFile() error {
	// Try to open the jsonFile
	jsonFile, err := os.Open(filepath)
	if err != nil {
		// Try to create the file because it doesn't exist
		jsonFile, err = os.Create(filepath)

		if err != nil {
			return errors.New("couldn't open the file for writing")
		}
	}

	// Close the file after usage
	defer jsonFile.Close()

	// Try to load the json from the file
	byteValue, _ := ioutil.ReadAll(jsonFile)
	err = json.Unmarshal(byteValue, settings)
	if err != nil {
		// Write the default settings into the file
		*settings = DefaultSettings
		jsonData, _ := json.MarshalIndent(settings, "", "")
		jsonFile.Write(jsonData)
	}

	// Return no errors
	return nil
}

// Write the changed settigs into the file
func (settings *Settings) WriteToFile(newSettings Settings) error {
	// Try to check if the settings are actually valid
	_, err := mines.GenerateBoard(newSettings.Width, newSettings.Height, newSettings.Bombs)
	if err != nil {
		return err
	}

	// Open the jsonFile
	jsonFile, err := os.OpenFile(filepath, os.O_WRONLY, os.ModeAppend)
	if err != nil {
		return errors.New("couldn't open the file")
	}

	// Marshall the json data and write it to the file
	jsonData, _ := json.MarshalIndent(newSettings, "", "")
	_, err = jsonFile.Write(jsonData)
	if err != nil {
		return errors.New("couldn't write the json data to the file")
	}

	// Make the newsettings be the new settings
	settings = &newSettings

	// Return no errors
	return nil
}

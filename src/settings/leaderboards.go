package settings

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
)

const leaderboardsFilepath string = "data/scores.json"

// Hold the data for one entry
type Entry struct {
	Date  int64  `json:"date"`
	Name  string `json:"name"`
	Score int    `json:"score"`
	Time  int    `json:"time"`
}

// Hold the data for multiple entries
type Scores struct {
	Entries []Entry
}

// Load the entires data from a file
func (scores *Scores) LoadFromFile() error {
	// Try to open the jsonFile
	jsonFile, err := os.Open(leaderboardsFilepath)
	if err != nil {
		// Try to create the file because it doesn't exist
		jsonFile, err = os.Create(leaderboardsFilepath)

		if err != nil {
			return errors.New("couldn't open the file for writing")
		}
	}

	// Close the file after usage
	defer jsonFile.Close()

	// Try to load the json from the file
	byteValue, _ := ioutil.ReadAll(jsonFile)
	err = json.Unmarshal(byteValue, scores)
	if err != nil {
		// Write the default leaderboard into the file
		jsonData, _ := json.MarshalIndent(scores, "", "")
		if _, err = jsonFile.Write(jsonData); err != nil {
			return err
		}
	}

	// Return no errors
	return nil
}

// Should the new score be at the scoreboard
func (scores *Scores) CanItBeInTheScoreboard(time int) (bool, int) {
	// If our time is bigger than the last entry, we can be in the scoreboard
	if time > scores.Entries[len(scores.Entries)-1].Time {
		return false, 0
	}

	// Get the place where the element should be
	var place int
	for pos, entry := range scores.Entries {
		if time < entry.Time {
			place = pos
			break
		}
	}

	// We can be in the scoreboard, for example in the 3rd place
	return true, place
}

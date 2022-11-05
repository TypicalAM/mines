package settings

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"time"
)

const (
	leaderboardsFilepath string = "data/scores.json"
	Beginner             int    = iota
	Intermediate
	Expert
	Custom
)

// Hold the data for one entry
type Entry struct {
	Date        int64  `json:"date"`
	Name        string `json:"name"`
	Time        int    `json:"time"`
	BoardWidth  int    `json:"board_width"`
	BoardHeight int    `json:"board_height"`
	BoardMines  int    `json:"board_mines"`
}

// Hold the data for multiple entries
type Scores struct {
	Entries []Entry
}

var defaultScores Scores = Scores{
	Entries: []Entry{
		{
			Date:        time.Now().Unix(),
			Name:        "Liam",
			Time:        220,
			BoardWidth:  30,
			BoardHeight: 16,
			BoardMines:  21,
		}, {
			Date:        time.Now().Unix() + 2137,
			Name:        "Noah",
			Time:        230,
			BoardWidth:  30,
			BoardHeight: 16,
			BoardMines:  21,
		}, {
			Date:        time.Now().Unix() + 123,
			Name:        "Oliver",
			Time:        240,
			BoardWidth:  30,
			BoardHeight: 16,
			BoardMines:  21,
		}, {
			Date:        time.Now().Unix() + 123,
			Name:        "Elijah",
			Time:        250,
			BoardWidth:  30,
			BoardHeight: 16,
			BoardMines:  21,
		}, {
			Date:        time.Now().Unix() + 123,
			Name:        "James",
			Time:        260,
			BoardWidth:  30,
			BoardHeight: 16,
			BoardMines:  21,
		}, {
			Date:        time.Now().Unix(),
			Name:        "Liam",
			Time:        200,
			BoardWidth:  16,
			BoardHeight: 16,
			BoardMines:  15,
		}, {
			Date:        time.Now().Unix() + 2137,
			Name:        "Noah",
			Time:        210,
			BoardWidth:  16,
			BoardHeight: 16,
			BoardMines:  15,
		}, {
			Date:        time.Now().Unix() + 123,
			Name:        "Oliver",
			Time:        220,
			BoardWidth:  16,
			BoardHeight: 16,
			BoardMines:  15,
		}, {
			Date:        time.Now().Unix() + 123,
			Name:        "Elijah",
			Time:        230,
			BoardWidth:  16,
			BoardHeight: 16,
			BoardMines:  15,
		}, {
			Date:        time.Now().Unix() + 123,
			Name:        "James",
			Time:        240,
			BoardWidth:  16,
			BoardHeight: 16,
			BoardMines:  15,
		}, {
			Date:        time.Now().Unix(),
			Name:        "Liam",
			Time:        140,
			BoardWidth:  8,
			BoardHeight: 8,
			BoardMines:  15,
		}, {
			Date:        time.Now().Unix() + 2137,
			Name:        "Noah",
			Time:        150,
			BoardWidth:  8,
			BoardHeight: 8,
			BoardMines:  15,
		}, {
			Date:        time.Now().Unix() + 123,
			Name:        "Oliver",
			Time:        160,
			BoardWidth:  8,
			BoardHeight: 8,
			BoardMines:  15,
		}, {
			Date:        time.Now().Unix() + 123,
			Name:        "Elijah",
			Time:        170,
			BoardWidth:  8,
			BoardHeight: 8,
			BoardMines:  15,
		}, {
			Date:        time.Now().Unix() + 123,
			Name:        "James",
			Time:        180,
			BoardWidth:  8,
			BoardHeight: 8,
			BoardMines:  15,
		},
	},
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
		// If we cant the scores, write the default ones to the file
		jsonData, _ := json.MarshalIndent(defaultScores, "", "")
		if _, err = jsonFile.Write(jsonData); err != nil {
			return err
		}
	}

	// Return no errors
	return nil
}

// Write the changed scores into the file
func (scores *Scores) WriteToFile() error {
	// Open the jsonFile
	jsonFile, err := os.OpenFile(leaderboardsFilepath, os.O_WRONLY, os.ModeAppend)
	if err != nil {
		return errors.New("couldn't open the file")
	}

	// Marshall the json data and write it to the file
	jsonData, _ := json.MarshalIndent(scores, "", "")
	_, err = jsonFile.Write(jsonData)
	if err != nil {
		return errors.New("couldn't write the json data to the file")
	}

	// Return no errors
	return nil
}

// Should the new score be at the scoreboard
func (scores *Scores) CanItBeInTheScoreboard(settings Settings, time int) (bool, int) {
	// Filter the scores by the settings
	var filter int

	if settings.Width == 8 && settings.Height == 8 && settings.Bombs == 15 {
		filter = Beginner
	} else if settings.Width == 16 && settings.Height == 16 && settings.Bombs == 15 {
		filter = Intermediate
	} else if settings.Width == 30 && settings.Height == 16 && settings.Bombs == 21 {
		filter = Expert
	} else {
		filter = Custom
	}

	entries := scores.FilterScores(filter)

	// If our time is bigger than the last entry, we can be in the scoreboard
	if time > entries[len(entries)-1].Time {
		return false, 0
	}

	// Get the place where the element should be
	var place int
	for pos, entry := range entries {
		if time <= entry.Time {
			place = pos
			break
		}
	}

	// We can be in the scoreboard, for example in the 3rd place
	return true, place
}

// Insert the new score to the scoreboard
func (scores *Scores) InsertNewScore(settings Settings, newScoreName string, gameTime int, scoreboardPlace int) error {
	// Add the new entry to the scoreboard and shfit its contents
	scores.Entries = append(scores.Entries[:scoreboardPlace+1], scores.Entries[scoreboardPlace:]...)
	scores.Entries[scoreboardPlace] = Entry{
		Date:        time.Now().Unix(),
		Name:        newScoreName,
		Time:        gameTime,
		BoardWidth:  settings.Width,
		BoardHeight: settings.Height,
		BoardMines:  settings.Bombs,
	}

	// Save the new score table
	return scores.WriteToFile()
}

// Filter the scores according to the category that they're in
func (scores *Scores) FilterScores(category int) []Entry {
	var entries []Entry

	// Loop over entries and add them to the shown ones
	for _, entry := range scores.Entries {
		switch category {
		case Beginner:
			if entry.BoardWidth == 8 && entry.BoardHeight == 8 && entry.BoardMines == 15 {
				entries = append(entries, entry)
			}
		case Intermediate:
			if entry.BoardWidth == 16 && entry.BoardHeight == 16 && entry.BoardMines == 15 {
				entries = append(entries, entry)
			}
		case Expert:
			if entry.BoardWidth == 30 && entry.BoardHeight == 16 && entry.BoardMines == 21 {
				entries = append(entries, entry)
			}
		case Custom:
			if entry.BoardWidth != 8 && entry.BoardWidth != 16 && entry.BoardWidth != 30 &&
				entry.BoardHeight != 8 && entry.BoardHeight != 16 &&
				entry.BoardMines != 15 && entry.BoardMines != 21 {
				entries = append(entries, entry)
			}
		}
	}

	return entries
}

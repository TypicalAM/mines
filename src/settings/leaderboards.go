package settings

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"sort"
	"time"
)

const (
	Beginner int = iota
	Intermediate
	Expert
	Custom
)

// Hold the data for one entry
type Entry struct {
	Name        string `json:"name"`
	Date        int64  `json:"date"`
	Time        int    `json:"time"`
	BoardWidth  int    `json:"board_width"`
	BoardHeight int    `json:"board_height"`
	BoardMines  int    `json:"board_mines"`
}

// Hold the data for multiple entries
type Scores struct {
	Entries []Entry
}

var defaultScores = Scores{
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
	cfgDir, err := os.UserConfigDir()
	if err != nil {
		return err
	}

	path := filepath.Join(cfgDir, "gomines", "leaderboards.json")
	data, err := os.ReadFile(path)
	if err == nil {
		return json.Unmarshal(data, scores)
	}

	if err = os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return err
	}

	// Write the default scores to the file
	*scores = defaultScores
	jsonData, _ := json.MarshalIndent(defaultScores, "", "  ")
	return os.WriteFile(path, jsonData, 0644)
}

// Write the changed scores into the file
func (scores *Scores) WriteToFile() error {
	cfgDir, err := os.UserConfigDir()
	if err != nil {
		return err
	}

	path := filepath.Join(cfgDir, "gomines", "leaderboards.json")
	data, err := json.MarshalIndent(scores, "", "  ")
	if err != nil {
		return errors.New("couldn't marshal the json data")
	}

	if err := os.WriteFile(path, data, 0644); err != nil {
		if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
			return err
		}

		if err := os.WriteFile(path, data, 0644); err != nil {
			return err
		}
	}

	return nil
}

// Should the new score be at the scoreboard
func (scores *Scores) CanItBeInTheScoreboard(filter int, time int) (bool, int) {
	// Filter the scores by the settings
	entries := scores.FilterScores(filter)

	// If there are no entries we are the first one
	if len(entries) == 0 {
		return true, 0
	}

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
func (scores *Scores) InsertNewScore(settings Settings, newScoreName string, gameTime int) {
	// Make the entry for the score
	entry := Entry{
		Date:        time.Now().Unix(),
		Name:        newScoreName,
		Time:        gameTime,
		BoardWidth:  settings.Width,
		BoardHeight: settings.Height,
		BoardMines:  settings.Bombs,
	}

	// Add the new entry to the scoreboard
	scores.Entries = append(scores.Entries, entry)
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
			if (entry.BoardWidth != 8 || entry.BoardHeight != 8 || entry.BoardMines != 15) &&
				(entry.BoardWidth != 16 || entry.BoardHeight != 16 || entry.BoardMines != 15) &&
				(entry.BoardWidth != 30 || entry.BoardHeight != 16 || entry.BoardMines != 21) {
				entries = append(entries, entry)
			}
		}
	}

	// Sort the slice by time
	sort.Slice(entries, func(i, j int) bool {
		return entries[i].Time < entries[j].Time
	})

	// Return the sorted entries
	return entries
}

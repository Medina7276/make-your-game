package model

type Scoreboard struct {
	Username string `json:"username"`
	Score    int    `json:"score"`
	Time     string `json:"time"`
}

func (store *PlayerDB) CreateUser(player *Scoreboard) error {
	defer store.Close()
	_, err := store.Exec(`INSERT INTO Scoreboard (username, score, time) VALUES (?, ?, ?)`,
		player.Username, player.Score, player.Time)
	return err
}

func (store *PlayerDB) GetUser(username string) (*Scoreboard, error) {
	defer store.Close()

	player := &Scoreboard{}

	row := store.QueryRow(`
	SELECT username, score, time 
	FROM Scoreboard
	WHERE username = ? 
	ORDER BY 
	score DESC`, username)

	err := row.Scan(&player.Username, &player.Score, &player.Time)
	return player, err
}

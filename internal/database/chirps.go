package database

import "os"

type Chirp struct {
	ID       int    `json:"id"`
	AuthorID int    `json:"author_id"`
	Body     string `json:"body"`
}

// GetChirps returns all chirps in the database
func (db *DB) GetChirps() ([]Chirp, error) {
	dbStructure, err := db.loadDB()
	if err != nil {
		return nil, err
	}
	chirps := make([]Chirp, 0, len(dbStructure.Chirps))
	for _, chirp := range dbStructure.Chirps {
		chirps = append(chirps, chirp)
	}
	return chirps, nil
}

// GetChirp gets a single Chirp using an ID
func (db *DB) GetChirp(id int) (Chirp, error) {
	dbStructure, err := db.loadDB()
	if err != nil {
		return Chirp{}, err
	}

	chirp, ok := dbStructure.Chirps[id]
	if !ok {
		return Chirp{}, os.ErrNotExist
	}
	return chirp, nil
}

// CreateChirp creates a new chirp and saves it to disk
func (db *DB) CreateChirp(body string, userID int) (Chirp, error) {
	dbStructure, err := db.loadDB()
	if err != nil {
		return Chirp{}, err
	}
	id := len(dbStructure.Chirps) + 1
	newChirp := Chirp{
		ID:       id,
		AuthorID: userID,
		Body:     body,
	}
	dbStructure.Chirps[id] = newChirp

	err = db.writeDB(dbStructure)
	if err != nil {
		return Chirp{}, err
	}
	return newChirp, nil
}

func (db *DB) DeleteChirp(id int) error {
	dbStructure, err := db.loadDB()
	if err != nil {
		return err
	}

	chirp, ok := dbStructure.Chirps[id]
	if !ok {
		return os.ErrNotExist
	}

	delete(dbStructure.Chirps, chirp.ID)
	err = db.writeDB(dbStructure)
	if err != nil {
		return err
	}
	return nil
}

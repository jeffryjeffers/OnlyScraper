package database

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"path"
)

type Database struct {
	DB *sqlx.DB
}

var medias = `
CREATE TABLE IF NOT EXISTS medias(
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    media_id INTEGER NOT NULL,
    post_id INTEGER NOT NULL,
    link VARCHAR,
    dir VARCHAR,
    filename VARCHAR);`

type Media struct {
	ID       int    `db:"id"`
	MediaId  int    `db:"media_id"`
	PostId   int    `db:"post_id"`
	Link     string `db:"link"`
	Dir      string `db:"dir"`
	Filename string `db:"filename"`
}

func CreateAndConnect(path string) (*Database, error) {
	db, err := sqlx.Connect("sqlite3", path)
	if err != nil {
		return nil, err
	}
	db.MustExec(medias)
	return &Database{DB: db}, nil
}
func (db *Database) InsertMedia(mediaId int, postId int, link string, dir string, filename string) error {
	tx := db.DB.MustBegin()
	_, err := tx.NamedExec("INSERT INTO medias (media_id, post_id, link, dir, filename) VALUES (:media_id, :post_id, :link, :dir, :filename)", &Media{
		MediaId:  mediaId,
		PostId:   postId,
		Link:     link,
		Dir:      dir,
		Filename: filename,
	})
	if err != nil {
		return err
	}
	err = tx.Commit()
	return err
}

func (db *Database) GetExisting() ([]Media, error) {
	var exists []Media
	var medias []Media
	err := db.DB.Select(&medias, "SELECT * from medias")
	if err != nil {
		return nil, err
	}
	for _, media := range medias {
		_, err := os.Stat(path.Join(media.Dir, media.Filename))
		if !os.IsNotExist(err) {
			exists = append(exists, media)
		}
	}
	return exists, nil
}

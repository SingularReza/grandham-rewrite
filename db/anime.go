package db

import (
	"fmt"
	"path/filepath"
	"strings"

	metadata "github.com/SingularReza/grandham-rewrite/metadata"
	_ "github.com/mattn/go-sqlite3"
)

// CreateAnimeEntry - creates an anime entry in ANIME table bote: change this to return (int64, err)
func CreateAnimeEntry(animeData metadata.AnimeMedia, animeFolderID string) int64 {
	statement, err := database.Prepare(`INSERT INTO ANIME (anime_id, anime_title_romaji,
										anime_title_english, anime_cover, anime_banner, anime_format,
										anime_episodes, anime_ep_duration, anime_genres) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`)
	checkErr(err)

	var genres string
	for _, genre := range animeData.Genres {
		genres += genre + ","
	}
	genres = strings.TrimLeft(genres, ",")

	result, err := statement.Exec(animeData.ID, animeData.Title.Romaji, animeData.Title.English,
		filepath.Base(animeData.CoverImage.Large), filepath.Base(animeData.BannerImage), animeData.Format,
		animeData.Episodes, animeData.EpisodeDuration, genres)
	fmt.Print(animeData.ID)
	checkErr(err)

	animeID, err := result.LastInsertId()
	checkErr(err)

	infoID := addAnimeExtraInfo(animeData.Description, animeData.StartDate, animeData.EndDate, animeID)

	fmt.Printf("animeid: %d, animeinfoid: %d\n", animeID, infoID)

	return animeID
}

func addAnimeExtraInfo(description string, startDate metadata.AnimeDate, endDate metadata.AnimeDate, animeID int64) int64 {
	statement, err := database.Prepare(`INSERT INTO ANIMEINFO (anime_desc, anime_startyear, anime_startmonth,
										anime_startday, anime_endyear, anime_endmonth, anime_endday, anime_id)
										VALUES (?, ?, ?, ?, ?, ?, ?, ?)`)
	checkErr(err)

	result, err := statement.Exec(description, startDate.Year, startDate.Month, startDate.Day,
		endDate.Year, endDate.Month, endDate.Day, animeID)
	checkErr(err)

	infoID, err := result.LastInsertId()
	checkErr(err)

	fmt.Println(infoID)

	return infoID
}

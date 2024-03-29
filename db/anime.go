package db

import (
	"fmt"
	"path/filepath"
	"strings"

	metadata "github.com/SingularReza/grandham-rewrite/metadata"
	_ "github.com/mattn/go-sqlite3"
	"google.golang.org/api/drive/v3"
)

// CreateAnimeEntry - creates an anime entry in ANIME table bote: change this to return (int64, err)
// note: get library_id and anime_id relation into a different table(different libraries may have same anime)
func CreateAnimeEntry(animeData metadata.AnimeMedia, animeFolderID string, libraryID int64) int64 {
	statement, err := database.Prepare(`INSERT OR IGNORE INTO ANIME (anime_id, anime_folderid, anime_title_romaji,
										anime_title_english, anime_cover, anime_banner, anime_format,
										anime_episodes, anime_ep_duration, anime_genres, library_id)
										VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`)
	checkErr(err)

	var genres string
	for _, genre := range animeData.Genres {
		genres += genre + ","
	}
	genres = strings.TrimLeft(genres, ",")

	result, err := statement.Exec(animeData.ID, animeFolderID, animeData.Title.Romaji, animeData.Title.English,
		filepath.Base(animeData.CoverImage.Large), filepath.Base(animeData.BannerImage), animeData.Format,
		animeData.Episodes, animeData.EpisodeDuration, genres, libraryID)
	fmt.Print(animeData.ID)
	checkErr(err)

	animeID, err := result.LastInsertId()
	checkErr(err)

	infoID := addAnimeExtraInfo(animeData.Description, animeData.StartDate, animeData.EndDate, animeID)
	fmt.Printf("animeid: %d, animeinfoid: %d\n", animeID, infoID)

	return animeID
}

func addAnimeExtraInfo(description string, startDate metadata.AnimeDate, endDate metadata.AnimeDate, animeID int64) int64 {
	statement, err := database.Prepare(`INSERT OR IGNORE INTO ANIMEINFO (anime_desc, anime_startyear, anime_startmonth,
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

// AddAnimeEpisodes - Reads video file metadata and records them
func AddAnimeEpisode(animeID int64, fileData *drive.File) int64 {
	statement, err := database.Prepare(`INSERT OR IGNORE INTO ANIMEEPISODES (ep_id, ep_title,
		ep_duration, ep_size, ep_height, ep_width, anime_id)
		VALUES (?, ?, ?, ?, ?, ?, ?)`)
	checkErr(err)
	videoMetadata := fileData.VideoMediaMetadata
	result, err := statement.Exec(fileData.Id, fileData.Name, videoMetadata.DurationMillis,
		fileData.Size, videoMetadata.Height, videoMetadata.Width, animeID)
	checkErr(err)

	infoID, err := result.LastInsertId()
	checkErr(err)

	fmt.Println(infoID)

	return infoID
}

// GetAnimeInfo - retrieves anime info from ANIME and ANIMEINFO tables
func GetAnimeInfo(animeID int) metadata.AnimeMedia {
	query := `SELECT ANIME.anime_id, anime_folderid, anime_title_romaji,
			  anime_title_english, anime_cover, anime_banner, anime_format,
			  anime_episodes, anime_ep_duration, anime_genres, anime_desc, anime_startyear,
			  anime_startmonth, anime_startday, anime_endyear, anime_endmonth, anime_endday
			  FROM ANIME LEFT JOIN ANIMEINFO ON ANIME.anime_id = ANIMEINFO.anime_id
			  WHERE ANIME.anime_id = ?`

	var genreString string

	animeData := metadata.AnimeMedia{}
	row := database.QueryRow(query, animeID)
	err := row.Scan(&animeData.ID, &animeData.FolderID, &animeData.Title.Romaji, &animeData.Title.English,
		&animeData.CoverImage.Large, &animeData.BannerImage, &animeData.Format, &animeData.Episodes,
		&animeData.EpisodeDuration, &genreString, &animeData.Description, &animeData.StartDate.Year,
		&animeData.StartDate.Month, &animeData.StartDate.Day, &animeData.EndDate.Year, &animeData.EndDate.Month,
		&animeData.EndDate.Day)
	checkErr(err)

	animeData.Genres = strings.Split(genreString, ",")

	episodeList := []metadata.AnimeEpisode{}
	rows, err := database.Query(`SELECT ep_id, ep_title,
			ep_duration, ep_size, ep_height, ep_width FROM ANIMEEPISODES
			WHERE anime_id = ?`, animeData.ID)
	for rows.Next() {
		episode := metadata.AnimeEpisode{}
		rows.Scan(&episode.FileID, &episode.FileName, &episode.FileDuration, &episode.FileSize,
			&episode.FileHeight, &episode.FileWidth)
		episodeList = append(episodeList, episode)
	}

	animeData.EpisodeList = episodeList

	return animeData
}

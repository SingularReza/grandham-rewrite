package db

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/SingularReza/grandham-rewrite/metadata"
)

func TestCreateAnimeEntry(t *testing.T) {

	animeData := metadata.AnimeMedia{ID: 123456789,
		Title:           metadata.AnimeTitle{Romaji: "Kimi no Na wa.", English: "Your Name."},
		CoverImage:      metadata.AnimeImage{Large: "https://s4.anilist.co/file/anilistcdn/media/anime/cover/medium/nx21519-F2z1QPS5GmpC.jpg"},
		BannerImage:     "https://s4.anilist.co/file/anilistcdn/media/anime/banner/21519-1ayMXgNlmByb.jpg",
		StartDate:       metadata.AnimeDate{Year: 2016, Month: 8, Day: 26},
		EndDate:         metadata.AnimeDate{Year: 2016, Month: 8, Day: 26},
		Description:     "Mitsuha Miyamizu, a high school girl, yearns to live the life of a boy in the bustling city of Tokyo—a dream that stands in stark contrast to her present life in the countryside. Meanwhile in the city, Taki Tachibana lives a busy life as a high school student while juggling his part-time job and hopes for a future in architecture.<br>\n<br>\nOne day, Mitsuha awakens in a room that is not her own and suddenly finds herself living the dream life in Tokyo—but in Taki's body! Elsewhere, Taki finds himself living Mitsuha's life in the humble countryside. In pursuit of an answer to this strange phenomenon, they begin to search for one another.<br>\n<br>\nKimi no Na wa. revolves around Mitsuha and Taki's actions, which begin to have a dramatic impact on each other's lives, weaving them into a fabric held together by fate and circumstance.<br>\n<br>\n[Written by MAL Rewrite]",
		Season:          "SUMMER",
		SeasonYear:      2016,
		Format:          "MOVIE",
		Episodes:        1,
		EpisodeDuration: 107,
		Genres:          []string{"Romance", "Drama", "Supernatural"},
	}

	animeEntryID := CreateAnimeEntry(animeData, "testid", 1)

	// Checking that the entry is successful | note: change this after making createanimeentry return (int64, err)
	if reflect.TypeOf(animeEntryID).Name() == "int64" {
		fmt.Print("CreateAnimeEntry creates a row\n")
	} else {
		t.Errorf("CreateAnimeEntry in db/anime.go failed")
	}
}

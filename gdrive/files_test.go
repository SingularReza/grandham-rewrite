package scan

import (
	"testing"

	cmp "github.com/google/go-cmp/cmp"
	drive "google.golang.org/api/drive/v3"
)

func TestGetItemsList(t *testing.T) {
	// change this later to a stable folder
	folderID := "1xx-dgZHjhtcYVc_ooeW5b9GN5pQGiMY5"

    itemList := GetItemsList(folderID)

	expected := []*drive.File{
								{Name:"Azhar",Id:"1_zUBkR3c-ZzN2W3d34zZbMbY9c0FPQ8S"},
								{Name:"Awe!",Id:"1ZViUXVbcAR_f_5DTPqQybHw2BiWjKhH5"},
								{Name:"Awarapan",Id:"121WWf-pX-ADLi9WegT6F38IP5exjvZFx"},
								{Name:"Avengers: Infinity War",Id:"1CUnEonUetdbze0IK5Z6b2Y_oetAqD1gt"},
								{Name:"Avengers: Endgame",Id:"103iv1yDxwpNrCZhmWHoRwu286zaFOgyU"},
								{Name:"Avengers: Age of Ultron",Id:"1BLkwtKA5uTZyIi_0-cI8W1cvDAxA0ue2"},
								{Name:"Avatar",Id:"11gzOivVoobL7_0cArW_xwvW54y6i5MBR"},
								{Name:"Asoka - Ashoka the Great",Id:"12-dceFZHUjQOuCElKgDK9jB1kVqGxFEW"},
								{Name:"As I Was Moving Ahead Occasionally I Saw Brief Glimpses of Beauty",Id:"1XvKXvLo7GaNNtunz410PtGsxlRatbcl-"},
								{Name:"Arrival",Id:"1mnvTdgbRQRcMisltA93VAlpABSmeaGtY"},
								{Name:"Arjun Reddy",Id:"14y9clWRCz7Bpg5pI_9mNdKS6ip1g_oAn"},
								{Name:"Aquaman",Id:"1uleJhJIdQ0he4mAUpqurBHqT3DGGN-HA"},
								{Name:"Apur Sansar",Id:"1KXafAcsuheG9go_aGtUEpQQtJBfv7iqQ"},
								{Name:"Aparajito",Id:"1e0hLpfi-ShRHWJEMUARnlFApII1ddZAb"},
								{Name:"Ant-Man",Id:"121gSmcwO-XcuAocWmzyh9s-AHXluTBZc"},
								{Name:"Andhadhun",Id:"1_go1HEPYaXihTzaoXdq892TG7SjSWz3b"},
								{Name:"Andaz Apna Apna",Id:"12rQ3T-4Yw1WxrVRqmQ4-l2AoRhmnQ67H"},
								{Name:"Anbe Sivam",Id:"1Br8Tj_0dnhYb5E63jKxasf4sZXaLgW6U"},
								{Name:"American History X",Id:"125_KX0NhFK7r7XTj-38Z-yBoJ6RX4p_T"},
								{Name:"Amar Akbar Anthony",Id:"1LGmBvl5flKT4adJ_Y8Xqz3kyqNXWIW6C"},
								{Name:"All About Lily Chou-Chou",Id:"1BW45mNvtD2fXy1DwJIGGlMT5M3Wu5lg9"},
								{Name:"Aliens",Id:"1CvmfDWHTaLXnagPY7gTVNOKetX2KeuFs"},
								{Name:"Alien Resurrection",Id:"1uxH2UyOuwKmqyuLCooHqas7QxnpwMH5f"},
								{Name:"Alien 3",Id:"1-yHsveyrfYgZEvm-8ogLR2iVjf_1uobs"},
								{Name:"Alien (Diector's Cut)",Id:"1d9sB3v_zsFeqH2a3PO2l8vB83k5mjrPL"},
								{Name:"Alien - Covenant",Id:"1cWTSgxwjxAFae3_MAHP_pVSnH3MgRH50"},
								{Name:"Alien",Id:"1uuR9X8nqpQpzUmib8c_tLbzJmFMuw9aA"},
								{Name:"Akira",Id:"1SMOJ8CMfOFET0QBYOcMi09EbvCa9688t"},
								{Name:"Agent Sai Srinivasa Athreya",Id:"1qB8XcXH8bVFm_gZZDdXhoHzLFYlRoXFZ"},
								{Name:"After Earth",Id:"1uzuSlDXyiihkyDXW-GRmQlkyTg_LrTry"},
								{Name:"Aata",Id:"11QOC0IpEMbLijfmfnFlq4kGMP1iWw4JU"},
								{Name:"Aashiqui 2",Id:"1NdE6HXDQMVjKbPuadlRTnC9XzfJ0z_Jj"},
								{Name:"Aashiqui",Id:"1LI08lMzrIvNfFVUDdcG5CQNmjyRwEF6v"},
								{Name:"A Wednesday",Id:"1Z2xFg9aVifWkpdEIO1q5W5hc4WUnZPLK"},
								{Name:"A Taxi Driver (korean)",Id:"1CgrepgP0WdIL-hlkORX-hmbeDd2tYNvI"},
								{Name:"A Tale of Two Sisters",Id:"1lYGqCu7JRnQdOykpIy-PU5XFq3htzy-Q"},
								{Name:"A Quiet Place",Id:"1S2RmmTNTNy-8mxreupXz_6eU1319RyRR"},
								{Name:"A Brighter Summer Day",Id:"1igqp1Bdxb5KGXTN-ROHmMZHwnMO00XbM"},
								{Name:"A Boy and His Atom - The World's Smallest Movie",Id:"1wOUFwpEulVrHGIcAeCk2YmPKM92B8iol"}}

	if !cmp.Equal(itemList, expected)  {
		t.Errorf("handler returned unexpected response")
	}
}

package gmusic

import (
	"strings"
)

type (
	Music struct {
		musicNames MusicNameList
		musicLines MusicLineList
	}
)

func NewMusic() *Music {
	mu := &Music{
		musicNames: make(MusicNameList, 0),
		musicLines: make(MusicLineList, 0),
	}

	//十二音名和位置
	musicNames := []map[string]int{
		{"C": 0},
		{"#C": 1},
		{"bD": 1},
		{"D": 2},
		{"#D": 3},
		{"bE": 3},
		{"E": 4},
		{"F": 5},
		{"#F": 6},
		{"bG": 6},
		{"G": 7},
		{"#G": 8},
		{"bA": 8},
		{"A": 9},
		{"#A": 10},
		{"bB": 10},
		{"B": 11},
	}

	//五线谱高音谱号对应的五条线 + 六条间
	musicLineHighTable := map[int][]string{
		0:  {"bD", "D", "#D"},
		1:  {"bE", "E"},
		2:  {"F", "#F"},
		3:  {"bG", "G", "#G"},
		4:  {"bA", "A", "#A"},
		5:  {"bB", "B"},
		6:  {"C", "#C"},
		7:  {"bD", "D", "#D"},
		8:  {"bE", "E"},
		9:  {"F", "#F"},
		10: {"bG", "G", "#G"},
	}

	//12音名集合
	for index, value := range musicNames {
		musicName := new(MusicName)
		musicName.Index = index

		for name, location := range value {
			musicName.Name = name
			musicName.Location = location
		}

		mu.musicNames = append(mu.musicNames, musicName)
	}

	//五线谱线和间集合
	for musicLineIndex, musicLineNames := range musicLineHighTable {
		musicLine := new(MusicLine)
		musicLine.Index = musicLineIndex
		musicLine.Names = make(MusicNameList, 0)

		for _, name := range musicLineNames {
			for _, musicName := range mu.musicNames {
				if name == musicName.Name {
					musicLine.Names = append(musicLine.Names, musicName)
					break
				}
			}
		}

		mu.musicLines = append(mu.musicLines, musicLine)
	}

	return mu
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 音名规范化
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s Music) MusicNameNormal(musicName string) string {
	if len(musicName) > 1 {
		names := make([]string, 0)
		for index, name := range musicName {
			if index == 0 {
				names = append(names, strings.ToLower(string(name)))
			} else {
				names = append(names, strings.ToUpper(string(name)))
			}
		}

		musicName = strings.Join(names, "")
	} else {
		musicName = strings.ToUpper(musicName)
	}

	return musicName
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 根据音名获取音名索引
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s Music) GetMusicNameByCode(musicName string) *MusicName {
	var result *MusicName
	musicName = s.MusicNameNormal(musicName)

	for _, musicNameItem := range s.musicNames {
		if musicNameItem.Name == musicName {
			result = musicNameItem
			break
		}
	}

	return result
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 根据音名索引获取音名
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s Music) GetMusicNameByIndex(musicNameIndex int) *MusicName {
	var result *MusicName

	for _, musicNameItem := range s.musicNames {
		if musicNameItem.Index == musicNameIndex {
			result = musicNameItem
			break
		}
	}

	return result
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 根据音名位置索引获取音名集合
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s Music) GetMusicNamesByLocation(musicNameLocation int) MusicNameList {
	musicNames := make(MusicNameList, 0)

	for _, musicName := range s.musicNames {
		if musicName.Location == musicNameLocation {
			musicNames = append(musicNames, musicName)
		}
	}

	return musicNames
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 根据五线谱索引获取音名集合
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s Music) GetMusicNamesByLine(musicLineIndex int) MusicNameList {
	musicNames := make(MusicNameList, 0)

	for _, musicLine := range s.musicLines {
		if musicLine.Index == musicLineIndex {
			musicNames = musicLine.Names
			break
		}
	}

	return musicNames
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 获取高音谱号音名五线谱索引
 * ~~~~~~~~~~~~~~~~~~~~~ 11    - A, #A, bA
 * 10                          - G, #G, bG
 * --------------------- 9     - F, #F
 * 8                           - E, bE
 * --------------------- 7     - D, #D, bD
 * 6                           - C, #C
 * --------------------- 5     - B, bB
 * 4                           - A, #A, bA
 * --------------------- 3     - G, #G, bG
 * 2                           - F, #F
 * --------------------- 1     - E, bE
 * 0                           - D, #D, bD
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s Music) GetMusicLinesByMusicName(musicName string) MusicLineList {
	musicLines := make(MusicLineList, 0)

	musicName = s.MusicNameNormal(musicName)

	for _, musicLine := range s.musicLines {
		for _, currentMusicName := range musicLine.Names {
			if currentMusicName.Name == musicName {
				musicLines = append(musicLines, musicLine)
				break
			}
		}
	}

	return musicLines
}

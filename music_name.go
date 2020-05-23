package gmusic

type (
	MusicNameList []*MusicName
	MusicName     struct {
		Name     string //音名
		Index    int    //7个音 + 5对同音异名变化音索引
		Location int    //12音位置索引
	}
)

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 判断音名是否黑键
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s MusicName) IsBlack() bool {
	isBlack := false
	if len(s.Name) > 1 {
		isBlack = true
	}

	return isBlack
}

package gmusic

type (
	MusicLineList []*MusicLine
	MusicLine     struct {
		Index int           //线和间的索引
		Names MusicNameList //音名集合（本音和升降音）
	}
)

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 判断是否线间
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s MusicLine) IsSpace() bool {
	isSpace := false

	if s.Index%2 == 0 {
		isSpace = true
	}

	return isSpace
}

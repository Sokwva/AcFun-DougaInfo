package rpcproto

import (
	"sokwva/acfun/dougaInfo/structs"
)

// goverter:converter
// goverter:extend IntToInt32
// goverter:ignoreUnexported
type Converter interface {
	ConvertItems(source []structs.DougaInfo) []DougaInfo

	// goverter:ignoreUnexported
	// goverter:map DougaBaseInfo.Title Title
	// goverter:map DougaBaseInfo.BananaCountShow BananaCountShow
	// goverter:map DougaBaseInfo.IsLike IsLike
	// goverter:map DougaBaseInfo.CommentCount CommentCount
	// goverter:map DougaBaseInfo.BananaCount BananaCount
	// goverter:map DougaBaseInfo.Mkey Mkey
	// goverter:map DougaBaseInfo.IsFavorite IsFavorite
	// goverter:map DougaBaseInfo.IsThrowBanana IsThrowBanana
	// goverter:map DougaBaseInfo.StowCount StowCount
	Convert(source structs.DougaInfo) DougaInfo
}

func IntToInt32(i int) int32 {
	return int32(i)
}

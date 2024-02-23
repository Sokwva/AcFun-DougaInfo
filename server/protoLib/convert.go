package rpcproto

import (
	"sokwva/acfun/dougaInfo/structs"
)

// goverter:converter
// goverter:extend IntToInt32 Int32ToInt ConvertEmptyChannel
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

	// goverter:ignoreUnexported
	// goverter:ignore DougaBaseInfo
	// goverter:useZeroValueOnPointerInconsistency yes
	// goverter:ignore HostName
	Revert(source DougaInfo) structs.DougaInfo
}

func IntToInt32(i int) int32 {
	return int32(i)
}

func Int32ToInt(i int32) int {
	return int(i)
}

func ConvertEmptyChannel(i *DougaInfo_Channel) structs.Channel {
	if i == nil {
		return structs.Channel{}
	}
	if i.ID != 0 {
		return structs.Channel{
			ParentName: i.ParentName,
			ParentID:   int(i.ParentID),
			Name:       i.Name,
			ID:         int(i.ID),
		}
	}
	return structs.Channel{}
}

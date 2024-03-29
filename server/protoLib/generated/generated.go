// Code generated by github.com/jmattheis/goverter, DO NOT EDIT.
//go:build !goverter

package generated

import (
	protolib "sokwva/acfun/dougaInfo/server/protoLib"
	structs "sokwva/acfun/dougaInfo/structs"
)

type ConverterImpl struct{}

func (c *ConverterImpl) Convert(source structs.DougaInfo) protolib.DougaInfo {
	var rpcprotoDougaInfo protolib.DougaInfo
	rpcprotoDougaInfo.CurrentVideoID = protolib.IntToInt32(source.CurrentVideoID)
	rpcprotoDougaInfo.Pctr = source.Pctr
	rpcprotoDougaInfo.CommentCountRealValue = protolib.IntToInt32(source.CommentCountRealValue)
	rpcprotoDougaInfo.GroupID = source.GroupID
	rpcprotoDougaInfo.BananaCountShow = source.DougaBaseInfo.BananaCountShow
	rpcprotoDougaInfo.StowCountShow = source.StowCountShow
	rpcprotoDougaInfo.GiftPeachCountShow = source.GiftPeachCountShow
	rpcprotoDougaInfo.StowCount = protolib.IntToInt32(source.DougaBaseInfo.StowCount)
	rpcprotoDougaInfo.Channel = c.structsChannelToPRpcprotoDougaInfo_Channel(source.Channel)
	rpcprotoDougaInfo.LikeCount = protolib.IntToInt32(source.LikeCount)
	rpcprotoDougaInfo.HasHotComment = source.HasHotComment
	rpcprotoDougaInfo.ShareCountShow = source.ShareCountShow
	rpcprotoDougaInfo.IsDislike = source.IsDislike
	rpcprotoDougaInfo.ShareCount = protolib.IntToInt32(source.ShareCount)
	rpcprotoDougaInfo.PicShareURL = source.PicShareURL
	rpcprotoDougaInfo.DanmakuCount = protolib.IntToInt32(source.DanmakuCount)
	rpcprotoDougaInfo.CurrentVideoInfo = c.structsCurrentVideoInfoToPRpcprotoDougaInfo_Currentvideoinfo(source.CurrentVideoInfo)
	rpcprotoDougaInfo.ViewCountShow = source.ViewCountShow
	rpcprotoDougaInfo.DougaID = source.DougaID
	rpcprotoDougaInfo.CommentCountTenThousandShow = source.CommentCountTenThousandShow
	rpcprotoDougaInfo.CoverImgInfo = c.structsCoverImgInfoToPRpcprotoDougaInfo_Coverimginfo(source.CoverImgInfo)
	rpcprotoDougaInfo.CommentCount = protolib.IntToInt32(source.DougaBaseInfo.CommentCount)
	rpcprotoDougaInfo.CoverURL = source.CoverURL
	rpcprotoDougaInfo.DisableEdit = source.DisableEdit
	rpcprotoDougaInfo.CreateTimeMillis = source.CreateTimeMillis
	rpcprotoDougaInfo.SuperUbb = source.SuperUbb
	rpcprotoDougaInfo.ShareURL = source.ShareURL
	rpcprotoDougaInfo.Status = protolib.IntToInt32(source.Status)
	rpcprotoDougaInfo.IsFavorite = source.DougaBaseInfo.IsFavorite
	rpcprotoDougaInfo.IsLike = source.DougaBaseInfo.IsLike
	rpcprotoDougaInfo.RecoReason = c.structsRecoReasonToPRpcprotoDougaInfo_Recoreason(source.RecoReason)
	rpcprotoDougaInfo.Description = source.Description
	rpcprotoDougaInfo.Title = source.DougaBaseInfo.Title
	rpcprotoDougaInfo.BelongToSpecifyArubamu = source.BelongToSpecifyArubamu
	rpcprotoDougaInfo.Result = protolib.IntToInt32(source.Result)
	rpcprotoDougaInfo.OriginalDeclare = protolib.IntToInt32(source.OriginalDeclare)
	var pRpcprotoDougaInfo_VideolistList []*protolib.DougaInfo_Videolist
	if source.VideoList != nil {
		pRpcprotoDougaInfo_VideolistList = make([]*protolib.DougaInfo_Videolist, len(source.VideoList))
		for i := 0; i < len(source.VideoList); i++ {
			pRpcprotoDougaInfo_VideolistList[i] = c.structsVideoListToPRpcprotoDougaInfo_Videolist(source.VideoList[i])
		}
	}
	rpcprotoDougaInfo.VideoList = pRpcprotoDougaInfo_VideolistList
	rpcprotoDougaInfo.IsThrowBanana = source.DougaBaseInfo.IsThrowBanana
	rpcprotoDougaInfo.ViewCount = protolib.IntToInt32(source.ViewCount)
	rpcprotoDougaInfo.BananaCount = protolib.IntToInt32(source.DougaBaseInfo.BananaCount)
	rpcprotoDougaInfo.LikeCountShow = source.LikeCountShow
	var pRpcprotoDougaInfo_CovercdnurlsList []*protolib.DougaInfo_Covercdnurls
	if source.CoverCdnUrls != nil {
		pRpcprotoDougaInfo_CovercdnurlsList = make([]*protolib.DougaInfo_Covercdnurls, len(source.CoverCdnUrls))
		for j := 0; j < len(source.CoverCdnUrls); j++ {
			pRpcprotoDougaInfo_CovercdnurlsList[j] = c.structsCoverCdnUrlsToPRpcprotoDougaInfo_Covercdnurls(source.CoverCdnUrls[j])
		}
	}
	rpcprotoDougaInfo.CoverCdnUrls = pRpcprotoDougaInfo_CovercdnurlsList
	rpcprotoDougaInfo.GiftPeachCount = protolib.IntToInt32(source.GiftPeachCount)
	rpcprotoDougaInfo.IsRewardSupportted = source.IsRewardSupportted
	rpcprotoDougaInfo.CommentCountShow = source.CommentCountShow
	rpcprotoDougaInfo.DurationMillis = protolib.IntToInt32(source.DurationMillis)
	rpcprotoDougaInfo.Priority = protolib.IntToInt32(source.Priority)
	var pRpcprotoDougaInfo_TaglistList []*protolib.DougaInfo_Taglist
	if source.TagList != nil {
		pRpcprotoDougaInfo_TaglistList = make([]*protolib.DougaInfo_Taglist, len(source.TagList))
		for k := 0; k < len(source.TagList); k++ {
			pRpcprotoDougaInfo_TaglistList[k] = c.structsTagListToPRpcprotoDougaInfo_Taglist(source.TagList[k])
		}
	}
	rpcprotoDougaInfo.TagList = pRpcprotoDougaInfo_TaglistList
	rpcprotoDougaInfo.DanmakuCountShow = source.DanmakuCountShow
	rpcprotoDougaInfo.CreateTime = source.CreateTime
	rpcprotoDougaInfo.User = c.structsUserToPRpcprotoDougaInfo_User(source.User)
	rpcprotoDougaInfo.Mkey = source.DougaBaseInfo.Mkey
	return rpcprotoDougaInfo
}
func (c *ConverterImpl) ConvertItems(source []structs.DougaInfo) []protolib.DougaInfo {
	var rpcprotoDougaInfoList []protolib.DougaInfo
	if source != nil {
		rpcprotoDougaInfoList = make([]protolib.DougaInfo, len(source))
		for i := 0; i < len(source); i++ {
			rpcprotoDougaInfoList[i] = c.Convert(source[i])
		}
	}
	return rpcprotoDougaInfoList
}
func (c *ConverterImpl) Revert(source protolib.DougaInfo) structs.DougaInfo {
	var structsDougaInfo structs.DougaInfo
	structsDougaInfo.CurrentVideoID = protolib.Int32ToInt(source.CurrentVideoID)
	structsDougaInfo.Pctr = source.Pctr
	structsDougaInfo.CommentCountRealValue = protolib.Int32ToInt(source.CommentCountRealValue)
	structsDougaInfo.GroupID = source.GroupID
	structsDougaInfo.StowCountShow = source.StowCountShow
	structsDougaInfo.GiftPeachCountShow = source.GiftPeachCountShow
	structsDougaInfo.Channel = protolib.ConvertEmptyChannel(source.Channel)
	structsDougaInfo.LikeCount = protolib.Int32ToInt(source.LikeCount)
	structsDougaInfo.HasHotComment = source.HasHotComment
	structsDougaInfo.ShareCountShow = source.ShareCountShow
	structsDougaInfo.IsDislike = source.IsDislike
	structsDougaInfo.ShareCount = protolib.Int32ToInt(source.ShareCount)
	structsDougaInfo.PicShareURL = source.PicShareURL
	structsDougaInfo.DanmakuCount = protolib.Int32ToInt(source.DanmakuCount)
	structsDougaInfo.CurrentVideoInfo = c.pRpcprotoDougaInfo_CurrentvideoinfoToStructsCurrentVideoInfo(source.CurrentVideoInfo)
	structsDougaInfo.ViewCountShow = source.ViewCountShow
	structsDougaInfo.DougaID = source.DougaID
	structsDougaInfo.CommentCountTenThousandShow = source.CommentCountTenThousandShow
	structsDougaInfo.CoverImgInfo = c.pRpcprotoDougaInfo_CoverimginfoToStructsCoverImgInfo(source.CoverImgInfo)
	structsDougaInfo.CoverURL = source.CoverURL
	structsDougaInfo.DisableEdit = source.DisableEdit
	structsDougaInfo.CreateTimeMillis = source.CreateTimeMillis
	structsDougaInfo.SuperUbb = source.SuperUbb
	structsDougaInfo.ShareURL = source.ShareURL
	structsDougaInfo.Status = protolib.Int32ToInt(source.Status)
	structsDougaInfo.RecoReason = c.pRpcprotoDougaInfo_RecoreasonToStructsRecoReason(source.RecoReason)
	structsDougaInfo.Description = source.Description
	structsDougaInfo.BelongToSpecifyArubamu = source.BelongToSpecifyArubamu
	structsDougaInfo.Result = protolib.Int32ToInt(source.Result)
	structsDougaInfo.OriginalDeclare = protolib.Int32ToInt(source.OriginalDeclare)
	var structsVideoListList []structs.VideoList
	if source.VideoList != nil {
		structsVideoListList = make([]structs.VideoList, len(source.VideoList))
		for i := 0; i < len(source.VideoList); i++ {
			structsVideoListList[i] = c.pRpcprotoDougaInfo_VideolistToStructsVideoList(source.VideoList[i])
		}
	}
	structsDougaInfo.VideoList = structsVideoListList
	structsDougaInfo.ViewCount = protolib.Int32ToInt(source.ViewCount)
	structsDougaInfo.LikeCountShow = source.LikeCountShow
	var structsCoverCdnUrlsList []structs.CoverCdnUrls
	if source.CoverCdnUrls != nil {
		structsCoverCdnUrlsList = make([]structs.CoverCdnUrls, len(source.CoverCdnUrls))
		for j := 0; j < len(source.CoverCdnUrls); j++ {
			structsCoverCdnUrlsList[j] = c.pRpcprotoDougaInfo_CovercdnurlsToStructsCoverCdnUrls(source.CoverCdnUrls[j])
		}
	}
	structsDougaInfo.CoverCdnUrls = structsCoverCdnUrlsList
	structsDougaInfo.GiftPeachCount = protolib.Int32ToInt(source.GiftPeachCount)
	structsDougaInfo.IsRewardSupportted = source.IsRewardSupportted
	structsDougaInfo.CommentCountShow = source.CommentCountShow
	structsDougaInfo.DurationMillis = protolib.Int32ToInt(source.DurationMillis)
	var structsTagListList []structs.TagList
	if source.TagList != nil {
		structsTagListList = make([]structs.TagList, len(source.TagList))
		for k := 0; k < len(source.TagList); k++ {
			structsTagListList[k] = c.pRpcprotoDougaInfo_TaglistToStructsTagList(source.TagList[k])
		}
	}
	structsDougaInfo.TagList = structsTagListList
	structsDougaInfo.DanmakuCountShow = source.DanmakuCountShow
	structsDougaInfo.CreateTime = source.CreateTime
	structsDougaInfo.User = c.pRpcprotoDougaInfo_UserToStructsUser(source.User)
	structsDougaInfo.Priority = protolib.Int32ToInt(source.Priority)
	return structsDougaInfo
}
func (c *ConverterImpl) pRpcprotoDougaInfo_Cdnurls1ToStructsCdnUrls(source *protolib.DougaInfo_Cdnurls1) structs.CdnUrls {
	var structsCdnUrls structs.CdnUrls
	if source != nil {
		var structsCdnUrls2 structs.CdnUrls
		structsCdnUrls2.URL = (*source).URL
		structsCdnUrls2.FreeTrafficProductAbbreviation = (*source).FreeTrafficProductAbbreviation
		structsCdnUrls2.FreeTrafficCdn = (*source).FreeTrafficCdn
		structsCdnUrls = structsCdnUrls2
	}
	return structsCdnUrls
}
func (c *ConverterImpl) pRpcprotoDougaInfo_Cdnurls2ToStructsCdnUrls(source *protolib.DougaInfo_Cdnurls2) structs.CdnUrls {
	var structsCdnUrls structs.CdnUrls
	if source != nil {
		var structsCdnUrls2 structs.CdnUrls
		structsCdnUrls2.URL = (*source).URL
		structsCdnUrls2.FreeTrafficProductAbbreviation = (*source).FreeTrafficProductAbbreviation
		structsCdnUrls2.FreeTrafficCdn = (*source).FreeTrafficCdn
		structsCdnUrls = structsCdnUrls2
	}
	return structsCdnUrls
}
func (c *ConverterImpl) pRpcprotoDougaInfo_Cdnurls3ToStructsCdnUrls(source *protolib.DougaInfo_Cdnurls3) structs.CdnUrls {
	var structsCdnUrls structs.CdnUrls
	if source != nil {
		var structsCdnUrls2 structs.CdnUrls
		structsCdnUrls2.URL = (*source).URL
		structsCdnUrls2.FreeTrafficProductAbbreviation = (*source).FreeTrafficProductAbbreviation
		structsCdnUrls2.FreeTrafficCdn = (*source).FreeTrafficCdn
		structsCdnUrls = structsCdnUrls2
	}
	return structsCdnUrls
}
func (c *ConverterImpl) pRpcprotoDougaInfo_CdnurlsToStructsCdnUrls(source *protolib.DougaInfo_Cdnurls) structs.CdnUrls {
	var structsCdnUrls structs.CdnUrls
	if source != nil {
		var structsCdnUrls2 structs.CdnUrls
		structsCdnUrls2.URL = (*source).URL
		structsCdnUrls2.FreeTrafficProductAbbreviation = (*source).FreeTrafficProductAbbreviation
		structsCdnUrls2.FreeTrafficCdn = (*source).FreeTrafficCdn
		structsCdnUrls = structsCdnUrls2
	}
	return structsCdnUrls
}
func (c *ConverterImpl) pRpcprotoDougaInfo_CovercdnurlsToStructsCoverCdnUrls(source *protolib.DougaInfo_Covercdnurls) structs.CoverCdnUrls {
	var structsCoverCdnUrls structs.CoverCdnUrls
	if source != nil {
		var structsCoverCdnUrls2 structs.CoverCdnUrls
		structsCoverCdnUrls2.URL = (*source).URL
		structsCoverCdnUrls2.FreeTrafficCdn = (*source).FreeTrafficCdn
		structsCoverCdnUrls = structsCoverCdnUrls2
	}
	return structsCoverCdnUrls
}
func (c *ConverterImpl) pRpcprotoDougaInfo_CoverimginfoToStructsCoverImgInfo(source *protolib.DougaInfo_Coverimginfo) structs.CoverImgInfo {
	var structsCoverImgInfo structs.CoverImgInfo
	if source != nil {
		var structsCoverImgInfo2 structs.CoverImgInfo
		structsCoverImgInfo2.Width = protolib.Int32ToInt((*source).Width)
		structsCoverImgInfo2.Height = protolib.Int32ToInt((*source).Height)
		structsCoverImgInfo2.Size = protolib.Int32ToInt((*source).Size)
		structsCoverImgInfo2.Type = protolib.Int32ToInt((*source).Type)
		structsCoverImgInfo2.Animated = (*source).Animated
		structsCoverImgInfo2.ThumbnailImage = c.pRpcprotoDougaInfo_ThumbnailimageToStructsThumbnailImage((*source).ThumbnailImage)
		structsCoverImgInfo2.ExpandedImage = c.pRpcprotoDougaInfo_ExpandedimageToStructsExpandedImage((*source).ExpandedImage)
		structsCoverImgInfo2.ThumbnailImageCdnURL = (*source).ThumbnailImageCdnURL
		structsCoverImgInfo = structsCoverImgInfo2
	}
	return structsCoverImgInfo
}
func (c *ConverterImpl) pRpcprotoDougaInfo_CurrentvideoinfoToStructsCurrentVideoInfo(source *protolib.DougaInfo_Currentvideoinfo) structs.CurrentVideoInfo {
	var structsCurrentVideoInfo structs.CurrentVideoInfo
	if source != nil {
		var structsCurrentVideoInfo2 structs.CurrentVideoInfo
		structsCurrentVideoInfo2.Priority = protolib.Int32ToInt((*source).Priority)
		structsCurrentVideoInfo2.Title = (*source).Title
		structsCurrentVideoInfo2.DanmakuCountShow = (*source).DanmakuCountShow
		structsCurrentVideoInfo2.UserPlayedSeconds = protolib.Int32ToInt((*source).UserPlayedSeconds)
		structsCurrentVideoInfo2.SizeType = protolib.Int32ToInt((*source).SizeType)
		structsCurrentVideoInfo2.DanmakuGuidePosition = protolib.Int32ToInt((*source).DanmakuGuidePosition)
		structsCurrentVideoInfo2.SourceStatus = protolib.Int32ToInt((*source).SourceStatus)
		structsCurrentVideoInfo2.DurationMillis = protolib.Int32ToInt((*source).DurationMillis)
		structsCurrentVideoInfo2.UploadTime = (*source).UploadTime
		structsCurrentVideoInfo2.VisibleType = protolib.Int32ToInt((*source).VisibleType)
		structsCurrentVideoInfo2.DanmakuCount = protolib.Int32ToInt((*source).DanmakuCount)
		structsCurrentVideoInfo2.FileName = (*source).FileName
		structsCurrentVideoInfo2.IsKsManifest = (*source).IsKsManifest
		var structsTranscodeInfosList []structs.TranscodeInfos
		if (*source).TranscodeInfos != nil {
			structsTranscodeInfosList = make([]structs.TranscodeInfos, len((*source).TranscodeInfos))
			for i := 0; i < len((*source).TranscodeInfos); i++ {
				structsTranscodeInfosList[i] = c.pRpcprotoDougaInfo_TranscodeinfosToStructsTranscodeInfos((*source).TranscodeInfos[i])
			}
		}
		structsCurrentVideoInfo2.TranscodeInfos = structsTranscodeInfosList
		structsCurrentVideoInfo2.ID = (*source).ID
		structsCurrentVideoInfo = structsCurrentVideoInfo2
	}
	return structsCurrentVideoInfo
}
func (c *ConverterImpl) pRpcprotoDougaInfo_Expandedimage1ToStructsExpandedImage(source *protolib.DougaInfo_Expandedimage1) structs.ExpandedImage {
	var structsExpandedImage structs.ExpandedImage
	if source != nil {
		var structsExpandedImage2 structs.ExpandedImage
		var structsCdnUrlsList []structs.CdnUrls
		if (*source).CdnUrls != nil {
			structsCdnUrlsList = make([]structs.CdnUrls, len((*source).CdnUrls))
			for i := 0; i < len((*source).CdnUrls); i++ {
				structsCdnUrlsList[i] = c.pRpcprotoDougaInfo_Cdnurls3ToStructsCdnUrls((*source).CdnUrls[i])
			}
		}
		structsExpandedImage2.CdnUrls = structsCdnUrlsList
		structsExpandedImage = structsExpandedImage2
	}
	return structsExpandedImage
}
func (c *ConverterImpl) pRpcprotoDougaInfo_ExpandedimageToStructsExpandedImage(source *protolib.DougaInfo_Expandedimage) structs.ExpandedImage {
	var structsExpandedImage structs.ExpandedImage
	if source != nil {
		var structsExpandedImage2 structs.ExpandedImage
		var structsCdnUrlsList []structs.CdnUrls
		if (*source).CdnUrls != nil {
			structsCdnUrlsList = make([]structs.CdnUrls, len((*source).CdnUrls))
			for i := 0; i < len((*source).CdnUrls); i++ {
				structsCdnUrlsList[i] = c.pRpcprotoDougaInfo_Cdnurls1ToStructsCdnUrls((*source).CdnUrls[i])
			}
		}
		structsExpandedImage2.CdnUrls = structsCdnUrlsList
		structsExpandedImage = structsExpandedImage2
	}
	return structsExpandedImage
}
func (c *ConverterImpl) pRpcprotoDougaInfo_HeadcdnurlsToStructsHeadCdnUrls(source *protolib.DougaInfo_Headcdnurls) structs.HeadCdnUrls {
	var structsHeadCdnUrls structs.HeadCdnUrls
	if source != nil {
		var structsHeadCdnUrls2 structs.HeadCdnUrls
		structsHeadCdnUrls2.URL = (*source).URL
		structsHeadCdnUrls2.FreeTrafficCdn = (*source).FreeTrafficCdn
		structsHeadCdnUrls = structsHeadCdnUrls2
	}
	return structsHeadCdnUrls
}
func (c *ConverterImpl) pRpcprotoDougaInfo_RecoreasonToStructsRecoReason(source *protolib.DougaInfo_Recoreason) structs.RecoReason {
	var structsRecoReason structs.RecoReason
	if source != nil {
		var structsRecoReason2 structs.RecoReason
		structsRecoReason2.Desc = (*source).Desc
		structsRecoReason2.Href = (*source).Href
		structsRecoReason2.Tag = (*source).Tag
		structsRecoReason2.LayoutType = protolib.Int32ToInt((*source).LayoutType)
		structsRecoReason2.DescType = protolib.Int32ToInt((*source).DescType)
		structsRecoReason2.Type = protolib.Int32ToInt((*source).Type)
		structsRecoReason = structsRecoReason2
	}
	return structsRecoReason
}
func (c *ConverterImpl) pRpcprotoDougaInfo_SocialmedalToStructsSocialMedal(source *protolib.DougaInfo_Socialmedal) structs.SocialMedal {
	var structsSocialMedal structs.SocialMedal
	if source != nil {
		var structsSocialMedal2 structs.SocialMedal
		_ = (*source)
		structsSocialMedal = structsSocialMedal2
	}
	return structsSocialMedal
}
func (c *ConverterImpl) pRpcprotoDougaInfo_TaglistToStructsTagList(source *protolib.DougaInfo_Taglist) structs.TagList {
	var structsTagList structs.TagList
	if source != nil {
		var structsTagList2 structs.TagList
		structsTagList2.Name = (*source).Name
		structsTagList2.ID = (*source).ID
		structsTagList = structsTagList2
	}
	return structsTagList
}
func (c *ConverterImpl) pRpcprotoDougaInfo_Thumbnailimage1ToStructsThumbnailImage(source *protolib.DougaInfo_Thumbnailimage1) structs.ThumbnailImage {
	var structsThumbnailImage structs.ThumbnailImage
	if source != nil {
		var structsThumbnailImage2 structs.ThumbnailImage
		var structsCdnUrlsList []structs.CdnUrls
		if (*source).CdnUrls != nil {
			structsCdnUrlsList = make([]structs.CdnUrls, len((*source).CdnUrls))
			for i := 0; i < len((*source).CdnUrls); i++ {
				structsCdnUrlsList[i] = c.pRpcprotoDougaInfo_Cdnurls2ToStructsCdnUrls((*source).CdnUrls[i])
			}
		}
		structsThumbnailImage2.CdnUrls = structsCdnUrlsList
		structsThumbnailImage2.Animated = protolib.Int32ToInt((*source).Animated)
		structsThumbnailImage = structsThumbnailImage2
	}
	return structsThumbnailImage
}
func (c *ConverterImpl) pRpcprotoDougaInfo_ThumbnailimageToStructsThumbnailImage(source *protolib.DougaInfo_Thumbnailimage) structs.ThumbnailImage {
	var structsThumbnailImage structs.ThumbnailImage
	if source != nil {
		var structsThumbnailImage2 structs.ThumbnailImage
		var structsCdnUrlsList []structs.CdnUrls
		if (*source).CdnUrls != nil {
			structsCdnUrlsList = make([]structs.CdnUrls, len((*source).CdnUrls))
			for i := 0; i < len((*source).CdnUrls); i++ {
				structsCdnUrlsList[i] = c.pRpcprotoDougaInfo_CdnurlsToStructsCdnUrls((*source).CdnUrls[i])
			}
		}
		structsThumbnailImage2.CdnUrls = structsCdnUrlsList
		structsThumbnailImage2.Animated = protolib.Int32ToInt((*source).Animated)
		structsThumbnailImage = structsThumbnailImage2
	}
	return structsThumbnailImage
}
func (c *ConverterImpl) pRpcprotoDougaInfo_TranscodeinfosToStructsTranscodeInfos(source *protolib.DougaInfo_Transcodeinfos) structs.TranscodeInfos {
	var structsTranscodeInfos structs.TranscodeInfos
	if source != nil {
		var structsTranscodeInfos2 structs.TranscodeInfos
		structsTranscodeInfos2.QualityType = (*source).QualityType
		structsTranscodeInfos2.SizeInBytes = protolib.Int32ToInt((*source).SizeInBytes)
		structsTranscodeInfos2.Hdr = (*source).Hdr
		structsTranscodeInfos = structsTranscodeInfos2
	}
	return structsTranscodeInfos
}
func (c *ConverterImpl) pRpcprotoDougaInfo_UserToStructsUser(source *protolib.DougaInfo_User) structs.User {
	var structsUser structs.User
	if source != nil {
		var structsUser2 structs.User
		structsUser2.Action = protolib.Int32ToInt((*source).Action)
		structsUser2.Href = (*source).Href
		var structsHeadCdnUrlsList []structs.HeadCdnUrls
		if (*source).HeadCdnUrls != nil {
			structsHeadCdnUrlsList = make([]structs.HeadCdnUrls, len((*source).HeadCdnUrls))
			for i := 0; i < len((*source).HeadCdnUrls); i++ {
				structsHeadCdnUrlsList[i] = c.pRpcprotoDougaInfo_HeadcdnurlsToStructsHeadCdnUrls((*source).HeadCdnUrls[i])
			}
		}
		structsUser2.HeadCdnUrls = structsHeadCdnUrlsList
		structsUser2.IsJoinUpCollege = (*source).IsJoinUpCollege
		structsUser2.FollowingCountValue = protolib.Int32ToInt((*source).FollowingCountValue)
		structsUser2.ContributeCountValue = protolib.Int32ToInt((*source).ContributeCountValue)
		structsUser2.FanCount = (*source).FanCount
		structsUser2.SocialMedal = c.pRpcprotoDougaInfo_SocialmedalToStructsSocialMedal((*source).SocialMedal)
		structsUser2.AvatarImage = (*source).AvatarImage
		structsUser2.UserHeadImgInfo = c.pRpcprotoDougaInfo_UserheadimginfoToStructsUserHeadImgInfo((*source).UserHeadImgInfo)
		structsUser2.IsFollowed = (*source).IsFollowed
		structsUser2.IsFollowing = (*source).IsFollowing
		structsUser2.AvatarFrame = protolib.Int32ToInt((*source).AvatarFrame)
		structsUser2.ContributeCount = (*source).ContributeCount
		structsUser2.AvatarFramePcImg = (*source).AvatarFramePcImg
		structsUser2.AvatarFrameMobileImg = (*source).AvatarFrameMobileImg
		structsUser2.SexTrend = protolib.Int32ToInt((*source).SexTrend)
		var intList []int
		if (*source).VerifiedTypes != nil {
			intList = make([]int, len((*source).VerifiedTypes))
			for j := 0; j < len((*source).VerifiedTypes); j++ {
				intList[j] = protolib.Int32ToInt((*source).VerifiedTypes[j])
			}
		}
		structsUser2.VerifiedTypes = intList
		structsUser2.HeadURL = (*source).HeadURL
		structsUser2.FanCountValue = protolib.Int32ToInt((*source).FanCountValue)
		structsUser2.FollowingStatus = protolib.Int32ToInt((*source).FollowingStatus)
		structsUser2.NameColor = protolib.Int32ToInt((*source).NameColor)
		structsUser2.FollowingCount = (*source).FollowingCount
		structsUser2.Gender = protolib.Int32ToInt((*source).Gender)
		structsUser2.Name = (*source).Name
		structsUser2.Signature = (*source).Signature
		structsUser2.ID = (*source).ID
		structsUser2.NameStyle = (*source).NameStyle
		structsUser = structsUser2
	}
	return structsUser
}
func (c *ConverterImpl) pRpcprotoDougaInfo_UserheadimginfoToStructsUserHeadImgInfo(source *protolib.DougaInfo_Userheadimginfo) structs.UserHeadImgInfo {
	var structsUserHeadImgInfo structs.UserHeadImgInfo
	if source != nil {
		var structsUserHeadImgInfo2 structs.UserHeadImgInfo
		structsUserHeadImgInfo2.Width = protolib.Int32ToInt((*source).Width)
		structsUserHeadImgInfo2.Height = protolib.Int32ToInt((*source).Height)
		structsUserHeadImgInfo2.Size = protolib.Int32ToInt((*source).Size)
		structsUserHeadImgInfo2.Type = protolib.Int32ToInt((*source).Type)
		structsUserHeadImgInfo2.Animated = (*source).Animated
		structsUserHeadImgInfo2.ThumbnailImage = c.pRpcprotoDougaInfo_Thumbnailimage1ToStructsThumbnailImage((*source).ThumbnailImage)
		structsUserHeadImgInfo2.ExpandedImage = c.pRpcprotoDougaInfo_Expandedimage1ToStructsExpandedImage((*source).ExpandedImage)
		structsUserHeadImgInfo2.ThumbnailImageCdnURL = (*source).ThumbnailImageCdnURL
		structsUserHeadImgInfo = structsUserHeadImgInfo2
	}
	return structsUserHeadImgInfo
}
func (c *ConverterImpl) pRpcprotoDougaInfo_VideolistToStructsVideoList(source *protolib.DougaInfo_Videolist) structs.VideoList {
	var structsVideoList structs.VideoList
	if source != nil {
		var structsVideoList2 structs.VideoList
		structsVideoList2.Priority = protolib.Int32ToInt((*source).Priority)
		structsVideoList2.Title = (*source).Title
		structsVideoList2.DanmakuCountShow = (*source).DanmakuCountShow
		structsVideoList2.UserPlayedSeconds = protolib.Int32ToInt((*source).UserPlayedSeconds)
		structsVideoList2.SizeType = protolib.Int32ToInt((*source).SizeType)
		structsVideoList2.DanmakuGuidePosition = protolib.Int32ToInt((*source).DanmakuGuidePosition)
		structsVideoList2.SourceStatus = protolib.Int32ToInt((*source).SourceStatus)
		structsVideoList2.DurationMillis = protolib.Int32ToInt((*source).DurationMillis)
		structsVideoList2.UploadTime = (*source).UploadTime
		structsVideoList2.VisibleType = protolib.Int32ToInt((*source).VisibleType)
		structsVideoList2.DanmakuCount = protolib.Int32ToInt((*source).DanmakuCount)
		structsVideoList2.FileName = (*source).FileName
		structsVideoList2.ID = (*source).ID
		structsVideoList = structsVideoList2
	}
	return structsVideoList
}
func (c *ConverterImpl) structsCdnUrlsToPRpcprotoDougaInfo_Cdnurls(source structs.CdnUrls) *protolib.DougaInfo_Cdnurls {
	var rpcprotoDougaInfo_Cdnurls protolib.DougaInfo_Cdnurls
	rpcprotoDougaInfo_Cdnurls.URL = source.URL
	rpcprotoDougaInfo_Cdnurls.FreeTrafficProductAbbreviation = source.FreeTrafficProductAbbreviation
	rpcprotoDougaInfo_Cdnurls.FreeTrafficCdn = source.FreeTrafficCdn
	return &rpcprotoDougaInfo_Cdnurls
}
func (c *ConverterImpl) structsCdnUrlsToPRpcprotoDougaInfo_Cdnurls1(source structs.CdnUrls) *protolib.DougaInfo_Cdnurls1 {
	var rpcprotoDougaInfo_Cdnurls1 protolib.DougaInfo_Cdnurls1
	rpcprotoDougaInfo_Cdnurls1.URL = source.URL
	rpcprotoDougaInfo_Cdnurls1.FreeTrafficProductAbbreviation = source.FreeTrafficProductAbbreviation
	rpcprotoDougaInfo_Cdnurls1.FreeTrafficCdn = source.FreeTrafficCdn
	return &rpcprotoDougaInfo_Cdnurls1
}
func (c *ConverterImpl) structsCdnUrlsToPRpcprotoDougaInfo_Cdnurls2(source structs.CdnUrls) *protolib.DougaInfo_Cdnurls2 {
	var rpcprotoDougaInfo_Cdnurls2 protolib.DougaInfo_Cdnurls2
	rpcprotoDougaInfo_Cdnurls2.URL = source.URL
	rpcprotoDougaInfo_Cdnurls2.FreeTrafficProductAbbreviation = source.FreeTrafficProductAbbreviation
	rpcprotoDougaInfo_Cdnurls2.FreeTrafficCdn = source.FreeTrafficCdn
	return &rpcprotoDougaInfo_Cdnurls2
}
func (c *ConverterImpl) structsCdnUrlsToPRpcprotoDougaInfo_Cdnurls3(source structs.CdnUrls) *protolib.DougaInfo_Cdnurls3 {
	var rpcprotoDougaInfo_Cdnurls3 protolib.DougaInfo_Cdnurls3
	rpcprotoDougaInfo_Cdnurls3.URL = source.URL
	rpcprotoDougaInfo_Cdnurls3.FreeTrafficProductAbbreviation = source.FreeTrafficProductAbbreviation
	rpcprotoDougaInfo_Cdnurls3.FreeTrafficCdn = source.FreeTrafficCdn
	return &rpcprotoDougaInfo_Cdnurls3
}
func (c *ConverterImpl) structsChannelToPRpcprotoDougaInfo_Channel(source structs.Channel) *protolib.DougaInfo_Channel {
	var rpcprotoDougaInfo_Channel protolib.DougaInfo_Channel
	rpcprotoDougaInfo_Channel.ParentName = source.ParentName
	rpcprotoDougaInfo_Channel.ParentID = protolib.IntToInt32(source.ParentID)
	rpcprotoDougaInfo_Channel.Name = source.Name
	rpcprotoDougaInfo_Channel.ID = protolib.IntToInt32(source.ID)
	return &rpcprotoDougaInfo_Channel
}
func (c *ConverterImpl) structsCoverCdnUrlsToPRpcprotoDougaInfo_Covercdnurls(source structs.CoverCdnUrls) *protolib.DougaInfo_Covercdnurls {
	var rpcprotoDougaInfo_Covercdnurls protolib.DougaInfo_Covercdnurls
	rpcprotoDougaInfo_Covercdnurls.URL = source.URL
	rpcprotoDougaInfo_Covercdnurls.FreeTrafficCdn = source.FreeTrafficCdn
	return &rpcprotoDougaInfo_Covercdnurls
}
func (c *ConverterImpl) structsCoverImgInfoToPRpcprotoDougaInfo_Coverimginfo(source structs.CoverImgInfo) *protolib.DougaInfo_Coverimginfo {
	var rpcprotoDougaInfo_Coverimginfo protolib.DougaInfo_Coverimginfo
	rpcprotoDougaInfo_Coverimginfo.Width = protolib.IntToInt32(source.Width)
	rpcprotoDougaInfo_Coverimginfo.Height = protolib.IntToInt32(source.Height)
	rpcprotoDougaInfo_Coverimginfo.Size = protolib.IntToInt32(source.Size)
	rpcprotoDougaInfo_Coverimginfo.Type = protolib.IntToInt32(source.Type)
	rpcprotoDougaInfo_Coverimginfo.Animated = source.Animated
	rpcprotoDougaInfo_Coverimginfo.ThumbnailImage = c.structsThumbnailImageToPRpcprotoDougaInfo_Thumbnailimage(source.ThumbnailImage)
	rpcprotoDougaInfo_Coverimginfo.ExpandedImage = c.structsExpandedImageToPRpcprotoDougaInfo_Expandedimage(source.ExpandedImage)
	rpcprotoDougaInfo_Coverimginfo.ThumbnailImageCdnURL = source.ThumbnailImageCdnURL
	return &rpcprotoDougaInfo_Coverimginfo
}
func (c *ConverterImpl) structsCurrentVideoInfoToPRpcprotoDougaInfo_Currentvideoinfo(source structs.CurrentVideoInfo) *protolib.DougaInfo_Currentvideoinfo {
	var rpcprotoDougaInfo_Currentvideoinfo protolib.DougaInfo_Currentvideoinfo
	rpcprotoDougaInfo_Currentvideoinfo.Priority = protolib.IntToInt32(source.Priority)
	rpcprotoDougaInfo_Currentvideoinfo.ID = source.ID
	var pRpcprotoDougaInfo_TranscodeinfosList []*protolib.DougaInfo_Transcodeinfos
	if source.TranscodeInfos != nil {
		pRpcprotoDougaInfo_TranscodeinfosList = make([]*protolib.DougaInfo_Transcodeinfos, len(source.TranscodeInfos))
		for i := 0; i < len(source.TranscodeInfos); i++ {
			pRpcprotoDougaInfo_TranscodeinfosList[i] = c.structsTranscodeInfosToPRpcprotoDougaInfo_Transcodeinfos(source.TranscodeInfos[i])
		}
	}
	rpcprotoDougaInfo_Currentvideoinfo.TranscodeInfos = pRpcprotoDougaInfo_TranscodeinfosList
	rpcprotoDougaInfo_Currentvideoinfo.IsKsManifest = source.IsKsManifest
	rpcprotoDougaInfo_Currentvideoinfo.Title = source.Title
	rpcprotoDougaInfo_Currentvideoinfo.DanmakuCountShow = source.DanmakuCountShow
	rpcprotoDougaInfo_Currentvideoinfo.UserPlayedSeconds = protolib.IntToInt32(source.UserPlayedSeconds)
	rpcprotoDougaInfo_Currentvideoinfo.SizeType = protolib.IntToInt32(source.SizeType)
	rpcprotoDougaInfo_Currentvideoinfo.DanmakuGuidePosition = protolib.IntToInt32(source.DanmakuGuidePosition)
	rpcprotoDougaInfo_Currentvideoinfo.SourceStatus = protolib.IntToInt32(source.SourceStatus)
	rpcprotoDougaInfo_Currentvideoinfo.DurationMillis = protolib.IntToInt32(source.DurationMillis)
	rpcprotoDougaInfo_Currentvideoinfo.UploadTime = source.UploadTime
	rpcprotoDougaInfo_Currentvideoinfo.VisibleType = protolib.IntToInt32(source.VisibleType)
	rpcprotoDougaInfo_Currentvideoinfo.DanmakuCount = protolib.IntToInt32(source.DanmakuCount)
	rpcprotoDougaInfo_Currentvideoinfo.FileName = source.FileName
	return &rpcprotoDougaInfo_Currentvideoinfo
}
func (c *ConverterImpl) structsExpandedImageToPRpcprotoDougaInfo_Expandedimage(source structs.ExpandedImage) *protolib.DougaInfo_Expandedimage {
	var rpcprotoDougaInfo_Expandedimage protolib.DougaInfo_Expandedimage
	var pRpcprotoDougaInfo_Cdnurls1List []*protolib.DougaInfo_Cdnurls1
	if source.CdnUrls != nil {
		pRpcprotoDougaInfo_Cdnurls1List = make([]*protolib.DougaInfo_Cdnurls1, len(source.CdnUrls))
		for i := 0; i < len(source.CdnUrls); i++ {
			pRpcprotoDougaInfo_Cdnurls1List[i] = c.structsCdnUrlsToPRpcprotoDougaInfo_Cdnurls1(source.CdnUrls[i])
		}
	}
	rpcprotoDougaInfo_Expandedimage.CdnUrls = pRpcprotoDougaInfo_Cdnurls1List
	return &rpcprotoDougaInfo_Expandedimage
}
func (c *ConverterImpl) structsExpandedImageToPRpcprotoDougaInfo_Expandedimage1(source structs.ExpandedImage) *protolib.DougaInfo_Expandedimage1 {
	var rpcprotoDougaInfo_Expandedimage1 protolib.DougaInfo_Expandedimage1
	var pRpcprotoDougaInfo_Cdnurls3List []*protolib.DougaInfo_Cdnurls3
	if source.CdnUrls != nil {
		pRpcprotoDougaInfo_Cdnurls3List = make([]*protolib.DougaInfo_Cdnurls3, len(source.CdnUrls))
		for i := 0; i < len(source.CdnUrls); i++ {
			pRpcprotoDougaInfo_Cdnurls3List[i] = c.structsCdnUrlsToPRpcprotoDougaInfo_Cdnurls3(source.CdnUrls[i])
		}
	}
	rpcprotoDougaInfo_Expandedimage1.CdnUrls = pRpcprotoDougaInfo_Cdnurls3List
	return &rpcprotoDougaInfo_Expandedimage1
}
func (c *ConverterImpl) structsHeadCdnUrlsToPRpcprotoDougaInfo_Headcdnurls(source structs.HeadCdnUrls) *protolib.DougaInfo_Headcdnurls {
	var rpcprotoDougaInfo_Headcdnurls protolib.DougaInfo_Headcdnurls
	rpcprotoDougaInfo_Headcdnurls.URL = source.URL
	rpcprotoDougaInfo_Headcdnurls.FreeTrafficCdn = source.FreeTrafficCdn
	return &rpcprotoDougaInfo_Headcdnurls
}
func (c *ConverterImpl) structsRecoReasonToPRpcprotoDougaInfo_Recoreason(source structs.RecoReason) *protolib.DougaInfo_Recoreason {
	var rpcprotoDougaInfo_Recoreason protolib.DougaInfo_Recoreason
	rpcprotoDougaInfo_Recoreason.Desc = source.Desc
	rpcprotoDougaInfo_Recoreason.Href = source.Href
	rpcprotoDougaInfo_Recoreason.Tag = source.Tag
	rpcprotoDougaInfo_Recoreason.LayoutType = protolib.IntToInt32(source.LayoutType)
	rpcprotoDougaInfo_Recoreason.DescType = protolib.IntToInt32(source.DescType)
	rpcprotoDougaInfo_Recoreason.Type = protolib.IntToInt32(source.Type)
	return &rpcprotoDougaInfo_Recoreason
}
func (c *ConverterImpl) structsSocialMedalToPRpcprotoDougaInfo_Socialmedal(source structs.SocialMedal) *protolib.DougaInfo_Socialmedal {
	var rpcprotoDougaInfo_Socialmedal protolib.DougaInfo_Socialmedal
	_ = source
	return &rpcprotoDougaInfo_Socialmedal
}
func (c *ConverterImpl) structsTagListToPRpcprotoDougaInfo_Taglist(source structs.TagList) *protolib.DougaInfo_Taglist {
	var rpcprotoDougaInfo_Taglist protolib.DougaInfo_Taglist
	rpcprotoDougaInfo_Taglist.Name = source.Name
	rpcprotoDougaInfo_Taglist.ID = source.ID
	return &rpcprotoDougaInfo_Taglist
}
func (c *ConverterImpl) structsThumbnailImageToPRpcprotoDougaInfo_Thumbnailimage(source structs.ThumbnailImage) *protolib.DougaInfo_Thumbnailimage {
	var rpcprotoDougaInfo_Thumbnailimage protolib.DougaInfo_Thumbnailimage
	var pRpcprotoDougaInfo_CdnurlsList []*protolib.DougaInfo_Cdnurls
	if source.CdnUrls != nil {
		pRpcprotoDougaInfo_CdnurlsList = make([]*protolib.DougaInfo_Cdnurls, len(source.CdnUrls))
		for i := 0; i < len(source.CdnUrls); i++ {
			pRpcprotoDougaInfo_CdnurlsList[i] = c.structsCdnUrlsToPRpcprotoDougaInfo_Cdnurls(source.CdnUrls[i])
		}
	}
	rpcprotoDougaInfo_Thumbnailimage.CdnUrls = pRpcprotoDougaInfo_CdnurlsList
	rpcprotoDougaInfo_Thumbnailimage.Animated = protolib.IntToInt32(source.Animated)
	return &rpcprotoDougaInfo_Thumbnailimage
}
func (c *ConverterImpl) structsThumbnailImageToPRpcprotoDougaInfo_Thumbnailimage1(source structs.ThumbnailImage) *protolib.DougaInfo_Thumbnailimage1 {
	var rpcprotoDougaInfo_Thumbnailimage1 protolib.DougaInfo_Thumbnailimage1
	var pRpcprotoDougaInfo_Cdnurls2List []*protolib.DougaInfo_Cdnurls2
	if source.CdnUrls != nil {
		pRpcprotoDougaInfo_Cdnurls2List = make([]*protolib.DougaInfo_Cdnurls2, len(source.CdnUrls))
		for i := 0; i < len(source.CdnUrls); i++ {
			pRpcprotoDougaInfo_Cdnurls2List[i] = c.structsCdnUrlsToPRpcprotoDougaInfo_Cdnurls2(source.CdnUrls[i])
		}
	}
	rpcprotoDougaInfo_Thumbnailimage1.CdnUrls = pRpcprotoDougaInfo_Cdnurls2List
	rpcprotoDougaInfo_Thumbnailimage1.Animated = protolib.IntToInt32(source.Animated)
	return &rpcprotoDougaInfo_Thumbnailimage1
}
func (c *ConverterImpl) structsTranscodeInfosToPRpcprotoDougaInfo_Transcodeinfos(source structs.TranscodeInfos) *protolib.DougaInfo_Transcodeinfos {
	var rpcprotoDougaInfo_Transcodeinfos protolib.DougaInfo_Transcodeinfos
	rpcprotoDougaInfo_Transcodeinfos.QualityType = source.QualityType
	rpcprotoDougaInfo_Transcodeinfos.SizeInBytes = protolib.IntToInt32(source.SizeInBytes)
	rpcprotoDougaInfo_Transcodeinfos.Hdr = source.Hdr
	return &rpcprotoDougaInfo_Transcodeinfos
}
func (c *ConverterImpl) structsUserHeadImgInfoToPRpcprotoDougaInfo_Userheadimginfo(source structs.UserHeadImgInfo) *protolib.DougaInfo_Userheadimginfo {
	var rpcprotoDougaInfo_Userheadimginfo protolib.DougaInfo_Userheadimginfo
	rpcprotoDougaInfo_Userheadimginfo.Width = protolib.IntToInt32(source.Width)
	rpcprotoDougaInfo_Userheadimginfo.Height = protolib.IntToInt32(source.Height)
	rpcprotoDougaInfo_Userheadimginfo.Size = protolib.IntToInt32(source.Size)
	rpcprotoDougaInfo_Userheadimginfo.Type = protolib.IntToInt32(source.Type)
	rpcprotoDougaInfo_Userheadimginfo.Animated = source.Animated
	rpcprotoDougaInfo_Userheadimginfo.ThumbnailImage = c.structsThumbnailImageToPRpcprotoDougaInfo_Thumbnailimage1(source.ThumbnailImage)
	rpcprotoDougaInfo_Userheadimginfo.ExpandedImage = c.structsExpandedImageToPRpcprotoDougaInfo_Expandedimage1(source.ExpandedImage)
	rpcprotoDougaInfo_Userheadimginfo.ThumbnailImageCdnURL = source.ThumbnailImageCdnURL
	return &rpcprotoDougaInfo_Userheadimginfo
}
func (c *ConverterImpl) structsUserToPRpcprotoDougaInfo_User(source structs.User) *protolib.DougaInfo_User {
	var rpcprotoDougaInfo_User protolib.DougaInfo_User
	rpcprotoDougaInfo_User.Action = protolib.IntToInt32(source.Action)
	rpcprotoDougaInfo_User.Href = source.Href
	var pRpcprotoDougaInfo_HeadcdnurlsList []*protolib.DougaInfo_Headcdnurls
	if source.HeadCdnUrls != nil {
		pRpcprotoDougaInfo_HeadcdnurlsList = make([]*protolib.DougaInfo_Headcdnurls, len(source.HeadCdnUrls))
		for i := 0; i < len(source.HeadCdnUrls); i++ {
			pRpcprotoDougaInfo_HeadcdnurlsList[i] = c.structsHeadCdnUrlsToPRpcprotoDougaInfo_Headcdnurls(source.HeadCdnUrls[i])
		}
	}
	rpcprotoDougaInfo_User.HeadCdnUrls = pRpcprotoDougaInfo_HeadcdnurlsList
	rpcprotoDougaInfo_User.IsJoinUpCollege = source.IsJoinUpCollege
	rpcprotoDougaInfo_User.FollowingCountValue = protolib.IntToInt32(source.FollowingCountValue)
	rpcprotoDougaInfo_User.ContributeCountValue = protolib.IntToInt32(source.ContributeCountValue)
	rpcprotoDougaInfo_User.FanCount = source.FanCount
	rpcprotoDougaInfo_User.SocialMedal = c.structsSocialMedalToPRpcprotoDougaInfo_Socialmedal(source.SocialMedal)
	rpcprotoDougaInfo_User.AvatarImage = source.AvatarImage
	rpcprotoDougaInfo_User.UserHeadImgInfo = c.structsUserHeadImgInfoToPRpcprotoDougaInfo_Userheadimginfo(source.UserHeadImgInfo)
	rpcprotoDougaInfo_User.IsFollowed = source.IsFollowed
	rpcprotoDougaInfo_User.IsFollowing = source.IsFollowing
	rpcprotoDougaInfo_User.AvatarFrame = protolib.IntToInt32(source.AvatarFrame)
	rpcprotoDougaInfo_User.ContributeCount = source.ContributeCount
	rpcprotoDougaInfo_User.AvatarFramePcImg = source.AvatarFramePcImg
	rpcprotoDougaInfo_User.AvatarFrameMobileImg = source.AvatarFrameMobileImg
	rpcprotoDougaInfo_User.NameStyle = source.NameStyle
	rpcprotoDougaInfo_User.SexTrend = protolib.IntToInt32(source.SexTrend)
	var int32List []int32
	if source.VerifiedTypes != nil {
		int32List = make([]int32, len(source.VerifiedTypes))
		for j := 0; j < len(source.VerifiedTypes); j++ {
			int32List[j] = protolib.IntToInt32(source.VerifiedTypes[j])
		}
	}
	rpcprotoDougaInfo_User.VerifiedTypes = int32List
	rpcprotoDougaInfo_User.HeadURL = source.HeadURL
	rpcprotoDougaInfo_User.FanCountValue = protolib.IntToInt32(source.FanCountValue)
	rpcprotoDougaInfo_User.FollowingStatus = protolib.IntToInt32(source.FollowingStatus)
	rpcprotoDougaInfo_User.NameColor = protolib.IntToInt32(source.NameColor)
	rpcprotoDougaInfo_User.FollowingCount = source.FollowingCount
	rpcprotoDougaInfo_User.Gender = protolib.IntToInt32(source.Gender)
	rpcprotoDougaInfo_User.Name = source.Name
	rpcprotoDougaInfo_User.Signature = source.Signature
	rpcprotoDougaInfo_User.ID = source.ID
	return &rpcprotoDougaInfo_User
}
func (c *ConverterImpl) structsVideoListToPRpcprotoDougaInfo_Videolist(source structs.VideoList) *protolib.DougaInfo_Videolist {
	var rpcprotoDougaInfo_Videolist protolib.DougaInfo_Videolist
	rpcprotoDougaInfo_Videolist.Priority = protolib.IntToInt32(source.Priority)
	rpcprotoDougaInfo_Videolist.Title = source.Title
	rpcprotoDougaInfo_Videolist.DanmakuCountShow = source.DanmakuCountShow
	rpcprotoDougaInfo_Videolist.UserPlayedSeconds = protolib.IntToInt32(source.UserPlayedSeconds)
	rpcprotoDougaInfo_Videolist.SizeType = protolib.IntToInt32(source.SizeType)
	rpcprotoDougaInfo_Videolist.DanmakuGuidePosition = protolib.IntToInt32(source.DanmakuGuidePosition)
	rpcprotoDougaInfo_Videolist.SourceStatus = protolib.IntToInt32(source.SourceStatus)
	rpcprotoDougaInfo_Videolist.DurationMillis = protolib.IntToInt32(source.DurationMillis)
	rpcprotoDougaInfo_Videolist.UploadTime = source.UploadTime
	rpcprotoDougaInfo_Videolist.VisibleType = protolib.IntToInt32(source.VisibleType)
	rpcprotoDougaInfo_Videolist.DanmakuCount = protolib.IntToInt32(source.DanmakuCount)
	rpcprotoDougaInfo_Videolist.FileName = source.FileName
	rpcprotoDougaInfo_Videolist.ID = source.ID
	return &rpcprotoDougaInfo_Videolist
}

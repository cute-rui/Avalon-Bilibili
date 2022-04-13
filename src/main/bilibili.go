package main

import (
    "avalon-bilibili/src/bilibili"
    "context"
    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/status"
)

const (
    NOTMATCHED = -1
    SHORTLINK  = 0
    AVVIDEO    = 1
    BVVIDEO    = 2
    SEASON     = 3
    EPISODE    = 4
    MEDIA      = 5
    AUDIO      = 6
    ARTICLE    = 7
    COLLECTION = 8
)

const (
    AV_REGEX         = `(?:av|AV)\d+`
    BV_REGEX         = `((bv|BV)[0-9A-z]{10})`
    DYNAMIC_REGEX    = `(t\.bilibili\.com/(h5/dynamic/detail/)?)([0-9]{18})`
    ROOM_REGEX       = `(live\.bilibili\.com/)(\d+)`
    SHORT_LINK_REGEX = `((b23\.tv\/)[0-9A-z]+)`
    SPACE_REGEX      = `(space\.bilibili\.com/)(\d+)`
    SEASON_REGEX     = `(?:bilibili\.com/bangumi/play/ss)\d+`
    EPISODE_REGEX    = `(?:bilibili\.com/bangumi/play/ep)\d+`
    MEDIA_REGEX      = `(?:bilibili\.com/bangumi/media/md)\d+`
    COLLECTION_REGEX = `(?:channel/collectiondetail\?sid=)\d+`
)

type BilibiliServiceServer struct {
    bilibili.UnimplementedBilibiliServer
}

func (server BilibiliServiceServer) DoDownloadQuery(ctx context.Context, param *bilibili.Param) (*bilibili.Query, error) {
    checkCollection := param.GetCheckCollection()
    id, kind, err := parseBilibiliURL(param.GetURL())
    if err != nil || kind == NOTMATCHED || id == `` {
        return nil, status.Error(codes.InvalidArgument, `NOT MATCHED`)
    }
    
    if kind == SHORTLINK {
        id, err = getB23TVLocation(id)
        if err != nil {
            return nil, status.Error(400, err.Error())
        }
        
        id, kind, err = parseBilibiliURL(id)
        if err != nil {
            return nil, status.Error(400, err.Error())
        }
    }
    
    var QueryResult bilibili.Query
    switch kind {
    case AVVIDEO:
        err := DoBilibiliAVQuery(id, &QueryResult, checkCollection)
        if err != nil {
            return nil, status.Error(400, err.Error())
        }
    case BVVIDEO:
        err := DoBilibiliBVQuery(id, &QueryResult, checkCollection)
        if err != nil {
            return nil, status.Error(400, err.Error())
        }
    case SEASON:
        err := DoBilibiliSeasonQuery(id, &QueryResult)
        if err != nil {
            return nil, status.Error(400, err.Error())
        }
    case EPISODE:
        err := DoBilibiliEpisodeQuery(id, &QueryResult)
        if err != nil {
            return nil, status.Error(400, err.Error())
        }
    
    case MEDIA:
        err := DoBilibiliMediaQuery(id, &QueryResult)
        if err != nil {
            return nil, status.Error(400, err.Error())
        }
    case COLLECTION:
        err := DoBilibiliCollectionQuery(id, &QueryResult)
        if err != nil {
            return nil, status.Error(400, err.Error())
        }
    default:
    }
    
    return &QueryResult, nil
}

func (server BilibiliServiceServer) GetInfo(ctx context.Context, param *bilibili.Param) (*bilibili.Info, error) {
    id, kind, err := parseBilibiliURL(param.GetURL())
    if err != nil || kind == NOTMATCHED || id == `` {
        return nil, status.Error(codes.InvalidArgument, `NOT MATCHED`)
    }
    
    if kind == SHORTLINK {
        id, err = getB23TVLocation(id)
        if err != nil {
            return nil, status.Error(400, err.Error())
        }
        
        id, kind, err = parseBilibiliURL(id)
        if err != nil {
            return nil, status.Error(400, err.Error())
        }
    }
    
    var m bilibili.Info
    switch kind {
    case AVVIDEO:
        err := fetchBilibiliAVInfo(id, &m)
        if err != nil {
            return nil, status.Error(400, err.Error())
        }
    case BVVIDEO:
        err := fetchBilibiliBVInfo(id, &m)
        if err != nil {
            return nil, status.Error(400, err.Error())
        }
    case SEASON:
        err := fetchBilibiliSeasonInfo(id, &m)
        if err != nil {
            return nil, status.Error(400, err.Error())
        }
    case EPISODE:
        err := fetchBilibiliEpisodeInfo(id, &m)
        if err != nil {
            return nil, status.Error(400, err.Error())
        }
    
    case MEDIA:
        err := fetchBilibiliMediaInfo(id, &m)
        if err != nil {
            return nil, status.Error(400, err.Error())
        }
    case COLLECTION:
        err := fetchBilibiliCollectionInfo(id, &m)
        if err != nil {
            return nil, status.Error(400, err.Error())
        }
    default:
        
    }
    
    return &m, nil
}

func (server BilibiliServiceServer) GetDownloadInfo(ctx context.Context, param *bilibili.Param) (*bilibili.DownloadInfo, error) {
    id, kind := param.GetID(), param.GetType()
    
    var DownloadInfoResult bilibili.DownloadInfo
    switch {
    case kind == bilibili.DataType_Video:
        err := fetchBilibiliBVDownloadInfo(id, param.GetCID(), &DownloadInfoResult)
        if err != nil {
            return nil, status.Error(400, err.Error())
        }
    case kind == bilibili.DataType_Season:
        err := fetchBilibiliEpisodeDownloadInfo(id, &DownloadInfoResult, param.GetRegion())
        if err != nil {
            return nil, status.Error(400, err.Error())
        }
    default:
    }
    
    return &DownloadInfoResult, nil
}

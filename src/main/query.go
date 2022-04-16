package main

import (
    "avalon-bilibili/src/bilibili"
    "strconv"
    "strings"
    "time"
)

func DoBilibiliAVQuery(str string, m *bilibili.Query, checkCollection bool) error {
    s, err := fetchAV(str)
    if err != nil {
        return err
    }
    
    if checkCollection {
        if s.Data.UgcSeason.Id != 0 {
            return DoBilibiliCollectionQuery(strconv.Itoa(s.Data.UgcSeason.Id), m)
        }
    }
    
    return videoQuery(s, m)
}

func DoBilibiliBVQuery(str string, m *bilibili.Query, checkCollection bool) error {
    s, err := fetchBV(str)
    if err != nil {
        return err
    }
    
    if checkCollection {
        if s.Data.UgcSeason.Id != 0 {
            return DoBilibiliCollectionQuery(strconv.Itoa(s.Data.UgcSeason.Id), m)
        }
    }
    
    return videoQuery(s, m)
}

func DoBilibiliSeasonQuery(str string, m *bilibili.Query) error {
    s, err, region := fetchSeason(str)
    if err != nil {
        return err
    }
    
    return seasonQuery(s, m, region)
}

func DoBilibiliEpisodeQuery(str string, m *bilibili.Query) error {
    s, err, region := fetchEpisode(str)
    if err != nil {
        return err
    }
    
    return seasonQuery(s, m, region)
}

func DoBilibiliMediaQuery(str string, m *bilibili.Query) error {
    s, err := fetchMedia(str)
    if err != nil {
        return err
    }
    
    return DoBilibiliSeasonQuery(strconv.Itoa(s.Result.Media.SeasonId), m)
    
}

func DoBilibiliCollectionQuery(str string, m *bilibili.Query) error {
    s, err := fetchCollection(str)
    if err != nil {
        return err
    }
    
    err = collectionQuery(s, m)
    if m != nil {
        m.Type = bilibili.DataType_Collection
    }
    
    return err
}

func videoQuery(s *singleVideoInfo, m *bilibili.Query) error {
    author := s.Data.Owner.Name
    for i := range s.Data.Pages {
        Detail := bilibili.QueryInfo{
            Index:  int32(i),
            ID:     strconv.Itoa(s.Data.Pages[i].Cid),
            Author: &author,
        }
        m.Detail = append(m.GetDetail(), &Detail)
    }
    
    m.Code = 0
    m.Msg = `OK`
    m.Type = bilibili.DataType_Video
    m.ID = s.Data.Bvid
    m.Author = &author
    
    return nil
}

func seasonQuery(s *episodeOrSeasonInfo, m *bilibili.Query, region bilibili.Region) error {
    if limit, r := checkRegion(s.Result.SeasonTitle); limit || region == bilibili.Region_CN {
        region = r
    }
    
    for i := range s.Result.Episodes {
        Detail := bilibili.QueryInfo{
            Index:  int32(i),
            ID:     strconv.Itoa(s.Result.Episodes[i].Id),
            BVID:   &s.Result.Episodes[i].Bvid,
            Region: &region,
        }
        m.Detail = append(m.GetDetail(), &Detail)
    }
    
    m.Code = 0
    m.Msg = `OK`
    m.Type = bilibili.DataType_Season
    m.ID = strconv.Itoa(s.Result.SeasonId)
    m.CollectionTitle = &s.Result.SeasonTitle
    
    if strings.Contains(s.Result.NewEp.Desc, `完结`) || strings.Contains(s.Result.NewEp.Desc, `全`) {
        var b = true
        m.IsEnd = &b
    }
    return nil
}

func collectionQuery(s *collectionType, m *bilibili.Query) error {
    for i := range s.Data.Archives {
        time.Sleep(1 * time.Second)
        videos, err := DoBilibiliCollectionVideoQuery(s.Data.Archives[i].Bvid)
        if err != nil {
            return err
        }
        
        m.Detail = append(m.GetDetail(), videos...)
    }
    
    m.Code = 0
    m.Msg = `OK`
    m.ID = strconv.Itoa(s.Data.Meta.SeasonId)
    m.Type = bilibili.DataType_Collection
    m.CollectionTitle = &s.Data.Meta.Name
    if len(m.Detail) > 0 {
        author := m.Detail[0].GetAuthor()
        m.Author = &author
    }
    
    return nil
}

func DoBilibiliCollectionVideoQuery(str string) ([]*bilibili.QueryInfo, error) {
    s, err := fetchBV(str)
    if err != nil {
        return nil, err
    }
    
    var info []*bilibili.QueryInfo
    for i := range s.Data.Pages {
        Detail := bilibili.QueryInfo{
            Index:  int32(i),
            ID:     strconv.Itoa(s.Data.Pages[i].Cid),
            BVID:   &s.Data.Bvid,
            Author: &s.Data.Owner.Name,
        }
        
        info = append(info, &Detail)
    }
    
    return info, nil
}

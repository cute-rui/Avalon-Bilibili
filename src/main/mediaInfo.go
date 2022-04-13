package main

import (
    "avalon-bilibili/src/bilibili"
    "strconv"
)

func fetchBilibiliAVInfo(str string, m *bilibili.Info) error {
    s, err := fetchAV(str)
    if err != nil {
        return err
    }
    
    m.Code = 0
    m.Msg = `OK`
    m.Type = bilibili.DataType_Video
    m.Picture = &s.Data.Pic
    m.BV = &s.Data.Bvid
    m.AV = &s.Data.Aid
    m.Title = &s.Data.Title
    m.Author = &s.Data.Owner.Name
    m.CreateTime = &s.Data.Ctime
    m.PublicTime = &s.Data.Pubdate
    m.Duration = &s.Data.Duration
    m.Description = &s.Data.Desc
    m.Dynamic = &s.Data.Dynamic
    
    return nil
}

func fetchBilibiliBVInfo(str string, m *bilibili.Info) error {
    s, err := fetchBV(str)
    if err != nil {
        return err
    }
    
    m.Code = 0
    m.Msg = `OK`
    m.Type = bilibili.DataType_Video
    m.Picture = &s.Data.Pic
    m.BV = &s.Data.Bvid
    m.AV = &s.Data.Aid
    m.Title = &s.Data.Title
    m.Author = &s.Data.Owner.Name
    m.CreateTime = &s.Data.Ctime
    m.PublicTime = &s.Data.Pubdate
    m.Duration = &s.Data.Duration
    m.Description = &s.Data.Desc
    m.Dynamic = &s.Data.Dynamic
    
    return nil
}

func fetchBilibiliEpisodeInfo(str string, m *bilibili.Info) error {
    s, err, _ := fetchEpisode(str)
    if err != nil {
        return err
    }
    
    m.Code = 0
    m.Msg = `OK`
    m.Type = bilibili.DataType_Season
    m.Picture = &s.Result.Cover
    m.Title = &s.Result.Title
    m.Evaluate = &s.Result.Evaluate
    
    var area string
    if s.Result.Areas != nil && len(s.Result.Areas) > 0 {
        area = s.Result.Areas[0].Name
        for i := range s.Result.Areas {
            if i == 0 {
                continue
            }
            area = StringBuilder(area, `，`, s.Result.Areas[i].Name)
        }
    }
    
    m.Area = &area
    m.SerialStatusDescription = &s.Result.NewEp.Desc
    m.ShareURL = &s.Result.ShareUrl
    
    return nil
}

func fetchBilibiliSeasonInfo(str string, m *bilibili.Info) error {
    s, err, _ := fetchSeason(str)
    if err != nil {
        return err
    }
    
    m.Code = 0
    m.Msg = `OK`
    m.Type = bilibili.DataType_Season
    m.Picture = &s.Result.Cover
    m.Title = &s.Result.Title
    m.Evaluate = &s.Result.Evaluate
    
    var area string
    if s.Result.Areas != nil && len(s.Result.Areas) > 0 {
        area = s.Result.Areas[0].Name
        for i := range s.Result.Areas {
            if i == 0 {
                continue
            }
            area = StringBuilder(area, `，`, s.Result.Areas[i].Name)
        }
    }
    
    m.Area = &area
    m.SerialStatusDescription = &s.Result.NewEp.Desc
    m.ShareURL = &s.Result.ShareUrl
    
    return nil
}

func fetchBilibiliMediaInfo(str string, m *bilibili.Info) error {
    s, err := fetchMedia(str)
    if err != nil {
        return err
    }
    
    m.Code = 0
    m.Msg = `OK`
    m.Type = bilibili.DataType_Media
    m.Picture = &s.Result.Media.Cover
    m.Title = &s.Result.Media.Title
    
    var area string
    if s.Result.Media.Areas != nil && len(s.Result.Media.Areas) > 0 {
        area = s.Result.Media.Areas[0].Name
        for i := range s.Result.Media.Areas {
            if i == 0 {
                continue
            }
            area = StringBuilder(area, `，`, s.Result.Media.Areas[i].Name)
        }
    }
    
    m.Area = &area
    m.SerialStatusDescription = &s.Result.Media.NewEp.IndexShow
    m.ShareURL = &s.Result.Media.ShareUrl
    m.MediaType = &s.Result.Media.TypeName
    
    return nil
}

func fetchBilibiliCollectionInfo(str string, m *bilibili.Info) error {
    s, err := fetchCollection(str)
    if err != nil {
        return err
    }
    
    m.Code = 0
    m.Code = 0
    m.Msg = `OK`
    m.Type = bilibili.DataType_Collection
    m.Picture = &s.Data.Meta.Cover
    m.Title = &s.Data.Meta.Name
    m.Description = &s.Data.Meta.Description
    url := StringBuilder(`https://space.bilibili.com/`, strconv.Itoa(s.Data.Meta.Mid), `/channel/collectiondetail?sid=`, strconv.Itoa(s.Data.Meta.SeasonId))
    m.ShareURL = &url
    
    return nil
}

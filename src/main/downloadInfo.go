package main

import (
    "avalon-bilibili/src/bilibili"
    "sort"
    "strconv"
)

func fetchBilibiliBVDownloadInfo(id, cid string, m *bilibili.DownloadInfo) error {
    s, err := fetchBV(id)
    if err != nil {
        return err
    }
    
    return videoDownloadInfoProc(cid, s, m)
}

func fetchBilibiliEpisodeDownloadInfo(id string, m *bilibili.DownloadInfo, region bilibili.Region) error {
    s, err, _ := fetchSpecificRegionEpisode(id, region)
    if err != nil {
        return err
    }
    
    return episodeDownloadInfoProc(id, s, m, region)
}

func videoDownloadInfoProc(id string, s *singleVideoInfo, downloadInfoResult *bilibili.DownloadInfo) error {
    for i := range s.Data.Pages {
        if strconv.Itoa(s.Data.Pages[i].Cid) != id {
            continue
        }
        p, e := reqPlayURL(s.Data.Bvid,
            s.Data.Pages[i].Cid,
            s.Data.Title,
            s.Data.Pages[i].Part,
            int32(i))
        if e != nil {
            return e
        }
        
        downloadInfoResult.Detail = append(downloadInfoResult.Detail, p...)
    }
    
    downloadInfoResult.Code = 0
    downloadInfoResult.Msg = `OK`
    downloadInfoResult.Type = bilibili.DataType_Video
    
    d := Details(downloadInfoResult.Detail)
    sort.Sort(d)
    downloadInfoResult.Detail = d
    
    return nil
}

func episodeDownloadInfoProc(id string, seasonInfo *episodeOrSeasonInfo, downloadInfoResult *bilibili.DownloadInfo, region bilibili.Region) error {
    for i := range seasonInfo.Result.Episodes {
        if strconv.Itoa(seasonInfo.Result.Episodes[i].Id) != id {
            continue
        }
        if region != bilibili.Region_TH {
            p, e := reqPgcURL(
                seasonInfo.Result.Episodes[i].Id,
                seasonInfo.Result.Episodes[i].Cid,
                seasonInfo.Result.Episodes[i].Bvid,
                seasonInfo.Result.Episodes[i].Title,
                seasonInfo.Result.Episodes[i].LongTitle,
                int32(i), region)
            if e != nil {
                return e
            }
            downloadInfoResult.Detail = append(downloadInfoResult.Detail, p...)
        }
        
    }
    
    downloadInfoResult.Code = 0
    downloadInfoResult.Msg = `OK`
    downloadInfoResult.Type = bilibili.DataType_Season
    
    d := Details(downloadInfoResult.Detail)
    sort.Sort(d)
    downloadInfoResult.Detail = d
    
    return nil
}

func GetJsonSubtitieURL(cid, bvid string) (bool, []*bilibili.Subtitle) {
    Subtitles, err := fetchSubtitle(StringBuilder(`cid=`, cid, `&bvid=`, bvid))
    
    if err != nil {
        return false, nil
    }
    
    if len(Subtitles) == 0 || Subtitles == nil {
        return false, nil
    }
    
    return true, Subtitles
}

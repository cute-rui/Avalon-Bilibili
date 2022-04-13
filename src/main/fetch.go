package main

import (
    "avalon-bilibili/src/bilibili"
    jsoniter "github.com/json-iterator/go"
    "github.com/pkg/errors"
    "io/ioutil"
    "math"
    "net/http"
    "strconv"
    "strings"
    "time"
)

func fetchAV(str string) (*singleVideoInfo, error) {
    var videoInfo singleVideoInfo
    
    str = StringBuilder(`aid=`, str)
    
    err := fetch(StringBuilder(`https://api.bilibili.com/x/web-interface/view?`, str), &videoInfo)
    if err != nil {
        return nil, err
    }
    
    if videoInfo.Code != 0 {
        return nil, errors.New(videoInfo.Message)
    }
    
    return &videoInfo, nil
}

func fetchBV(str string) (*singleVideoInfo, error) {
    var videoInfo singleVideoInfo
    
    str = StringBuilder(`bvid=`, str)
    
    err := fetch(StringBuilder(`https://api.bilibili.com/x/web-interface/view?`, str), &videoInfo)
    if err != nil {
        return nil, err
    }
    
    if videoInfo.Code != 0 {
        return nil, errors.New(videoInfo.Message)
    }
    
    return &videoInfo, nil
}

func fetchSeason(str string) (*episodeOrSeasonInfo, error, bilibili.Region) {
    var s episodeOrSeasonInfo
    
    str = StringBuilder(`season_id=`, str)
    
    err := fetch(StringBuilder(`https://api.bilibili.com/pgc/view/web/season?`, str), &s)
    if err != nil {
        return nil, err, bilibili.Region_CN
    }
    
    if s.Code != 0 {
        if s.Code == -404 {
            hostname := Conf.GetString(`proxy.region.hk`)
            if hostname == `` {
                return nil, errors.New(`not having proxy server`), bilibili.Region_CN
            }
            
            season, err, region := fetchHongKongEpisode(StringBuilder(`https://`, hostname, `/pgc/view/web/season?`, str))
            if err != nil {
                return nil, err, bilibili.Region_CN
            }
            
            if region != bilibili.Region_CN {
                return season, err, region
            }
            
            hostname = Conf.GetString(`proxy.region.tw`)
            if hostname == `` {
                return nil, errors.New(`not having proxy server`), bilibili.Region_CN
            }
            
            season, err, region = fetchTaiwanEpisode(StringBuilder(`https://`, hostname, `/pgc/view/web/season?`, str))
            if err != nil {
                return nil, err, bilibili.Region_CN
            }
            
            if region != bilibili.Region_CN {
                return season, err, region
            }
        }
        
        return nil, errors.New(s.Message), bilibili.Region_CN
    }
    
    if l, region := checkRegion(s.Result.SeasonTitle); l {
        if region == bilibili.Region_HK {
            hostname := Conf.GetString(`proxy.region.hk`)
            if hostname == `` {
                return nil, errors.New(`not having proxy server`), bilibili.Region_HK
            }
            
            return fetchHongKongEpisode(StringBuilder(`https://`, hostname, `/pgc/view/web/season?`, str))
        }
        
        if region == bilibili.Region_TW {
            hostname := Conf.GetString(`proxy.region.tw`)
            if hostname == `` {
                return nil, errors.New(`not having proxy server`), bilibili.Region_TW
            }
            
            return fetchTaiwanEpisode(StringBuilder(`https://`, hostname, `/pgc/view/web/season?`, str))
        }
    }
    
    return &s, nil, bilibili.Region_CN
}

func fetchSpecificRegionEpisode(str string, region bilibili.Region) (*episodeOrSeasonInfo, error, bilibili.Region) {
    var s episodeOrSeasonInfo
    
    str = StringBuilder(`ep_id=`, str)
    
    switch region {
    case bilibili.Region_CN:
        err := fetch(StringBuilder(`https://api.bilibili.com/pgc/view/web/season?`, str), &s)
        if err != nil {
            return nil, err, bilibili.Region_CN
        }
    case bilibili.Region_HK:
        hostname := Conf.GetString(`proxy.region.hk`)
        if hostname == `` {
            return nil, errors.New(`not having proxy server`), bilibili.Region_CN
        }
        
        err := fetch(StringBuilder(`https://`, hostname, `/pgc/view/web/season?`, str), &s)
        if err != nil {
            return nil, err, bilibili.Region_CN
        }
    case bilibili.Region_TW:
        hostname := Conf.GetString(`proxy.region.tw`)
        if hostname == `` {
            return nil, errors.New(`not having proxy server`), bilibili.Region_CN
        }
        
        err := fetch(StringBuilder(`https://`, hostname, `/pgc/view/web/season?`, str), &s)
        if err != nil {
            return nil, err, bilibili.Region_CN
        }
    }
    
    if s.Code != 0 {
        return nil, errors.New(s.Message), bilibili.Region_CN
    }
    
    return &s, nil, region
}

func fetchEpisode(str string) (*episodeOrSeasonInfo, error, bilibili.Region) {
    var s episodeOrSeasonInfo
    
    str = StringBuilder(`ep_id=`, str)
    
    err := fetch(StringBuilder(`https://api.bilibili.com/pgc/view/web/season?`, str), &s)
    if err != nil {
        return nil, err, bilibili.Region_CN
    }
    
    if s.Code != 0 {
        if s.Code == -404 {
            hostname := Conf.GetString(`proxy.region.hk`)
            if hostname == `` {
                return nil, errors.New(`not having proxy server`), bilibili.Region_CN
            }
            
            season, err, region := fetchHongKongEpisode(StringBuilder(`https://`, hostname, `/pgc/view/web/season?`, str))
            if err != nil {
                return nil, err, bilibili.Region_CN
            }
            
            if region != bilibili.Region_CN {
                return season, err, region
            }
            
            hostname = Conf.GetString(`proxy.region.tw`)
            if hostname == `` {
                return nil, errors.New(`not having proxy server`), bilibili.Region_CN
            }
            
            season, err, region = fetchTaiwanEpisode(StringBuilder(`https://`, hostname, `/pgc/view/web/season?`, str))
            if err != nil {
                return nil, err, bilibili.Region_CN
            }
            
            if region != bilibili.Region_CN {
                return season, err, region
            }
        }
        
        return nil, errors.New(s.Message), bilibili.Region_CN
    }
    
    if l, region := checkRegion(s.Result.SeasonTitle); l {
        if region == bilibili.Region_HK {
            hostname := Conf.GetString(`proxy.region.hk`)
            if hostname == `` {
                return nil, errors.New(`not having proxy server`), bilibili.Region_HK
            }
            
            return fetchHongKongEpisode(StringBuilder(`https://`, hostname, `/pgc/view/web/season?`, str))
        }
        
        if region == bilibili.Region_TW {
            hostname := Conf.GetString(`proxy.region.tw`)
            if hostname == `` {
                return nil, errors.New(`not having proxy server`), bilibili.Region_TW
            }
            
            return fetchTaiwanEpisode(StringBuilder(`https://`, hostname, `/pgc/view/web/season?`, str))
        }
    }
    
    return &s, nil, bilibili.Region_CN
}

func fetchHongKongEpisode(url string) (*episodeOrSeasonInfo, error, bilibili.Region) {
    var seasonInfo episodeOrSeasonInfo
    
    err := fetch(url, &seasonInfo)
    if err != nil {
        return nil, err, bilibili.Region_HK
    }
    
    if seasonInfo.Code == 0 {
        return &seasonInfo, nil, bilibili.Region_HK
    } else if seasonInfo.Code != -404 {
        return nil, errors.New(seasonInfo.Message), bilibili.Region_HK
    }
    
    return nil, nil, bilibili.Region_CN
}

func fetchTaiwanEpisode(url string) (*episodeOrSeasonInfo, error, bilibili.Region) {
    var seasonInfo episodeOrSeasonInfo
    
    err := fetch(url, &seasonInfo)
    if err != nil {
        return nil, err, bilibili.Region_TW
    }
    
    if seasonInfo.Code == 0 {
        return &seasonInfo, nil, bilibili.Region_TW
    } else if seasonInfo.Code != -404 {
        return nil, errors.New(seasonInfo.Message), bilibili.Region_TW
    }
    
    return nil, errors.New(seasonInfo.Message), bilibili.Region_CN
}

func fetchMedia(str string) (*mediaType, error) {
    var s mediaType
    
    err := fetch(StringBuilder(`https://api.bilibili.com/pgc/review/user?media_id=`, str), &s)
    if err != nil {
        return nil, err
    }
    
    if s.Code != 0 {
        return nil, errors.New(s.Message)
    }
    
    return &s, nil
}

func fetchCollection(str string) (*collectionType, error) {
    var s collectionType
    
    err := fetch(StringBuilder(`https://api.bilibili.com/x/polymer/space/seasons_archives_list?mid=1&sort_reverse=false&page_num=1&page_size=100&season_id=`, str), &s)
    if err != nil {
        return nil, err
    }
    
    if s.Code != 0 {
        return nil, errors.New(s.Message)
    }
    
    pages := math.Ceil(float64(s.Data.Page.Total) / 100)
    for i := 2; float64(i) <= pages; i++ {
        var tmp collectionType
        err := fetch(StringBuilder(`https://api.bilibili.com/x/polymer/space/seasons_archives_list?mid=1&sort_reverse=false&page_num=`, strconv.Itoa(i), `&page_size=100&season_id=`, str), &tmp)
        if err != nil {
            return nil, err
        }
        
        if s.Code != 0 {
            return nil, errors.New(s.Message)
        }
        
        s.Data.Aids = append(s.Data.Aids, tmp.Data.Aids...)
        s.Data.Archives = append(s.Data.Archives, tmp.Data.Archives...)
    }
    
    return &s, nil
}

func fetchSubtitle(str string) ([]*bilibili.Subtitle, error) {
    var player playerInfo
    
    err := fetch(StringBuilder(`https://api.bilibili.com/x/player/v2?`, str), &player)
    if err != nil {
        return nil, err
    }
    
    if player.Code != 0 {
        return nil, errors.New(player.Message)
    }
    
    var SubtitleInfos []*bilibili.Subtitle
    for i := range player.Data.Subtitle.Subtitles {
        url := player.Data.Subtitle.Subtitles[i].SubtitleUrl
        if strings.HasPrefix(url, `//`) {
            url = StringBuilder(`https:`, url)
        }
        SubtitleInfos = append(SubtitleInfos, &bilibili.Subtitle{
            Index:       int32(i),
            Locale:      player.Data.Subtitle.Subtitles[i].Lan,
            LocaleText:  player.Data.Subtitle.Subtitles[i].LanDoc,
            SubtitleUrl: url,
        })
    }
    
    return SubtitleInfos, nil
}

type FetchOption func(*http.Request)

func WithHeader(n, v string) FetchOption {
    return func(r *http.Request) {
        r.Header.Add(n, v)
    }
}

func SetHost(v string) FetchOption {
    return func(request *http.Request) {
        request.Host = v
    }
}

func fetch(URL string, result interface{}, fetchOption ...FetchOption) error {
    req, err := http.NewRequest(`GET`, URL, nil)
    if err != nil {
        return err
    }
    
    req.AddCookie(&http.Cookie{Name: `SESSDATA`, Value: Conf.GetString(`bilibili.sessdata`), Path: `/`, Expires: time.Now().AddDate(0, 0, 1)})
    req.AddCookie(&http.Cookie{Name: `CURRENT_FNVAL`, Value: `4048`, Path: `/`, Expires: time.Now().AddDate(0, 0, 1)})
    req.AddCookie(&http.Cookie{Name: `CURRENT_QUALITY`, Value: `120`, Path: `/`, Expires: time.Now().AddDate(0, 0, 1)})
    req.Header.Add(`Referer`, `https://www.bilibili.com`)
    req.Header.Add(`User-Agent`, "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/97.0.4692.99 Safari/537.36")
    
    for i := range fetchOption {
        fetchOption[i](req)
    }
    
    client := http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return err
    }
    
    defer resp.Body.Close()
    
    respBytes, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return err
    }
    
    err = jsoniter.Unmarshal(respBytes, result)
    if err != nil {
        return err
    }
    
    return nil
    
}

func getB23TVLocation(u string) (string, error) {
    if !strings.HasPrefix(u, "http") {
        u = StringBuilder("https://", u)
    }
    
    client := &http.Client{
        CheckRedirect: func(req *http.Request, via []*http.Request) error {
            return http.ErrUseLastResponse
        },
    }
    
    res, err := client.Get(u)
    if err != nil {
        return ``, err
    }
    
    if !(res.StatusCode == 301 || res.StatusCode == 302) {
        return ``, errors.New("no redirection")
    }
    
    return res.Header.Get(`Location`), nil
}

func reqPlayURL(bv string, cid int, title, subTitle string, index int32) ([]*bilibili.PartInfo, error) {
    var d VideoDetail
    
    err := fetch(StringBuilder(`https://api.bilibili.com/x/player/playurl?cid=`, strconv.Itoa(cid), `&fourk=1&fnval=4048&bvid=`, bv), &d)
    if err != nil {
        return nil, err
    }
    
    hasSubtitle, Subtitles := GetJsonSubtitieURL(strconv.Itoa(cid), bv)
    
    var (
        list            []*bilibili.PartInfo
        avc1list        []*bilibili.PartInfo
        dolbyAtmos      bool
        dolbyAtomsIndex int
    )
    
    if d.Data.Dash.Dolby.Type != 0 {
        dolbyAtmos = true
        var tmp int
        for j := range d.Data.Dash.Dolby.Audio {
            if d.Data.Dash.Dolby.Audio[j].Id > tmp {
                tmp = d.Data.Dash.Dolby.Audio[j].Id
                dolbyAtomsIndex = j
            }
        }
    }
    
    tmp, audioIndex := 0, 0
    for j := range d.Data.Dash.Audio {
        if d.Data.Dash.Audio[j].Id > tmp {
            tmp = d.Data.Dash.Audio[j].Id
            audioIndex = j
        }
    }
    
    videoquality := getBestQualityVideo(d.Data.AcceptQuality)
    if videoquality == 126 {
        if deleteDolbyAndGetBestQualityVideo(d.Data.AcceptQuality) != videoquality {
            videoquality = deleteDolbyAndGetBestQualityVideo(d.Data.AcceptQuality)
        }
    }
    
    for j := range d.Data.Dash.Video {
        if d.Data.Dash.Video[j].Id == videoquality {
            var p bilibili.PartInfo
            
            if strings.Contains(d.Data.Dash.Video[j].Codecs, `avc1`) {
                if len(list) != 0 {
                    continue
                }
                
                p.Title = title
                p.SubTitle = subTitle
                p.Index = index
                p.ID = bv
                p.CID = strconv.Itoa(cid)
                p.VideoQuality = getQualityString(videoquality)
                p.VideoURL = d.Data.Dash.Video[j].BaseUrl
                p.AudioURL = d.Data.Dash.Audio[audioIndex].BaseUrl
                if hasSubtitle {
                    p.Subtitles = Subtitles
                }
                
                avc1list = append(avc1list, &p)
                
                if dolbyAtmos {
                    dolby := bilibili.PartInfo{
                        Index:        index,
                        ID:           bv,
                        CID:          strconv.Itoa(cid),
                        Title:        title,
                        SubTitle:     StringBuilder(p.GetSubTitle(), ` - Dolby Atmos`),
                        VideoQuality: getQualityString(videoquality),
                        VideoURL:     d.Data.Dash.Video[j].BaseUrl,
                        AudioURL:     d.Data.Dash.Dolby.Audio[dolbyAtomsIndex].BaseUrl,
                    }
                    
                    if hasSubtitle {
                        dolby.Subtitles = Subtitles
                    }
                    
                    avc1list = append(avc1list, &dolby)
                }
            } else {
                if strings.Contains(d.Data.Dash.Video[j].Codecs, `av01`) {
                    p.SubTitle = StringBuilder(subTitle, ` - av01 codec - unstable`)
                } else {
                    p.SubTitle = subTitle
                }
                
                p.Title = title
                p.Index = index
                p.ID = bv
                p.CID = strconv.Itoa(cid)
                p.VideoQuality = getQualityString(videoquality)
                p.VideoURL = d.Data.Dash.Video[j].BaseUrl
                p.AudioURL = d.Data.Dash.Audio[audioIndex].BaseUrl
                
                if hasSubtitle {
                    p.Subtitles = Subtitles
                }
                
                list = append(list, &p)
                
                if dolbyAtmos {
                    dolby := bilibili.PartInfo{
                        Index:        index,
                        ID:           bv,
                        CID:          strconv.Itoa(cid),
                        Title:        title,
                        SubTitle:     StringBuilder(p.GetSubTitle(), ` - Dolby Atmos`),
                        VideoQuality: getQualityString(videoquality),
                        VideoURL:     d.Data.Dash.Video[j].BaseUrl,
                        AudioURL:     d.Data.Dash.Dolby.Audio[dolbyAtomsIndex].BaseUrl,
                    }
                    
                    if hasSubtitle {
                        dolby.Subtitles = Subtitles
                    }
                    
                    list = append(list, &dolby)
                }
            }
            
        } else if d.Data.Dash.Video[j].Id == 126 {
            var p bilibili.PartInfo
            
            p.Title = title
            p.SubTitle = StringBuilder(subTitle, ` - Dolby Vision`)
            p.Index = index
            p.ID = bv
            p.CID = strconv.Itoa(cid)
            p.VideoQuality = getQualityString(126)
            p.VideoURL = d.Data.Dash.Video[j].BaseUrl
            
            if dolbyAtmos {
                p.AudioURL = d.Data.Dash.Dolby.Audio[dolbyAtomsIndex].BaseUrl
            } else {
                p.AudioURL = d.Data.Dash.Audio[audioIndex].BaseUrl
            }
            
            if hasSubtitle {
                p.Subtitles = Subtitles
            }
            
            list = append(list, &p)
        }
    }
    
    if len(list) == 0 && len(avc1list) != 0 {
        return avc1list, nil
    }
    
    return list, nil
}

func reqPgcURL(ep, cid int, bvid, shortTitle, subTitle string, index int32, region bilibili.Region) ([]*bilibili.PartInfo, error) {
    var d MediaDetail
    
    switch region {
    case bilibili.Region_CN:
        err := fetch(StringBuilder(`https://api.bilibili.com/pgc/player/web/playurl?fourk=1&fnval=4048&ep_id=`, strconv.Itoa(ep)), &d)
        if err != nil {
            return nil, err
        }
    case bilibili.Region_HK:
        hostname := Conf.GetString(`proxy.region.hk`)
        if hostname == `` {
            return nil, errors.New(`not having proxy server`)
        }
        
        err := fetch(StringBuilder(`https://`, hostname, `/pgc/player/web/playurl?fourk=1&fnval=4048&ep_id=`, strconv.Itoa(ep)), &d)
        if err != nil {
            return nil, err
        }
    case bilibili.Region_TW:
        hostname := Conf.GetString(`proxy.region.tw`)
        if hostname == `` {
            return nil, errors.New(`not having proxy server`)
        }
        
        err := fetch(StringBuilder(`https://`, hostname, `/pgc/player/web/playurl?fourk=1&fnval=4048&ep_id=`, strconv.Itoa(ep)), &d)
        if err != nil {
            return nil, err
        }
    }
    
    hasSubtitle, Subtitles := GetJsonSubtitieURL(strconv.Itoa(cid), bvid)
    
    var (
        list            []*bilibili.PartInfo
        avc1list        []*bilibili.PartInfo
        dolbyAtmos      bool
        dolbyAtomsIndex int
    )
    
    if d.Data.Dash.Dolby.Type != 0 {
        dolbyAtmos = true
        var tmp int
        for j := range d.Data.Dash.Dolby.Audio {
            if d.Data.Dash.Dolby.Audio[j].Id > tmp {
                tmp = d.Data.Dash.Dolby.Audio[j].Id
                dolbyAtomsIndex = j
            }
        }
    }
    
    tmp, audioIndex := 0, 0
    for j := range d.Data.Dash.Audio {
        if d.Data.Dash.Audio[j].Id > tmp {
            tmp = d.Data.Dash.Audio[j].Id
            audioIndex = j
        }
    }
    
    videoquality := getBestQualityVideo(d.Data.AcceptQuality)
    if videoquality == 126 {
        if deleteDolbyAndGetBestQualityVideo(d.Data.AcceptQuality) != videoquality {
            videoquality = deleteDolbyAndGetBestQualityVideo(d.Data.AcceptQuality)
        }
    }
    
    for j := range d.Data.Dash.Video {
        if d.Data.Dash.Video[j].Id == videoquality {
            var p bilibili.PartInfo
            
            if strings.Contains(d.Data.Dash.Video[j].Codecs, `avc1`) {
                if len(list) != 0 {
                    continue
                }
                
                p.ID = bvid
                p.CID = strconv.Itoa(cid)
                p.Title = subTitle
                p.SubTitle = shortTitle
                p.Index = index
                p.VideoQuality = getQualityString(videoquality)
                p.VideoURL = d.Data.Dash.Video[j].BaseUrl
                p.AudioURL = d.Data.Dash.Audio[audioIndex].BaseUrl
                
                if hasSubtitle {
                    p.Subtitles = Subtitles
                }
                
                avc1list = append(avc1list, &p)
                
                if dolbyAtmos {
                    dolby := bilibili.PartInfo{
                        Index:        index,
                        ID:           bvid,
                        CID:          strconv.Itoa(cid),
                        Title:        StringBuilder(p.GetTitle(), ` - Dolby Atmos`),
                        SubTitle:     shortTitle,
                        VideoQuality: getQualityString(videoquality),
                        VideoURL:     d.Data.Dash.Video[j].BaseUrl,
                        AudioURL:     d.Data.Dash.Dolby.Audio[dolbyAtomsIndex].BaseUrl,
                    }
                    
                    if hasSubtitle {
                        dolby.Subtitles = Subtitles
                    }
                    
                    avc1list = append(avc1list, &dolby)
                }
            } else {
                if strings.Contains(d.Data.Dash.Video[j].Codecs, `av01`) {
                    p.Title = StringBuilder(subTitle, ` - av01 codec - unstable`)
                } else {
                    p.Title = subTitle
                }
                
                p.ID = bvid
                p.CID = strconv.Itoa(cid)
                p.SubTitle = shortTitle
                p.Index = index
                p.VideoQuality = getQualityString(videoquality)
                p.VideoURL = d.Data.Dash.Video[j].BaseUrl
                p.AudioURL = d.Data.Dash.Audio[audioIndex].BaseUrl
                
                if hasSubtitle {
                    p.Subtitles = Subtitles
                }
                
                list = append(list, &p)
                
                if dolbyAtmos {
                    dolby := bilibili.PartInfo{
                        Index:        index,
                        ID:           bvid,
                        CID:          strconv.Itoa(cid),
                        Title:        StringBuilder(p.GetTitle(), ` - Dolby Atmos`),
                        SubTitle:     shortTitle,
                        VideoQuality: getQualityString(videoquality),
                        VideoURL:     d.Data.Dash.Video[j].BaseUrl,
                        AudioURL:     d.Data.Dash.Dolby.Audio[dolbyAtomsIndex].BaseUrl,
                    }
                    
                    if hasSubtitle {
                        dolby.Subtitles = Subtitles
                    }
                    
                    list = append(list, &dolby)
                }
            }
        } else if d.Data.Dash.Video[j].Id == 126 {
            var p bilibili.PartInfo
            
            p.Title = StringBuilder(subTitle, ` - Dolby Vision`)
            p.ID = bvid
            p.CID = strconv.Itoa(cid)
            p.SubTitle = shortTitle
            p.Index = index
            p.VideoQuality = getQualityString(videoquality)
            p.VideoURL = d.Data.Dash.Video[j].BaseUrl
            
            if dolbyAtmos {
                p.AudioURL = d.Data.Dash.Dolby.Audio[dolbyAtomsIndex].BaseUrl
            } else {
                p.AudioURL = d.Data.Dash.Audio[audioIndex].BaseUrl
            }
            
            if hasSubtitle {
                p.Subtitles = Subtitles
            }
            
            list = append(list, &p)
        }
    }
    
    if len(list) == 0 && len(avc1list) != 0 {
        return avc1list, nil
    }
    
    return list, nil
}

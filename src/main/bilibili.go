package main

import (
    "avalon-bilibili/src/bilibili"
    "context"
    jsoniter "github.com/json-iterator/go"
    "github.com/pkg/errors"
    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/status"
    "io/ioutil"
    "math"
    "net/http"
    "regexp"
    "sort"
    "strconv"
    "strings"
    "sync"
    "time"
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

func (server BilibiliServiceServer) GetInfo(ctx context.Context, param *bilibili.Param) (*bilibili.Info, error) {
    s, t, err := parseBilibiliURL(param.GetMark())
    if err != nil || t == NOTMATCHED || s == `` {
        return nil, status.Error(codes.InvalidArgument, `NOT MATCHED`)
    }
    
    if t == SHORTLINK {
        s, err = getB23TVLocation(s)
        if err != nil {
            return nil, status.Error(400, err.Error())
        }
        
        s, t, err = parseBilibiliURL(s)
        if err != nil {
            return nil, status.Error(400, err.Error())
        }
    }
    
    var m bilibili.Info
    switch t {
    case AVVIDEO:
        err := fetchBilibiliAVInfo(s, &m)
        if err != nil {
            return nil, status.Error(400, err.Error())
        }
    case BVVIDEO:
        err := fetchBilibiliBVInfo(s, &m)
        if err != nil {
            return nil, status.Error(400, err.Error())
        }
    case SEASON:
        err := fetchBilibiliSeasonInfo(s, &m)
        if err != nil {
            return nil, status.Error(400, err.Error())
        }
    case EPISODE:
        err := fetchBilibiliEpisodeInfo(s, &m)
        if err != nil {
            return nil, status.Error(400, err.Error())
        }
    
    case MEDIA:
        err := fetchBilibiliMediaInfo(s, &m)
        if err != nil {
            return nil, status.Error(400, err.Error())
        }
    case COLLECTION:
        err := fetchBilibiliCollectionInfo(s, &m)
        if err != nil {
            return nil, status.Error(400, err.Error())
        }
    default:
        
    }
    
    return &m, nil
}

func (server BilibiliServiceServer) GetDownloadInfo(ctx context.Context, param *bilibili.Param) (*bilibili.DownloadInfo, error) {
    checkCollection := param.GetCheckCollection()
    s, t, err := parseBilibiliURL(param.GetMark())
    if err != nil || t == NOTMATCHED || s == `` {
        return nil, status.Error(codes.InvalidArgument, `NOT MATCHED`)
    }
    
    if t == SHORTLINK {
        s, err = getB23TVLocation(s)
        if err != nil {
            return nil, status.Error(400, err.Error())
        }
        
        s, t, err = parseBilibiliURL(s)
        if err != nil {
            return nil, status.Error(400, err.Error())
        }
    }
    
    var m bilibili.DownloadInfo
    switch t {
    case AVVIDEO:
        err := fetchBilibiliAVDownloadInfo(s, &m, checkCollection)
        if err != nil {
            return nil, status.Error(400, err.Error())
        }
    case BVVIDEO:
        err := fetchBilibiliBVDownloadInfo(s, &m, checkCollection)
        if err != nil {
            return nil, status.Error(400, err.Error())
        }
    case SEASON:
        err := fetchBilibiliSeasonDownloadInfo(s, &m)
        if err != nil {
            return nil, status.Error(400, err.Error())
        }
    case EPISODE:
        err := fetchBilibiliEpisodeDownloadInfo(s, &m)
        if err != nil {
            return nil, status.Error(400, err.Error())
        }
    
    case MEDIA:
        err := fetchBilibiliMediaDownloadInfo(s, &m)
        if err != nil {
            return nil, status.Error(400, err.Error())
        }
    case COLLECTION:
        err := fetchBilibiliCollectionDownloadInfo(s, &m)
        if err != nil {
            return nil, status.Error(400, err.Error())
        }
    default:
    }
    
    return &m, nil
}

func fetchBilibiliAVDownloadInfo(str string, m *bilibili.DownloadInfo, checkCollection bool) error {
    s, err := fetchAV(str)
    if err != nil {
        return err
    }
    
    if checkCollection {
        if s.Data.UgcSeason.Id != 0 {
            return fetchBilibiliCollectionDownloadInfo(strconv.Itoa(s.Data.UgcSeason.Id), m)
        }
    }
    
    return videoDownloadInfoProc(s, m)
    
}

func fetchBilibiliBVDownloadInfo(str string, m *bilibili.DownloadInfo, checkCollection bool) error {
    s, err := fetchBV(str)
    if err != nil {
        return err
    }
    
    if checkCollection {
        if s.Data.UgcSeason.Id != 0 {
            return fetchBilibiliCollectionDownloadInfo(strconv.Itoa(s.Data.UgcSeason.Id), m)
        }
    }
    
    return videoDownloadInfoProc(s, m)
}

func fetchBilibiliSeasonDownloadInfo(str string, m *bilibili.DownloadInfo) error {
    s, err := fetchSeason(str)
    if err != nil {
        return err
    }
    
    return seasonDownloadInfoProc(s, m)
}

func fetchBilibiliEpisodeDownloadInfo(str string, m *bilibili.DownloadInfo) error {
    s, err := fetchEpisode(str)
    if err != nil {
        return err
    }
    
    return seasonDownloadInfoProc(s, m)
}

func fetchBilibiliMediaDownloadInfo(str string, m *bilibili.DownloadInfo) error {
    s, err := fetchMedia(str)
    if err != nil {
        return err
    }
    
    return fetchBilibiliSeasonDownloadInfo(strconv.Itoa(s.Result.Media.SeasonId), m)
    
}

func fetchBilibiliCollectionDownloadInfo(str string, m *bilibili.DownloadInfo) error {
    s, err := fetchCollection(str)
    if err != nil {
        return err
    }
    
    err = collectionDownloadInfoProc(s, m)
    if m != nil {
        m.Type = bilibili.DataType_Collection
    }
    
    return err
}

func videoDownloadInfoProc(s *videoType, m *bilibili.DownloadInfo) error {
    var (
        WG   sync.WaitGroup
        Lock sync.RWMutex
    )
    var errStr string
    for i := range s.Data.Pages {
        WG.Add(1)
        go func(index int) {
            p, e := reqPlayURL(s.Data.Bvid,
                s.Data.Pages[index].Cid,
                s.Data.Title,
                s.Data.Pages[index].Part,
                int32(index))
            if e != nil {
                errStr += e.Error()
            }
            Lock.Lock()
            m.Detail = append(m.Detail, p...)
            Lock.Unlock()
            WG.Done()
        }(i)
    }
    
    WG.Wait()
    
    if errStr != `` {
        return errors.New(errStr)
    }
    
    m.Code = 0
    m.Msg = `OK`
    m.Type = bilibili.DataType_Video
    m.ID = s.Data.Bvid
    m.Author = &s.Data.Owner.Name
    
    d := Details(m.Detail)
    sort.Sort(d)
    m.Detail = d
    
    return nil
}

func seasonDownloadInfoProc(s *episodeOrSeasonType, m *bilibili.DownloadInfo) error {
    var (
        WG   sync.WaitGroup
        Lock sync.RWMutex
    )
    var errStr string
    for i := range s.Result.Episodes {
        WG.Add(1)
        go func(index int) {
            p, e := reqPgcURL(
                s.Result.Episodes[index].Id,
                s.Result.Episodes[index].Cid,
                s.Result.Episodes[index].Bvid,
                s.Result.Episodes[index].Title,
                s.Result.Episodes[index].LongTitle,
                int32(index))
            if e != nil {
                errStr += e.Error()
            }
            Lock.Lock()
            m.Detail = append(m.Detail, p...)
            Lock.Unlock()
            WG.Done()
        }(i)
    }
    
    WG.Wait()
    
    if errStr != `` {
        return errors.New(errStr)
    }
    
    m.Code = 0
    m.Msg = `OK`
    m.Type = bilibili.DataType_Season
    m.ID = strconv.Itoa(s.Result.SeasonId)
    m.CollectionTitle = &s.Result.SeasonTitle
    
    d := Details(m.Detail)
    sort.Sort(d)
    m.Detail = d
    
    if strings.Contains(s.Result.NewEp.Desc, `完结`) || strings.Contains(s.Result.NewEp.Desc, `全`) {
        var b = true
        m.IsEnd = &b
    }
    
    return nil
}

func collectionDownloadInfoProc(s *collectionType, m *bilibili.DownloadInfo) error {
    for i := range s.Data.Archives {
        time.Sleep(1 * time.Second)
        var tmp bilibili.DownloadInfo
        err := fetchBilibiliBVDownloadInfo(s.Data.Archives[i].Bvid, &tmp, false)
        if err != nil {
            return err
        }
        
        if m.GetType() == bilibili.DataType_Video {
            m.Code = 0
            m.Msg = `OK`
            m.Type = bilibili.DataType_Collection
            m.ID = strconv.Itoa(s.Data.Meta.SeasonId)
            author := tmp.GetAuthor()
            m.Author = &author
            m.CollectionTitle = &s.Data.Meta.Name
        }
        
        m.Detail = append(m.GetDetail(), tmp.GetDetail()...)
    }
    
    d := Details(m.Detail)
    sort.Sort(d)
    m.Detail = d
    
    return nil
}

func reqPlayURL(bv string, cid int, title, subTitle string, index int32) ([]*bilibili.PartInfo, error) {
    var d VideoDetail
    
    err := fetch(StringBuilder(`http://api.bilibili.com/x/player/playurl?cid=`, strconv.Itoa(cid), `&fourk=1&fnval=4048&bvid=`, bv), &d)
    if err != nil {
        return nil, err
    }
    
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
            
            list = append(list, &p)
        }
    }
    
    if len(list) == 0 && len(avc1list) != 0 {
        return avc1list, nil
    }
    
    return list, nil
}

func reqPgcURL(ep, cid int, bvid, shortTitle, subTitle string, index int32) ([]*bilibili.PartInfo, error) {
    var d MediaDetail
    
    err := fetch(StringBuilder(`https://api.bilibili.com/pgc/player/web/playurl?fourk=1&fnval=4048&ep_id=`, strconv.Itoa(ep)), &d)
    if err != nil {
        return nil, err
    }
    
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
                avc1list = append(avc1list, &p)
                
                if dolbyAtmos {
                    dolby := bilibili.PartInfo{
                        Index:        index,
                        ID:           bvid,
                        CID:          strconv.Itoa(cid),
                        Title:        StringBuilder(p.GetSubTitle(), ` - Dolby Atmos`),
                        SubTitle:     shortTitle,
                        VideoQuality: getQualityString(videoquality),
                        VideoURL:     d.Data.Dash.Video[j].BaseUrl,
                        AudioURL:     d.Data.Dash.Dolby.Audio[dolbyAtomsIndex].BaseUrl,
                    }
                    avc1list = append(avc1list, &dolby)
                }
            } else {
                if strings.Contains(d.Data.Dash.Video[j].Codecs, `av01`) {
                    p.Title = StringBuilder(subTitle, `av01 codec - unstable`)
                } else {
                    p.Title = subTitle
                }
                
                p.ID = bvid
                p.CID = strconv.Itoa(cid)
                p.Title = subTitle
                p.SubTitle = shortTitle
                p.Index = index
                p.VideoQuality = getQualityString(videoquality)
                p.VideoURL = d.Data.Dash.Video[j].BaseUrl
                p.AudioURL = d.Data.Dash.Audio[audioIndex].BaseUrl
                list = append(list, &p)
                
                if dolbyAtmos {
                    dolby := bilibili.PartInfo{
                        Index:        index,
                        ID:           bvid,
                        CID:          strconv.Itoa(cid),
                        Title:        StringBuilder(p.GetSubTitle(), ` - Dolby Atmos`),
                        SubTitle:     shortTitle,
                        VideoQuality: getQualityString(videoquality),
                        VideoURL:     d.Data.Dash.Video[j].BaseUrl,
                        AudioURL:     d.Data.Dash.Dolby.Audio[dolbyAtomsIndex].BaseUrl,
                    }
                    list = append(list, &dolby)
                }
            }
        } else if d.Data.Dash.Video[j].Id == 126 {
            var p bilibili.PartInfo
            
            p.Title = StringBuilder(subTitle, ` - Dolby Vision`)
            p.ID = bvid
            p.CID = strconv.Itoa(cid)
            p.Title = subTitle
            p.SubTitle = shortTitle
            p.Index = index
            p.VideoQuality = getQualityString(videoquality)
            p.VideoURL = d.Data.Dash.Video[j].BaseUrl
            
            if dolbyAtmos {
                p.AudioURL = d.Data.Dash.Dolby.Audio[dolbyAtomsIndex].BaseUrl
            } else {
                p.AudioURL = d.Data.Dash.Audio[audioIndex].BaseUrl
            }
            
            list = append(list, &p)
        }
    }
    
    if len(list) == 0 && len(avc1list) != 0 {
        return avc1list, nil
    }
    
    return list, nil
}

func parseBilibiliURL(str string) (string, int, error) {
    switch {
    case strings.Contains(str, "b23.tv"):
        r, err := regexp.Compile(SHORT_LINK_REGEX)
        if err != nil {
            return "", NOTMATCHED, err
        }
        return r.FindString(str), SHORTLINK, err
    case strings.Contains(str, `bilibili.com/bangumi/play/ss`):
        r, err := regexp.Compile(SEASON_REGEX)
        if err != nil {
            return "", NOTMATCHED, err
        }
        return strings.TrimPrefix(r.FindString(str), `bilibili.com/bangumi/play/ss`), SEASON, err
    case strings.Contains(str, `bilibili.com/bangumi/play/ep`):
        r, err := regexp.Compile(EPISODE_REGEX)
        if err != nil {
            return "", NOTMATCHED, err
        }
        return strings.TrimPrefix(r.FindString(str), `bilibili.com/bangumi/play/ep`), EPISODE, err
    case strings.Contains(str, `bilibili.com/bangumi/media/md`):
        r, err := regexp.Compile(MEDIA_REGEX)
        if err != nil {
            return "", NOTMATCHED, err
        }
        return strings.TrimPrefix(r.FindString(str), `bilibili.com/bangumi/media/md`), MEDIA, err
    case strings.Contains(str, `av`):
        r, err := regexp.Compile(AV_REGEX)
        if err != nil {
            return "", NOTMATCHED, err
        }
        return strings.TrimPrefix(r.FindString(str), `av`), AVVIDEO, err
    case strings.Contains(str, `AV`):
        r, err := regexp.Compile(AV_REGEX)
        if err != nil {
            return "", NOTMATCHED, err
        }
        return strings.TrimPrefix(r.FindString(str), `AV`), AVVIDEO, err
    case (strings.Contains(str, `BV`) || strings.Contains(str, `bv`)):
        r, err := regexp.Compile(BV_REGEX)
        if err != nil {
            return "", NOTMATCHED, err
        }
        return r.FindString(str), BVVIDEO, err
    case strings.Contains(str, `channel/collectiondetail?sid=`):
        r, err := regexp.Compile(COLLECTION_REGEX)
        if err != nil {
            return "", NOTMATCHED, err
        }
        return strings.TrimPrefix(r.FindString(str), `channel/collectiondetail?sid=`), COLLECTION, err
    default:
        return "", NOTMATCHED, nil
    }
}

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
    s, err := fetchEpisode(str)
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
    s, err := fetchSeason(str)
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

func fetchAV(str string) (*videoType, error) {
    var s videoType
    
    str = StringBuilder(`aid=`, str)
    
    err := fetch(StringBuilder(`https://api.bilibili.com/x/web-interface/view?`, str), &s)
    if err != nil {
        return nil, err
    }
    
    if s.Code != 0 {
        return nil, errors.New(s.Message)
    }
    
    return &s, nil
}

func fetchBV(str string) (*videoType, error) {
    var s videoType
    
    str = StringBuilder(`bvid=`, str)
    
    err := fetch(StringBuilder(`https://api.bilibili.com/x/web-interface/view?`, str), &s)
    if err != nil {
        return nil, err
    }
    
    if s.Code != 0 {
        return nil, errors.New(s.Message)
    }
    
    return &s, nil
}

func fetchSeason(str string) (*episodeOrSeasonType, error) {
    var s episodeOrSeasonType
    
    str = StringBuilder(`season_id=`, str)
    
    err := fetch(StringBuilder(`https://api.bilibili.com/pgc/view/web/season?`, str), &s)
    if err != nil {
        return nil, err
    }
    
    if s.Code != 0 {
        return nil, errors.New(s.Message)
    }
    
    return &s, nil
}

func fetchEpisode(str string) (*episodeOrSeasonType, error) {
    var s episodeOrSeasonType
    
    str = StringBuilder(`ep_id=`, str)
    
    err := fetch(StringBuilder(`https://api.bilibili.com/pgc/view/web/season?`, str), &s)
    if err != nil {
        return nil, err
    }
    
    if s.Code != 0 {
        return nil, errors.New(s.Message)
    }
    
    return &s, nil
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

func fetch(URL string, result interface{}) error {
    req, err := http.NewRequest(`GET`, URL, nil)
    if err != nil {
        return err
    }
    
    req.AddCookie(&http.Cookie{Name: `SESSDATA`, Value: Conf.GetString(`bilibili.sessdata`), Path: `/`, Expires: time.Now().AddDate(0, 0, 1)})
    req.AddCookie(&http.Cookie{Name: `CURRENT_FNVAL`, Value: `4048`, Path: `/`, Expires: time.Now().AddDate(0, 0, 1)})
    req.AddCookie(&http.Cookie{Name: `CURRENT_QUALITY`, Value: `120`, Path: `/`, Expires: time.Now().AddDate(0, 0, 1)})
    req.Header.Add(`Referer`, `https://www.bilibili.com`)
    req.Header.Add(`User-Agent`, "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/97.0.4692.99 Safari/537.36")
    
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

func StringBuilder(p ...string) string {
    var (
        b strings.Builder
        c int
    )
    l := len(p)
    for i := 0; i < l; i++ {
        c += len(p[i])
    }
    b.Grow(c)
    for i := 0; i < l; i++ {
        b.WriteString(p[i])
    }
    return b.String()
}

func getBestQualityVideo(qlist []int) int {
    var tmp int
    
    for i := range qlist {
        if qlist[i] > tmp {
            tmp = qlist[i]
        }
    }
    
    return tmp
}

func deleteDolbyAndGetBestQualityVideo(qlist []int) int {
    var tmp int
    
    for i := range qlist {
        if qlist[i] == 126 {
            continue
        }
        
        if qlist[i] > tmp {
            tmp = qlist[i]
        }
    }
    
    return tmp
}

func getQualityString(i int) string {
    switch i {
    case 127:
        return `超高清 8K`
    case 126:
        return `杜比视界`
    case 125:
        return `HDR 真彩色`
    case 120:
        return `4K 超清`
    case 116:
        return `1080P60 高帧率`
    case 112:
        return `1080P+ 高码率`
    case 80:
        return `1080P 高清`
    case 74:
        return `720P60 高帧率`
    case 64:
        return `720P 高清`
    case 32:
        return `480P 清晰`
    case 16:
        return `360P 流畅`
    case 6:
        return `240P 极速`
    default:
        return ``
    }
}

type Details []*bilibili.PartInfo

func (c Details) Len() int {
    return len(c)
}
func (c Details) Swap(i, j int) {
    c[i], c[j] = c[j], c[i]
}
func (c Details) Less(i, j int) bool {
    return c[i].Index < c[j].Index
}

type episodeOrSeasonType struct {
    Code    int    `json:"code"`
    Message string `json:"message,omitempty"`
    Result  struct {
        Areas []struct {
            Id   int    `json:"id,omitempty"`
            Name string `json:"name,omitempty"`
        } `json:"areas,omitempty"`
        Evaluate string `json:"evaluate,omitempty"`
        Cover    string `json:"cover,omitempty"`
        NewEp    struct {
            Desc string `json:"desc,omitempty"`
        } `json:"new_ep,omitempty"`
        Title       string `json:"title,omitempty"`
        ShareUrl    string `json:"share_url,omitempty"`
        SeasonId    int    `json:"season_id,omitempty"`
        SeasonTitle string `json:"season_title,omitempty"`
        Episodes    []struct {
            Bvid      string `json:"bvid,omitempty"`
            Cid       int    `json:"cid,omitempty"`
            Title     string `json:"title,omitempty"`
            LongTitle string `json:"long_title,omitempty"`
            Id        int    `json:"id,omitempty"`
        } `json:"episodes,omitempty"`
    } `json:"result,omitempty"`
}

type mediaType struct {
    Code    int    `json:"code"`
    Message string `json:"message,omitempty"`
    Result  struct {
        Media struct {
            Areas []struct {
                Name string `json:"name,omitempty"`
            } `json:"areas,omitempty"`
            Cover string `json:"cover,omitempty"`
            NewEp struct {
                IndexShow string `json:"index_show,omitempty"`
            } `json:"new_ep,omitempty"`
            ShareUrl string `json:"share_url,omitempty"`
            Title    string `json:"title,omitempty"`
            TypeName string `json:"type_name,omitempty"`
            SeasonId int    `json:"season_id,omitempty"`
        } `json:"media,omitempty"`
    } `json:"result,omitempty"`
}

type videoType struct {
    Code    int    `json:"code"`
    Message string `json:"message,omitempty"`
    Data    struct {
        Bvid     string `json:"bvid,omitempty"`
        Aid      int64  `json:"aid,omitempty"`
        Pic      string `json:"pic,omitempty"`
        Title    string `json:"title,omitempty"`
        Cid      int    `json:"cid,omitempty"`
        Pubdate  int64  `json:"pubdate,omitempty"`
        Ctime    int64  `json:"ctime,omitempty"`
        Desc     string `json:"desc,omitempty"`
        Duration int64  `json:"duration,omitempty"`
        Owner    struct {
            Name string `json:"name,omitempty"`
        } `json:"owner,omitempty"`
        Dynamic string `json:"dynamic,omitempty"`
        Pages   []struct {
            Cid  int    `json:"cid,omitempty"`
            Part string `json:"part,omitempty"`
        } `json:"pages,omitempty"`
        UgcSeason struct {
            Id int `json:"id"`
        } `json:"ugc_season,omitempty"`
    } `json:"data,omitempty"`
}

type VideoDetail struct {
    Code int `json:"code"`
    Data struct {
        Quality       int    `json:"quality"`
        AcceptFormat  string `json:"accept_format"`
        AcceptQuality []int  `json:"accept_quality"`
        Durl          []struct {
            Order     int      `json:"order"`
            Url       string   `json:"url"`
            BackupUrl []string `json:"backup_url"`
        } `json:"durl,omitempty"`
        Dash struct {
            Video []struct {
                Id         int      `json:"id"`
                BaseUrl    string   `json:"baseUrl"`
                Base_Url   string   `json:"base_url"`
                BackupUrl  []string `json:"backupUrl"`
                Backup_Url []string `json:"backup_url"`
                Codecs     string   `json:"codecs"`
            } `json:"video"`
            Audio []struct {
                Id         int      `json:"id"`
                BaseUrl    string   `json:"baseUrl"`
                Base_Url   string   `json:"base_url"`
                BackupUrl  []string `json:"backupUrl"`
                Backup_Url []string `json:"backup_url"`
                Bandwidth  int      `json:"bandwidth"`
            } `json:"audio"`
            Dolby struct {
                Type  int `json:"type"`
                Audio []struct {
                    Id          int      `json:"id"`
                    BaseUrl     string   `json:"base_url"`
                    BackupUrl   []string `json:"backup_url"`
                    Bandwidth   int      `json:"bandwidth"`
                    MimeType    string   `json:"mime_type"`
                    Codecs      string   `json:"codecs"`
                    SegmentBase struct {
                        Initialization string `json:"initialization"`
                        IndexRange     string `json:"index_range"`
                    } `json:"segment_base"`
                    Size int `json:"size"`
                } `json:"audio"`
            } `json:"dolby,omitempty"`
        } `json:"dash,omitempty"`
    } `json:"data"`
}

type MediaDetail struct {
    Code int `json:"code"`
    Data struct {
        Quality       int    `json:"quality"`
        AcceptFormat  string `json:"accept_format"`
        AcceptQuality []int  `json:"accept_quality"`
        Durl          []struct {
            Order     int      `json:"order"`
            Url       string   `json:"url"`
            BackupUrl []string `json:"backup_url"`
        } `json:"durl,omitempty"`
        Dash struct {
            Video []struct {
                Id         int      `json:"id"`
                BaseUrl    string   `json:"baseUrl"`
                Base_Url   string   `json:"base_url"`
                BackupUrl  []string `json:"backupUrl"`
                Backup_Url []string `json:"backup_url"`
                Codecs     string   `json:"codecs"`
            } `json:"video"`
            Audio []struct {
                Id         int      `json:"id"`
                BaseUrl    string   `json:"baseUrl"`
                Base_Url   string   `json:"base_url"`
                BackupUrl  []string `json:"backupUrl"`
                Backup_Url []string `json:"backup_url"`
                Bandwidth  int      `json:"bandwidth"`
            } `json:"audio"`
            Dolby struct {
                Type  int `json:"type"`
                Audio []struct {
                    Id          int      `json:"id"`
                    BaseUrl     string   `json:"base_url"`
                    BackupUrl   []string `json:"backup_url"`
                    Bandwidth   int      `json:"bandwidth"`
                    MimeType    string   `json:"mime_type"`
                    Codecs      string   `json:"codecs"`
                    SegmentBase struct {
                        Initialization string `json:"initialization"`
                        IndexRange     string `json:"index_range"`
                    } `json:"segment_base"`
                    Size int `json:"size"`
                } `json:"audio"`
            } `json:"dolby,omitempty"`
        } `json:"dash,omitempty"`
    } `json:"result"`
}

type collectionType struct {
    Code    int    `json:"code"`
    Message string `json:"message"`
    Ttl     int    `json:"ttl"`
    Data    struct {
        Aids     []int `json:"aids"`
        Archives []struct {
            Aid              int    `json:"aid"`
            Bvid             string `json:"bvid"`
            Ctime            int    `json:"ctime"`
            Duration         int    `json:"duration"`
            InteractiveVideo bool   `json:"interactive_video"`
            Pic              string `json:"pic"`
            Pubdate          int    `json:"pubdate"`
            Stat             struct {
                View int `json:"view"`
            } `json:"stat"`
            State  int    `json:"state"`
            Title  string `json:"title"`
            UgcPay int    `json:"ugc_pay"`
        } `json:"archives"`
        Meta struct {
            Category    int    `json:"category"`
            Cover       string `json:"cover"`
            Description string `json:"description"`
            Mid         int    `json:"mid"`
            Name        string `json:"name"`
            Ptime       int    `json:"ptime"`
            SeasonId    int    `json:"season_id"`
            Total       int    `json:"total"`
        } `json:"meta"`
        Page struct {
            PageNum  int `json:"page_num"`
            PageSize int `json:"page_size"`
            Total    int `json:"total"`
        } `json:"page"`
    } `json:"data"`
}

package main

import (
    "avalon-bilibili/src/bilibili"
    "regexp"
    "strings"
)

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

func checkRegion(text string) (limited bool, region bilibili.Region) {
    switch {
    case strings.Contains(text, `僅限港澳`):
        return true, bilibili.Region_HK
    case strings.Contains(text, `僅限台灣`):
        return true, bilibili.Region_TW
    }
    
    return false, bilibili.Region_CN
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

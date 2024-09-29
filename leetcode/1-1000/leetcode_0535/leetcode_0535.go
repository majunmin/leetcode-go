package leetcode_0535

import (
	"math/rand/v2"
	"strings"
)

const (
	chars      = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	size       = len(chars)
	tinyPrefix = "http://tinyurl.com/"
)

type Codec struct {
	key2Origin map[string]string
	origin2Key map[string]string
}

func Constructor() Codec {
	return Codec{
		key2Origin: make(map[string]string),
		origin2Key: make(map[string]string),
	}
}

// Encodes a URL to a shortened URL.
func (this *Codec) encode(longUrl string) string {
	var urlKey string
	if urlKey, exist := this.origin2Key[longUrl]; exist {
		return tinyPrefix + urlKey
	}
	for { // 这里可以限制一下次数,不要无限循环.
		urlKey = generate(longUrl, 6)
		if _, exist := this.key2Origin[urlKey]; exist {
			continue
		}
		this.key2Origin[urlKey] = longUrl
		this.origin2Key[longUrl] = urlKey
		break
	}
	return urlKey
}

func generate(url string, cnt int) string {
	var sb strings.Builder
	defer sb.Reset()
	for i := 0; i < cnt; i++ {
		item := chars[rand.IntN(size)]
		sb.WriteByte(item)
	}
	return sb.String()
}

// Decodes a shortened URL to its original URL.
func (this *Codec) decode(shortUrl string) string {
	reqUrlKey := shortUrl[strings.LastIndexByte(shortUrl, '/')+1:]
	if originUrl, exist := this.key2Origin[reqUrlKey]; exist {
		return originUrl
	}
	return ""
}

/**
 * Your Codec object will be instantiated and called as such:
 * obj := Constructor();
 * url := obj.encode(longUrl);
 * ans := obj.decode(url);
 */

package live

import (
	"bytes"
	"compress/gzip"
	"danmu-core/core/jsScript"
	"danmu-core/logger"
	"danmu-core/utils"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"time"

	"github.com/spf13/cast"
)

func (d *DouyinLive) fetchTTWID() error {
	res, err := http.Get("https://live.douyin.com/")
	if err != nil {
		return fmt.Errorf("获取直播 URL 失败: %w", err)
	}

	for _, cookie := range res.Cookies() {
		if cookie.Name == "ttwid" {
			d.ttwid = cookie.Value
			return nil
		}
	}
	return fmt.Errorf("未找到 ttwid cookie")
}

// fetchRoomID 获取 roomID
func (d *DouyinLive) fetchRoomID() {
	if d.roomId != "" {
		return
	}

	ttwid := &http.Cookie{
		Name:  "ttwid",
		Value: "ttwid=" + d.ttwid + "&msToken=" + utils.GenerateMsToken(107),
	}
	acNonce := &http.Cookie{
		Name:  "__ac_nonce",
		Value: "0123407cc00a9e438deb4",
	}
	res, err := d.c.R().SetCookies(ttwid, acNonce).Get(d.liveurl)
	if err != nil {
		logger.Warn().Str("liveurl", d.liveurl).Err(err).Msg("获取房间 ID 失败")
		return
	}

	var roomIDRegexp = regexp.MustCompile(`roomId\\":\\"(\d+)\\"`)
	var pushIDRegexp = regexp.MustCompile(`user_unique_id\\":\\"(\d+)\\"`)
	d.roomId = extractMatch(roomIDRegexp, res.String())
	d.pushid = extractMatch(pushIDRegexp, res.String())
}

// extractMatch 从字符串中提取正则表达式匹配的内容
func extractMatch(re *regexp.Regexp, s string) string {
	match := re.FindStringSubmatch(s)
	if len(match) > 1 {
		return match[1]
	}
	return ""
}

// GzipUnzipReset 使用 Reset 方法解压 gzip 数据
// GzipUnzipReset 解压 gzip 数据
func (d *DouyinLive) GzipUnzipReset(compressedData []byte) ([]byte, error) {
	// 创建一个新的 bytes.Buffer 来读取压缩数据
	buffer := bytes.NewBuffer(compressedData)

	var err error
	// 重用或创建 gzip reader
	if d.gzip != nil {
		err = d.gzip.Reset(buffer)
		if err != nil {
			d.gzip.Close()
			d.gzip = nil
			return nil, err
		}
	} else {
		d.gzip, err = gzip.NewReader(buffer)
		if err != nil {
			return nil, err
		}
	}
	defer d.gzip.Close()

	// 读取解压后的数据
	uncompressedBuffer := &bytes.Buffer{}
	_, err = io.Copy(uncompressedBuffer, d.gzip)
	if err != nil {
		return nil, err
	}

	return uncompressedBuffer.Bytes(), nil
}

// StitchUrl 构建 WebSocket 连接的 URL
func (d *DouyinLive) StitchUrl() string {
	smap := utils.NewOrderedMap(d.roomId, d.pushid)
	signaturemd5 := utils.GetxMSStub(smap)
	signature := jsScript.ExecuteJS(signaturemd5)
	browserInfo := strings.Split(d.userAgent, "Mozilla")[1]
	parsedURL := strings.Replace(browserInfo[1:], " ", "%20", -1)
	fetchTime := time.Now().UnixNano() / int64(time.Millisecond)
	return "wss://webcast5-ws-web-lf.douyin.com/webcast/im/push/v2/?app_name=douyin_web&version_code=180800&" +
		"webcast_sdk_version=1.0.14-beta.0&update_version_code=1.0.14-beta.0&compress=gzip&device_platform" +
		"=web&cookie_enabled=true&screen_width=1920&screen_height=1080&browser_language=zh-CN&browser_platform=Win32&" +
		"browser_name=Mozilla&browser_version=" + parsedURL + "&browser_online=true" +
		"&tz_name=Asia/Shanghai&cursor=d-1_u-1_fh-7383731312643626035_t-1719159695790_r-1&internal_ext" +
		"=internal_src:dim|wss_push_room_id:" + d.roomId + "|wss_push_did:" + d.pushid + "|first_req_ms:" + cast.ToString(fetchTime) + "|fetch_time:" + cast.ToString(fetchTime) + "|seq:1|wss_info:0-" + cast.ToString(fetchTime) + "-0-0|" +
		"wrds_v:7382620942951772256&host=https://live.douyin.com&aid=6383&live_id=1&did_rule=3" +
		"&endpoint=live_pc&support_wrds=1&user_unique_id=" + d.pushid + "&im_path=/webcast/im/fetch/" +
		"&identity=audience&need_persist_msg_count=15&insert_task_id=&live_reason=&room_id=" + d.roomId + "&heartbeatDuration=0&signature=" + signature
}

func (d *DouyinLive) distinct(messageLogId uint64) bool {
	if cacheTime, exists := d.distinctCache.Get(messageLogId); exists {
		if time.Since(cacheTime.(time.Time)) > 10*time.Second {
			d.distinctCache.Remove(messageLogId)
		} else {
			logger.Debug().Uint64("id", messageLogId).Str("liveurl", d.liveurl)
			return false
		}
	}
	d.distinctCache.Add(messageLogId, time.Now())
	return true
}

func (d *DouyinLive) getWebRoomInfo(webRid string) (map[string]interface{}, error) {
	// 构建请求 URL
	targetUrl := d.buildRequestUrl(fmt.Sprintf("https://live.douyin.com/webcast/room/web/enter/?web_rid=%s", webRid))

	// 使用 GET 请求获取数据
	resp, err := d.c.R().Get(targetUrl)
	defer func() {
		if resp.Body != nil {
			resp.Body.Close()
		}
	}()
	if err != nil {
		return nil, err
	}
	var result map[string]interface{}
	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (d *DouyinLive) getRoomInfo(secUid string, roomId string) (map[string]interface{}, error) {
	if secUid == "" {
		return nil, fmt.Errorf("sec_uid is nil")
	}
	params := url.Values{}
	params.Add("type_id", "0")
	params.Add("live_id", "1")
	params.Add("version_code", "99.99.99")
	params.Add("app_id", "1128")
	if roomId == "" {
		params.Add("room_id", "2") // 默认值
	} else {
		params.Add("room_id", roomId)
	}
	params.Add("sec_user_id", secUid)
	baseURL := "https://webcast.amemv.com/webcast/room/reflow/info/"
	requestURL := baseURL + "?" + params.Encode()
	resp, err := d.c.R().Get(requestURL)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Response.Body)
	// 解析返回的 JSON 数据
	if err != nil {
		return nil, err
	}
	// 定义一个空的 map 来存储 JSON 数据
	var result map[string]interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (d *DouyinLive) buildRequestUrl(urlStr string) string {
	// 解析 URL
	parsedUrl, err := url.Parse(urlStr)
	if err != nil {
		logger.Error().Str("liveurl", d.liveurl).Err(err).Msg("无法解析Url:" + urlStr)
	}

	// 获取现有的查询参数
	queryParams := parsedUrl.Query()

	// 添加/更新参数
	queryParams.Set("aid", "6383")
	queryParams.Set("device_platform", "web")
	queryParams.Set("browser_language", "zh-CN")
	queryParams.Set("enter_from", "web_live")
	queryParams.Set("browser_platform", "Win32")

	// 获取 Douyin 用户代理信息
	uaParts := strings.Split(d.userAgent, "/")
	if len(uaParts) > 1 {
		queryParams.Set("browser_name", uaParts[0])
		queryParams.Set("browser_version", uaParts[1])
	}

	// 将新的查询参数编码到 URL 中
	parsedUrl.RawQuery = queryParams.Encode()

	// 返回新的 URL
	return parsedUrl.String()
}

package live

import (
	"compress/gzip"
	"context"
	"danmu-core/generated/douyin"
	"github.com/gorilla/websocket"
	lru "github.com/hashicorp/golang-lru"
	"github.com/imroc/req/v3"
	"net/http"
	"sync"
	"sync/atomic"
)

const (
	WebcastChatMessage        = "WebcastChatMessage"
	WebcastGiftMessage        = "WebcastGiftMessage"
	WebcastLikeMessage        = "WebcastLikeMessage"
	WebcastMemberMessage      = "WebcastMemberMessage"
	WebcastSocialMessage      = "WebcastSocialMessage"
	WebcastRoomUserSeqMessage = "WebcastRoomUserSeqMessage"
	WebcastFansclubMessage    = "WebcastFansclubMessage"
	WebcastControlMessage     = "WebcastControlMessage"
	WebcastEmojiChatMessage   = "WebcastEmojiChatMessage"
	WebcastRoomStatsMessage   = "WebcastRoomStatsMessage"
	WebcastRoomMessage        = "WebcastRoomMessage"
	WebcastRoomRankMessage    = "WebcastRoomRankMessage"

	Default = "Default"
)

type EventHandler func(eventData *douyin.Message)
type DouyinLive struct {
	ttwid         string
	liveid        string
	liveurl       string
	userAgent     string
	c             *req.Client
	eventHandlers []EventHandler
	headers       http.Header
	gzip          *gzip.Reader
	Conn          *websocket.Conn
	wssurl        string
	pushid        string
	recvMsgChan   chan *douyin.Response
	ctx           context.Context
	cancelFunc    context.CancelFunc
	enable        atomic.Bool
	distinctCache *lru.Cache
	mu            sync.Mutex
	connMu        sync.Mutex
	isLive        atomic.Bool

	secUid string
	roomId string
	webRid string
}

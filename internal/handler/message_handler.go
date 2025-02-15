package handler

import (
	"danmu-core/core/live"
	"danmu-core/generated/douyin"
	"danmu-core/internal/model"
	"danmu-core/logger"
	"danmu-core/utils"
	"fmt"

	lru "github.com/hashicorp/golang-lru"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type DouyinHandler struct {
	cache         *lru.Cache
	roomDisplayId string
	roomName      string
	liveUrl       string
}

func NewDouyinHandler(conf *model.LiveConf) (*DouyinHandler, error) {
	cache, err := lru.New(1000)
	if err != nil {
		return nil, fmt.Errorf("DouyinHandler Init Cache failure, err:%v", err)
	}
	return &DouyinHandler{
		cache:         cache,
		roomDisplayId: conf.RoomDisplayID,
		roomName:      conf.Name,
		liveUrl:       conf.URL,
	}, nil
}

func (h *DouyinHandler) HandleDouyinMessage(eventData *douyin.Message) {
	msg, err := utils.MatchMethod(eventData.Method)
	if err != nil || msg == nil {
		return
	}
	if _, exists := h.cache.Get(eventData.MsgId); exists {
		return
	}
	if proto.Unmarshal(eventData.Payload, msg) != nil {
		logger.Warn().Str("liveid", h.roomDisplayId).Err(err).Msg("反序列化失败")
		return
	}
	if err := h.saveToDB(msg, eventData.Method, eventData.MsgId); err != nil {
		h.cache.Add(eventData.MsgId, true)
	}
}

func (h *DouyinHandler) saveToDB(msg protoreflect.ProtoMessage, method string, id int64) error {
	var common *model.CommonMessage
	switch method {
	case live.WebcastGiftMessage:
		m := msg.(*douyin.GiftMessage)
		// 先处理用户信息
		user := model.NewUser(m.User)
		if err := user.CheckAndInsert(); err != nil {
			logger.Warn().Str("liveid", h.roomDisplayId).Err(err).Msg("Failed to process user")
			// 不返回错误，继续处理礼物消息
		}

		if m.RepeatEnd == 1 {
			return nil
		}
		common = &model.CommonMessage{
			MessageType:   method,
			UserName:      m.User.NickName,
			UserID:        m.User.Id,
			UserDisplayId: m.User.DisplayId,
			RoomID:        m.Common.RoomId,
			Content:       m.Common.Describe,
			Timestamp:     m.Common.CreateTime,
			RoomName:      h.roomName,
			RoomDisplayId: h.roomDisplayId,
		}
		giftMessage := model.NewGiftMessage(m)
		giftMessage.ID = id
		giftMessage.RoomDisplayId = h.roomDisplayId
		giftMessage.RoomName = h.roomName
		if err := giftMessage.Insert(); err != nil {
			logger.Warn().Str("liveid", h.roomDisplayId).Err(err).
				Msgf("Failed to insert gift message: %v", m)
			return err
		}
	case live.WebcastChatMessage:
		m := msg.(*douyin.ChatMessage)
		common = &model.CommonMessage{
			MessageType:   method,
			UserName:      m.User.NickName,
			UserID:        m.User.Id,
			UserDisplayId: m.User.DisplayId,
			RoomID:        m.Common.RoomId,
			RoomDisplayId: h.roomDisplayId,
			RoomName:      h.roomName,
			Content:       fmt.Sprintf("[%v]: %v", m.User.NickName, m.Content),
			Timestamp:     int64(m.EventTime),
		}
		/*	case live.WebcastMemberMessage:
				m := msg.(*douyin.MemberMessage)
				common = &model.CommonMessage{
					MessageType:   method,
					UserName:      m.User.NickName,
					UserID:        m.User.ShortId,
					UserDisplayId: m.User.DisplayId,
					RoomID:        m.Common.RoomId,
					Content:       fmt.Sprintf("%v 来了, 人数 %v", m.User.NickName, m.MemberCount),
					Timestamp:     time.Now().Unix(),
				}
			case live.WebcastSocialMessage:
				m := msg.(*douyin.SocialMessage)
				common = &model.CommonMessage{
					MessageType:   method,
					UserName:      m.User.NickName,
					UserID:        m.User.ShortId,
					UserDisplayId: m.User.DisplayId,
					RoomID:        m.Common.RoomId,
					Content:       fmt.Sprintf("%v 关注了，Follow Count: %v", m.User.NickName, m.FollowCount),
					Timestamp:     time.Now().Unix(),
				}*/

		/*	case live.WebcastLikeMessage:
			m := msg.(*douyin.LikeMessage)
			common = &model.CommonMessage{
				MessageType:   method,
				UserName:      m.User.NickName,
				UserID:        m.User.ShortId,
				UserDisplayId: m.User.DisplayId,
				RoomID:        m.Common.RoomId,
				Content:       fmt.Sprintf("%v 为主播点赞， Total: %v", m.User.NickName, m.Total),
				Timestamp:     time.Now().Unix(),
			}*/
	default:
		return nil
	}
	if common != nil {
		common.ID = uint64(id)
		common.Timestamp = utils.NormalizeTimestamp(common.Timestamp)
		if err := common.Insert(); err != nil {
			logger.Warn().Str("liveid", h.roomDisplayId).Err(err).
				Msgf("Failed to insert common message: %v", common)
			return err
		}
		//	logger.Debug().Uint64("commonid", common.ID).Msg("save to db")
	}
	return nil
}

package live

// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    welcome, err := UnmarshalWelcome(bytes)
//    bytes, err = welcome.Marshal()

import "encoding/json"

func UnmarshalRoomInfo(data []byte) (Welcome, error) {
	var r Welcome
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Welcome) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Welcome struct {
	Data struct {
		Data []struct {
			IDStr        string `json:"id_str"`
			Status       int    `json:"status"`
			StatusStr    string `json:"status_str"`
			Title        string `json:"title"`
			UserCountStr string `json:"user_count_str"`
			Cover        struct {
				URLList []string `json:"url_list"`
			} `json:"cover"`
			StreamURL struct {
				FlvPullURL struct {
					FULLHD1 string `json:"FULL_HD1"`
					HD1     string `json:"HD1"`
					SD1     string `json:"SD1"`
					SD2     string `json:"SD2"`
				} `json:"flv_pull_url"`
				DefaultResolution string `json:"default_resolution"`
				HlsPullURLMap     struct {
					FULLHD1 string `json:"FULL_HD1"`
					HD1     string `json:"HD1"`
					SD1     string `json:"SD1"`
					SD2     string `json:"SD2"`
				} `json:"hls_pull_url_map"`
				HlsPullURL        string `json:"hls_pull_url"`
				StreamOrientation int    `json:"stream_orientation"`
				LiveCoreSdkData   struct {
					PullData struct {
						Options struct {
							DefaultQuality struct {
								Name              string `json:"name"`
								SdkKey            string `json:"sdk_key"`
								VCodec            string `json:"v_codec"`
								Resolution        string `json:"resolution"`
								Level             int    `json:"level"`
								VBitRate          int    `json:"v_bit_rate"`
								AdditionalContent string `json:"additional_content"`
								Fps               int    `json:"fps"`
								Disable           int    `json:"disable"`
							} `json:"default_quality"`
							Qualities []struct {
								Name              string `json:"name"`
								SdkKey            string `json:"sdk_key"`
								VCodec            string `json:"v_codec"`
								Resolution        string `json:"resolution"`
								Level             int    `json:"level"`
								VBitRate          int    `json:"v_bit_rate"`
								AdditionalContent string `json:"additional_content"`
								Fps               int    `json:"fps"`
								Disable           int    `json:"disable"`
							} `json:"qualities"`
						} `json:"options"`
						StreamData string `json:"stream_data"`
					} `json:"pull_data"`
				} `json:"live_core_sdk_data"`
				Extra struct {
					Height                  int  `json:"height"`
					Width                   int  `json:"width"`
					Fps                     int  `json:"fps"`
					MaxBitrate              int  `json:"max_bitrate"`
					MinBitrate              int  `json:"min_bitrate"`
					DefaultBitrate          int  `json:"default_bitrate"`
					BitrateAdaptStrategy    int  `json:"bitrate_adapt_strategy"`
					AnchorInteractProfile   int  `json:"anchor_interact_profile"`
					AudienceInteractProfile int  `json:"audience_interact_profile"`
					HardwareEncode          bool `json:"hardware_encode"`
					VideoProfile            int  `json:"video_profile"`
					H265Enable              bool `json:"h265_enable"`
					GopSec                  int  `json:"gop_sec"`
					BframeEnable            bool `json:"bframe_enable"`
					Roi                     bool `json:"roi"`
					SwRoi                   bool `json:"sw_roi"`
					Bytevc1Enable           bool `json:"bytevc1_enable"`
				} `json:"extra"`
				PullDatas struct {
				} `json:"pull_datas"`
			} `json:"stream_url"`
			MosaicStatus    int      `json:"mosaic_status"`
			MosaicStatusStr string   `json:"mosaic_status_str"`
			AdminUserIds    []int64  `json:"admin_user_ids"`
			AdminUserIdsStr []string `json:"admin_user_ids_str"`
			Owner           struct {
				IDStr       string `json:"id_str"`
				SecUID      string `json:"sec_uid"`
				Nickname    string `json:"nickname"`
				AvatarThumb struct {
					URLList []string `json:"url_list"`
				} `json:"avatar_thumb"`
				FollowInfo struct {
					FollowStatus    int    `json:"follow_status"`
					FollowStatusStr string `json:"follow_status_str"`
				} `json:"follow_info"`
				Subscribe struct {
					IsMember     bool `json:"is_member"`
					Level        int  `json:"level"`
					IdentityType int  `json:"identity_type"`
					BuyType      int  `json:"buy_type"`
					Open         int  `json:"open"`
				} `json:"subscribe"`
				ForeignUser int    `json:"foreign_user"`
				OpenIDStr   string `json:"open_id_str"`
			} `json:"owner"`
			RoomAuth struct {
				Chat                      bool `json:"Chat"`
				Danmaku                   bool `json:"Danmaku"`
				Gift                      bool `json:"Gift"`
				LuckMoney                 bool `json:"LuckMoney"`
				Digg                      bool `json:"Digg"`
				RoomContributor           bool `json:"RoomContributor"`
				Props                     bool `json:"Props"`
				UserCard                  bool `json:"UserCard"`
				POI                       bool `json:"POI"`
				MoreAnchor                int  `json:"MoreAnchor"`
				Banner                    int  `json:"Banner"`
				Share                     int  `json:"Share"`
				UserCorner                int  `json:"UserCorner"`
				Landscape                 int  `json:"Landscape"`
				LandscapeChat             int  `json:"LandscapeChat"`
				PublicScreen              int  `json:"PublicScreen"`
				GiftAnchorMt              int  `json:"GiftAnchorMt"`
				RecordScreen              int  `json:"RecordScreen"`
				DonationSticker           int  `json:"DonationSticker"`
				HourRank                  int  `json:"HourRank"`
				CommerceCard              int  `json:"CommerceCard"`
				AudioChat                 int  `json:"AudioChat"`
				DanmakuDefault            int  `json:"DanmakuDefault"`
				KtvOrderSong              int  `json:"KtvOrderSong"`
				SelectionAlbum            int  `json:"SelectionAlbum"`
				Like                      int  `json:"Like"`
				MultiplierPlayback        int  `json:"MultiplierPlayback"`
				DownloadVideo             int  `json:"DownloadVideo"`
				Collect                   int  `json:"Collect"`
				TimedShutdown             int  `json:"TimedShutdown"`
				Seek                      int  `json:"Seek"`
				Denounce                  int  `json:"Denounce"`
				Dislike                   int  `json:"Dislike"`
				OnlyTa                    int  `json:"OnlyTa"`
				CastScreen                int  `json:"CastScreen"`
				CommentWall               int  `json:"CommentWall"`
				BulletStyle               int  `json:"BulletStyle"`
				ShowGamePlugin            int  `json:"ShowGamePlugin"`
				VSGift                    int  `json:"VSGift"`
				VSTopic                   int  `json:"VSTopic"`
				VSRank                    int  `json:"VSRank"`
				AdminCommentWall          int  `json:"AdminCommentWall"`
				CommerceComponent         int  `json:"CommerceComponent"`
				DouPlus                   int  `json:"DouPlus"`
				GamePointsPlaying         int  `json:"GamePointsPlaying"`
				Poster                    int  `json:"Poster"`
				Highlights                int  `json:"Highlights"`
				TypingCommentState        int  `json:"TypingCommentState"`
				StrokeUpDownGuide         int  `json:"StrokeUpDownGuide"`
				UpRightStatsFloatingLayer int  `json:"UpRightStatsFloatingLayer"`
				CastScreenExplicit        int  `json:"CastScreenExplicit"`
				Selection                 int  `json:"Selection"`
				IndustryService           int  `json:"IndustryService"`
				VerticalRank              int  `json:"VerticalRank"`
				EnterEffects              int  `json:"EnterEffects"`
				FansClub                  int  `json:"FansClub"`
				EmojiOutside              int  `json:"EmojiOutside"`
				CanSellTicket             int  `json:"CanSellTicket"`
				DouPlusPopularityGem      int  `json:"DouPlusPopularityGem"`
				MissionCenter             int  `json:"MissionCenter"`
				ExpandScreen              int  `json:"ExpandScreen"`
				FansGroup                 int  `json:"FansGroup"`
				Topic                     int  `json:"Topic"`
				AnchorMission             int  `json:"AnchorMission"`
				Teleprompter              int  `json:"Teleprompter"`
				LongTouch                 int  `json:"LongTouch"`
				FirstFeedHistChat         int  `json:"FirstFeedHistChat"`
				MoreHistChat              int  `json:"MoreHistChat"`
				TaskBanner                int  `json:"TaskBanner"`
				SpecialStyle              struct {
					Chat struct {
						UnableStyle             int    `json:"UnableStyle"`
						Content                 string `json:"Content"`
						OffType                 int    `json:"OffType"`
						AnchorSwitchForPaidLive int    `json:"AnchorSwitchForPaidLive"`
						ContentForPaidLive      string `json:"ContentForPaidLive"`
					} `json:"Chat"`
					Like struct {
						UnableStyle             int    `json:"UnableStyle"`
						Content                 string `json:"Content"`
						OffType                 int    `json:"OffType"`
						AnchorSwitchForPaidLive int    `json:"AnchorSwitchForPaidLive"`
						ContentForPaidLive      string `json:"ContentForPaidLive"`
					} `json:"Like"`
				} `json:"SpecialStyle"`
				FixedChat             int `json:"FixedChat"`
				QuizGamePointsPlaying int `json:"QuizGamePointsPlaying"`
			} `json:"room_auth"`
			LiveRoomMode int `json:"live_room_mode"`
			Stats        struct {
				TotalUserDesp string `json:"total_user_desp"`
				LikeCount     int    `json:"like_count"`
				TotalUserStr  string `json:"total_user_str"`
				UserCountStr  string `json:"user_count_str"`
			} `json:"stats"`
			HasCommerceGoods bool `json:"has_commerce_goods"`
			LinkerMap        struct {
			} `json:"linker_map"`
			LinkerDetail struct {
				LinkerPlayModes             []interface{} `json:"linker_play_modes"`
				BigPartyLayoutConfigVersion int           `json:"big_party_layout_config_version"`
				AcceptAudiencePreApply      bool          `json:"accept_audience_pre_apply"`
				LinkerUILayout              int           `json:"linker_ui_layout"`
				EnableAudienceLinkmic       int           `json:"enable_audience_linkmic"`
				FunctionType                string        `json:"function_type"`
				LinkerMapStr                struct {
				} `json:"linker_map_str"`
				KtvLyricMode             string `json:"ktv_lyric_mode"`
				InitSource               string `json:"init_source"`
				ForbidApplyFromOther     bool   `json:"forbid_apply_from_other"`
				KtvExhibitMode           int    `json:"ktv_exhibit_mode"`
				EnlargeGuestTurnOnSource int    `json:"enlarge_guest_turn_on_source"`
				PlaymodeDetail           struct {
				} `json:"playmode_detail"`
				ClientUIInfo string        `json:"client_ui_info"`
				ManualOpenUI int           `json:"manual_open_ui"`
				FeatureList  []interface{} `json:"feature_list"`
			} `json:"linker_detail"`
			RoomViewStats struct {
				IsHidden            bool   `json:"is_hidden"`
				DisplayShort        string `json:"display_short"`
				DisplayMiddle       string `json:"display_middle"`
				DisplayLong         string `json:"display_long"`
				DisplayValue        int    `json:"display_value"`
				DisplayVersion      int    `json:"display_version"`
				Incremental         bool   `json:"incremental"`
				DisplayType         int    `json:"display_type"`
				DisplayShortAnchor  string `json:"display_short_anchor"`
				DisplayMiddleAnchor string `json:"display_middle_anchor"`
				DisplayLongAnchor   string `json:"display_long_anchor"`
			} `json:"room_view_stats"`
			SceneTypeInfo struct {
				IsUnionLiveRoom              bool `json:"is_union_live_room"`
				IsLife                       bool `json:"is_life"`
				IsProtectedRoom              int  `json:"is_protected_room"`
				IsLastedGoodsRoom            int  `json:"is_lasted_goods_room"`
				IsDesireRoom                 int  `json:"is_desire_room"`
				CommentaryType               bool `json:"commentary_type"`
				IsSubOrientationVerticalRoom int  `json:"is_sub_orientation_vertical_room"`
			} `json:"scene_type_info"`
			ToolbarData struct {
				EntranceList []struct {
					GroupID       int    `json:"group_id"`
					ComponentType int    `json:"component_type"`
					OpType        int    `json:"op_type"`
					Text          string `json:"text"`
					SchemaURL     string `json:"schema_url"`
					ShowType      int    `json:"show_type"`
					DataStatus    int    `json:"data_status"`
					Extra         string `json:"extra"`
					Icon          struct {
						URLList         []string      `json:"url_list"`
						URI             string        `json:"uri"`
						Height          int           `json:"height"`
						Width           int           `json:"width"`
						AvgColor        string        `json:"avg_color"`
						ImageType       int           `json:"image_type"`
						OpenWebURL      string        `json:"open_web_url"`
						IsAnimated      bool          `json:"is_animated"`
						FlexSettingList []interface{} `json:"flex_setting_list"`
						TextSettingList []interface{} `json:"text_setting_list"`
					} `json:"icon,omitempty"`
				} `json:"entrance_list"`
				MorePanel []struct {
					GroupID       int    `json:"group_id"`
					ComponentType int    `json:"component_type"`
					OpType        int    `json:"op_type"`
					Text          string `json:"text"`
					SchemaURL     string `json:"schema_url"`
					ShowType      int    `json:"show_type"`
					DataStatus    int    `json:"data_status"`
					Extra         string `json:"extra"`
				} `json:"more_panel"`
				MaxEntranceCnt   int           `json:"max_entrance_cnt"`
				LandscapeUpRight []interface{} `json:"landscape_up_right"`
				SkinResource     struct {
				} `json:"skin_resource"`
				MaxEntranceCntLandscape int `json:"max_entrance_cnt_landscape"`
				Permutation             struct {
					General struct {
						GroupPriority     []int `json:"GroupPriority"`
						ComponentSequence []int `json:"ComponentSequence"`
					} `json:"general"`
					OnDemandComponentList []interface{} `json:"on_demand_component_list"`
				} `json:"permutation"`
				ExtraInfo struct {
					GamePromotionCoexist int `json:"game_promotion_coexist"`
				} `json:"extra_info"`
			} `json:"toolbar_data"`
			EcomData struct {
				RedsShowInfos []interface{} `json:"reds_show_infos"`
				InstantType   int           `json:"instant_type"`
			} `json:"ecom_data"`
			RoomCart struct {
				ContainCart bool   `json:"contain_cart"`
				Total       int    `json:"total"`
				FlashTotal  int    `json:"flash_total"`
				CartIcon    string `json:"cart_icon"`
				ShowCart    int    `json:"show_cart"`
			} `json:"room_cart"`
			AnchorABMap struct {
				AbAdminCommentOnWall              string `json:"ab_admin_comment_on_wall"`
				AbFriendChat                      string `json:"ab_friend_chat"`
				AdminPrivilegeRefine              string `json:"admin_privilege_refine"`
				AllowSharedToFans                 string `json:"allow_shared_to_fans"`
				AudienceLinkmicContinue           string `json:"audience_linkmic_continue"`
				AudioDoubleEnlargeEnable          string `json:"audio_double_enlarge_enable"`
				AudioRoomSubtitleOpt              string `json:"audio_room_subtitle_opt"`
				BattleMatchRebuildAnchor          string `json:"battle_match_rebuild_anchor"`
				BigPartyEnableOpenCamera          string `json:"big_party_enable_open_camera"`
				ChatIntercommunicateMultiAnchor   string `json:"chat_intercommunicate_multi_anchor"`
				ChatIntercommunicatePk            string `json:"chat_intercommunicate_pk"`
				DoubleEnlargeEnable               string `json:"double_enlarge_enable"`
				EcomRoomDisableGift               string `json:"ecom_room_disable_gift"`
				EnableEnterBySharing              string `json:"enable_enter_by_sharing"`
				EnableLinkGuestEnter              string `json:"enable_link_guest_enter"`
				EnterMessageTipRelation           string `json:"enter_message_tip_relation"`
				EnterSourceMark                   string `json:"enter_source_mark"`
				FrequentlyChatAbValue             string `json:"frequently_chat_ab_value"`
				FriendRoomAudioTuning             string `json:"friend_room_audio_tuning"`
				FriendRoomSupportNsMode           string `json:"friend_room_support_ns_mode"`
				FriendShareVideoFeatureType       string `json:"friend_share_video_feature_type"`
				GameLinkEntrance                  string `json:"game_link_entrance"`
				GiftHideTip                       string `json:"gift_hide_tip"`
				GuestBattleCrownUpgrade           string `json:"guest_battle_crown_upgrade"`
				GuestBattleExpand                 string `json:"guest_battle_expand"`
				GuestBattleScoreExpand            string `json:"guest_battle_score_expand"`
				GuestBattleUpgrade                string `json:"guest_battle_upgrade"`
				InteractAnchorGuide               string `json:"interact_anchor_guide"`
				KtvAnchorEnableAddAll             string `json:"ktv_anchor_enable_add_all"`
				KtvAutoMuteSelf                   string `json:"ktv_auto_mute_self"`
				KtvChallengeMinusGift             string `json:"ktv_challenge_minus_gift"`
				KtvComponentNewMidi               string `json:"ktv_component_new_midi"`
				KtvEnableAvatar                   string `json:"ktv_enable_avatar"`
				KtvEnableOpenCamera               string `json:"ktv_enable_open_camera"`
				KtvFragmentSong                   string `json:"ktv_fragment_song"`
				KtvGrabGuideSong                  string `json:"ktv_grab_guide_song"`
				KtvGuideSongSwitch                string `json:"ktv_guide_song_switch"`
				KtvKickWhenLinkerFull             string `json:"ktv_kick_when_linker_full"`
				KtvMcHostShowTag                  string `json:"ktv_mc_host_show_tag"`
				KtvNewChallenge                   string `json:"ktv_new_challenge"`
				KtvRoomAtmosphere                 string `json:"ktv_room_atmosphere"`
				KtvSingingHotRank                 string `json:"ktv_singing_hot_rank"`
				KtvVideoStreamOptimize            string `json:"ktv_video_stream_optimize"`
				KtvWantListenEnable               string `json:"ktv_want_listen_enable"`
				LinkmicMultiChorus                string `json:"linkmic_multi_chorus"`
				LinkmicOrderSingSearchFingerprint string `json:"linkmic_order_sing_search_fingerprint"`
				LinkmicOrderSingUpgrade           string `json:"linkmic_order_sing_upgrade"`
				LinkmicStarwish                   string `json:"linkmic_starwish"`
				LiveAnchorEnableChorus            string `json:"live_anchor_enable_chorus"`
				LiveAnchorEnableCustomPosition    string `json:"live_anchor_enable_custom_position"`
				LiveAnchorHitNewAudienceLinkmic   string `json:"live_anchor_hit_new_audience_linkmic"`
				LiveAnchorHitPositionOpt          string `json:"live_anchor_hit_position_opt"`
				LiveAnchorHitVideoBidPaid         string `json:"live_anchor_hit_video_bid_paid"`
				LiveAnchorHitVideoTeamfight       string `json:"live_anchor_hit_video_teamfight"`
				LiveAnswerOnWall                  string `json:"live_answer_on_wall"`
				LiveAudienceLinkmicPreApplyV2     string `json:"live_audience_linkmic_pre_apply_v2"`
				LiveDouPlusEnter                  string `json:"live_dou_plus_enter"`
				LiveKtvEnableBeat                 string `json:"live_ktv_enable_beat"`
				LiveKtvGroup                      string `json:"live_ktv_group"`
				LiveKtvShowSingerIcon             string `json:"live_ktv_show_singer_icon"`
				LiveKtvSingingChallenge           string `json:"live_ktv_singing_challenge"`
				LiveLinkmicBattleOptimize         string `json:"live_linkmic_battle_optimize"`
				LiveLinkmicKtvAnchorLyricMode     string `json:"live_linkmic_ktv_anchor_lyric_mode"`
				LiveLinkmicOrderSingMicroOpt      string `json:"live_linkmic_order_sing_micro_opt"`
				LiveLinkmicOrderSingV3            string `json:"live_linkmic_order_sing_v3"`
				LivePcHelperNewLayout             string `json:"live_pc_helper_new_layout"`
				LiveRoomManageStyle               string `json:"live_room_manage_style"`
				LiveTeamFightFlexible             string `json:"live_team_fight_flexible"`
				LiveVideoEnableCPosition          string `json:"live_video_enable_c_position"`
				LiveVideoEnableSelfDiscipline     string `json:"live_video_enable_self_discipline"`
				LiveVideoHostIdentityEnable       string `json:"live_video_host_identity_enable"`
				LiveVideoShare                    string `json:"live_video_share"`
				LonelyRoomEnterMsgUnfold          string `json:"lonely_room_enter_msg_unfold"`
				MarkUser                          string `json:"mark_user"`
				MergeKtvModeEnable                string `json:"merge_ktv_mode_enable"`
				MergeKtvOptimizeEnable            string `json:"merge_ktv_optimize_enable"`
				OptAudienceLinkmic                string `json:"opt_audience_linkmic"`
				OptPaidLinkFeatureSwitch          string `json:"opt_paid_link_feature_switch"`
				OptranPaidLinkmic                 string `json:"optran_paid_linkmic"`
				OrderSingMv                       string `json:"order_sing_mv"`
				PlayModeOpt24                     string `json:"play_mode_opt_24"`
				PsUseNewPanel                     string `json:"ps_use_new_panel"`
				RadioPrepareApply                 string `json:"radio_prepare_apply"`
				RoomDoubleLike                    string `json:"room_double_like"`
				SelfDisciplineV2                  string `json:"self_discipline_v2"`
				SelfDisciplineV3                  string `json:"self_discipline_v3"`
				SocialShareVideoAdjustVolume      string `json:"social_share_video_adjust_volume"`
				SupportMultipleAddPrice           string `json:"support_multiple_add_price"`
				ThemedCompetitionV2               string `json:"themed_competition_v2"`
				TrafficStrategy                   string `json:"traffic_strategy"`
				VideoEqual1V8FixSwitch            string `json:"video_equal_1v8fix_switch"`
				VideoKtvChallenge                 string `json:"video_ktv_challenge"`
				VideoTalkEnableAvatar             string `json:"video_talk_enable_avatar"`
			} `json:"AnchorABMap"`
			LikeCount      int    `json:"like_count"`
			OwnerUserIDStr string `json:"owner_user_id_str"`
			PaidLiveData   struct {
				PaidType           int  `json:"paid_type"`
				ViewRight          int  `json:"view_right"`
				Duration           int  `json:"duration"`
				Delivery           int  `json:"delivery"`
				NeedDeliveryNotice bool `json:"need_delivery_notice"`
				AnchorRight        int  `json:"anchor_right"`
				PayAbType          int  `json:"pay_ab_type"`
				PrivilegeInfo      struct {
				} `json:"privilege_info"`
				PrivilegeInfoMap struct {
				} `json:"privilege_info_map"`
				MaxPreviewDuration int `json:"max_preview_duration"`
			} `json:"paid_live_data"`
			Basis struct {
				NextPing             int  `json:"next_ping"`
				IsCustomizeAudioRoom bool `json:"is_customize_audio_room"`
				NeedRequestLuckybox  int  `json:"need_request_luckybox"`
				SecretRoom           int  `json:"secret_room"`
				ForeignUserRoom      int  `json:"foreign_user_room"`
			} `json:"basis"`
			ShortTouchAreaConfig struct {
				Elements struct {
					Num1 struct {
						Type     int `json:"type"`
						Priority int `json:"priority"`
					} `json:"1"`
					Num2 struct {
						Type     int `json:"type"`
						Priority int `json:"priority"`
					} `json:"2"`
					Num3 struct {
						Type     int `json:"type"`
						Priority int `json:"priority"`
					} `json:"3"`
					Num4 struct {
						Type     int `json:"type"`
						Priority int `json:"priority"`
					} `json:"4"`
					Num5 struct {
						Type     int `json:"type"`
						Priority int `json:"priority"`
					} `json:"5"`
					Num6 struct {
						Type     int `json:"type"`
						Priority int `json:"priority"`
					} `json:"6"`
					Num7 struct {
						Type     int `json:"type"`
						Priority int `json:"priority"`
					} `json:"7"`
					Num8 struct {
						Type     int `json:"type"`
						Priority int `json:"priority"`
					} `json:"8"`
					Num9 struct {
						Type     int `json:"type"`
						Priority int `json:"priority"`
					} `json:"9"`
					Num10 struct {
						Type     int `json:"type"`
						Priority int `json:"priority"`
					} `json:"10"`
					Num12 struct {
						Type     int `json:"type"`
						Priority int `json:"priority"`
					} `json:"12"`
					Num22 struct {
						Type     int `json:"type"`
						Priority int `json:"priority"`
					} `json:"22"`
					Num27 struct {
						Type     int `json:"type"`
						Priority int `json:"priority"`
					} `json:"27"`
					Num30 struct {
						Type     int `json:"type"`
						Priority int `json:"priority"`
					} `json:"30"`
				} `json:"elements"`
				ForbiddenTypesMap struct {
				} `json:"forbidden_types_map"`
				TempStateConditionMap struct {
					Num1 struct {
						Type struct {
							StrategyType int `json:"strategy_type"`
							Priority     int `json:"priority"`
						} `json:"type"`
						MinimumGap int `json:"minimum_gap"`
					} `json:"1"`
					Num2 struct {
						Type struct {
							StrategyType int `json:"strategy_type"`
							Priority     int `json:"priority"`
						} `json:"type"`
						MinimumGap int `json:"minimum_gap"`
					} `json:"2"`
					Num3 struct {
						Type struct {
							StrategyType int `json:"strategy_type"`
							Priority     int `json:"priority"`
						} `json:"type"`
						MinimumGap int `json:"minimum_gap"`
					} `json:"3"`
					Num4 struct {
						Type struct {
							StrategyType int `json:"strategy_type"`
							Priority     int `json:"priority"`
						} `json:"type"`
						MinimumGap int `json:"minimum_gap"`
					} `json:"4"`
					Num5 struct {
						Type struct {
							StrategyType int `json:"strategy_type"`
							Priority     int `json:"priority"`
						} `json:"type"`
						MinimumGap int `json:"minimum_gap"`
					} `json:"5"`
					Num6 struct {
						Type struct {
							StrategyType int `json:"strategy_type"`
							Priority     int `json:"priority"`
						} `json:"type"`
						MinimumGap int `json:"minimum_gap"`
					} `json:"6"`
					Num7 struct {
						Type struct {
							StrategyType int `json:"strategy_type"`
							Priority     int `json:"priority"`
						} `json:"type"`
						MinimumGap int `json:"minimum_gap"`
					} `json:"7"`
				} `json:"temp_state_condition_map"`
				TempStateStrategy struct {
					Num4 struct {
						ShortTouchType int `json:"short_touch_type"`
						StrategyMap    struct {
							Num1 struct {
								Type struct {
									StrategyType int `json:"strategy_type"`
									Priority     int `json:"priority"`
								} `json:"type"`
								Duration       int    `json:"duration"`
								StrategyMethod string `json:"strategy_method"`
							} `json:"1"`
							Num2 struct {
								Type struct {
									StrategyType int `json:"strategy_type"`
									Priority     int `json:"priority"`
								} `json:"type"`
								Duration       int    `json:"duration"`
								StrategyMethod string `json:"strategy_method"`
							} `json:"2"`
							Num3 struct {
								Type struct {
									StrategyType int `json:"strategy_type"`
									Priority     int `json:"priority"`
								} `json:"type"`
								Duration       int    `json:"duration"`
								StrategyMethod string `json:"strategy_method"`
							} `json:"3"`
							Num6 struct {
								Type struct {
									StrategyType int `json:"strategy_type"`
									Priority     int `json:"priority"`
								} `json:"type"`
								Duration       int    `json:"duration"`
								StrategyMethod string `json:"strategy_method"`
							} `json:"6"`
							Num7 struct {
								Type struct {
									StrategyType int `json:"strategy_type"`
									Priority     int `json:"priority"`
								} `json:"type"`
								Duration       int    `json:"duration"`
								StrategyMethod string `json:"strategy_method"`
							} `json:"7"`
						} `json:"strategy_map"`
					} `json:"4"`
					Num7 struct {
						ShortTouchType int `json:"short_touch_type"`
						StrategyMap    struct {
							Num1 struct {
								Type struct {
									StrategyType int `json:"strategy_type"`
									Priority     int `json:"priority"`
								} `json:"type"`
								Duration       int    `json:"duration"`
								StrategyMethod string `json:"strategy_method"`
							} `json:"1"`
							Num2 struct {
								Type struct {
									StrategyType int `json:"strategy_type"`
									Priority     int `json:"priority"`
								} `json:"type"`
								Duration       int    `json:"duration"`
								StrategyMethod string `json:"strategy_method"`
							} `json:"2"`
							Num3 struct {
								Type struct {
									StrategyType int `json:"strategy_type"`
									Priority     int `json:"priority"`
								} `json:"type"`
								Duration       int    `json:"duration"`
								StrategyMethod string `json:"strategy_method"`
							} `json:"3"`
							Num4 struct {
								Type struct {
									StrategyType int `json:"strategy_type"`
									Priority     int `json:"priority"`
								} `json:"type"`
								Duration       int    `json:"duration"`
								StrategyMethod string `json:"strategy_method"`
							} `json:"4"`
							Num5 struct {
								Type struct {
									StrategyType int `json:"strategy_type"`
									Priority     int `json:"priority"`
								} `json:"type"`
								Duration       int    `json:"duration"`
								StrategyMethod string `json:"strategy_method"`
							} `json:"5"`
							Num6 struct {
								Type struct {
									StrategyType int `json:"strategy_type"`
									Priority     int `json:"priority"`
								} `json:"type"`
								Duration       int    `json:"duration"`
								StrategyMethod string `json:"strategy_method"`
							} `json:"6"`
						} `json:"strategy_map"`
					} `json:"7"`
					Num8 struct {
						ShortTouchType int `json:"short_touch_type"`
						StrategyMap    struct {
							Num1 struct {
								Type struct {
									StrategyType int `json:"strategy_type"`
									Priority     int `json:"priority"`
								} `json:"type"`
								Duration       int    `json:"duration"`
								StrategyMethod string `json:"strategy_method"`
							} `json:"1"`
							Num2 struct {
								Type struct {
									StrategyType int `json:"strategy_type"`
									Priority     int `json:"priority"`
								} `json:"type"`
								Duration       int    `json:"duration"`
								StrategyMethod string `json:"strategy_method"`
							} `json:"2"`
						} `json:"strategy_map"`
					} `json:"8"`
					Num97 struct {
						ShortTouchType int `json:"short_touch_type"`
						StrategyMap    struct {
							Num1 struct {
								Type struct {
									StrategyType int `json:"strategy_type"`
									Priority     int `json:"priority"`
								} `json:"type"`
								Duration       int    `json:"duration"`
								StrategyMethod string `json:"strategy_method"`
							} `json:"1"`
							Num2 struct {
								Type struct {
									StrategyType int `json:"strategy_type"`
									Priority     int `json:"priority"`
								} `json:"type"`
								Duration       int    `json:"duration"`
								StrategyMethod string `json:"strategy_method"`
							} `json:"2"`
							Num3 struct {
								Type struct {
									StrategyType int `json:"strategy_type"`
									Priority     int `json:"priority"`
								} `json:"type"`
								Duration       int    `json:"duration"`
								StrategyMethod string `json:"strategy_method"`
							} `json:"3"`
							Num5 struct {
								Type struct {
									StrategyType int `json:"strategy_type"`
									Priority     int `json:"priority"`
								} `json:"type"`
								Duration       int    `json:"duration"`
								StrategyMethod string `json:"strategy_method"`
							} `json:"5"`
							Num6 struct {
								Type struct {
									StrategyType int `json:"strategy_type"`
									Priority     int `json:"priority"`
								} `json:"type"`
								Duration       int    `json:"duration"`
								StrategyMethod string `json:"strategy_method"`
							} `json:"6"`
							Num7 struct {
								Type struct {
									StrategyType int `json:"strategy_type"`
									Priority     int `json:"priority"`
								} `json:"type"`
								Duration       int    `json:"duration"`
								StrategyMethod string `json:"strategy_method"`
							} `json:"7"`
						} `json:"strategy_map"`
					} `json:"97"`
					Num136 struct {
						ShortTouchType int `json:"short_touch_type"`
						StrategyMap    struct {
							Num1 struct {
								Type struct {
									StrategyType int `json:"strategy_type"`
									Priority     int `json:"priority"`
								} `json:"type"`
								Duration       int    `json:"duration"`
								StrategyMethod string `json:"strategy_method"`
							} `json:"1"`
							Num2 struct {
								Type struct {
									StrategyType int `json:"strategy_type"`
									Priority     int `json:"priority"`
								} `json:"type"`
								Duration       int    `json:"duration"`
								StrategyMethod string `json:"strategy_method"`
							} `json:"2"`
						} `json:"strategy_map"`
					} `json:"136"`
					Num141 struct {
						ShortTouchType int `json:"short_touch_type"`
						StrategyMap    struct {
							Num1 struct {
								Type struct {
									StrategyType int `json:"strategy_type"`
									Priority     int `json:"priority"`
								} `json:"type"`
								Duration       int    `json:"duration"`
								StrategyMethod string `json:"strategy_method"`
							} `json:"1"`
							Num2 struct {
								Type struct {
									StrategyType int `json:"strategy_type"`
									Priority     int `json:"priority"`
								} `json:"type"`
								Duration       int    `json:"duration"`
								StrategyMethod string `json:"strategy_method"`
							} `json:"2"`
							Num3 struct {
								Type struct {
									StrategyType int `json:"strategy_type"`
									Priority     int `json:"priority"`
								} `json:"type"`
								Duration       int    `json:"duration"`
								StrategyMethod string `json:"strategy_method"`
							} `json:"3"`
						} `json:"strategy_map"`
					} `json:"141"`
					Num149 struct {
						ShortTouchType int `json:"short_touch_type"`
						StrategyMap    struct {
							Num1 struct {
								Type struct {
									StrategyType int `json:"strategy_type"`
									Priority     int `json:"priority"`
								} `json:"type"`
								Duration       int    `json:"duration"`
								StrategyMethod string `json:"strategy_method"`
							} `json:"1"`
							Num2 struct {
								Type struct {
									StrategyType int `json:"strategy_type"`
									Priority     int `json:"priority"`
								} `json:"type"`
								Duration       int    `json:"duration"`
								StrategyMethod string `json:"strategy_method"`
							} `json:"2"`
						} `json:"strategy_map"`
					} `json:"149"`
					Num152 struct {
						ShortTouchType int `json:"short_touch_type"`
						StrategyMap    struct {
							Num1 struct {
								Type struct {
									StrategyType int `json:"strategy_type"`
									Priority     int `json:"priority"`
								} `json:"type"`
								Duration       int    `json:"duration"`
								StrategyMethod string `json:"strategy_method"`
							} `json:"1"`
							Num2 struct {
								Type struct {
									StrategyType int `json:"strategy_type"`
									Priority     int `json:"priority"`
								} `json:"type"`
								Duration       int    `json:"duration"`
								StrategyMethod string `json:"strategy_method"`
							} `json:"2"`
						} `json:"strategy_map"`
					} `json:"152"`
					Num153 struct {
						ShortTouchType int `json:"short_touch_type"`
						StrategyMap    struct {
							Num1 struct {
								Type struct {
									StrategyType int `json:"strategy_type"`
									Priority     int `json:"priority"`
								} `json:"type"`
								Duration       int    `json:"duration"`
								StrategyMethod string `json:"strategy_method"`
							} `json:"1"`
							Num2 struct {
								Type struct {
									StrategyType int `json:"strategy_type"`
									Priority     int `json:"priority"`
								} `json:"type"`
								Duration       int    `json:"duration"`
								StrategyMethod string `json:"strategy_method"`
							} `json:"2"`
							Num4 struct {
								Type struct {
									StrategyType int `json:"strategy_type"`
									Priority     int `json:"priority"`
								} `json:"type"`
								Duration       int    `json:"duration"`
								StrategyMethod string `json:"strategy_method"`
							} `json:"4"`
						} `json:"strategy_map"`
					} `json:"153"`
					Num159 struct {
						ShortTouchType int `json:"short_touch_type"`
						StrategyMap    struct {
							Num1 struct {
								Type struct {
									StrategyType int `json:"strategy_type"`
									Priority     int `json:"priority"`
								} `json:"type"`
								Duration       int    `json:"duration"`
								StrategyMethod string `json:"strategy_method"`
							} `json:"1"`
						} `json:"strategy_map"`
					} `json:"159"`
					Num161 struct {
						ShortTouchType int `json:"short_touch_type"`
						StrategyMap    struct {
							Num1 struct {
								Type struct {
									StrategyType int `json:"strategy_type"`
									Priority     int `json:"priority"`
								} `json:"type"`
								Duration       int    `json:"duration"`
								StrategyMethod string `json:"strategy_method"`
							} `json:"1"`
							Num2 struct {
								Type struct {
									StrategyType int `json:"strategy_type"`
									Priority     int `json:"priority"`
								} `json:"type"`
								Duration       int    `json:"duration"`
								StrategyMethod string `json:"strategy_method"`
							} `json:"2"`
						} `json:"strategy_map"`
					} `json:"161"`
					Num210 struct {
						ShortTouchType int `json:"short_touch_type"`
						StrategyMap    struct {
							Num1 struct {
								Type struct {
									StrategyType int `json:"strategy_type"`
									Priority     int `json:"priority"`
								} `json:"type"`
								Duration       int    `json:"duration"`
								StrategyMethod string `json:"strategy_method"`
							} `json:"1"`
						} `json:"strategy_map"`
					} `json:"210"`
					Num306 struct {
						ShortTouchType int `json:"short_touch_type"`
						StrategyMap    struct {
							Num3 struct {
								Type struct {
									StrategyType int `json:"strategy_type"`
									Priority     int `json:"priority"`
								} `json:"type"`
								Duration       int    `json:"duration"`
								StrategyMethod string `json:"strategy_method"`
							} `json:"3"`
						} `json:"strategy_map"`
					} `json:"306"`
					Num307 struct {
						ShortTouchType int `json:"short_touch_type"`
						StrategyMap    struct {
							Num4 struct {
								Type struct {
									StrategyType int `json:"strategy_type"`
									Priority     int `json:"priority"`
								} `json:"type"`
								Duration       int    `json:"duration"`
								StrategyMethod string `json:"strategy_method"`
							} `json:"4"`
						} `json:"strategy_map"`
					} `json:"307"`
					Num308 struct {
						ShortTouchType int `json:"short_touch_type"`
						StrategyMap    struct {
							Num5 struct {
								Type struct {
									StrategyType int `json:"strategy_type"`
									Priority     int `json:"priority"`
								} `json:"type"`
								Duration       int    `json:"duration"`
								StrategyMethod string `json:"strategy_method"`
							} `json:"5"`
						} `json:"strategy_map"`
					} `json:"308"`
					Num311 struct {
						ShortTouchType int `json:"short_touch_type"`
						StrategyMap    struct {
							Num3 struct {
								Type struct {
									StrategyType int `json:"strategy_type"`
									Priority     int `json:"priority"`
								} `json:"type"`
								Duration       int    `json:"duration"`
								StrategyMethod string `json:"strategy_method"`
							} `json:"3"`
						} `json:"strategy_map"`
					} `json:"311"`
					Num312 struct {
						ShortTouchType int `json:"short_touch_type"`
						StrategyMap    struct {
							Num1 struct {
								Type struct {
									StrategyType int `json:"strategy_type"`
									Priority     int `json:"priority"`
								} `json:"type"`
								Duration       int    `json:"duration"`
								StrategyMethod string `json:"strategy_method"`
							} `json:"1"`
						} `json:"strategy_map"`
					} `json:"312"`
					Num313 struct {
						ShortTouchType int `json:"short_touch_type"`
						StrategyMap    struct {
							Num2 struct {
								Type struct {
									StrategyType int `json:"strategy_type"`
									Priority     int `json:"priority"`
								} `json:"type"`
								Duration       int    `json:"duration"`
								StrategyMethod string `json:"strategy_method"`
							} `json:"2"`
						} `json:"strategy_map"`
					} `json:"313"`
				} `json:"temp_state_strategy"`
				StrategyFeatWhitelist    []string `json:"strategy_feat_whitelist"`
				TempStateGlobalCondition struct {
					DurationGap         int   `json:"duration_gap"`
					AllowCount          int   `json:"allow_count"`
					IgnoreStrategyTypes []int `json:"ignore_strategy_types"`
				} `json:"temp_state_global_condition"`
			} `json:"short_touch_area_config"`
			ReqUser struct {
				UserShareRoomScore  int `json:"user_share_room_score"`
				EnterUserDeviceType int `json:"enter_user_device_type"`
			} `json:"req_user"`
			Others struct {
				DecoDetail struct {
				} `json:"deco_detail"`
				MorePanelInfo struct {
					LoadStrategy int `json:"load_strategy"`
				} `json:"more_panel_info"`
				AppointmentInfo struct {
					AppointmentID int  `json:"appointment_id"`
					IsSubscribe   bool `json:"is_subscribe"`
				} `json:"appointment_info"`
				WebSkin struct {
					EnableSkin bool `json:"enable_skin"`
				} `json:"web_skin"`
				Programme struct {
					EnableProgramme bool `json:"enable_programme"`
				} `json:"programme"`
				WebLivePortOptimization struct {
					StrategyConfig struct {
						Background struct {
							StrategyType         int    `json:"strategy_type"`
							UseConfigDuration    bool   `json:"use_config_duration"`
							PauseMonitorDuration string `json:"pause_monitor_duration"`
						} `json:"background"`
						Detail struct {
							StrategyType         int    `json:"strategy_type"`
							UseConfigDuration    bool   `json:"use_config_duration"`
							PauseMonitorDuration string `json:"pause_monitor_duration"`
						} `json:"detail"`
						Tab struct {
							StrategyType         int    `json:"strategy_type"`
							UseConfigDuration    bool   `json:"use_config_duration"`
							PauseMonitorDuration string `json:"pause_monitor_duration"`
						} `json:"tab"`
					} `json:"strategy_config"`
					StrategyExtra string `json:"strategy_extra"`
				} `json:"web_live_port_optimization"`
				LvideoItemID          int `json:"lvideo_item_id"`
				RecognitionContainers struct {
					RecognitionCandidates []interface{} `json:"recognition_candidates"`
				} `json:"recognition_containers"`
				AnchorTogetherLive struct {
					IsTogetherLive int           `json:"is_together_live"`
					UserList       []interface{} `json:"user_list"`
					Title          string        `json:"title"`
					SchemaURL      string        `json:"schema_url"`
					Scene          int           `json:"scene"`
					IsShow         bool          `json:"is_show"`
				} `json:"anchor_together_live"`
				MosaicVersion         int           `json:"mosaic_version"`
				MetricTrackerDataList []interface{} `json:"metric_tracker_data_list"`
			} `json:"others"`
			AdminUserOpenIds    []interface{} `json:"admin_user_open_ids"`
			AdminUserOpenIdsStr []interface{} `json:"admin_user_open_ids_str"`
			OwnerOpenIDStr      string        `json:"owner_open_id_str"`
		} `json:"data"`
		EnterRoomID string `json:"enter_room_id"`
		Extra       struct {
			DiggColor         string `json:"digg_color"`
			PayScores         string `json:"pay_scores"`
			IsOfficialChannel bool   `json:"is_official_channel"`
			Signature         string `json:"signature"`
		} `json:"extra"`
		User struct {
			IDStr       string `json:"id_str"`
			SecUID      string `json:"sec_uid"`
			Nickname    string `json:"nickname"`
			AvatarThumb struct {
				URLList []string `json:"url_list"`
			} `json:"avatar_thumb"`
			FollowInfo struct {
				FollowStatus    int    `json:"follow_status"`
				FollowStatusStr string `json:"follow_status_str"`
			} `json:"follow_info"`
			ForeignUser int    `json:"foreign_user"`
			OpenIDStr   string `json:"open_id_str"`
		} `json:"user"`
		QrcodeURL        string `json:"qrcode_url"`
		EnterMode        int    `json:"enter_mode"`
		RoomStatus       int    `json:"room_status"`
		PartitionRoadMap struct {
			Partition struct {
				IDStr string `json:"id_str"`
				Type  int    `json:"type"`
				Title string `json:"title"`
			} `json:"partition"`
			SubPartition struct {
				Partition struct {
					IDStr string `json:"id_str"`
					Type  int    `json:"type"`
					Title string `json:"title"`
				} `json:"partition"`
			} `json:"sub_partition"`
		} `json:"partition_road_map"`
		SimilarRooms []struct {
			Room struct {
				IDStr        string `json:"id_str"`
				Status       int    `json:"status"`
				StatusStr    string `json:"status_str"`
				Title        string `json:"title"`
				UserCountStr string `json:"user_count_str"`
				Cover        struct {
					URLList []string `json:"url_list"`
				} `json:"cover"`
				StreamURL struct {
					FlvPullURL struct {
						FULLHD1 string `json:"FULL_HD1"`
						HD1     string `json:"HD1"`
						SD1     string `json:"SD1"`
						SD2     string `json:"SD2"`
					} `json:"flv_pull_url"`
					DefaultResolution string `json:"default_resolution"`
					HlsPullURLMap     struct {
						FULLHD1 string `json:"FULL_HD1"`
						HD1     string `json:"HD1"`
						SD1     string `json:"SD1"`
						SD2     string `json:"SD2"`
					} `json:"hls_pull_url_map"`
					HlsPullURL        string `json:"hls_pull_url"`
					StreamOrientation int    `json:"stream_orientation"`
					LiveCoreSdkData   struct {
						PullData struct {
							Options struct {
								DefaultQuality struct {
									Name              string `json:"name"`
									SdkKey            string `json:"sdk_key"`
									VCodec            string `json:"v_codec"`
									Resolution        string `json:"resolution"`
									Level             int    `json:"level"`
									VBitRate          int    `json:"v_bit_rate"`
									AdditionalContent string `json:"additional_content"`
									Fps               int    `json:"fps"`
									Disable           int    `json:"disable"`
								} `json:"default_quality"`
								Qualities []interface{} `json:"qualities"`
							} `json:"options"`
							StreamData string `json:"stream_data"`
						} `json:"pull_data"`
					} `json:"live_core_sdk_data"`
					Extra struct {
						Height                  int  `json:"height"`
						Width                   int  `json:"width"`
						Fps                     int  `json:"fps"`
						MaxBitrate              int  `json:"max_bitrate"`
						MinBitrate              int  `json:"min_bitrate"`
						DefaultBitrate          int  `json:"default_bitrate"`
						BitrateAdaptStrategy    int  `json:"bitrate_adapt_strategy"`
						AnchorInteractProfile   int  `json:"anchor_interact_profile"`
						AudienceInteractProfile int  `json:"audience_interact_profile"`
						HardwareEncode          bool `json:"hardware_encode"`
						VideoProfile            int  `json:"video_profile"`
						H265Enable              bool `json:"h265_enable"`
						GopSec                  int  `json:"gop_sec"`
						BframeEnable            bool `json:"bframe_enable"`
						Roi                     bool `json:"roi"`
						SwRoi                   bool `json:"sw_roi"`
						Bytevc1Enable           bool `json:"bytevc1_enable"`
					} `json:"extra"`
					PullDatas struct {
					} `json:"pull_datas"`
				} `json:"stream_url"`
				MosaicStatus    int           `json:"mosaic_status"`
				MosaicStatusStr string        `json:"mosaic_status_str"`
				AdminUserIds    []interface{} `json:"admin_user_ids"`
				AdminUserIdsStr []interface{} `json:"admin_user_ids_str"`
				Owner           struct {
					IDStr       string `json:"id_str"`
					SecUID      string `json:"sec_uid"`
					Nickname    string `json:"nickname"`
					AvatarThumb struct {
						URLList []string `json:"url_list"`
					} `json:"avatar_thumb"`
					FollowInfo struct {
						FollowStatus    int    `json:"follow_status"`
						FollowStatusStr string `json:"follow_status_str"`
					} `json:"follow_info"`
					Subscribe struct {
						IsMember     bool `json:"is_member"`
						Level        int  `json:"level"`
						IdentityType int  `json:"identity_type"`
						BuyType      int  `json:"buy_type"`
						Open         int  `json:"open"`
					} `json:"subscribe"`
					ForeignUser int    `json:"foreign_user"`
					OpenIDStr   string `json:"open_id_str"`
				} `json:"owner"`
				LiveRoomMode int `json:"live_room_mode"`
				Stats        struct {
					TotalUserDesp string `json:"total_user_desp"`
					LikeCount     int    `json:"like_count"`
					TotalUserStr  string `json:"total_user_str"`
					UserCountStr  string `json:"user_count_str"`
				} `json:"stats"`
				HasCommerceGoods bool `json:"has_commerce_goods"`
				LinkerMap        struct {
				} `json:"linker_map"`
				RoomViewStats struct {
					IsHidden            bool   `json:"is_hidden"`
					DisplayShort        string `json:"display_short"`
					DisplayMiddle       string `json:"display_middle"`
					DisplayLong         string `json:"display_long"`
					DisplayValue        int    `json:"display_value"`
					DisplayVersion      int    `json:"display_version"`
					Incremental         bool   `json:"incremental"`
					DisplayType         int    `json:"display_type"`
					DisplayShortAnchor  string `json:"display_short_anchor"`
					DisplayMiddleAnchor string `json:"display_middle_anchor"`
					DisplayLongAnchor   string `json:"display_long_anchor"`
				} `json:"room_view_stats"`
				EcomData struct {
					RedsShowInfos []interface{} `json:"reds_show_infos"`
					RoomCartV2    struct {
						ShowCart int `json:"show_cart"`
					} `json:"room_cart_v2"`
					InstantType int `json:"instant_type"`
				} `json:"ecom_data"`
				AnchorABMap struct {
				} `json:"AnchorABMap"`
				LikeCount      int    `json:"like_count"`
				OwnerUserIDStr string `json:"owner_user_id_str"`
				PaidLiveData   struct {
					PaidType           int  `json:"paid_type"`
					ViewRight          int  `json:"view_right"`
					Duration           int  `json:"duration"`
					Delivery           int  `json:"delivery"`
					NeedDeliveryNotice bool `json:"need_delivery_notice"`
					AnchorRight        int  `json:"anchor_right"`
					PayAbType          int  `json:"pay_ab_type"`
					PrivilegeInfo      struct {
					} `json:"privilege_info"`
					PrivilegeInfoMap struct {
					} `json:"privilege_info_map"`
					MaxPreviewDuration int `json:"max_preview_duration"`
				} `json:"paid_live_data"`
				Others struct {
					WebLivePortOptimization struct {
						StrategyConfig struct {
							Background struct {
								StrategyType         int    `json:"strategy_type"`
								UseConfigDuration    bool   `json:"use_config_duration"`
								PauseMonitorDuration string `json:"pause_monitor_duration"`
							} `json:"background"`
							Detail struct {
								StrategyType         int    `json:"strategy_type"`
								UseConfigDuration    bool   `json:"use_config_duration"`
								PauseMonitorDuration string `json:"pause_monitor_duration"`
							} `json:"detail"`
							Tab struct {
								StrategyType         int    `json:"strategy_type"`
								UseConfigDuration    bool   `json:"use_config_duration"`
								PauseMonitorDuration string `json:"pause_monitor_duration"`
							} `json:"tab"`
						} `json:"strategy_config"`
						StrategyExtra string `json:"strategy_extra"`
					} `json:"web_live_port_optimization"`
					LvideoItemID  int `json:"lvideo_item_id"`
					MosaicVersion int `json:"mosaic_version"`
					WebData       struct {
						AdditionalStreamURL struct {
							Provider       int    `json:"provider"`
							ID             int64  `json:"id"`
							IDStr          string `json:"id_str"`
							ResolutionName struct {
								FULLHD1 string `json:"FULL_HD1"`
								HD1     string `json:"HD1"`
								ORIGION string `json:"ORIGION"`
								SD1     string `json:"SD1"`
								SD2     string `json:"SD2"`
							} `json:"resolution_name"`
							DefaultResolution string `json:"default_resolution"`
							Extra             struct {
								Height                  int  `json:"height"`
								Width                   int  `json:"width"`
								Fps                     int  `json:"fps"`
								MaxBitrate              int  `json:"max_bitrate"`
								MinBitrate              int  `json:"min_bitrate"`
								DefaultBitrate          int  `json:"default_bitrate"`
								BitrateAdaptStrategy    int  `json:"bitrate_adapt_strategy"`
								AnchorInteractProfile   int  `json:"anchor_interact_profile"`
								AudienceInteractProfile int  `json:"audience_interact_profile"`
								HardwareEncode          bool `json:"hardware_encode"`
								VideoProfile            int  `json:"video_profile"`
								H265Enable              bool `json:"h265_enable"`
								GopSec                  int  `json:"gop_sec"`
								BframeEnable            bool `json:"bframe_enable"`
								Roi                     bool `json:"roi"`
								SwRoi                   bool `json:"sw_roi"`
								Bytevc1Enable           bool `json:"bytevc1_enable"`
							} `json:"extra"`
							RtmpPushURL string `json:"rtmp_push_url"`
							RtmpPullURL string `json:"rtmp_pull_url"`
							FlvPullURL  struct {
								FULLHD1 string `json:"FULL_HD1"`
								HD1     string `json:"HD1"`
								SD1     string `json:"SD1"`
								SD2     string `json:"SD2"`
							} `json:"flv_pull_url"`
							CandidateResolution []string `json:"candidate_resolution"`
							HlsPullURL          string   `json:"hls_pull_url"`
							HlsPullURLParams    string   `json:"hls_pull_url_params"`
							RtmpPullURLParams   string   `json:"rtmp_pull_url_params"`
							FlvPullURLParams    struct {
								HD1 string `json:"HD1"`
								SD1 string `json:"SD1"`
								SD2 string `json:"SD2"`
							} `json:"flv_pull_url_params"`
							RtmpPushURLParams string        `json:"rtmp_push_url_params"`
							PushUrls          []interface{} `json:"push_urls"`
							LiveCoreSdkData   struct {
								PullData struct {
									StreamData string `json:"stream_data"`
									Options    struct {
										DefaultQuality struct {
											Name              string `json:"name"`
											SdkKey            string `json:"sdk_key"`
											VCodec            string `json:"v_codec"`
											Resolution        string `json:"resolution"`
											Level             int    `json:"level"`
											VBitRate          int    `json:"v_bit_rate"`
											AdditionalContent string `json:"additional_content"`
											Fps               int    `json:"fps"`
											Disable           int    `json:"disable"`
										} `json:"default_quality"`
										Qualities []struct {
											Name              string `json:"name"`
											SdkKey            string `json:"sdk_key"`
											VCodec            string `json:"v_codec"`
											Resolution        string `json:"resolution"`
											Level             int    `json:"level"`
											VBitRate          int    `json:"v_bit_rate"`
											AdditionalContent string `json:"additional_content"`
											Fps               int    `json:"fps"`
											Disable           int    `json:"disable"`
										} `json:"qualities"`
										VpassDefault bool `json:"vpass_default"`
									} `json:"options"`
									Version            int `json:"version"`
									HlsDataUnencrypted struct {
									} `json:"hls_data_unencrypted"`
									Kind int `json:"kind"`
									Hls  []struct {
										URL         string `json:"url"`
										QualityName string `json:"quality_name"`
										Params      string `json:"params"`
									} `json:"Hls"`
									Flv []struct {
										URL         string `json:"url"`
										QualityName string `json:"quality_name"`
										Params      string `json:"params"`
									} `json:"Flv"`
									Codec            string `json:"codec"`
									CompensatoryData string `json:"compensatory_data"`
								} `json:"pull_data"`
								PushData struct {
									ResolutionParams struct {
									} `json:"resolution_params"`
									PushStreamLevel int    `json:"push_stream_level"`
									PreSchedule     bool   `json:"pre_schedule"`
									RtmpPushURL     string `json:"rtmp_push_url"`
									PushParams      string `json:"push_params"`
									Kind            int    `json:"kind"`
									StreamID        int    `json:"stream_id"`
									StreamIDStr     string `json:"stream_id_str"`
								} `json:"push_data"`
								Size string `json:"size"`
							} `json:"live_core_sdk_data"`
							HlsPullURLMap struct {
								FULLHD1 string `json:"FULL_HD1"`
								HD1     string `json:"HD1"`
								SD1     string `json:"SD1"`
								SD2     string `json:"SD2"`
							} `json:"hls_pull_url_map"`
							CompletePushUrls  []interface{} `json:"complete_push_urls"`
							StreamControlType int           `json:"stream_control_type"`
							StreamOrientation int           `json:"stream_orientation"`
							PushStreamType    int           `json:"push_stream_type"`
							PullDatas         struct {
							} `json:"pull_datas"`
							Play struct {
								Horizontal string `json:"horizontal"`
								Vertical   string `json:"vertical"`
							} `json:"play"`
							PushDatas struct {
							} `json:"push_datas"`
							VrType int `json:"vr_type"`
						} `json:"additional_stream_url"`
					} `json:"web_data"`
					MetricTrackerDataList []interface{} `json:"metric_tracker_data_list"`
				} `json:"others"`
				AdminUserOpenIds    []interface{} `json:"admin_user_open_ids"`
				AdminUserOpenIdsStr []interface{} `json:"admin_user_open_ids_str"`
				OwnerOpenIDStr      string        `json:"owner_open_id_str"`
			} `json:"room"`
			TagName     string `json:"tag_name"`
			UniqID      string `json:"uniq_id"`
			WebRid      string `json:"web_rid"`
			IsRecommend int    `json:"is_recommend"`
			TitleType   int    `json:"title_type"`
			CoverType   int    `json:"cover_type"`
		} `json:"similar_rooms"`
		SharkDecisionConf string `json:"shark_decision_conf"`
		WebStreamURL      struct {
			FlvPullURL struct {
			} `json:"flv_pull_url"`
			DefaultResolution string `json:"default_resolution"`
			HlsPullURLMap     struct {
			} `json:"hls_pull_url_map"`
			HlsPullURL        string `json:"hls_pull_url"`
			StreamOrientation int    `json:"stream_orientation"`
			PullDatas         struct {
			} `json:"pull_datas"`
		} `json:"web_stream_url"`
		LoginLead struct {
			IsLogin bool `json:"is_login"`
			Level   int  `json:"level"`
			Items   struct {
			} `json:"items"`
		} `json:"login_lead"`
	} `json:"data"`
	Extra struct {
		Now int64 `json:"now"`
	} `json:"extra"`
	StatusCode int `json:"status_code"`
}

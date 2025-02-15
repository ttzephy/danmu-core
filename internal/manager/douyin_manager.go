package manager

import (
	"danmu-core/core/live"
	"danmu-core/internal/handler"
	"danmu-core/internal/model"
	"danmu-core/logger"
	"danmu-core/utils"
	"fmt"
	"sync"
	"time"
)

var (
	taskList map[int64]*DouyinTask
	mu       sync.RWMutex
)

type DouyinTask struct {
	live *live.DouyinLive
	conf *model.LiveConf
}

func InitDouyinManager() {
	logger.Info().Msg("INIT DOYIN MANAGER")
	taskList = make(map[int64]*DouyinTask)
	confs, err := model.GetAllLiveConf()
	if err != nil {
		logger.Fatal().Err(err).Msg("get live conf failed")
	}
	for _, conf := range confs {
		dylive, err := live.NewDouyinLive(conf)
		if err != nil {
			logger.Warn().Err(err).Str("room_display_id", conf.RoomDisplayID).Str("room_name", conf.Name).Msg("init live failed")
			continue
		}
		dyhandler, err := handler.NewDouyinHandler(conf)
		if err != nil {
			logger.Warn().Err(err).Str("room_display_id", conf.RoomDisplayID).Str("room_name", conf.Name)
			continue
		}
		dylive.Subscribe(dyhandler.HandleDouyinMessage)
		taskList[conf.ID] = &DouyinTask{
			live: dylive,
			conf: conf,
		}
	}
	go checkAllLiveTimer()
}

func AddDouyinTask(conf *model.LiveConf) error {
	mu.Lock()
	defer mu.Unlock()
	if _, ok := taskList[conf.ID]; ok {
		return fmt.Errorf("live task already exists: %s", conf.RoomDisplayID)
	}
	dylive, err := live.NewDouyinLive(conf)
	if err != nil {
		return fmt.Errorf("init live failed: %w", err)
	}
	dyhandler, err := handler.NewDouyinHandler(conf)
	if err != nil {
		logger.Error().Err(err).Str("room_display_id", conf.RoomDisplayID).Str("room_name", conf.Name)
	}
	dylive.Subscribe(dyhandler.HandleDouyinMessage)
	taskList[conf.ID] = &DouyinTask{
		live: dylive,
		conf: conf,
	}
	isLive, err := dylive.CheckStream()
	if err != nil {
		return err
	}
	if isLive {
		go utils.SafeRun(dylive.Start)
	}
	logger.Info().Str("room_display_id", conf.RoomDisplayID).Str("room_name", conf.Name).Msg("add live task success")
	return nil
}

func DeleteDouyinTask(id int64) error {
	mu.Lock()
	defer mu.Unlock()
	task, ok := taskList[id]
	if !ok {
		return fmt.Errorf("task not found: %d", id)
	}
	task.live.Stop()
	delete(taskList, id)
	logger.Info().Str("room_display_id", task.conf.RoomDisplayID).Str("room_name", task.conf.Name).Msg("delete live task success")
	return nil
}

func UpdateDouyinTask(conf *model.LiveConf) error {
	mu.Lock()
	defer mu.Unlock()
	task, ok := taskList[conf.ID]
	if !ok {
		return fmt.Errorf("task not found: %d", conf.ID)
	}
	logger.Info().Str("room_display_id", conf.RoomDisplayID).Str("room_name", conf.Name).Msg("update live task")

	if conf.URL != task.conf.URL {
		dylive, err := live.NewDouyinLive(conf)
		if err != nil {
			return fmt.Errorf("init live failed: %w", err)
		}
		dyhandler, err := handler.NewDouyinHandler(conf)
		if err != nil {
			logger.Error().Err(err).Str("room_display_id", conf.RoomDisplayID).Str("room_name", conf.Name)
		}
		dylive.Subscribe(dyhandler.HandleDouyinMessage)
		taskList[conf.ID].live.Stop()
		taskList[conf.ID] = &DouyinTask{
			live: dylive,
			conf: conf,
		}
		return nil
	}

	if conf.Enable != task.conf.Enable {
		err := task.live.SetEnable(conf.Enable)
		if err != nil {
			return err
		}
	}
	task.conf = conf
	return nil
}

func CloseDouyinManager() {
	for _, task := range taskList {
		task.live.Stop()
		delete(taskList, task.conf.ID)
	}
}

func checkAllLiveTimer() {
	for {
		logger.Info().Msg("BEGIN TO CHECK ALL LIVE")
		for _, task := range taskList {
			if !task.conf.Enable {
				continue
			}
			utils.SafeRun(func() {
				isLive, err := task.live.CheckStream()
				if err != nil {
					logger.Warn().Err(err)
					return
				}
				if isLive {
					logger.Info().Str("url", task.conf.URL).Msg("CheckStream: live is living")
					go utils.SafeRun(task.live.Start)
				} else {
					logger.Info().Str("url", task.conf.URL).Msg("CheckStream: live is closed")
					task.live.Close()
				}
			})
		}
		time.Sleep(time.Minute * 10)
	}
}

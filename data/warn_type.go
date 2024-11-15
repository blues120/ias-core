package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"gitlab.ctyuncdn.cn/ias/ias-core/data/ent/algorithm"

	"gitlab.ctyuncdn.cn/ias/ias-core/biz"
)

var parentWarnIDCounter uint32
var subWarnIDCounter uint32

type warnTypeRepo struct {
	data   *Data
	alarms map[uint32]*biz.WarnType
	log    *log.Helper
}

func NewWarnTypeRepo(data *Data, logger log.Logger) biz.WarnTypeRepo {
	return &warnTypeRepo{
		data:   data,
		alarms: make(map[uint32]*biz.WarnType),
		log:    log.NewHelper(logger),
	}
}
func getNextParenrAlarmID() uint32 {
	parentWarnIDCounter++
	return parentWarnIDCounter
}

func getNextSubAlarmID() uint32 {
	subWarnIDCounter++
	return subWarnIDCounter
}

func (r *warnTypeRepo) List(ctx context.Context, filter *biz.WarnTypeFilter) ([]*biz.SubAlarm, error) {
	query := r.data.db.Algorithm(ctx).Query()
	//组装过滤条件

	if filter.NameEq != "" {
		query = query.Where(algorithm.AppNameEQ(filter.NameEq))
	}

	results, err := query.All(ctx)
	if err != nil {
		return nil, err
	}

	warns := make(map[string]*biz.WarnType)
	parentWarnIDCounter = 0
	subWarnIDCounter = 0
	// 构建层级结构
	for _, alarm := range results {
		warn, exists := warns[alarm.AlarmType]

		if !exists {
			warn = &biz.WarnType{
				ID:        parentWarnIDCounter,
				AppName:   alarm.AppName,
				AlarmType: alarm.AlarmType,
				AlarmName: alarm.AlarmName,
			}
			getNextParenrAlarmID()
			//类型在子列表中第一个
			warns[alarm.AlarmType] = warn
			subAlarm := biz.SubAlarm{ //类型
				ID:          getNextSubAlarmID(),
				AlarmName:   alarm.AlarmType,
				AlarmTypeID: 0,
			}
			warn.SubAlarms = append(warn.SubAlarms, subAlarm)

		}

		//子列表中的告警名称 从list下标1开始 0为类型
		subAlarm := biz.SubAlarm{ //具体告警名称
			ID:          getNextSubAlarmID(),
			AlarmName:   alarm.AlarmName,
			AlarmTypeID: warn.SubAlarms[0].ID,
		}
		warn.SubAlarms = append(warn.SubAlarms, subAlarm)

	}

	//从map改为list

	list := make([]*biz.SubAlarm, 0, len(results))
	for _, alarm := range warns {
		for _, subAlarm := range alarm.SubAlarms {
			list = append(list, &biz.SubAlarm{
				ID:          subAlarm.ID,
				AlarmName:   subAlarm.AlarmName,
				AlarmTypeID: subAlarm.AlarmTypeID,
			})
		}
	}

	return list, nil
}

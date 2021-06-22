package qbuilder_test

import (
	"testing"

	"github.com/saitofun/qlib/database/qbuilder"
	"github.com/saitofun/qlib/os/qtime"
)

type Event struct {
	qbuilder.Primary
	EventID     string     `db:"f_event_id"                  json:"eventID"`     // 事件ID
	Seq         int        `db:"f_seq,default='0'"           json:"-"`           // 分析端序号
	DevName     string     `db:"f_device_name"               json:"deviceName"`  // 设备名称
	DevID       string     `db:"f_device_id"                 json:"deviceID"`    // 设备ID
	Scene       string     `db:"f_scene"                     json:"scene"`       // 事件场景类型
	Path        string     `db:"f_path,default=''"           json:"path"`        // 事件存储路径
	Version     string     `db:"f_version,default=''"        json:"version"`     // 场景版本
	CaptureTime qtime.Time `db:"f_capture_time"              json:"captureTime"` // 事件捕获时间
	IotPushed   bool       `db:"f_iot_pushed,default='2'"    json:"iotPushed"`   // IOT推送成功标记
	IotRetry    int        `db:"f_iot_retry,default='0'"     json:"iotRetry"`    // 重试次数
	IotLastPush qtime.Time `db:"f_iot_last_push,default='0'" json:"iotLastPush"` // 事件IOT最后一次推送时间
	BusPushed   bool       `db:"f_bus_pushed,default='2'"    json:"busPushed"`   // 业务推送成功标记
	BusRetry    int        `db:"f_bus_retry,default='0'"     json:"busRetry"`    // 业务推送重试次数
	BusLastPush qtime.Time `db:"f_bus_last_push,default='0'" json:"busLastPush"` // 事件最后一次推送时间
	qbuilder.OperationTime
}

func TestDefaultNaming(t *testing.T) {
	t.Log(qbuilder.TableName("Event"))
	t.Log(qbuilder.ColumnName("Event", "CaptureTime"))
	t.Log(qbuilder.IndexName("Event", "CaptureTime"))
	t.Log(qbuilder.IndexName("Event", "CaptureTime", "IotRetry"))
	t.Log(qbuilder.UniqueIndexName("Event", "CaptureTime"))
	t.Log(qbuilder.UniqueIndexName("Event", "CaptureTime", "IotRetry"))
}

func TestModel(t *testing.T) {
	qbuilder.Model(&Event{})
}

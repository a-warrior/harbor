package notifier

import (
	"errors"
	"reflect"

	"github.com/vmware/harbor/src/common/models"
	"github.com/vmware/harbor/src/common/utils"
)

//WatchConfigChanges is used to watch the configuration changes.
func WatchConfigChanges(cfg map[string]interface{}) error {
	if cfg == nil {
		return errors.New("Empty configurations")
	}

	//Currently only watch the scan all policy change.
	if v, ok := cfg[ScanAllPolicyTopic]; ok {
		policyCfg := &models.ScanAllPolicy{}
		if err := utils.ConvertMapToStruct(policyCfg, v); err != nil {
			return err
		}

		policyNotification := ScanPolicyNotification{
			Type:      policyCfg.Type,
			DailyTime: 0,
		}

		if t, yes := policyCfg.Parm["daily_time"]; yes {
			if reflect.TypeOf(t).Kind() == reflect.Int {
				policyNotification.DailyTime = (int64)(t.(int))
			}
		}

		return Publish(ScanAllPolicyTopic, policyNotification)
	}

	return nil
}
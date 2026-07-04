package activityService

import (
	"context"
	"fmt"
	"ostadbun/pkg/constants"
)

func (a Activity) Trigger(ctx context.Context, userid int, activity Activityconstants.ActivityTriggersName) {

	defer func() {
		errC := a.UpdateRedisCash(ctx, userid)
		if errC != nil {
			//TODO log here
		}
	}()

	err := a.repo.TriggerSetter(userid, activity)

	fmt.Println("⚡️Trigger :", err)

}

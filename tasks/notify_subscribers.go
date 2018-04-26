package tasks

import (
	"fmt"
	"strings"
	"time"
)

func NotifySubscriber(args ...string) error {
	fmt.Println(fmt.Sprintf("Notifying to : %s", args[0]))
	if strings.HasPrefix(args[0], "a") || strings.HasPrefix(args[0], "c") {
		return tasks.NewErrRetryTaskLater("some error", 4 * time.Hour)
		return fmt.Errorf("forced error to requeue task")
	}
	return nil
}

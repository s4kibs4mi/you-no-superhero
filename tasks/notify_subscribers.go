package tasks

import (
	"fmt"
	"strings"
	"time"
	"github.com/RichardKnop/machinery/v1/tasks"
)

func NotifySubscriber(args ...string) error {
	fmt.Println(fmt.Sprintf("Notifying to : %s", args[0]))
	if strings.HasPrefix(args[0], "a") || strings.HasPrefix(args[0], "c") {
		return tasks.NewErrRetryTaskLater("forced error to requeue task", 10*time.Second)
	}
	return nil
}

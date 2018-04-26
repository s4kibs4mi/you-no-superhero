package main

import (
	"github.com/RichardKnop/machinery/v1/config"
	"github.com/RichardKnop/machinery/v1"
	"sync"
	"github.com/s4kibs4mi/you-no-superhero/tasks"
	t "github.com/RichardKnop/machinery/v1/tasks"
	"time"
)

var Wg sync.WaitGroup
var server *machinery.Server
var err error

const TaskNotifySubscriber = "notify_subscriber"

func main() {
	var cfg = &config.Config{
		Broker:        "amqp://batman:batman@localhost:5672/",
		DefaultQueue:  "machinery_tasks",
		ResultBackend: "amqp://batman:batman@localhost:5672/",
		AMQP: &config.AMQPConfig{
			Exchange:     "machinery_exchange",
			ExchangeType: "direct",
			BindingKey:   "machinery_task",
		},
	}

	server, err = machinery.NewServer(cfg)
	if err != nil {
		// do something with the error
	}

	server.RegisterTask(TaskNotifySubscriber, tasks.NotifySubscriber)

	go func() {
		worker := server.NewWorker("yo_no_superhero", 10)
		err := worker.Launch()
		if err != nil {
			// do something with the error
		}
	}()

	go processSubscriptionEmail()

	Wg.Add(1)
	Wg.Wait()
}

func processSubscriptionEmail() {
	time.Sleep(time.Second * 10)
	// Pull subscription emails from database
	emails := []string{"a@email.com", "b@email.com", "c@email.com", "d@email.com", "e@email.com"}

	for _, email := range emails {
		signature := &t.Signature{
			Name: TaskNotifySubscriber,
			Args: []t.Arg{
				{
					Type:  "string",
					Value: email,
				},
			},
		}
		server.SendTask(signature)
	}
}

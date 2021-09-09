package ansiblelogparserforcloudwatch

import (
	"strings"

	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs/types"
	ansiblelogparser "github.com/yukihiko-shinoda/go-ansible-log-parser"
)

func PickupMessage(events []types.OutputLogEvent) string {
	message := pickupChangedTasks(events)
	message += pickupPlayRecap(events)
	return message
}

func pickupChangedTasks(events []types.OutputLogEvent) string {
	var sliceAllChangedTask []string
	var sliceChangedTask []string
	var latestTaskName string
	for _, event := range events {
		sliceChangedTask, latestTaskName = ansiblelogparser.PickUpChangedTasks(*event.Message, latestTaskName)
		sliceAllChangedTask = append(sliceAllChangedTask, sliceChangedTask...)
	}
	return strings.Join(sliceAllChangedTask, "\n") + "\n"
}

func pickupPlayRecap(events []types.OutputLogEvent) string {
	return ansiblelogparser.TrimRecap(*last(events).Message)
}

func last(slice []types.OutputLogEvent) types.OutputLogEvent {
	return slice[len(slice)-1]
}

func PickupNumberPlayRecap(events []types.OutputLogEvent) (*ansiblelogparser.StructPlayRecap, error) {
	return ansiblelogparser.PickupNumberPlayRecap(*last(events).Message)
}

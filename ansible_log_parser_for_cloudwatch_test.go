package ansiblelogparserforcloudwatch

import (
	"reflect"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs/types"
	ansiblelogparser "github.com/yukihiko-shinoda/go-ansible-log-parser"
	"github.com/yukihiko-shinoda/go-ansible-log-parser-for-cloudwatch/_testlibraries"
)

func TestPickupMessage(t *testing.T) {
	expected := "TASK [file] *********************************************************************************************************************************************************************************************************************\n" +
		"TASK [file] *********************************************************************************************************************************************************************************************************************\n" +
		"PLAY RECAP **********************************************************************************************************************************************************************************************************************\n" +
		"localhost                  : ok=3    changed=1    unreachable=0    failed=0    skipped=0    rescued=0    ignored=0   \n\n"
	message, err := _testlibraries.LoadMessage()
	if err != nil {
		t.Errorf("%v", err)
	}
	events := []types.OutputLogEvent{
		{Message: message},
		{Message: message},
	}
	actual := PickupMessage(events)
	if actual != expected {
		t.Errorf("%v", actual)
	}
}

func TestPickupNumberPlayRecap(t *testing.T) {
	message, err := _testlibraries.LoadMessage()
	if err != nil {
		t.Errorf("%v", err)
	}
	events := []types.OutputLogEvent{
		{Message: message},
		{Message: message},
	}
	actual, err := PickupNumberPlayRecap(events)
	if err != nil {
		t.Errorf("%v", err)
	}
	if !reflect.DeepEqual(*actual, ansiblelogparser.StructPlayRecap{
		Ok:          3,
		Changed:     1,
		Unreachable: 0,
		Failed:      0,
		Skipped:     0,
		Rescued:     0,
		Ignored:     0,
	}) {
		t.Errorf("%v", *actual)
	}
}

func TestPickupNumberPlayRecapNotMatch(t *testing.T) {
	message := ""
	events := []types.OutputLogEvent{
		{Message: &message},
	}
	actual, err := PickupNumberPlayRecap(events)
	if actual != nil {
		t.Errorf("%v", actual)
	}
	if err != nil {
		t.Errorf("%v", err)
	}
}

func TestPickupChangedTasks(t *testing.T) {
	expected := "TASK [file] *********************************************************************************************************************************************************************************************************************\n" +
		"TASK [file] *********************************************************************************************************************************************************************************************************************\n"
	message, err := _testlibraries.LoadMessage()
	if err != nil {
		t.Errorf("%v", err)
	}
	events := []types.OutputLogEvent{
		{Message: message},
		{Message: message},
	}
	actual := pickupChangedTasks(events)
	if actual != expected {
		t.Errorf("%v", actual)
	}
}

func TestPickupPlayRecapLog(t *testing.T) {
	expected := "PLAY RECAP **********************************************************************************************************************************************************************************************************************\n" +
		"localhost                  : ok=3    changed=1    unreachable=0    failed=0    skipped=0    rescued=0    ignored=0   \n\n"
	message, err := _testlibraries.LoadMessage()
	if err != nil {
		t.Errorf("%v", err)
	}
	events := []types.OutputLogEvent{
		{},
		{Message: message},
	}
	actual := pickupPlayRecap(events)
	if actual != expected {
		t.Errorf("%v", actual)
	}
}

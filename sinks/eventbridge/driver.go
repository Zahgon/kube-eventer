package eventbridge

import (
	"net/url"

	"github.com/AliyunContainerService/kube-eventer/core"
	"github.com/AliyunContainerService/kube-eventer/sinks/utils"
	"github.com/alibabacloud-go/eventbridge-sdk/eventbridge"
	v1 "k8s.io/api/core/v1"
)

const (
	eventBridgeSinkName               = "EventBridgeSink"
	defaultBusName                    = "default"
	eventBridgeEndpointSchema         = "%v.eventbridge.%v.aliyuncs.com"
	eventBridgeInternalEndpointSchema = "%v.eventbridge.%v-vpc.aliyuncs.com"
	aliyunContainerServiceSource      = "acs.cs"
	eventbridgeMaxBatchSize           = 4
	eventTypeSchema                   = "cs:k8s:%vRelatedEvent"
	unknownEventType                  = "cs:k8s:UnknownTypeEvent"
)

type eventBridgeSink struct {
	client          *eventbridge.Client
	akInfo          *utils.AKInfo
	clusterId       string
	regionId        string
	accountId       string
	accessKeyId     string
	accessKeySecret string
	eventBusName    string
	internal        bool
}

type putEventsImpl func(events []*eventbridge.CloudEvent) error

func NewEventBridgeSink(uri *url.URL) (core.EventSink, error) {
	_ = "STUB: not implemented"
	return *new(core.EventSink), nil
}

func (ebSink *eventBridgeSink) Name() string { _ = "STUB: not implemented"; return "" }

// Exports data to the external storage. The function should be synchronous/blocking and finish only
// after the given EventBatch was written. This will allow sink manager to push data only to these
// sinks that finished writing the previous data.
func (ebSink *eventBridgeSink) ExportEvents(batch *core.EventBatch) {
	_ = "STUB: not implemented"
	return
}

func (ebSink *eventBridgeSink) Stop() {
	_ = "STUB: not implemented"
	// no background task, no need to implement
	return
}

func (ebSink *eventBridgeSink) toCloudEvent(event *v1.Event) (*eventbridge.CloudEvent, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ebSink *eventBridgeSink) putEvents(events []*eventbridge.CloudEvent) error {
	_ = "STUB: not implemented"
	return nil
}

func (ebSink *eventBridgeSink) exportEventsInBatch(batch *core.EventBatch, putEvents putEventsImpl) {
	_ = "STUB: not implemented"
	return
}

func (ebSink *eventBridgeSink) getClient() (*eventbridge.Client, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ebSink *eventBridgeSink) newClient() (*eventbridge.Client, error) {
	_ = "STUB: not implemented"
	// region from env
	return nil, nil
}

// region from client

// region from meta data

func (ebSink *eventBridgeSink) isAkValid() bool { _ = "STUB: not implemented"; return false }

// Creates a cloudevents subject of the form found in object metadata selfLinks
// like: acs:cs:${Region}:${Account}:${ClusterId}/${selfLink}
func (ebSink *eventBridgeSink) createEventSubject(o v1.ObjectReference) string {
	_ = "STUB: not implemented"
	return ""
}

// Core API types don't have a separate package name and only have a version string (e.g. /apis/v1/namespaces/default/pods/myPod)
// To avoid weird looking strings like "v1/versionUnknown" we'll sniff for a "." in the version

func (ebSink *eventBridgeSink) parseEventBridgeEndpoint(region string) string {
	_ = "STUB: not implemented"
	return ""
}

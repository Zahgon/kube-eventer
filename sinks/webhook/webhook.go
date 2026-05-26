package webhook

import (
	"net/url"

	"github.com/AliyunContainerService/kube-eventer/common/filters"
	"github.com/AliyunContainerService/kube-eventer/core"
	v1 "k8s.io/api/core/v1"
)

const (
	SinkName = "webHook"
	Warning  = "Warning"
	Normal   = "Normal"
)

var (
	// body template of event
	defaultBodyTemplate = `
{
	"EventType": "{{ .Type }}",
	"EventKind": "{{ .InvolvedObject.Kind }}",
	"EventReason": "{{ .Reason }}",
	"EventTime": "{{ .LastTimestamp }}",
	"EventMessage": "{{ .Message }}"
}`
)

type WebHookSink struct {
	filters                map[string]filters.Filter
	headerMap              map[string]string
	endpoint               string
	method                 string
	bodyTemplate           string
	bodyConfigMapName      string
	bodyConfigMapNamespace string
}

func (ws *WebHookSink) Name() string { _ = "STUB: not implemented"; return "" }

func (ws *WebHookSink) ExportEvents(batch *core.EventBatch) { _ = "STUB: not implemented"; return }

// send msg to generic webHook
func (ws *WebHookSink) Send(event *v1.Event) (err error) { _ = "STUB: not implemented"; return nil }

// append header to http request

func (ws *WebHookSink) RenderBodyTemplate(event *v1.Event) (body string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (ws *WebHookSink) Stop() {
	_ = "STUB: not implemented"
	// not implement
	return
}

func getLevels(level string) []string { _ = "STUB: not implemented"; return nil }

// init WebHookSink with url params
func NewWebHookSink(uri *url.URL) (*WebHookSink, error) {
	_ = "STUB: not implemented"

	// default http method
	return nil, nil
}

// set header of webHook

// namespace filter doesn't support regexp

// such as node,pod,component and so on
// kinds:https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#lists-and-simple-kinds

// reason filter support regexp.

func parseHeaders(headers []string) map[string]string { _ = "STUB: not implemented"; return nil }

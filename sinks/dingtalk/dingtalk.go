// Copyright 2015 Google Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package dingtalk

import (
	"net/url"

	"github.com/AliyunContainerService/kube-eventer/core"
	"k8s.io/api/core/v1"
	v1 "k8s.io/api/core/v1"
)

const (
	DINGTALK_SINK         = "DingTalkSink"
	WARNING           int = 2
	NORMAL            int = 1
	DEFAULT_MSG_TYPE      = "text"
	CONTENT_TYPE_JSON     = "application/json"
	LABEL_TEMPLATE        = "%s\n"
)

var (
	MSG_TEMPLATE = "Level:%s \nKind:%s \nNamespace:%s \nName:%s \nReason:%s \nTimestamp:%s \nMessage:%s"

	MSG_TEMPLATE_ARR = [][]string{
		{"Level"},
		{"Kind"},
		{"Namespace"},
		{"Name"},
		{"Reason"},
		{"Timestamp"},
		{"Message"},
	}
)

/*
*
dingtalk msg struct
*/
type DingTalkMsg struct {
	MsgType  string           `json:"msgtype"`
	Text     DingTalkText     `json:"text"`
	Markdown DingTalkMarkdown `json:"markdown"`
}

type DingTalkMarkdown struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}

type DingTalkText struct {
	Content string `json:"content"`
}

/*
*
dingtalk sink usage
--sink:dingtalk:https://oapi.dingtalk.com/robot/send?access_token=[access_token]&level=Warning&label=[label]

level: Normal or Warning. The event level greater than global level will emit.
label: some thing unique when you want to distinguish different k8s clusters.
*/
type DingTalkSink struct {
	Endpoint   string
	Namespaces []string
	Kinds      []string
	Token      string
	Level      int
	Labels     []string
	MsgType    string
	ClusterID  string
	Secret     string
	Region     string
}

func (d *DingTalkSink) Name() string { _ = "STUB: not implemented"; return "" }

func (d *DingTalkSink) Stop() {
	_ = "STUB: not implemented"
	// do nothing
	return
}

func (d *DingTalkSink) ExportEvents(batch *core.EventBatch) { _ = "STUB: not implemented"; return }

// add threshold

func (d *DingTalkSink) isEventLevelDangerous(level string) bool {
	_ = "STUB: not implemented"
	return false
}

func (d *DingTalkSink) Ding(event *v1.Event) { _ = "STUB: not implemented"; return }

func getLevel(level string) int { _ = "STUB: not implemented"; return 0 }

//score will remain 0

func createMsgFromEvent(d *DingTalkSink, event *v1.Event) *DingTalkMsg {
	_ = "STUB: not implemented"
	return nil
}

//https://open-doc.dingtalk.com/microapp/serverapi2/ye8tup#-6

//title 加不加其实没所谓,最终不会显示

//默认按文本模式推送

func NewDingTalkSink(uri *url.URL) (*DingTalkSink, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// get ding talk sign

//add extra labels

//向下兼容,覆盖以前的版本,没有这个参数的情况

// kinds:https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#lists-and-simple-kinds
// such as node,pod,component and so on

func getValues(o []string) []string { _ = "STUB: not implemented"; return nil }

func sign(t int64, secret string) string { _ = "STUB: not implemented"; return "" }

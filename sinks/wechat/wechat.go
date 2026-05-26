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

package wechat

import (
	"net/url"

	"github.com/AliyunContainerService/kube-eventer/core"
	v1 "k8s.io/api/core/v1"
)

const (
	WECHAT_SINK           = "WechatSink"
	WARNING           int = 2
	NORMAL            int = 1
	DEFAULT_MSG_TYPE      = "text"
	CONTENT_TYPE_JSON     = "application/json"
	LABEL_TEMPLATE        = "%s\n"
	//发送消息使用的url
	SEND_MSG_URL = `https://qyapi.weixin.qq.com/cgi-bin/message/send?access_token=`
	//获取token使用的url
	GET_TOKEN_URL = `https://qyapi.weixin.qq.com/cgi-bin/gettoken?corpid=`
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
wechat msg struct
*/
type WechatMsg struct {
	ToUser  string     `json:"touser"`
	ToParty string     `json:"toparty"`
	ToTag   string     `json:"totag"`
	MsgType string     `json:"msgtype"`
	AgentID int        `json:"agentid"`
	Text    WechatText `json:"text"`
	Safe    int        `json:"safe"`
}

type WechatText struct {
	Content string `json:"content"`
}

type Token struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

/*
*
wechat sink usage
--sink:wechat:https://qyapi.weixin.qq.com/cgi-bin/message/send?access_token=[access_token]&level=Warning&label=[label]

level: Normal or Warning. The event level greater than global level will emit.
label: some thing unique when you want to distinguish different k8s clusters.
*/
type WechatSink struct {
	Namespaces []string
	Kinds      []string
	CorpID     string
	CorpSecret string
	AgentID    int
	ToUser     []string
	Level      int
	Labels     []string
}

func (d *WechatSink) Name() string { _ = "STUB: not implemented"; return "" }

func (d *WechatSink) Stop() {
	_ = "STUB: not implemented"
	// do nothing
	return
}

func (d *WechatSink) ExportEvents(batch *core.EventBatch) { _ = "STUB: not implemented"; return }

// add threshold

func (d *WechatSink) isEventLevelDangerous(level string) bool {
	_ = "STUB: not implemented"
	return false
}

func (d *WechatSink) Send(event *v1.Event) { _ = "STUB: not implemented"; return }

func getToken(corp_id, corp_secret string) (at Token, err error) {
	_ = "STUB: not implemented"
	return *new(Token), nil
}

func getLevel(level string) int { _ = "STUB: not implemented"; return 0 }

//score will remain 0

func createMsgFromEvent(d *WechatSink, event *v1.Event) *WechatMsg {
	_ = "STUB: not implemented"
	return nil
}

//默认按文本模式推送

func NewWechatSink(uri *url.URL) (*WechatSink, error) { _ = "STUB: not implemented"; return nil, nil }

//使用逗号分隔需要通知的用户，如果为空则通知所有当前组下的所有用户

//add extra labels

func getValues(o []string) []string { _ = "STUB: not implemented"; return nil }

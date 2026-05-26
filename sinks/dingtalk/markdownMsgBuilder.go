package dingtalk

import (
	v1 "k8s.io/api/core/v1"
)

const (
	MARKDOWN_MSG_TYPE      = "markdown"
	MARKDOWN_TEMPLATE      = "Level: %s \n\nKind: %s \n\nNamespace: %s \n\nName: %s \n\nReason: %s \n\nTimestamp: %s \n\nMessage: %s"
	MARKDOWN_LINK_TEMPLATE = "[%s](%s)"
	MARKDOWN_TEXT_BOLD     = "**%s**"
	MARKDOWN_NEW_LINE      = "\n\n"

	URL_ALIYUN_K8S_CONSULE = "https://cs.console.aliyun.com/#/k8s"
	//阿里云 kubernetes 管理控制台, Deployment,StatefulSet,DaemonSet 有同样的URL规律
	URL_ALIYUN_RESOURCE_DETAIL_TEMPLATE = URL_ALIYUN_K8S_CONSULE + "/%s/detail/%s/%s/%s/%s/pods"
	URL_ALIYUN_POD_TEMPLATE             = URL_ALIYUN_K8S_CONSULE + "/pod/%s/%s/%s/container"
	URL_ALIYUN_CROBJOB_TEMPLATE         = URL_ALIYUN_K8S_CONSULE + "/cronjob/detail/%s/%s/%s/%s/jobs"
	URL_ALIYUN_SVC_TEMPLATE             = URL_ALIYUN_K8S_CONSULE + "/service/detail/%s/%s/%s/%s"
	URL_ALIYUN_NAMESPACE_TEMPLATE       = URL_ALIYUN_K8S_CONSULE + "/namespace"
	URL_ALIYUN_ECS_TEMPLATE             = "https://ecs.console.aliyun.com/#/server/%s/detail?regionId=%s"
)

type MarkdownMsgBuilder struct {
	Labels     []string
	Region     string
	ClusterID  string
	OutputText string
}

func NewMarkdownMsgBuilder(clusterID, region string, event *v1.Event) *MarkdownMsgBuilder {
	_ = "STUB: not implemented"
	return nil
}

//fixme:覆盖所有 event.InvolvedObject.Kind

// removeDotContent 每个 Event 由 <resource>.<UnixNano> 组成,需要去掉.后面的部分,得到 <resource>
func removeDotContent(s string) string { _ = "STUB: not implemented"; return "" }

func (m *MarkdownMsgBuilder) AddLabels(labels []string) { _ = "STUB: not implemented"; return }

func (m *MarkdownMsgBuilder) AddNodeName(nodeName string) { _ = "STUB: not implemented"; return }

func (m *MarkdownMsgBuilder) Build() string { _ = "STUB: not implemented"; return "" }

package prometheus

import (
	v1 "k8s.io/api/core/v1"
)

// always can be used when no finer-grained reason is reached. It should be the last one when used.
func always(_ *v1.Event) bool { _ = "STUB: not implemented"; return false }

func isPodImagePullBackOff(event *v1.Event) bool { _ = "STUB: not implemented"; return false }

func isFailCreatePodExceedQuota(event *v1.Event) bool { _ = "STUB: not implemented"; return false }

func isResourceInsufficient(event *v1.Event) bool { _ = "STUB: not implemented"; return false }

func isPodFailStart(event *v1.Event) bool { _ = "STUB: not implemented"; return false }

func isPodCrash(event *v1.Event) bool { _ = "STUB: not implemented"; return false }

func isDiskProvisionFailSize(event *v1.Event) bool { _ = "STUB: not implemented"; return false }

func isDiskProvisionFail(event *v1.Event) bool { _ = "STUB: not implemented"; return false }

func isFailedBindingNoStorageClass(event *v1.Event) bool { _ = "STUB: not implemented"; return false }

func isNodePLEGUnhealthy(event *v1.Event) bool { _ = "STUB: not implemented"; return false }

func isNodeNotReady(event *v1.Event) bool { _ = "STUB: not implemented"; return false }

func isClusterIPNotEnough(event *v1.Event) bool { _ = "STUB: not implemented"; return false }

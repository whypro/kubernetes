/*
Copyright 2018 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package collectors

import (
	"github.com/prometheus/client_golang/prometheus"
	stats "k8s.io/kubernetes/pkg/kubelet/apis/stats/v1alpha1"
	"k8s.io/kubernetes/pkg/kubelet/metrics"
	serverstats "k8s.io/kubernetes/pkg/kubelet/server/stats"
)

var (
	volumeStatsCapacityBytesDesc = prometheus.NewDesc(
		prometheus.BuildFQName("", metrics.KubeletSubsystem, metrics.VolumeStatsCapacityBytesKey),
		"Capacity in bytes of the volume",
		[]string{"namespace", "persistentvolumeclaim"}, nil,
	)
	volumeStatsAvailableBytesDesc = prometheus.NewDesc(
		prometheus.BuildFQName("", metrics.KubeletSubsystem, metrics.VolumeStatsAvailableBytesKey),
		"Number of available bytes in the volume",
		[]string{"namespace", "persistentvolumeclaim"}, nil,
	)
	volumeStatsUsedBytesDesc = prometheus.NewDesc(
		prometheus.BuildFQName("", metrics.KubeletSubsystem, metrics.VolumeStatsUsedBytesKey),
		"Number of used bytes in the volume",
		[]string{"namespace", "persistentvolumeclaim"}, nil,
	)
	volumeStatsInodesDesc = prometheus.NewDesc(
		prometheus.BuildFQName("", metrics.KubeletSubsystem, metrics.VolumeStatsInodesKey),
		"Maximum number of inodes in the volume",
		[]string{"namespace", "persistentvolumeclaim"}, nil,
	)
	volumeStatsInodesFreeDesc = prometheus.NewDesc(
		prometheus.BuildFQName("", metrics.KubeletSubsystem, metrics.VolumeStatsInodesFreeKey),
		"Number of free inodes in the volume",
		[]string{"namespace", "persistentvolumeclaim"}, nil,
	)
	volumeStatsInodesUsedDesc = prometheus.NewDesc(
		prometheus.BuildFQName("", metrics.KubeletSubsystem, metrics.VolumeStatsInodesUsedKey),
		"Number of used inodes in the volume",
		[]string{"namespace", "persistentvolumeclaim"}, nil,
	)
)

type volumeStatsCollecotr struct {
	summaryProvider serverstats.SummaryProvider
}

// NewVolumeStatsCollector creates a volume stats prometheus collector.
func NewVolumeStatsCollector(summaryProvider serverstats.SummaryProvider) prometheus.Collector {
	return &volumeStatsCollecotr{summaryProvider: summaryProvider}
}

// Describe implements the prometheus.Collector interface.
func (collector *volumeStatsCollecotr) Describe(ch chan<- *prometheus.Desc) {
	ch <- volumeStatsCapacityBytesDesc
	ch <- volumeStatsAvailableBytesDesc
	ch <- volumeStatsUsedBytesDesc
	ch <- volumeStatsInodesDesc
	ch <- volumeStatsInodesFreeDesc
	ch <- volumeStatsInodesUsedDesc
}

// Collect implements the prometheus.Collector interface.
func (collector *volumeStatsCollecotr) Collect(ch chan<- prometheus.Metric) {
	statsSummary, err := collector.summaryProvider.Get()
	if err != nil {
		return
	}
	addGauge := func(desc *prometheus.Desc, pvcRef *stats.PVCReference, v float64, lv ...string) {
		lv = append([]string{pvcRef.Namespace, pvcRef.Name}, lv...)
		ch <- prometheus.MustNewConstMetric(desc, prometheus.GaugeValue, v, lv...)
	}
	for _, podStat := range statsSummary.Pods {
		if podStat.VolumeStats == nil {
			continue
		}
		for _, volumeStat := range podStat.VolumeStats {
			pvcRef := volumeStat.PVCRef
			if pvcRef == nil {
				// ignore if no PVC reference
				continue
			}
			addGauge(volumeStatsCapacityBytesDesc, pvcRef, float64(*volumeStat.CapacityBytes))
			addGauge(volumeStatsAvailableBytesDesc, pvcRef, float64(*volumeStat.AvailableBytes))
			addGauge(volumeStatsUsedBytesDesc, pvcRef, float64(*volumeStat.UsedBytes))
			addGauge(volumeStatsInodesDesc, pvcRef, float64(*volumeStat.Inodes))
			addGauge(volumeStatsInodesFreeDesc, pvcRef, float64(*volumeStat.InodesFree))
			addGauge(volumeStatsInodesUsedDesc, pvcRef, float64(*volumeStat.InodesUsed))
		}
	}
}

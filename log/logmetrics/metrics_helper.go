///
//    Copyright 2019 Insolar Technologies
//
//    Licensed under the Apache License, Version 2.0 (the "License");
//    you may not use this file except in compliance with the License.
//    You may obtain a copy of the License at
//
//        http://www.apache.org/licenses/LICENSE-2.0
//
//    Unless required by applicable law or agreed to in writing, software
//    distributed under the License is distributed on an "AS IS" BASIS,
//    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//    See the License for the specific language governing permissions and
//    limitations under the License.
///

package logmetrics

import (
	"context"
	"github.com/insolar/insolar/insolar"
	"go.opencensus.io/stats"
	"time"
)

func NewMetricsHelper(recorder insolar.LogMetricsRecorder) *MetricsHelper {
	return &MetricsHelper{recorder}
}

type MetricsHelper struct {
	recorder insolar.LogMetricsRecorder
}

type DurationReportFunc func(d time.Duration)

func (p *MetricsHelper) OnNewEvent(ctx context.Context, level insolar.LogLevel) {
	if p == nil {
		return
	}
	if ctx == nil {
		ctx = GetLogLevelContext(level)
	}
	stats.Record(ctx, statLogCalls.M(1))
	if p.recorder != nil {
		p.recorder.RecordLogEvent(level)
	}
}

func (p *MetricsHelper) OnFilteredEvent(ctx context.Context, level insolar.LogLevel) {
	if p == nil {
		return
	}
	if ctx == nil {
		ctx = GetLogLevelContext(level)
	}
	stats.Record(ctx, statLogWrites.M(1))
	if p.recorder != nil {
		p.recorder.RecordLogWrite(level)
	}
}

func (p *MetricsHelper) OnWriteDuration(d time.Duration) {
	if p == nil {
		return
	}
	stats.Record(context.Background(), statLogWriteDelays.M(int64(d)))
	if p.recorder != nil {
		p.recorder.RecordLogDelay(insolar.NoLevel, d)
	}
}

func (p *MetricsHelper) GetOnWriteDurationReport() DurationReportFunc {
	if p == nil {
		return nil
	}
	return p.OnWriteDuration
}

func (p *MetricsHelper) OnWriteSkip(skippedCount int) {
	if p == nil {
		return
	}
	stats.Record(context.Background(), statLogSkips.M(int64(skippedCount)))
}

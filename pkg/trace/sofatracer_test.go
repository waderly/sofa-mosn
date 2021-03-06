/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package trace

import (
	"runtime"
	"testing"
	"time"
)

func init() {
	tracer := CreateTracer("SOFATracer")
	SetTracer(tracer)
}

func TestSofaTracerStartFinish(t *testing.T) {
	span := Tracer().Start(time.Now())
	span.SetTag(TRACE_ID, IdGen().GenerateTraceId())
	span.FinishSpan()
}

func TestSofaTracerPrintSpan(t *testing.T) {
	Tracer().PrintSpan(&SofaTracerSpan{})
}

func TestSofaTracerPrintIngressSpan(t *testing.T) {
	span := &SofaTracerSpan{}
	span.tags[DOWNSTEAM_HOST_ADDRESS] = "127.0.0.1:43"
	span.tags[SPAN_TYPE] = "ingress"
	Tracer().PrintSpan(span)
}

func TestSofaTracerPrintEgressSpan(t *testing.T) {
	span := &SofaTracerSpan{}
	span.tags[SPAN_TYPE] = "egress"
	Tracer().PrintSpan(span)
}

func BenchmarkSofaTracer(b *testing.B) {
	runtime.GOMAXPROCS(runtime.NumCPU())
	for n := 0; n < b.N; n++ {
		span := &SofaTracerSpan{}
		span.SetTag(TRACE_ID, "BenchmarkSofaTracer")
		span.SetTag(PARENT_SPAN_ID, "BenchmarkSofaTracer")
		span.SetTag(SERVICE_NAME, "BenchmarkSofaTracer")
		span.SetTag(METHOD_NAME, "BenchmarkSofaTracer")
		span.SetTag(PROTOCOL, "BenchmarkSofaTracer")
		span.SetTag(RESULT_STATUS, "BenchmarkSofaTracer")
		span.SetTag(REQUEST_SIZE, "BenchmarkSofaTracer")
		span.SetTag(RESPONSE_SIZE, "BenchmarkSofaTracer")
		span.SetTag(UPSTREAM_HOST_ADDRESS, "BenchmarkSofaTracer")
		span.SetTag(DOWNSTEAM_HOST_ADDRESS, "BenchmarkSofaTracer")
		span.SetTag(APP_NAME, "BenchmarkSofaTracer")
		span.SetTag(SPAN_TYPE, "BenchmarkSofaTracer")
		span.SetTag(BAGGAGE_DATA, "BenchmarkSofaTracer")
		span.SetTag(REQUEST_URL, "BenchmarkSofaTracer")

		span.SetTag(SPAN_TYPE, "ingress")

		Tracer().PrintSpan(span)
	}
}

func BenchmarkSofaTracerParallel(b *testing.B) {
	runtime.GOMAXPROCS(runtime.NumCPU())
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			span := &SofaTracerSpan{}
			span.SetTag(TRACE_ID, "BenchmarkSofaTracer")
			span.SetTag(PARENT_SPAN_ID, "BenchmarkSofaTracer")
			span.SetTag(SERVICE_NAME, "BenchmarkSofaTracer")
			span.SetTag(METHOD_NAME, "BenchmarkSofaTracer")
			span.SetTag(PROTOCOL, "BenchmarkSofaTracer")
			span.SetTag(RESULT_STATUS, "BenchmarkSofaTracer")
			span.SetTag(REQUEST_SIZE, "BenchmarkSofaTracer")
			span.SetTag(RESPONSE_SIZE, "BenchmarkSofaTracer")
			span.SetTag(UPSTREAM_HOST_ADDRESS, "BenchmarkSofaTracer")
			span.SetTag(DOWNSTEAM_HOST_ADDRESS, "BenchmarkSofaTracer")
			span.SetTag(APP_NAME, "BenchmarkSofaTracer")
			span.SetTag(SPAN_TYPE, "BenchmarkSofaTracer")
			span.SetTag(BAGGAGE_DATA, "BenchmarkSofaTracer")
			span.SetTag(REQUEST_URL, "BenchmarkSofaTracer")

			span.SetTag(SPAN_TYPE, "ingress")

			Tracer().PrintSpan(span)
		}
	})
}

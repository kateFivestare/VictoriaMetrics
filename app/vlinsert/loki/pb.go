// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: push_request.proto
// source: https://raw.githubusercontent.com/grafana/loki/main/pkg/push/push_request.proto
// Licensed under the Apache License, Version 2.0 (the "License");
// https://github.com/grafana/loki/blob/main/pkg/push/LICENSE

package loki

import (
	"fmt"
	"time"

	"github.com/VictoriaMetrics/easyproto"
)

var mp easyproto.MarshalerPool

// PushRequest represents Loki PushRequest
//
// See https://github.com/grafana/loki/blob/4220737a52da7ab6c9346b12d5a5d7bedbcd641d/pkg/push/push.proto#L14C1-L14C20
type PushRequest struct {
	Streams []Stream

	entriesBuf   []Entry
	labelPairBuf []LabelPair
}

func (pr *PushRequest) reset() {
	pr.Streams = pr.Streams[:0]

	pr.entriesBuf = pr.entriesBuf[:0]
	pr.labelPairBuf = pr.labelPairBuf[:0]
}

// UnmarshalProtobuf unmarshals pr from protobuf message at src.
//
// pr remains valid until src is modified.
func (pr *PushRequest) UnmarshalProtobuf(src []byte) error {
	pr.reset()
	var err error
	pr.entriesBuf, pr.labelPairBuf, err = pr.unmarshalProtobuf(pr.entriesBuf, pr.labelPairBuf, src)
	return err
}

// MarshalProtobuf marshals r to protobuf message, appends it to dst and returns the result.
func (pr *PushRequest) MarshalProtobuf(dst []byte) []byte {
	m := mp.Get()
	pr.marshalProtobuf(m.MessageMarshaler())
	dst = m.Marshal(dst)
	mp.Put(m)
	return dst
}

func (pr *PushRequest) marshalProtobuf(mm *easyproto.MessageMarshaler) {
	for _, s := range pr.Streams {
		s.marshalProtobuf(mm.AppendMessage(1))
	}
}

func (pr *PushRequest) unmarshalProtobuf(entriesBuf []Entry, labelPairBuf []LabelPair, src []byte) ([]Entry, []LabelPair, error) {
	// message PushRequest {
	//   repeated Stream streams = 1;
	// }
	var err error
	var fc easyproto.FieldContext
	for len(src) > 0 {
		src, err = fc.NextField(src)
		if err != nil {
			return entriesBuf, labelPairBuf, fmt.Errorf("cannot read next field in PushRequest: %w", err)
		}
		switch fc.FieldNum {
		case 1:
			data, ok := fc.MessageData()
			if !ok {
				return entriesBuf, labelPairBuf, fmt.Errorf("cannot read Stream data")
			}
			pr.Streams = append(pr.Streams, Stream{})
			s := &pr.Streams[len(pr.Streams)-1]
			entriesBuf, labelPairBuf, err = s.unmarshalProtobuf(entriesBuf, labelPairBuf, data)
			if err != nil {
				return entriesBuf, labelPairBuf, fmt.Errorf("cannot unmarshal Stream: %w", err)
			}
		}
	}
	return entriesBuf, labelPairBuf, nil
}

// Stream represents Loki stream.
//
// See https://github.com/grafana/loki/blob/4220737a52da7ab6c9346b12d5a5d7bedbcd641d/pkg/push/push.proto#L23
type Stream struct {
	Labels  string
	Entries []Entry
}

func (s *Stream) marshalProtobuf(mm *easyproto.MessageMarshaler) {
	mm.AppendString(1, s.Labels)
	for _, e := range s.Entries {
		e.marshalProtobuf(mm.AppendMessage(2))
	}
}

func (s *Stream) unmarshalProtobuf(entriesBuf []Entry, labelPairBuf []LabelPair, src []byte) ([]Entry, []LabelPair, error) {
	// message Stream {
	//   string labels = 1;
	//   repeated Entry entries = 2;
	// }
	var err error
	var fc easyproto.FieldContext
	entriesBufLen := len(entriesBuf)
	for len(src) > 0 {
		src, err = fc.NextField(src)
		if err != nil {
			return entriesBuf, labelPairBuf, fmt.Errorf("cannot read next field in Stream: %w", err)
		}
		switch fc.FieldNum {
		case 1:
			labels, ok := fc.String()
			if !ok {
				return entriesBuf, labelPairBuf, fmt.Errorf("cannot read labels")
			}
			s.Labels = labels
		case 2:
			data, ok := fc.MessageData()
			if !ok {
				return entriesBuf, labelPairBuf, fmt.Errorf("cannot read Entry data")
			}
			entriesBuf = append(entriesBuf, Entry{})
			e := &entriesBuf[len(entriesBuf)-1]
			labelPairBuf, err = e.unmarshalProtobuf(labelPairBuf, data)
			if err != nil {
				return entriesBuf, labelPairBuf, fmt.Errorf("cannot unmarshal Entry: %w", err)
			}
		}
	}
	s.Entries = entriesBuf[entriesBufLen:]
	return entriesBuf, labelPairBuf, nil
}

// Entry represents Loki entry.
//
// See https://github.com/grafana/loki/blob/4220737a52da7ab6c9346b12d5a5d7bedbcd641d/pkg/push/push.proto#L38
type Entry struct {
	Timestamp          time.Time
	Line               string
	StructuredMetadata []LabelPair
}

func (e *Entry) marshalProtobuf(mm *easyproto.MessageMarshaler) {
	marshalTime(mm, 1, e.Timestamp)
	mm.AppendString(2, e.Line)
	for _, lp := range e.StructuredMetadata {
		lp.marshalProtobuf(mm.AppendMessage(3))
	}
}

func (e *Entry) unmarshalProtobuf(labelPairBuf []LabelPair, src []byte) ([]LabelPair, error) {
	// message Entry {
	//   Timestamp timestamp = 1;
	//   string line = 2;
	//   repeated LabelPair structuredMetadata = 3;
	// }
	var err error
	var fc easyproto.FieldContext
	labelPairBufLen := len(labelPairBuf)
	for len(src) > 0 {
		src, err = fc.NextField(src)
		if err != nil {
			return labelPairBuf, fmt.Errorf("cannot read next field in Entry: %w", err)
		}
		switch fc.FieldNum {
		case 1:
			data, ok := fc.MessageData()
			if !ok {
				return labelPairBuf, fmt.Errorf("cannot read Timestamp data")
			}
			timestamp, err := unmarshalTime(data)
			if err != nil {
				return labelPairBuf, fmt.Errorf("cannot unmarshal Timestamp: %w", err)
			}
			e.Timestamp = timestamp
		case 2:
			line, ok := fc.String()
			if !ok {
				return labelPairBuf, fmt.Errorf("cannot read Line")
			}
			e.Line = line
		case 3:
			data, ok := fc.MessageData()
			if !ok {
				return labelPairBuf, fmt.Errorf("cannot read StructuredMetadata")
			}
			labelPairBuf = append(labelPairBuf, LabelPair{})
			lp := &labelPairBuf[len(labelPairBuf)-1]
			if err := lp.unmarshalProtobuf(data); err != nil {
				return labelPairBuf, fmt.Errorf("cannot unmarshal StructuredMetadata: %w", err)
			}
		}
	}
	e.StructuredMetadata = labelPairBuf[labelPairBufLen:]
	return labelPairBuf, nil
}

// LabelPair represents Loki label pair.
//
// See https://github.com/grafana/loki/blob/4220737a52da7ab6c9346b12d5a5d7bedbcd641d/pkg/push/push.proto#L33
type LabelPair struct {
	Name  string
	Value string
}

func (lp *LabelPair) marshalProtobuf(mm *easyproto.MessageMarshaler) {
	mm.AppendString(1, lp.Name)
	mm.AppendString(2, lp.Value)
}

func (lp *LabelPair) unmarshalProtobuf(src []byte) (err error) {
	// message LabelPair {
	//   string name = 1;
	//   string value = 2;
	// }
	var fc easyproto.FieldContext
	for len(src) > 0 {
		src, err = fc.NextField(src)
		if err != nil {
			return fmt.Errorf("cannot read next field in LabelPair: %w", err)
		}
		switch fc.FieldNum {
		case 1:
			name, ok := fc.String()
			if !ok {
				return fmt.Errorf("cannot read name")
			}
			lp.Name = name
		case 2:
			value, ok := fc.String()
			if !ok {
				return fmt.Errorf("cannot unmarshal value")
			}
			lp.Value = value
		}
	}
	return nil
}

func marshalTime(mm *easyproto.MessageMarshaler, fieldNum uint32, timestamp time.Time) {
	nsecs := timestamp.UnixNano()
	ts := Timestamp{
		Seconds: nsecs / 1e9,
		Nanos:   int32(nsecs % 1e9),
	}
	ts.marshalProtobuf(mm.AppendMessage(fieldNum))
}

func unmarshalTime(src []byte) (time.Time, error) {
	var ts Timestamp
	if err := ts.unmarshalProtobuf(src); err != nil {
		return time.Time{}, err
	}
	timestamp := time.Unix(ts.Seconds, int64(ts.Nanos)).UTC()
	return timestamp, nil
}

// Timestamp is protobuf well-known timestamp type.
type Timestamp struct {
	Seconds int64
	Nanos   int32
}

func (ts *Timestamp) marshalProtobuf(mm *easyproto.MessageMarshaler) {
	mm.AppendInt64(1, ts.Seconds)
	mm.AppendInt32(2, ts.Nanos)
}

func (ts *Timestamp) unmarshalProtobuf(src []byte) (err error) {
	// message Timestamp {
	//   int64 seconds = 1;
	//   int32 nanos = 2;
	// }
	var fc easyproto.FieldContext
	for len(src) > 0 {
		src, err = fc.NextField(src)
		if err != nil {
			return fmt.Errorf("cannot read next field in Timestamp: %w", err)
		}
		switch fc.FieldNum {
		case 1:
			seconds, ok := fc.Int64()
			if !ok {
				return fmt.Errorf("cannot read Seconds")
			}
			ts.Seconds = seconds
		case 2:
			nanos, ok := fc.Int32()
			if !ok {
				return fmt.Errorf("cannot read Nanos")
			}
			ts.Nanos = nanos
		}
	}
	return nil
}
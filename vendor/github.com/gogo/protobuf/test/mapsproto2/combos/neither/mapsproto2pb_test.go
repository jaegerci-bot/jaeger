// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: combos/neither/mapsproto2.proto

/*
Package proto2_maps is a generated protocol buffer package.

It is generated from these files:
	combos/neither/mapsproto2.proto

It has these top-level messages:
	FloatingPoint
	CustomMap
	AllMaps
	AllMapsOrdered
*/
package proto2_maps

import testing "testing"
import rand "math/rand"
import time "time"
import proto "github.com/gogo/protobuf/proto"
import jsonpb "github.com/gogo/protobuf/jsonpb"
import fmt "fmt"
import parser "go/parser"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func TestFloatingPointProto(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := rand.New(rand.NewSource(seed))
	p := NewPopulatedFloatingPoint(popr, false)
	dAtA, err := proto.Marshal(p)
	if err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	msg := &FloatingPoint{}
	if err := proto.Unmarshal(dAtA, msg); err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	littlefuzz := make([]byte, len(dAtA))
	copy(littlefuzz, dAtA)
	for i := range dAtA {
		dAtA[i] = byte(popr.Intn(256))
	}
	if err := p.VerboseEqual(msg); err != nil {
		t.Fatalf("seed = %d, %#v !VerboseProto %#v, since %v", seed, msg, p, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("seed = %d, %#v !Proto %#v", seed, msg, p)
	}
	if len(littlefuzz) > 0 {
		fuzzamount := 100
		for i := 0; i < fuzzamount; i++ {
			littlefuzz[popr.Intn(len(littlefuzz))] = byte(popr.Intn(256))
			littlefuzz = append(littlefuzz, byte(popr.Intn(256)))
		}
		// shouldn't panic
		_ = proto.Unmarshal(littlefuzz, msg)
	}
}

func BenchmarkFloatingPointProtoMarshal(b *testing.B) {
	popr := rand.New(rand.NewSource(616))
	total := 0
	pops := make([]*FloatingPoint, 10000)
	for i := 0; i < 10000; i++ {
		pops[i] = NewPopulatedFloatingPoint(popr, false)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		dAtA, err := proto.Marshal(pops[i%10000])
		if err != nil {
			panic(err)
		}
		total += len(dAtA)
	}
	b.SetBytes(int64(total / b.N))
}

func BenchmarkFloatingPointProtoUnmarshal(b *testing.B) {
	popr := rand.New(rand.NewSource(616))
	total := 0
	datas := make([][]byte, 10000)
	for i := 0; i < 10000; i++ {
		dAtA, err := proto.Marshal(NewPopulatedFloatingPoint(popr, false))
		if err != nil {
			panic(err)
		}
		datas[i] = dAtA
	}
	msg := &FloatingPoint{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		total += len(datas[i%10000])
		if err := proto.Unmarshal(datas[i%10000], msg); err != nil {
			panic(err)
		}
	}
	b.SetBytes(int64(total / b.N))
}

func TestCustomMapProto(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := rand.New(rand.NewSource(seed))
	p := NewPopulatedCustomMap(popr, false)
	dAtA, err := proto.Marshal(p)
	if err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	msg := &CustomMap{}
	if err := proto.Unmarshal(dAtA, msg); err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	littlefuzz := make([]byte, len(dAtA))
	copy(littlefuzz, dAtA)
	for i := range dAtA {
		dAtA[i] = byte(popr.Intn(256))
	}
	if err := p.VerboseEqual(msg); err != nil {
		t.Fatalf("seed = %d, %#v !VerboseProto %#v, since %v", seed, msg, p, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("seed = %d, %#v !Proto %#v", seed, msg, p)
	}
	if len(littlefuzz) > 0 {
		fuzzamount := 100
		for i := 0; i < fuzzamount; i++ {
			littlefuzz[popr.Intn(len(littlefuzz))] = byte(popr.Intn(256))
			littlefuzz = append(littlefuzz, byte(popr.Intn(256)))
		}
		// shouldn't panic
		_ = proto.Unmarshal(littlefuzz, msg)
	}
}

func BenchmarkCustomMapProtoMarshal(b *testing.B) {
	popr := rand.New(rand.NewSource(616))
	total := 0
	pops := make([]*CustomMap, 10000)
	for i := 0; i < 10000; i++ {
		pops[i] = NewPopulatedCustomMap(popr, false)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		dAtA, err := proto.Marshal(pops[i%10000])
		if err != nil {
			panic(err)
		}
		total += len(dAtA)
	}
	b.SetBytes(int64(total / b.N))
}

func BenchmarkCustomMapProtoUnmarshal(b *testing.B) {
	popr := rand.New(rand.NewSource(616))
	total := 0
	datas := make([][]byte, 10000)
	for i := 0; i < 10000; i++ {
		dAtA, err := proto.Marshal(NewPopulatedCustomMap(popr, false))
		if err != nil {
			panic(err)
		}
		datas[i] = dAtA
	}
	msg := &CustomMap{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		total += len(datas[i%10000])
		if err := proto.Unmarshal(datas[i%10000], msg); err != nil {
			panic(err)
		}
	}
	b.SetBytes(int64(total / b.N))
}

func TestAllMapsProto(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := rand.New(rand.NewSource(seed))
	p := NewPopulatedAllMaps(popr, false)
	dAtA, err := proto.Marshal(p)
	if err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	msg := &AllMaps{}
	if err := proto.Unmarshal(dAtA, msg); err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	littlefuzz := make([]byte, len(dAtA))
	copy(littlefuzz, dAtA)
	for i := range dAtA {
		dAtA[i] = byte(popr.Intn(256))
	}
	if err := p.VerboseEqual(msg); err != nil {
		t.Fatalf("seed = %d, %#v !VerboseProto %#v, since %v", seed, msg, p, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("seed = %d, %#v !Proto %#v", seed, msg, p)
	}
	if len(littlefuzz) > 0 {
		fuzzamount := 100
		for i := 0; i < fuzzamount; i++ {
			littlefuzz[popr.Intn(len(littlefuzz))] = byte(popr.Intn(256))
			littlefuzz = append(littlefuzz, byte(popr.Intn(256)))
		}
		// shouldn't panic
		_ = proto.Unmarshal(littlefuzz, msg)
	}
}

func BenchmarkAllMapsProtoMarshal(b *testing.B) {
	popr := rand.New(rand.NewSource(616))
	total := 0
	pops := make([]*AllMaps, 10000)
	for i := 0; i < 10000; i++ {
		pops[i] = NewPopulatedAllMaps(popr, false)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		dAtA, err := proto.Marshal(pops[i%10000])
		if err != nil {
			panic(err)
		}
		total += len(dAtA)
	}
	b.SetBytes(int64(total / b.N))
}

func BenchmarkAllMapsProtoUnmarshal(b *testing.B) {
	popr := rand.New(rand.NewSource(616))
	total := 0
	datas := make([][]byte, 10000)
	for i := 0; i < 10000; i++ {
		dAtA, err := proto.Marshal(NewPopulatedAllMaps(popr, false))
		if err != nil {
			panic(err)
		}
		datas[i] = dAtA
	}
	msg := &AllMaps{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		total += len(datas[i%10000])
		if err := proto.Unmarshal(datas[i%10000], msg); err != nil {
			panic(err)
		}
	}
	b.SetBytes(int64(total / b.N))
}

func TestAllMapsOrderedProto(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := rand.New(rand.NewSource(seed))
	p := NewPopulatedAllMapsOrdered(popr, false)
	dAtA, err := proto.Marshal(p)
	if err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	msg := &AllMapsOrdered{}
	if err := proto.Unmarshal(dAtA, msg); err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	littlefuzz := make([]byte, len(dAtA))
	copy(littlefuzz, dAtA)
	for i := range dAtA {
		dAtA[i] = byte(popr.Intn(256))
	}
	if err := p.VerboseEqual(msg); err != nil {
		t.Fatalf("seed = %d, %#v !VerboseProto %#v, since %v", seed, msg, p, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("seed = %d, %#v !Proto %#v", seed, msg, p)
	}
	if len(littlefuzz) > 0 {
		fuzzamount := 100
		for i := 0; i < fuzzamount; i++ {
			littlefuzz[popr.Intn(len(littlefuzz))] = byte(popr.Intn(256))
			littlefuzz = append(littlefuzz, byte(popr.Intn(256)))
		}
		// shouldn't panic
		_ = proto.Unmarshal(littlefuzz, msg)
	}
}

func BenchmarkAllMapsOrderedProtoMarshal(b *testing.B) {
	popr := rand.New(rand.NewSource(616))
	total := 0
	pops := make([]*AllMapsOrdered, 10000)
	for i := 0; i < 10000; i++ {
		pops[i] = NewPopulatedAllMapsOrdered(popr, false)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		dAtA, err := proto.Marshal(pops[i%10000])
		if err != nil {
			panic(err)
		}
		total += len(dAtA)
	}
	b.SetBytes(int64(total / b.N))
}

func BenchmarkAllMapsOrderedProtoUnmarshal(b *testing.B) {
	popr := rand.New(rand.NewSource(616))
	total := 0
	datas := make([][]byte, 10000)
	for i := 0; i < 10000; i++ {
		dAtA, err := proto.Marshal(NewPopulatedAllMapsOrdered(popr, false))
		if err != nil {
			panic(err)
		}
		datas[i] = dAtA
	}
	msg := &AllMapsOrdered{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		total += len(datas[i%10000])
		if err := proto.Unmarshal(datas[i%10000], msg); err != nil {
			panic(err)
		}
	}
	b.SetBytes(int64(total / b.N))
}

func TestFloatingPointJSON(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := rand.New(rand.NewSource(seed))
	p := NewPopulatedFloatingPoint(popr, true)
	marshaler := jsonpb.Marshaler{}
	jsondata, err := marshaler.MarshalToString(p)
	if err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	msg := &FloatingPoint{}
	err = jsonpb.UnmarshalString(jsondata, msg)
	if err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	if err := p.VerboseEqual(msg); err != nil {
		t.Fatalf("seed = %d, %#v !VerboseProto %#v, since %v", seed, msg, p, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("seed = %d, %#v !Json Equal %#v", seed, msg, p)
	}
}
func TestCustomMapJSON(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := rand.New(rand.NewSource(seed))
	p := NewPopulatedCustomMap(popr, true)
	marshaler := jsonpb.Marshaler{}
	jsondata, err := marshaler.MarshalToString(p)
	if err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	msg := &CustomMap{}
	err = jsonpb.UnmarshalString(jsondata, msg)
	if err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	if err := p.VerboseEqual(msg); err != nil {
		t.Fatalf("seed = %d, %#v !VerboseProto %#v, since %v", seed, msg, p, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("seed = %d, %#v !Json Equal %#v", seed, msg, p)
	}
}
func TestAllMapsJSON(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := rand.New(rand.NewSource(seed))
	p := NewPopulatedAllMaps(popr, true)
	marshaler := jsonpb.Marshaler{}
	jsondata, err := marshaler.MarshalToString(p)
	if err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	msg := &AllMaps{}
	err = jsonpb.UnmarshalString(jsondata, msg)
	if err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	if err := p.VerboseEqual(msg); err != nil {
		t.Fatalf("seed = %d, %#v !VerboseProto %#v, since %v", seed, msg, p, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("seed = %d, %#v !Json Equal %#v", seed, msg, p)
	}
}
func TestAllMapsOrderedJSON(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := rand.New(rand.NewSource(seed))
	p := NewPopulatedAllMapsOrdered(popr, true)
	marshaler := jsonpb.Marshaler{}
	jsondata, err := marshaler.MarshalToString(p)
	if err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	msg := &AllMapsOrdered{}
	err = jsonpb.UnmarshalString(jsondata, msg)
	if err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	if err := p.VerboseEqual(msg); err != nil {
		t.Fatalf("seed = %d, %#v !VerboseProto %#v, since %v", seed, msg, p, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("seed = %d, %#v !Json Equal %#v", seed, msg, p)
	}
}
func TestFloatingPointProtoText(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := rand.New(rand.NewSource(seed))
	p := NewPopulatedFloatingPoint(popr, true)
	dAtA := proto.MarshalTextString(p)
	msg := &FloatingPoint{}
	if err := proto.UnmarshalText(dAtA, msg); err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	if err := p.VerboseEqual(msg); err != nil {
		t.Fatalf("seed = %d, %#v !VerboseProto %#v, since %v", seed, msg, p, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("seed = %d, %#v !Proto %#v", seed, msg, p)
	}
}

func TestFloatingPointProtoCompactText(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := rand.New(rand.NewSource(seed))
	p := NewPopulatedFloatingPoint(popr, true)
	dAtA := proto.CompactTextString(p)
	msg := &FloatingPoint{}
	if err := proto.UnmarshalText(dAtA, msg); err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	if err := p.VerboseEqual(msg); err != nil {
		t.Fatalf("seed = %d, %#v !VerboseProto %#v, since %v", seed, msg, p, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("seed = %d, %#v !Proto %#v", seed, msg, p)
	}
}

func TestCustomMapProtoText(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := rand.New(rand.NewSource(seed))
	p := NewPopulatedCustomMap(popr, true)
	dAtA := proto.MarshalTextString(p)
	msg := &CustomMap{}
	if err := proto.UnmarshalText(dAtA, msg); err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	if err := p.VerboseEqual(msg); err != nil {
		t.Fatalf("seed = %d, %#v !VerboseProto %#v, since %v", seed, msg, p, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("seed = %d, %#v !Proto %#v", seed, msg, p)
	}
}

func TestCustomMapProtoCompactText(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := rand.New(rand.NewSource(seed))
	p := NewPopulatedCustomMap(popr, true)
	dAtA := proto.CompactTextString(p)
	msg := &CustomMap{}
	if err := proto.UnmarshalText(dAtA, msg); err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	if err := p.VerboseEqual(msg); err != nil {
		t.Fatalf("seed = %d, %#v !VerboseProto %#v, since %v", seed, msg, p, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("seed = %d, %#v !Proto %#v", seed, msg, p)
	}
}

func TestAllMapsProtoText(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := rand.New(rand.NewSource(seed))
	p := NewPopulatedAllMaps(popr, true)
	dAtA := proto.MarshalTextString(p)
	msg := &AllMaps{}
	if err := proto.UnmarshalText(dAtA, msg); err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	if err := p.VerboseEqual(msg); err != nil {
		t.Fatalf("seed = %d, %#v !VerboseProto %#v, since %v", seed, msg, p, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("seed = %d, %#v !Proto %#v", seed, msg, p)
	}
}

func TestAllMapsProtoCompactText(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := rand.New(rand.NewSource(seed))
	p := NewPopulatedAllMaps(popr, true)
	dAtA := proto.CompactTextString(p)
	msg := &AllMaps{}
	if err := proto.UnmarshalText(dAtA, msg); err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	if err := p.VerboseEqual(msg); err != nil {
		t.Fatalf("seed = %d, %#v !VerboseProto %#v, since %v", seed, msg, p, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("seed = %d, %#v !Proto %#v", seed, msg, p)
	}
}

func TestAllMapsOrderedProtoText(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := rand.New(rand.NewSource(seed))
	p := NewPopulatedAllMapsOrdered(popr, true)
	dAtA := proto.MarshalTextString(p)
	msg := &AllMapsOrdered{}
	if err := proto.UnmarshalText(dAtA, msg); err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	if err := p.VerboseEqual(msg); err != nil {
		t.Fatalf("seed = %d, %#v !VerboseProto %#v, since %v", seed, msg, p, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("seed = %d, %#v !Proto %#v", seed, msg, p)
	}
}

func TestAllMapsOrderedProtoCompactText(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := rand.New(rand.NewSource(seed))
	p := NewPopulatedAllMapsOrdered(popr, true)
	dAtA := proto.CompactTextString(p)
	msg := &AllMapsOrdered{}
	if err := proto.UnmarshalText(dAtA, msg); err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	if err := p.VerboseEqual(msg); err != nil {
		t.Fatalf("seed = %d, %#v !VerboseProto %#v, since %v", seed, msg, p, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("seed = %d, %#v !Proto %#v", seed, msg, p)
	}
}

func TestMapsproto2Description(t *testing.T) {
	Mapsproto2Description()
}
func TestFloatingPointVerboseEqual(t *testing.T) {
	popr := rand.New(rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedFloatingPoint(popr, false)
	dAtA, err := proto.Marshal(p)
	if err != nil {
		panic(err)
	}
	msg := &FloatingPoint{}
	if err := proto.Unmarshal(dAtA, msg); err != nil {
		panic(err)
	}
	if err := p.VerboseEqual(msg); err != nil {
		t.Fatalf("%#v !VerboseEqual %#v, since %v", msg, p, err)
	}
}
func TestCustomMapVerboseEqual(t *testing.T) {
	popr := rand.New(rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedCustomMap(popr, false)
	dAtA, err := proto.Marshal(p)
	if err != nil {
		panic(err)
	}
	msg := &CustomMap{}
	if err := proto.Unmarshal(dAtA, msg); err != nil {
		panic(err)
	}
	if err := p.VerboseEqual(msg); err != nil {
		t.Fatalf("%#v !VerboseEqual %#v, since %v", msg, p, err)
	}
}
func TestAllMapsVerboseEqual(t *testing.T) {
	popr := rand.New(rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedAllMaps(popr, false)
	dAtA, err := proto.Marshal(p)
	if err != nil {
		panic(err)
	}
	msg := &AllMaps{}
	if err := proto.Unmarshal(dAtA, msg); err != nil {
		panic(err)
	}
	if err := p.VerboseEqual(msg); err != nil {
		t.Fatalf("%#v !VerboseEqual %#v, since %v", msg, p, err)
	}
}
func TestAllMapsOrderedVerboseEqual(t *testing.T) {
	popr := rand.New(rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedAllMapsOrdered(popr, false)
	dAtA, err := proto.Marshal(p)
	if err != nil {
		panic(err)
	}
	msg := &AllMapsOrdered{}
	if err := proto.Unmarshal(dAtA, msg); err != nil {
		panic(err)
	}
	if err := p.VerboseEqual(msg); err != nil {
		t.Fatalf("%#v !VerboseEqual %#v, since %v", msg, p, err)
	}
}
func TestFloatingPointFace(t *testing.T) {
	popr := rand.New(rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedFloatingPoint(popr, true)
	msg := p.TestProto()
	if !p.Equal(msg) {
		t.Fatalf("%#v !Face Equal %#v", msg, p)
	}
}
func TestCustomMapFace(t *testing.T) {
	popr := rand.New(rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedCustomMap(popr, true)
	msg := p.TestProto()
	if !p.Equal(msg) {
		t.Fatalf("%#v !Face Equal %#v", msg, p)
	}
}
func TestAllMapsFace(t *testing.T) {
	popr := rand.New(rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedAllMaps(popr, true)
	msg := p.TestProto()
	if !p.Equal(msg) {
		t.Fatalf("%#v !Face Equal %#v", msg, p)
	}
}
func TestAllMapsOrderedFace(t *testing.T) {
	popr := rand.New(rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedAllMapsOrdered(popr, true)
	msg := p.TestProto()
	if !p.Equal(msg) {
		t.Fatalf("%#v !Face Equal %#v", msg, p)
	}
}
func TestFloatingPointGoString(t *testing.T) {
	popr := rand.New(rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedFloatingPoint(popr, false)
	s1 := p.GoString()
	s2 := fmt.Sprintf("%#v", p)
	if s1 != s2 {
		t.Fatalf("GoString want %v got %v", s1, s2)
	}
	_, err := parser.ParseExpr(s1)
	if err != nil {
		t.Fatal(err)
	}
}
func TestCustomMapGoString(t *testing.T) {
	popr := rand.New(rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedCustomMap(popr, false)
	s1 := p.GoString()
	s2 := fmt.Sprintf("%#v", p)
	if s1 != s2 {
		t.Fatalf("GoString want %v got %v", s1, s2)
	}
	_, err := parser.ParseExpr(s1)
	if err != nil {
		t.Fatal(err)
	}
}
func TestAllMapsGoString(t *testing.T) {
	popr := rand.New(rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedAllMaps(popr, false)
	s1 := p.GoString()
	s2 := fmt.Sprintf("%#v", p)
	if s1 != s2 {
		t.Fatalf("GoString want %v got %v", s1, s2)
	}
	_, err := parser.ParseExpr(s1)
	if err != nil {
		t.Fatal(err)
	}
}
func TestAllMapsOrderedGoString(t *testing.T) {
	popr := rand.New(rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedAllMapsOrdered(popr, false)
	s1 := p.GoString()
	s2 := fmt.Sprintf("%#v", p)
	if s1 != s2 {
		t.Fatalf("GoString want %v got %v", s1, s2)
	}
	_, err := parser.ParseExpr(s1)
	if err != nil {
		t.Fatal(err)
	}
}
func TestFloatingPointSize(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := rand.New(rand.NewSource(seed))
	p := NewPopulatedFloatingPoint(popr, true)
	size2 := proto.Size(p)
	dAtA, err := proto.Marshal(p)
	if err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	size := p.Size()
	if len(dAtA) != size {
		t.Errorf("seed = %d, size %v != marshalled size %v", seed, size, len(dAtA))
	}
	if size2 != size {
		t.Errorf("seed = %d, size %v != before marshal proto.Size %v", seed, size, size2)
	}
	size3 := proto.Size(p)
	if size3 != size {
		t.Errorf("seed = %d, size %v != after marshal proto.Size %v", seed, size, size3)
	}
}

func BenchmarkFloatingPointSize(b *testing.B) {
	popr := rand.New(rand.NewSource(616))
	total := 0
	pops := make([]*FloatingPoint, 1000)
	for i := 0; i < 1000; i++ {
		pops[i] = NewPopulatedFloatingPoint(popr, false)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		total += pops[i%1000].Size()
	}
	b.SetBytes(int64(total / b.N))
}

func TestCustomMapSize(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := rand.New(rand.NewSource(seed))
	p := NewPopulatedCustomMap(popr, true)
	size2 := proto.Size(p)
	dAtA, err := proto.Marshal(p)
	if err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	size := p.Size()
	if len(dAtA) != size {
		t.Errorf("seed = %d, size %v != marshalled size %v", seed, size, len(dAtA))
	}
	if size2 != size {
		t.Errorf("seed = %d, size %v != before marshal proto.Size %v", seed, size, size2)
	}
	size3 := proto.Size(p)
	if size3 != size {
		t.Errorf("seed = %d, size %v != after marshal proto.Size %v", seed, size, size3)
	}
}

func BenchmarkCustomMapSize(b *testing.B) {
	popr := rand.New(rand.NewSource(616))
	total := 0
	pops := make([]*CustomMap, 1000)
	for i := 0; i < 1000; i++ {
		pops[i] = NewPopulatedCustomMap(popr, false)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		total += pops[i%1000].Size()
	}
	b.SetBytes(int64(total / b.N))
}

func TestAllMapsSize(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := rand.New(rand.NewSource(seed))
	p := NewPopulatedAllMaps(popr, true)
	size2 := proto.Size(p)
	dAtA, err := proto.Marshal(p)
	if err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	size := p.Size()
	if len(dAtA) != size {
		t.Errorf("seed = %d, size %v != marshalled size %v", seed, size, len(dAtA))
	}
	if size2 != size {
		t.Errorf("seed = %d, size %v != before marshal proto.Size %v", seed, size, size2)
	}
	size3 := proto.Size(p)
	if size3 != size {
		t.Errorf("seed = %d, size %v != after marshal proto.Size %v", seed, size, size3)
	}
}

func BenchmarkAllMapsSize(b *testing.B) {
	popr := rand.New(rand.NewSource(616))
	total := 0
	pops := make([]*AllMaps, 1000)
	for i := 0; i < 1000; i++ {
		pops[i] = NewPopulatedAllMaps(popr, false)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		total += pops[i%1000].Size()
	}
	b.SetBytes(int64(total / b.N))
}

func TestAllMapsOrderedSize(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := rand.New(rand.NewSource(seed))
	p := NewPopulatedAllMapsOrdered(popr, true)
	size2 := proto.Size(p)
	dAtA, err := proto.Marshal(p)
	if err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	size := p.Size()
	if len(dAtA) != size {
		t.Errorf("seed = %d, size %v != marshalled size %v", seed, size, len(dAtA))
	}
	if size2 != size {
		t.Errorf("seed = %d, size %v != before marshal proto.Size %v", seed, size, size2)
	}
	size3 := proto.Size(p)
	if size3 != size {
		t.Errorf("seed = %d, size %v != after marshal proto.Size %v", seed, size, size3)
	}
}

func BenchmarkAllMapsOrderedSize(b *testing.B) {
	popr := rand.New(rand.NewSource(616))
	total := 0
	pops := make([]*AllMapsOrdered, 1000)
	for i := 0; i < 1000; i++ {
		pops[i] = NewPopulatedAllMapsOrdered(popr, false)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		total += pops[i%1000].Size()
	}
	b.SetBytes(int64(total / b.N))
}

func TestFloatingPointStringer(t *testing.T) {
	popr := rand.New(rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedFloatingPoint(popr, false)
	s1 := p.String()
	s2 := fmt.Sprintf("%v", p)
	if s1 != s2 {
		t.Fatalf("String want %v got %v", s1, s2)
	}
}
func TestCustomMapStringer(t *testing.T) {
	popr := rand.New(rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedCustomMap(popr, false)
	s1 := p.String()
	s2 := fmt.Sprintf("%v", p)
	if s1 != s2 {
		t.Fatalf("String want %v got %v", s1, s2)
	}
}
func TestAllMapsStringer(t *testing.T) {
	popr := rand.New(rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedAllMaps(popr, false)
	s1 := p.String()
	s2 := fmt.Sprintf("%v", p)
	if s1 != s2 {
		t.Fatalf("String want %v got %v", s1, s2)
	}
}
func TestAllMapsOrderedStringer(t *testing.T) {
	popr := rand.New(rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedAllMapsOrdered(popr, false)
	s1 := p.String()
	s2 := fmt.Sprintf("%v", p)
	if s1 != s2 {
		t.Fatalf("String want %v got %v", s1, s2)
	}
}

//These tests are generated by github.com/gogo/protobuf/plugin/testgen

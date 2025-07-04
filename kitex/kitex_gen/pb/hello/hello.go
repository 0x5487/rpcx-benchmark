// Code generated by Kitex v0.14.1. DO NOT EDIT.

package hello

import (
	"context"
	"errors"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	streaming "github.com/cloudwego/kitex/pkg/streaming"
	proto "github.com/cloudwego/prutal"
	pb "github.com/rpcxio/rpcx-benchmark/kitex/kitex_gen/pb"
)

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"Say": kitex.NewMethodInfo(
		sayHandler,
		newSayArgs,
		newSayResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
}

var (
	helloServiceInfo                = NewServiceInfo()
	helloServiceInfoForClient       = NewServiceInfoForClient()
	helloServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

// for server
func serviceInfo() *kitex.ServiceInfo {
	return helloServiceInfo
}

// for stream client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return helloServiceInfoForStreamClient
}

// for client
func serviceInfoForClient() *kitex.ServiceInfo {
	return helloServiceInfoForClient
}

// NewServiceInfo creates a new ServiceInfo containing all methods
func NewServiceInfo() *kitex.ServiceInfo {
	return newServiceInfo(false, true, true)
}

// NewServiceInfo creates a new ServiceInfo containing non-streaming methods
func NewServiceInfoForClient() *kitex.ServiceInfo {
	return newServiceInfo(false, false, true)
}
func NewServiceInfoForStreamClient() *kitex.ServiceInfo {
	return newServiceInfo(true, true, false)
}

func newServiceInfo(hasStreaming bool, keepStreamingMethods bool, keepNonStreamingMethods bool) *kitex.ServiceInfo {
	serviceName := "Hello"
	handlerType := (*pb.Hello)(nil)
	methods := map[string]kitex.MethodInfo{}
	for name, m := range serviceMethods {
		if m.IsStreaming() && !keepStreamingMethods {
			continue
		}
		if !m.IsStreaming() && !keepNonStreamingMethods {
			continue
		}
		methods[name] = m
	}
	extra := map[string]interface{}{
		"PackageName": "pb",
	}
	if hasStreaming {
		extra["streaming"] = hasStreaming
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Protobuf,
		KiteXGenVersion: "v0.14.1",
		Extra:           extra,
	}
	return svcInfo
}

func sayHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(pb.BenchmarkMessage)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(pb.Hello).Say(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *SayArgs:
		success, err := handler.(pb.Hello).Say(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*SayResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newSayArgs() interface{} {
	return &SayArgs{}
}

func newSayResult() interface{} {
	return &SayResult{}
}

type SayArgs struct {
	Req *pb.BenchmarkMessage
}

func (p *SayArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *SayArgs) Unmarshal(in []byte) error {
	msg := new(pb.BenchmarkMessage)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var SayArgs_Req_DEFAULT *pb.BenchmarkMessage

func (p *SayArgs) GetReq() *pb.BenchmarkMessage {
	if !p.IsSetReq() {
		return SayArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *SayArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *SayArgs) GetFirstArgument() interface{} {
	return p.Req
}

type SayResult struct {
	Success *pb.BenchmarkMessage
}

var SayResult_Success_DEFAULT *pb.BenchmarkMessage

func (p *SayResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *SayResult) Unmarshal(in []byte) error {
	msg := new(pb.BenchmarkMessage)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *SayResult) GetSuccess() *pb.BenchmarkMessage {
	if !p.IsSetSuccess() {
		return SayResult_Success_DEFAULT
	}
	return p.Success
}

func (p *SayResult) SetSuccess(x interface{}) {
	p.Success = x.(*pb.BenchmarkMessage)
}

func (p *SayResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *SayResult) GetResult() interface{} {
	return p.Success
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) Say(ctx context.Context, Req *pb.BenchmarkMessage) (r *pb.BenchmarkMessage, err error) {
	var _args SayArgs
	_args.Req = Req
	var _result SayResult
	if err = p.c.Call(ctx, "Say", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

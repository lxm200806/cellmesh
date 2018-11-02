// Auto generated by github.com/davyxu/cellmesh/protogen
// DO NOT EDIT!

package proto

import (
	"github.com/davyxu/cellnet"
	"github.com/davyxu/cellnet/codec"
	_ "github.com/davyxu/cellnet/codec/gogopb"
	_ "github.com/davyxu/cellnet/codec/protoplus"
	"reflect"
)

// Make compiler import happy
var (
	_ cellnet.Event
	_ codec.CodecRecycler
	_ reflect.Type
)

// agent
var (
	Handle_Agent_BindBackendACK = func(ev cellnet.Event) { panic("'BindBackendACK' not handled") }
	Handle_Agent_CloseClientACK = func(ev cellnet.Event) { panic("'CloseClientACK' not handled") }
	Handle_Agent_Default        func(ev cellnet.Event)
)

// game
var (
	Handle_Game_ChatREQ   = func(ev cellnet.Event) { panic("'ChatREQ' not handled") }
	Handle_Game_VerifyREQ = func(ev cellnet.Event) { panic("'VerifyREQ' not handled") }
	Handle_Game_Default   func(ev cellnet.Event)
)

// hub
var (
	Handle_Hub_SubscribeChannelREQ = func(ev cellnet.Event) { panic("'SubscribeChannelREQ' not handled") }
	Handle_Hub_Default             func(ev cellnet.Event)
)

// login
var (
	Handle_Login_LoginREQ     = func(ev cellnet.Event) { panic("'LoginREQ' not handled") }
	Handle_Login_SvcStatusACK = func(ev cellnet.Event) { panic("'SvcStatusACK' not handled") }
	Handle_Login_Default      func(ev cellnet.Event)
)

// match
var (
	Handle_Match_SvcStatusACK = func(ev cellnet.Event) { panic("'SvcStatusACK' not handled") }
	Handle_Match_Default      func(ev cellnet.Event)
)

func GetMessageHandler(svcName string) cellnet.EventCallback {

	switch svcName {
	case "agent":
		return func(ev cellnet.Event) {
			switch ev.Message().(type) {
			case *BindBackendACK:
				Handle_Agent_BindBackendACK(ev)
			case *CloseClientACK:
				Handle_Agent_CloseClientACK(ev)
			default:
				if Handle_Agent_Default != nil {
					Handle_Agent_Default(ev)
				}
			}
		}
	case "game":
		return func(ev cellnet.Event) {
			switch ev.Message().(type) {
			case *ChatREQ:
				Handle_Game_ChatREQ(ev)
			case *VerifyREQ:
				Handle_Game_VerifyREQ(ev)
			default:
				if Handle_Game_Default != nil {
					Handle_Game_Default(ev)
				}
			}
		}
	case "hub":
		return func(ev cellnet.Event) {
			switch ev.Message().(type) {
			case *SubscribeChannelREQ:
				Handle_Hub_SubscribeChannelREQ(ev)
			default:
				if Handle_Hub_Default != nil {
					Handle_Hub_Default(ev)
				}
			}
		}
	case "login":
		return func(ev cellnet.Event) {
			switch ev.Message().(type) {
			case *LoginREQ:
				Handle_Login_LoginREQ(ev)
			case *SvcStatusACK:
				Handle_Login_SvcStatusACK(ev)
			default:
				if Handle_Login_Default != nil {
					Handle_Login_Default(ev)
				}
			}
		}
	case "match":
		return func(ev cellnet.Event) {
			switch ev.Message().(type) {
			case *SvcStatusACK:
				Handle_Match_SvcStatusACK(ev)
			default:
				if Handle_Match_Default != nil {
					Handle_Match_Default(ev)
				}
			}
		}
	}

	return nil
}

func init() {

	cellnet.RegisterMessageMeta(&cellnet.MessageMeta{
		Codec: codec.MustGetCodec("gogopb"),
		Type:  reflect.TypeOf((*PingACK)(nil)).Elem(),
		ID:    16241,
	})
	cellnet.RegisterMessageMeta(&cellnet.MessageMeta{
		Codec: codec.MustGetCodec("gogopb"),
		Type:  reflect.TypeOf((*LoginREQ)(nil)).Elem(),
		ID:    18837,
	})
	cellnet.RegisterMessageMeta(&cellnet.MessageMeta{
		Codec: codec.MustGetCodec("gogopb"),
		Type:  reflect.TypeOf((*LoginACK)(nil)).Elem(),
		ID:    46204,
	})
	cellnet.RegisterMessageMeta(&cellnet.MessageMeta{
		Codec: codec.MustGetCodec("gogopb"),
		Type:  reflect.TypeOf((*VerifyREQ)(nil)).Elem(),
		ID:    13457,
	})
	cellnet.RegisterMessageMeta(&cellnet.MessageMeta{
		Codec: codec.MustGetCodec("gogopb"),
		Type:  reflect.TypeOf((*VerifyACK)(nil)).Elem(),
		ID:    40824,
	})
	cellnet.RegisterMessageMeta(&cellnet.MessageMeta{
		Codec: codec.MustGetCodec("gogopb"),
		Type:  reflect.TypeOf((*ChatREQ)(nil)).Elem(),
		ID:    29052,
	})
	cellnet.RegisterMessageMeta(&cellnet.MessageMeta{
		Codec: codec.MustGetCodec("gogopb"),
		Type:  reflect.TypeOf((*ChatACK)(nil)).Elem(),
		ID:    56419,
	})
	cellnet.RegisterMessageMeta(&cellnet.MessageMeta{
		Codec: codec.MustGetCodec("protoplus"),
		Type:  reflect.TypeOf((*BindBackendACK)(nil)).Elem(),
		ID:    5768,
	})
	cellnet.RegisterMessageMeta(&cellnet.MessageMeta{
		Codec: codec.MustGetCodec("protoplus"),
		Type:  reflect.TypeOf((*CloseClientACK)(nil)).Elem(),
		ID:    58040,
	})
	cellnet.RegisterMessageMeta(&cellnet.MessageMeta{
		Codec: codec.MustGetCodec("protoplus"),
		Type:  reflect.TypeOf((*ClientClosedACK)(nil)).Elem(),
		ID:    50844,
	})
	cellnet.RegisterMessageMeta(&cellnet.MessageMeta{
		Codec: codec.MustGetCodec("protoplus"),
		Type:  reflect.TypeOf((*SubscribeChannelREQ)(nil)).Elem(),
		ID:    27927,
	})
	cellnet.RegisterMessageMeta(&cellnet.MessageMeta{
		Codec: codec.MustGetCodec("protoplus"),
		Type:  reflect.TypeOf((*SubscribeChannelACK)(nil)).Elem(),
		ID:    55294,
	})
	cellnet.RegisterMessageMeta(&cellnet.MessageMeta{
		Codec: codec.MustGetCodec("protoplus"),
		Type:  reflect.TypeOf((*SvcStatusACK)(nil)).Elem(),
		ID:    50227,
	})
}

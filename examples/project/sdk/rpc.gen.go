// Code generated by 'micro gen' command.
// DO NOT EDIT!

package sdk

import (
	"fmt"

	tp "github.com/henrylee2cn/teleport"
	"github.com/xiaoenai/tp-micro/clientele"
)

var _ = fmt.Sprintf

// Stat handler
func Stat(ctx clientele.Ctx, arg *StatArg, setting ...tp.MessageSetting) *tp.Rerror {
	setting = append(setting, tp.WithQuery("ts", fmt.Sprintf("%v", arg.Ts)))
	return clientele.DynamicPush(ctx, "/project/stat", arg, setting...)
}

// Home handler
func Home(ctx clientele.Ctx, arg *EmptyStruct, setting ...tp.MessageSetting) (*HomeResult, *tp.Rerror) {
	result := new(HomeResult)
	rerr := clientele.DynamicCall(ctx, "/project/home", arg, result, setting...).Rerror()
	return result, rerr
}

// Divide handler
func Math_Divide(ctx clientele.Ctx, arg *DivideArg, setting ...tp.MessageSetting) (*DivideResult, *tp.Rerror) {
	result := new(DivideResult)
	rerr := clientele.DynamicCall(ctx, "/project/math/divide", arg, result, setting...).Rerror()
	return result, rerr
}

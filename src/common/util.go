package common

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"github.com/dylanpeng/golib/coder"
	"github.com/dylanpeng/nbc-chat/common/config"
	"github.com/dylanpeng/nbc-chat/common/consts"
	"github.com/gin-gonic/gin"
	"image"
	"io"
	"runtime/debug"
	"strings"
)

func GetKey(prefix string, items ...interface{}) string {
	format := config.GetConfig().App.GetKeyPrefix() + prefix + strings.Repeat(":%v", len(items))
	return fmt.Sprintf(format, items...)
}

func ConvertStruct(a interface{}, b interface{}) (err error) {
	defer func() {
		if err != nil {
			Logger.Debugf("convert data failed | data: %s | error: %s", a, err)
		}
	}()

	data, err := json.Marshal(a)

	if err != nil {
		return
	}

	return json.Unmarshal(data, b)
}

func ConvertStructs(items ...fmt.Stringer) (err error) {
	for i := 0; i < len(items)-1; i += 2 {
		if err = ConvertStruct(items[i], items[i+1]); err != nil {
			return
		}
	}

	return
}

func CatchPanic() {
	if err := recover(); err != nil {
		Logger.Errorf("catch panic | %s\n%s", err, debug.Stack())
	}
}

func GetCtxCoder(ctx *gin.Context) coder.ICoder {
	name := ctx.GetString(consts.CtxCoderKey)

	if name == coder.EncodingJson {
		return coder.JsonCoder
	} else {
		return config.GetConfig().App.HttpCoder
	}
}

func SetCtxCoder(ctx *gin.Context, encoding string) {
	if encoding == coder.EncodingJson {
		ctx.Set(consts.CtxCoderKey, encoding)
	}
}

func GetTraceId(ctx context.Context) string {
	trace := ctx.Value(consts.CtxValueTraceId)
	traceId, ok := trace.(string)

	if ok {
		return traceId
	}

	return ""
}

func IsRGBAImage(reader io.Reader) (isRGBA bool, err error) {
	img, _, err := image.Decode(bufio.NewReader(reader))
	if err != nil {
		return false, err
	}

	// 判断是否是 RGBA 格式
	isRGBA = true
	bounds := img.Bounds()
	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			_, _, _, a := img.At(x, y).RGBA()
			if a != 65535 {
				isRGBA = false
				break
			}
		}
		if !isRGBA {
			break
		}
	}

	return
}

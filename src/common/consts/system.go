package consts

const (
	CtxCoderKey      = "context.value.request.context.coder"
	CtxValueRequest  = "context.value.request.value"
	CtxValueAuth     = "context.value.auth.user"
	CtxValueLanguage = "context.value.language"
	CtxValueTraceId  = "context.value.trace.id"
	CtxValueImage    = "context.value.image"
)

const (
	HeaderKeyAuth        = "Authorization"
	HeaderKeyUserId      = "UserId"
	HeaderKeyVersionCode = "Version_code"
	HeaderKeyVersionName = "Version_name"
	HeaderKeyTokenPrefix = "Bearer "
	HeaderKeyLanguage    = "Language"
	HeaderKeyPlatform    = "Platform"
	HeaderKeyDeviceId    = "Device_id"
	HeaderKeyPackageName = "pn"
	HeaderKeyTraceId     = "TraceId"
)

const (
	SystemFileSizeLimit = 1024 * 1024 * 4
)

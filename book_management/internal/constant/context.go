package constant

type wrapCtx string
type userCtx string

const (
	RequestIDKey    = "X-Request-Id"
	RequestIDPrefix = "RID"
	UserContext userCtx = "authorizedUser"

	TxCtx = wrapCtx("TxCtx")
)

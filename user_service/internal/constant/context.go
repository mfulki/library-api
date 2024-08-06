package constant

type wrapCtx string

const (
	UserContext = "authorizedUser"
	RequestIDKey    = "X-Request-Id"
	RequestIDPrefix = "RID"

	TxCtx = wrapCtx("TxCtx")
)

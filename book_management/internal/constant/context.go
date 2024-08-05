package constant

type wrapCtx string

const (
	RequestIDKey    = "X-Request-Id"
	RequestIDPrefix = "RID"

	TxCtx = wrapCtx("TxCtx")
)

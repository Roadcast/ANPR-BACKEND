package constant

type contextKey string

const UserCtxKey = contextKey("user")
const TokenContextKey = contextKey("token")
const BypassPrivacyKey = contextKey("bypass_privacy")

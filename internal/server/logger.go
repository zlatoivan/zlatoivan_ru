package server

//// GetReqID returns a request ID from the given context if one is present.
//// Returns the empty string if a request ID cannot be found.
//func GetReqID(ctx context.Context) string {
//	if ctx == nil {
//		return ""
//	}
//	if reqID, ok := ctx.Value(RequestIDKey).(string); ok {
//		return reqID
//	}
//	return ""
//}

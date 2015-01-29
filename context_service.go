package jiaoxue

type Context struct {
	users          *Users
	live_manager   *LiveManager
	record_manager *RecordManager
}

var jiao_context *Context

func GetContext() *Context {
	if jiao_context == nil {
		jiao_context = new(Context)
	}

	return jiao_context
}

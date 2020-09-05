package mobi

var options Options

//Options Struct
type Options struct {
	debug bool
}

//IsDebug return bool if is debug
func (o Options) IsDebug() bool {
	return o.debug
}

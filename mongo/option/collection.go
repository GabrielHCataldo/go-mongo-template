package option

type Drop struct {
	DisableAutoCloseSession bool
}

func NewDrop() Drop {
	return Drop{}
}

func (d Drop) SetDisableAutoCloseTransaction(b bool) Drop {
	d.DisableAutoCloseSession = b
	return d
}

func GetDropOptionByParams(opts []Drop) Drop {
	result := Drop{}
	for _, opt := range opts {
		if opt.DisableAutoCloseSession {
			result.DisableAutoCloseSession = opt.DisableAutoCloseSession
		}
	}
	return result
}

package option

type Drop struct {
	DisableAutoCloseTransaction bool
}

func NewDrop() Drop {
	return Drop{}
}

func (d Drop) SetDisableAutoCloseTransaction(b bool) Drop {
	d.DisableAutoCloseTransaction = b
	return d
}

func GetDropOptionByParams(opts []Drop) Drop {
	result := Drop{}
	for _, opt := range opts {
		if opt.DisableAutoCloseTransaction {
			result.DisableAutoCloseTransaction = opt.DisableAutoCloseTransaction
		}
	}
	return result
}

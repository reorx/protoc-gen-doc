package hugo

import "github.com/pseudomuto/protoc-gen-doc/extensions"

func init() {
	extensions.SetTransformer("hugo.page", func(payload interface{}) interface{} {
		v, ok := payload.(*Page)
		if !ok {
			return nil
		}
		return v
	})
}

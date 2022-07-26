package images

import "go.uber.org/fx"

func Module() fx.Option {
	return fx.Options(
		fx.Provide(newSocialPreviewGenerator),
		fx.Provide(NewImageProxyHandler),
		fx.Provide(NewSocialPreviewHandler),
	)
}

package fixtures

import "github.com/elgohr/counterfeiter/fixtures/aliased_package"

//go:generate counterfeiter . SomethingWithForeignInterface
type SomethingWithForeignInterface interface {
	the_aliased_package.InAliasedPackage
}

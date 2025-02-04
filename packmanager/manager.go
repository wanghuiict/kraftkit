// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2022, Unikraft GmbH and The KraftKit Authors.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.
package packmanager

import (
	"context"

	"kraftkit.sh/pack"
	"kraftkit.sh/unikraft/component"
)

type PackageManager interface {
	// Update retrieves and stores locally a cache of the upstream registry.
	Update(context.Context) error

	// Pack turns the provided component into the distributable package.  Since
	// components can comprise of other components, it is possible to return more
	// than one package.  It is possible to disable this and "flatten" a component
	// into a single package by setting a relevant `pack.PackOption`.
	Pack(context.Context, component.Component, ...PackOption) ([]pack.Package, error)

	// Unpack turns a given package into a usable component.  Since a package can
	// compromise of a multiple components, it is possible to return multiple
	// components.
	Unpack(context.Context, pack.Package, ...UnpackOption) ([]component.Component, error)

	// Catalog returns all packages known to the manager via given query
	Catalog(context.Context, CatalogQuery) ([]pack.Package, error)

	// Add a source to the package manager
	AddSource(context.Context, string) error

	// Remove a source from the package manager
	RemoveSource(context.Context, string) error

	// IsCompatible checks whether the provided source is compatible with the
	// package manager
	IsCompatible(context.Context, string) (PackageManager, error)

	// From is used to retrieve a sub-package manager.  For now, this is a small
	// hack used for the umbrella.
	From(string) (PackageManager, error)

	// Format returns the name of the implementation.
	Format() string
}

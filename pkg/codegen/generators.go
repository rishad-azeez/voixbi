package codegen

import (
	"bytes"
	"fmt"
	"path/filepath"

	"github.com/grafana/codejen"
	"github.com/grafana/grafana/pkg/kindsys"
	"github.com/grafana/thema"
)

type OneToOne codejen.OneToOne[*DeclForGen]
type OneToMany codejen.OneToMany[*DeclForGen]
type ManyToOne codejen.ManyToOne[*DeclForGen]
type ManyToMany codejen.ManyToMany[*DeclForGen]

// ForGen is a codejen input transformer that converts a pure kindsys.SomeDecl into
// a DeclForGen by binding its contained lineage.
func ForGen(rt *thema.Runtime, decl *kindsys.SomeDecl) (*DeclForGen, error) {
	lin, err := decl.BindKindLineage(rt)
	if err != nil {
		return nil, err
	}

	return &DeclForGen{
		SomeDecl: decl,
		lin:      lin,
	}, nil
}

// RawForGen produces a [DeclForGen] from a [kindsys.Raw] kind.
//
// Useful for grafana-external code generators, which depend on the Kind
// representation in registries produced by grafana core (such as
// ["github.com/grafana/grafana/pkg/registry/corekind".NewBase]).
func RawForGen(k kindsys.Raw) *DeclForGen {
	return &DeclForGen{
		SomeDecl: k.Decl().Some(),
	}
}

// StructuredForGen produces a [DeclForGen] from a [kindsys.Structured] kind.
//
// Useful for grafana-external code generators, which depend on the Kind
// representation in registries produced by grafana core (such as
// ["github.com/grafana/grafana/pkg/registry/corekind".NewBase]).
func StructuredForGen(k kindsys.Structured) *DeclForGen {
	return &DeclForGen{
		SomeDecl: k.Decl().Some(),
		lin:      k.Lineage(),
	}
}

// DeclForGen wraps [kindsys.SomeDecl] to provide trivial caching of
// the lineage declared by the kind (nil for raw kinds).
type DeclForGen struct {
	*kindsys.SomeDecl
	lin thema.Lineage
	sch thema.Lineage
}

// Lineage returns the [thema.Lineage] for the underlying [kindsys.SomeDecl].
func (decl *DeclForGen) Lineage() thema.Lineage {
	return decl.lin
}

// Schema returns the [thema.Schema] that a jenny should operate against, for those
// jennies that target a single schema.
func (decl *DeclForGen) Schema() thema.Lineage {
	return decl.sch
}

// SlashHeaderMapper produces a FileMapper that injects a comment header onto
// a [codejen.File] indicating the main generator that produced it (via the provided
// maingen, which should be a path) and the jenny or jennies that constructed the
// file.
func SlashHeaderMapper(maingen string) codejen.FileMapper {
	return func(f codejen.File) (codejen.File, error) {
		// Never inject on certain filetypes, it's never valid
		switch filepath.Ext(f.RelativePath) {
		case ".json", ".yml", ".yaml":
			return f, nil
		default:
			b := new(bytes.Buffer)
			fmt.Fprintf(b, headerTmpl, filepath.ToSlash(maingen), f.FromString())
			fmt.Fprint(b, string(f.Data))
			f.Data = b.Bytes()
		}
		return f, nil
	}
}

var headerTmpl = `// THIS FILE IS GENERATED. EDITING IS FUTILE.
//
// Generated by:
//     %s
// Using jennies:
//     %s
//
// Run 'make gen-cue' from repository root to regenerate.

`

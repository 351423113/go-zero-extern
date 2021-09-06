package generator

import (
	conf "github.com/lukebull/go-zero-extern/tools/goctl/config"
	"github.com/lukebull/go-zero-extern/tools/goctl/rpc/parser"
)

// Generator defines a generator interface to describe how to generate rpc service
type Generator interface {
	Prepare() error
	GenMain(ctx DirContext, proto parser.Proto, cfg *conf.Config) error
	GenCall(ctx DirContext, proto parser.Proto, cfg *conf.Config) error
	GenEtc(ctx DirContext, proto parser.Proto, cfg *conf.Config) error
	GenConfig(ctx DirContext, proto parser.Proto, cfg *conf.Config) error
	GenLogic(ctx DirContext, proto parser.Proto, cfg *conf.Config) error
	GenServer(ctx DirContext, proto parser.Proto, cfg *conf.Config) error
	GenSvc(ctx DirContext, proto parser.Proto, cfg *conf.Config) error
	GenPb(ctx DirContext, protoImportPath []string, proto parser.Proto, cfg *conf.Config, goOptions ...string) error
}

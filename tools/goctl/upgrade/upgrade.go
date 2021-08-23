package upgrade

import (
	"fmt"

	"github.com/351423113/go-zero-extern/tools/goctl/rpc/execx"
	"github.com/urfave/cli"
)

// Upgrade gets the latest goctl by
// go get -u github.com/351423113/go-zero-extern/tools/goctl
func Upgrade(_ *cli.Context) error {
	info, err := execx.Run("GO111MODULE=on GOPROXY=https://goproxy.cn/,direct go get -u github.com/351423113/go-zero-extern/tools/goctl", "")
	if err != nil {
		return err
	}

	fmt.Print(info)
	return nil
}

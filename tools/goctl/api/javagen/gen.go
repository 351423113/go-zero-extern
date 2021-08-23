package javagen

import (
	"errors"
	"fmt"
	"strings"

	"github.com/351423113/go-zero-extern/core/logx"
	"github.com/351423113/go-zero-extern/tools/goctl/api/parser"
	"github.com/351423113/go-zero-extern/tools/goctl/util"
	"github.com/logrusorgru/aurora"
	"github.com/urfave/cli"
)

// JavaCommand the generate java code command entrance
func JavaCommand(c *cli.Context) error {
	apiFile := c.String("api")
	dir := c.String("dir")
	if len(apiFile) == 0 {
		return errors.New("missing -api")
	}
	if len(dir) == 0 {
		return errors.New("missing -dir")
	}

	api, err := parser.Parse(apiFile)
	if err != nil {
		return err
	}

	packetName := api.Service.Name
	if strings.HasSuffix(packetName, "-api") {
		packetName = packetName[:len(packetName)-4]
	}

	logx.Must(util.MkdirIfNotExist(dir))
	logx.Must(genPacket(dir, packetName, api))
	logx.Must(genComponents(dir, packetName, api))

	fmt.Println(aurora.Green("Done."))
	return nil
}

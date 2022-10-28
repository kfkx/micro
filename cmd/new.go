package cmd

import (
	"fmt"
	"os"
	"path"

	"github.com/urfave/cli/v2"
	"gitlab.com/repod/micro/generator"
	"gitlab.com/repod/micro/generator/tmpls"
)

var (
	srvId          int
	srvName        string
	srvDescription string
	srvWorkspace   string
)

func init() {

}

// newService
func newService(ctx *cli.Context) error {

	fmt.Print("请输入服务ID (例如： 1024): ")
	fmt.Scanln(&srvId)

	fmt.Print("请输入服务名称 (例如： core): ")
	fmt.Scanln(&srvName)

	fmt.Print("请输入服务介绍 (例如： 核心服务): ")
	fmt.Scanln(&srvDescription)

	fmt.Print("请输入服务地址 (例如： $GOPATH/src/demo): ")
	fmt.Scanln(&srvWorkspace)

	if path.IsAbs(srvName) {
		fmt.Println("must provide a relative path as service name")
		return nil
	}

	if _, err := os.Stat(srvName); !os.IsNotExist(err) {
		return fmt.Errorf("%s already exists", err)
	}

	fmt.Printf("creating %d %s\n", srvId, srvName)

	g := generator.New(
		generator.Service(srvName),
		generator.Directory(srvWorkspace),
	)

	files := []generator.File{
		{Path: ".gitignore", Template: tmpls.Gitignore},
		{Path: "Dockerfile", Template: tmpls.Gitignore},
		{Path: ".dockerignore", Template: tmpls.DockerIgnore},
		{Path: "go.mod", Template: tmpls.Module},
		{Path: "service.yml", Template: tmpls.ServiceYml},
	}

	/*
		app/
			admin/
				proto/
					admin.proto
				init.go
			client/
				proto/
					client.proto
				init.go
			rpc/
				proto/
					rpc.proto
				init.go
		vm/
			db.go
			log.go
		logic/


	*/

	var dirs = []string{"admin", "client", "rpc"}
	for _, dir := range dirs {
		files = append(files, []generator.File{
			{Path: dir, Template: ""},
		}...)
	}

	if err := g.Generate(files); err != nil {
		return err
	}

	return nil
}

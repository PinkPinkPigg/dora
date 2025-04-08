package main

import (
	//"fmt"
	"github.com/PinkPinkPigg/dora/services/executor/gen"
	"github.com/PinkPinkPigg/dora/services/executor/impl"
	//"github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	//TIP <p>Press <shortcut actionId="ShowIntentionActions"/> when your caret is at the underlined text
	// to see how GoLand suggests fixing the warning.</p><p>Alternatively, if available, click the lightbulb to view possible fixes.</p>
	grpcServer := grpc.NewServer()
	//注册服务
	gen.RegisterExecutorServiceServer(grpcServer, impl.GetExecutorService())
}

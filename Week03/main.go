package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"net/http"
	"os"
	"os/signal"
)

func main() {
	group, ctx := errgroup.WithContext(context.Background())
	group.Go(func() error {
		return httpServer(ctx, "127.0.0.1:8001",&requestHandler{})
	})
	group.Go(func() error {
		return relatedSignal(ctx)
	})
	if err := group.Wait(); err != nil {
		fmt.Println("错误:", err.Error())
	}
	fmt.Println("退出!")
}

/**
获取信号进行处理
 */
func relatedSignal(ctx context.Context) error {
	c := make(chan os.Signal)
	signal.Notify(c) //监听幸好
	fmt.Println("开始监听信号！")
	for {
		select {
		case s := <-c:
			return fmt.Errorf("获得信号:%v", s)
		case <-ctx.Done():
			return fmt.Errorf("完成操作！")
		}
	}
}

/**
http启动服务
 */
func httpServer(ctx context.Context, addr string,h *requestHandler) error {
	s := http.Server{
		Addr:    addr,
		Handler : h,
	}

	go func(ctx context.Context) {
		<-ctx.Done()
		fmt.Println("操作完成！")
		s.Shutdown(ctx)
	}(ctx)
	fmt.Println("启动http服务")
	return s.ListenAndServe()
}

/**
定义一个结构体
 */
type requestHandler struct {

}

/**
实现接口
 */
func (h *requestHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}
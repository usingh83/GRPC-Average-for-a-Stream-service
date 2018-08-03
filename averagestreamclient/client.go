package main
import (
	"time"
	"fmt"
	"context"
	"log"
	"google.golang.org/grpc"
	"grpc-practice/averagestream/averagestreampb"
)
type server struct{}
func main(){
	cc,err:=grpc.Dial("localhost:50054",grpc.WithInsecure())
	if err!=nil{
		log.Fatalf("connection error %v",err)
	}
	defer cc.Close()
	c:=averagestreampb.NewAveragestreamClient(cc)
	req:=[]*averagestreampb.NumberRequest{
		&averagestreampb.NumberRequest{
			Request:int64(1),
		},
		&averagestreampb.NumberRequest{
			Request:int64(2),
		},
		&averagestreampb.NumberRequest{
			Request:int64(3),
		},
		&averagestreampb.NumberRequest{
			Request:int64(4),
		},
	}
	stream,err:=c.Average(context.Background())
	if err!=nil{
		log.Fatalf("Streaming error %v",err)
	}
	for _, request := range req{
		fmt.Println("Request is ",request)
		stream.Send(request)
		time.Sleep(1000*time.Millisecond)
	}
	response,err:=stream.CloseAndRecv()
	if err!=nil{
		log.Fatalf("Streaming error %v",err)
	}
	fmt.Println("Response is",response)
}
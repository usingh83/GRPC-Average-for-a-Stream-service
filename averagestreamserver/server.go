package main

import (
	"io"
	"log"
	"net"
	"google.golang.org/grpc"
	"grpc-practice/averagestream/averagestreampb"
)

type server struct{}

func (*server) Average(stream averagestreampb.Averagestream_AverageServer) error{
	sum:=float64(0)
	count:=float64(0)
	for {
		request,err:=stream.Recv()
		if err==io.EOF{
			return stream.SendAndClose(&averagestreampb.NumberResponse{
				Response:float32(sum/count),
			})
		}
		if err!=nil{
			log.Fatalf("stream error %v",err)
		}
		req:=request.GetRequest()
		sum+=float64(req)
		count+=1;
	}
}

func main(){
	lis,err:=net.Listen("tcp","0.0.0.0:50054")
	if err!=nil{
		log.Println("can not create listner %v",err)
	}
	s:=grpc.NewServer()
	averagestreampb.RegisterAveragestreamServer(s,&server{})
	if err:=s.Serve(lis);err!=nil{
		log.Println("can not create listner %v",err)
	}
}
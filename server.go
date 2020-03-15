package main

import "net/http"

/*
实现一个简单 http服务器
*/

func handler(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("hello,httpserver"))
}


func main(){
	
	   //1.注册处理函数
    /*
    http包下的Handlefunc函数，需要两个参数，一个是pattern，即“/”，为根
    目录，这里加入什么，访问时候就得在地址后面加入什么。第二个参数为一个回调
    函数，一个负责读写客户端的函数，这里将函数名写入即可，由系统自动按需调用
    */
    http.HandleFunc("/test",handler)
	http.Handle("/web/", http.StripPrefix("/web/", http.FileServer(http.Dir("web"))))
	http.Handle("/", http.StripPrefix("/web/", http.FileServer(http.Dir("/"))))
    //2.绑定服务器地址结构
    /*
    这一步也需要两个参数，第一个参数提供一个IP地址和端口，我们这里为了方便，
    直接引入回送地址，方便查看结果，第二个参数最好使用nil，剩下的有系统为我们
    分配，除非有特殊需求
   */
    http.ListenAndServe("127.0.0.1:7777",nil)

}

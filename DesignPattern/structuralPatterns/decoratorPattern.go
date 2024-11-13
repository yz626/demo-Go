package structuralPatterns

import (
	"log"
	"net/http"
	"time"
)

// 装饰器模式
// 结构设计模式
// 允许你通过将对象放入特殊封装对象中来为原对象增加新的行为，
// 原来的对象和修饰器的对象具有相同的接口，同样可以对修饰器进行修饰，
// 从而使得原本的对象获得所有修饰器的功能。

// example-1：
// Station 作为最基础的站点接口
// subwayStation 地铁站点,假设提供基础的站点信息
// securityCheckDecorator 安全检查修饰器，可以修饰站点，增加安全检查的功能
// epidemicCheckDecorator 疫情检查修饰器，可以修饰站点，增加疫情检查的功能

// Station 站点接口
// 作为接口，Station接口定义了所有站点的公共方法
type Station interface {
	Enter() string
}

// subwayStation 地铁站点
// 定义了基础的站点信息
type subwayStation struct {
	name string
}

func (s *subwayStation) Enter() string {
	return "subway station " + s.name
}

var _ Station = (*subwayStation)(nil)

func NewSubwayStation(name string) *subwayStation {
	return &subwayStation{name: name}
}

// securityCheckDecorator 安全检查修饰器
// 实现了Station接口
// 增加了站点安全检查的功能
type securityCheckDecorator struct {
	station Station
}

func (s *securityCheckDecorator) Enter() string {
	return "security check " + s.station.Enter()
}

var _ Station = (*securityCheckDecorator)(nil)

func NewSecurityCheckDecorator(station Station) *securityCheckDecorator {
	return &securityCheckDecorator{station: station}
}

// epidemicCheckDecorator 疫情检查修饰器
// 实现了Station接口
// 增加了站点疫情检查的功能
type epidemicCheckDecorator struct {
	station Station
}

func (e *epidemicCheckDecorator) Enter() string {
	return "epidemic check " + e.station.Enter()
}

var _ Station = (*epidemicCheckDecorator)(nil)

func NewEpidemicCheckDecorator(station Station) *epidemicCheckDecorator {
	return &epidemicCheckDecorator{station: station}
}

// example-2：
// 对于每个url请求路径，我希望添加一个日志，记录请求处理函数的运行时间，
// 并且希望这个日志可以动态的开启和关闭，
// 可以使用装饰器模式，而不用修改所有的请求处理函数

type HttpHandler func(w http.ResponseWriter, r *http.Request)

// Logger 装饰器
func Logger(handler HttpHandler) HttpHandler {
	return func(w http.ResponseWriter, r *http.Request) {
		now := time.Now()
		handler(w, r)
		log.Println("request time:", time.Since(now))
	}
}

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("hello world"))
	w.WriteHeader(http.StatusOK)
}

func Hello(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("hello"))
	w.WriteHeader(http.StatusOK)
}

func main() {
	test := http.NewServeMux()
	test.HandleFunc("/hello", Logger(Hello))
	test.HandleFunc("/helloWorld", Logger(HelloWorld))

	err := http.ListenAndServe(":8080", test)
	if err != nil {
		return
	}
}

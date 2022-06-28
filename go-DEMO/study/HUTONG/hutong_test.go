package hutong

import (
	"encoding/binary"
	"fmt"
	"io"
	"net"
	"sync"
	"sync/atomic"
	"testing"
)

var z1 = " 吃了没，您吶?"
var z2 = " 嗨！吃饱了溜溜弯儿。"
var z3 = " 回头去给老太太请安！"

var l4 = " 刚吃。"
var l5 = " 您这，嘛去？"
var l6 = " 有空家里坐坐啊。"

var TotalCount = 1000000
var PackCount = 10000

//协议结构
type Msg struct {
	Index uint32
	Data  string
}

func (m *Msg) Pack() []byte {
	length := 4 + len(m.Data)
	data := make([]byte, length+4, length+4)
	binary.BigEndian.PutUint32(data, uint32(length))
	binary.BigEndian.PutUint32(data[4:], m.Index)
	copy(data[8:], []byte(m.Data))
	return data
}
func (m *Msg) UnPack(conn net.Conn) {
	l := make([]byte, 4, 4)
	_, err := io.ReadFull(conn, l)
	if err != nil {
		return
	}
	length := binary.BigEndian.Uint32(l)
	d := make([]byte, length, length)
	_, err = io.ReadFull(conn, d)
	if err != nil {
		return
	}
	m.Index = binary.BigEndian.Uint32(d)
	m.Data = string(d[4:])
}
func (m *Msg) UnPack2(conn net.Conn, u *DaYe) {
	l := make([]byte, 4, 4)
	io.ReadFull(conn, l)
	length := binary.BigEndian.Uint32(l)
	d := make([]byte, length, length)
	io.ReadFull(conn, d)
	go func() {
		m.Index = binary.BigEndian.Uint32(d)
		m.Data = string(d[4:])
		u.Listen(m)
	}()
}

//大爷结构
type DaYe struct {
	Name        string
	ListenCount int32
	SpeakCount  int32
	conn        net.Conn
	group       sync.WaitGroup
}

func (dy *DaYe) Init(name string) {
	dy.Name = name
	dy.group.Add(2)
}

func (dy *DaYe) Speak(times int) {
	buf := make([]byte, 0, 128*times)
	for j := 0; j < times; j++ {
		switch dy.Name {
		case "张":
			buf = append(buf, (&Msg{Index: 1, Data: z1}).Pack()...)
			buf = append(buf, (&Msg{Index: 2, Data: z2}).Pack()...)
			buf = append(buf, (&Msg{Index: 3, Data: z3}).Pack()...)
			//dy.conn.Write((&Msg{Index: 1, Data: z1}).Pack())
			//dy.conn.Write((&Msg{Index: 2, Data: z2}).Pack())
			//dy.conn.Write((&Msg{Index: 3, Data: z3}).Pack())

		case "李":
			buf = append(buf, (&Msg{Index: 4, Data: l4}).Pack()...)
			buf = append(buf, (&Msg{Index: 5, Data: l5}).Pack()...)
			buf = append(buf, (&Msg{Index: 6, Data: l6}).Pack()...)
			//dy.conn.Write((&Msg{Index: 4, Data: l4}).Pack())
			//dy.conn.Write((&Msg{Index: 5, Data: l5}).Pack())
			//dy.conn.Write((&Msg{Index: 6, Data: l6}).Pack())

		}

	}
	dy.conn.Write(buf)

	atomic.AddInt32(&dy.SpeakCount, 3*int32(times))
	//fmt.Printf("SpeakCount: %d\n", dy.SpeakCount)
	if dy.SpeakCount == int32(TotalCount*3) {
		//大爷说完了
		dy.group.Done()
		fmt.Printf("大爷·%s 说完了\n", dy.Name)
	}
}
func (dy *DaYe) Listen(msg *Msg) {
	atomic.AddInt32(&dy.ListenCount, 1)
	//fmt.Printf("ListenCount: %d\n", dy.ListenCount)
	if dy.ListenCount == int32(TotalCount*3) {
		//大爷听完了
		dy.group.Done()
		fmt.Printf("大爷·%s 听完了\n", dy.Name)
	}
}

func (dy *DaYe) Go() {
	go func() {
		for i := 0; i < TotalCount/PackCount; i++ {
			dy.Speak(PackCount)
			//go dy.Speak(PackCount)
		}
	}()

	go func() {
		for {
			m := &Msg{}
			m.UnPack(dy.conn)
			dy.Listen(m)
			//m.UnPack2(dy.conn, dy)
		}
	}()
}
func (dy *DaYe) Wait() {
	dy.group.Wait()
}

//运行部分
func TestRun(t *testing.T) {
	zhang := &DaYe{}
	li := &DaYe{}
	zhang.Init("张")
	li.Init("李")

	//创建一个服务器
	serAddr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:8080")
	if err != nil {
		//fmt.Println(err)
		return
	}
	listen, err := net.ListenTCP("tcp", serAddr)
	if err != nil {
		//fmt.Println(err)
		return
	}
	go func() {
		for {
			conn, err := listen.Accept()
			if err != nil {
				//fmt.Println(err)
				continue
			}
			fmt.Printf("大爷来了 %s\n", conn.RemoteAddr().String())

			//服务器给张大爷
			zhang.conn = conn
			go zhang.Go()
		}
	}()

	//创建一个客户端
	conn, err := net.DialTCP("tcp", nil, serAddr)
	if err != nil {
		//fmt.Println(err)
		return
	}

	//客户端给李大爷
	li.conn = conn
	go li.Go()

	//等待两个大爷完事
	zhang.Wait()
	li.Wait()

	listen.Close()
	conn.Close()
}

//--- PASS: TestRun (8.66s)

package main

import (
	"ModbusMQTT/Library"
	"fmt"

	//"bytes"
	//"encoding/binary"
	//"fmt"
	"sync"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var Json []byte

func ModbusRead(wg *sync.WaitGroup, mbc Library.MBClient) interface{} {

	var add uint16
	var leng uint16
	leng = 10
	add = 0

	for {
		Regdata, err := mbc.ReadHoldReg(1, 0, 10)
		if err != nil {
			fmt.Print(err)
		}
		Coildata, err := mbc.ReadCoil(1, 0, 10)
		if err != nil {
			fmt.Print(err)
		}
		RegIndata, err := mbc.ReadRegIn(1, 0, 10)
		if err != nil {
			fmt.Print(err)
		}
		CoilIndata, err := mbc.ReadCoilIn(1, 0, 10)
		if err != nil {
			fmt.Print(err)
		}

		time.Sleep(time.Second)
		println("Modbus Read")
		Json = Library.JsonMaker(add, leng, Regdata, RegIndata, Coildata, CoilIndata)
		//b := []byte{RegIndata}
		//buf := bytes.NewBuffer(b)
		//var data uint32
		//binary.Read(buf, binary.LittleEndian, &data)
		//
		//fmt.Printf("0x%x\n", data)

	}
	return Json
}

func MQTTPublish(wg *sync.WaitGroup, mbc Library.MBClient, client mqtt.Client, topic string) {

	for {

		Library.MQTTPublish(client, topic, Json)
		println("data Publish")
		time.Sleep(time.Second)
	}
}

func main() {
	var wg sync.WaitGroup

	//Modbus Client Creat and TCP access
	mbc := Library.NewClient("127.0.0.1", 502)
	mbc.Open()

	//MQTT Client Creat and MQTT Broker access
	uri := "tcp://broker.hivemq.com:1883"
	topic := "test/topic12/1"
	client := Library.Connect("enitt", uri)

	println("@@ModbusRead Start@@")
	wg.Add(1)
	go ModbusRead(&wg, *mbc)

	println("@@MQTTPublish Start@@")
	wg.Add(2)
	go MQTTPublish(&wg, *mbc, client, topic)

	wg.Wait()
}

/*func Scrclr() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func ErrorCli() {
	println("-------------------------------------")
	fmt.Println("error")
	println("-------------------------------------")
	fmt.Println("You entered it iolncorrectly.\nReturn to the Output Coils menu.")
	println("-------------------------------------\n\n\n")
	fmt.Print("1 : Back\n", "2 : Main menu\n\n\n")
	fmt.Print("Select number Enter:")
}
func ContinueCli() {
	println("\n\n\n-------------------------------------")
	fmt.Println("1 : Back", "\n2 : Main menu")
	println("-------------------------------------\n\n\n")
	fmt.Print("Select number Enter:")
}*/

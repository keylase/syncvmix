package main
import (
    // "net"
    "os"
    "time"
    // "strings"
    "encoding/json"
    "fmt"
    "net/http"
    "io/ioutil"
    "log"
    "encoding/xml"
)

type Configuration struct {
  InVMIX string
  OutVMIX string
}
type VmixStatus struct {
    XMLName xml.Name `xml:"vmix"`
    Version string   `xml:"version"`
    Streaming bool `xml:"streaming"`
}

// type Result struct {
//   		XMLName xml.Name `xml:"Person"`
//   		Name    string   `xml:"FullName"`
//   		Phone   string
//   		Email   []Email
//   		Groups  []string `xml:"Group>Value"`
//   		Address
// }


// <vmix>
// <version>20.0.0.27</version>
// <edition>Basic</edition>
// <inputs>
// <input key="a4bd31df-5932-4b48-b638-913a483f05dc" number="1" type="Blank" title="Blank" state="Paused" position="0" duration="0" loop="False">Blank</input>
// <input key="7408daa7-5827-47b3-9527-7318f95bcfbe" number="2" type="Blank" title="Blank" state="Paused" position="0" duration="0" loop="False">Blank</input>
// </inputs>
// <overlays>
// <overlay number="1"/>
// <overlay number="2"/>
// <overlay number="3"/>
// <overlay number="4"/>
// <overlay number="5"/>
// <overlay number="6"/>
// </overlays>
// <preview>1</preview>
// <active>2</active>
// <fadeToBlack>False</fadeToBlack>
// <transitions>
// <transition number="1" effect="Fly" duration="500"/>
// <transition number="2" effect="Merge" duration="1000"/>
// <transition number="3" effect="CrossZoom" duration="1000"/>
// <transition number="4" effect="Merge" duration="1000"/>
// </transitions>
// <recording>False</recording>
// <external>False</external>
// <streaming>False</streaming>
// <playList>False</playList>
// <multiCorder>False</multiCorder>
// <fullscreen>False</fullscreen>
// <audio>
// <master volume="100" muted="False" meterF1="0" meterF2="0" headphonesVolume="100"/>
// </audio>
// </vmix>

func getXML(url string) (string, error) {
    resp, err := http.Get(url)
    if err != nil {
        return "", fmt.Errorf("GET error: %v", err)
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return "", fmt.Errorf("Status error: %v", resp.StatusCode)
    }

    data, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return "", fmt.Errorf("Read body: %v", err)
    }

    // log.Println(data)

    return string(data), nil
}

func main() {
  config, _ := os.Open("conf.json")
  decoder := json.NewDecoder(config)
  configuration := Configuration{}

  err := decoder.Decode(&configuration)
  if err != nil {
    fmt.Println("error:", err)
  }
  fmt.Println(configuration.InVMIX);

  for {

    //VMIX IN
    // servAddrIn := configuration.InVMIX
    // tcpAddrIn, err := net.ResolveTCPAddr("tcp", servAddrIn)
    // if err != nil {
    //     println("ResolveTCPAddrIN failed:", err.Error())
    //     continue
    // }
    // connIn, err := net.DialTCP("tcp", nil, tcpAddrIn)
    // if err != nil {
    //     println("Dial VMIX IN failed:", err.Error())
    //     continue
    // }


    if xmlStr, err := getXML(configuration.InVMIX); err != nil {
        log.Printf("Failed to get XML: %v", err)
        getXML(configuration.OutVMIX + "?Function=StopStreaming")
    } else {

        v:=VmixStatus{Version:""}
        err := xml.Unmarshal([]byte(xmlStr), &v)
        if err != nil {
          fmt.Printf("error: %v", err)
          // return
        }
        fmt.Printf("STREAMING: %#v\n", v.Streaming)

        if (v.Streaming){
          getXML(configuration.OutVMIX + "?Function=StartStreaming")
        } else {
          getXML(configuration.OutVMIX + "?Function=StopStreaming")
        }
    }



    //VMIX OUT
    // servAddrOut := configuration.OutVMIX
    // tcpAddrOut, err := net.ResolveTCPAddr("tcp", servAddrOut)
    // if err != nil {
    //     println("ResolveTCPAddr OUT failed:", err.Error())
    //     continue
    // }
    // connOut, err := net.DialTCP("tcp", nil, tcpAddrOut)
    // if err != nil {
    //     println("Dial VMIX OUT failed:", err.Error())
    //     continue
    // }


    // reply := make([]byte, 50)
    // _, err = connIn.Read(reply)
    // if err != nil {
    //     println("Read from vmix IN failed:", err.Error())
    //     continue
    // }
    // println("reply from VMIX In  server=", string(reply))
    //
    //
    // reply = make([]byte, 50)
    // _, err = connOut.Read(reply)
    // if err != nil {
    //     println("Read from vmix IN  failed:", err.Error())
    //     continue
    // }
    //
    // println("reply from VMIX Out server=", string(reply))
    //
    //
    //
    //
    // strEcho := "XMLTEXT vmix/streaming\r\n"
    // // strEcho := "FUNCTION StartStreaming\r\n";
    //
    // _, err = connIn.Write([]byte(strEcho))
    // if err != nil {
    //     println("Get data from VMIX In server failed:", err.Error())
    //     continue
    // }
    //
    // // println("write to server = ", strEcho)
    //
    // reply = make([]byte, 50)
    //
    // _, err = connIn.Read(reply)
    // // if err != nil {
    // //     println("Write to server failed:", err.Error())
    // //     os.Exit(1)
    // // }
    //
    // // println("reply from server=", string(reply))
    // if (strings.Contains(string(reply), "True")){
    //   println("Streaming ok")
    //   // strEcho = "FUNCTION StopStreaming\r\n";
    //   strEcho = "FUNCTION StartStreaming\r\n";
    //
    //   _,err = connOut.Write([]byte(strEcho))
    // } else {
    //   println("Not streaming")
    //   // strEcho = "FUNCTION StartStreaming\r\n";
    //   strEcho = "FUNCTION StopStreaming\r\n";
    //
    //   _,err = connOut.Write([]byte(strEcho))
    //
    //
    // }
    time.Sleep(500 * time.Millisecond)
    //
    // connIn.Close()
    // connOut.Close()
  }

}

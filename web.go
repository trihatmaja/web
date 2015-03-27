package main

import (
    "os"
    "fmt"
    "net/http"
    "github.com/gorilla/mux"
    "html/template"
    "github.com/mrmorphic/hwio"
    "golang.org/x/net/websocket"
    "time"
    "strconv"
    "math"
    "encoding/json"
)

var (
    in string
)

type Response map[string]interface{}

func (r Response) String() (s string) {
        b, err := json.Marshal(r)
        if err != nil {
                s = ""
                return
        }
        s = string(b)
        return
}

func handler(w http.ResponseWriter, r *http.Request) {
    t, _ := template.ParseFiles("index.html")
    t.Execute(w, nil)
}

func writeblaster(pin string, value string){
    f, err := os.Create("/dev/pi-blaster")
    if err != nil {panic(err)}
    _, err = f.WriteString(fmt.Sprintf("%s=%s\n", pin, value))
    if err != nil {panic(err)}
    f.Sync()
}

func on (w http.ResponseWriter, r *http.Request) {
    myPin1, err:= hwio.GetPinWithMode("gpio18", hwio.INPUT)
    if err !=nil{panic(err)}
    w.Header().Set("Content-Type", "application/json")
    r.ParseForm()
    
    if r.FormValue("turnon") == "true" {
	writeblaster("17", "1")
	time.Sleep(100 * time.Millisecond)
        val, _ := hwio.DigitalRead(myPin1)
	if val == hwio.HIGH{
            fmt.Fprint(w, Response{"success": true})}else {
            fmt.Fprint(w, Response{"success": false})}
    } else {
	writeblaster("17", "0")
        time.Sleep(100 * time.Millisecond)
        val, _ := hwio.DigitalRead(myPin1)
        if val == hwio.LOW{
            fmt.Fprint(w, Response{"success": true})}else{
	    fmt.Fprint(w, Response{"success": false})}
    }
    hwio.ClosePin(myPin1)
}

func socketHandler(ws *websocket.Conn){
    for{
    	if err:= websocket.Message.Receive(ws, &in);err!=nil{
		return
    	}
    	writeblaster("17", in)
    }
}

func currentHandler(ws *websocket.Conn){
    m, e := hwio.GetModule("i2c")
    if e != nil {fmt.Printf("could not get i2c module: %s\n", e)
    return}
    i2c := m.(hwio.I2CModule)

    i2c.Enable()
    defer i2c.Disable()
    device := i2c.GetDevice(0x48)

    for {
    	a, _ := device.Read(0x48, 1)
        s := fmt.Sprintf("%d", a[0])
        i, _ := strconv.ParseFloat(s, 64)
        val := math.Abs(i-124)*0.039215
    	websocket.Message.Send(ws, fmt.Sprintf("%v", val))
	time.Sleep(1 * time.Second)
    }
}

func init(){
    defer hwio.CloseAll()
}

func main() {
    r := mux.NewRouter()
    r.HandleFunc("/", handler)
    r.HandleFunc("/control/on", on).Methods("POST")
    r.Handle("/ws/slide", websocket.Handler(socketHandler))
    r.Handle("/ws/current", websocket.Handler(currentHandler))
    fmt.Println("listen and server at port: 8080")
    http.ListenAndServe(":8080", r)
}

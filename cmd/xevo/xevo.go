package main

import (
  "encoding/json"
  "fmt"
  "io/ioutil"
  "os"
  //"strings"
  //"plugin"
)

var config map[string]interface{}

func main() {
  args := os.Args
  config = getJson("/etc/xevo/config.json")

  getPlugins("./plugins")

  if (len(args) <= 1) {
    help()
  } else {
    switch args[1] {
    case "help":
      help()
    default:
      help()
    }
  }
}

func getPlugins(plugin_dir string) map[string]interface{} {
  var plugins map[string]interface{}
  pluginFiles, err := ioutil.ReadDir(plugin_dir)
  if err !=nil {
    panic(err)
  }

  for _, file := range pluginFiles {
    plugins[file.Name()], err = plugin.Open("plugin_name.so")
    if err != nil {
       panic(err)
    }
    Execute, err := plugins[file.Name()].(interface{}).Lookup("Execute")
    if err != nil {
      panic(err)
    }
    Help, err := plugins[file.Name()].(interface{}).Lookup("Help")
    if err != nil {
      panic(err)
    }
    Execute.(func(args []string))(pluginFiles)
    Help.(func())()
  }

  return plugins
}

func getJson(file string) map[string]interface{} {
  var dat map[string]interface{}
  raw, err := ioutil.ReadFile(file)
  if err != nil {
    fmt.Println(err.Error())
    os.Exit(1)
  }
  json.Unmarshal(raw, &dat)
  return dat
}

func help() {
  fmt.Printf("xevo Version: %s development\n\n", config["version"])
  fmt.Printf(" Usage: xevo [options] command [command options....]\n\n")
  fmt.Printf(" Commands:\n")
  fmt.Printf("\thelp  Print this menu\n")
}
////////////////////////////////STUFFF
/*argsWithProg := os.Args
p, err := plugin.Open("plugin_name.so")
if err != nil {
   panic(err)
}
v, err := p.Lookup("V")
if err != nil {
  panic(err)
}
f, err := p.Lookup("F")
if err != nil {
  panic(err)
}
test, err := p.Lookup("Test")
if err != nil {
  panic(err)
}
*v.(*int) = 7
f.(func())() // prints "Hello, number 7"
test.(func(printthis int))(*v.(*int))
fmt.Println(argsWithProg)
fmt.Println(len(argsWithProg))

config := getJson("/etc/xevo/config.json")
plugin := getJson("./download.json")

fmt.Println(plugin["file"])
fmt.Println(config["os"])
keywords := config["os"].(map[string]interface{})["keywords"]
fmt.Println(keywords.([]interface{})[1])
enc := json.NewEncoder(os.Stdout)
enc.Encode(plugin)
enc.Encode(config)*/

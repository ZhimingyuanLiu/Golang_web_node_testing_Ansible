package handler


import (  
	"strconv"
	"html/template"
	"net/http"
	"encoding/json"
    "fmt"  
    "io/ioutil"  
    // "os"  
	"strings"
)

type testMes struct {
	Name string
	Status string
	Mes string
}

type nodeTest struct {
	List []testMes
	Name string
}

type testMap struct {
	Good []nodeTest
	Err []nodeTest
}

type feature struct {
	Name string
	Mes []featureInfo
}

type nodeFeature struct {
	Name string
	Features []feature
}

type allTest struct {
	Test testMap
	Feature []nodeFeature
}

type featureInfo struct {
	Author string
	Name string
	Date string
	Desc string
}

func ReadTest(fileName string, i int) testMap{
	var list testMap
    // if err != nil {  
    //     return nil  
    // }
	ra, _ := ioutil.ReadFile(fileName)
	linelist := strings.Split(string(ra),"\n")
	fmt.Println(len(linelist))
	var mes testMes
	var gnodetest nodeTest
	var enodetest nodeTest
	gnodetest.Name = "node" + strconv.Itoa(i)
	enodetest.Name = "node" + strconv.Itoa(i)
	for i:= range linelist {
		mess := strings.Split(linelist[i],"#")
		if len(mess)==3{
			mes.Name = mess[0]
			mes.Status = mess[1]
			mes.Mes = mess[2]
			if mess[1] == "pass"{
				gnodetest.List = append(gnodetest.List,mes)
			}else{
				enodetest.List = append(enodetest.List,mes)
			}
		}
	 }
	 list.Good = append(list.Good,gnodetest)
	 list.Err = append(list.Err,enodetest)
    // for {
	// 	line, _ := buf.ReadString('\n')
	// 	line = strings.Replace(line, "\n", "", -1)
	// 	mess := strings.Split(line,"#")
	// 	mes.Name = mess[0]
	// 	mes.Status = mess[1]
	// 	mes.Mes = mess[2]
	// 	list.List = append(list.List,mes)
	// }
	fmt.Println(list)
    return list  
}  

func ReadFeature(fileName string, i int) nodeFeature{
	var node nodeFeature
	var feature feature
    // if err != nil {  
    //     return nil  
    // }
	ra, _ := ioutil.ReadFile(fileName)
	linelist := strings.Split(string(ra),"\n")
	fmt.Println(len(linelist))
	node.Name = "node" + strconv.Itoa(i)
	feaitem := [...]string {"dashboard","heapster","fsmon","heapster"}
	for i,_ := range linelist {
		var fi []featureInfo
		json.Unmarshal([]byte(linelist[i]),&fi)
		feature.Name = feaitem[i]
		feature.Mes = fi
		node.Features = append(node.Features, feature)
	}
    // for {
	// 	line, _ := buf.ReadString('\n')
	// 	line = strings.Replace(line, "\n", "", -1)
	// 	mess := strings.Split(line,"#")
	// 	mes.Name = mess[0]
	// 	mes.Status = mess[1]
	// 	mes.Mes = mess[2]
	// 	list.List = append(list.List,mes)
	// }
    return node  
}  

func GetMes(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "text/html")
	w.Header().Set("charset", "utf-8")
	var alltest allTest
	var testmap testMap
	dir_list, e := ioutil.ReadDir("logs")
    if e != nil {
		fmt.Println("read dir error")
		w.WriteHeader(403)
        return
	}
	len := len(dir_list)/2
    for a := 0; a < len; a++ {
		path := "logs/ansible_error_node"
		path += strconv.Itoa(a+1)
		path += ".txt"
		testmap = ReadTest(path,a+1)
		path = "logs/feature_node"
		path += strconv.Itoa(a+1)
		path += ".txt"
		alltest.Feature = append(alltest.Feature,ReadFeature(path,a+1))
    }
	// var tlist testMap
	// err := json.Unmarshal(wlist, &tlist)
  
	// if err!=nil {
	//   fmt.Println("读取文件失败")
	//   w.WriteHeader(403)
	//   return
	// }
	
	alltest.Test = testmap
	t, _ := template.ParseFiles("tpl/tpl.html")
	if t != nil{
		t.Execute(w, alltest)
		// w.Write([]byte(t.Execute(os.Stdout, alltest)))
	}else{
		w.WriteHeader(403)
		w.Write([]byte("err"))
	}
	// jn,_ := json.Marshal(alltest)
	// w.Write(jn)
}

// func GetMesHandler() http.Handler {
// 	return http.HandlerFunc(GetMes)
//   }
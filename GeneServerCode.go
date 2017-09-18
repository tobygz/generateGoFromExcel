package main

import (
    "fmt"
    "os"
    "strings"
    "path/filepath"

    "os/signal"
    "github.com/xuri/excelize"
)



var g_outdir string = "../GeneteServerCode/"


func writeToFile(fname string, content string) int {
    fmt.Println("\twritefile: ", fname)
    f, _ := os.Create(fname)
    defer f.Close()

    n, _ := f.Write([]byte(content ) )
    f.Close()
    return n    
}

func parseXlsmFile(filename string) {
    xlsx, err := excelize.OpenFile( filename )
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

    fmt.Println(fmt.Sprintf("\r\nprocess file name:%s",  filename ))
    for idx:=1; idx<=xlsx.SheetCount;idx++{
        content := "package cfg\r\n\r\n"
        SheetName := xlsx.GetSheetName(idx)
        if !strings.Contains( SheetName, "Base" ) {
            continue
        }

        map_name := fmt.Sprintf("%s_cfgMap", SheetName)
        content = fmt.Sprintf("%s\r\nfunc Get_%s(id int, key string) string {\r\n\tval1, ok1 := %s[id]\r\n\tif !ok1 {\r\n\t\treturn \"\"\r\n\t}\r\n\tval2, ok2 := (*val1)[key]\r\n\tif !ok2 {\r\n\t\treturn \"\"\r\n\t}\r\n\treturn val2\r\n}\r\n\r\n\r\n", content, map_name,map_name )


        fmt.Println(fmt.Sprintf("\tprocess sheet name:%s", SheetName ))
        rows := xlsx.GetRows( fmt.Sprintf("Sheet%d",idx) )
        isServMap := make(map[int]bool, 0)
        //获取服务器配置项
        for _, row := range rows {
            r_idx := 0
            for _, cell := range row {
                if cell == "YES" {
                    isServMap[r_idx] = true
                }  
                r_idx++              
            }            
            break
        }        

        //获取字段名        
        fieldKeyMap := make( map[int]string, 0)

        for _, row := range rows {            
            if row[0] != "KEY"{
                continue
            }
            r_idx := 0
            for _, cell := range row {                
                fieldKeyMap[r_idx] = cell   
                //fmt.Println(r_idx,":", cell)             
                r_idx++             
            }            
            break          
        }        

        fieldNameMap := make( map[int]string, 0)
        for _, row := range rows {            
            if row[0] != "NAME"{
                continue
            }
            r_idx := 0
            for _, cell := range row {                
                fieldNameMap[r_idx] = cell                
                r_idx++             
            }            
            break          
        }        
        
        /*
        format eg
            SheetName_cfgMap := map[int]*map[string]string{
                1: &map[string]string {
                    "a":"a","b":"b",
                    },
            }
        */
        //value
        content = fmt.Sprintf("%svar %s_cfgMap map[int]*map[string]string = map[int]*map[string]string{ \r\n",content, SheetName)               
        for _, row := range rows {            
            if row[0] != "VALUE"{
                continue
            }            
            r_idx := 0            
            for _, cell := range row {       
                if r_idx == 0 {
                    r_idx++
                    continue
                }  
                if v,ok := isServMap[r_idx]; !ok || !v {
                    r_idx++
                    continue
                }
                if r_idx == 1 {
                    content = fmt.Sprintf( "%s%s:&map[string]string {\r\n", content, cell)
                }                            
                content = fmt.Sprintf("%s\"%s\":\"%v\",\r\n", content, fieldKeyMap[r_idx],cell )
                r_idx++             
            } 
            content = fmt.Sprintf("%s},\r\n", content)                             
        }                
        content = fmt.Sprintf("%s}\r\n", content)                             

        //write file
        outFileName := fmt.Sprintf("%s%s.go", g_outdir, SheetName)
        writeToFile(outFileName, content)               
    }
}

func pause() {  
    fmt.Println("CTRL+C 退出...")  
    c := make(chan os.Signal, 1)
    signal.Notify(c, os.Interrupt, os.Kill )
    <-c 
} 


func main() {
    //g_outdir
    _, err0 := os.Stat(g_outdir)
    if os.IsNotExist(err0) {
        os.Mkdir(g_outdir, os.ModePerm)
    }

    path := "./"
    ext := "xlsm"
    err := filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
            if ( f == nil ) {return err}
            if f.IsDir() {return nil}
            if strings.Contains(path, ext){
                //println(path)
                parseXlsmFile(path)
            }
            return nil
    })
    
    if err != nil {
            fmt.Printf("filepath.Walk() returned %v\n", err)
    }    

    pause()
}

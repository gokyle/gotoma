package main

import (
        "fmt"
        "os"
        "time"
)

const timeFormat = "2006-01-_2 15:04"

func logEvent(event string) {
        if logfile != "" {
                err := logToFile(event)
                if err != nil {
                        fmt.Println("[!] logEvent: ", err)
                }
        }
}

func logToFile(event string) (err error) {
        nonExist := fmt.Sprintf("open %s: no such file or directory",
                logfile)
        file, err := os.OpenFile(logfile, os.O_WRONLY | os.O_APPEND, 0600)
        if err != nil && err.Error() == nonExist {
                file, err = os.Create(logfile)
        } 
        
        if err != nil {
                return
        }
        defer file.Close()

        event = fmt.Sprintf("%s %s\n", time.Now().Format(timeFormat), event)
        _, err = file.WriteString(event)
        return                
}

package main

import (
    "fmt"      
    "net/url"  
    "log"      
    "os"       
)

func main(){
    banner()
    
    if len(os.Args) < 2 {
        help()      
        os.Exit(1) 
    }
    
    comando := os.Args[1]
    
    switch comando {
    case "u", "uncode":  
        if len(os.Args) < 3 {
            fmt.Println("[!]Error no url")
            os.Exit(1)
        }
        urlEncoded := os.Args[2]
        decoded := decodeURL(urlEncoded)
        fmt.Println(decoded)
        
    case "c", "code":  
        if len(os.Args) < 3 {
            fmt.Println("[!]Error no url")
            os.Exit(1)
        }
        urlDecoded := os.Args[2]
        encoded := encodeURL(urlDecoded)
        fmt.Println(encoded)
        
    case "h", "-h", "--help":
        help()
        
    default:
        fmt.Println("Comando no reconocido. Usa 'h' para ayuda.")
        os.Exit(1)
    }
}

func banner(){
    banner := `   __  ___   ____________  ____  ________  ______  __ 
  / / / / | / / ____/ __ \/ __ \/ ____/ / / / __ \/ / 
 / / / /  |/ / /   / / / / / / / __/ / / / / /_/ / /  
/ /_/ / /|  / /___/ /_/ / /_/ / /___/ /_/ / _, _/ /___
\____/_/ |_/\____/\____/_____/_____/\____/_/ |_/_____/`
    
    fmt.Println(banner)
    fmt.Println()
}

func help(){
    fmt.Println("Use: <command> <url>")
    fmt.Println("  u, uncode")
    fmt.Println("  c, code")
    fmt.Println("  h, help")
}

// Decodifica URL (quita los %)
func decodeURL(encodedURL string) string {
    decodedURL, err := url.QueryUnescape(encodedURL)
    
    if err != nil {
        log.Fatal("Error al decodificar:", err)
    }
    
    return decodedURL
}

// Codifica URL (pone los %)
func encodeURL(decodedURL string) string {
    return url.QueryEscape(decodedURL)
}
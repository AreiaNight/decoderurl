package main 
import (
    "fmt"
    "log"
    "os"
    "encoding/base64"
    "net/url" 
    "github.com/common-nighthawk/go-figure" //Repo for ascii letters
)
func main(){
    banner()
    if len(os.Args) < 2 { // Se queda como 2 porque solo requiere un -argumento 
        fmt.Println("[!]Error no command")
        os.Exit(1)
    }
    comando := os.Args[1]
    switch comando { 
        case "-h", "--help":
            help()
        case "-b64", "--base64":
            if len(os.Args) < 3 {
                fmt.Println("[!]Error no file")
                os.Exit(1)
            }
            filePath := os.Args[2] // El contenido del file es igual al argumento 2, en este caso lo que viene después del -b64 o --base64
            codedFile, err := os.ReadFile(filePath) // Lee el contenido del file, el pathFile indica dónde se encuentra el file a leer, si esta en la misma carpeta, solo se pone el nombre del file
            if err != nil { // Si hay un error al leer el file, se muestra el error
                log.Fatal("[!]Error:", err)
            }
            decodedString, err := decodedBase64(string(codedFile)) // Se llama a la función decodedBase64 para decodificar el contenido del file
            if err != nil {
                log.Fatal("[!]Error:", err)
            }
            fmt.Println(decodedString) // Se imprime el resultado de la decodificación
        case "-u", "--uncode":
            if len(os.Args) < 3 {
                fmt.Println("[!]Error no url")
                os.Exit(1)
            }
            urlEncoded := os.Args[2]
            decoded := decodeURL(urlEncoded)
            fmt.Println(decoded)
        case "-c", "--code":
            if len(os.Args) < 3 {
                fmt.Println("[!]Error no url")
                os.Exit(1)
            }
            urlDecoded := os.Args[2]
            encoded := encodeURL(urlDecoded)
            fmt.Println("\n" + encoded + "\n")
        default:
            fmt.Println("[!]No command, exiting.")
            os.Exit(1)
    }
}
// ----- FUNCIONES -----
func banner(){
    myFigure := figure.NewFigure("DeGOded", "slant", true)
    myFigure.Print()
}


func help(){
    fmt.Println("\nWhat's DecodedGo?")
    fmt.Println("DecodedGo is a simple tool to encode and decode urls and base64 strings.")
    fmt.Println("This tool will get constant updates, depends on what I need")
    fmt.Println("Commands:")
    fmt.Println("  -h, --help       Show this help message")
    fmt.Println("  -b64, --base64   Decode a base64 file")
    fmt.Println("  -u, --uncode     Decode a URL")
    fmt.Println("  -c, --code       Encode a URL")
}


// Decodifica URL (quita los %)
func decodeURL(encodedURL string) string {
    decodedURL, err := url.PathUnescape(encodedURL) 
    if err != nil {
        log.Fatal("[!]Error:", err)
    }
    return decodedURL
}


// Codifica URL (pone los %)
func encodeURL(decodedURL string) string {
    return url.PathEscape(decodedURL)
}


func decodedBase64(encodedString string) (string, error) {
    decoded, err := base64.StdEncoding.DecodeString(encodedString) // Decodifica la cadena base64
    if err != nil {
        return "", err
    }
    return string(decoded), nil
}
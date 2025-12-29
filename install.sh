#!/bin/bash

# Colores
greenColour="\e[0;32m\033[1m"
endColour="\033[0m\e[0m"
redColour="\e[0;31m\033[1m"
blueColour="\e[0;34m\033[1m"
yellowColour="\e[0;33m\033[1m"
purpleColour="\e[0;35m\033[1m"
turquoiseColour="\e[0;36m\033[1m"
grayColour="\e[0;37m\033[1m"

# Define color variables para ASCII
RED='\033[0;31m'
GREEN='\033[0;32m'
BLUE='\033[0;34m'
PURPLE='\033[0;35m'
CYAN='\033[0;36m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

function ctrl_c(){
    echo -e "\n[!] Saliendo...\n"
    tput cnorm
    exit 1
}


# Capturar Ctrl+C
trap ctrl_c SIGINT


# Actualización del sistema
echo -e "${redColour}\n\n[!]Updating system, please wait... This will taka a while so why not go for a snack or pet your pets?"
sudo pacman -Syu --noconfirm > /dev/null 2>&1

# Instalación de go si es necesario  
if ! command -v go > /dev/null 2>&1; then
    echo -e "${greenColour}[+] Installing go${endColour}"
    sudo pacman -S --noconfirm go > /dev/null 2>&1
fi

# Compilar el programa
echo -e "${yellowColour}[+] Compiling uncodeurl...${endColour}"
go build -o uncodeurl main.go

if [ $? -ne 0 ]; then
    echo -e "${redColour}[!] Error! Try again${endColour}"
    exit 1
fi

# Mover el binario a /usr/local/bin
echo -e "${greenColour}[+] Instalando uncodeurl en /usr/local/bin${endColour}"
sudo mv uncodeurl /usr/local/bin/

# Dar permisos de ejecución
sudo chmod +x /usr/local/bin/uncodeurl

# Verificar instalación
if command -v uncodeurl > /dev/null 2>&1; then
    echo -e "${greenColour}\n[!] Done!${endColour}"
else
    echo -e "${redColour}[!] Error: No se pudo instalar correctamente${endColour}"
    exit 1
fi
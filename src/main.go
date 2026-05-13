package main

import (
	"bytes"
	"image"

	//	"image/png"
	//"io/ioutil"
	"bufio"
	"log"
	"os"

	"github.com/auyer/steganography"
	// "github.com/auyer/steganography"
)

func main() {
	// Ler a Mensagem do arquivo txt
	message, err := os.ReadFile("/home/otaviormc/Desktop/steganography/src/messageInputFile.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Ler a imagem
	inputFile, err := os.Open("/home/otaviormc/Desktop/steganography/src/inputImage.png")
	if err != nil {
		log.Fatalf("Erro ao abrir a imagem: %v", err)
	}

	defer inputFile.Close()

	// Ler dados binários da imagem
	reader := bufio.NewReader(inputFile)

	// Decodificar dados binários
	img, _, err := image.Decode(reader)

	if err != nil {
		log.Fatalf("Erro ao decodificar a imagem %v",err)
	}

	encondedImg := new(bytes.Buffer)

	// Codifica a mensagem na imagem
	err = steganography.Encode(encondedImg, img, message)

	if err != nil {
		log.Fatalf("Erro ao codificar a mensagem na imagem %v",err)
	}

	// Cria o arquivo output
	outputFile, err := os.Create("encoded_img.png")
	if err != nil {
		log.Fatalf("Erro ao criar arquivo de output %v", err)
	}

	// Escreve a imagem no disco
	bufio.NewWriter(outputFile).Write(encondedImg.Bytes())	
}

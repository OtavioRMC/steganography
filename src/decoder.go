package main

import (
	"bufio"
	"fmt"
	"image"
	"log"
	"os"

	"github.com/auyer/steganography"
)

func main() {
	inFile, err := os.Open("/home/otaviormc/Desktop/steganography/src/encoded_img.png")
	if err != nil {
		log.Fatalf("Erro ao abrir o arquivo codificado %v", err)
	}

	defer inFile.Close()

	reader := bufio.NewReader(inFile)
	img, _, err := image.Decode(reader)
	if err != nil {
		log.Fatalf("Erro ao decodificar a imagem %v", err)
	}

	// Retorna o tamanho da mensagem
	sizeOfMessage := steganography.GetMessageSizeFromImage(img)

	msg := steganography.Decode(sizeOfMessage, img)
	for i := range msg {
		fmt.Printf("%c", msg[i])
	}
}

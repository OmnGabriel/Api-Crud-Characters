package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
)

type Testsuite struct {
	XMLName  xml.Name `xml:"testsuite"`
	Name     string   `xml:"name,attr"`
	Tests    int      `xml:"tests,attr"`
	Coverage string   `xml:"system-out"` // Adicione um campo para a cobertura de testes
}

type Testsuites struct {
	XMLName    xml.Name    `xml:"testsuites"`
	Testsuites []Testsuite `xml:"testsuite"`
}

func main() {
	// Abra o arquivo XML
	xmlFile, err := os.Open("report.xml")
	if err != nil {
		fmt.Println("Erro ao abrir o arquivo XML:", err)
		return
	}
	defer xmlFile.Close()

	// Leia o conteúdo do arquivo XML
	xmlData, err := io.ReadAll(xmlFile)
	if err != nil {
		fmt.Println("Erro ao ler o arquivo XML:", err)
		return
	}

	var testSuites Testsuites
	if err := xml.Unmarshal(xmlData, &testSuites); err != nil {
		fmt.Println("Erro ao analisar o XML:", err)
		return
	}

	// Crie um cabeçalho de tabela Markdown
	fmt.Println("|	Suite Name	|   Tests	|	Coverage	|")
	fmt.Println("|-----|-----|-----|")

	// Preencha a tabela Markdown com dados
	for _, suite := range testSuites.Testsuites {
		suiteName := suite.Name
		tests := suite.Tests
		coverage := suite.Coverage // A cobertura de testes

		// Crie uma linha na tabela Markdown
		row := fmt.Sprintf("| %s | %d | %s |", suiteName, tests, coverage)
		fmt.Println(row)
	}
}

package main

import (
	"io"
	"net/http"
	"os"
	"sync"
)


func downloadFile(link string, outputPath string, wg *sync.WaitGroup)  {
	out, _ := os.Create(outputPath) // Create a file
	defer out.Close() // Close the file (defer is telling us to close at te end of the function)

	resp, err := http.Get(link) // Get request

	if err != nil {
		print(err)
		return
	}

	defer resp.Body.Close() // Close request

	io.Copy(out, resp.Body) // Save the file with request content

	wg.Done() // decrement counter for goroutines
}

func main() {
	m := map[string]string{
		"/Users/leonardocastro/Projects/go/Download File/download/output_1.csv":"https://www.portaltransparencia.gov.br/beneficios/consulta/baixar?paginacaoSimples=true&direcaoOrdenacao=asc&de=01%2F02%2F2022&ate=28%2F02%2F2022&colunasSelecionadas=linkDetalhamento%2ClinguagemCidada%2CmesAno%2Cuf%2Cmunicipio%2Cvalor",
		"/Users/leonardocastro/Projects/go/Download File/download/output_2.csv":"https://www.portaltransparencia.gov.br/beneficios/consulta/baixar?paginacaoSimples=true&direcaoOrdenacao=asc&de=01%2F02%2F2022&ate=28%2F02%2F2022&colunasSelecionadas=linkDetalhamento%2ClinguagemCidada%2CmesAno%2Cuf%2Cmunicipio%2Cvalor",
		"/Users/leonardocastro/Projects/go/Download File/download/output_3.csv":"https://www.portaltransparencia.gov.br/beneficios/consulta/baixar?paginacaoSimples=true&direcaoOrdenacao=asc&de=01%2F02%2F2022&ate=28%2F02%2F2022&colunasSelecionadas=linkDetalhamento%2ClinguagemCidada%2CmesAno%2Cuf%2Cmunicipio%2Cvalor",
		"/Users/leonardocastro/Projects/go/Download File/download/output_4.csv":"https://www.portaltransparencia.gov.br/beneficios/consulta/baixar?paginacaoSimples=true&direcaoOrdenacao=asc&de=01%2F02%2F2022&ate=28%2F02%2F2022&colunasSelecionadas=linkDetalhamento%2ClinguagemCidada%2CmesAno%2Cuf%2Cmunicipio%2Cvalor"}

	// This WaitGroup is used to wait for all the goroutines launched here to finish
	var wg sync.WaitGroup
	for k, v := range m {
		wg.Add(1) // incrment counter for goroutines
		go downloadFile(v, k, &wg)
	}
	wg.Wait()  // Main Goroutine will wait till incremnet counter is zero
}

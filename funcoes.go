package componentes

import (
	"bytes"
	_ "bytes"
	_ "database/sql"
	"encoding/csv"
	_ "encoding/json"
	"fmt"
	"io"
	_ "io"
	"io/ioutil"
	"log"
	"math"
	_ "net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	_ "strconv"
	"time"
	"unicode"

	_ "github.com/gocraft/dbr"
	_ "github.com/lib/pq"
	"github.com/pdfcpu/pdfcpu/pkg/api"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

// Leitura de Arquivos diretamente do sistema operacional
func LeituraArquivo(arquivo string) ([]byte, error) {
	jsonFile, err := os.Open(arquivo)
	if err != nil {
		fmt.Println("Error opening JSON file:", err)
		return nil, err
	}

	defer jsonFile.Close()
	jsonData, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println("Error reading JSON data:", err)
		return nil, err

	}

	return jsonData, err

}

// Diretorio e a extensão, somente inserir jpeg, png, etc
func ListaDiretorioExtensao(diretorio string, extensao string) ([]string, error) {
	files, err := filepath.Glob(diretorio + "*." + extensao)

	return files, err
}

// Converte String para Inteiro
func String2Int(s string) (int, error) {

	i, err := strconv.Atoi(s)
	return i, err

}

// Grava array de bytes em filesystem
// Entrada: Array de Bytes, nome de arquivo
// Saida : Arquivo gravado no sistema de arquivo
func GravarArquivo(post []byte, arquivo string) error {

	// output, err := xml.Marshal(&post)
	err := ioutil.WriteFile(arquivo, post, 0644)
	if err != nil {
		fmt.Println("Error writing  to file:", err)
		//		return
	}
	return err
}

// Condição de Checagem de erro
func CheckErr(err error) {
	if err != nil {
		log.Println("Erro Geral: ", err)

	}
}

// Log de Erro com Finalização
func LogErros(err error) {
	log.Fatal(err.Error())
}

// Changed to csvExport, as it doesn't make much sense to export things from
// package main
func CsvExport(data [][]string) error {
	file, err := os.Create("result.csv")
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, value := range data {
		if err := writer.Write(value); err != nil {
			return err // let's return errors if necessary, rather than having a one-size-fits-all error handler
		}
	}
	return nil
}

// Exporta em Arquivo CSV formato 2
func csvExport2(data [][]string) error {
	file, err := os.Create("result.csv")
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()
	writer.WriteAll(data)

	return nil
}

// Realiza copia de dois arquivos
func CopyFile(src, dst string) (int64, error) {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return 0, err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer destination.Close()
	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}

// Função auxiliar para acentuação
func isMn(r rune) bool {
	return unicode.Is(unicode.Mn, r) // Mn: nonspacing marks
}

// Remove acentuação de palavras
func RemoverAcentos(palavra string) string {
	t := transform.Chain(norm.NFD, transform.RemoveFunc(isMn), norm.NFC)
	result, _, err := transform.String(t, palavra)
	if err != nil {
		return palavra
	} else {
		return result
	}
}

// Executa programa, aguardando até a finalizacao
func ExecutarPrograma(prg string, arg ...string) error {
	fmt.Println("Executacao: ", prg, arg)
	cmd := exec.Command(prg, arg...)
	stdout, err := cmd.Output()

	// err := cmd.Start()
	if err != nil {
		fmt.Println("Error: ", err)
	}

	fmt.Print(string(stdout))

	return err

}

// exists returns whether the given file or directory exists
func Exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// converter de IO Reader para []bytes
func StreamToByte(stream io.Reader) []byte {
	buf := new(bytes.Buffer)
	buf.ReadFrom(stream)
	return buf.Bytes()
}

// Converte de Io Reader para string
func StreamToString(stream io.Reader) string {
	buf := new(bytes.Buffer)
	buf.ReadFrom(stream)
	return buf.String()
}

/*
ação: Recebe os nomes dos arquivos em pdf exportados no filesystem e gera um append dos arquivos
Apaga os arquivos ao utilizar
Entrada: Lista de Nome de Arquivos PDF exportados em Filesystem
Saida: Arquivo PDF salvo no disco
*/
func Merge(arquivos []string, retorno string) error {
	fmt.Println("Merge: ", arquivos, retorno)
	err := api.MergeCreateFile(arquivos, retorno, nil)
	//	err := api.MergeFile(arquivos, retorno, nil)

	for _, arquivo := range arquivos {
		fmt.Println("Arquivo:", arquivo)
		os.Remove(arquivo)

	}
	return err
}

// Diferença de Duas Datas String no Formato YYYYMMDD
func DiffDatasString(Inicio string, Final string) (int, error) {
	layout := "20060102"
	dataInicio, err := time.Parse(layout, Inicio)
	if err != nil {
		return 0, err
	}
	dataFinal, err := time.Parse(layout, Final)
	if err != nil {
		return 0, err

	}
	diff := math.Round(dataFinal.Sub(dataInicio).Hours() / 24)

	return int(diff), nil

}

// Diferença de Duas Datas String no Formato YYYYMMDD
func DiffDatas(dataInicio time.Time, dataFinal time.Time) int {

	diff := math.Round(dataFinal.Sub(dataInicio).Hours() / 24)

	return int(diff)

}

// Convert Data no padrão YYYYMMDD para golang time
func ConvertData(data string) (time.Time, error) {
	layout := "20060102"
	dataTime, err := time.Parse(layout, data)
	if err != nil {
		return dataTime, err
	}
	return dataTime, nil

}

func GerarUID() string {
	myuuid, err := uuid.NewV4()
	if err != nil {
		fmt.Println("Erro ao Gerar UID ", err)
		return ""

	}

	return myuuid.String()

}

func RemoveDuplicateStr(strSlice []string) []string {
	allKeys := make(map[string]bool)
	list := []string{}
	for _, item := range strSlice {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}

func RemoveDuplicateInt(intSlice []int) []int {
	allKeys := make(map[int]bool)
	list := []int{}
	for _, item := range intSlice {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}

// Retorna o caminho absoluto do aplicativo.
// Usado para executar o executavel fora do diretorio correte da apliação
func CaminhoAplicativo(arquivo string) string {
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	caminho := filepath.FromSlash(dir + "/" + arquivo)
	return caminho

}

// Formatar CPF, colocando os Pontos e traços
func CPFMask(cpf string) string {
	if len(cpf) != 11 {
		return ""
	}

	mask := fmt.Sprintf("%s.%s.%s-%s", cpf[0:3], cpf[3:6], cpf[6:9], cpf[9:11])
	return mask

}

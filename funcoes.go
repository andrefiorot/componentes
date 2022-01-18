package componentes

import (
	_ "bytes"
	_ "database/sql"
	"encoding/csv"
	_ "encoding/json"
	"fmt"
	"io"
	_ "io"
	"io/ioutil"
	"log"
	_ "net/http"
	"os"
	"os/exec"
	"strconv"
	_ "strconv"
	"unicode"

	_ "github.com/gocraft/dbr"
	_ "github.com/lib/pq"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

func LeituraArquivo(arquivo string) ([]byte, error) {
	jsonFile, err := os.Open(arquivo)
	if err != nil {
		fmt.Println("Error opening JSON file:", err)
		// return nil, err
	}

	defer jsonFile.Close()
	jsonData, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println("Error reading JSON data:", err)
		// return nil, err

	}

	return jsonData, err

}

func String2Int(s string) (int, error) {

	i, err := strconv.Atoi(s)
	return i, err

}

func GravarArquivo(post []byte, arquivo string) error {

	// output, err := xml.Marshal(&post)
	err := ioutil.WriteFile(arquivo, post, 0644)
	if err != nil {
		fmt.Println("Error writing  to file:", err)
		//		return
	}
	return err
}

func CheckErr(err error) {
	if err != nil {
		fmt.Println("Erro de Panic: ", err)
		panic(err)
	}
}

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

func isMn(r rune) bool {
	return unicode.Is(unicode.Mn, r) // Mn: nonspacing marks
}

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
	cmd := exec.Command(prg, arg...)
	stdout, err := cmd.Output()

	// err := cmd.Start()
	if err != nil {
		fmt.Println("Error: ", err)
	}
	/*
		if err != nil {
			fmt.Println(err.Error())

		}
	*/

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

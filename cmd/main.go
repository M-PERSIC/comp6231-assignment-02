package main

import (
	_ "embed"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/m-persic/comp6231-assignment-02/database"
	"github.com/m-persic/comp6231-assignment-02/fmp"
	"github.com/m-persic/comp6231-assignment-02/ftp"
)

//go:embed FMP.xlsx
var excelData []byte

func main() {
	fmpFlag := flag.Bool("fmp", false, "")
	ftpFlag := flag.Bool("ftp", false, "")
	flag.Parse()
	var microservice string
	if *fmpFlag && *ftpFlag {
		log.Fatal("Select either the FruitMonthPrice or FruitTotalPrice microservice!")
	} else if *ftpFlag {
		microservice = "ftp"
	} else {
		microservice = "fmp"
	}
	switch microservice {
	case "fmp":
		//FruitMonthPrice microservice
		fmpPort := os.Getenv("FMP_PORT")
		if fmpPort == "" {
			fmpPort = "8000"
		}
		db, err := database.InitDB()
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()
		err = database.LoadExcelData(db, excelData, "FruitPriceMonth", "fruit_prices")
		if err != nil {
			log.Fatal(err)
		}
		fmpService := fmp.NewService(db, "fruit_prices", fmpPort)
		err = fmp.StartFMPServer(fmpService, fmpPort)
		if err != nil {
			log.Fatal(err)
		}
	case "ftp":
		//FruitTotalPrice microservice
		ftpPort := os.Getenv("FTP_PORT")
		if ftpPort == "" {
			ftpPort = "8100"
		}
		fmpServiceUrl := os.Getenv("FMP_SERVICE_URL")
		if fmpServiceUrl == "" {
			fmpPort := os.Getenv("FMP_PORT")
			if fmpPort == "" {
				fmpPort = "8000"
			}
			fmpServiceUrl = fmt.Sprintf("http://localhost:%s", fmpPort)
		}
		ftpService := ftp.NewService(ftpPort, fmpServiceUrl)
		err := ftp.StartFTPServer(ftpService, ftpPort)
		if err != nil {
			log.Fatal(err)
		}
	default:
		log.Fatal("Incorrect flag (--fmp or --ftp only)!")
	}
}

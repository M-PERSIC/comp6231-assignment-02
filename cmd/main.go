package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/m-persic/comp6231-assignment-02/database"
	"github.com/m-persic/comp6231-assignment-02/fmp"
)

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
		fmp_port := os.Getenv("FMP_PORT")
		if fmp_port == "" {
			fmp_port = "8000"
		}
		db, err := database.InitDB()
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()
		err = database.LoadExcelFile(db, "data/FMP.xlsx", "FruitPriceMonth", "fruit_prices")
		if err != nil {
			log.Fatal(err)
		}
		fmp_svc := fmp.NewService(db, "fruit_prices", fmp_port)
		err = fmp.StartFMPServer(fmp_svc, fmp_port)
		if err != nil {
			log.Fatal(err)
		}
	case "ftp":
		fmt.Println("hello!")
	default:
		log.Fatal("Incorrect flag (--fmp or --ftp allowed)!")
	}

}

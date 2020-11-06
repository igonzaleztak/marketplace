package libs

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"

	// mysql Driver
	_ "github.com/go-sql-driver/mysql"
)

// StoreInDB stores the measurement in the database and returns the URL
// where is located.
func StoreInDB(configParams map[string]interface{}, measurement []byte, hash, iotName string) (string, error) {

	iotName = strings.ToLower(iotName)

	// The format of the URL to the database is:
	// 	user:password@tcp(dbHost)/dbName
	connParams := configParams["dbUsername"].(string) + ":" +
		configParams["dbPassword"].(string) + "@tcp(" +
		configParams["dbHost"].(string) + ")/" +
		configParams["dbName"].(string)

	// Connect to the MySQL database
	db, err := sql.Open("mysql", connParams)
	if err != nil {
		return "", err
	}

	defer db.Close()

	// Create table for the IoT producer
	query := fmt.Sprintf("CREATE TABLE %s(Hash CHAR(64), Measurement MEDIUMTEXT);", iotName)

	// Do query
	// Error 1050 means that the table has already been created,
	// so ignore the error message.
	_, err = db.Query(query)
	if err != nil {
		errorText := err.Error()
		errorCode := errorText[:10]
		if errorCode != "Error 1050" {
			return "", err
		}
	}

	// Check the measurement has not been stored yet
	query = fmt.Sprintf("SELECT Hash FROM %s WHERE Hash='%s';", iotName, hash)
	dbResp, err := db.Query(query)
	if err != nil {
		return "", err
	}
	// Returns True if there is a match in the DB.
	isStored := dbResp.Next()
	if isStored {
		return "", errors.New("The measurement has already been stored in the database")
	}

	// Insert the measurement in the Table
	query = fmt.Sprintf("INSERT INTO %s VALUES('%s', '%s')", iotName, hash, measurement)

	// Do Query
	_, err = db.Query(query)
	if err != nil {
		return "", err
	}

	fmt.Printf("\n%s: Stored in the Database\n", hash)

	// Create URL
	url := fmt.Sprintf("http://10.10.46.21:5051/measurements/%s/%s", iotName, hash)
	fmt.Printf("%s: Measurement can be found in %s", hash, url)

	return url, nil
}

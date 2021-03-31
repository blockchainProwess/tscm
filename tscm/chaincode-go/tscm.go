package main

import (
    "bytes"
    "encoding/json"
    "fmt"
    "strconv"
    "time"
    "strings"
    "github.com/hyperledger/fabric-chaincode-go/shim"
    sc "github.com/hyperledger/fabric-protos-go/peer"
)

// SmartContract struct
type SmartContract struct {
}

//ProducerTransaction struct
type ProducerTransaction struct {
    TxID                     string `json:"tx_id"`
    Stage                    int    `json:"stage_id"`
    AssetID                  string `json:"asset_id"`
    TotalWeightShippedPerLot string `json:"total_weight_shipped_per_lot"`
    FarmerDetails            string `json:"farmer_details"`
    DispatchDate             string `json:"pr_dispatch_date"`
    ArrivalDate              string `json:"pr_arrival_date"`
    ProducerName             string `json:"producer_name"`
	Location				 string `json:"location"`
}

//ProcessorTransaction struct
type ProcessorTransaction struct {
    TxID                	 string `json:"tx_id"`
    Stage               	 int    `json:"stage_id"`
    AssetID             	 string `json:"asset_id"`
    DispatchDate        	 string `json:"po_dispatch_date"`
    ArrivalDate         	 string `json:"po_arrival_date"`
    GinningDetails			 string `json:"ginning_details"`
	SpinningDetails			 string `json:"spinning_details"`
	WeavingDetails			 string `json:"weaving_details"`
	TotalWeightShippedPerLot string `json:"total_weight_shipped_per_lot"`
	ProcessorName       	 string `json:"processor_name"`
	Location				 string `json:"location"`
}



//ManufacturerTransaction struct
type ManufacturerTransaction struct {
    TxID                	 string `json:"tx_id"`
    Stage               	 int    `json:"stage_id"`
    AssetID             	 string `json:"asset_id"`
    DispatchDate        	 string `json:"mf_dispatch_date"`
    ArrivalDate         	 string `json:"mf_arrival_date"`
    DyeingDetails			 string `json:"ginning_details"`
	ProductName			 	 string `json:"product_name"`
	QCStatus				 string `json:"qc_status"`
	TotalPiecesShippedPerLot string `json:"total_pieces_shipped_per_lot"`
	ManufacturerName       	 string `json:"manufacturer_name"`
	Location				 string `json:"location"`
}


//DistributorTransaction struct
type DistributorTransaction struct {
    TxID                  string `json:"tx_id"`
    Stage                 int    `json:"stage_id"`
    AssetID               string `json:"asset_id"`
    DispatchDate          string `json:"di_dispatch_date"`
    ArrivalDate           string `json:"di_arrival_date"`
    DistributorName       string `json:"distributor_name"`
	Location			  string `json:"location"`
}

//RetailerTransaction struct
type RetailerTransaction struct {
    TxID          string `json:"tx_id"`
    Stage         int    `json:"stage_id"`
    AssetID       string `json:"asset_id"`
    ArrivalDate   string `json:"go_arrival_date"`
    RetailerName    string `json:"retailer_name"`
	Location      string `json:"location"`
}

//PrAsset map variable
var PrAsset map[string]string

//PoAsset map variable
var PoAsset map[string]string

//MfAsset map variable
var MfAsset map[string]string

// DiAsset map variable
var DiAsset map[string]string

// GoAsset map variable
var GoAsset map[string]string

func init() {
    PrAsset = make(map[string]string)
    PoAsset = make(map[string]string)
	MfAsset = make(map[string]string)
    DiAsset = make(map[string]string)
    GoAsset = make(map[string]string)

}

// Init function
func (s *SmartContract) Init(APIstub shim.ChaincodeStubInterface) sc.Response {
    return shim.Success(nil)
}

// Invoke function
func (s *SmartContract) Invoke(APIstub shim.ChaincodeStubInterface) sc.Response {

    // Retrieve the requested Smart Contract function and arguments
    function, args := APIstub.GetFunctionAndParameters()
    // Route to the appropriate handler function to interact with the ledger
    if function == "initLedger" {
        return s.initLedger(APIstub)
    } else if function == "recordProducerTransaction" {
        return s.recordProducerTransaction(APIstub, args)
    } else if function == "recordProcessorTransaction" {
        return s.recordProcessorTransaction(APIstub, args)
    }else if function == "recordManufacturerTransaction" {
        return s.recordManufacturerTransaction(APIstub, args)
    }else if function == "recordDistributorTransaction" {
        return s.recordDistributorTransaction(APIstub, args)
    } else if function == "recordRetailerTransaction" {
        return s.recordRetailerTransaction(APIstub, args)
    } else if function == "queryTransaction" {
        return s.queryTransaction(APIstub, args)
    } else if function == "queryTransactionHistory" {
        return s.queryTransactionHistory(APIstub, args)
    } else if function == "queryAllTransactions" {
        return s.queryAllTransactions(APIstub, args)
    } else if function == "queryAssetID" {
        return s.queryAssetID(APIstub, args)
    } else if function == "initAssetID" {
        return s.initAssetID(APIstub)
    }

    return shim.Error("Invalid Smart Contract function name.")
}

/*
 * The initLedger method *
Will add test data (10 asset catches)to our network
*/
func (s *SmartContract) initLedger(APIstub shim.ChaincodeStubInterface) sc.Response {

   
    return shim.Success(nil)
}

/*
 * The initLedger method *
Will add test data (10 asset catches)to our network
*/
func (s *SmartContract) initAssetID(APIstub shim.ChaincodeStubInterface) sc.Response {
   
    PrAsset["0"] = "0"
    asBytes, _ := json.Marshal(PrAsset)
    APIstub.PutState("PrAssetID", asBytes)

    PoAsset["0"] = "0"
    poasBytes, _ := json.Marshal(PoAsset)
    APIstub.PutState("PoAssetID", poasBytes)

	MfAsset["0"] = "0"
    mfasBytes, _ := json.Marshal(MfAsset)
    APIstub.PutState("MfAssetID", mfasBytes)

    DiAsset["0"] = "0"
    diasBytes, _ := json.Marshal(DiAsset)
    APIstub.PutState("DiAssetID", diasBytes)

    GoAsset["0"] = "0"
    goasBytes, _ := json.Marshal(GoAsset)
    APIstub.PutState("GoAssetID", goasBytes)

    return shim.Success(nil)
}

/*
 * The query method *
 */
func (s *SmartContract) queryAssetID(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

    if len(args) != 1 {
        return shim.Error("Incorrect number of arguments. Expecting 1")
    }

    assetAsBytes, _ := APIstub.GetState(args[0])
    if assetAsBytes == nil {
        return shim.Error("Could not locate asset")
    }

    return shim.Success(assetAsBytes)
}

/*
 * The recordProducerTransaction method *
 */
func (s *SmartContract) recordProducerTransaction(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

    if len(args) != 7 {
        return shim.Error("Incorrect number of arguments. Expecting 7")
    }
    stageID := 1

    var producertx = ProducerTransaction{TxID: args[0], Stage: stageID, AssetID: args[0], TotalWeightShippedPerLot: args[1], FarmerDetails: args[2], DispatchDate: args[3], ArrivalDate: args[4], ProducerName: args[5], Location: args[6]}

    producertxAsBytes, _ := json.Marshal(producertx)
    err := APIstub.PutState(args[0], producertxAsBytes)
    if err != nil {
        return shim.Error(fmt.Sprintf("Failed to record producertx catch: %s", args[0]))
    }

    prassetid := PrAsset
    assetAsBytes, _ := APIstub.GetState("PrAssetID")
    if assetAsBytes == nil {
        return shim.Error("Could not locate asset")
    }
    json.Unmarshal(assetAsBytes, &prassetid)
    asst_id := strings.Split(args[0],"_")
    key := asst_id[1]+"_"+asst_id[2][:6]
    PrAsset[key] = asst_id[2][6:]
    assetAsBytes, _ = json.Marshal(PrAsset)
    nerr := APIstub.PutState("PrAssetID", assetAsBytes)
    if nerr != nil {
        return shim.Error(fmt.Sprintf("Producer stage error: %s", args[0]))
    }


    return shim.Success(nil)
}



/*
 * The recordProcessorTransaction method *
 */
func (s *SmartContract) recordProcessorTransaction(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

    if len(args) != 10 {
        return shim.Error("Incorrect number of arguments. Expecting 10")
    }

    stageID := 2

    var processortx = ProcessorTransaction{TxID: args[0], Stage: stageID, AssetID: args[1], TotalWeightShippedPerLot: args[2], DispatchDate: args[3], ArrivalDate: args[4], GinningDetails: args[5], SpinningDetails: args[6], WeavingDetails: args[7], ProcessorName: args[8],Location: args[9]}

    processortxAsBytes, _ := json.Marshal(processortx)
    err := APIstub.PutState(args[0], processortxAsBytes)
    if err != nil {
        return shim.Error(fmt.Sprintf("Failed to record processortx catch: %s", args[0]))
    }

    poassetid := PoAsset
    assetAsBytes, _ := APIstub.GetState("PoAssetID")
    if assetAsBytes == nil {
        return shim.Error("Could not locate asset")
    }
    json.Unmarshal(assetAsBytes, &poassetid)
    asst_id := strings.Split(args[1],"_")
    key := asst_id[1]+"_"+asst_id[3][:6]
    PoAsset[key] = asst_id[3][6:]
   
    assetAsBytes, _ = json.Marshal(PoAsset)
    nerr := APIstub.PutState("PoAssetID", assetAsBytes)
    if nerr != nil {
        return shim.Error(fmt.Sprintf("Processor stage error: %s", args[1]))
    }



    return shim.Success(nil)
}






/*
 * The recordManufacturerTransaction method *
 */
 func (s *SmartContract) recordManufacturerTransaction(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

    if len(args) != 10 {
        return shim.Error("Incorrect number of arguments. Expecting 10")
    }

    stageID := 3

    var manufacturertx = ManufacturerTransaction{TxID: args[0], Stage: stageID, AssetID: args[1], DispatchDate: args[2], ArrivalDate: args[3], DyeingDetails: args[4], ProductName: args[5], QCStatus: args[6], TotalPiecesShippedPerLot: args[7], ManufacturerName: args[8], Location: args[9]}

    manufacturertxAsBytes, _ := json.Marshal(manufacturertx)
    err := APIstub.PutState(args[0], manufacturertxAsBytes)
    if err != nil {
        return shim.Error(fmt.Sprintf("Failed to record manufacturertx catch: %s", args[0]))
    }

    mfassetid := MfAsset
    assetAsBytes, _ := APIstub.GetState("MfAssetID")
    if assetAsBytes == nil {
        return shim.Error("Could not locate asset")
    }
    json.Unmarshal(assetAsBytes, &mfassetid)
    asst_id := strings.Split(args[1],"_")
    key := asst_id[1]+"_"+asst_id[4][:6]
    MfAsset[key] = asst_id[4][6:]
   
    assetAsBytes, _ = json.Marshal(MfAsset)
    nerr := APIstub.PutState("MfAssetID", assetAsBytes)
    if nerr != nil {
        return shim.Error(fmt.Sprintf("Manufacturer stage error: %s", args[1]))
    }



    return shim.Success(nil)
}



/*
 * The recordDistributorTransaction method *
 */
func (s *SmartContract) recordDistributorTransaction(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

    if len(args) != 6 {
        return shim.Error("Incorrect number of arguments. Expecting 6")
    }

    stageID := 4
    var distributortx = DistributorTransaction{TxID: args[0], Stage: stageID, AssetID: args[1], DispatchDate: args[2], ArrivalDate: args[3], DistributorName: args[4], Location: args[5]}

    distributortxAsBytes, _ := json.Marshal(distributortx)
    err := APIstub.PutState(args[0], distributortxAsBytes)
    if err != nil {
        return shim.Error(fmt.Sprintf("Failed to record distributortx catch: %s", args[0]))
    }

    diassetid := DiAsset
    assetAsBytes, _ := APIstub.GetState("DiAssetID")
    if assetAsBytes == nil {
        return shim.Error("Could not locate asset")
    }
    json.Unmarshal(assetAsBytes, &diassetid)
    asst_id := strings.Split(args[1],"_")
    key := asst_id[1]+"_"+asst_id[5][:6]
    DiAsset[key] = asst_id[5][6:]
   
    assetAsBytes, _ = json.Marshal(DiAsset)
    nerr := APIstub.PutState("DiAssetID", assetAsBytes)
    if nerr != nil {
        return shim.Error(fmt.Sprintf("Distributor stage error: %s", args[1]))
    }


    return shim.Success(nil)
}

/*
 * The recordRetailerTransaction method *
 */
func (s *SmartContract) recordRetailerTransaction(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

    if len(args) != 5 {
        return shim.Error("Incorrect number of arguments. Expecting 5")
    }

    stageID := 5
    var retailertx = RetailerTransaction{TxID: args[0], Stage: stageID, AssetID: args[1],ArrivalDate: args[2], RetailerName: args[3],Location:args[4]}

    retailertxAsBytes, _ := json.Marshal(retailertx)
    err := APIstub.PutState(args[0], retailertxAsBytes)
    if err != nil {
        return shim.Error(fmt.Sprintf("Failed to record retailertx catch: %s", args[0]))
    }

    goassetid := GoAsset
    assetAsBytes, _ := APIstub.GetState("GoAssetID")
    if assetAsBytes == nil {
        return shim.Error("Could not locate asset")
    }
    json.Unmarshal(assetAsBytes, &goassetid)
    asst_id := strings.Split(args[1],"_")
    key := asst_id[1]+"_"+asst_id[6][:6]
    GoAsset[key] = asst_id[6][6:]
    assetAsBytes, _ = json.Marshal(GoAsset)
    nerr := APIstub.PutState("GoAssetID", assetAsBytes)
    if nerr != nil {
        return shim.Error(fmt.Sprintf("Retailer stage error: %s", args[1]))
    }

    return shim.Success(nil)
}

/*
 * The queryAllTuna method *
 */
func (s *SmartContract) queryTransactionHistory(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

    if len(args) < 1 {
        return shim.Error("Incorrect number of arguments. Expecting 1")
    }

    assetID := args[0]

    fmt.Printf("- start getAsset History: %s\n", assetID)

    resultsIterator, err := APIstub.GetHistoryForKey(assetID)
    if err != nil {
        return shim.Error(err.Error())
    }
    defer resultsIterator.Close()

    // buffer is a JSON array containing historic values for the asset
    var buffer bytes.Buffer
    buffer.WriteString("[")

    bArrayMemberAlreadyWritten := false
    for resultsIterator.HasNext() {
        response, err := resultsIterator.Next()
        if err != nil {
            return shim.Error(err.Error())
        }
        // Add a comma before array members, suppress it for the first array member
        if bArrayMemberAlreadyWritten == true {
            buffer.WriteString(",")
        }
        buffer.WriteString("{\"TxId\":")
        buffer.WriteString("\"")
        buffer.WriteString(response.TxId)
        buffer.WriteString("\"")

        buffer.WriteString(", \"Value\":")
        // if it was a delete operation on given key, then we need to set the
        //corresponding value null. Else, we will write the response.Value
        //as-is (as the Value itself a JSON asset)
        if response.IsDelete {
            buffer.WriteString("null")
        } else {
            buffer.WriteString(string(response.Value))
        }

        buffer.WriteString(", \"Timestamp\":")
        buffer.WriteString("\"")
        buffer.WriteString(time.Unix(response.Timestamp.Seconds, int64(response.Timestamp.Nanos)).String())
        buffer.WriteString("\"")

        buffer.WriteString(", \"IsDelete\":")
        buffer.WriteString("\"")
        buffer.WriteString(strconv.FormatBool(response.IsDelete))
        buffer.WriteString("\"")

        buffer.WriteString("}")
        bArrayMemberAlreadyWritten = true
    }
    buffer.WriteString("]")

    fmt.Printf("- get Asset History returning:\n%s\n", buffer.String())

    return shim.Success(buffer.Bytes())

}

/*
 * The queryAllTransactions *
 */
func (s *SmartContract) queryAllTransactions(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

    if len(args) < 2 {
        return shim.Error("Incorrect number of arguments. Expecting 2")
    }

    startKey := ""
    endKey := ""

    resultsIterator, err := APIstub.GetStateByRange(startKey, endKey)
    if err != nil {
        return shim.Error(err.Error())
    }
    defer resultsIterator.Close()

    // buffer is a JSON array containing QueryResults
    var buffer bytes.Buffer
    buffer.WriteString("[")

    // bArrayMemberAlreadyWritten := false
    for resultsIterator.HasNext() {

        queryResponse, err := resultsIterator.Next()

        if err != nil {
            return shim.Error(err.Error())
        }

        assetID := queryResponse.Key

        fmt.Printf("- start get All Assets History: %s\n", assetID)

        resultsIterator, err := APIstub.GetHistoryForKey(assetID)
        if err != nil {
            return shim.Error(err.Error())
        }
        defer resultsIterator.Close()

        buffer.WriteString("[")

        bArrayMemberAlreadyWritten := false
        for resultsIterator.HasNext() {

            response, err := resultsIterator.Next()

            if err != nil {
                return shim.Error(err.Error())
            }
            // Add a comma before array members, suppress it for the first array member
            if bArrayMemberAlreadyWritten == true {
                buffer.WriteString(",")
            }
            buffer.WriteString("{\"TxId\":")
            buffer.WriteString("\"")
            buffer.WriteString(response.TxId)
            buffer.WriteString("\"")

            buffer.WriteString(", \"Value\":")
            // if it was a delete operation on given key, then we need to set the
            //corresponding value null. Else, we will write the response.Value
            //as-is (as the Value itself a JSON asset)
            if response.IsDelete {
                buffer.WriteString("null")
            } else {
                buffer.WriteString(string(response.Value))
            }

            buffer.WriteString(", \"Timestamp\":")
            buffer.WriteString("\"")
            buffer.WriteString(time.Unix(response.Timestamp.Seconds, int64(response.Timestamp.Nanos)).String())
            buffer.WriteString("\"")

            buffer.WriteString(", \"IsDelete\":")
            buffer.WriteString("\"")
            buffer.WriteString(strconv.FormatBool(response.IsDelete))
            buffer.WriteString("\"")

            buffer.WriteString("}")
            bArrayMemberAlreadyWritten = true
        }
        buffer.WriteString("],")
        // buffer.WriteString("]")
        // fmt.Printf("- get Asset History returning:\n%s\n", buffer.String())
        // bArrayMemberAlreadyWritten = true
    }
    buffer.WriteString("]")
    fmt.Printf("- get All Assets History returning:\n%s\n", buffer.String())
    return shim.Success(buffer.Bytes())

}

/*
 * The query method *
 */
func (s *SmartContract) queryTransaction(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

    if len(args) != 1 {
        return shim.Error("Incorrect number of arguments. Expecting 1")
    }

    assetAsBytes, _ := APIstub.GetState(args[0])
    if assetAsBytes == nil {
        return shim.Error("Could not locate asset")
    }
    return shim.Success(assetAsBytes)
}


func main() {

    // Create a new Smart Contract
    err := shim.Start(new(SmartContract))
    if err != nil {
        fmt.Printf("Error creating new Smart Contract: %s", err)
    }
}


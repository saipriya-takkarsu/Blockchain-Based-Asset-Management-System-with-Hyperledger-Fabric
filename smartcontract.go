package chaincode

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/v2/contractapi"
)

// SmartContract provides functions for managing an Asset
type SmartContract struct {
	contractapi.Contract
}

// Asset describes basic details of what makes up a simple asset
// Insert struct field in alphabetic order => to achieve determinism across languages
// golang keeps the order when marshal to json but doesn't order automatically
type Asset struct {
	DEALERID       int    `json:"DEALERID"`
	MSISDN         string `json:"MSISDN"`
	MPIN           string `json:"ID"`
	BALANCE        int    `json:"BALANCE"`
	STATUS         string `json:"Status"`
	TRANSAMOUNT    int    `json:"TRANSAM"`
	Transtype      string `json:"TRANS"`
	REMARKS        string `json:"REMARKS"`
}

// InitLedger adds a base set of assets to the ledger
func (s *SmartContract) InitLedger(ctx contractapi.TransactionContextInterface) error {
    assets := []Asset{
        {DEALERID: 1301, MSISDN: "+91 111111", MPIN: "1111", BALANCE: 50, STATUS: "SUCCESS", TRANSAMOUNT: 100, TRANSTYPE: "ONLINE", REMARKS: "ACTIVE"},
        {DEALERID: 1302, MSISDN: "+91 222222", MPIN: "2222", BALANCE: 200, STATUS: "FAILURE", TRANSAMOUNT: 150, TRANSTYPE: "OFFLINE", REMARKS: "INACTIVE"},
        {DEALERID: 1303, MSISDN: "+91 333333", MPIN: "3333", BALANCE: 400, STATUS: "SUCCESS", TRANSAMOUNT: 300, TRANSTYPE: "ONLINE", REMARKS: "ACTIVE"},
        {DEALERID: 1304, MSISDN: "+91 444444", MPIN: "4444", BALANCE: 250, STATUS: "SUCCESS", TRANSAMOUNT: 200, TRANSTYPE: "OFFLINE", REMARKS: "ACTIVE"},
        {DEALERID: 1305, MSISDN: "+91 555555", MPIN: "5555", BALANCE: 350, STATUS: "FAILURE", TRANSAMOUNT: 500, TRANSTYPE: "ONLINE", REMARKS: "INACTIVE"},
        {DEALERID: 1306, MSISDN: "+91 666666", MPIN: "6666", BALANCE: 450, STATUS: "SUCCESS", TRANSAMOUNT: 600, TRANSTYPE: "OFFLINE", REMARKS: "ACTIVE"},

    }

    for _, asset := range assets {
        assetJSON, err := json.Marshal(asset)
        if err != nil {
            return err
        }

        err = ctx.GetStub().PutState(strconv.Itoa(asset.DEALERID), assetJSON)
        if err != nil {
            return fmt.Errorf("failed to put to world state: %v", err)
        }
    }

    return nil
}

// CreateAsset adds a new asset to the ledger
func (s *SmartContract) CreateAsset(ctx contractapi.TransactionContextInterface, dealerID int, msisdn, mpin string, balance, transAmount int, status, transType, remarks string) error {
    asset := Asset{
        DEALERID:    dealerID,
        MSISDN:      msisdn,
        MPIN:        mpin,
        BALANCE:     balance,
        STATUS:      status,
        TRANSAMOUNT: transAmount,
        TRANSTYPE:   transType,
        REMARKS:     remarks,
    }

    assetJSON, err := json.Marshal(asset)
    if err != nil {
        return err
    }

    return ctx.GetStub().PutState(strconv.Itoa(dealerID), assetJSON)
}

// UpdateAsset modifies an existing asset in the ledger
func (s *SmartContract) UpdateAsset(ctx contractapi.TransactionContextInterface, dealerID int, msisdn, mpin string, balance, transAmount int, status, transType, remarks string) error {
    assetJSON, err := ctx.GetStub().GetState(strconv.Itoa(dealerID))
    if err != nil {
        return fmt.Errorf("failed to read asset: %v", err)
    }
    if assetJSON == nil {
        return fmt.Errorf("asset %d does not exist", dealerID)
    }

    var asset Asset
    err = json.Unmarshal(assetJSON, &asset)
    if err != nil {
        return err
    }

    asset.MSISDN = msisdn
    asset.MPIN = mpin
    asset.BALANCE = balance
    asset.STATUS = status
    asset.TRANSAMOUNT = transAmount
    asset.TRANSTYPE = transType
    asset.REMARKS = remarks

    updatedAssetJSON, err := json.Marshal(asset)
    if err != nil {
        return err
    }

    return ctx.GetStub().PutState(strconv.Itoa(dealerID), updatedAssetJSON)
}

// ReadAsset retrieves an asset from the ledger
func (s *SmartContract) ReadAsset(ctx contractapi.TransactionContextInterface, dealerID int) (*Asset, error) {
    assetJSON, err := ctx.GetStub().GetState(strconv.Itoa(dealerID))
    if err != nil {
        return nil, fmt.Errorf("failed to read asset: %v", err)
    }
    if assetJSON == nil {
        return nil, fmt.Errorf("asset %d does not exist", dealerID)
    }

    var asset Asset
    err = json.Unmarshal(assetJSON, &asset)
    if err != nil {
        return nil, err
    }

    return &asset, nil
}

// GetAllAssets retrieves all assets from the ledger
func (s *SmartContract) GetAllAssets(ctx contractapi.TransactionContextInterface) ([]*Asset, error) {
    iterator, err := ctx.GetStub().GetStateByRange("", "")
    if err != nil {
        return nil, err
    }
    defer iterator.Close()

    var assets []*Asset
    for iterator.HasNext() {
        queryResponse, err := iterator.Next()
        if err != nil {
            return nil, err
        }

        var asset Asset
        err = json.Unmarshal(queryResponse.Value, &asset)
        if err != nil {
            return nil, err
        }
        assets = append(assets, &asset)
    }

    return assets, nil
}

// GetAssetHistory retrieves the history of an asset's transactions
func (s *SmartContract) GetAssetHistory(ctx contractapi.TransactionContextInterface, dealerID int) ([]*Asset, error) {
    historyIterator, err := ctx.GetStub().GetHistoryForKey(strconv.Itoa(dealerID))
    if err != nil {
        return nil, fmt.Errorf("failed to retrieve asset history: %v", err)
    }
    defer historyIterator.Close()

    var assetHistory []*Asset
    for historyIterator.HasNext() {
        historyData, err := historyIterator.Next()
        if err != nil {
            return nil, err
        }

        var asset Asset
        if historyData.Value != nil {
            err = json.Unmarshal(historyData.Value, &asset)
            if err != nil {
                return nil, err
            }
            assetHistory = append(assetHistory, &asset)
        }
    }

    return assetHistory, nil
}

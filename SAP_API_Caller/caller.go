package sap_api_caller

import (
	"fmt"
	"io/ioutil"
	"net/http"
	sap_api_output_formatter "sap-api-integrations-supplier-quotation-reads/SAP_API_Output_Formatter"
	"strings"
	"sync"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	"golang.org/x/xerrors"
)

type SAPAPICaller struct {
	baseURL string
	apiKey  string
	log     *logger.Logger
}

func NewSAPAPICaller(baseUrl string, l *logger.Logger) *SAPAPICaller {
	return &SAPAPICaller{
		baseURL: baseUrl,
		apiKey:  GetApiKey(),
		log:     l,
	}
}

func (c *SAPAPICaller) AsyncGetSupplierQuotation(supplierQuotation, supplierQuotationItem string, accepter []string) {
	wg := &sync.WaitGroup{}
	wg.Add(len(accepter))
	for _, fn := range accepter {
		switch fn {
		case "Header":
			func() {
				c.Header(supplierQuotation)
				wg.Done()
			}()
		case "Item":
			func() {
				c.Item(supplierQuotation, supplierQuotationItem)
				wg.Done()
			}()
		default:
			wg.Done()
		}
	}
	wg.Wait()
}

func (c *SAPAPICaller) Header(supplierQuotation string) {
	data, err := c.callWarehouseSrvAPIRequirementHeader("SupplierQuotation", supplierQuotation)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(data)

}

func (c *SAPAPICaller) callWarehouseSrvAPIRequirementHeader(api, supplierQuotation string) ([]sap_api_output_formatter.Header, error) {
	url := strings.Join([]string{c.baseURL, "api_supplierquotation_2/srvd_a2x/sap/supplierquotation/0001", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	c.setHeaderAPIKeyAccept(req)
	c.getQueryWithHeader(req, supplierQuotation)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToHeader(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) Item(supplierQuotation, supplierQuotationItem string) {
	data, err := c.callWarehouseSrvAPIRequirementItem("SupplierQuotation", supplierQuotation, supplierQuotationItem)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(data)

}

func (c *SAPAPICaller) callWarehouseSrvAPIRequirementItem(api, supplierQuotation, supplierQuotationItem string) ([]sap_api_output_formatter.Item, error) {
	url := strings.Join([]string{c.baseURL, "api_supplierquotation_2/srvd_a2x/sap/supplierquotation/0001", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	c.setHeaderAPIKeyAccept(req)
	c.getQueryWithItem(req, supplierQuotation, supplierQuotationItem)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToItem(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) setHeaderAPIKeyAccept(req *http.Request) {
	req.Header.Set("APIKey", c.apiKey)
	req.Header.Set("Accept", "application/json")
}

func (c *SAPAPICaller) getQueryWithHeader(req *http.Request, supplierQuotation string) {
	params := req.URL.Query()
	params.Add("$filter", fmt.Sprintf("SupplierQuotation eq '%s'", supplierQuotation))
	req.URL.RawQuery = params.Encode()
}

func (c *SAPAPICaller) getQueryWithItem(req *http.Request, supplierQuotation, supplierQuotationItem string) {
	params := req.URL.Query()
	params.Add("$filter", fmt.Sprintf("SupplierQuotation eq '%s' and SupplierQuotationItem eq '%s'", supplierQuotation, supplierQuotationItem))
	req.URL.RawQuery = params.Encode()
}

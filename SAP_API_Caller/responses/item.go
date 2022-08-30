package responses

type Item struct {
	Value []struct {
		SupplierQuotation              string `json:"SupplierQuotation"`
		SupplierQuotationItem          string `json:"SupplierQuotationItem"`
		PurchasingDocumentCategory     string `json:"PurchasingDocumentCategory"`
		PurchasingDocumentItemText     string `json:"PurchasingDocumentItemText"`
		Material                       string `json:"Material"`
		ProductTypeCode                string `json:"ProductTypeCode"`
		ManufacturerMaterial           string `json:"ManufacturerMaterial"`
		SupplierMaterialNumber         string `json:"SupplierMaterialNumber"`
		ManufacturerPartNmbr           string `json:"ManufacturerPartNmbr"`
		Manufacturer                   string `json:"Manufacturer"`
		MaterialGroup                  string `json:"MaterialGroup"`
		Plant                          string `json:"Plant"`
		ManualDeliveryAddressID        string `json:"ManualDeliveryAddressID"`
		ReferenceDeliveryAddressID     string `json:"ReferenceDeliveryAddressID"`
		AddressID                      string `json:"AddressID"`
		ItemDeliveryAddressID          string `json:"ItemDeliveryAddressID"`
		IncotermsClassification        string `json:"IncotermsClassification"`
		IncotermsTransferLocation      string `json:"IncotermsTransferLocation"`
		IncotermsLocation1             string `json:"IncotermsLocation1"`
		IncotermsLocation2             string `json:"IncotermsLocation2"`
		ScheduleLineDeliveryDate       string `json:"ScheduleLineDeliveryDate"`
		ScheduleLineOrderQuantity      int    `json:"ScheduleLineOrderQuantity"`
		AwardedQuantity                int    `json:"AwardedQuantity"`
		PerformancePeriodStartDate     string `json:"PerformancePeriodStartDate"`
		PerformancePeriodEndDate       string `json:"PerformancePeriodEndDate"`
		OrderPriceUnit                 string `json:"OrderPriceUnit"`
		OrderPriceUnitToOrderUnitNmrtr int    `json:"OrderPriceUnitToOrderUnitNmrtr"`
		OrdPriceUnitToOrderUnitDnmntr  int    `json:"OrdPriceUnitToOrderUnitDnmntr"`
		OrderQuantityUnit              string `json:"OrderQuantityUnit"`
		OrderItemQtyToBaseQtyNmrtr     int    `json:"OrderItemQtyToBaseQtyNmrtr"`
		OrderItemQtyToBaseQtyDnmntr    int    `json:"OrderItemQtyToBaseQtyDnmntr"`
		PurgDocPriceDate               string `json:"PurgDocPriceDate"`
		BaseUnit                       string `json:"BaseUnit"`
		NetAmount                      int    `json:"NetAmount"`
		GrossAmount                    int    `json:"GrossAmount"`
		EffectiveAmount                int    `json:"EffectiveAmount"`
		NetPriceAmount                 int    `json:"NetPriceAmount"`
		NetPriceQuantity               int    `json:"NetPriceQuantity"`
		DocumentCurrency               string `json:"DocumentCurrency"`
		PurchaseRequisition            string `json:"PurchaseRequisition"`
		PurchaseRequisitionItem        string `json:"PurchaseRequisitionItem"`
		RequestForQuotation            string `json:"RequestForQuotation"`
		RequestForQuotationItem        string `json:"RequestForQuotationItem"`
		PurchasingInfoRecordUpdateCode string `json:"PurchasingInfoRecordUpdateCode"`
		PurchasingInfoRecord           string `json:"PurchasingInfoRecord"`
		PurchasingDocumentItemCategory string `json:"PurchasingDocumentItemCategory"`
		LastChangeDateTime             string `json:"LastChangeDateTime"`
	} `json:"value"`
}

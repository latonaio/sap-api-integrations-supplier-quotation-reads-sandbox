package sap_api_output_formatter

import (
	"encoding/json"
	"sap-api-integrations-supplier-quotation-reads/SAP_API_Caller/responses"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	"golang.org/x/xerrors"
)

func ConvertToHeader(raw []byte, l *logger.Logger) ([]Header, error) {
	pm := &responses.Header{}
	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to Header. unmarshal error: %w", err)
	}
	if len(pm.Value) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.Value) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.Value))
	}

	header := make([]Header, 0, 10)
	for i := 0; i < 10 && i < len(pm.Value); i++ {
		data := pm.Value[i]
		header = append(header, Header{
			SupplierQuotation:              data.SupplierQuotation,
			CompanyCode:                    data.CompanyCode,
			PurchasingDocumentCategory:     data.PurchasingDocumentCategory,
			PurchasingDocumentType:         data.PurchasingDocumentType,
			Supplier:                       data.Supplier,
			CreatedByUser:                  data.CreatedByUser,
			CreationDate:                   data.CreationDate,
			Language:                       data.Language,
			DocumentCurrency:               data.DocumentCurrency,
			IncotermsClassification:        data.IncotermsClassification,
			IncotermsTransferLocation:      data.IncotermsTransferLocation,
			IncotermsVersion:               data.IncotermsVersion,
			IncotermsLocation1:             data.IncotermsLocation1,
			IncotermsLocation2:             data.IncotermsLocation2,
			PaymentTerms:                   data.PaymentTerms,
			CashDiscount1Days:              data.CashDiscount1Days,
			CashDiscount2Days:              data.CashDiscount2Days,
			CashDiscount1Percent:           data.CashDiscount1Percent,
			CashDiscount2Percent:           data.CashDiscount2Percent,
			NetPaymentDays:                 data.NetPaymentDays,
			PricingProcedure:               data.PricingProcedure,
			PurchasingOrganization:         data.PurchasingOrganization,
			PurchasingGroup:                data.PurchasingGroup,
			PurchasingDocumentOrderDate:    data.PurchasingDocumentOrderDate,
			AbsoluteExchangeRate:           data.AbsoluteExchangeRate,
			ExchRateIsIndirectQuotation:    data.ExchRateIsIndirectQuotation,
			EffectiveExchangeRate:          data.EffectiveExchangeRate,
			ExchangeRateIsFixed:            data.ExchangeRateIsFixed,
			PurContrValidityStartDate:      data.PurContrValidityStartDate,
			PurContrValidityEndDate:        data.PurContrValidityEndDate,
			IsEndOfPurposeBlocked:          data.IsEndOfPurposeBlocked,
			PurchasingDocumentDeletionCode: data.PurchasingDocumentDeletionCode,
			RequestForQuotation:            data.RequestForQuotation,
			SupplierQuotationExternalID:    data.SupplierQuotationExternalID,
			QuotationSubmissionDate:        data.QuotationSubmissionDate,
			QuotationLatestSubmissionDate:  data.QuotationLatestSubmissionDate,
			BindingPeriodValidityEndDate:   data.BindingPeriodValidityEndDate,
			QtnLifecycleStatus:             data.QtnLifecycleStatus,
			FollowOnDocumentCategory:       data.FollowOnDocumentCategory,
			PurgDocFollowOnDocumentType:    data.PurgDocFollowOnDocumentType,
		})
	}

	return header, nil
}

func ConvertToItem(raw []byte, l *logger.Logger) ([]Item, error) {
	pm := &responses.Item{}
	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to Item. unmarshal error: %w", err)
	}
	if len(pm.Value) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.Value) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.Value))
	}

	item := make([]Item, 0, 10)
	for i := 0; i < 10 && i < len(pm.Value); i++ {
		data := pm.Value[i]
		item = append(item, Item{
			SupplierQuotation:              data.SupplierQuotation,
			SupplierQuotationItem:          data.SupplierQuotationItem,
			PurchasingDocumentCategory:     data.PurchasingDocumentCategory,
			PurchasingDocumentItemText:     data.PurchasingDocumentItemText,
			Material:                       data.Material,
			ProductTypeCode:                data.ProductTypeCode,
			ManufacturerMaterial:           data.ManufacturerMaterial,
			SupplierMaterialNumber:         data.SupplierMaterialNumber,
			ManufacturerPartNmbr:           data.ManufacturerPartNmbr,
			Manufacturer:                   data.Manufacturer,
			MaterialGroup:                  data.MaterialGroup,
			Plant:                          data.Plant,
			ManualDeliveryAddressID:        data.ManualDeliveryAddressID,
			ReferenceDeliveryAddressID:     data.ReferenceDeliveryAddressID,
			AddressID:                      data.AddressID,
			ItemDeliveryAddressID:          data.ItemDeliveryAddressID,
			IncotermsClassification:        data.IncotermsClassification,
			IncotermsTransferLocation:      data.IncotermsTransferLocation,
			IncotermsLocation1:             data.IncotermsLocation1,
			IncotermsLocation2:             data.IncotermsLocation2,
			ScheduleLineDeliveryDate:       data.ScheduleLineDeliveryDate,
			ScheduleLineOrderQuantity:      data.ScheduleLineOrderQuantity,
			AwardedQuantity:                data.AwardedQuantity,
			PerformancePeriodStartDate:     data.PerformancePeriodStartDate,
			PerformancePeriodEndDate:       data.PerformancePeriodEndDate,
			OrderPriceUnit:                 data.OrderPriceUnit,
			OrderPriceUnitToOrderUnitNmrtr: data.OrderPriceUnitToOrderUnitNmrtr,
			OrdPriceUnitToOrderUnitDnmntr:  data.OrdPriceUnitToOrderUnitDnmntr,
			OrderQuantityUnit:              data.OrderQuantityUnit,
			OrderItemQtyToBaseQtyNmrtr:     data.OrderItemQtyToBaseQtyNmrtr,
			OrderItemQtyToBaseQtyDnmntr:    data.OrderItemQtyToBaseQtyDnmntr,
			PurgDocPriceDate:               data.PurgDocPriceDate,
			BaseUnit:                       data.BaseUnit,
			NetAmount:                      data.NetAmount,
			GrossAmount:                    data.GrossAmount,
			EffectiveAmount:                data.EffectiveAmount,
			NetPriceAmount:                 data.NetPriceAmount,
			NetPriceQuantity:               data.NetPriceQuantity,
			DocumentCurrency:               data.DocumentCurrency,
			PurchaseRequisition:            data.PurchaseRequisition,
			PurchaseRequisitionItem:        data.PurchaseRequisitionItem,
			RequestForQuotation:            data.RequestForQuotation,
			RequestForQuotationItem:        data.RequestForQuotationItem,
			PurchasingInfoRecordUpdateCode: data.PurchasingInfoRecordUpdateCode,
			PurchasingInfoRecord:           data.PurchasingInfoRecord,
			PurchasingDocumentItemCategory: data.PurchasingDocumentItemCategory,
			LastChangeDateTime:             data.LastChangeDateTime,
		})
	}

	return item, nil
}

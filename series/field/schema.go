package field

import (
	"github.com/lindb/lindb/aggregation/function"
)

const SimpleFieldPFieldID = PrimitiveID(1)

// Schema represents the field schema internal definition
type Schema interface {
	// GetAggFunc gets agg func type by primitive field id
	GetAggFunc(pFieldID PrimitiveID) AggFunc
	// GetAllPrimitiveFields returns all primitive field ids, sort by field id
	GetAllPrimitiveFields() []PrimitiveID
	// getPrimitiveFields gets need extract primitive fields
	getPrimitiveFields(funcType function.FuncType) PrimitiveFields
	// getDefaultPrimitiveFields gets the default extract primitive fields
	getDefaultPrimitiveFields() PrimitiveFields
}

type sumSchema struct {
	primitiveFieldID PrimitiveID
	fieldIDs         []PrimitiveID
}

func newSumSchema() Schema {
	return &sumSchema{
		primitiveFieldID: SimpleFieldPFieldID,
		fieldIDs:         []PrimitiveID{SimpleFieldPFieldID},
	}
}

func (s *sumSchema) GetAggFunc(pFieldID PrimitiveID) AggFunc {
	return sumAggregator
}

func (s *sumSchema) GetAllPrimitiveFields() []PrimitiveID {
	return s.fieldIDs
}

func (s *sumSchema) getPrimitiveFields(funcType function.FuncType) PrimitiveFields {
	switch funcType {
	case function.Sum:
		return s.getDefaultPrimitiveFields()
	default:
		return nil
	}
}

func (s *sumSchema) getDefaultPrimitiveFields() PrimitiveFields {
	return PrimitiveFields{
		{FieldID: s.primitiveFieldID, AggType: Sum},
	}
}

type minSchema struct {
	primitiveFieldID PrimitiveID
	fieldIDs         []PrimitiveID
}

func newMinSchema() Schema {
	return &minSchema{
		primitiveFieldID: SimpleFieldPFieldID,
		fieldIDs:         []PrimitiveID{SimpleFieldPFieldID},
	}
}

func (s *minSchema) GetAggFunc(pFieldID PrimitiveID) AggFunc {
	return minAggregator
}

func (s *minSchema) GetAllPrimitiveFields() []PrimitiveID {
	return s.fieldIDs
}

func (s *minSchema) getPrimitiveFields(funcType function.FuncType) PrimitiveFields {
	switch funcType {
	case function.Min:
		return s.getDefaultPrimitiveFields()
	default:
		return nil
	}
}

func (s *minSchema) getDefaultPrimitiveFields() PrimitiveFields {
	return PrimitiveFields{
		{FieldID: s.primitiveFieldID, AggType: Min},
	}
}

type maxSchema struct {
	primitiveFieldID PrimitiveID
	fieldIDs         []PrimitiveID
}

func newMaxSchema() Schema {
	return &maxSchema{
		primitiveFieldID: SimpleFieldPFieldID,
		fieldIDs:         []PrimitiveID{SimpleFieldPFieldID},
	}
}

func (s *maxSchema) GetAggFunc(pFieldID PrimitiveID) AggFunc {
	return maxAggregator
}

func (s *maxSchema) GetAllPrimitiveFields() []PrimitiveID {
	return s.fieldIDs
}

func (s *maxSchema) getPrimitiveFields(funcType function.FuncType) PrimitiveFields {
	switch funcType {
	case function.Max:
		return s.getDefaultPrimitiveFields()
	default:
		return nil
	}
}

func (s *maxSchema) getDefaultPrimitiveFields() PrimitiveFields {
	return PrimitiveFields{
		{FieldID: s.primitiveFieldID, AggType: Max},
	}
}

type gaugeSchema struct {
	primitiveFieldID PrimitiveID
	fieldIDs         []PrimitiveID
}

func newGaugeSchema() Schema {
	return &gaugeSchema{
		primitiveFieldID: SimpleFieldPFieldID,
		fieldIDs:         []PrimitiveID{SimpleFieldPFieldID},
	}
}

func (s *gaugeSchema) GetAggFunc(pFieldID PrimitiveID) AggFunc {
	return replaceAggregator
}

func (s *gaugeSchema) GetAllPrimitiveFields() []PrimitiveID {
	return s.fieldIDs
}

func (s *gaugeSchema) getPrimitiveFields(funcType function.FuncType) PrimitiveFields {
	switch funcType {
	case function.Replace:
		return s.getDefaultPrimitiveFields()
	default:
		return nil
	}
}

func (s *gaugeSchema) getDefaultPrimitiveFields() PrimitiveFields {
	return PrimitiveFields{
		{FieldID: s.primitiveFieldID, AggType: Replace},
	}
}

type summarySchema struct {
	sumFieldID, countFieldID, minFieldID, maxFieldID PrimitiveID
	fieldIDs                                         []PrimitiveID
}

func newSummarySchema() Schema {
	return &summarySchema{
		sumFieldID:   PrimitiveID(1),
		countFieldID: PrimitiveID(2),
		maxFieldID:   PrimitiveID(3),
		minFieldID:   PrimitiveID(4),
		fieldIDs:     []PrimitiveID{1, 2, 3, 4},
	}
}

func (s *summarySchema) GetAggFunc(pFieldID PrimitiveID) AggFunc {
	switch pFieldID {
	case PrimitiveID(1), PrimitiveID(2):
		return sumAggregator
	case PrimitiveID(3):
		return maxAggregator
	case PrimitiveID(4):
		return minAggregator
	default:
		return replaceAggregator
	}
}

func (s *summarySchema) GetAllPrimitiveFields() []PrimitiveID {
	return s.fieldIDs
}

func (s *summarySchema) getPrimitiveFields(funcType function.FuncType) PrimitiveFields {
	switch funcType {
	case function.Sum:
		return PrimitiveFields{
			{FieldID: s.sumFieldID, AggType: Sum},
		}
	case function.Min:
		return PrimitiveFields{
			{FieldID: s.minFieldID, AggType: Min},
		}
	case function.Max:
		return PrimitiveFields{
			{FieldID: s.maxFieldID, AggType: Max},
		}
	case function.Count:
		return PrimitiveFields{
			{FieldID: s.countFieldID, AggType: Sum},
		}
	case function.Avg:
		return PrimitiveFields{
			{FieldID: s.sumFieldID, AggType: Sum},
			{FieldID: s.countFieldID, AggType: Sum},
		}
	default:
		return nil
	}
}

func (s *summarySchema) getDefaultPrimitiveFields() PrimitiveFields {
	return PrimitiveFields{
		{FieldID: s.countFieldID, AggType: Sum},
	}
}

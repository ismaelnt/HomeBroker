package entity

type Investor struct {
	ID             string
	Name           string
	AssetPositions []*InvestorAssetPosition
}

func NewInvestor(id string) *Investor {
	return &Investor{
		ID:             id,
		AssetPositions: []*InvestorAssetPosition{},
	}
}

func (i *Investor) GetAssetPosition(assetID string) *InvestorAssetPosition {
	for _, assetPosition := range i.AssetPositions {
		if assetPosition.AssetID == assetID {
			return assetPosition
		}
	}
	return nil
}

func (i *Investor) AddAssetPosition(assetPosition *InvestorAssetPosition) {
	i.AssetPositions = append(i.AssetPositions, assetPosition)
}

func (i *Investor) UpdateAssetPosition(AssetID string, Shares int) {
	assetPosition := i.GetAssetPosition(AssetID)
	if assetPosition == nil {
		i.AssetPositions = append(i.AssetPositions, NewInvestorAssetPosition(AssetID, Shares))
	} else {
		assetPosition.Shares += Shares
	}
}

type InvestorAssetPosition struct {
	AssetID string
	Shares  int
}

func NewInvestorAssetPosition(assetID string, shares int) *InvestorAssetPosition {
	return &InvestorAssetPosition{
		AssetID: assetID,
		Shares:  shares,
	}
}

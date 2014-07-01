package openrtb

import (
	"errors"
)

// At least one of Bid is required.
// A bid response can contain multiple “seatbid” objects, each on behalf of a different bidder seat.
// Seatbid object can contain multiple bids each pertaining to a different impression on behalf of a seat.
// Each “bid” object must include the impression ID to which it pertains as well as the bid price.
// Group attribute can be used to specify if a seat is willing to accept any impressions that it can win (default) or if it is
// only interested in winning any if it can win them all (i.e., all or nothing).
type Seatbid struct {
	Bid   []Bid      `json:"bid"`             // Array of bid objects; each realtes to an imp, if exchange supported can have many bid objects.
	Seat  *string    `json:"seat,omitempty"`  // ID of the bidder seat optional string ID of the bidder seat on whose behalf this bid is made.
	Group *int       `json:"group,omitempty"` // '1' means impression must be won-lost as a group; default is '0'.
	Ext   Extensions `json:"ext,omitempty"`
}

// Validation errors
var (
	ErrInvalidSeatbidBid = errors.New("openrtb response: seatbid is missing bids")
)

// Validate Seatbid required attributes
func (sb *Seatbid) Valid() (bool, error) {

	if sb.Bid == nil || len(sb.Bid) < 1 {
		return false, ErrInvalidSeatbidBid
	}

	for _, bid := range sb.Bid {
		if ok, err := bid.Valid(); !ok {
			return ok, err
		}
	}

	return true, nil
}

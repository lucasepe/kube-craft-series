package v1alpha1

import (
	kubelibapis "github.com/lucasepe/kubelib/apis"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// SpotPriceSpec defines the desired state of a SpotPrice
type SpotPriceSpec struct {
	// Symbol of the Metal; one of: AU, AG, PT, PD, RH
	//+kubebuilder:validation:Enum=AU;AG;PT;PD;RH
	Symbol string `json:"symbol"`
	// Currency of this spot price
	// One of: USD, AUD, CAD, EUR, GBP, JPY, CHF,
	//         CNY, HKD, BRL, INR, MXN, RUB, ZAR
	//+kubebuilder:validation:Enum=USD;AUD;CAD;EUR;GBP;JPY;CHF;CNY;HKD;BRL;INR;MXN;RUB;ZAR
	Currency string `json:"currency"`
	// Gold weight measurement; one of: oz, gram, kilo, tola
	//+kubebuilder:validation:Enum=oz;gram;kilo;tola
	Unit string `json:"unit"`
}

// SpotPriceStatus defines the observed state of a SpotPrice
type SpotPriceStatus struct {
	kubelibapis.Conditioned `json:",inline"`

	Date metav1.Time `json:"date"`
	// Bid is the price at which you can SELL.
	Bid resource.Quantity `json:"bid"`
	// Ask is the price at which you can BUY.
	Ask resource.Quantity `json:"ask"`
	// The change value indicates the difference between the current price and the previous closing price.
	// It represents the price movement, either up or down, since the last closing of the market.
	Change resource.Quantity `json:"change"`
}

// +genclient
//
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:shortName=msp;sp,scope=Namespaced,categories=metals
// +kubebuilder:printcolumn:JSONPath=".spec.symbol",name="Symbol",type="string"
// +kubebuilder:printcolumn:JSONPath=".spec.currency",name="Currency",type="string"
// +kubebuilder:printcolumn:JSONPath=".spec.unit",name="Unit",type="string"
// +kubebuilder:printcolumn:JSONPath=".status.bid",name="Bid",type="string",format="float",priority=1
// +kubebuilder:printcolumn:JSONPath=".status.ask",name="Ask",type="string",format="float",priority=1
// +kubebuilder:printcolumn:JSONPath=".status.change",name="Change",type="string",format="float",priority=1
// +kubebuilder:printcolumn:name="Updated",type="date",JSONPath=".status.date",priority=1
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status",priority=1

// SpotPrice is the Schema for the SpotPrice API
// SpotPrice is the price at which a metal may be bought and sold right now.
type SpotPrice struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SpotPriceSpec   `json:"spec,omitempty"`
	Status SpotPriceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SpotPriceList contains a list of SpotPrice
type SpotPriceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SpotPrice `json:"items"`
}

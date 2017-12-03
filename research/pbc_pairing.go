package research

import (
	"gobdc/pairing"
)

func getCompatParams() (params pairing.PairingParameters) {

	/*
	   *** this is now actually *also* cribbed from the bdc project ...
	   taken from PBC-provided param file 'a.param'
	   type a
	   q 8780710799663312522437781984754049815806883199414208211028653399266475630880222957078625179422662221423155858769582317459277713367317481324925129998224791
	   h 12016012264891146079388821366740534204802954401251311822919615131047207289359704531102844802183906537786776
	   r 730750818665451621361119245571504901405976559617
	   exp2 159
	   exp1 107
	   sign1 1
	   sign0 1
	*/

	params = pairing.PairingParameters{}
	params["type"] = "a"
	params["q"] = "8780710799663312522437781984754049815806883199414208211028653399266475630880222957078625179422662221423155858769582317459277713367317481324925129998224791"
	params["r"] = "730750818665451621361119245571504901405976559617"
	params["h"] = "12016012264891146079388821366740534204802954401251311822919615131047207289359704531102844802183906537786776"
	params["exp1"] = "107"
	params["exp2"] = "159"
	params["sign0"] = "1"
	params["sign1"] = "1"

	params["genNoCofac"] = "WLeuxaO0DxaW+oJ4vrLKgkq91prZNLGQUVoXH4gIx6AGIS7vrU7Fq3/5DfYTRHfpnOCIuo96hfRwTzUTf2+EUndlGtVaI05vjWxsIaCqKSPtq+xYpr0jaGVVwnXojhjbi0AeR/JvjiIaF9WFjSRzqEvR1WHp0LkJRrtBfNcA0k4="

	return
}

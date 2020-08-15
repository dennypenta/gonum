package countries

import "errors"

type Country uint8

const (  
   CountryRu = iota + 1 	  
   CountryEn 	
)

var countryTable = map[string]Country{ 
    "RU": CountryRu, 
    "EN": CountryEn, 
}

func (t Country) String() string {
	switch t {     
    case CountryRu:
        return "RU"	    
    case CountryEn:
        return "EN"	
    }
	return "unknown"
}

var ErrUnknownCountry = errors.New("unknown country")

func NewCountry(str string) (Country, error) {
	if v, ok := countryTable[str]; ok {
		return v, nil
	} else {
		return Country(0), ErrUnknownCountry
	}
}
package main

import "encoding/json"

func CreateDraftInvoice() Invoice {
	invoice := Invoice{
		FaturaUuidstring:             "",
		BelgeNumarasi:                "",
		FaturaTarihi:                 "",
		Saat:                         "",
		ParaBirimi:                   "",
		DovzTLkur:                    "",
		FaturaTipi:                   "",
		HangiTip:                     "",
		SiparisNumarasi:              "",
		SiparisTarihi:                "",
		IrsaliyeNumarasi:             "",
		IrsaliyeTarihi:               "",
		FisNo:                        "",
		FisTarihi:                    "",
		FisSaati:                     "",
		FisTipi:                      "",
		ZRaporNo:                     "",
		OkcSeriNo:                    "",
		VknTckn:                      "",
		AliciUnvan:                   "",
		AliciAdi:                     "",
		AliciSoyadi:                  "",
		Bulvarcaddesokak:             "",
		BinaAdi:                      "",
		BinaNo:                       "",
		KapiNo:                       "",
		KasabaKoy:                    "",
		MahalleSemtIlce:              "",
		Sehir:                        "",
		Ulke:                         "",
		PostaKodu:                    "",
		Tel:                          "",
		Fax:                          "",
		Eposta:                       "",
		Websitesi:                    "",
		VergiDairesi:                 "",
		KomisyonOrani:                "",
		NavlunOrani:                  "",
		HammaliyeOrani:               "",
		NakliyeOrani:                 "",
		KomisyonTutari:               "",
		NavlunTutari:                 "",
		HammaliyeTutari:              "",
		NakliyeTutari:                "",
		KomisyonKDVOrani:             "",
		NavlunKDVOrani:               "",
		HammaliyeKDVOrani:            "",
		NakliyeKDVOrani:              "",
		KomisyonKDVTutari:            "",
		NavlunKDVTutari:              "",
		HammaliyeKDVTutari:           "",
		NakliyeKDVTutari:             "",
		GelirVergisiOrani:            "",
		BagkurTevkifatiOrani:         "",
		GelirVergisiTevkifatiTutari:  "",
		BagkurTevkifatiOrani2:        "",
		GelirVergisiTevkifatiTutari2: "",
		BagkurTevkifatiTutari:        "",
		HalRusumuOrani:               "",
		TicaretBorsasiOrani:          "",
		MilliSavunmaFonuOrani:        "",
		DigerOrani:                   "",
		HalRusumuTutari:              "",
		TicaretBorsasiTutari:         "",
		MilliSavunmaFonuTutari:       "",
		DigerTutari:                  "",
		HalRusumuKDVOrani:            "",
		TicaretBorsasiKDVOrani:       "",
		MilliSavunmaFonuKDVOrani:     "",
		DigerKDVOrani:                "",
		HalRusumuKDVTutari:           "",
		TicaretBorsasiKDVTutari:      "",
		MilliSavunmaFonuKDVTutari:    "",
		DigerKDVTutari:               "",
		IadeTable:                    "",
		OzelMatrahTutari:             "",
		OzelMatrahOrani:              "",
		OzelMatrahVergiTutari:        "",
		VergiCesidi:                  "",
		MalHizmetTable:               nil,
		Tip:                          "",
		Matrah:                       "",
		MalhizmetToplamTutari:        "",
		ToplamIskonto:                "",
		Hesaplanankdv:                "",
		VergilerToplami:              "",
		VergilerDahilToplamTutar:     "",
		ToplamMasraflar:              "",
		OdenecekTutar:                "",
		Not:                          "",
	}
	return invoice
}

func CreateDraftInvoiceAsJson() string {
	invoice := CreateDraftInvoice()
	json, _ := json.Marshal(invoice)
	return string(json)
}

type InvoiceDetail struct {
	IskontoArttm      string `json:"iskontoArttm"`
	MalHizmet         string `json:"malHizmet"`
	Miktar            string `json:"miktar"`
	Birim             string `json:"birim"`
	BirimFiyat        string `json:"birimFiyat"`
	Fiyat             string `json:"fiyat"`
	IskontoOrani      string `json:"iskontoOrani"`
	IskontoTutari     string `json:"iskontoTutari"`
	IskontoNedeni     string `json:"iskontoNedeni"`
	MalHizmetTutari   string `json:"malHizmetTutari"`
	KdvOrani          string `json:"kdvOrani"`
	VergiOrani        string `json:"vergiOrani"`
	KdvTutari         string `json:"kdvTutari"`
	VergininKdvTutari string `json:"vergininKdvTutari"`
}

type Invoice struct {
	FaturaUuidstring             string          `json:"faturaUuidstring"`
	BelgeNumarasi                string          `json:"belgeNumarasi"`
	FaturaTarihi                 string          `json:"faturaTarihi"`
	Saat                         string          `json:"saat"`
	ParaBirimi                   string          `json:"paraBirimi"`
	DovzTLkur                    string          `json:"dovzTLkur"`
	FaturaTipi                   string          `json:"faturaTipi"`
	HangiTip                     string          `json:"hangiTip"`
	SiparisNumarasi              string          `json:"siparisNumarasi"`
	SiparisTarihi                string          `json:"siparisTarihi"`
	IrsaliyeNumarasi             string          `json:"irsaliyeNumarasi"`
	IrsaliyeTarihi               string          `json:"irsaliyeTarihi"`
	FisNo                        string          `json:"fisNo"`
	FisTarihi                    string          `json:"fisTarihi"`
	FisSaati                     string          `json:"fisSaati"`
	FisTipi                      string          `json:"fisTipi"`
	ZRaporNo                     string          `json:"zRaporNo"`
	OkcSeriNo                    string          `json:"okcSeriNo"`
	VknTckn                      string          `json:"vknTckn"`
	AliciUnvan                   string          `json:"aliciUnvan"`
	AliciAdi                     string          `json:"aliciAdi"`
	AliciSoyadi                  string          `json:"aliciSoyadi"`
	Bulvarcaddesokak             string          `json:"bulvarcaddesokak"`
	BinaAdi                      string          `json:"binaAdi"`
	BinaNo                       string          `json:"binaNo"`
	KapiNo                       string          `json:"kapiNo"`
	KasabaKoy                    string          `json:"kasabaKoy"`
	MahalleSemtIlce              string          `json:"mahalleSemtIlce"`
	Sehir                        string          `json:"sehir"`
	Ulke                         string          `json:"ulke"`
	PostaKodu                    string          `json:"postaKodu"`
	Tel                          string          `json:"tel"`
	Fax                          string          `json:"fax"`
	Eposta                       string          `json:"eposta"`
	Websitesi                    string          `json:"websitesi"`
	VergiDairesi                 string          `json:"vergiDairesi"`
	KomisyonOrani                string          `json:"komisyonOrani"`
	NavlunOrani                  string          `json:"navlunOrani"`
	HammaliyeOrani               string          `json:"hammaliyeOrani"`
	NakliyeOrani                 string          `json:"nakliyeOrani"`
	KomisyonTutari               string          `json:"komisyonTutari"`
	NavlunTutari                 string          `json:"navlunTutari"`
	HammaliyeTutari              string          `json:"hammaliyeTutari"`
	NakliyeTutari                string          `json:"nakliyeTutari"`
	KomisyonKDVOrani             string          `json:"komisyonKDVOrani"`
	NavlunKDVOrani               string          `json:"navlunKDVOrani"`
	HammaliyeKDVOrani            string          `json:"hammaliyeKDVOrani"`
	NakliyeKDVOrani              string          `json:"nakliyeKDVOrani"`
	KomisyonKDVTutari            string          `json:"komisyonKDVTutari"`
	NavlunKDVTutari              string          `json:"navlunKDVTutari"`
	HammaliyeKDVTutari           string          `json:"hammaliyeKDVTutari"`
	NakliyeKDVTutari             string          `json:"nakliyeKDVTutari"`
	GelirVergisiOrani            string          `json:"gelirVergisiOrani"`
	BagkurTevkifatiOrani         string          `json:"bagkurTevkifatiOrani"`
	GelirVergisiTevkifatiTutari  string          `json:"gelirVergisiTevkifatiTutari"`
	BagkurTevkifatiOrani2        string          `json:"bagkurTevkifatiOrani"`
	GelirVergisiTevkifatiTutari2 string          `json:"gelirVergisiTevkifatiTutari"`
	BagkurTevkifatiTutari        string          `json:"bagkurTevkifatiTutari"`
	HalRusumuOrani               string          `json:"halRusumuOrani"`
	TicaretBorsasiOrani          string          `json:"ticaretBorsasiOrani"`
	MilliSavunmaFonuOrani        string          `json:"milliSavunmaFonuOrani"`
	DigerOrani                   string          `json:"digerOrani"`
	HalRusumuTutari              string          `json:"halRusumuTutari"`
	TicaretBorsasiTutari         string          `json:"ticaretBorsasiTutari"`
	MilliSavunmaFonuTutari       string          `json:"milliSavunmaFonuTutari"`
	DigerTutari                  string          `json:"digerTutari"`
	HalRusumuKDVOrani            string          `json:"halRusumuKDVOrani"`
	TicaretBorsasiKDVOrani       string          `json:"ticaretBorsasiKDVOrani"`
	MilliSavunmaFonuKDVOrani     string          `json:"milliSavunmaFonuKDVOrani"`
	DigerKDVOrani                string          `json:"digerKDVOrani"`
	HalRusumuKDVTutari           string          `json:"halRusumuKDVTutari"`
	TicaretBorsasiKDVTutari      string          `json:"ticaretBorsasiKDVTutari"`
	MilliSavunmaFonuKDVTutari    string          `json:"milliSavunmaFonuKDVTutari"`
	DigerKDVTutari               string          `json:"digerKDVTutari"`
	IadeTable                    string          `json:"iadeTable"`
	OzelMatrahTutari             string          `json:"ozelMatrahTutari"`
	OzelMatrahOrani              string          `json:"ozelMatrahOrani"`
	OzelMatrahVergiTutari        string          `json:"ozelMatrahVergiTutari"`
	VergiCesidi                  string          `json:"vergiCesidi"`
	MalHizmetTable               []InvoiceDetail `json:"malHizmetTable"`
	Tip                          string          `json:"tip"`
	Matrah                       string          `json:"matrah"`
	MalhizmetToplamTutari        string          `json:"malhizmetToplamTutari"`
	ToplamIskonto                string          `json:"toplamIskonto"`
	Hesaplanankdv                string          `json:"hesaplanankdv"`
	VergilerToplami              string          `json:"vergilerToplami"`
	VergilerDahilToplamTutar     string          `json:"vergilerDahilToplamTutar"`
	ToplamMasraflar              string          `json:"toplamMasraflar"`
	OdenecekTutar                string          `json:"odenecekTutar"`
	Not                          string          `json:"not"`
}

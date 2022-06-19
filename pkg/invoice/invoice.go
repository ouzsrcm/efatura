package invoice

import (
	"encoding/json"
	"fmt"
	"invoicer/config"
	"invoicer/pkg/common"
	"time"
)

var TOKEN string

func Init() {
	fmt.Println("tokenization...")
	TOKEN = GetToken()
	fmt.Print(TOKEN)
}

func GetAllInvoicesByDateRange(token string, startdate time.Time, enddate time.Time) {
	fmt.Println("get all invoices by date range...")

}

func CreateDraftInvoice() *Invoice {
	invoice := &Invoice{
		FaturaUuid:                   common.GetUUID(),
		BelgeNumarasi:                "",
		FaturaTarihi:                 "",
		Saat:                         "",
		ParaBirimi:                   "TRY",
		DovzTLkur:                    "0",
		FaturaTipi:                   "5000/30000",
		HangiTip:                     "Buyuk",
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
		VknTckn:                      "11111111111",
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
		KomisyonOrani:                "0",
		NavlunOrani:                  "0",
		HammaliyeOrani:               "0",
		NakliyeOrani:                 "0",
		KomisyonTutari:               "0",
		NavlunTutari:                 "0",
		HammaliyeTutari:              "0",
		NakliyeTutari:                "0",
		KomisyonKDVOrani:             "0",
		NavlunKDVOrani:               "0",
		HammaliyeKDVOrani:            "0",
		NakliyeKDVOrani:              "0",
		KomisyonKDVTutari:            "0",
		NavlunKDVTutari:              "0",
		HammaliyeKDVTutari:           "0",
		NakliyeKDVTutari:             "0",
		GelirVergisiOrani:            "0",
		BagkurTevkifatiOrani:         "0",
		GelirVergisiTevkifatiTutari:  "0",
		BagkurTevkifatiOrani2:        "0",
		GelirVergisiTevkifatiTutari2: "0",
		BagkurTevkifatiTutari:        "0",
		HalRusumuOrani:               "0",
		TicaretBorsasiOrani:          "0",
		MilliSavunmaFonuOrani:        "0",
		DigerOrani:                   "0",
		HalRusumuTutari:              "0",
		TicaretBorsasiTutari:         "0",
		MilliSavunmaFonuTutari:       "0",
		DigerTutari:                  "0",
		HalRusumuKDVOrani:            "0",
		TicaretBorsasiKDVOrani:       "0",
		MilliSavunmaFonuKDVOrani:     "0",
		DigerKDVOrani:                "0",
		HalRusumuKDVTutari:           "0",
		TicaretBorsasiKDVTutari:      "0",
		MilliSavunmaFonuKDVTutari:    "0",
		DigerKDVTutari:               "0",
		IadeTable:                    "",
		OzelMatrahTutari:             "0",
		OzelMatrahOrani:              "0",
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
	invoice.MalHizmetTable = append(invoice.MalHizmetTable, InvoiceDetail{
		IskontoArttm:      "İskonto",
		MalHizmet:         "",
		Miktar:            "",
		Birim:             "",
		BirimFiyat:        "",
		Fiyat:             "",
		IskontoOrani:      "",
		IskontoTutari:     "",
		IskontoNedeni:     "",
		MalHizmetTutari:   "",
		KdvOrani:          "",
		VergiOrani:        "",
		KdvTutari:         "",
		VergininKdvTutari: "",
	})
	invoice_byte, _ := json.Marshal(invoice)
	invoice_json := string(invoice_byte)
	command := config.GetCommand("createDraftInvoice")
	RunCommand(TOKEN, command.Name, command.Value, invoice_json)
	return invoice
}

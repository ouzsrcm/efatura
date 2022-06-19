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

	temp, draft := CreateDraftInvoice()
	fmt.Println("temp: " + temp.FaturaUuid)
	fmt.Println("draft: " + draft.FaturaUuid)

	fmt.Println("get all invoices...")
	invoices_json := GetAllInvoicesByDateRange(time.Now(), time.Now().Add(time.Hour*24*30))
	fmt.Println(invoices_json)
}

func GetAllInvoicesByDateRange(startdate time.Time, enddate time.Time) string {
	fmt.Println("get all invoices by date range...")
	command := config.GetCommand("getAllInvoicesByDateRange")
	data := map[string]interface{}{
		"baslangic": startdate,
		"bitis":     enddate,
		"hangiTip":  "5000/30000",
	}
	data_json, _ := json.Marshal(data)
	invoices_json := RunCommand(TOKEN, command.Name, command.Value, string(data_json))
	return invoices_json
}

func CreateDraftInvoice() (*Invoice, *Invoice) {
	tempinvoice := &Invoice{
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
	tempinvoice.MalHizmetTable = append(tempinvoice.MalHizmetTable, InvoiceDetail{
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

	invoice_byte, _ := json.Marshal(tempinvoice)
	invoice_json := string(invoice_byte)

	command := config.GetCommand("createDraftInvoice")

	draft_text := RunCommand(TOKEN, command.Name, command.Value, invoice_json)
	draft := Invoice{}
	json.Unmarshal([]byte(draft_text), draft)

	return tempinvoice, &draft
}

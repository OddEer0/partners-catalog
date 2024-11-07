package scripts

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type Upload struct {
	Json string `json:"json"`
}

func SetPartnerCatalog() error {
	reqBody := Upload{Json: jsonStr}
	reqBodyByte, err := json.Marshal(reqBody)
	if err != nil {
		return err
	}
	body := bytes.NewBuffer(reqBodyByte)

	contentType := "application/json"
	resp, err := http.Post("http://localhost:10100/api/smarthome/v1/partners-catalog", contentType, body)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("SetPartnerCatalog returned status code %d", resp.StatusCode)
	}

	f, err := os.OpenFile("set_partner_catalog.json", os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		return err
	}

	_, err = f.Write(reqBodyByte)
	if err != nil {
		return err
	}

	return nil
}

const jsonStr = `{
    "categories": [
        {
            "name": "Вентиляторы",
            "id": "hvac-fan"
        },
        {
            "name": "Ворота",
            "id": "gate"
        },
        {
            "name": "Выключатели и реле",
            "id": "relay"
        },
        {
            "name": "Датчики газа",
            "id": "sensor-gas"
        },
        {
            "name": "Датчики движения",
            "id": "sensor-pir"
        },
        {
            "name": "Датчики дыма",
            "id": "sensor-smoke"
        },
        {
            "name": "Датчики открытия",
            "id": "sensor-door"
        },
        {
            "name": "Датчики протечки",
            "id": "sensor-water-leak"
        },
        {
            "name": "Датчики температуры и влажности",
            "id": "sensor-temp"
        },
        {
            "name": "Жалюзи и рулонные шторы",
            "id": "window-blind"
        },
        {
            "name": "Кондиционеры",
            "id": "hvac-ac"
        },
        {
            "name": "Контроллеры и управляющие устройства",
            "id": "controller"
        },
        {
            "name": "Котлы, контроллеры отопления",
            "id": "hvac-boiler"
        },
        {
            "name": "Краны с электроприводами",
            "id": "valve"
        },
        {
            "name": "Лампы",
            "id": "light"
        },
        {
            "name": "Обогреватели",
            "id": "hvac-heater"
        },
        {
            "name": "Пылесосы",
            "id": "vacuum-cleaner"
        },
        {
            "name": "Раздвижные шторы",
            "id": "curtain"
        },
        {
            "name": "Розетки",
            "id": "socket"
        },
        {
            "name": "Светодиодные ленты",
            "id": "led-strip"
        },
        {
            "name": "Сценарные кнопки",
            "id": "scenario-button"
        },
        {
            "name": "Термоголовки и терморегуляторы для радиаторов",
            "id": "hvac-radiator"
        },
        {
            "name": "Тёплые полы",
            "id": "hvac-underfloor-heating"
        },
        {
            "name": "Увлажнители",
            "id": "hvac-humidifier"
        },
        {
            "name": "Хабы (шлюзы)",
            "id": "hub"
        },
        {
            "name": "Чайники",
            "id": "kettle"
        },
        {
            "name": "Другое",
            "id": "other"
        }
    ],
    "protocols": [
        {
            "name": "Bluetooth",
            "id": "bluetooth"
        },
        {
            "name": "MQTT",
            "id": "mqtt"
        },
        {
            "name": "Wi-Fi",
            "id": "wifi"
        },
        {
            "name": "Zigbee",
            "id": "zigbee"
        },
        {
            "name": "Z-Wave",
            "id": "zwave"
        },
        {
            "name": "Собственный протокол",
            "id": "own"
        },
        {
            "name": "Другой",
            "id": "other"
        }
    ],
    "manufacturers": [
        {
            "name": "Sber",
            "id": "sber",
            "brand_name": "Умный дом Sber"
        },
        {
            "name": "Aero",
            "id": "aero",
            "brand_name": "Daichi Comfort"
        },
        {
            "name": "Aeronik",
            "id": "aeronik",
            "brand_name": "Daichi Comfort"
        },
        {
            "name": "Airwave",
            "id": "airwave",
            "brand_name": "Daichi Comfort"
        },
        {
            "name": "Airwell",
            "id": "airwell",
            "brand_name": "Daichi Comfort"
        },
        {
            "name": "Aqara",
            "id": "aqara",
            "brand_name": "Aqara"
        },
        {
            "name": "Aurum",
            "id": "aurum",
            "brand_name": "Daichi Comfort"
        },
        {
            "name": "Axioma",
            "id": "axioma",
            "brand_name": "Daichi Comfort"
        },
        {
            "name": "Ballu",
            "id": "ballu",
            "brand_name": "Hommyn"
        },
        {
            "name": "Bosch",
            "id": "bosch",
            "brand_name": "Daichi Comfort"
        },
        {
            "name": "Candy",
            "id": "candy",
            "brand_name": "evo"
        },
        {
            "name": "Casarte",
            "id": "casarte",
            "brand_name": "evo"
        },
        {
            "name": "Centek",
            "id": "centek",
            "brand_name": "Daichi Comfort"
        },
        {
            "name": "Citilux",
            "id": "сitilux",
            "brand_name": "Citilux SMART"
        },
        {
            "name": "Comfee",
            "id": "comfee",
            "brand_name": "Daichi Comfort"
        },
        {
            "name": "CoolUp",
            "id": "coolup",
            "brand_name": "Daichi Comfort"
        },
        {
            "name": "Сooper&Hunter",
            "id": "cooper-hunter",
            "brand_name": "Daichi Comfort"
        },
        {
            "name": "Daichi Comfort",
            "id": "daichi",
            "brand_name": "Daichi Comfort"
        },
        {
            "name": "Daikin",
            "id": "daikin",
            "brand_name": "Daichi Comfort"
        },
        {
            "name": "Dantex",
            "id": "dantex",
            "brand_name": "Daichi Comfort"
        },
        {
            "name": "Digma",
            "id": "digma",
            "brand_name": "DIGMA"
        },
        {
            "name": "Ecoletta",
            "id": "ecoletta",
            "brand_name": "Daichi Comfort"
        },
        {
            "name": "EKF Connect",
            "id": "ekf",
            "brand_name": "EKF Connect Home"
        },
        {
            "name": "ELARI Smart Home",
            "id": "elari",
            "brand_name": "ELARI SmartHome"
        },
        {
            "name": "Electrolux",
            "id": "electrolux",
            "brand_name": "Hommyn"
        },
        {
            "name": "Elektrostandard",
            "id": "elektrostandard",
            "brand_name": "Minimir Home"
        },
        {
            "name": "Eltex Home",
            "id": "eltexhome",
            "brand_name": "Eltex Home"
        },
        {
            "name": "FUNAI",
            "id": "funai",
            "brand_name": "Daichi Comfort"
        },
        {
            "name": "GEOZON",
            "id": "geozon",
            "brand_name": "GEOZON"
        },
        {
            "name": "Giulia Novars",
            "id": "giulianovars",
            "brand_name": "Giulia Novars Smart"
        },
        {
            "name": "Gosund",
            "id": "gosund",
            "brand_name": "Gosund"
        },
        {
            "name": "Haier",
            "id": "haier",
            "brand_name": "evo"
        },
        {
            "name": "Halsen",
            "id": "halsen",
            "brand_name": "Daichi Comfort"
        },
        {
            "name": "HDL Automation",
            "id": "hdl",
            "brand_name": "HDL Automation"
        },
        {
            "name": "HEC",
            "id": "hec",
            "brand_name": "evo"
        },
        {
            "name": "Hi",
            "id": "hi",
            "brand_name": "Daichi Comfort"
        },
        {
            "name": "HIPER IoT",
            "id": "hiper",
            "brand_name": "HIPER IoT"
        },
        {
            "name": "HiTE PRO",
            "id": "hitepro",
            "brand_name": "HiTE PRO - Нет повода для провода"
        },
        {
            "name": "Hunberg",
            "id": "hunberg",
            "brand_name": "Daichi Comfort"
        },
        {
            "name": "iFEEL",
            "id": "ifeel",
            "brand_name": "iFEEL Safe+Smart"
        },
        {
            "name": "IGC",
            "id": "igc",
            "brand_name": "Daichi Comfort"
        },
        {
            "name": "inSmart",
            "id": "insmart",
            "brand_name": "inSmart"
        },
        {
            "name": "IRBIS",
            "id": "irbis",
            "brand_name": "IRBIS SMART HOME"
        },
        {
            "name": "iRidi",
            "id": "iridi",
            "brand_name": "iRidi"
        },
        {
            "name": "iTEQ",
            "id": "iteq",
            "brand_name": "iTEQ"
        },
        {
            "name": "Kentatsu",
            "id": "kentatsu",
            "brand_name": "Daichi Comfort"
        },
        {
            "name": "Komanchi",
            "id": "komanchi",
            "brand_name": "Daichi Comfort"
        },
        {
            "name": "LESSAR",
            "id": "lessar",
            "brand_name": "Daichi Comfort"
        },
        {
            "name": "MDV",
            "id": "mdv",
            "brand_name": "Daichi Comfort"
        },
        {
            "name": "Midea",
            "id": "midea",
            "brand_name": "Daichi Comfort"
        },
        {
            "name": "Mitsubishi Electric",
            "id": "mitsubishi-electric",
            "brand_name": "Daichi Comfort"
        },
        {
            "name": "MOES",
            "id": "moes",
            "brand_name": "MOES"
        },
        {
            "name": "MOiO",
            "id": "moio",
            "brand_name": "MOiO"
        },
        {
            "name": "One Air",
            "id": "one-air",
            "brand_name": "Daichi Comfort"
        },
        {
            "name": "Philips Hue",
            "id": "philips",
            "brand_name": "Philips"
        },
        {
            "name": "Polaris IQ Home",
            "id": "polaris",
            "brand_name": "Polaris"
        },
        {
            "name": "Primera",
            "id": "primera",
            "brand_name": "Daichi Comfort"
        },
        {
            "name": "RED Solution",
            "id": "red",
            "brand_name": "Sky умный дом и сервисы"
        },
        {
            "name": "Roximo IoT",
            "id": "roximo",
            "brand_name": "Roximo IoT"
        },
        {
            "name": "ROYAL Premium",
            "id": "royal-premium",
            "brand_name": "Daichi Comfort"
        },
        {
            "name": "Scarlett",
            "id": "scarlett",
            "brand_name": "Home Intellect"
        },
        {
            "name": "Shuft",
            "id": "shuft",
            "brand_name": "Daichi Comfort"
        },
        {
            "name": "Sibling",
            "id": "sibling",
            "brand_name": "Sibling"
        },
        {
            "name": "SmartHome Tricolor",
            "id": "tricolor",
            "brand_name": "SmartHome Tricolor"
        },
        {
            "name": "Sonoff",
            "id": "sonoff",
            "brand_name": "eWeLink smart home"
        },
        {
            "name": "Tesla",
            "id": "tesla",
            "brand_name": "Daichi Comfort"
        },
        {
            "name": "Timberk",
            "id": "timberk",
            "brand_name": "Home Intellect"
        },
        {
            "name": "Tuya",
            "id": "tuya",
            "brand_name": "Tuya Smart"
        },
        {
            "name": "Ujin",
            "id": "ujin",
            "brand_name": "Ujin"
        },
        {
            "name": "Vixion Home",
            "id": "vixion",
            "brand_name": "Vixion Home"
        },
        {
            "name": "Welrok",
            "id": "welrok",
            "brand_name": "Welrok"
        },
        {
            "name": "Wise Home",
            "id": "wisehome",
            "brand_name": "Wise Home"
        },
        {
            "name": "Yeelight",
            "id": "yeelight",
            "brand_name": "Yeelight"
        },
        {
            "name": "Yeelight Pro",
            "id": "yeelightpro",
            "brand_name": "Yeelight Pro"
        },
        {
            "name": "Zanussi",
            "id": "zanussi",
            "brand_name": "Daichi Comfort"
        },
        {
            "name": "ZONT",
            "id": "zont",
            "brand_name": "ZONT"
        },
        {
            "name": "Мажордом",
            "id": "majordom",
            "brand_name": "Мажордом"
        },
        {
            "name": "МТС",
            "id": "mts",
            "brand_name": "МТС"
        },
        {
            "name": "ЦЕНТРСВЕТ",
            "id": "centersvet",
            "brand_name": "ЦЕНТРСВЕТ"
        },
        {
            "name": "ЭРА Smart",
            "id": "era",
            "brand_name": "ЭРА Smart"
        },
        {
            "name": "Яндекс",
            "id": "yandex",
            "brand_name": "Яндекс"
        }
    ],
    "direct": [
        {
            "name": "Работает с Zigbee-колонкой или хабом Sber",
            "id": "hub"
        },
        {
            "name": "Работает с контроллером умного дома Sber",
            "id": "controller"
        }
    ],
	"partners": [
		{
			"brand_name": "Philips",
			"aliases": ["fillips"]
		}
	],
    "cloud": [  
        {  
          "name": "Подключается к облаку производителя через его хаб", 
          "id": "с2с-hub"  
        },  
        {  
          "name": "Подключается к облаку производителя напрямую",  
          "id": "с2с-app"  
        }  
    ],
    "devices": [
        {
            "manufacturer": "aero",
            "model": "",
            "name": "Кондиционеры Aero. Для управления по Wi-Fi должны быть оснащены модулем Daichi DW21-B/DW22-B",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/aero-hvac-ac.svg",
            "link": "https://daichicloud.ru/split-lineup/search.php?search=AERO",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "aeronik",
            "model": "",
            "name": "Кондиционеры Aeronik. Для управления по Wi-Fi должны быть оснащены модулем Daichi DW21-B/DW22-B",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/aeronik-hvac-ac.svg",
            "link": "https://daichicloud.ru/split-lineup/search.php?search=Aeronik",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "airwave",
            "model": "",
            "name": "Кондиционеры Airwave. Для управления по Wi-Fi должны быть оснащены модулем Daichi DW21-B/DW22-B",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/airwave-hvac-ac.svg",
            "link": "https://daichicloud.ru/split-lineup/search.php?search=AIRWAVE",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "airwell",
            "model": "",
            "name": "Кондиционеры Airwell. Для управления по Wi-Fi должны быть оснащены модулем Daichi DW21-B/DW22-B",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/airwell-hvac-ac.svg",
            "link": "https://daichicloud.ru/split-lineup/search.php?search=AIRWELL",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "aqara",
            "model": "LLKZMK11LM",
            "name": "Aqara Wireless Relay. Подключается через хаб Aqara",
            "categories": ["relay"],
            "protocols": ["zigbee"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/a6419c50_aqara-relay-1.svg",
            "link": "https://aqara.ru/product/беспроводное-реле-aqara-wireless-relay/",
            "direct": [],
            "cloud": ["с2с-hub"]
        },
        {
            "manufacturer": "aqara",
            "model": "QBKG03LM",
            "name": "Aqara Wall Switch (No Neutral, Double Rocker). Подключается через хаб Aqara",
            "categories": ["relay"],
            "protocols": ["zigbee"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/85dce993_aqara-relay-2.png",
            "link": "https://aqara.ru/product/aqara-wall-switch-double-rocker/",
            "direct": [],
            "cloud": ["с2с-hub"]
        },
        {
            "manufacturer": "aqara",
            "model": "QBKG04LM",
            "name": "Aqara Wall Switch (No Neutral, Single Rocker). Подключается через хаб Aqara",
            "categories": ["relay"],
            "protocols": ["zigbee"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/fda83fcf_aqara-relay-3.png",
            "link": "https://aqara.ru/product/aqara-wall-switch-single-rocker/",
            "direct": [],
            "cloud": ["с2с-hub"]
        },
        {
            "manufacturer": "aqara",
            "model": "SSM-U01",
            "name": "Aqara Single Switch Module T1 (With Neutral). Подключается через хаб Aqara, Zigbee-колонку Sber или хаб Sber",
            "categories": ["relay"],
            "protocols": ["zigbee"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/34a4724b_aqara-relay-4.png",
            "link": "https://aqara.ru/product/модуль-реле-одноканальный-с-нейтраль/",
            "direct": ["hub"],
            "cloud": ["с2с-hub"]
        },
        {
            "manufacturer": "aqara",
            "model": "SSM-U02",
            "name": "Aqara Single Switch Module T1 (No Neutral). Подключается через хаб Aqara, Zigbee-колонку Sber или хаб Sber",
            "categories": ["relay"],
            "protocols": ["zigbee"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/4f4f192d_aqara-relay-5.png",
            "link": "https://aqara.ru/product/реле-одноканальное-без-нейтрали-aqara-single-switch-modu/",
            "direct": ["hub"],
            "cloud": ["с2с-hub"]
        },
        {
            "manufacturer": "aqara",
            "model": "WS-EUK01",
            "name": "Aqara Smart Wall Switch H1 EU (No Neutral, Single Rocker). Подключается через хаб Aqara, Zigbee-колонку Sber или хаб Sber",
            "categories": ["relay"],
            "protocols": ["zigbee"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/cc478a91_aqara-relay-6.png",
            "link": "https://aqara.ru/product/aqara-smart-wall-switch-h1-no-neutral-single-rocker/",
            "direct": ["hub"],
            "cloud": ["с2с-hub"]
        },
        {
            "manufacturer": "aqara",
            "model": "WS-EUK02",
            "name": "Aqara Smart Wall Switch H1 EU (No Neutral, Double Rocker). Подключается через хаб Aqara, Zigbee-колонку Sber или хаб Sber",
            "categories": ["relay"],
            "protocols": ["zigbee"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/aae45cd2_aqara-relay-7.png",
            "link": "https://aqara.ru/product/aqara-smart-wall-switch-h1-no-neutral-double-rocker/",
            "direct": ["hub"],
            "cloud": ["с2с-hub"]
        },
        {
            "manufacturer": "aqara",
            "model": "WS-EUK03",
            "name": "Aqara Smart Wall Switch H1 EU (With Neutral, Single Rocker). Подключается через хаб Aqara, Zigbee-колонку Sber или хаб Sber",
            "categories": ["relay"],
            "protocols": ["zigbee"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/cc478a91_aqara-relay-6.png",
            "link": "https://aqara.ru/product/aqara-smart-wall-switch-h1-with-neutral-single-rocker/",
            "direct": ["hub"],
            "cloud": ["с2с-hub"]
        },
        {
            "manufacturer": "aqara",
            "model": "WS-EUK04",
            "name": "Aqara Smart Wall Switch H1 EU (With Neutral, Double Rocker). Подключается через хаб Aqara, Zigbee-колонку Sber или хаб Sber",
            "categories": ["relay"],
            "protocols": ["zigbee"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/aae45cd2_aqara-relay-7.png",
            "link": "https://aqara.ru/product/aqara-smart-wall-switch-h1-with-neutral-double-rocker/",
            "direct": ["hub"],
            "cloud": ["с2с-hub"]
        },
        {
            "manufacturer": "aqara",
            "model": "WRS-R02",
            "name": "Aqara Wireless Remote Switch H1 (Double Rocker). Подключается через хаб Aqara, Zigbee-колонку Sber или хаб Sber",
            "categories": ["relay"],
            "protocols": ["zigbee"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/bbb96de6_aqara-relay-10-2.svg",
            "link": "https://aqara.ru/product/aqara-wireless-remote-switch-h1/",
            "direct": ["hub"],
            "cloud": ["с2с-hub"]
        },
        {
            "manufacturer": "aqara",
            "model": "WXKG03LM",
            "name": "Aqara Wireless Remote Switch (Single Rocker). Подключается через хаб Aqara",
            "categories": ["relay"],
            "protocols": ["zigbee"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/aqara-wireless-remote-switch.svg",
            "link": "https://aqara.ru/product/aqara-wireless-remote-switch/",
            "direct": [],
            "cloud": ["с2с-hub"]
        },
        {
            "manufacturer": "aqara",
            "model": "WXKG02LM",
            "name": "Aqara Wireless Remote Switch (Double Rocker). Подключается через хаб Aqara",
            "categories": ["relay"],
            "protocols": ["zigbee"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/aqara-wireless-remote-switch-double-rocker.svg",
            "link": "https://aqara.ru/product/aqara-wireless-remote-switch-double-rocker/",
            "direct": [],
            "cloud": ["с2с-hub"]
        },
        {
            "manufacturer": "aqara",
            "model": "DCM-K01",
            "name": "Aqara Dual Relay Module T2. Подключается через хаб Aqara",
            "categories": ["relay"],
            "protocols": ["zigbee"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/aqara-dual-relay-module-t2.svg",
            "link": "https://aqara.ru/product/реле-двухканальное-aqara-dual-relay-module-t2/",
            "direct": [],
            "cloud": ["с2с-hub"]
        },
        {
            "manufacturer": "aqara",
            "model": "JT-BZ-03AQ/A",
            "name": "Aqara Smart Natural Gas Detector. Подключается через хаб Aqara, Zigbee-колонку Sber или хаб Sber",
            "categories": ["sensor-gas"],
            "protocols": ["zigbee"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/aqara-smart-natural-gas-detector.svg",
            "link": "https://aqara.ru/product/умный-датчик-газа-aqara-smart-natural-gas-detector/",
            "direct": ["hub"],
            "cloud": ["с2с-hub"]
        },
        {
            "manufacturer": "aqara",
            "model": "RTCGQ11LM",
            "name": "Aqara Motion Sensor. Подключается через хаб Aqara",
            "categories": ["sensor-pir"],
            "protocols": ["zigbee"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/2d816e79_aqara-sensor-pir-1.jpeg",
            "link": "https://aqara.ru/product/aqara-motion-sensor/",
            "direct": [],
            "cloud": ["с2с-hub"]
        },
        {
            "manufacturer": "aqara",
            "model": "MS-S02",
            "name": "Aqara Motion Sensor P1. Подключается через хаб Aqara, Zigbee-колонку Sber или хаб Sber",
            "categories": ["sensor-pir"],
            "protocols": ["zigbee"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/aqara-motion-sensor-p1.svg",
            "link": "https://aqara.ru/product/датчик-движения-aqara-motion-sensor-p1/",
            "direct": ["hub"],
            "cloud": ["с2с-hub"]
        },
        {
            "manufacturer": "aqara",
            "model": "JY-GZ-03AQ",
            "name": "Aqara Smart Smoke Detector. Подключается через хаб Aqara, Zigbee-колонку Sber или хаб Sber",
            "categories": ["sensor-smoke"],
            "protocols": ["zigbee"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/aqara-smart-smoke-detector.svg",
            "link": "https://aqara.ru/product/умный-датчик-дыма-aqara-smart-smoke-detector/",
            "direct": ["hub"],
            "cloud": ["с2с-hub"]
        },
        {
            "manufacturer": "aqara",
            "model": "MCCGQ11LM",
            "name": "Aqara Door and Window Sensor. Подключается через хаб Aqara, Zigbee-колонку Sber или хаб Sber",
            "categories": ["sensor-door"],
            "protocols": ["zigbee"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/b052308a_aqara-sensor-door-1.jpeg",
            "link": "https://aqara.ru/product/aqara-door-and-window-sensor/",
            "direct": ["hub"],
            "cloud": ["с2с-hub"]
        },
        {
            "manufacturer": "aqara",
            "model": "DW-S03D",
            "name": "Aqara Door and Window Sensor T1. Подключается через хаб Aqara, Zigbee-колонку Sber или хаб Sber",
            "categories": ["sensor-door"],
            "protocols": ["zigbee"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/aqara-door-and-window-sensor-t1.svg",
            "link": "https://aqara.ru/product/датчик-открытия-дверей-и-окон-aqara-door-and-window-sensor-t1/",
            "direct": ["hub"],
            "cloud": ["с2с-hub"]
        },
        {
            "manufacturer": "aqara",
            "model": "SJCGQ11LM",
            "name": "Aqara Water Leak Sensor. Подключается через хаб Aqara",
            "categories": ["sensor-water-leak"],
            "protocols": ["zigbee"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/b4b1e9b8_aqara-sensor-water-leak-1.jpeg",
            "link": "https://aqara.ru/product/aqara-water-leak-sensor/",
            "direct": [],
            "cloud": ["с2с-hub"]
        },
        {
            "manufacturer": "aqara",
            "model": "WL-S02D",
            "name": "Aqara Water Leak Sensor T1. Подключается через хаб Aqara, Zigbee-колонку Sber или хаб Sber",
            "categories": ["sensor-water-leak"],
            "protocols": ["zigbee"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/b4b1e9b8_aqara-sensor-water-leak-1.jpeg",
            "link": "https://aqara.ru/product/датчик-протечки-aqara-water-leak-sensor-t1/",
            "direct": ["hub"],
            "cloud": ["с2с-hub"]
        },
        {
            "manufacturer": "aqara",
            "model": "WSDCGQ11LM",
            "name": "Aqara Temperature and Humidity Sensor. Подключается через хаб Aqara, Zigbee-колонку Sber или хаб Sber",
            "categories": ["sensor-temp"],
            "protocols": ["zigbee"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/59a7cf67_aqara-sensor-temp-1.jpeg",
            "link": "https://aqara.ru/product/aqara-temperature-and-humidity-sensor/",
            "direct": ["hub"],
            "cloud": ["с2с-hub"]
        },
        {
            "manufacturer": "aqara",
            "model": "TH-S02D",
            "name": "Aqara Temperature and Humidity Sensor T1. Подключается через хаб Aqara, Zigbee-колонку Sber или хаб Sber",
            "categories": ["sensor-temp"],
            "protocols": ["zigbee"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/aqara-temperature-and-humidity-s.svg",
            "link": "https://aqara.ru/product/датчик-температуры-и-влажности-aqara-temperature-and-humidity-s/",
            "direct": ["hub"],
            "cloud": ["с2с-hub"]
        },
        {
            "manufacturer": "aqara",
            "model": "SRSC-M01",
            "name": "Aqara Roller Shade. Подключается через хаб Aqara",
            "categories": ["window-blind"],
            "protocols": ["zigbee"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/524f8324_aqara-window-blind-1.svg",
            "link": "https://aqara.ru/product/мотор-для-рулонных-штор-aqara-roller-shade/",
            "direct": [],
            "cloud": ["с2с-hub"]
        },
        {
            "manufacturer": "aqara",
            "model": "RSD-M01",
            "name": "Aqara Roller Shade Driver E1. Подключается через хаб Aqara",
            "categories": ["window-blind"],
            "protocols": ["zigbee"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/b6498f62_aqara-window-blind-2.svg",
            "link": "https://aqara.ru/product/мотор-для-рулонных-штор-и-жалюзи-aqara-roller-shade-driver-e1/",
            "direct": [],
            "cloud": ["с2с-hub"]
        },
        {
            "manufacturer": "aqara",
            "model": "CD-M01D",
            "name": "Aqara Smart Curtain Controller. Подключается через хаб Aqara",
            "categories": ["window-blind"],
            "protocols": ["zigbee"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/aqara-curtain-controller-2.svg",
            "link": "https://aqara.ru/product/мотор-для-раздвижных-штор-aqara-curtain-controller-2/",
            "direct": [],
            "cloud": ["с2с-hub"]
        },
        {
            "manufacturer": "aqara",
            "model": "ZNLDP12LM",
            "name": "Aqara LED Light Bulb. Подключается через хаб Aqara",
            "categories": ["light"],
            "protocols": ["zigbee"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/b44051b1_aqara-light-1.jpeg",
            "link": "https://aqara.ru/product/aqara-led-light-bulb/",
            "direct": [],
            "cloud": ["с2с-hub"]
        },
        {
            "manufacturer": "aqara",
            "model": "LEDLBT1-L01",
            "name": "Aqara LED Bulb T1 (Tunable White). Подключается через хаб Aqara",
            "categories": ["light"],
            "protocols": ["zigbee"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/aqara-light-bulb-t1.svg",
            "link": "https://aqara.ru/product/умная-лампа-aqara-light-bulb-t1/",
            "direct": [],
            "cloud": ["с2с-hub"]
        },
        {
            "manufacturer": "aqara",
            "model": "MZSD11LM",
            "name": "Aqara Surface mounted spotlight T1 (24). Подключается через хаб Aqara",
            "categories": ["light"],
            "protocols": ["zigbee"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/aqara-light-t1.svg",
            "link": "https://aqara.ru/product/накладные-точечные-светильники-aqara-t1-с-уг/",
            "direct": [],
            "cloud": ["с2с-hub"]
        },
        {
            "manufacturer": "aqara",
            "model": "MZSD12LM",
            "name": "Aqara Surface mounted spotlight T1 (36). Подключается через хаб Aqara",
            "categories": ["light"],
            "protocols": ["zigbee"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/aqara-light-t1.svg",
            "link": "https://aqara.ru/product/накладные-точечные-светильники-aqara-t1-с-уг/",
            "direct": [],
            "cloud": ["с2с-hub"]
        },
        {
            "manufacturer": "aqara",
            "model": "MZTD11LM",
            "name": "Aqara Surface mounted spotlight T1 (60). Подключается через хаб Aqara",
            "categories": ["light"],
            "protocols": ["zigbee"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/aqara-light-t1.svg",
            "link": "https://aqara.ru/product/накладные-точечные-светильники-aqara-t1-с-уг/",
            "direct": [],
            "cloud": ["с2с-hub"]
        },
        {
            "manufacturer": "aqara",
            "model": "CM-M01",
            "name": "Aqara Curtain Driver E1. Подключается через хаб Aqara",
            "categories": ["curtain"],
            "protocols": ["zigbee"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/fdb56ecc_aqara-curtain-1-1.svg",
            "link": "https://aqara.ru/product/мотор-раздвижных-штор-e1-aqara-curtain-driver-e1/",
            "direct": [],
            "cloud": ["с2с-hub"]
        },
        {
            "manufacturer": "aqara",
            "model": "ZNCLDJ11LM",
            "name": "Aqara Curtain Controller. Подключается через хаб Aqara",
            "categories": ["curtain"],
            "protocols": ["zigbee"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/d4bbb971_aqara-curtain-2.svg",
            "link": "https://aqara.ru/product/мотор-для-раздвижных-штор-aqara-curtain-controller/",
            "direct": [],
            "cloud": ["с2с-hub"]
        },
        {
            "manufacturer": "aqara",
            "model": "SP-EUC01",
            "name": "Aqara Smart Plug. Подключается через хаб Aqara, Zigbee-колонку Sber или хаб Sber",
            "categories": ["socket"],
            "protocols": ["zigbee"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/26bec896_aqara-socket-1.jpeg",
            "link": "https://aqara.ru/product/aqara-smart-plug/",
            "direct": ["hub"],
            "cloud": ["с2с-hub"]
        },
        {
            "manufacturer": "aqara",
            "model": "WP-P01D",
            "name": "Aqara Wall outlet H2 EU. Подключается через хаб Aqara, Zigbee-колонку Sber или хаб Sber",
            "categories": ["socket"],
            "protocols": ["zigbee"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/aqara-wall-outlet-h2-eu.svg",
            "link": "https://aqara.ru/product/умная-встраиваемая-zigbee-розетка-aqara-wall-outlet-h2-eu/",
            "direct": ["hub"],
            "cloud": ["с2с-hub"]
        },
        {
            "manufacturer": "aqara",
            "model": "WXKG11LM",
            "name": "Aqara Wireless Mini Switch. Подключается через хаб Aqara",
            "categories": ["scenario-button"],
            "protocols": ["zigbee"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/5f91f680_aqara-scenario-button-1.jpeg",
            "link": "https://aqara.ru/product/aqara-wireless-mini-switch/",
            "direct": [],
            "cloud": ["с2с-hub"]
        },
        {
            "manufacturer": "aqara",
            "model": "WB-R02D",
            "name": "Aqara Wireless mini switch T1. Подключается через хаб Aqara, Zigbee-колонку Sber или хаб Sber",
            "categories": ["scenario-button"],
            "protocols": ["zigbee"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/aqara-wireless-mini-switch-t1.svg",
            "link": "https://aqara.ru/product/беспроводная-кнопка-aqara-wireless-mini-switch-t1/",
            "direct": ["hub"],
            "cloud": ["с2с-hub"]
        },
        {
            "manufacturer": "aqara",
            "model": "SRTS-A01",
            "name": "Aqara Smart Radiator Thermostat E1. Подключается через хаб Aqara",
            "categories": ["hvac-radiator"],
            "protocols": ["zigbee"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/08cf9dc2_aqara-hvac-radiator-1.svg",
            "link": "https://aqara.ru/product/терморегулятор-для-радиатора-aqara-e1/",
            "direct": [],
            "cloud": ["с2с-hub"]
        },
        {
            "manufacturer": "aqara",
            "model": "ZHWG11LM",
            "name": "Aqara Hub",
            "categories": ["hub"],
            "protocols": ["zigbee"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/99c389ef_aqara-hub-1.jpeg",
            "link": "https://aqara.ru/product/aqara-hub/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "aqara",
            "model": "HE1-G01",
            "name": "Aqara Hub E1",
            "categories": ["hub"],
            "protocols": ["zigbee"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/b38414e8_aqara-hub-2.png",
            "link": "https://aqara.ru/product/usb-центр-управления-умным-домом-aqara-hub-e1/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "aqara",
            "model": "HM1S-G01",
            "name": "Aqara Hub M1S",
            "categories": ["hub"],
            "protocols": ["zigbee"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/d4eceb96_aqara-hub-3.png",
            "link": "https://aqara.ru/product/центр-умного-дома-aqara-hub-m1s-eu/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "aqara",
            "model": "HM2-G01",
            "name": "Aqara Hub M2",
            "categories": ["hub"],
            "protocols": ["zigbee"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/6e5b7884_aqara-hub-4.jpeg",
            "link": "https://aqara.ru/product/центр-умного-дома-aqara-hub-m2/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "aqara",
            "model": "HM1S-G02",
            "name": "Aqara Hub M1S Gen2",
            "categories": ["hub"],
            "protocols": ["zigbee"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/aqara-hub-m1s-gen-2.svg",
            "link": "https://aqara.ru/product/центр-умного-дома-aqara-hub-m1s-gen-2/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "aurum",
            "model": "",
            "name": "Кондиционеры Aurum. Для управления по Wi-Fi должны быть оснащены модулем Daichi DW21-B/DW22-B",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/aurum-hvac-ac.svg",
            "link": "https://daichicloud.ru/split-lineup/search.php?search=AURUM",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "axioma",
            "model": "",
            "name": "Кондиционеры Axioma. Для управления по Wi-Fi должны быть оснащены модулем Daichi DW21-B/DW22-B",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/axioma-hvac-ac.svg",
            "link": "https://daichicloud.ru/split-lineup/search.php?search=Axioma",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "ballu",
            "model": "Boho Full-DC",
            "name": "Кондиционеры серии Ballu Boho DC. Для управления по Wi-Fi должны быть оснащены модулем Hommyn HDN/WFN-02-01",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/f2af7eae_ballu-boho-full-dc.svg",
            "link": "https://www.ballu.ru/catalog/tekhnika_dlya_doma_i_ofisa/konditsionery_vozdukha/dc_invertornye_split_sistemy/seriya_boho_full_dc_inverter/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "ballu",
            "model": "Eco Smart DC",
            "name": "Кондиционеры серии Ballu Eco Smart DC. Для управления по Wi-Fi должны быть оснащены модулем Hommyn HDN/WFN-02-01",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/db20a075_ballu-eco-smart-dc.svg",
            "link": "https://www.ballu.ru/catalog/tekhnika_dlya_doma_i_ofisa/konditsionery_vozdukha/dc_invertornye_split_sistemy/seriya_eco_smart_dc_inverter/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "ballu",
            "model": "Ice Peak Full-DC",
            "name": "Кондиционеры серии Ballu Ice Peak DC. Для управления по Wi-Fi должны быть оснащены модулем Hommyn HDN/WFN-02-01",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/154750e9_ballu-ice-peak-full-dc.svg",
            "link": "https://www.ballu.ru/catalog/tekhnika_dlya_doma_i_ofisa/konditsionery_vozdukha/dc_invertornye_split_sistemy/seriya_ice_peak_full_dc_inverter/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "ballu",
            "model": "BCT/EVU-3.1I",
            "name": "Блок управления Transformer Digital Inverter Ballu BCT/EVU-3.1I",
            "categories": ["hvac-heater"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/b335d9e2_hommyn-hvac-heater-1.svg",
            "link": "https://www.ballu.ru/catalog/tekhnika_dlya_doma_i_ofisa/obogrevateli/elektricheskie_konvektory/bloki_upravleniya_dlya_elektricheskikh_konvektorov/blok_upravleniya_transformer_digital_inverter_ballu_bct_evu_3_1i/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "ballu",
            "model": "BCT/EVU-4I",
            "name": "Блок управления Transformer Digital Inverter Ballu BCT/EVU-4I",
            "categories": ["hvac-heater"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/fcc35be4_ballu-hvac-heater-4i.svg",
            "link": "https://www.ballu.ru/catalog/tekhnika_dlya_doma_i_ofisa/obogrevateli/elektricheskie_konvektory/bloki_upravleniya_dlya_elektricheskikh_konvektorov/blok_upravleniya_transformer_digital_inverter_ballu_bct_evu_4i/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "ballu",
            "model": "BEC/ATI-1503",
            "name": "Конвектор электрический Ballu Apollo digital INVERTER Black Infinity BEC/ATI-1503",
            "categories": ["hvac-heater"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/1c59567a_hommyn-hvac-heater-3.png",
            "link": "https://www.ballu.ru/catalog/tekhnika_dlya_doma_i_ofisa/obogrevateli/elektricheskie_konvektory/seriya_apollo_digital_inverter/konvektor_elektricheskiy_ballu_apollo_digital_inverter_black_infinity_bec_ati_1503/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "ballu",
            "model": "BEC/ATI-2003",
            "name": "Конвектор электрический Ballu Apollo digital INVERTER Black Infinity BEC/ATI-2003",
            "categories": ["hvac-heater"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/1c59567a_hommyn-hvac-heater-3.png",
            "link": "https://www.ballu.ru/catalog/tekhnika_dlya_doma_i_ofisa/obogrevateli/elektricheskie_konvektory/seriya_apollo_digital_inverter/konvektor_elektricheskiy_ballu_apollo_digital_inverter_black_infinity_bec_ati_2003/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "ballu",
            "model": "BEC/ATI-2503",
            "name": "Конвектор электрический Ballu Apollo digital INVERTER Black Infinity BEC/ATI-2503",
            "categories": ["hvac-heater"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/1c59567a_hommyn-hvac-heater-3.png",
            "link": "https://www.ballu.ru/catalog/tekhnika_dlya_doma_i_ofisa/obogrevateli/elektricheskie_konvektory/seriya_apollo_digital_inverter/konvektor_elektricheskiy_ballu_apollo_digital_inverter_black_infinity_bec_ati_2503/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "ballu",
            "model": "BEC/ATI-1501",
            "name": "Конвектор электрический Ballu Apollo digital INVERTER Moon Gray BEC/ATI-1501",
            "categories": ["hvac-heater"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/f39b9d6d_hommyn-hvac-heater-4.jpeg",
            "link": "https://www.ballu.ru/catalog/tekhnika_dlya_doma_i_ofisa/obogrevateli/elektricheskie_konvektory/seriya_apollo_digital_inverter/konvektor_elektricheskiy_ballu_apollo_digital_inverter_moon_gray_bec_ati_1501/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "ballu",
            "model": "BEC/ATI-2001",
            "name": "Конвектор электрический Ballu Apollo digital INVERTER Moon Gray BEC/ATI-2001",
            "categories": ["hvac-heater"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/f39b9d6d_hommyn-hvac-heater-4.jpeg",
            "link": "https://www.ballu.ru/catalog/tekhnika_dlya_doma_i_ofisa/obogrevateli/elektricheskie_konvektory/seriya_apollo_digital_inverter/konvektor_elektricheskiy_ballu_apollo_digital_inverter_moon_gray_bec_ati_2001/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "ballu",
            "model": "BEC/ATI-2501",
            "name": "Конвектор электрический Ballu Apollo digital INVERTER Moon Gray BEC/ATI-2501",
            "categories": ["hvac-heater"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/f39b9d6d_hommyn-hvac-heater-4.jpeg",
            "link": "https://www.ballu.ru/catalog/tekhnika_dlya_doma_i_ofisa/obogrevateli/elektricheskie_konvektory/seriya_apollo_digital_inverter/konvektor_elektricheskiy_ballu_apollo_digital_inverter_moon_gray_bec_ati_2501/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "ballu",
            "model": "BEC/ATI-1502",
            "name": "Конвектор электрический Ballu Apollo digital INVERTER Space Black BEC/ATI-1502",
            "categories": ["hvac-heater"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/f743c248_hommyn-hvac-heater-5.jpeg",
            "link": "https://www.ballu.ru/catalog/tekhnika_dlya_doma_i_ofisa/obogrevateli/elektricheskie_konvektory/seriya_apollo_digital_inverter/konvektor_elektricheskiy_ballu_apollo_digital_inverter_space_black_bec_ati_1502/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "ballu",
            "model": "BEC/ATI-2002",
            "name": "Конвектор электрический Ballu Apollo digital INVERTER Space Black BEC/ATI-2002",
            "categories": ["hvac-heater"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/f743c248_hommyn-hvac-heater-5.jpeg",
            "link": "https://www.ballu.ru/catalog/tekhnika_dlya_doma_i_ofisa/obogrevateli/elektricheskie_konvektory/seriya_apollo_digital_inverter/konvektor_elektricheskiy_ballu_apollo_digital_inverter_space_black_bec_ati_2002/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "ballu",
            "model": "BEC/ATI-2502",
            "name": "Конвектор электрический Ballu Apollo digital INVERTER Space Black BEC/ATI-2502",
            "categories": ["hvac-heater"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/f743c248_hommyn-hvac-heater-5.jpeg",
            "link": "https://www.ballu.ru/catalog/tekhnika_dlya_doma_i_ofisa/obogrevateli/elektricheskie_konvektory/seriya_apollo_digital_inverter/konvektor_elektricheskiy_ballu_apollo_digital_inverter_space_black_bec_ati_2502/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "bosch",
            "model": "",
            "name": "Кондиционеры Bosch. Для управления по Wi-Fi должны быть оснащены модулем Daichi DW21-B/DW22-B",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/bosch-hvac-ac.svg",
            "link": "https://daichicloud.ru/split-lineup/search.php?search=Bosch",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "candy",
            "model": "AABFH9E169RU",
            "name": "Кондиционер Candy AC-07HTA303/R2-IN",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/candy-hvac-ac-1.svg",
            "link": "https://candy-home.ru/catalog/ventilation-and-heating-candy/acs_cd/split-sistema-candy-t-on-off-ac-07hta303-r2/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "candy",
            "model": "AABFH9E171RU",
            "name": "Кондиционер Candy AC-09HTA303/R2-IN",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/candy-hvac-ac-2.svg",
            "link": "https://candy-home.ru/catalog/ventilation-and-heating-candy/acs_cd/split-sistema-candy-t-on-off-ac-09hta303-r2/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "candy",
            "model": "AABFH9E173RU",
            "name": "Кондиционер Candy AC-12HTA303/R2-IN",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/candy-hvac-ac-2.svg",
            "link": "https://candy-home.ru/catalog/ventilation-and-heating-candy/acs_cd/split-sistema-candy-t-on-off-ac-12hta303-r2/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "candy",
            "model": "AABFH9E175RU",
            "name": "Кондиционер Candy AC-18HTA303/R2-IN",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/candy-hvac-ac-2.svg",
            "link": "https://candy-home.ru/catalog/ventilation-and-heating-candy/acs_cd/split-sistema-candy-t-on-off-ac-18hta303-r2/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "candy",
            "model": "AABFH9E177RU",
            "name": "Кондиционер Candy AC-24HTA303/R2-IN",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/candy-hvac-ac-2.svg",
            "link": "https://candy-home.ru/catalog/ventilation-and-heating-candy/acs_cd/split-sistema-candy-t-on-off-ac-24hta303-r2/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "candy",
            "model": "AABFH9E179RU",
            "name": "Кондиционер Candy ACI-09HRR203/R3-IN",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/candy-hvac-ac-3.svg",
            "link": "https://candy-home.ru/catalog/ventilation-and-heating-candy/acs_cd/split-sistema-candy-inverter-aci-09hrr203-r3/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "candy",
            "model": "AABFH9E181RU",
            "name": "Кондиционер Candy ACI-12HRR203/R3-IN",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/candy-hvac-ac-3.svg",
            "link": "https://candy-home.ru/catalog/ventilation-and-heating-candy/acs_cd/split-sistema-candy-inverter-aci-12hrr203-r3/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "candy",
            "model": "AABFH9E183RU",
            "name": "Кондиционер Candy ACI-18HRR203/R3-IN",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/candy-hvac-ac-3.svg",
            "link": "https://candy-home.ru/catalog/ventilation-and-heating-candy/acs_cd/split-sistema-candy-inverter-aci-18hrr203-r3/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "candy",
            "model": "AABFH9E185RU",
            "name": "Кондиционер Candy ACI-24HRR203/R3-IN",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/candy-hvac-ac-3.svg",
            "link": "https://candy-home.ru/catalog/ventilation-and-heating-candy/acs_cd/split-sistema-candy-inverter-aci-24hrr203-r3/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "casarte",
            "model": "",
            "name": "Кондиционер Casarte CAS25CX1/R3-W",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/casarte-hvac-ac-5.svg",
            "link": "https://casarte.ru/catalog/air-conditioners/eletto/belyj-matovyj-25-kv-m/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "casarte",
            "model": "",
            "name": "Кондиционер Casarte CAS25CX1/R3-G",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/casarte-hvac-ac-6.svg",
            "link": "https://casarte.ru/catalog/air-conditioners/eletto/belyj-matovyj-25-kv-m/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "casarte",
            "model": "",
            "name": "Кондиционер Casarte CAS35CX1/R3-W",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/casarte-hvac-ac-5.svg",
            "link": "https://casarte.ru/catalog/air-conditioners/eletto/belyj-matovyj-35-kv-m/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "casarte",
            "model": "",
            "name": "Кондиционер Casarte CAS50CX1/R3-W",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/casarte-hvac-ac-5.svg",
            "link": "https://casarte.ru/catalog/air-conditioners/eletto/belyj-matovyj-50-kv-m/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "casarte",
            "model": "",
            "name": "Кондиционер Casarte CAS25CX1/R3-B",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/casarte-hvac-ac-7.svg",
            "link": "https://casarte.ru/catalog/air-conditioners/eletto/chernyj-matovyj-25-kv-m/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "casarte",
            "model": "",
            "name": "Кондиционер Casarte CAS35CX1/R3-B",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/casarte-hvac-ac-7.svg",
            "link": "https://casarte.ru/catalog/air-conditioners/eletto/chernyj-matovyj-35-kv-m/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "casarte",
            "model": "",
            "name": "Кондиционер Casarte CAS50CX1/R3-B",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/casarte-hvac-ac-7.svg",
            "link": "https://casarte.ru/catalog/air-conditioners/eletto/chernyj-matovyj-50-kv-m/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "casarte",
            "model": "",
            "name": "Кондиционер Casarte CAS35CX1/R3-G",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/casarte-hvac-ac-6.svg",
            "link": "https://casarte.ru/catalog/air-conditioners/eletto/shampan-35-kv-m/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "casarte",
            "model": "",
            "name": "Кондиционер Casarte CAS50CX1/R3-G",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/casarte-hvac-ac-6.svg",
            "link": "https://casarte.ru/catalog/air-conditioners/eletto/shampan-50-kv-m/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "casarte",
            "model": "",
            "name": "Кондиционер Casarte CAS25MW1/R3-W",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/casarte-hvac-ac-1.svg",
            "link": "https://casarte.ru/catalog/air-conditioners/triano/glyancevyj-belyj-25-kv-m/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "casarte",
            "model": "",
            "name": "Кондиционер Casarte CAS35MW1/R3-W",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/casarte-hvac-ac-4.svg",
            "link": "https://casarte.ru/catalog/air-conditioners/triano/glyancevyj-belyj-35-kv-m/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "casarte",
            "model": "",
            "name": "Кондиционер Casarte CAS35MW1/R3-B",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/casarte-hvac-ac-3.svg",
            "link": "https://casarte.ru/catalog/air-conditioners/triano/glyancevyj-belyj-35-kv-m/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "casarte",
            "model": "",
            "name": "Кондиционер Casarte CAS25MW1/R3-B",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/casarte-hvac-ac-3.svg",
            "link": "https://casarte.ru/catalog/air-conditioners/triano/goluboj-antracit-25-kv-m/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "casarte",
            "model": "",
            "name": "Кондиционер Casarte CAS25MW1/R3-G",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/casarte-hvac-ac-2.svg",
            "link": "https://casarte.ru/catalog/air-conditioners/triano/shampan-25-kv-m/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "casarte",
            "model": "",
            "name": "Кондиционер Casarte CAS35MW1/R3-G",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/casarte-hvac-ac-2.svg",
            "link": "https://casarte.ru/catalog/air-conditioners/triano/shampan-35-kv-m/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "centek",
            "model": "",
            "name": "Кондиционеры Centek. Для управления по Wi-Fi должны быть оснащены модулем Daichi DW21-B/DW22-B",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/centek-hvac-ac.svg",
            "link": "https://daichicloud.ru/split-lineup/search.php?search=Centek",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "сitilux",
            "model": "CLR6S",
            "name": "Citilux Контроллер CWRGB-BK",
            "categories": ["controller"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/5f2933d4_sitilux-controller.svg",
            "link": "https://citilux.ru/store/clr6s/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "сitilux",
            "model": "CL228A931",
            "name": "Citilux Умный торшер CWRGB",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/f38a077e_sitilux-light-2.svg",
            "link": "https://citilux.ru/store/cl228a931/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "сitilux",
            "model": "CL228A031",
            "name": "Citilux Умная люстра CWRGB",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/sitilux-light-13.svg",
            "link": "https://citilux.ru/store/CL228A031/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "сitilux",
            "model": "CL228A051",
            "name": "Citilux Умная люстра CWRGB",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/sitilux-light-14.svg",
            "link": "https://citilux.ru/store/CL228A051/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "сitilux",
            "model": "CL228A141",
            "name": "Citilux Умная люстра CWRGB",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/sitilux-light-94.svg",
            "link": "https://citilux.ru/store/CL228A141/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "сitilux",
            "model": "CL228A161",
            "name": "Citilux Умная люстра CWRGB",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/sitilux-light-16.svg",
            "link": "https://citilux.ru/store/CL228A161/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "сitilux",
            "model": "CL228A181",
            "name": "Citilux Умная люстра CWRGB",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/sitilux-light-17.svg",
            "link": "https://citilux.ru/store/CL228A181/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "сitilux",
            "model": "CL228A811",
            "name": "Citilux Умная лампа CWRGB",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/sitilux-light-18.svg",
            "link": "https://citilux.ru/store/CL228A811/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "сitilux",
            "model": "CL703A100G",
            "name": "Citilux Умная люстра CWRGB",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/sitilux-light-64.svg",
            "link": "https://citilux.ru/store/CL703A100G/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "сitilux",
            "model": "CL703A101G",
            "name": "Citilux Умная люстра CWRGB",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/sitilux-light-64.svg",
            "link": "https://citilux.ru/store/CL703A101G/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "сitilux",
            "model": "CL703A103G",
            "name": "Citilux Умная люстра CWRGB",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/sitilux-light-65.svg",
            "link": "https://citilux.ru/store/CL703A103G/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "сitilux",
            "model": "CL703A105G",
            "name": "Citilux Умная люстра CWRGB",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/sitilux-light-66.svg",
            "link": "https://citilux.ru/store/CL703A105G/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "сitilux",
            "model": "CL703A10G",
            "name": "Citilux Умная люстра CWRGB",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/sitilux-light-67.svg",
            "link": "https://citilux.ru/store/CL703A10G/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "сitilux",
            "model": "CL703A11G",
            "name": "Citilux Умная люстра CWRGB",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/sitilux-light-68.svg",
            "link": "https://citilux.ru/store/CL703A11G/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "сitilux",
            "model": "CL703A15G",
            "name": "Citilux Умная люстра CWRGB",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/sitilux-light-70.svg",
            "link": "https://citilux.ru/store/CL703A15G/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "сitilux",
            "model": "CL703A30G",
            "name": "Citilux Умная люстра CWRGB",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/sitilux-light-71.svg",
            "link": "https://citilux.ru/store/CL703A30G/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "сitilux",
            "model": "CL703A31G",
            "name": "Citilux Умная люстра CWRGB",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/sitilux-light-72.svg",
            "link": "https://citilux.ru/store/CL703A31G/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "сitilux",
            "model": "CL703A33G",
            "name": "Citilux Умная люстра CWRGB",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/sitilux-light-73.svg",
            "link": "https://citilux.ru/store/CL703A33G/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "сitilux",
            "model": "CL703A35G",
            "name": "Citilux Умная люстра CWRGB",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/sitilux-light-74.svg",
            "link": "https://citilux.ru/store/CL703A35G/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "сitilux",
            "model": "CL703A40G",
            "name": "Citilux Умная люстра CWRGB",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/sitilux-light-75.svg",
            "link": "https://citilux.ru/store/CL703A40G/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "сitilux",
            "model": "CL703A41G",
            "name": "Citilux Умная люстра CWRGB",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/sitilux-light-75.svg",
            "link": "https://citilux.ru/store/CL703A41G/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "сitilux",
            "model": "CL703A43G",
            "name": "Citilux Умная люстра CWRGB",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/sitilux-light-76.svg",
            "link": "https://citilux.ru/store/CL703A43G/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "сitilux",
            "model": "CL703A45G",
            "name": "Citilux Умная люстра CWRGB",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/sitilux-light-76.svg",
            "link": "https://citilux.ru/store/CL703A45G/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "сitilux",
            "model": "CL703A60G",
            "name": "Citilux Умная люстра CWRGB",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/sitilux-light-75.svg",
            "link": "https://citilux.ru/store/CL703A60G/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "сitilux",
            "model": "CL703A61G",
            "name": "Citilux Умная люстра CWRGB",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/sitilux-light-71.svg",
            "link": "https://citilux.ru/store/CL703A61G/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "сitilux",
            "model": "CL703A63G",
            "name": "Citilux Умная люстра CWRGB",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/sitilux-light-65.svg",
            "link": "https://citilux.ru/store/CL703A63G/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "сitilux",
            "model": "CL703A65G",
            "name": "Citilux Умная люстра CWRGB",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/sitilux-light-74.svg",
            "link": "https://citilux.ru/store/CL703A65G/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "сitilux",
            "model": "CL703A80G",
            "name": "Citilux Умная люстра CWRGB",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/sitilux-light-68.svg",
            "link": "https://citilux.ru/store/CL703A80G/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "сitilux",
            "model": "CL703A81G",
            "name": "Citilux Умная люстра CWRGB",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/sitilux-light-68.svg",
            "link": "https://citilux.ru/store/CL703A81G/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "сitilux",
            "model": "CL703A83G",
            "name": "Citilux Умная люстра CWRGB",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/sitilux-light-73.svg",
            "link": "https://citilux.ru/store/CL703A83G/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "сitilux",
            "model": "CL703A85G",
            "name": "Citilux Умная люстра CWRGB",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/sitilux-light-74.svg",
            "link": "https://citilux.ru/store/CL703A85G/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "сitilux",
            "model": "CL703A140G",
            "name": "Citilux Умная люстра CWRGB",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/sitilux-light-98.svg",
            "link": "https://citilux.ru/store/CL703A140G/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "сitilux",
            "model": "CL703A141G",
            "name": "Citilux Умная люстра CWRGB",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/sitilux-light-98.svg",
            "link": "https://citilux.ru/store/CL703A141G/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "сitilux",
            "model": "CL703A143G",
            "name": "Citilux Умная люстра CWRGB",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/sitilux-light-98.svg",
            "link": "https://citilux.ru/store/CL703A143G/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "сitilux",
            "model": "CL703A145G",
            "name": "Citilux Умная люстра CWRGB",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/sitilux-light-98.svg",
            "link": "https://citilux.ru/store/CL703A145G/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "сitilux",
            "model": "CL703A200G",
            "name": "Citilux Умная люстра CWRGB",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/sitilux-light-98.svg",
            "link": "https://citilux.ru/store/CL703A200G/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "сitilux",
            "model": "CL703A201G",
            "name": "Citilux Умная люстра CWRGB",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/sitilux-light-98.svg",
            "link": "https://citilux.ru/store/CL703A201G/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "сitilux",
            "model": "CL703A203G",
            "name": "Citilux Умная люстра CWRGB",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/sitilux-light-98.svg",
            "link": "https://citilux.ru/store/CL703A203G/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "сitilux",
            "model": "CL703A205G",
            "name": "Citilux Умная люстра CWRGB",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/sitilux-light-98.svg",
            "link": "https://citilux.ru/store/CL703A205G/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "сitilux",
            "model": "CL703AK50G",
            "name": "Citilux Умная люстра CWRGB",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/sitilux-light-77.svg",
            "link": "https://citilux.ru/store/CL703AK50G/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "сitilux",
            "model": "CL703AK51G",
            "name": "Citilux Умная люстра CWRGB",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/sitilux-light-77.svg",
            "link": "https://citilux.ru/store/CL703AK51G/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "сitilux",
            "model": "CL703AK53G",
            "name": "Citilux Умная люстра CWRGB",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/sitilux-light-78.svg",
            "link": "https://citilux.ru/store/CL703AK53G/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "сitilux",
            "model": "CL703AK55G",
            "name": "Citilux Умная люстра CWRGB",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/sitilux-light-79.svg",
            "link": "https://citilux.ru/store/CL703AK55G/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "сitilux",
            "model": "CL703AK80G",
            "name": "Citilux Умная люстра CWRGB",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/sitilux-light-79.svg",
            "link": "https://citilux.ru/store/CL703AK80G/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "сitilux",
            "model": "CL703AK81G",
            "name": "Citilux Умная люстра CWRGB",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/sitilux-light-79.svg",
            "link": "https://citilux.ru/store/CL703AK81G/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "сitilux",
            "model": "CL703AK83G",
            "name": "Citilux Умная люстра CWRGB",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/sitilux-light-80.svg",
            "link": "https://citilux.ru/store/CL703AK83G/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "сitilux",
            "model": "CL703AK85G",
            "name": "Citilux Умная люстра CWRGB",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/sitilux-light-81.svg",
            "link": "https://citilux.ru/store/CL703AK85G/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "сitilux",
            "model": "CL713A100G",
            "name": "Citilux Умная люстра CWRGB",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/sitilux-light-82.svg",
            "link": "https://citilux.ru/store/CL713A100G/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "сitilux",
            "model": "CL713A10G",
            "name": "Citilux Умная люстра CWRGB",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/sitilux-light-83.svg",
            "link": "https://citilux.ru/store/CL713A10G/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "сitilux",
            "model": "CL713A30G",
            "name": "Citilux Умная люстра CWRGB",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/sitilux-light-82.svg",
            "link": "https://citilux.ru/store/CL713A30G/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "сitilux",
            "model": "CL713A40G",
            "name": "Citilux Умная люстра CWRGB",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/sitilux-light-84.svg",
            "link": "https://citilux.ru/store/CL713A40G/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "сitilux",
            "model": "CL713A60G",
            "name": "Citilux Умная люстра CWRGB",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/sitilux-light-84.svg",
            "link": "https://citilux.ru/store/CL713A60G/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "сitilux",
            "model": "CL713A80G",
            "name": "Citilux Умная люстра CWRGB",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/sitilux-light-84.svg",
            "link": "https://citilux.ru/store/CL713A80G/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "сitilux",
            "model": "CL718A100G",
            "name": "Citilux Умная люстра CWRGB",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/sitilux-light-85.svg",
            "link": "https://citilux.ru/store/CL718A100G/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "сitilux",
            "model": "CL718A12G",
            "name": "Citilux Умная люстра CWRGB",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/sitilux-light-86.svg",
            "link": "https://citilux.ru/store/CL718A12G/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "сitilux",
            "model": "CL718A40G",
            "name": "Citilux Умная люстра CWRGB",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/sitilux-light-87.svg",
            "link": "https://citilux.ru/store/CL718A40G/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "сitilux",
            "model": "CL718A60G",
            "name": "Citilux Умная люстра CWRGB",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/sitilux-light-88.svg",
            "link": "https://citilux.ru/store/CL718A60G/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "сitilux",
            "model": "CL718A80G",
            "name": "Citilux Умная люстра CWRGB",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/sitilux-light-85.svg",
            "link": "https://citilux.ru/store/CL718A80G/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "сitilux",
            "model": "CL732A520G",
            "name": "Citilux Умная люстра CWRGB",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/sitilux-light-89.svg",
            "link": "https://citilux.ru/store/CL732A520G/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "сitilux",
            "model": "CL732A660G",
            "name": "Citilux Умная люстра CWRGB",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/sitilux-light-89.svg",
            "link": "https://citilux.ru/store/CL732A660G/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "сitilux",
            "model": "CL732A800G",
            "name": "Citilux Умная люстра CWRGB",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/sitilux-light-90.svg",
            "link": "https://citilux.ru/store/CL732A800G/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "сitilux",
            "model": "CL225A140E",
            "name": "Citilux Умная люстра CW+RGB",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/sitilux-light-1.svg",
            "link": "https://citilux.ru/store/cl225a140e/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "сitilux",
            "model": "CL225A145E",
            "name": "Citilux Умная люстра CW+RGB",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/sitilux-light-2.svg",
            "link": "https://citilux.ru/store/cl225a145e/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "сitilux",
            "model": "CL225A160E",
            "name": "Citilux Умная люстра CW+RGB",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/sitilux-light-3.svg",
            "link": "https://citilux.ru/store/CL225A160E/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "сitilux",
            "model": "CL225A165E",
            "name": "Citilux Умная люстра CW+RGB",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/sitilux-light-4.svg",
            "link": "https://citilux.ru/store/CL225A165E/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "сitilux",
            "model": "CL225A190E",
            "name": "Citilux Умная люстра CW+RGB",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/sitilux-light-5.svg",
            "link": "https://citilux.ru/store/CL225A190E/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "сitilux",
            "model": "CL225A195E",
            "name": "Citilux Умная люстра CW+RGB",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/sitilux-light-6.svg",
            "link": "https://citilux.ru/store/CL225A195E/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "сitilux",
            "model": "CL225A240E",
            "name": "Citilux Умная люстра CW+RGB",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/sitilux-light-7.svg",
            "link": "https://citilux.ru/store/CL225A240E/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "сitilux",
            "model": "CL225A245E",
            "name": "Citilux Умная люстра CW+RGB",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/sitilux-light-8.svg",
            "link": "https://citilux.ru/store/CL225A245E/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "сitilux",
            "model": "CL225A250E",
            "name": "Citilux Умная люстра CW+RGB",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/sitilux-light-9.svg",
            "link": "https://citilux.ru/store/CL225A250E/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "сitilux",
            "model": "CL225A255E",
            "name": "Citilux Умная люстра CW+RGB",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/sitilux-light-10.svg",
            "link": "https://citilux.ru/store/CL225A255E/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "сitilux",
            "model": "CL225A290E",
            "name": "Citilux Умная люстра CW+RGB",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/sitilux-light-11.svg",
            "link": "https://citilux.ru/store/CL225A290E/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "сitilux",
            "model": "CL225A295E",
            "name": "Citilux Умная люстра CW+RGB",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/sitilux-light-12.svg",
            "link": "https://citilux.ru/store/CL225A295E/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "сitilux",
            "model": "CL229A151E",
            "name": "Citilux Умная люстра CW+RGB",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/sitilux-light-92.svg",
            "link": "https://citilux.ru/store/CL229A151E/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "сitilux",
            "model": "CL229A155E",
            "name": "Citilux Умная люстра CW+RGB",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/sitilux-light-93.svg",
            "link": "https://citilux.ru/store/CL229A155E/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "сitilux",
            "model": "CL229A161E",
            "name": "Citilux Умная люстра CW+RGB",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/sitilux-light-92.svg",
            "link": "https://citilux.ru/store/CL229A161E/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "сitilux",
            "model": "CL229A165E",
            "name": "Citilux Умная люстра CW+RGB",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/sitilux-light-93.svg",
            "link": "https://citilux.ru/store/CL229A165E/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "сitilux",
            "model": "CL232A140E",
            "name": "Citilux Умная люстра CW+RGB",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/sitilux-light-23.svg",
            "link": "https://citilux.ru/store/CL232A140E/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "сitilux",
            "model": "CL232A145E",
            "name": "Citilux Умная люстра CW+RGB",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/sitilux-light-24.svg",
            "link": "https://citilux.ru/store/CL232A145E/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "сitilux",
            "model": "CL232A150E",
            "name": "Citilux Умная люстра CW+RGB",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/sitilux-light-25.svg",
            "link": "https://citilux.ru/store/CL232A150E/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "сitilux",
            "model": "CL232A155E",
            "name": "Citilux Умная люстра CW+RGB",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/sitilux-light-26.svg",
            "link": "https://citilux.ru/store/cl232a155e/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "сitilux",
            "model": "CL232A180E",
            "name": "Citilux Умная люстра CW+RGB",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/sitilux-light-27.svg",
            "link": "https://citilux.ru/store/CL232A180E/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "сitilux",
            "model": "CL232A185E",
            "name": "Citilux Умная люстра CW+RGB",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/sitilux-light-28.svg",
            "link": "https://citilux.ru/store/CL232A185E/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "сitilux",
            "model": "CL233A150E",
            "name": "Citilux Умная люстра CW+RGB",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/sitilux-light-29.svg",
            "link": "https://citilux.ru/store/CL233A150E/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "сitilux",
            "model": "CL233A155E",
            "name": "Citilux Умная люстра CW+RGB",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/sitilux-light-30.svg",
            "link": "https://citilux.ru/store/CL233A155E/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "сitilux",
            "model": "CL233A170E",
            "name": "Citilux Умная люстра CW+RGB",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/sitilux-light-31.svg",
            "link": "https://citilux.ru/store/CL233A170E/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "сitilux",
            "model": "CL233A175E",
            "name": "Citilux Умная люстра CW+RGB",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/sitilux-light-30.svg",
            "link": "https://citilux.ru/store/CL233A175E/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "сitilux",
            "model": "CL233A250E",
            "name": "Citilux Умная люстра CW+RGB",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/sitilux-light-32.svg",
            "link": "https://citilux.ru/store/CL233A250E/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "сitilux",
            "model": "CL233A255E",
            "name": "Citilux Умная люстра CW+RGB",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/sitilux-light-33.svg",
            "link": "https://citilux.ru/store/CL233A255E/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "сitilux",
            "model": "CL233A270E",
            "name": "Citilux Умная люстра CW+RGB",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/sitilux-light-34.svg",
            "link": "https://citilux.ru/store/CL233A270E/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "сitilux",
            "model": "CL233A275E",
            "name": "Citilux Умная люстра CW+RGB",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/sitilux-light-33.svg",
            "link": "https://citilux.ru/store/CL233A275E/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "сitilux",
            "model": "CL234A150E",
            "name": "Citilux Умная люстра CW+RGB",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/sitilux-light-35.svg",
            "link": "https://citilux.ru/store/CL234A150E/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "сitilux",
            "model": "CL234A250E",
            "name": "Citilux Умная люстра CW+RGB",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/sitilux-light-35.svg",
            "link": "https://citilux.ru/store/CL234A250E/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "сitilux",
            "model": "CL234A290E",
            "name": "Citilux Умная люстра CW+RGB",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/sitilux-light-36.svg",
            "link": "https://citilux.ru/store/CL234A290E/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "сitilux",
            "model": "CL235A145E",
            "name": "Citilux Умная люстра CW+RGB",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/sitilux-light-37.svg",
            "link": "https://citilux.ru/store/CL235A145E/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "сitilux",
            "model": "CL235A150E",
            "name": "Citilux Умная люстра CW+RGB",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/sitilux-light-38.svg",
            "link": "https://citilux.ru/store/CL235A150E/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "сitilux",
            "model": "CL235A165E",
            "name": "Citilux Умная люстра CW+RGB",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/sitilux-light-39.svg",
            "link": "https://citilux.ru/store/CL235A165E/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "сitilux",
            "model": "CL235A180E",
            "name": "Citilux Умная люстра CW+RGB",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/sitilux-light-40.svg",
            "link": "https://citilux.ru/store/CL235A180E/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "сitilux",
            "model": "CL235A195E",
            "name": "Citilux Умная люстра CW+RGB",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/sitilux-light-41.svg",
            "link": "https://citilux.ru/store/CL235A195E/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "сitilux",
            "model": "CL235A200E",
            "name": "Citilux Умная люстра CW+RGB",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/sitilux-light-42.svg",
            "link": "https://citilux.ru/store/CL235A200E/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "сitilux",
            "model": "CL236A140E",
            "name": "Citilux Умная люстра CW+RGB",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/sitilux-light-43.svg",
            "link": "https://citilux.ru/store/CL236A140E/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "сitilux",
            "model": "CL236A160E",
            "name": "Citilux Умная люстра CW+RGB",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/sitilux-light-44.svg",
            "link": "https://citilux.ru/store/CL236A160E/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "сitilux",
            "model": "CL236A190E",
            "name": "Citilux Умная люстра CW+RGB",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/sitilux-light-45.svg",
            "link": "https://citilux.ru/store/CL236A190E/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "сitilux",
            "model": "CL719A700",
            "name": "Citilux Умная люстра CW+RGB",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/sitilux-light-46.svg",
            "link": "https://citilux.ru/store/CL719A700/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "сitilux",
            "model": "CL719A701",
            "name": "Citilux Умная люстра CW+RGB",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/sitilux-light-47.svg",
            "link": "https://citilux.ru/store/CL719A701/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "сitilux",
            "model": "CL731A090E",
            "name": "Citilux Умная люстра CW+RGB",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/sitilux-light-48.svg",
            "link": "https://citilux.ru/store/CL731A090E/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "сitilux",
            "model": "CL731A095E",
            "name": "Citilux Умная люстра CW+RGB",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/sitilux-light-49.svg",
            "link": "https://citilux.ru/store/CL731A095E/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "сitilux",
            "model": "CL731A330E",
            "name": "Citilux Умная люстра CW+RGB",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/sitilux-light-50.svg",
            "link": "https://citilux.ru/store/CL731A330E/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "сitilux",
            "model": "CL731AK110E",
            "name": "Citilux Умная люстра CW+RGB",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/sitilux-light-51.svg",
            "link": "https://citilux.ru/store/CL731AK110E/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "сitilux",
            "model": "CL731AK115E",
            "name": "Citilux Умная люстра CW+RGB",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/sitilux-light-52.svg",
            "link": "https://citilux.ru/store/CL731AK115E/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "сitilux",
            "model": "CL737A34E",
            "name": "Citilux Умная люстра CW+RGB",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/sitilux-light-91.svg",
            "link": "https://citilux.ru/store/cl737a34e/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "сitilux",
            "model": "CL737A080E",
            "name": "Citilux Умная люстра CW+RGB",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/sitilux-light-53.svg",
            "link": "https://citilux.ru/store/CL737A080E/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "сitilux",
            "model": "CL737A100E",
            "name": "Citilux Умная люстра CW+RGB",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/sitilux-light-53.svg",
            "link": "https://citilux.ru/store/CL737A100E/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "сitilux",
            "model": "CL737A35E",
            "name": "Citilux Умная люстра CW+RGB",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/sitilux-light-54.svg",
            "link": "https://citilux.ru/store/CL737A35E/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "сitilux",
            "model": "CL737A44E",
            "name": "Citilux Умная люстра CW+RGB",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/sitilux-light-55.svg",
            "link": "https://citilux.ru/store/CL737A44E/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "сitilux",
            "model": "CL737A45E",
            "name": "Citilux Умная люстра CW+RGB",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/sitilux-light-56.svg",
            "link": "https://citilux.ru/store/CL737A45E/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "сitilux",
            "model": "CL737A54E",
            "name": "Citilux Умная люстра CW+RGB",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/sitilux-light-57.svg",
            "link": "https://citilux.ru/store/CL737A54E/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "сitilux",
            "model": "CL737A55E",
            "name": "Citilux Умная люстра CW+RGB",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/sitilux-light-58.svg",
            "link": "https://citilux.ru/store/CL737A55E/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "сitilux",
            "model": "CL710A104S",
            "name": "Citilux Умная люстра CW+UP",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/sitilux-light-95.svg",
            "link": "https://citilux.ru/store/cl710a104s/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "сitilux",
            "model": "CL710A105S",
            "name": "Citilux Умная люстра CW+UP",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/sitilux-light-59.svg",
            "link": "https://citilux.ru/store/cl710a105s/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "сitilux",
            "model": "CL730A090S",
            "name": "Citilux Умная люстра CW+UP",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/sitilux-light-60.svg",
            "link": "https://citilux.ru/store/CL730A090S/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "сitilux",
            "model": "CL730A095S",
            "name": "Citilux Умная люстра CW+UP",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/sitilux-light-61.svg",
            "link": "https://citilux.ru/store/CL730A095S/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "сitilux",
            "model": "CL730A150S",
            "name": "Citilux Умная люстра CW+UP",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/sitilux-light-62.svg",
            "link": "https://citilux.ru/store/CL730A150S/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "сitilux",
            "model": "CL730A155S",
            "name": "Citilux Умная люстра CW+UP",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/sitilux-light-63.svg",
            "link": "https://citilux.ru/store/CL730A155S/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "comfee",
            "model": "",
            "name": "Кондиционеры Comfee. Для управления по Wi-Fi должны быть оснащены модулем Daichi DW21-B/DW22-B",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/comfee-hvac-ac.svg",
            "link": "https://daichicloud.ru/split-lineup/search.php?search=COMFEE",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "coolup",
            "model": "",
            "name": "Кондиционеры CoolUp. Для управления по Wi-Fi должны быть оснащены модулем Daichi DW21-B/DW22-B",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/coolup-hvac-ac.svg",
            "link": "https://daichicloud.ru/split-lineup/search.php?search=CoolUp",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "cooper-hunter",
            "model": "",
            "name": "Кондиционеры Сooper&Hunter. Для управления по Wi-Fi должны быть оснащены модулем Daichi DW21-B/DW22-B",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/cooper-hunter-hvac-ac.svg",
            "link": "https://daichicloud.ru/split-lineup/search.php?search=COOPER&HUNTER",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "daichi",
            "model": "Без модуля управления по Wi-Fi",
            "name": "Кондиционеры Daichi. Для управления по Wi-Fi нужно докупить модуль Daichi DW21-B/DW22-B",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/daichi-hvac-ac-ex-wifi.svg",
            "link": "https://daichicloud.ru/split-lineup/search.php?search=daichi",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "daichi",
            "model": "С модулем управления по Wi-Fi",
            "name": "Кондиционеры Daichi с предустановленным модулем Daichi DW21-B/DW22-B",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/daichi-hvac-ac-wifi.svg",
            "link": "https://daichi.cloud/catalog",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "daichi",
            "model": "38700DC",
            "name": "Wi-Fi-модуль Daichi DW21-B/DW22-B для кондиционеров. Совместим с десятками брендов разных производителей",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/18470f34_daichi-hvac-ac-1.jpeg",
            "link": "https://daichicloud.ru/split-lineup/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "daikin",
            "model": "",
            "name": "Кондиционеры Daikin. Для управления по Wi-Fi должны быть оснащены модулем Daichi DW21-B/DW22-B",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/daikin-hvac-ac.svg",
            "link": "https://daichicloud.ru/split-lineup/search.php?search=DAIKIN",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "dantex",
            "model": "",
            "name": "Кондиционеры Dantex. Для управления по Wi-Fi должны быть оснащены модулем Daichi DW21-B/DW22-B",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/dantex-hvac-ac.svg",
            "link": "https://daichicloud.ru/split-lineup/search.php?search=DANTEX",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "digma",
            "model": "1118520",
            "name": "Умная лампа Digma DiLight E27 N1 E27 8Вт 800lm Wi-Fi",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/73216567_digma-light-1.svg",
            "link": "https://digma.ru/catalog/lifestyle/smart-sockets/smart_bulbs/dilight-e27-n1-1118520/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "digma",
            "model": "1118521",
            "name": "Умная лампа Digma DiLight E27 N1 RGB E27 8Вт 800lm Wi-Fi",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/c9efba13_digma-light-2.svg",
            "link": "https://digma.ru/catalog/lifestyle/smart-sockets/smart_bulbs/dilight-e27-n1-rgb-1118521/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "digma",
            "model": "1118519",
            "name": "Умная лампа Digma DiLight E27 W1 E27 8Вт 850lm Wi-Fi",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/d7cf120c_digma-light-3.svg",
            "link": "https://digma.ru/catalog/lifestyle/smart-sockets/smart_bulbs/dilight-e27-w1-1118519/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "digma",
            "model": "1182653",
            "name": "Сетевой фильтр Digma DiPlug Strip 40 дистанц. вкл/выкл белый",
            "categories": ["socket"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/4beb65c6_digma-socket-1.svg",
            "link": "https://digma.ru/catalog/lifestyle/smart-sockets/smart-extension-cords/diplug-strip-40-1182653/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "digma",
            "model": "1004452",
            "name": "Умная розетка Digma DiPlug 100 EU Wi-Fi белый (DPL100)",
            "categories": ["socket"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/4e036bf8_digma-socket-2.svg",
            "link": "https://digma.ru/catalog/lifestyle/smart-sockets/smart-sockets/diplug-100-1004452/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "digma",
            "model": "1079363",
            "name": "Умная розетка Digma DiPlug 100S EU Wi-Fi белый (DPL101)",
            "categories": ["socket"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/1fd82544_digma-socket-3.svg",
            "link": "https://digma.ru/catalog/lifestyle/smart-sockets/smart-sockets/diplug-100s-1079363/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "digma",
            "model": "1182652",
            "name": "Умная розетка Digma DiPlug 200S EU Wi-Fi белый (DPL200S)",
            "categories": ["socket"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/cd9693ac_digma-socket-4.svg",
            "link": "https://digma.ru/catalog/lifestyle/smart-sockets/smart-sockets/diplug-200s-1182652/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "digma",
            "model": "1408329",
            "name": "Умная розетка Digma DiPlug 400 EU Wi-Fi белый (TY1932)",
            "categories": ["socket"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/703541c7_digma-socket-5.svg",
            "link": "https://digma.ru/catalog/lifestyle/smart-sockets/smart-sockets/diplug-400-1408329/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "digma",
            "model": "1408330",
            "name": "Умная розетка Digma DiPlug 500 EU Wi-Fi белый (TY1910)",
            "categories": ["socket"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/577ff66a_digma-socket-6.svg",
            "link": "https://digma.ru/catalog/lifestyle/smart-sockets/smart-sockets/diplug-500-1408330/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "digma",
            "model": "1099931",
            "name": "Умная розетка Digma diplug110s Wi-Fi белый",
            "categories": ["socket"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/a9ef65a9_digma-socket-7.svg",
            "link": "https://digma.ru/catalog/lifestyle/smart-sockets/smart-sockets/diplug-110s-1099931/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "ecoletta",
            "model": "",
            "name": "Кондиционеры Ecoletta. Для управления по Wi-Fi должны быть оснащены модулем Daichi DW21-B/DW22-B",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/ecoletta-hvac-ac.svg",
            "link": "https://daichicloud.ru/split-lineup/search.php?search=Ecoletta",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "ekf",
            "model": "ss-16a-wf",
            "name": "Умное реле 16А PRO Wi-FI EKF Connect",
            "categories": ["relay"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/ekf-relay-1.svg",
            "link": "https://ekfgroup.com/ru/catalog/products/umnoe-rele-16a-pro-s-monitoringom-potrebleniya-wi-fi",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "ekf",
            "model": "ss-10a-wf",
            "name": "Умное реле 10А Wi-FI EKF Connect",
            "categories": ["relay"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/ekf-relay-1.svg",
            "link": "https://ekfgroup.com/ru/catalog/products/umnoe-rele-10a-wi-fi",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "ekf",
            "model": "sdsh-1g-wf",
            "name": "Умный диммер в подрозетник 1-канальный Wi-Fi EKF Connect",
            "categories": ["relay"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/ekf-relay-2.svg",
            "link": "https://ekfgroup.com/ru/catalog/products/umnyj-dimmer-v-podrozetnik-1-kanalnyj-wi-fi",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "ekf",
            "model": "sdsh-2g-wf",
            "name": "Умный диммер в подрозетник 2-канальный Wi-Fi EKF Connect",
            "categories": ["relay"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/ekf-relay-3.svg",
            "link": "https://ekfgroup.com/ru/catalog/products/umnyj-dimmer-v-podrozetnik-2-kanalnyj-wi-fi",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "ekf",
            "model": "ssh-1g-wf",
            "name": "Умное реле в подрозетник 1-канальное Wi-Fi EKF Connect",
            "categories": ["relay"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/ekf-relay-4.svg",
            "link": "https://ekfgroup.com/ru/catalog/products/umnoe-rele-v-podrozetnik-1-kanalnoe-wi-fi",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "ekf",
            "model": "ssh-2g-wf",
            "name": "Умное реле в подрозетник 2-канальное Wi-Fi EKF Connect",
            "categories": ["relay"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/ekf-relay-5.svg",
            "link": "https://ekfgroup.com/ru/catalog/products/umnoe-rele-v-podrozetnik-2-kanalnoe-wi-fi",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "ekf",
            "model": "RCV-ST1-WD-ZB",
            "name": "Умный выключатель Стокгольм 1-кл. белый Zigbee EKF Сonnect. Подключается через хаб EKF, Zigbee-колонку Sber или хаб Sber",
            "categories": ["relay"],
            "protocols": ["zigbee"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/ekf-relay-1z.svg",
            "link": "https://ekfgroup.com/ru/catalog/products/umnyj-vyklyuchatel-stokgolm-1-kl-belyj-zigbee-ekf-sonnect",
            "direct": ["hub"],
            "cloud": ["с2с-hub"]
        },
        {
            "manufacturer": "ekf",
            "model": "RCV-ST2-WD-ZB",
            "name": "Умный выключатель Стокгольм 2-кл. белый Zigbee EKF Сonnect. Подключается через хаб EKF, Zigbee-колонку Sber или хаб Sber",
            "categories": ["relay"],
            "protocols": ["zigbee"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/ekf-relay-2z.svg",
            "link": "https://ekfgroup.com/ru/catalog/products/umnyj-vyklyuchatel-stokgolm-2-kl-belyj-zigbee-ekf-sonnect",
            "direct": ["hub"],
            "cloud": ["с2с-hub"]
        },
        {
            "manufacturer": "ekf",
            "model": "scsh-1g-wf",
            "name": "Умное реле для штор в подрозетник Wi-Fi EKF Connect",
            "categories": ["relay", "curtain", "window-blind"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/ekf-curtain-1.svg",
            "link": "https://ekfgroup.com/ru/catalog/products/umnoe-rele-dlya-shtor-v-podrozetnik-wi-fi",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "ekf",
            "model": "is-ga-zb",
            "name": "Умный датчик газа Zigbee EKF Connect. Подключается через хаб EKF",
            "categories": ["sensor-gas"],
            "protocols": ["zigbee"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/ekf-sensor-gas-1.svg",
            "link": "https://ekfgroup.com/ru/catalog/products/umnyj-datchik-gaza-zigbee",
            "direct": [],
            "cloud": ["с2с-hub"]
        },
        {
            "manufacturer": "ekf",
            "model": "is-pir-zb-1",
            "name": "Умный датчик движения Zigbee EKF Connect. Подключается через хаб EKF, Zigbee-колонку Sber или хаб Sber",
            "categories": ["sensor-pir"],
            "protocols": ["zigbee"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/ekf-sensor-pir-1.svg",
            "link": "https://ekfgroup.com/ru/catalog/products/umnyj-datchik-dvizheniya-zigbee",
            "direct": ["hub"],
            "cloud": ["с2с-hub"]
        },
        {
            "manufacturer": "ekf",
            "model": "is-sm-zb",
            "name": "Умный датчик дыма Zigbee EKF Connect. Подключается через хаб EKF, Zigbee-колонку Sber или хаб Sber",
            "categories": ["sensor-smoke"],
            "protocols": ["zigbee"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/ekf-sensor-smoke-1.svg",
            "link": "https://ekfgroup.com/ru/catalog/products/umnyj-datchik-dyma-zigbee",
            "direct": ["hub"],
            "cloud": ["с2с-hub"]
        },
        {
            "manufacturer": "ekf",
            "model": "is-dw-zb",
            "name": "Умный датчик открытия Zigbee EKF Connect. Подключается через хаб EKF, Zigbee-колонку Sber или хаб Sber",
            "categories": ["sensor-door"],
            "protocols": ["zigbee"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/ekf-sensor-door-1.svg",
            "link": "https://ekfgroup.com/ru/catalog/products/umnyj-datchik-otkrytiya-zigbee",
            "direct": ["hub"],
            "cloud": ["с2с-hub"]
        },
        {
            "manufacturer": "ekf",
            "model": "is-fl-zb",
            "name": "Умный датчик протечки Zigbee EKF Connect. Подключается через хаб EKF, Zigbee-колонку Sber или хаб Sber",
            "categories": ["sensor-water-leak"],
            "protocols": ["zigbee"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/ekf-sensor-water-leak-1.svg",
            "link": "https://ekfgroup.com/ru/catalog/products/umnayj-datchik-protechki-zigbee",
            "direct": ["hub"],
            "cloud": ["с2с-hub"]
        },
        {
            "manufacturer": "ekf",
            "model": "irr-ths",
            "name": "Умный пульт EKF Connect с датчиками температуры и влажности. Управляет устройствами из категорий, поддерживаемых умным домом Sber",
            "categories": ["sensor-temp"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/ekf-temp-1.svg",
            "link": "https://ekfgroup.com/ru/catalog/products/umnyjpultsdatchikamitemperaturyivlazhnostiekfconnectwi-fi",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "ekf",
            "model": "is-th-nd-zb",
            "name": "Умный датчик температуры и влажности Zigbee EKF Connect. Подключается через хаб EKF, Zigbee-колонку Sber или хаб Sber",
            "categories": ["sensor-temp"],
            "protocols": ["zigbee"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/ekf-sensor-temp-2.svg",
            "link": "https://ekfgroup.com/ru/catalog/products/umnyj-datchik-temperatury-i-vlazhnosti-zigbee-ekf-connect",
            "direct": ["hub"],
            "cloud": ["с2с-hub"]
        },
        {
            "manufacturer": "ekf",
            "model": "is-th-zb",
            "name": "Умный датчик температуры и влажности с экраном Zigbee EKF Connect. Подключается через хаб EKF, Zigbee-колонку Sber или хаб Sber",
            "categories": ["sensor-temp"],
            "protocols": ["zigbee"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/ekf-sensor-temp-3.svg",
            "link": "https://ekfgroup.com/ru/catalog/products/umnyj-datchik-temperatury-i-vlazhnosti-zigbee",
            "direct": ["hub"],
            "cloud": ["с2с-hub"]
        },
        {
            "manufacturer": "ekf",
            "model": "is-thpl-zb",
            "name": "Умный датчик 4 в 1 Zigbee EKF Connect. Подключается через хаб EKF",
            "categories": ["sensor-pir", "sensor-temp"],
            "protocols": ["zigbee"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/ekf-sensor-temp-4.svg",
            "link": "https://ekfgroup.com/ru/catalog/products/umnyj-datchik-4v1-zigbee-ekf-connect",
            "direct": [],
            "cloud": ["с2с-hub"]
        },
        {
            "manufacturer": "ekf",
            "model": "ett-5",
            "name": "Умный терморегулятор для водяного/газового бойлера EKF Connect",
            "categories": ["hvac-boiler"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/ekf-boiler-1.svg",
            "link": "https://ekfgroup.com/ru/catalog/products/umnyj-termoregulyator-dlya-vodyanogo-gazovogo-bojlera-ekf-connect",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "ekf",
            "model": "slwf-gu10-rgbw",
            "name": "Умная лампа GU10 EKF Connect 5W WIFI RGBW",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/de90ac3b_ekf-light-1.svg",
            "link": "https://ekfgroup.com/ru/catalog/products/umnaya-lampa-gu10-ekf-connect-5w-wifi-rgbw",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "ekf",
            "model": "slwf-gu-53-rgbw",
            "name": "Умная лампа GU5.3 EKF Connect 4,5W WIFI RGBW",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/ekf-light-4.svg",
            "link": "https://ekfgroup.com/ru/catalog/products/umnaya-lampa-gu5-3-ekf-connect-5w-wifi-rgbw",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "ekf",
            "model": "slwf-e27-fil",
            "name": "Умная филаментная лампа EKF Connect E27 Wi-Fi",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/69839af4_ekf-light-2.svg",
            "link": "https://ekfgroup.com/ru/catalog/products/umnaya-filamentnaya-lampa-ekf-connect-e27-wi-fi",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "ekf",
            "model": "slwf-e27-st64-fil-rgbw",
            "name": "Умная филаментная лампа EKF Connect E27 ST64 RGB",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/ekf-light-6.svg",
            "link": "https://ekfgroup.com/ru/catalog/products/umnaya-filamentnaya-rgb-lampa-e27-st64-ekf",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "ekf",
            "model": "slwf-e27-rgbw",
            "name": "Умная лампа EKF Connect 8W WIFI RGBW E27",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/ekf-light-7.svg",
            "link": "https://ekfgroup.com/ru/catalog/products/umnaya-led-lampa-ekf-connect-8w-wifi-rgbw-e27",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "ekf",
            "model": "slwf-e14-rgbw",
            "name": "Умная лампа EKF Connect 5W WIFI RGBW E14",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/ekf-light-8.svg",
            "link": "https://ekfgroup.com/ru/catalog/products/umnaya-led-lampa-ekf-homeconnect-8w-wifi-rgbw-e14",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "ekf",
            "model": "slwf-gx53-cct",
            "name": "Умная лампа GX53 EKF Connect Wi-Fi",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/f5bd1a61_ekf-light-3.svg",
            "link": "https://ekfgroup.com/ru/catalog/products/umnaya-lampa-ekf-connect-gx53-wi-fi",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "ekf",
            "model": "slwf-gx53-rgbw",
            "name": "Умная лампа GX53 EKF Connect RGBW Wi-Fi",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/ekf-light-5.svg",
            "link": "https://ekfgroup.com/ru/catalog/products/umnaya-lampa-gx53-ekf-connect-rgbw-wi-fi",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "ekf",
            "model": "sclwf-230-cct",
            "name": "Умный потолочный светильник 230 мм 18 W EKF Connect",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/ekf-light-11.svg",
            "link": "https://ekfgroup.com/ru/catalog/products/umnyj-potolochnyj-svetilnik-230-mm-ekf-connect",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "ekf",
            "model": "sclwf-300-cct",
            "name": "Умный потолочный светильник 300 мм 24 W EKF Connect",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/ekf-light-9.svg",
            "link": "https://ekfgroup.com/ru/catalog/products/umnyj-potolochnyj-svetilnik-300-mm-ekf-connect",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "ekf",
            "model": "sclwf-400-cct",
            "name": "Умный потолочный светильник 400 мм 32 W EKF Connect",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/ekf-light-10.svg",
            "link": "https://ekfgroup.com/ru/catalog/products/umnyj-potolochnyj-svetilnik-400-mm-belyj-ekf",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "ekf",
            "model": "sclwf-600-cct",
            "name": "Умный потолочный светильник 600 мм 45 W EKF Connect",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/ekf-light-10.svg",
            "link": "https://ekfgroup.com/ru/catalog/products/umnyj-potolochnyj-svetilnik-600-mm-belyj-ekf",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "ekf",
            "model": "sclwf-480-cct",
            "name": "Умный потолочный светильник 480 мм 36 W EKF Connect",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/ekf-light-10.svg",
            "link": "https://ekfgroup.com/ru/catalog/products/umnyj-potolochnyj-svetilnik-480-mm-belyj-ekf",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "ekf",
            "model": "RCS-1-WF",
            "name": "Умная розетка EKF Сonnect Wi-Fi белая",
            "categories": ["socket"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/e14f48be_ekf-socket-1-2.svg",
            "link": "https://ekfgroup.com/ru/catalog/products/umnaya-rozetka-wi-fi-ekfhomesonnect-belaya",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "ekf",
            "model": "RCS-2-WF",
            "name": "Умная розетка EKF Сonnect PRO Wi-Fi чёрная",
            "categories": ["socket"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/ekf-socket-4.svg",
            "link": "https://ekfgroup.com/ru/catalog/products/umnaya-rozetka-wi-fi-pro-ekfhomesonnect-chernaya",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "ekf",
            "model": "RCS-ST16-WD-ZB",
            "name": "Умная розетка Стокгольм 1-местная 16А EKF Сonnect. Подключается через Zigbee-колонку Sber или хаб Sber",
            "categories": ["socket"],
            "protocols": ["zigbee"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/b7742e25_EKF-RCS-ST16-WD-ZB.png",
            "link": "https://ekfgroup.com/ru/catalog/products/umnaya-rozetka-stokgolm-1-mestnaya-16a-belyj-zigbee-ekf-sonnect",
            "direct": ["hub"],
            "cloud": []
        },
        {
            "manufacturer": "ekf",
            "model": "RCE-1-WF",
            "name": "Умный удлинитель c USB Wi-Fi EKF Connect",
            "categories": ["socket"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/ekf-socket-2.svg",
            "link": "https://ekfgroup.com/ru/catalog/products/umnyj-udlinitel-c-usb-wi-fi-ekfhomeconnect",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "ekf",
            "model": "RCE-2-WF",
            "name": "Умный удлинитель c USB Wi-Fi PRO EKF Connect",
            "categories": ["socket"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/ekf-socket-3.svg",
            "link": "https://ekfgroup.com/ru/catalog/products/umnyj-udlinitel-c-usb-wi-fi-pro-ekfhomeconnect",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "ekf",
            "model": "slswf-3-rgbw",
            "name": "Умная светодиодная лента EKF Connect RGBW 3m",
            "categories": ["led-strip"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/ae4ca6ad_ekf-led-strip-1.jpeg",
            "link": "https://ekfgroup.com/ru/catalog/products/umnaya-svetodiodnaya-lenta-ekf-homeconnect-18w-rgbw-3m",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "ekf",
            "model": "slswf-5-rgbw",
            "name": "Умная светодиодная лента EKF Connect RGBW 5m",
            "categories": ["led-strip"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/ae4ca6ad_ekf-led-strip-1.jpeg",
            "link": "https://ekfgroup.com/ru/catalog/products/umnaya-svetodiodnaya-lenta-ekf-connect-rgbw-5m",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "ekf",
            "model": "RCV-ST1-WS-ZB",
            "name": "Умный беспроводной выключатель Стокгольм 1-кл белый Zigbee EKF Сonnect. Подключается через хаб EKF, Zigbee-колонку Sber или хаб Sber",
            "categories": ["scenario-button"],
            "protocols": ["zigbee"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/ekf-relay-1z.svg",
            "link": "https://ekfgroup.com/ru/catalog/products/umnyj-vyklyuchatel-stokgolm-1-kl-besprovodnoj-belyj-zigbee-ekf-sonnect",
            "direct": ["hub"],
            "cloud": ["с2с-hub"]
        },
        {
            "manufacturer": "ekf",
            "model": "RCV-ST2-WS-ZB",
            "name": "Умный беспроводной выключатель Стокгольм 2-кл. белый Zigbee EKF Сonnect. Подключается через Zigbee-колонку Sber или хаб Sber",
            "categories": ["scenario-button"],
            "protocols": ["zigbee"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/ekf-relay-2z.svg",
            "link": "https://ekfgroup.com/ru/catalog/products/umnyj-vyklyuchatel-stokgolm-2-kl-besprovodnoj-belyj-zigbee-ekf-sonnect",
            "direct": ["hub"],
            "cloud": []
        },
        {
            "manufacturer": "ekf",
            "model": "ett-4",
            "name": "Умный терморегулятор для теплых полов Wi-Fi EKF Connect",
            "categories": ["hvac-underfloor-heating"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/ekf-underfloor-1.svg",
            "link": "https://ekfgroup.com/ru/catalog/products/umnyj-termostat-dlya-teply-polov-elektronnyj-16-a-230-v-wi-fi-ekfhomeconnect",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "ekf",
            "model": "szh-t",
            "name": "Умный хаб EKF Connect",
            "categories": ["hub"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/ekf-hub-1.svg",
            "link": "https://ekfgroup.com/ru/catalog/products/umnyj-ab-ekf-connect",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "elari",
            "model": "SWT-ZB/12",
            "name": "Умный выключатель ELARI SmartNRG Switch",
            "categories": ["relay"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/a554c41d_elari-relay-1.png",
            "link": "https://elari.net/catalog/smart-home/smartnrg-switch/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "elari",
            "model": "SWT-ZB/22",
            "name": "Умный выключатель ELARI SmartNRG Switch Double",
            "categories": ["relay"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/cab08c4f_elari-relay-2.png",
            "link": "https://elari.net/catalog/smart-home/smartnrg-switch-double/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "elari",
            "model": "NRG-LR2C-1",
            "name": "Умный выключатель-модуль ELARI SmartNRG Light Module",
            "categories": ["relay"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/15d491ba_elari-relay-3.png",
            "link": "https://elari.net/catalog/smart-home/smartnrg-light-module/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "elari",
            "model": "NRG-NR1C-1",
            "name": "Умный выключатель-модуль ELARI SmartNRG Module",
            "categories": ["relay"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/15d491ba_elari-relay-3.png",
            "link": "https://elari.net/catalog/smart-home/smartnrg-module/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "elari",
            "model": "LMS-14RGB",
            "name": "Умная лампа ELARI SmartLED Color E14",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/d44524f0_elari-light-1.png",
            "link": "https://elari.net/catalog/smart-home/smartled-color-e14/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "elari",
            "model": "LMS-27",
            "name": "Умная лампа ELARI SmartLED Color E27",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/e3c8d427_elari-light-2.png",
            "link": "https://elari.net/catalog/smart-home/smartled-color-e27/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "elari",
            "model": "LMS-27 / LMS-27RGB",
            "name": "Умная лампа ELARI SmartLED Color E27",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/e3c8d427_elari-light-2.png",
            "link": "https://elari.net/catalog/smart-home/smartled-color-e27/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "elari",
            "model": "LMS-01",
            "name": "Умная лампа ELARI SmartLED Filament E27",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/3b3027f0_elari-light-3.png",
            "link": "https://elari.net/catalog/smart-home/smartled-filament-e27/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "elari",
            "model": "LMS-10CCT",
            "name": "Умная лампа ELARI SmartLED Warm&Cold GU10",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/9fcd0ebb_elari-light-4.svg",
            "link": "https://elari.net/catalog/smart-home/smartled-warm-cold-gu10/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "elari",
            "model": "SMS-EU10USB",
            "name": "Умная розетка ELARI Smart Socket",
            "categories": ["socket"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/0965956b_elari-socket-2.svg",
            "link": "https://elari.net/catalog/smart-home/smartsocket/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "elari",
            "model": "NRG-DEU16PM",
            "name": "Умная двойная розетка ELARI Dual Smart Socket",
            "categories": ["socket"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/372ecb04_elari-socket-1.png",
            "link": "https://elari.net/catalog/smart-home/dual-smart-socket/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "electrolux",
            "model": "Smartline",
            "name": "Кондиционеры серии Electrolux Smartline. Для управления по Wi-Fi должны быть оснащены модулем Hommyn HDN/WFN-02-01",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/1dad8d1f_electrolux-smartline.svg",
            "link": "https://www.rusklimat.ru/search/?how=r&q=Electrolux+Smartline",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "electrolux",
            "model": "ECH/TUI3.1",
            "name": "Блок управления Electrolux Transformer Digital Inverter 3.1",
            "categories": ["hvac-heater"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/92fa6755_hommyn-hvac-heater-2.svg",
            "link": "https://www.rusklimat.ru/product-_transformer_digital_inverter_electrolux_ech_tui3_1/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "electrolux",
            "model": "ECH/TUI4",
            "name": "Блок управления конвектора Electrolux Transformer Digital Inverter 4",
            "categories": ["hvac-heater"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/31a12015_electrolux-hvac-heater-4.svg",
            "link": "https://www.rusklimat.ru/product-blok_upravleniya_konvektora_electrolux_transformer_digital_inverter_ech_tui4/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "elektrostandard",
            "model": "76007/00",
            "name": "Elektrostandard 76007/00 реле Умный дом черный",
            "categories": ["relay"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/a4050e3c_minimir-relay-1.svg",
            "link": "https://minimir.ru/catalog/smart-home/wi-fi-rele/umnoe-dvukhkanalnoe-rele-76007-00-a059326",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "elektrostandard",
            "model": "76010/00",
            "name": "Elektrostandard 76010/00 реле Умный дом с независимым контактом",
            "categories": ["relay"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/7febb77e_minimir-relay-2.svg",
            "link": "https://minimir.ru/catalog/smart-home/wi-fi-rele/rele-umnyy-dom-s-nezavisimym-kontaktom-76010-00-a067444",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "elektrostandard",
            "model": "76001/00",
            "name": "Elektrostandard 76001/00 реле Умный дом",
            "categories": ["relay"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/bf82a7e5_minimir-relay-3.svg",
            "link": "https://minimir.ru/catalog/smart-home/wi-fi-rele/rele-umnyy-dom-76001-0-a055189",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "elektrostandard",
            "model": "76004/00",
            "name": "Elektrostandard 76004/00 реле Умный дом",
            "categories": ["relay"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/ad16f97d_minimir-relay-4.svg",
            "link": "https://minimir.ru/catalog/smart-home/wi-fi-rele/umnoe-trekhkanalnoe-rele-76004-00-a056203",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "elektrostandard",
            "model": "76009/00",
            "name": "Elektrostandard 76009/00 реле Умный дом с мониторингом энергопотребления",
            "categories": ["relay"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/85635b46_minimir-relay-5.svg",
            "link": "hhttps://minimir.ru/catalog/smart-home/rele-umnyy-dom-s-monitoringom-energopotrebleniya-76009-00-a062688",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "elektrostandard",
            "model": "76005/00",
            "name": "Elektrostandard 76005/00 реле Умный дом",
            "categories": ["relay"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/53cc8ac7_minimir-relay-6.svg",
            "link": "https://minimir.ru/catalog/smart-home/wi-fi-rele/umnoe-chetyrekhkanalnoe-rele-76005-00-a059230",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "elektrostandard",
            "model": "76006/00",
            "name": "Elektrostandard 76006/00 реле Умный дом черный",
            "categories": ["relay"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/60304363_minimir-relay-7.svg",
            "link": "https://minimir.ru/catalog/smart-home/wi-fi-rele/umnoe-odnokanalnoe-rele-76006-00-a059324",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "elektrostandard",
            "model": "76008/00",
            "name": "Elektrostandard 76008/00 реле Умный дом для жалюзи и штор",
            "categories": ["window-blind"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/2878de4f_minimir-blind.svg",
            "link": "https://minimir.ru/catalog/smart-home/wi-fi-rele/umnoe-rele-dlya-zhalyuzi-i-shtor-76008-00-a060692",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "elektrostandard",
            "model": "95000/00",
            "name": "Elektrostandard 95000/00 Умный контроллер для светодиодных лент RGBWW 12-24V",
            "categories": ["controller"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/d5881042_minimir-control-1.svg",
            "link": "https://minimir.ru/catalog/smart-home/umnye-kontrollery-dlya-svetodiodnoy-lenty/umnyy-kontroller-dlya-svetodiodnykh-lent-rgbww-12-24-v-95000-00-a055252",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "elektrostandard",
            "model": "95001/00",
            "name": "Elektrostandard 95001/00 Умный контроллер для светодиодных лент RGBW 12-24V",
            "categories": ["controller"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/e1f47cdb_minimir-control-2.svg",
            "link": "https://minimir.ru/catalog/electronics/umnyy-kontroller-dlya-svetodiodnykh-lent-rgbw-12-24-v-95001-00-a055253",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "elektrostandard",
            "model": "95002/00",
            "name": "Elektrostandard 95002/00 Умный контроллер для светодиодных лент RGB 12-24V",
            "categories": ["controller"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/ba64d49d_minimir-control-3.svg",
            "link": "https://minimir.ru/catalog/street-lighting/svetodiodnaya-lenta/lenta-12-v/kontrollery-dlya-led-lenty-12-24v/umnyy-kontroller-dlya-svetodiodnykh-lent-rgb-12-24-v-95002-00-a055254",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "elektrostandard",
            "model": "95003/00",
            "name": "Elektrostandard 95003/00 Умный контроллер для светодиодных лент MIX 12-24V",
            "categories": ["controller"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/cf51abe3_minimir-control-4.svg",
            "link": "https://minimir.ru/catalog/electronics/kontrollery-dlya-svetodiodnoy-lenty/umnyy-kontroller-dlya-svetodiodnykh-lent-mix-12-24-v-95003-00-a055255",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "elektrostandard",
            "model": "95004/00",
            "name": "Elektrostandard 95004/00 Умный контроллер для светодиодных лент dimming 12-24V",
            "categories": ["controller"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/6b5dcc66_minimir-control-5.svg",
            "link": "https://minimir.ru/catalog/street-lighting/svetodiodnaya-lenta/lenta-12-v/kontrollery-dlya-led-lenty-12-24v/umnyy-kontroller-dlya-svetodiodnykh-lent-dimming-12-24-v-95004-00-a055256",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "elektrostandard",
            "model": "BLE1437",
            "name": "Elektrostandard BLE1437 Умная лампа свеча F C37 Е14 5W 3300К-6500К CCT+DIM",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/75181b64_minimir-light-1.svg",
            "link": "https://minimir.ru/catalog/smart-home/umnye-lampy/umnaya-filamentnaya-svetodiodnaya-lampa-ble1437-a055921",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "elektrostandard",
            "model": "BLE2754",
            "name": "Elektrostandard BLE2754 Умная лампа Classic F А60 Е27 6,5W 3300К-6500К CCT+DIM",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/d11dcfd8_minimir-light-2.svg",
            "link": "https://minimir.ru/catalog/smart-home/umnye-lampy/umnaya-filamentnaya-svetodiodnaya-lampa-ble2754-a055920",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "elektrostandard",
            "model": "BLE2755",
            "name": "Elektrostandard BLE2755 Умная лампа Classic LED А60 Е27 10W 3300К-6500К CCT+DIM",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/bebacceb_minimir-light-3.svg",
            "link": "https://minimir.ru/catalog/smart-home/umnye-lampy/umnaya-svetodiodnaya-lampa-ble2755-a055923",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "elektrostandard",
            "model": "BLG5316",
            "name": "Elektrostandard BLG5316 Умная лампа G5.3 LED 5W 3300К-6500К CCT+DIM",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/870a44c5_minimir-light-4.svg",
            "link": "https://minimir.ru/catalog/smart-home/umnye-lampy/umnaya-svetodiodnaya-lampa-blg5316-a055926",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "elektrostandard",
            "model": "BLGU1016",
            "name": "Elektrostandard BLGU1016 Умная лампа GU10 LED 5W 3300К-6500К CCT+DIM",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/51b926a8_minimir-light-5.svg",
            "link": "https://minimir.ru/catalog/smart-home/umnye-lampy/umnaya-svetodiodnaya-lampa-blgu1016-a055925",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "elektrostandard",
            "model": "BLGX5316",
            "name": "Elektrostandard BLGX5316 Умная лампа GX53 10W 3300К-6500К CCT+DIM",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/1ff40bca_minimir-light-6.svg",
            "link": "https://minimir.ru/catalog/smart-home/umnye-lampy/umnaya-svetodiodnaya-lampa-blgx5316-a061026",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "elektrostandard",
            "model": "BLE1438",
            "name": "Elektrostandard BLE1438 Умная лампа Свеча  LED C37 Е14 5W 3300К-6500К CCT+DIM",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/75303967_minimir-light-7.svg",
            "link": "https://minimir.ru/catalog/smart-home/umnye-lampy/umnaya-svetodiodnaya-lampa-ble1438-a055924",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "elektrostandard",
            "model": "85076/01",
            "name": "Elektrostandard 85076/01 Slim Magnetic Умный трековый светильник 10W 2700-6500K Dim L01 (чёрный)",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/b6df8f88_minimir-light-8.svg",
            "link": "https://minimir.ru/catalog/interernoe-osveshchenie/tracks/umnyy-trekovyy-svetilnik-10w-2700-6500k-dim-l01-slim-magnetic-85076-01-a063539",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "elektrostandard",
            "model": "85077/01",
            "name": "Elektrostandard 85077/01 Slim Magnetic Умный трековый светильник 20W 2700-6500K Dim L02 (чёрный)",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/3b433ce8_minimir-light-9.svg",
            "link": "https://minimir.ru/catalog/interernoe-osveshchenie/tracks/trekovaya-magnitnaya-sistema-slim-magnetic/umnyye-trekovyye-svetilniki-slim-magnetic/umnyy-trekovyy-svetilnik-20w-2700-6500k-dim-l02-slim-magnetic-85077-01-a063540",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "elektrostandard",
            "model": "85080/01",
            "name": "Elektrostandard 85080/01 Slim Magnetic Умный трековый светильник 30W 2700-6500K Dim L03 (чёрный)",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/93013fda_minimir-light-10.svg",
            "link": "https://minimir.ru/catalog/interernoe-osveshchenie/tracks/trekovaya-magnitnaya-sistema-slim-magnetic/umnyy-trekovyy-svetilnik-30w-2700-6500k-dim-l03-slim-magnetic-85080-01-a063541",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "elektrostandard",
            "model": "85056/01",
            "name": "Elektrostandard 85056/01 Slim Magnetic Умный трековый светильник 14W 2700-6500K Dim Dual (латунь)",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/24814fc2_minimir-light-11.svg",
            "link": "https://minimir.ru/catalog/interernoe-osveshchenie/tracks/trekovaya-magnitnaya-sistema-slim-magnetic/umnyy-trekovyy-svetilnik-14-vt-2700-6500k-dim-dual-slim-magnetic-85056-01-a063527",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "elektrostandard",
            "model": "85056/01",
            "name": "Elektrostandard 85056/01 Slim Magnetic Умный трековый светильник 14W 2700-6500K Dim Dual (чёрный)",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/940a4faa_minimir-light-12.svg",
            "link": "https://minimir.ru/catalog/interernoe-osveshchenie/tracks/trekovaya-magnitnaya-sistema-slim-magnetic/umnyye-trekovyye-svetilniki-slim-magnetic/umnyy-trekovyy-svetilnik-14w-2700-6500k-dim-dual-slim-magnetic-85056-01-a063528",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "elektrostandard",
            "model": "85070/01",
            "name": "Elektrostandard 85070/01 Slim Magnetic Умный трековый светильник 7W 2700-6500K Dim Cubo (латунь)",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/551a1fee_minimir-light-13.svg",
            "link": "https://minimir.ru/catalog/interernoe-osveshchenie/tracks/trekovaya-magnitnaya-sistema-slim-magnetic/umnyy-trekovyy-svetilnik-7w-2700-6500k-dim-cubo-slim-magnetic-85070-01-a063529",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "elektrostandard",
            "model": "85070/01",
            "name": "Elektrostandard 85070/01 Slim Magnetic Умный трековый светильник 7W 2700-6500K Dim Cubo (чёрный)",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/380b8112_minimir-light-14.svg",
            "link": "https://minimir.ru/catalog/interernoe-osveshchenie/tracks/umnyy-trekovyy-svetilnik-7w-2700-6500k-dim-cubo-slim-magnetic-85070-01-a063530",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "elektrostandard",
            "model": "85071/01",
            "name": "Elektrostandard 85071/01 Slim Magnetic Умный трековый светильник 5W 2700-6500K Dim Cantors (латунь)",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/0fb266ba_minimir-light-15.svg",
            "link": "https://minimir.ru/catalog/interernoe-osveshchenie/tracks/trekovaya-magnitnaya-sistema-slim-magnetic/umnyye-trekovyye-svetilniki-slim-magnetic/umnyy-trekovyy-svetilnik-5w-2700-6500k-dim-cantors-slim-magnetic-85071-01-a063531",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "elektrostandard",
            "model": "85071/01",
            "name": "Elektrostandard 85071/01 Slim Magnetic Умный трековый светильник 5W 2700-6500K Dim Cantors (чёрный)",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/2aa02024_minimir-light-16.svg",
            "link": "https://minimir.ru/catalog/interernoe-osveshchenie/tracks/trekovaya-magnitnaya-sistema-slim-magnetic/umnyy-trekovyy-svetilnik-5w-2700-6500k-dim-cantors-slim-magnetic-85071-01-a063532",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "elektrostandard",
            "model": "85072/01",
            "name": "Elektrostandard 85072/01 Slim Magnetic Умный трековый светильник 7W 2700-6500K Dim Amend (чёрный)",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/38066c39_minimir-light-17.svg",
            "link": "https://minimir.ru/catalog/interernoe-osveshchenie/tracks/umnyy-trekovyy-svetilnik-7w-2700-6500k-dim-amend-slim-magnetic-85072-01-a063533",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "elektrostandard",
            "model": "85073/01",
            "name": "Elektrostandard 85073/01 Slim Magnetic Умный трековый светильник 15W 2700-6500K Dim Amend (чёрный)",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/75ae1ed4_minimir-light-18.svg",
            "link": "https://minimir.ru/catalog/interernoe-osveshchenie/tracks/trekovaya-magnitnaya-sistema-slim-magnetic/umnyye-trekovyye-svetilniki-slim-magnetic/umnyy-trekovyy-svetilnik-15w-2700-6500k-dim-amend-ch-rnyy-slim-magnetic-85073-01-a063534",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "elektrostandard",
            "model": "85074/01",
            "name": "Elektrostandard 85074/01 Slim Magnetic Умный трековый светильник 6W 2700-6500K Dim R01 (чёрный)",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/5ac9a99a_minimir-light-19.jpg",
            "link": "https://minimir.ru/catalog/interernoe-osveshchenie/tracks/trekovaya-magnitnaya-sistema-slim-magnetic/trekovyye-svetilniki-slim-mag/umnyy-trekovyy-svetilnik-6w-2700-6500k-dim-r01-slim-magnetic-85074-01-a063537",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "elektrostandard",
            "model": "85075/01",
            "name": "Elektrostandard 85075/01 Slim Magnetic Умный трековый светильник 12W 2700-6500K Dim R02 (чёрный)",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/d5bbbba7_minimir-light-20.png",
            "link": "https://minimir.ru/catalog/interernoe-osveshchenie/tracks/trekovaya-magnitnaya-sistema-slim-magnetic/umnyye-trekovyye-svetilniki-slim-magnetic/umnyy-trekovyy-svetilnik-12w-2700-6500k-dim-r02-slim-magnetic-85075-01-a063538",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "elektrostandard",
            "model": "76100/00",
            "name": "Elektrostandard 76100/00 Умная розетка белый",
            "categories": ["socket"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/a4add6a8_minimir-socket-1.svg",
            "link": "https://minimir.ru/catalog/smart-home/umnye-rozetki/umnaya-rozetka-76100-00-a054335",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "elektrostandard",
            "model": "76100/00",
            "name": "Elektrostandard 76100/00 Умная розетка черный матовый",
            "categories": ["socket"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/6f889ac9_minimir-socket-2.svg",
            "link": "https://minimir.ru/catalog/smart-home/umnye-rozetki/umnaya-rozetka-76100-00-a054336",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "elektrostandard",
            "model": "76102/00",
            "name": "Elektrostandard 76102/00 Умная розетка белый",
            "categories": ["socket"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/b13e3b28_minimir-socket-3.svg",
            "link": "https://minimir.ru/catalog/smart-home/umnye-rozetki/umnaya-rozetka-76102-00-a060311",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "elektrostandard",
            "model": "76103/00",
            "name": "Elektrostandard 76103/00 Умная розетка с мониторингом энергопотребления черный IP44",
            "categories": ["socket"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/41002732_minimir-socket-4.svg",
            "link": "https://minimir.ru/catalog/smart-home/umnye-rozetki/umnaya-rozetka-ip44-76103-00-a065376",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "eltexhome",
            "model": "SW-RLY01",
            "name": "Умное реле Eltex Home для управления освещением SW-RLY01",
            "categories": ["relay"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/6416f422_eltexhome-relay-1.svg",
            "link": "https://eltex-doma.ru/catalog/smart-devices/sw-rly01/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "eltexhome",
            "model": "SZ-PIR",
            "name": "Датчик движения Eltex Home SZ-PIR REV.C. Подключается через хаб Eltex Home",
            "categories": ["sensor-pir"],
            "protocols": ["zwave"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/1d9bc7e1_eltexhome-sensor-pir-1.png",
            "link": "https://eltex-doma.ru/catalog/sensors/sz-pir/",
            "direct": [],
            "cloud": ["с2с-hub"]
        },
        {
            "manufacturer": "eltexhome",
            "model": "SZ-MCT",
            "name": "Датчик открытия Eltex Home SZ-MCT REV.C. Подключается через хаб Eltex Home",
            "categories": ["sensor-door"],
            "protocols": ["zwave"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/7781eacf_eltexhome-sensor-door-1.svg",
            "link": "https://eltex-doma.ru/catalog/sensors/sz-mct/",
            "direct": [],
            "cloud": ["с2с-hub"]
        },
        {
            "manufacturer": "eltexhome",
            "model": "SZ-WLK REV.B",
            "name": "Беспроводной датчик протечки воды Eltex Home SZ-WLK REV.B. Подключается через хаб Eltex Home",
            "categories": ["sensor-water-leak"],
            "protocols": ["zwave"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/c266c5dd_eltexhome-sensor-water-leak-1.png",
            "link": "https://eltex-doma.ru/catalog/sensors/sz-wlk/",
            "direct": [],
            "cloud": ["с2с-hub"]
        },
        {
            "manufacturer": "eltexhome",
            "model": "SW-PLG01",
            "name": "Eltex Home Wi-Fi-розетка",
            "categories": ["socket"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/b3fa65c3_eltexhome-socket-1.svg",
            "link": "https://eltex-doma.ru/support/knowledge/umnaya-wi-fi-rozetka-eltex-sw-plg01/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "funai",
            "model": "",
            "name": "Кондиционеры FUNAI. Для управления по Wi-Fi должны быть оснащены модулем Daichi DW21-B/DW22-B",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/funai-hvac-ac.svg",
            "link": "https://daichicloud.ru/split-lineup/search.php?search=FUNAI",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "geozon",
            "model": "GSH-SLF01",
            "name": "Умная лампочка Geozon FL-01",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/geozon-light-filamentnaya-fl-01.svg",
            "link": "https://geozon.ru/catalog/smarthome/umnaya-lampochka-filamentnaya-fl-01/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "geozon",
            "model": "GSH-SLF02",
            "name": "Умная лампочка Geozon FL-02",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/geozon-light-filamentnaya-fl-02.svg",
            "link": "https://geozon.ru/catalog/smarthome/umnaya-lampochka-filamentnaya-fl-02/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "geozon",
            "model": "GSH-SLF03",
            "name": "Умная лампочка Geozon FL-03",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/geozon-light-filamentnaya-fl-03.svg",
            "link": "https://geozon.ru/catalog/smarthome/umnaya-lampochka-filamentnaya-fl-03/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "geozon",
            "model": "GSH-SLF04",
            "name": "Умная лампочка Geozon FL-04",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/geozon-light-filamentnaya-fl-04.svg",
            "link": "https://geozon.ru/catalog/smarthome/umnaya-lampochka-filamentnaya-fl-04/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "geozon",
            "model": "GSH-SLF05",
            "name": "Умная лампочка Geozon FL-05",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/geozon-light-filamentnaya-fl-05.svg",
            "link": "https://geozon.ru/catalog/smarthome/umnaya-lampochka-filamentnaya-fl-05/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "geozon",
            "model": "GSH-SLR01",
            "name": "Умная лампочка Geozon RG-01",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/geozon-light-rgb-rg-01.svg",
            "link": "https://geozon.ru/catalog/smarthome/umnaya-lampochka-rgb-rg-01/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "geozon",
            "model": "GSH-SLR02",
            "name": "Умная лампочка Geozon RG-02",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/geozon-light-rgb-rg-02.svg",
            "link": "https://geozon.ru/catalog/smarthome/umnaya-lampochka-rgb-rg-02/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "geozon",
            "model": "GSH-SLR03",
            "name": "Умная лампочка Geozon RG-03",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/geozon-light-rgb-rg-03.svg",
            "link": "https://geozon.ru/catalog/smarthome/umnaya-lampochka-rgb-rg-03/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "geozon",
            "model": "GSH-SSW01",
            "name": "Встраиваемая розетка GEOZON WP-01",
            "categories": ["socket"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/geozon-socket-rozetka-wp-01.svg",
            "link": "https://geozon.ru/catalog/smarthome/umnaya-vstraivaemaya-rozetka-wp-01/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "geozon",
            "model": "GSH-SSW02",
            "name": "Встраиваемая розетка GEOZON WP-02",
            "categories": ["socket"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/geozon-socket-rozetka-wp-02.svg",
            "link": "https://geozon.ru/catalog/smarthome/umnaya-vstraivaemaya-rozetka-wp-02/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "geozon",
            "model": "GSH-SSP03",
            "name": "Умная розетка GEOZON PE-02",
            "categories": ["socket"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/geozon-socket-rozetka-pe-02.svg",
            "link": "https://geozon.ru/catalog/smarthome/umnaya-vneshnyaya-rozetka-pe-02/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "giulianovars",
            "model": "SmartKitchen",
            "name": "Контроллер Giulia Novars Smart Kitchen",
            "categories": ["controller"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/7aedd66e_giulianovars-controller-1.svg",
            "link": null,
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "gosund",
            "model": "EP8",
            "name": "Розетка Gosund Plug",
            "categories": ["socket"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/dc2fe50e_gosund-socket.svg",
            "link": "https://megamarket.ru/catalog/details/umnaya-rozetka-gosund-plug-ep8-600009618372_7/#?related_search=gosund%20ep8",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "gosund",
            "model": "SL1",
            "name": "Умная светодиодная лента Gosund Nitebird Smart LED Light Strip 2,8 м RGB Wi-Fi",
            "categories": ["led-strip"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/81620353_gosund-led.svg",
            "link": "https://megamarket.ru/catalog/details/umnaya-svetodiodnaya-lenta-gosund-smart-sl1-28m-white-600003325427_7/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "gosund",
            "model": "SL2",
            "name": "Умная светодиодная лента Gosund Nitebird Smart LED Light Strip 5 м RGB Wi-Fi",
            "categories": ["led-strip"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/81620353_gosund-led.svg",
            "link": "https://megamarket.ru/catalog/details/umnaya-svetodiodnaya-lenta-nitebird-sl2-5-m-100062269763_739/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "haier",
            "model": "",
            "name": "Кондиционеры Haier. Для управления по Wi-Fi должны быть оснащены модулем Daichi DW21-B/DW22-B",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/haier-hvac-ac.svg",
            "link": "https://daichicloud.ru/split-lineup/search.php?search=Haier",
            "direct": [],
            "cloud": ["с2с-app"],
            "exceptions": [
                {
                    "id": "brand_name",
                    "type": "string_value",
                    "value": "Daichi Comfort"
                },
                {
                    "id": "is_visible_app",
                    "type": "bool_value",
                    "value": true
                }
            ]
        },
        {
            "manufacturer": "haier",
            "model": "AABL65E01RU",
            "name": "Кондиционер Haier AS07TS6HRA-M",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/haier-hvac-ac-4.svg",
            "link": "https://haierproff.ru/catalog/cond/products/2970",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "haier",
            "model": "AABL64E01RU",
            "name": "Кондиционер Haier AS09TS6HRA-M",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/haier-hvac-ac-4.svg",
            "link": "https://haierproff.ru/catalog/cond/products/2971",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "haier",
            "model": "AABL8WE00RU",
            "name": "Кондиционер Haier AS12TS6HRA-M",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/haier-hvac-ac-4.svg",
            "link": "https://haierproff.ru/catalog/cond/products/2972",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "haier",
            "model": "AAANQRE00RU",
            "name": "Кондиционер Haier AS18TS5HRA-M",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/haier-hvac-ac-4.svg",
            "link": "https://haierproff.ru/catalog/cond/products/2973",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "haier",
            "model": "AAAZRJE00RU",
            "name": "Кондиционер Haier AS24TS5HRA-M",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/haier-hvac-ac-4.svg",
            "link": "https://haierproff.ru/catalog/cond/products/2974",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "haier",
            "model": "AAACMRE00RU",
            "name": "Кондиционер Haier AS70S2SF2FA-G",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/haier-hvac-ac-9.svg",
            "link": null,
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "haier",
            "model": "AAACMNE00RU",
            "name": "Кондиционер Haier AS70S2SF2FA-W",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/haier-hvac-ac-8.svg",
            "link": null,
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "haier",
            "model": "",
            "name": "Кондиционер Haier AS25S2SF3FA-S",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/haier-hvac-ac-11.svg",
            "link": null,
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "haier",
            "model": "",
            "name": "Кондиционер Haier AS35S2SF3FA-G",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/haier-hvac-ac-12.svg",
            "link": null,
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "haier",
            "model": "AABQGBE00RU",
            "name": "Кондиционер Haier AS70HPL1HRA",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/haier-hvac-ac-5.svg",
            "link": "https://haierproff.ru/catalog/cond/products/2949",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "haier",
            "model": "AABQG9E00RU",
            "name": "Кондиционер Haier AS70PHP2HRA",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/haier-hvac-ac-3.svg",
            "link": "https://haierproff.ru/catalog/cond/products/2944",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "haier",
            "model": "AAC0M1E00RU",
            "name": "Кондиционер Haier AS20HPL1HRA",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/haier-hvac-ac-5.svg",
            "link": "https://haieronline.ru/catalog/ventilation-and-heating/acs/konditsioner-haier-coral-dc-as20hpl1hra/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "haier",
            "model": "AAC0M6E00RU",
            "name": "Кондиционер Haier AS20HPL2HRA",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/haier-hvac-ac-3.svg",
            "link": "https://haieronline.ru/catalog/ventilation-and-heating/acs/konditsioner-haier-coral-dc-as20hpl2hra/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "haier",
            "model": "AAC0N1E00RU",
            "name": "Кондиционер Haier AS25HPL1HRA",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/haier-hvac-ac-5.svg",
            "link": "https://haieronline.ru/catalog/ventilation-and-heating/acs/konditsioner-haier-coral-dc-as25hpl1hra/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "haier",
            "model": "AAC0NBE00RU",
            "name": "Кондиционер Haier AS25HPL2HRA",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/haier-hvac-ac-3.svg",
            "link": "https://haieronline.ru/catalog/ventilation-and-heating/acs/konditsioner-haier-coral-dc-as25hpl2hra/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "haier",
            "model": "AAC1A1E00RU",
            "name": "Кондиционер Haier AS35HPL1HRA",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/haier-hvac-ac-5.svg",
            "link": "https://haieronline.ru/catalog/ventilation-and-heating/acs/konditsioner-haier-coral-dc-as35hpl1hra/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "haier",
            "model": "AAC0X8E00RU",
            "name": "Кондиционер Haier AS35HPL2HRA",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/haier-hvac-ac-3.svg",
            "link": "https://haieronline.ru/catalog/ventilation-and-heating/acs/konditsioner-haier-coral-dc-as35hpl2hra/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "haier",
            "model": "AAC1B2E00RU",
            "name": "Кондиционер Haier AS50HPL1HRA",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/haier-hvac-ac-5.svg",
            "link": "https://haieronline.ru/catalog/ventilation-and-heating/acs/konditsioner-haier-coral-dc-as50hpl1hra/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "haier",
            "model": "AAC1BCE00RU",
            "name": "Кондиционер Haier AS50HPL2HRA",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/haier-hvac-ac-5.svg",
            "link": "https://haieronline.ru/catalog/ventilation-and-heating/acs/konditsioner-haier-coral-dc-as50hpl2hra/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "haier",
            "model": "AABF1QE00RU",
            "name": "Кондиционер Haier AS20PHP2HRA",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/haier-hvac-ac-3.svg",
            "link": "https://haieronline.ru/catalog/ventilation-and-heating/acs/konditsioner-haier-coral-expert-20s-as20php2hra/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "haier",
            "model": "",
            "name": "Кондиционер Haier AS20PHP3HRA",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/haier-hvac-ac-3.svg",
            "link": "https://haieronline.ru/catalog/ventilation-and-heating/acs/konditsioner-haier-coral-expert-20s-as20php3hra/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "haier",
            "model": "AABF1RE00RU",
            "name": "Кондиционер Haier AS25PHP2HRA",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/haier-hvac-ac-3.svg",
            "link": "https://haieronline.ru/catalog/ventilation-and-heating/acs/konditsioner-haier-coral-expert-20s-as25php2hra/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "haier",
            "model": "",
            "name": "Кондиционер Haier AS25PHP3HRA",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/haier-hvac-ac-3.svg",
            "link": "https://haieronline.ru/catalog/ventilation-and-heating/acs/konditsioner-haier-coral-expert-20s-as25php3hra/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "haier",
            "model": "AABF3ME00RU",
            "name": "Кондиционер Haier AS35PHP2HRA",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/haier-hvac-ac-3.svg",
            "link": "https://haieronline.ru/catalog/ventilation-and-heating/acs/konditsioner-haier-coral-expert-20s-as35php2hra/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "haier",
            "model": "",
            "name": "Кондиционер Haier AS35PHP3HRA",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/haier-hvac-ac-3.svg",
            "link": "https://haieronline.ru/catalog/ventilation-and-heating/acs/konditsioner-haier-coral-expert-20s-as35php3hra/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "haier",
            "model": "AABF5HE00RU",
            "name": "Кондиционер Haier AS50PHP2HRA",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/haier-hvac-ac-3.svg",
            "link": "https://haieronline.ru/catalog/ventilation-and-heating/acs/konditsioner-haier-coral-expert-20s-as50php2hra/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "haier",
            "model": "",
            "name": "Кондиционер Haier AS50PHP3HRA",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/haier-hvac-ac-3.svg",
            "link": "https://haieronline.ru/catalog/ventilation-and-heating/acs/konditsioner-haier-coral-expert-20s-as50php3hra/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "haier",
            "model": "AABFAEE00RU",
            "name": "Кондиционер Haier HSU-07HPL203/R3(IN)",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/haier-hvac-ac-3.svg",
            "link": "https://haieronline.ru/catalog/ventilation-and-heating/acs/konditsioner-haier-coral-on-off-hsu-07hpl203-r3/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "haier",
            "model": "AABFGCE00RU",
            "name": "Кондиционер Haier HSU-09HPL203/R3(IN)",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/haier-hvac-ac-3.svg",
            "link": "https://haieronline.ru/catalog/ventilation-and-heating/acs/konditsioner-haier-coral-on-off-hsu-09hpl203-r3/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "haier",
            "model": "AABFH9E00RU",
            "name": "Кондиционер Haier HSU-12HPL203/R3(IN)",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/haier-hvac-ac-3.svg",
            "link": "https://haieronline.ru/catalog/ventilation-and-heating/acs/konditsioner-haier-coral-on-off-hsu-12hpl203-r3/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "haier",
            "model": "AABFH9E02RU",
            "name": "Кондиционер Haier HSU-18HPL203/R3(IN)",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/haier-hvac-ac-3.svg",
            "link": "https://haieronline.ru/catalog/ventilation-and-heating/acs/konditsioner-haier-coral-on-off-hsu-18hpl203-r3/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "haier",
            "model": "AAA86KE01RU",
            "name": "Кондиционер Haier AS25S2SF2FA-B",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/haier-hvac-ac-10.svg",
            "link": null,
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "haier",
            "model": "",
            "name": "Кондиционер Haier AS25S2SF2FA-B",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/haier-hvac-ac-10.svg",
            "link": "https://haieronline.ru/catalog/ventilation-and-heating/acs/konditsioner-haier-flexis-super-match-as25s2sf2fa-b/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "haier",
            "model": "AAA86LE01RU",
            "name": "Кондиционер Haier AS25S2SF2FA-G",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/haier-hvac-ac-9.svg",
            "link": null,
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "haier",
            "model": "AAA86JE01RU",
            "name": "Кондиционер Haier AS25S2SF2FA-W",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/haier-hvac-ac-6.svg",
            "link": null,
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "haier",
            "model": "",
            "name": "Кондиционер Haier AS25S2SF2FA-W",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/haier-hvac-ac-11.svg",
            "link": "https://haieronline.ru/catalog/ventilation-and-heating/acs/konditsioner-haier-flexis-super-match-as25s2sf2fa-w/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "haier",
            "model": "AAA6SAE01RU",
            "name": "Кондиционер Haier AS35S2SF2FA-B",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/haier-hvac-ac-10.svg",
            "link": null,
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "haier",
            "model": "",
            "name": "Кондиционер Haier AS35S2SF2FA-B",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/haier-hvac-ac-10.svg",
            "link": "https://haieronline.ru/catalog/ventilation-and-heating/acs/konditsioner-haier-flexis-super-match-as35s2sf2fa-b/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "haier",
            "model": "AAA6SCE01RU",
            "name": "Кондиционер Haier AS35S2SF2FA-G",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/haier-hvac-ac-12.svg",
            "link": null,
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "haier",
            "model": "AAA6SDE01RU",
            "name": "Кондиционер Haier AS35S2SF2FA-W",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/haier-hvac-ac-11.svg",
            "link": null,
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "haier",
            "model": "",
            "name": "Кондиционер Haier AS35S2SF2FA-W",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/haier-hvac-ac-11.svg",
            "link": "https://haieronline.ru/catalog/ventilation-and-heating/acs/konditsioner-haier-flexis-super-match-as35s2sf2fa-w/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "haier",
            "model": "AAACL0E01RU",
            "name": "Кондиционер Haier AS50S2SF2FA-B",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/haier-hvac-ac-10.svg",
            "link": null,
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "haier",
            "model": "AABFH9E68RU",
            "name": "Кондиционер Haier AS50S2SF2FA-B",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/haier-hvac-ac-10.svg",
            "link": "https://haieronline.ru/catalog/ventilation-and-heating/acs/konditsioner-haier-flexis-super-match-as50s2sf2fa-b/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "haier",
            "model": "AAACL2E01RU",
            "name": "Кондиционер Haier AS50S2SF2FA-G",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/haier-hvac-ac-12.svg",
            "link": null,
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "haier",
            "model": "AAACL1E01RU",
            "name": "Кондиционер Haier AS50S2SF2FA-W",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/haier-hvac-ac-11.svg",
            "link": null,
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "haier",
            "model": "AABFH9E66RU",
            "name": "Кондиционер Haier AS50S2SF2FA-W",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/haier-hvac-ac-6.svg",
            "link": "https://haieronline.ru/catalog/ventilation-and-heating/acs/konditsioner-haier-flexis-super-match-as50s2sf2fa-w/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "haier",
            "model": "",
            "name": "Кондиционер Haier AS07TT5HRA",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/haier-hvac-ac-4.svg",
            "link": "https://haieronline.ru/catalog/ventilation-and-heating/acs/konditsioner-haier-tundra-dc-as07tt5hra/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "haier",
            "model": "",
            "name": "Кондиционер Haier AS09TT5HRA",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/haier-hvac-ac-4.svg",
            "link": "https://haieronline.ru/catalog/ventilation-and-heating/acs/konditsioner-haier-tundra-dc-as09tt5hra/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "haier",
            "model": "",
            "name": "Кондиционер Haier AS12TT5HRA",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/haier-hvac-ac-4.svg",
            "link": "https://haieronline.ru/catalog/ventilation-and-heating/acs/konditsioner-haier-tundra-dc-as12tt5hra/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "haier",
            "model": "",
            "name": "Кондиционер Haier AS18TT5HRA",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/haier-hvac-ac-4.svg",
            "link": "https://haieronline.ru/catalog/ventilation-and-heating/acs/konditsioner-haier-tundra-dc-as18tt5hra/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "haier",
            "model": "",
            "name": "Кондиционер Haier AS24TT5HRA",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/haier-hvac-ac-3.svg",
            "link": "https://haierproff.ru/catalog/cond/products/3305",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "haier",
            "model": "",
            "name": "Кондиционер Haier HSU-07HTT03/R3(IN)",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/haier-hvac-ac-4.svg",
            "link": "https://haieronline.ru/catalog/ventilation-and-heating/acs/konditsioner-haier-tundra-on-off-hsu-07htt03-r3/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "haier",
            "model": "",
            "name": "Кондиционер Haier HSU-12HTT03/R3(IN)",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/haier-hvac-ac-4.svg",
            "link": "https://haieronline.ru/catalog/ventilation-and-heating/acs/konditsioner-haier-tundra-on-off-hsu-12htt03-r3/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "haier",
            "model": "",
            "name": "Кондиционер Haier HSU-18HTT03/R3(IN)",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/haier-hvac-ac-4.svg",
            "link": "https://haieronline.ru/catalog/ventilation-and-heating/acs/konditsioner-haier-tundra-on-off-hsu-18htt03-r3/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "haier",
            "model": "",
            "name": "Кондиционер Haier AS12CB2HRA",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/haier-hvac-ac-2.svg",
            "link": null,
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "haier",
            "model": "AABEWTE00RU",
            "name": "Кондиционер Haier AS25S2SJ2FA-G",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/haier-hvac-ac-14.svg",
            "link": "https://haieronline.ru/catalog/ventilation-and-heating/acs/konditsioner-split-sistema-haier-jade-super-match-as25s2sj2fa-g/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "haier",
            "model": "",
            "name": "Кондиционер Haier AS25S2SJ2FA-G",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/haier-hvac-ac-14.svg",
            "link": "https://haieronline.ru/catalog/ventilation-and-heating/acs/konditsioner-split-sistema-haier-jade-super-match-as25s2sj2fa-g/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "haier",
            "model": "AABEWRE00RU",
            "name": "Кондиционер Haier AS25S2SJ2FA-S",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/haier-hvac-ac-15.svg",
            "link": "https://haieronline.ru/catalog/ventilation-and-heating/acs/konditsioner-split-sistema-haier-jade-super-match-as25s2sj2fa-s/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "haier",
            "model": "",
            "name": "Кондиционер Haier AS25S2SJ2FA-S",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/haier-hvac-ac-15.svg",
            "link": "https://haieronline.ru/catalog/ventilation-and-heating/acs/konditsioner-split-sistema-haier-jade-super-match-as25s2sj2fa-s/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "haier",
            "model": "",
            "name": "Кондиционер Haier AS25S2SJ2FA-W",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/haier-hvac-ac-16.svg",
            "link": "https://haieronline.ru/catalog/ventilation-and-heating/acs/konditsioner-split-sistema-haier-jade-super-match-as25s2sj2fa-w/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "haier",
            "model": "AABEYCE00RU",
            "name": "Кондиционер Haier AS35S2SJ2FA-G",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/haier-hvac-ac-14.svg",
            "link": "https://haieronline.ru/catalog/ventilation-and-heating/acs/konditsioner-split-sistema-haier-jade-super-match-as35s2sj2fa-g/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "haier",
            "model": "",
            "name": "Кондиционер Haier AS35S2SJ2FA-G",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/haier-hvac-ac-14.svg",
            "link": "https://haieronline.ru/catalog/ventilation-and-heating/acs/konditsioner-split-sistema-haier-jade-super-match-as35s2sj2fa-g/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "haier",
            "model": "AABEYDE00RU",
            "name": "Кондиционер Haier AS35S2SJ2FA-S",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/haier-hvac-ac-15.svg",
            "link": "https://haieronline.ru/catalog/ventilation-and-heating/acs/konditsioner-split-sistema-haier-jade-super-match-as35s2sj2fa-s/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "haier",
            "model": "",
            "name": "Кондиционер Haier AS35S2SJ2FA-S",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/haier-hvac-ac-15.svg",
            "link": "https://haieronline.ru/catalog/ventilation-and-heating/acs/konditsioner-split-sistema-haier-jade-super-match-as35s2sj2fa-s/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "haier",
            "model": "AABEYBE00RU",
            "name": "Кондиционер Haier AS35S2SJ2FA-W",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/haier-hvac-ac-16.svg",
            "link": "https://haieronline.ru/catalog/ventilation-and-heating/acs/konditsioner-split-sistema-haier-jade-super-match-as35s2sj2fa-w/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "haier",
            "model": "",
            "name": "Кондиционер Haier AS35S2SJ2FA-W",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/haier-hvac-ac-16.svg",
            "link": "https://haieronline.ru/catalog/ventilation-and-heating/acs/konditsioner-split-sistema-haier-jade-super-match-as35s2sj2fa-w/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "haier",
            "model": "AABF0BE00RU",
            "name": "Кондиционер Haier AS50S2SJ2FA-G",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/haier-hvac-ac-14.svg",
            "link": "https://haieronline.ru/catalog/ventilation-and-heating/acs/konditsioner-split-sistema-haier-jade-super-match-as50s2sj2fa-g/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "haier",
            "model": "",
            "name": "Кондиционер Haier AS50S2SJ2FA-G",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/haier-hvac-ac-14.svg",
            "link": "https://haieronline.ru/catalog/ventilation-and-heating/acs/konditsioner-split-sistema-haier-jade-super-match-as50s2sj2fa-g/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "haier",
            "model": "AABF09E00RU",
            "name": "Кондиционер Haier AS50S2SJ2FA-S",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/haier-hvac-ac-17.svg",
            "link": "https://haieronline.ru/catalog/ventilation-and-heating/acs/konditsioner-split-sistema-haier-jade-super-match-as50s2sj2fa-s/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "haier",
            "model": "",
            "name": "Кондиционер Haier AS50S2SJ2FA-S",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/haier-hvac-ac-15.svg",
            "link": "https://haieronline.ru/catalog/ventilation-and-heating/acs/konditsioner-split-sistema-haier-jade-super-match-as50s2sj2fa-s/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "haier",
            "model": "AABF0AE00RU",
            "name": "Кондиционер Haier AS50S2SJ2FA-W",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/haier-hvac-ac-16.svg",
            "link": "https://haieronline.ru/catalog/ventilation-and-heating/acs/konditsioner-split-sistema-haier-jade-super-match-as50s2sj2fa-w/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "haier",
            "model": "",
            "name": "Кондиционер Haier AS50S2SJ2FA-W",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/haier-hvac-ac-13.svg",
            "link": "https://haieronline.ru/catalog/ventilation-and-heating/acs/konditsioner-split-sistema-haier-jade-super-match-as50s2sj2fa-w/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "haier",
            "model": "AAACMPE00RU",
            "name": "Кондиционер Haier HSU-24HFM203/R3(SDB)-IN",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/haier-hvac-ac-7.svg",
            "link": null,
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "haier",
            "model": "AAA86ME01RU",
            "name": "Кондиционер Haier HSU-09HFDN103/R3(SDB)-IN",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/haier-hvac-ac-8.svg",
            "link": null,
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "haier",
            "model": "AAA6SEE01RU",
            "name": "Кондиционер Haier HSU-12HFDN103/R3(SDB)-IN",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/haier-hvac-ac-8.svg",
            "link": null,
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "haier",
            "model": "AAACL3E01RU",
            "name": "Кондиционер Haier HSU-18HFDN103/R3(SDB)-IN",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/haier-hvac-ac-8.svg",
            "link": null,
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "haier",
            "model": "AAACMSE00RU",
            "name": "Кондиционер Haier HSU-24HFDN103/R3(SDB)-IN",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/haier-hvac-ac-8.svg",
            "link": null,
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "haier",
            "model": "AABQGXE00RU",
            "name": "Кондиционер Haier AS70HPL2HRA",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/haier-hvac-ac-3.svg",
            "link": null,
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "haier",
            "model": "AABFH9E04RU",
            "name": "Кондиционер Haier HSU-24HPL203/R3(IN)",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/haier-hvac-ac-3.svg",
            "link": "https://haierproff.ru/catalog/cond/products/3283",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "haier",
            "model": "",
            "name": "Кондиционер Haier HSU-07HFF103/R3-W(IN)",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/haier-hvac-ac-18.svg",
            "link": null,
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "haier",
            "model": "",
            "name": "Кондиционер Haier HSU-07HFF103/R3-G(IN)",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/haier-hvac-ac-19.svg",
            "link": null,
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "haier",
            "model": "",
            "name": "Кондиционер Haier HSU-07HFF103/R3-B(IN)",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/haier-hvac-ac-20.svg",
            "link": null,
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "haier",
            "model": "",
            "name": "Кондиционер Haier HSU-09HFF103/R3-W(IN)",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/haier-hvac-ac-18.svg",
            "link": null,
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "haier",
            "model": "",
            "name": "Кондиционер Haier HSU-09HFF103/R3-G(IN)",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/haier-hvac-ac-19.svg",
            "link": null,
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "haier",
            "model": "",
            "name": "Кондиционер Haier HSU-09HFF103/R3-B(IN)",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/haier-hvac-ac-20.svg",
            "link": null,
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "haier",
            "model": "",
            "name": "Кондиционер Haier HSU-12HFF103/R3-W(IN)",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/haier-hvac-ac-18.svg",
            "link": null,
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "haier",
            "model": "",
            "name": "Кондиционер Haier HSU-12HFF103/R3-G(IN)",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/haier-hvac-ac-19.svg",
            "link": null,
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "haier",
            "model": "",
            "name": "Кондиционер Haier HSU-12HFF103/R3-B(IN)",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/haier-hvac-ac-20.svg",
            "link": null,
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "haier",
            "model": "",
            "name": "Кондиционер Haier HSU-18HFF103/R3-W(IN)",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/haier-hvac-ac-18.svg",
            "link": null,
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "haier",
            "model": "",
            "name": "Кондиционер Haier HSU-18HFF103/R3-G(IN)",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/haier-hvac-ac-19.svg",
            "link": null,
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "haier",
            "model": "",
            "name": "Кондиционер Haier HSU-18HFF103/R3-B(IN)",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/haier-hvac-ac-20.svg",
            "link": null,
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "haier",
            "model": "",
            "name": "Кондиционер Haier HSU-24HFF103/R3-W(IN)",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/haier-hvac-ac-11.svg",
            "link": null,
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "haier",
            "model": "",
            "name": "Кондиционер Haier HSU-24HFF103/R3-G(IN)",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/haier-hvac-ac-12.svg",
            "link": null,
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "haier",
            "model": "",
            "name": "Кондиционер Haier HSU-24HFF103/R3-B(IN)",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/haier-hvac-ac-10.svg",
            "link": null,
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "haier",
            "model": "",
            "name": "Кондиционер Haier AS70PHP3HRA",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/haier-hvac-ac-3.svg",
            "link": "https://haierproff.ru/catalog/cond/products/3289",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "haier",
            "model": "",
            "name": "Кондиционер Haier AS25S2SF3FA-G",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/haier-hvac-ac-12.svg",
            "link": null,
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "haier",
            "model": "",
            "name": "Кондиционер Haier AS35S2SF3FA-S",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/haier-hvac-ac-2.svg",
            "link": null,
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "haier",
            "model": "AABFH9E67RU",
            "name": "Кондиционер Haier AS50S2SF3FA-G",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/haier-hvac-ac-19.svg",
            "link": null,
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "haier",
            "model": "",
            "name": "Кондиционер Haier AS50S2SF3FA-S",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/haier-hvac-ac-2.svg",
            "link": null,
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "haier",
            "model": "AABFH9E71RU",
            "name": "Кондиционер Haier AS70S2SF3FA-G",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/haier-hvac-ac-19.svg",
            "link": null,
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "haier",
            "model": "",
            "name": "Кондиционер Haier AS70S2SF3FA-S",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/haier-hvac-ac-2.svg",
            "link": null,
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "haier",
            "model": "",
            "name": "Кондиционер Haier HSU-09HTT103/R3(IN)",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/haier-hvac-ac-4.svg",
            "link": "https://haierproff.ru/catalog/cond/products/3297",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "haier",
            "model": "",
            "name": "Кондиционер Haier HSU-24HTT103/R3(IN)",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/haier-hvac-ac-4.svg",
            "link": "https://haierproff.ru/catalog/cond/products/3300",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "haier",
            "model": "AABF1RE01RU",
            "name": "Кондиционер Haier AS20PS1HRA-M",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/haier-hvac-ac-3.svg",
            "link": null,
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "haier",
            "model": "AABF1QE01RU",
            "name": "Кондиционер Haier AS25PS1HRA-M",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/haier-hvac-ac-3.svg",
            "link": null,
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "haier",
            "model": "AABF3JE01RU",
            "name": "Кондиционер Haier AS35PS1HRA-M",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/haier-hvac-ac-3.svg",
            "link": null,
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "haier",
            "model": "AABF58E01RU",
            "name": "Кондиционер Haier AS50PS1HRA-M",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/haier-hvac-ac-3.svg",
            "link": null,
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "haier",
            "model": "AABQGFE01RU",
            "name": "Кондиционер Haier AS70PS1HRA-M",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/haier-hvac-ac-3.svg",
            "link": null,
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "haier",
            "model": "AABFH9E127RU",
            "name": "Кондиционер Haier HSU-09HRM203/R3(IN)",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/haier-hvac-ac-21.svg",
            "link": "https://megamarket.ru/catalog/details/split-sistema-haier-hsu-09hrm203-r3-hsu-09hrm103-r3-100065446468/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "haier",
            "model": "AABFH9E129RU",
            "name": "Кондиционер Haier HSU-12HRM203/R3(IN)",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/haier-hvac-ac-22.svg",
            "link": null,
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "haier",
            "model": "AABFH9E137RU",
            "name": "Кондиционер Haier HSU-24HRM203/R3(DB)-IN",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/haier-hvac-ac-22.svg",
            "link": null,
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "haier",
            "model": "AABFH9E139RU",
            "name": "Кондиционер HHaier SU-09HFM303/R3(SDB)-IN",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/haier-hvac-ac-6.svg",
            "link": null,
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "haier",
            "model": "AABFH9E141RU",
            "name": "Кондиционер Haier HSU-12HFM303/R3(SDB)-IN",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/haier-hvac-ac-6.svg",
            "link": null,
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "haier",
            "model": "AABFH9E143RU",
            "name": "Кондиционер Haier HSU-18HFM303/R3(SDB)-IN",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/haier-hvac-ac-6.svg",
            "link": null,
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "haier",
            "model": "AABFH9E145RU",
            "name": "Кондиционер Haier HSU-24HFM303/R3(SDB)-IN",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/haier-hvac-ac-6.svg",
            "link": null,
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "haier",
            "model": "AACA32E00RU",
            "name": "Кондиционер Haier HSU-33HPL03/R3(IN)",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/haier-hvac-ac-3.svg",
            "link": "https://haierproff.ru/catalog/cond/products/3295",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "haier",
            "model": "AAAVD3E00RU",
            "name": "Кондиционер Haier AS100HPL1HRA",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/haier-hvac-ac-3.svg",
            "link": "https://haierproff.ru/catalog/cond/products/3284",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "haier",
            "model": "AAACMQE00RU",
            "name": "Кондиционер Haier AS70S2SF2FA-B",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/haier-hvac-ac-10.svg",
            "link": null,
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "haier",
            "model": "AABEWSE00RU",
            "name": "Кондиционер Haier AS25S2SJ2FA-W",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/haier-hvac-ac-6.svg",
            "link": "https://haierproff.ru/catalog/cond/products/2961",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "halsen",
            "model": "",
            "name": "Кондиционеры Halsen. Для управления по Wi-Fi должны быть оснащены модулем Daichi DW21-B/DW22-B",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/halsen-hvac-ac.svg",
            "link": "https://daichicloud.ru/split-lineup/search.php?search=HALSEN",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "hdl",
            "model": "HDL-MR1210.433",
            "name": "DIN реле",
            "categories": ["relay"],
            "protocols": ["own"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/hdl-1.png",
            "link": "https://hdlautomation.ru/catalog/hdl-buspro/din-rele-12-kanalnoe-10a-na-kanal/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "hdl",
            "model": "HDL-MR1210.434",
            "name": "DIN реле",
            "categories": ["relay"],
            "protocols": ["own"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/hdl-1.png",
            "link": "https://hdlautomation.ru/catalog/hdl-buspro/din-rele-12-kanalnoe-10a-na-kanal/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "hdl",
            "model": "HDL-MD0602.432",
            "name": "DIN диммер",
            "categories": ["relay"],
            "protocols": ["own"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/hdl-2.png",
            "link": "https://hdlautomation.ru/catalog/hdl-buspro/din-dimmer-6-kanalnyj-2a-na-kanal/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "hdl",
            "model": "HDL-MSP08M.4C",
            "name": "Сенсор 8 в одном",
            "categories": ["sensor-temp"],
            "protocols": ["own"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/hdl-7.png",
            "link": "https://hdlautomation.ru/catalog/hdl-buspro/sensor-8-v-odnom/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "hdl",
            "model": "CoolMaster",
            "name": "HDL Универсальный шлюз к системе кондиционирования",
            "categories": ["controller"],
            "protocols": ["own"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/hdl-3.png",
            "link": "https://hdlautomation.ru/catalog/coolautomation/universalnyj-shlyuz-k-sisteme-konditsionirovaniya/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "hdl",
            "model": "HDL-MAC01.431",
            "name": "HDL Модуль управления климатом, фанкойлами",
            "categories": ["controller"],
            "protocols": ["own"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/hdl-4.png",
            "link": "https://hdlautomation.ru/catalog/hdl-buspro/modul-upravleniya-klimatom-fankojlami/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "hdl",
            "model": "HDL-MWM65B.20",
            "name": "Привод штор HDL Buspro",
            "categories": ["curtain"],
            "protocols": ["own"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/hdl-5.png",
            "link": "https://hdlautomation.ru/catalog/hdl-buspro/privod-shtor-hdl-buspro/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "hdl",
            "model": "MW02.431",
            "name": "HDL DIN контроллер моторизованных штор, жалюзи, роллет на 2 мотора, 10А",
            "categories": ["curtain"],
            "protocols": ["own"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/hdl-6.png",
            "link": "https://hdlautomation.ru/catalog/hdl-buspro/din-kontroller-motorizovannykh-shtor-zhalyuzi-rollet-na-2-motora-10a/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "hdl",
            "model": "HDL-MFH06.432",
            "name": "Модуль управления отоплением, 6 каналов",
            "categories": ["hvac-underfloor-heating"],
            "protocols": ["own"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/hdl-8.png",
            "link": "https://hdlautomation.ru/catalog/hdl-buspro/modul-upravleniya-otopleniem-6-kanalov/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "hec",
            "model": "AAC0M4E00RU",
            "name": "Кондиционер HEC-07HRC03/R3(DB)-IN",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/hec-hvac-ac-1.svg",
            "link": "https://hec-rus.ru/catalog/cond/products/2967",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "hec",
            "model": "AABFA7E00RU",
            "name": "Кондиционер HEC-07HRC03/R3(IN)",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/hec-hvac-ac-1.svg",
            "link": "https://hec-rus.ru/catalog/cond/products/2962",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "hec",
            "model": "AAC0N7E00RU",
            "name": "Кондиционер HEC-09HRC03/R3(DB)-IN",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/hec-hvac-ac-1.svg",
            "link": "https://hec-rus.ru/catalog/cond/products/2968",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "hec",
            "model": "AABFG5E00RU",
            "name": "Кондиционер HEC-09HRC03/R3(IN)",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/hec-hvac-ac-1.svg",
            "link": "https://hec-rus.ru/catalog/cond/products/2963",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "hec",
            "model": "AAC1A7E00RU",
            "name": "Кондиционер HEC-12HRC03/R3(DB)-IN",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/hec-hvac-ac-1.svg",
            "link": "https://hec-rus.ru/catalog/cond/products/2969",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "hec",
            "model": "AABFH5E00RU",
            "name": "Кондиционер HEC-12HRC03/R3(IN)",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/hec-hvac-ac-1.svg",
            "link": "https://hec-rus.ru/catalog/cond/products/2964",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "hec",
            "model": "AAC1B7E00RU",
            "name": "Кондиционер HEC-18HRC03/R3(DB)-IN",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/hec-hvac-ac-1.svg",
            "link": "https://hec-rus.ru/catalog/cond/products/2970",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "hec",
            "model": "AABFJ4E00RU",
            "name": "Кондиционер HEC-18HRC03/R3(IN)",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/hec-hvac-ac-1.svg",
            "link": "https://hec-rus.ru/catalog/cond/products/2965",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "hec",
            "model": "AABQGKE00RU",
            "name": "Кондиционер HEC-24HRC03/R3(DB)-IN",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/hec-hvac-ac-1.svg",
            "link": "https://hec-rus.ru/catalog/cond/products/2971",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "hec",
            "model": "AABFK5E00RU",
            "name": "Кондиционер HEC-24HRC03/R3(IN)",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/hec-hvac-ac-1.svg",
            "link": "https://hec-rus.ru/catalog/cond/products/2966",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "hi",
            "model": "",
            "name": "Кондиционеры Hi. Для управления по Wi-Fi должны быть оснащены модулем Daichi DW21-B/DW22-B",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/hi-hvac-ac.svg",
            "link": "https://daichicloud.ru/split-lineup/device-page.php?line=621",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "hiper",
            "model": "HIPER IoT Switch B01",
            "name": "Умный встраиваемый кнопочный выключатель — 1 канал",
            "categories": ["relay"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/2d97f10b_hiper-relay-b01.svg",
            "link": "https://hiper-power.com/ru/catalog/umnyy-dom/vyklyuchateli/umnyy-vyklyuchatel-iot-switch-b01/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "hiper",
            "model": "HIPER IoT Switch B02",
            "name": "Умный встраиваемый кнопочный выключатель — 2 канала",
            "categories": ["relay"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/eef97f0c_hiper-relay-b02.jpeg",
            "link": "https://hiper-power.com/ru/catalog/umnyy-dom/vyklyuchateli/umnyy-vstraivaemyy-wi-fi-vyklyuchatel-hiper-iot-switch-b01/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "hiper",
            "model": "HIPER IoT Switch B03",
            "name": "Умный встраиваемый кнопочный выключатель — 3 канала",
            "categories": ["relay"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/78509d5f_hiper-relay-b03.jpeg",
            "link": "https://hiper-power.com/ru/catalog/umnyy-dom/umnyy-vstraivaemyy-vyklyuchatel-switch-b03/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "hiper",
            "model": "HIPER IoT Switch T01W",
            "name": "Умный встраиваемый сенсорный выключатель — 1 канал",
            "categories": ["relay"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/77ad7556_hiper-relay-t01w.jpeg",
            "link": "https://hiper-power.com/ru/catalog/umnyy-dom/umnyy-vstraivaemyy-wi-fi-vyklyuchatel-hiper-iot-switch-t01w/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "hiper",
            "model": "HIPER IoT Switch T02W",
            "name": "Умный встраиваемый сенсорный выключатель — 2 канала",
            "categories": ["relay"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/d1119898_hiper-relay-t02w.jpeg",
            "link": "https://hiper-power.com/ru/catalog/umnyy-dom/vyklyuchateli/umnyy-vstraivaemyy-wi-fi-vyklyuchatel-hiper-iot-switch-t02w/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "hiper",
            "model": "HIPER IoT Switch T03W",
            "name": "Умный встраиваемый сенсорный выключатель — 3 канала",
            "categories": ["relay"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/c715fe56_hiper-relay-t03w.jpeg",
            "link": "https://hiper-power.com/ru/catalog/umnyy-dom/vyklyuchateli/umnyy-vstraivaemyy-wi-fi-vyklyuchatel-hiper-iot-switch-t03w/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "hiper",
            "model": "HIPER IoT Switch M01",
            "name": "Умный модуль — 1 канал",
            "categories": ["relay"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/585faa3f_hiper-relay-m01.jpeg",
            "link": "https://hiper-power.com/ru/catalog/umnyy-dom/vyklyuchateli/umnyy-modul-vyklyuchatel-iot-switch-m01/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "hiper",
            "model": "HIPER IoT Switch M03",
            "name": "Умный модуль — 1 канал с управлением от IoT Switch Sх серии выключателей (беспроводные)",
            "categories": ["relay"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/74342a65_hiper-relay-m03.svg",
            "link": "https://hiper-power.com/ru/catalog/umnyy-dom/umnyy-modul-vyklyuchatel-hiper-iot-switch-m03/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "hiper",
            "model": "HIPER IoT Switch M04",
            "name": "Умный модуль — 1 канал с управлением от IoT Switch Sх серии выключателей (беспроводные)",
            "categories": ["relay"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/0d06acf9_hiper-relay-m04.svg",
            "link": "https://hiper-power.com/ru/catalog/umnyy-dom/vyklyuchateli/umnyy-modul-vyklyuchatel-hiper-iot-switch-m04/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "hiper",
            "model": "HIPER IoT Switch M02",
            "name": "Умный модуль — 2 канала, для установки в стакан обычного выключателя",
            "categories": ["relay"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/f7377903_hiper-relay-m02.jpeg",
            "link": "https://hiper-power.com/ru/catalog/umnyy-dom/vyklyuchateli/umnyy-modul-vyklyuchatel-iot-switch-m02/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "hiper",
            "model": "HIPER IoT Light DL772",
            "name": "Умная люстра, белый свет, регулировка температуры и яркости",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/5ae2dea2_hiper-light-1.svg",
            "link": "https://hiper-power.com/ru/catalog/umnyy-dom/osveshchenie/svetilniki/umnaya-potolochnaya-lampa-iot-light-dl772/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "hiper",
            "model": "HIPER IoT Light DL442",
            "name": "Умная люстра, белый свет, регулировка температуры и яркоcти",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/b5a07a7e_hiper-light-2.jpeg",
            "link": "https://hiper-power.com/ru/catalog/umnyy-dom/osveshchenie/svetilniki/umnaya-led-potolochnaya-lampa-hiper-iot-light-dl442/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "hiper",
            "model": "HIPER IoT DL221/331",
            "name": "Умная настольная лампа, белый свет, регулировка температуры и яркости",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/0c26176e_hiper-light-3.jpeg",
            "link": "https://hiper-power.com/ru/catalog/umnyy-dom/osveshchenie/svetilniki/umnyy-svetilnik-hiper-iot-dl221/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "hiper",
            "model": "HIPER IoT A60 Filament",
            "name": "Умная филаментная лампа, белый свет, регулировка температуры и яркости",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/04c1e5d3_hiper-light-4.jpeg",
            "link": "https://hiper-power.com/ru/catalog/umnyy-dom/osveshchenie/lampochki/umnaya-lampochka-iot-a60-filament/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "hiper",
            "model": "HIPER IoT A60 Filament Vintage",
            "name": "Умная филаментная лампа, белый свет, регулировка температуры и яркости",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/95dba10a_hiper-light-5.jpeg",
            "link": "https://hiper-power.com/ru/catalog/umnyy-dom/osveshchenie/lampochki/umnaya-lampochka-iot-a60-filament-vintage/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "hiper",
            "model": "HIPER IoT ST64 Filament Vintage",
            "name": "Умная филаментная лампа, белый свет, регулировка температуры и яркости",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/27371a50_hiper-light-6.jpeg",
            "link": "https://hiper-power.com/ru/catalog/umnyy-dom/osveshchenie/lampochki/umnaya-lampochka-iot-st64-filament-vintage/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "hiper",
            "model": "HIPER IoT G80 Filament Vintage",
            "name": "Умная филаментная лампа, белый свет, регулировка температуры и яркости",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/baa7a4d2_hiper-light-7.jpeg",
            "link": "https://hiper-power.com/ru/catalog/umnyy-dom/osveshchenie/lampochki/umnaya-lampochka-iot-g80-filament-vintage/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "hiper",
            "model": "HIPER IoT G95 Filament Vintage",
            "name": "Умная филаментная лампа, белый свет, регулировка температуры и яркости",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/47ab029a_hiper-light-8.jpeg",
            "link": "https://hiper-power.com/ru/catalog/umnyy-dom/osveshchenie/lampochki/umnaya-lampochka-iot-g95-filament-vintage/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "hiper",
            "model": "HIPER IoT LED A1 RGB",
            "name": "Умная LED-лампа, белый + RGB, регулировка температуры, яркоcти и RGB",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/9813413e_hiper-light-9.jpeg",
            "link": "https://hiper-power.com/ru/catalog/umnyy-dom/osveshchenie/lampochki/umnaya-lampochka-c-raznotsvetnoy-podsvetkoy-hiper-iot-led-a1-rgb/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "hiper",
            "model": "HIPER IoT A60 RGB",
            "name": "Умная LED-лампа, белый + RGB, регулировка температуры, яркоcти и RGB",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/ec7efdd4_hiper-light-10.jpeg",
            "link": "https://hiper-power.com/ru/catalog/umnyy-dom/osveshchenie/lampochki/umnaya-lampochka-c-raznotsvetnoy-podsvetkoy-hiper-iot-a60/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "hiper",
            "model": "HIPER IoT B1 RGB",
            "name": "Умная LED-лампа, белый + RGB, регулировка температуры, яркоcти и RGB",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/72f581aa_hiper-light-11.svg",
            "link": "https://hiper-power.com/ru/catalog/umnyy-dom/osveshchenie/lampochki/umnaya-led-lampochka-wi-fi-hiper-iot-b1-rgb/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "hiper",
            "model": "HIPER IoT C1 RGB",
            "name": "Умная LED-лампа, белый + RGB, регулировка температуры, яркоcти и RGB",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/bd320c23_hiper-light-12.svg",
            "link": "https://hiper-power.com/ru/catalog/umnyy-dom/osveshchenie/lampochki/umnaya-led-lampochka-wi-fi-hiper-iot-c1-rgb/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "hiper",
            "model": "HIPER IoT LED C2 RGB",
            "name": "Умная LED-лампа, белый + RGB, регулировка температуры, яркоcти и RGB",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/166999ff_hiper-light-13.svg",
            "link": "https://hiper-power.com/ru/catalog/umnyy-dom/osveshchenie/lampochki/umnaya-lampochka-c-raznotsvetnoy-podsvetkoy-hiper-iot-led-c2-rgb/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "hiper",
            "model": "HIPER IoT LED R1 RGB",
            "name": "Умная LED-лампа, белый + RGB, регулировка температуры, яркоcти и RGB",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/0400a6d5_hiper-light-14.jpeg",
            "link": "https://hiper-power.com/ru/catalog/umnyy-dom/osveshchenie/lampochki/umnaya-tsvetnaya-lampochka-r1-rgb/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "hiper",
            "model": "HIPER IoT LED R2 RGB",
            "name": "Умная LED-лампа, белый + RGB, регулировка температуры, яркоcти и RGB",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/0b5efb35_hiper-light-15.svg",
            "link": "https://hiper-power.com/ru/catalog/umnyy-dom/osveshchenie/lampochki/umnaya-lampochka-c-raznotsvetnoy-podsvetkoy-hiper-iot-led-r2-rgb/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "hiper",
            "model": "HIPER IoT Outlet W01",
            "name": "Умная встраиваемая розетка 3600 Вт",
            "categories": ["socket"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/361bc772_hiper-socker-1.svg",
            "link": "https://hiper-power.com/ru/catalog/umnyy-dom/rozetki/umnaya-vstraivaemaya-rozetka-iot-outlet-w01/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "hiper",
            "model": "HIPER IoT Outlet W02 Duo",
            "name": "Умная встраиваемая розетка 3600 Вт — 2 шт в одном корпусе",
            "categories": ["socket"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/0f23faf4_hiper-socker-2.jpeg",
            "link": "https://hiper-power.com/ru/catalog/umnyy-dom/rozetki/umnaya-rozetka-iot-outlet-w02-duo/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "hiper",
            "model": "HIPER IoT Panel E01",
            "name": "Умная накладная влагозащищённая розетка 3600 Вт",
            "categories": ["socket"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/67999efc_hiper-socker-5.svg",
            "link": "https://hiper-power.com/ru/catalog/umnyy-dom/rozetki/umnaya-nakladnaya-zashchishchennaya-ot-vody-ip55-rozetka-3-6kvt/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "hiper",
            "model": "HIPER IoT Panel E02",
            "name": "Умная двухканальная влагозащищённая розетка-переходник — 3600 Вт",
            "categories": ["socket"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/925cc58c_hiper-socker-4.jpeg",
            "link": "https://hiper-power.com/ru/catalog/umnyy-dom/rozetki/umnaya-zashchishchennaya-ot-vody-ip55-dvoynaya-rozetka-3-6kvt/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "hiper",
            "model": "HIPER IoT P01",
            "name": "Умная розетка переходник до 2500 Вт",
            "categories": ["socket"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/4a604b6c_hiper-socker-6.svg",
            "link": "https://hiper-power.com/ru/catalog/umnyy-dom/rozetki/umnaya-rozetka-hiper-iot-p01/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "hiper",
            "model": "HIPER IoT P02",
            "name": "Умная розетка переходник до 3600 Вт + мониторинг энергии",
            "categories": ["socket"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/21be79fe_hiper-socker-9.svg",
            "link": "https://hiper-power.com/ru/catalog/umnyy-dom/rozetki/umnaya-rozetka-hiper-iot-p02/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "hiper",
            "model": "HIPER IoT P03",
            "name": "Умная розетка переходник до 3600 Вт + мониторинг энергии",
            "categories": ["socket"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/19ca2142_hiper-socker-10.svg",
            "link": "https://hiper-power.com/ru/catalog/umnyy-dom/rozetki/umnaya-rozetka-wi-fi-hiper-iot-p03/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "hiper",
            "model": "HIPER IoT P04",
            "name": "Умная розетка переходник до 3600 Вт + мониторинг энергии",
            "categories": ["socket"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/7082ab4a_hiper-socker-11.png",
            "link": "https://hiper-power.com/ru/catalog/umnyy-dom/rozetki/umnaya-rozetka-wi-fi-hiper-iot-p04/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "hiper",
            "model": "HIPER IoT P07",
            "name": "Умная розетка переходник до 3600 Вт + мониторинг энергии",
            "categories": ["socket"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/ea481c79_hiper-socker-12.svg",
            "link": "https://hiper-power.com/ru/catalog/umnyy-dom/rozetki/umnaya-rozetka-p07-c-monitoringom-energii/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "hiper",
            "model": "HIPER IoT P08",
            "name": "Умная двойная розетка переходник до 3600 Вт, 2 независимых канала",
            "categories": ["socket"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/769989b1_hiper-socker-3.jpeg",
            "link": "https://hiper-power.com/ru/catalog/umnyy-dom/rozetki/umnaya-dvoynaya-rozetka-3600-vt/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "hiper",
            "model": "HIPER IoT P09",
            "name": "Умная розетка переходник до 2500 Вт + USB-порт",
            "categories": ["socket"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/9fb51bc8_hiper-socker-8.jpeg",
            "link": "https://hiper-power.com/ru/catalog/umnyy-dom/rozetki/umnaya-rozetka-s-usb-porotom/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "hiper",
            "model": "HIPER IoT PL01",
            "name": "Умная розетка переходник до 2500 Вт",
            "categories": ["socket"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/a2558f77_hiper-socker-7.jpeg",
            "link": "https://hiper-power.com/ru/catalog/umnyy-dom/rozetki/umnaya-rozetka-pl01/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "hiper",
            "model": "HIPER IoT PS44",
            "name": "Умный сетевой фильтр: 4 розетки + USB порты — 6 каналов",
            "categories": ["socket"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/4d2ca797_hiper-socker-13.jpeg",
            "link": "https://hiper-power.com/ru/catalog/umnyy-dom/rozetki/umnyy-setevoy-filtr-hiper-iot-ps44/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "hiper",
            "model": "HIPER IoT PS45",
            "name": "Умный сетевой фильтр: 4 розетки + USB порты — 6 каналов",
            "categories": ["socket"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/6b66f22f_hiper-socker-14.jpeg",
            "link": "https://hiper-power.com/ru/catalog/umnyy-dom/rozetki/umnyy-setevoy-filtr-hiper-iot-ps45/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "hiper",
            "model": "HIPER IoT Light DL115",
            "name": "Умная светодиодная лента, белый свет + RGB, регулировка температуры, яркоcти и RGB",
            "categories": ["led-strip"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/3c0bd76b_hiper-led-strip-1.jpeg",
            "link": "https://hiper-power.com/ru/catalog/umnyy-dom/osveshchenie/umnaya-svetodiodnaya-lenta-hiper-iot-light-dl115/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "hitepro",
            "model": "HiTE PRO Relay-4M",
            "name": "Блок радиореле HiTE PRO Relay-4M",
            "categories": ["relay"],
            "protocols": ["own"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/115cb661_hitepro-relay-4m.svg",
            "link": "https://www.hite-pro.ru/shop/goods/blok-upravleniya-relay-4m",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "hitepro",
            "model": "HP-Relay-DIM",
            "name": "Блок радиореле HiTE PRO Relay-DIM",
            "categories": ["relay"],
            "protocols": ["own"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/3f30be5c_hitepro-relay-dim.png",
            "link": "https://www.hite-pro.ru/shop/goods/blok-upravleniya-relay-dim",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "hitepro",
            "model": "HP-Relay-Drive",
            "name": "Блок радиореле HiTE PRO Relay-DRIVE",
            "categories": ["relay"],
            "protocols": ["own"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/ea9ce68e_hitepro-relay-drive.png",
            "link": "https://www.hite-pro.ru/shop/goods/blok-upravleniya-relay-drive-relay-drive-12v",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "hitepro",
            "model": "HP-Relay-LED",
            "name": "Блок радиореле HiTE PRO Relay-LED",
            "categories": ["relay"],
            "protocols": ["own"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/2e96cecf_hitepro-relay-led.png",
            "link": "https://www.hite-pro.ru/shop/goods/blok-upravleniya-relay-led",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "hitepro",
            "model": "Relay-LED3S",
            "name": "Блок радиореле HiTE PRO Relay-LED3S трехканальное",
            "categories": ["relay"],
            "protocols": ["own"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/57fd4ea5_hitepro-relay-led3s.svg",
            "link": "https://www.hite-pro.ru/shop/goods/blok-upravleniya-relay-led3s",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "hitepro",
            "model": "HP-Smart Motion",
            "name": "Датчик движения HiTE PRO Smart Motion",
            "categories": ["sensor-pir"],
            "protocols": ["own"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/10068beb_hitepro-sensor-pir-1-1.png",
            "link": "https://www.hite-pro.ru/shop/goods/datchik-dvizheniya-i-osveshhennosti-smart-motion",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "hitepro",
            "model": "HP-Smart Water",
            "name": "Датчик протечки HiTE PRO Smart Water",
            "categories": ["sensor-water-leak"],
            "protocols": ["own"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/941609c2_hitepro-sensor-water-leak-1-1.png",
            "link": "https://www.hite-pro.ru/shop/goods/datchik-protechki-smart-water",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "hitepro",
            "model": "HP-Smart Air",
            "name": "Датчик температуры и влажности HiTE PRO Smart Air",
            "categories": ["sensor-temp"],
            "protocols": ["own"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/5790bd83_hitepro-sensor-temp-1-1.png",
            "link": "https://www.hite-pro.ru/shop/goods/datchik-temperatury-i-vlazhnosti-smart-air",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "hitepro",
            "model": "HP-Gateway",
            "name": "HiTE PRO Gateway — сервер для управления умным домом",
            "categories": ["hub"],
            "protocols": ["own"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/c854247d_hitepro-hub-1.png",
            "link": "https://www.hite-pro.ru/shop/goods/server-umnogo-doma-gateway",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "hunberg",
            "model": "",
            "name": "Кондиционеры Hunberg. Для управления по Wi-Fi должны быть оснащены модулем Daichi DW21-B/DW22-B",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/hunberg-hvac-ac.svg",
            "link": "https://daichicloud.ru/split-lineup/search.php?search=Hunberg",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "ifeel",
            "model": "IFS-SB001",
            "name": "Умная лампочка iFEEL Globe E27 IFS-SB001",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/ifeel-light-1.svg",
            "link": "https://ifeel.systems/catalog/elektrika/umnaya_lampochka_ifeel_globe_e27_ifs_sb001/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "ifeel",
            "model": "IFS-SB002",
            "name": "Умная лампочка iFEEL Candle E14 IFS-SB002",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/ifeel-light-2.svg",
            "link": "https://ifeel.systems/catalog/elektrika/umnaya_lampochka_ifeel_candle_e14_ifs_sb002/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "ifeel",
            "model": "IFS-SB003",
            "name": "Умная лампочка iFEEL Globe E14 IFS-SB003",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/ifeel-light-3.svg",
            "link": "https://ifeel.systems/catalog/elektrika/umnaya_lampochka_ifeel_globe_e14_ifs_sb003/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "ifeel",
            "model": "IFS-SB004",
            "name": "Умная лампочка iFEEL Spot GU10 IFS-SB004",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/ifeel-light-4.svg",
            "link": "https://ifeel.systems/catalog/elektrika/umnaya_lampochka_ifeel_spot_gu10_ifs_sb004/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "ifeel",
            "model": "IFS-SP002",
            "name": "Умная розетка iFEEL Electra Plus IFS-SP002 с измерением энергопотребления",
            "categories": ["socket"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/ifeel-socket-1.svg",
            "link": "https://ifeel.systems/catalog/elektrika/umnaya_rozetka_ifeel_electra_plus_ifs_sp002_s_izmereniem_energopotrebleniya_i_alisoy/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "igc",
            "model": "",
            "name": "Кондиционеры IGC. Для управления по Wi-Fi должны быть оснащены модулем Daichi DW21-B/DW22-B",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/igc-hvac-ac.svg",
            "link": "https://daichicloud.ru/split-lineup/search.php?search=IGC",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "insmart",
            "model": "",
            "name": "Платформа inSmart. Управляет выключателями, реле, датчиками температуры и влажности, кранами с электроприводами и раздвижными шторами",
            "categories": ["other", "relay", "sensor-temp", "valve", "curtain"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/79f812e4_insmart-logo.svg",
            "link": "https://inpromauto.ru/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "irbis",
            "model": "IRB0121",
            "name": "Робот-пылесос IRBIS BEAN IRB0121",
            "categories": ["vacuum-cleaner"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/irbis-IRB0121.jpg",
            "link": "https://www.irbis.su/category/portable/roboty-pylesosy/product/637",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "irbis",
            "model": "IRB0321",
            "name": "Робот-пылесос IRBIS BEAN IRB0321",
            "categories": ["vacuum-cleaner"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/irbis-IRB0121.jpg",
            "link": "https://www.irbis.su/category/portable/roboty-pylesosy/product/637",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "irbis",
            "model": "IRB0421",
            "name": "Робот-пылесос IRBIS BEAN IRB0421",
            "categories": ["vacuum-cleaner"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/irbis-IRB0421.jpg",
            "link": "https://www.irbis.su/category/portable/roboty-pylesosy/product/918",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "irbis",
            "model": "IRP0121",
            "name": "Робот-пылесос IRBIS PEACH IRP0121",
            "categories": ["vacuum-cleaner"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/irbis-IRP0121.jpg",
            "link": "https://www.irbis.su/category/portable/roboty-pylesosy/product/919",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "irbis",
            "model": "IRP0421",
            "name": "Робот-пылесос IRBIS PEACH IRP0421",
            "categories": ["vacuum-cleaner"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/12fd7037_irbis-IRP0421.jpg",
            "link": "https://www.irbis.su/category/portable/roboty-pylesosy/product/920",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "irbis",
            "model": "IRHB10",
            "name": "Умная лампа IRBIS IRHB10",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/48766018_irbis-IRHB10_2.jpg",
            "link": "https://www.irbis-digital.ru/category/portable/umnyi-dom/product/604",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "iridi",
            "model": "",
            "name": "О совместимости устройств iRidi с умным домом Sber уточняйте у продавца или производителя",
            "categories": ["other"],
            "protocols": ["mqtt"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/04adc0ef_unknown.svg",
            "link": null,
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "iteq",
            "model": "IT-SM1ZI-K01",
            "name": "iTEQ SMART-выключатель 1-кл. 16А б/нейтр. ZigBee бел. ONI. Подключается через хаб EKF",
            "categories": ["relay"],
            "protocols": ["zigbee"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/fe36a931_iteq-relay-1-1.svg",
            "link": null,
            "direct": [],
            "cloud": ["с2с-hub"]
        },
        {
            "manufacturer": "iteq",
            "model": "IT-R16SFWBG-K01",
            "name": "iTEQ SMART-розетка 16А WIFI+BLE стекл/р. з/ш. бел. ONI",
            "categories": ["socket"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/481bc089_iteq-socker-1.svg",
            "link": null,
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "iteq",
            "model": "IT-R16SFWBG-K22",
            "name": "iTEQ SMART-розетка 16А WIFI+BLE стекл/р. з/ш. зол. ONI",
            "categories": ["socket"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/14bb0373_iteq-socker-2-1.svg",
            "link": null,
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "iteq",
            "model": "IT-R16SFWBG-K02",
            "name": "iTEQ SMART-розетка 16А WIFI+BLE стекл/р. з/ш. черн. ONI",
            "categories": ["socket"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/8a40d525_iteq-socker-3.svg",
            "link": null,
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "kentatsu",
            "model": "",
            "name": "Кондиционеры Kentatsu. Для управления по Wi-Fi должны быть оснащены модулем Daichi DW21-B/DW22-B",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/kentatsu-hvac-ac.svg",
            "link": "https://daichicloud.ru/split-lineup/search.php?search=Kentatsu",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "komanchi",
            "model": "",
            "name": "Кондиционеры Komanchi. Для управления по Wi-Fi должны быть оснащены модулем Daichi DW21-B/DW22-B",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/komanchi-hvac-ac.svg",
            "link": "https://daichicloud.ru/split-lineup/search.php?search=KOMANCHI",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "lessar",
            "model": "",
            "name": "Кондиционеры LESSAR. Для управления по Wi-Fi должны быть оснащены модулем Daichi DW21-B/DW22-B",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/lessar-hvac-ac.svg",
            "link": "https://daichicloud.ru/split-lineup/search.php?search=LESSAR",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "mdv",
            "model": "",
            "name": "Кондиционеры MDV. Для управления по Wi-Fi должны быть оснащены модулем Daichi DW21-B/DW22-B",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/mdv-hvac-ac.svg",
            "link": "https://daichicloud.ru/split-lineup/search.php?search=MDV",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "midea",
            "model": "",
            "name": "Кондиционеры Midea. Для управления по Wi-Fi должны быть оснащены модулем Daichi DW21-B/DW22-B",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/midea-hvac-ac.svg",
            "link": "https://daichicloud.ru/split-lineup/search.php?search=Midea",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "mitsubishi-electric",
            "model": "",
            "name": "Кондиционеры Mitsubishi Electric. Для управления по Wi-Fi должны быть оснащены модулем Daichi DW21-B/DW22-B",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/mitsubishi-electric-hvac-ac.svg",
            "link": "https://daichicloud.ru/split-lineup/search.php?search=MITSUBISHI+ELECTRIC",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "moes",
            "model": "WB-TDA9-RCW-E27",
            "name": "Умная лампочка MOES WiFi LED Bulb E27 (RGB+CW) 9W",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/f98dd57a_moes-light-1.svg",
            "link": "https://megamarket.ru/catalog/details/umnaya-svetodiodnaya-lampochka-moes-wb-tda9-rcw-e27-smart-led-bulb-e27-a60-600011602228/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "moes",
            "model": "WB-TDC6-RCW-E14",
            "name": "Умная лампочка MOES WiFi LED Bulb E14 (RGB+CW) 6W",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/d67e142e_moes-light-2.svg",
            "link": "https://megamarket.ru/catalog/details/umnaya-lampochka-moes-wifi-led-bulb-e14-6w-wb-tdc6-rcw-e14-600016373988/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "moes",
            "model": "WP-X-EU16M",
            "name": "Умная розетка MOES WiFi Plug 16A",
            "categories": ["socket"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/6d7df218_moes-socket-1.svg",
            "link": "https://megamarket.ru/catalog/details/umnaya-rozetka-moes-wifi-plug-16a-s-energomonitoringom-wp-x-eu16m-600016376448/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "moio",
            "model": "M32001",
            "name": "Контроллер MOiO 3ch для выключателей, датчиков и розеток",
            "categories": ["hvac-fan", "relay", "controller", "socket"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/f84d3c73_moio-img.svg",
            "link": "https://moio.pro/equipment/tproduct/795693832-859804111521-moio-3ch",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "moio",
            "model": "M32006",
            "name": "Контроллер MOiO Door + 2ch для дверного замка",
            "categories": ["gate", "controller"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/f84d3c73_moio-img.svg",
            "link": "https://moio.pro/equipment/tproduct/795693832-369450206441-moio-door-2ch",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "moio",
            "model": "M32003",
            "name": "Контроллер MOiO Dimmer AC для ламп освещения",
            "categories": ["controller", "light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/f84d3c73_moio-img.svg",
            "link": "https://moio.pro/equipment/tproduct/795693832-230988363271-moio-dimmer-ac",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "moio",
            "model": "M32004",
            "name": "Контроллер MOiO Dimmer DC для светодиодной ленты",
            "categories": ["controller", "light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/f84d3c73_moio-img.svg",
            "link": "https://moio.pro/equipment/tproduct/795693832-101094919571-moio-dimmer-dc",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "moio",
            "model": "M32007",
            "name": "Контроллер MOiO Valve для систем водоснабжения",
            "categories": ["controller", "valve"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/f84d3c73_moio-img.svg",
            "link": "https://moio.pro/equipment/tproduct/795693832-257844290111-moio-valve",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "moio",
            "model": "M32009",
            "name": "Контроллер MOiO Heater для систем отопления",
            "categories": ["controller", "hvac-boiler"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/f84d3c73_moio-img.svg",
            "link": "https://moio.pro/equipment/tproduct/795693832-286659988981-moio-heater",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "one-air",
            "model": "",
            "name": "Кондиционеры One Air. Для управления по Wi-Fi должны быть оснащены модулем Daichi DW21-B/DW22-B",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/one-air-hvac-ac.svg",
            "link": "https://daichicloud.ru/split-lineup/search.php?search=One+Air",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "philips",
            "model": "",
            "name": "О совместимости устройств Philips Hue с умным домом Sber уточняйте у продавца или производителя",
            "categories": ["other"],
            "protocols": ["other"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/04adc0ef_unknown.svg",
            "link": null,
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "polaris",
            "model": "PVCR 0726 WI-FI IQ Home GYRO",
            "name": "Polaris PVCR 0726 WI-FI IQ Home GYRO",
            "categories": ["vacuum-cleaner"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/9884aa54_polaris-vacuum-cleaner-1.png",
            "link": "https://polaris.ru/catalog/robot-pylesosy/polaris-pvcr-0726-wi-fi-iq-home-gyro-blue/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "polaris",
            "model": "PVCR 1028 WI-FI IQ Home",
            "name": "Polaris PVCR 1028 WI-FI IQ Home",
            "categories": ["vacuum-cleaner"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/d927a226_polaris-vacuum-cleaner-2.png",
            "link": "https://polaris.ru/catalog/robot-pylesosy/polaris-pvcr-1028-wi-fi-iq-home/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "polaris",
            "model": "PVCR 1050 WI-FI IQ Home Aqua",
            "name": "Polaris PVCR 1050 WI-FI IQ Home Aqua",
            "categories": ["vacuum-cleaner"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/4cb9e446_polaris-vacuum-cleaner-3.png",
            "link": "https://polaris.ru/catalog/robot-pylesosy/polaris-pvcr-1050/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "polaris",
            "model": "PVCR 1226 WI-FI IQ Home GYRO",
            "name": "Polaris PVCR 1226 WI-FI IQ Home GYRO",
            "categories": ["vacuum-cleaner"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/55db4aa9_polaris-vacuum-cleaner-4.png",
            "link": "https://polaris.ru/catalog/robot-pylesosy/polaris-pvcr-1226-wi-fi-iq-home-gray/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "polaris",
            "model": "PVCR 1229 WI-FI IQ Home Aqua",
            "name": "Polaris PVCR 1229 WI-FI IQ Home Aqua",
            "categories": ["vacuum-cleaner"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/335fba60_polaris-vacuum-cleaner-5.png",
            "link": "https://polaris.ru/catalog/robot-pylesosy/polaris-pvcr-1229-iq-home-aqua-white/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "polaris",
            "model": "PVCR 3100 IQ Home Aqua",
            "name": "Polaris PVCR 3100 IQ Home Aqua",
            "categories": ["vacuum-cleaner"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/26feeebc_polaris-vacuum-cleaner-6.png",
            "link": "https://polaris.ru/catalog/robot-pylesosy/polaris-pvcr-3100-iq-home-aqua/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "polaris",
            "model": "PVCR 3200 IQ Home Aqua",
            "name": "Polaris PVCR 3200 IQ Home Aqua",
            "categories": ["vacuum-cleaner"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/b5d76b3b_polaris-vacuum-cleaner-7.png",
            "link": "https://polaris.ru/catalog/robot-pylesosy/polaris-pvcr-3200-iq-home-aqua-darkgray/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "polaris",
            "model": "PVCR 3300 IQ Home Aqua",
            "name": "Polaris PVCR 3300 IQ Home Aqua",
            "categories": ["vacuum-cleaner"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/3445c822_polaris-vacuum-cleaner-8.png",
            "link": "https://polaris.ru/catalog/robot-pylesosy/robot-pylesos-polaris-pvcr-3300-iq-home-aqua/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "polaris",
            "model": "PVCR 3400 IQ Home Aqua",
            "name": "Polaris PVCR 3400 IQ Home Aqua",
            "categories": ["vacuum-cleaner"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/b534202f_polaris-vacuum-cleaner-9.png",
            "link": "https://polaris.ru/catalog/robot-pylesosy/polaris-pvcr-3400-iq-home-aqua/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "polaris",
            "model": "PVCR 3700 WI-FI IQ Home Aqua",
            "name": "Polaris PVCR 3700 WI-FI IQ Home Aqua",
            "categories": ["vacuum-cleaner"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/e94a9cb2_polaris-vacuum-cleaner-10.png",
            "link": "https://polaris.ru/catalog/robot-pylesosy/polaris-pvcr-3700-iq-home-wave-aqua/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "polaris",
            "model": "PVCR 3900 IQ Home Panorama Aqua",
            "name": "Polaris PVCR 3900 IQ Home Panorama Aqua",
            "categories": ["vacuum-cleaner"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/ca83c854_polaris-vacuum-cleaner-11.png",
            "link": "https://polaris.ru/catalog/robot-pylesosy/polaris-pvcr-3900-iq-home-panorama-aqua/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "polaris",
            "model": "PVCR 4000 Wi-Fi IQ Home Envision AQUA",
            "name": "Polaris PVCR 4000 Wi-Fi IQ Home Envision AQUA",
            "categories": ["vacuum-cleaner"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/1b251bdb_polaris-vacuum-cleaner-12.png",
            "link": "https://polaris.ru/catalog/robot-pylesosy/pvcr-4000-wi-fi-iq-home-envision-aqua/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "polaris",
            "model": "PVCR 4105 WI-FI IQ Home Aqua",
            "name": "Polaris PVCR 4105 WI-FI IQ Home Aqua",
            "categories": ["vacuum-cleaner"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/5b60e6b6_polaris-vacuum-cleaner-13.png",
            "link": "https://polaris.ru/catalog/robot-pylesosy/pvcr-4105-wi-fi-iq-home-aqua-white/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "polaris",
            "model": "PVCR Wave 15 WI-FI IQ Home Aqua",
            "name": "Polaris PVCR Wave 15 WI-FI IQ Home Aqua",
            "categories": ["vacuum-cleaner"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/d1e56a0a_polaris-vacuum-cleaner-14.png",
            "link": "https://polaris.ru/catalog/robot-pylesosy/polaris-pvcr-wave-15-wi-fi-iq-home-aqua-black/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "polaris",
            "model": "PVCR 0735 WI-FI IQ Home Aqua",
            "name": "Polaris PVCR 0735 WI-FI IQ Home Aqua",
            "categories": ["vacuum-cleaner"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/2806ac29_polaris-vacuum-cleaner-pvcr-0735.png",
            "link": "https://polaris.ru/catalog/robot-pylesosy/polaris-pvcr-0735/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "polaris",
            "model": "PVCR 0833 WI-FI IQ Home",
            "name": "Polaris PVCR 0833 WI-FI IQ Home",
            "categories": ["vacuum-cleaner"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/eb81d291_polaris-vacuum-cleaner-15.png",
            "link": "https://polaris.ru/catalog/robot-pylesosy/polaris-pvcr-0833/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "polaris",
            "model": "PVCR 5001 WIFI IQ Home",
            "name": "Polaris PVCR 5001 WIFI IQ Home",
            "categories": ["vacuum-cleaner"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/171de87e_polaris-vacuum-cleaner-16.png",
            "link": "https://polaris.ru/catalog/robot-pylesosy/polaris-pvcr-5001-iq-home-aqua/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "polaris",
            "model": "PVCRDC 5002 WIFI IQ Home",
            "name": "Polaris PVCRDC 5002 WIFI IQ Home",
            "categories": ["vacuum-cleaner"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/51f9e340_polaris-vacuum-cleaner-17.png",
            "link": "https://polaris.ru/catalog/robot-pylesosy/polaris-pvcrdc-5002-wi-fi-iq-home/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "polaris",
            "model": "PVCR 6001 WIFI IQ Home",
            "name": "Polaris PVCR 6001 WIFI IQ Home",
            "categories": ["vacuum-cleaner"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/5cf96e06_polaris-vacuum-cleaner-18.png",
            "link": "https://polaris.ru/catalog/robot-pylesosy/polaris-pvcrdc-6001-wifi-iq-home/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "polaris",
            "model": "PVCRDC 6002 WIFI IQ Home",
            "name": "Polaris PVCRDC 6002 WIFI IQ Home",
            "categories": ["vacuum-cleaner"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/c06a78cb_polaris-vacuum-cleaner-19.png",
            "link": "https://polaris.ru/catalog/robot-pylesosy/polaris-pvcrdc-6002-wifi-iq-home/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "polaris",
            "model": "PVCR G2 0726W WIFI IQ Home",
            "name": "Polaris PVCR G2 0726W WIFI IQ Home",
            "categories": ["vacuum-cleaner"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/fccc3d2a_polaris-vacuum-cleaner-20.png",
            "link": "https://polaris.ru/catalog/robot-pylesosy/polaris-pvcr-g2-0726w-wi-fi-iq-home/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "polaris",
            "model": "PVCR G2 0926W WIFI IQ Home",
            "name": "Polaris PVCR G2 0926W WIFI IQ Home",
            "categories": ["vacuum-cleaner"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/fccc3d2a_polaris-vacuum-cleaner-21.png",
            "link": "https://polaris.ru/catalog/robot-pylesosy/polaris-pvcr-g2-0926-wi-fi-iq-home-white/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "polaris",
            "model": "PVCR G2 1226 WIFI IQ Home",
            "name": "Polaris PVCR G2 1226 WIFI IQ Home",
            "categories": ["vacuum-cleaner"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/32bc4963_polaris-vacuum-cleaner-22.png",
            "link": "https://polaris.ru/catalog/robot-pylesosy/polaris-pvcr-g2-1226-wi-fi-iq-home-graphit/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "polaris",
            "model": "PVCR 3600 WIFI IQ Home",
            "name": "Polaris PVCR 3600 WIFI IQ Home",
            "categories": ["vacuum-cleaner"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/86d77e9b_polaris-vacuum-cleaner-23.png",
            "link": "https://polaris.ru/catalog/robot-pylesosy/polaris-pvcr-3600-wi-fi-iq-home/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "polaris",
            "model": "PVCR 0905 WIFI IQ Home Panorama Aqua",
            "name": "Polaris PVCR 0905 WIFI IQ Home Panorama Aqua",
            "categories": ["vacuum-cleaner"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/ce37accf_polaris-vacuum-cleaner-24.png",
            "link": "https://polaris.ru/catalog/robot-pylesosy/polaris-pvcr-0905-wifi-iq-home-panorama-aqua/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "polaris",
            "model": "PVCRDC 0101 WIFI IQ Home Panorama Extra",
            "name": "Polaris PVCRDC 0101 WIFI IQ Home Panorama Extra",
            "categories": ["vacuum-cleaner"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/7d499308_polaris-vacuum-cleaner-25.png",
            "link": "https://polaris.ru/catalog/robot-pylesosy/polaris-pvrdc-0101-wifi-iq-home-panorama-extra/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "polaris",
            "model": "PUH 9105 IQ Home",
            "name": "Polaris Ультразвуковой увлажнитель воздуха PUH 9105 IQ Home",
            "categories": ["hvac-humidifier"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/937d964f_polaris-hvac-humidifier-1.png",
            "link": "https://polaris.ru/catalog/humidifiers/uvlazhnitel-vozdukha-polaris-puh-9105-iq-home/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "polaris",
            "model": "PUH 1010 IQ Home",
            "name": "Polaris Увлажнитель воздуха PUH 1010 IQ Home",
            "categories": ["hvac-humidifier"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/1c57eaa8_polaris-hvac-humidifier-2.png",
            "link": "hhttps://polaris.ru/catalog/humidifiers/polaris-puh-1010-wifi-iq-home/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "polaris",
            "model": "PUH 2300 IQ Home",
            "name": "Polaris Увлажнитель воздуха PUH 2300 IQ Home",
            "categories": ["hvac-humidifier"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/31fe3272_polaris-hvac-humidifier-3.png",
            "link": "https://polaris.ru/catalog/humidifiers/polaris-puh-2300-wifi-iq-home/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "polaris",
            "model": "PUH 4040 IQ Home",
            "name": "Polaris Ультразвуковой увлажнитель воздуха PUH 4040 IQ Home",
            "categories": ["hvac-humidifier"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/e0e2e020_polaris-hvac-humidifier-puh-4040.png",
            "link": "https://polaris.ru/catalog/humidifiers/polaris-puh-4040-wifi-iq-home/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "polaris",
            "model": "PUH 9009 WIFI IQ Home",
            "name": "Polaris Увлажнитель воздуха PUH 9009 WIFI IQ Home",
            "categories": ["hvac-humidifier"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/2aaa7f35_polaris-hvac-humidifier-4.png",
            "link": "https://polaris.ru/catalog/humidifiers/polaris-puh-9009-wifi-iq-home/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "polaris",
            "model": "PUH 8080 WIFI IQ Home",
            "name": "Polaris Увлажнитель воздуха PUH 8080 WIFI IQ Home",
            "categories": ["hvac-humidifier"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/ca07e6a6_polaris-hvac-humidifier-5.png",
            "link": "https://polaris.ru/catalog/humidifiers/polaris-puh-8080-wifi-iq-home/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "polaris",
            "model": "PWK 1725CGLD IQ Home",
            "name": "Чайник Polaris PWK 1725CGLD IQ Home",
            "categories": ["kettle"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/6bd96064_polaris-kettle-1.png",
            "link": "https://polaris.ru/catalog/chayniki/chaynik-polaris-pwk-1725cgld-wifi-iq-home/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "polaris",
            "model": "PWK 1775CGLD IQ Home",
            "name": "Чайник Polaris PWK 1775CGLD IQ Home",
            "categories": ["kettle"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/polaris-kettle-2.svg",
            "link": "https://polaris.ru/catalog/chayniki/chaynik-polaris-pwk-1775cgld-wifi-iq-home/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "polaris",
            "model": "PWK 1755CAD IQ Home",
            "name": "Чайник Polaris PWK 1755CAD IQ Home",
            "categories": ["kettle"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/polaris-kettle-3.svg",
            "link": "https://polaris.ru/catalog/chayniki/chaynik-polaris-pwk-1755cad-wifi-iq-home-grey/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "polaris",
            "model": "PWK 1712CGLD IQ Home",
            "name": "Чайник Polaris PWK 1712CGLD IQ Home",
            "categories": ["kettle"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/polaris-kettle-4.svg",
            "link": "https://polaris.ru/catalog/chayniki/polaris-pwk-1712cgld-wifi-iq-home/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "polaris",
            "model": "PWK 1720CGLD IQ Home",
            "name": "Чайник Polaris PWK 1720CGLD IQ Home",
            "categories": ["kettle"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/polaris-kettle-5.svg",
            "link": "https://polaris.ru/catalog/chayniki/polaris-pwk-1720cgld-wifi-iq-home-blue/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "primera",
            "model": "",
            "name": "Кондиционеры Primera. Для управления по Wi-Fi должны быть оснащены модулем Daichi DW21-B/DW22-B",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/primera-hvac-ac.svg",
            "link": "https://daichicloud.ru/split-lineup/search.php?search=Primera",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "red",
            "model": "RED Solution RV-R6030S",
            "name": "Умный робот-пылесос RED Solution RV-R6030S",
            "categories": ["vacuum-cleaner"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/vacuum-cleaner-red-rv-r6030s.svg",
            "link": "https://redsolution.company/products/house/pylesosy/roboty_pylesosy/robot_pylesos_rv_r6030s_wi_fi/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "red",
            "model": "RED Solution RV-R6040S",
            "name": "Умный робот-пылесос RED Solution RV-R6040S",
            "categories": ["vacuum-cleaner"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/vacuum-cleaner-red-rv-r6040s.svg",
            "link": "https://redsolution.company/products/house/pylesosy/roboty_pylesosy/robot_pylesos_rv_r6040s_wi_fi/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "red",
            "model": "RED Solution RV-R6050S",
            "name": "Умный робот-пылесос RED Solution RV-R6050S",
            "categories": ["vacuum-cleaner"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/vacuum-cleaner-red-rv-r6050s.svg",
            "link": "https://redsolution.company/products/house/pylesosy/roboty_pylesosy/robot_pylesos_rv_r6050s_wi_fi/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "red",
            "model": "RED Solution RV-RL6000S",
            "name": "Умный робот-пылесос RED Solution RV-RL6000S",
            "categories": ["vacuum-cleaner"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/cfd61fa9_vacuum-cleaner-red-rv-rl6000s.png",
            "link": "https://redsolution.company/products/house/pylesosy/roboty_pylesosy/robot_pylesos_rv_rl6000s_wi_fi/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "red",
            "model": "RED Solution RV-R6060S",
            "name": "Умный робот-пылесос RED Solution RV-R6060S",
            "categories": ["vacuum-cleaner"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/35ecbad2_vacuum-cleaner-red-rv-rl6060s.png",
            "link": "https://redsolution.company/products/house/pylesosy/roboty_pylesosy/rv-r6060s-wifi/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "red",
            "model": "RED Solution RV-R6070S",
            "name": "Умный робот-пылесос RED Solution RV-R6070S",
            "categories": ["vacuum-cleaner"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/b476e1e5_vacuum-cleaner-red-rv-rl6070s.png",
            "link": "https://redsolution.company/products/house/pylesosy/roboty_pylesosy/robot_pylesos_rv_r6070s_wi_fi/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "red",
            "model": "RED Solution RV-RL6100S",
            "name": "Умный робот-пылесос RED Solution RV-RL6100S",
            "categories": ["vacuum-cleaner"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/077609f2_vacuum-cleaner-red-rv-rl6100s.png",
            "link": "https://redsolution.company/products/house/pylesosy/roboty_pylesosy/robot_pylesos_rv_rl6100s_wi_fi/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "red",
            "model": "RED solution SkyKettle G200S",
            "name": "Умный чайник. Подключается через хаб RED SkyCenter 12S",
            "categories": ["kettle"],
            "protocols": ["bluetooth"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/red-kettle-1.svg",
            "link": "https://redsolution.company/products/kitchen/chayniki/chaynik-rk-g200s/",
            "direct": [],
            "cloud": ["с2с-hub"]
        },
        {
            "manufacturer": "red",
            "model": "RED solution SkyKettle RK-G203S",
            "name": "Умный чайник. Подключается через хаб RED SkyCenter 12S",
            "categories": ["kettle"],
            "protocols": ["bluetooth"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/red-kettle-2.svg",
            "link": "https://redsolution.company/products/kitchen/chayniki/chaynik-rk-g203s/",
            "direct": [],
            "cloud": ["с2с-hub"]
        },
        {
            "manufacturer": "red",
            "model": "RED solution SkyKettle RK-G210S",
            "name": "Умный чайник. Подключается через хаб RED SkyCenter 12S",
            "categories": ["kettle"],
            "protocols": ["bluetooth"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/red-kettle-3.svg",
            "link": "https://redsolution.company/products/kitchen/chayniki/chaynik_red_solution_rk_g210s/",
            "direct": [],
            "cloud": ["с2с-hub"]
        },
        {
            "manufacturer": "red",
            "model": "RED solution SkyKettle RK-G212S",
            "name": "Умный чайник. Подключается через хаб RED SkyCenter 12S",
            "categories": ["kettle"],
            "protocols": ["bluetooth"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/red-kettle-4.svg",
            "link": "https://redsolution.company/products/kitchen/chayniki/umnyy_chaynik_red_solution_rk_g212s/",
            "direct": [],
            "cloud": ["с2с-hub"]
        },
        {
            "manufacturer": "red",
            "model": "RED solution SkyKettle RK-M216S",
            "name": "Умный чайник. Подключается через хаб RED SkyCenter 12S",
            "categories": ["kettle"],
            "protocols": ["bluetooth"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/red-kettle-5.svg",
            "link": "https://redsolution.company/products/kitchen/chayniki/chaynik-rk-m216s/",
            "direct": [],
            "cloud": ["с2с-hub"]
        },
        {
            "manufacturer": "roximo",
            "model": "SZBTN01-2",
            "name": "Умный Zigbee выключатель ROXIMO, двухкнопочный, белый. Подключается через Zigbee-колонку Sber или хаб Sber",
            "categories": ["relay"],
            "protocols": ["zigbee"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/roximo-szbtn01-2.png",
            "link": "https://iot.roximo.ru/smarthome/zigbee-button-switch-2",
            "direct": ["hub"],
            "cloud": []
        },
        {
            "manufacturer": "roximo",
            "model": "SZBTN01-2W",
            "name": "Умный Zigbee выключатель ROXIMO, двухкнопочный, белый. Подключается через Zigbee-колонку Sber или хаб Sber",
            "categories": ["relay"],
            "protocols": ["zigbee"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/roximo-szbtn01-2w.png",
            "link": "https://iot.roximo.ru/smarthome/zigbee-button-switch-szbtn01-2w",
            "direct": ["hub"],
            "cloud": []
        },
        {
            "manufacturer": "roximo",
            "model": "SZBTN01-2S",
            "name": "Умный Zigbee выключатель ROXIMO, двухкнопочный, серый. Подключается через Zigbee-колонку Sber или хаб Sber",
            "categories": ["relay"],
            "protocols": ["zigbee"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/roximo-szbtn01-2s.png",
            "link": "https://iot.roximo.ru/smarthome/zigbee-button-switch-szbtn01-2s",
            "direct": ["hub"],
            "cloud": []
        },
        {
            "manufacturer": "roximo",
            "model": "SZBTN01-2B",
            "name": "Умный Zigbee выключатель ROXIMO, двухкнопочный, чёрный. Подключается через Zigbee-колонку Sber или хаб Sber",
            "categories": ["relay"],
            "protocols": ["zigbee"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/roximo-szbtn01-2b.png",
            "link": "https://iot.roximo.ru/smarthome/zigbee-button-switch-szbtn01-2b",
            "direct": ["hub"],
            "cloud": []
        },
        {
            "manufacturer": "roximo",
            "model": "SZBTN01-2P",
            "name": "Умный Zigbee выключатель ROXIMO, двухкнопочный, платиновый. Подключается через Zigbee-колонку Sber или хаб Sber",
            "categories": ["relay"],
            "protocols": ["zigbee"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/roximo-szbtn01-2p.png",
            "link": "https://iot.roximo.ru/smarthome/zigbee-button-switch-szbtn01-2p",
            "direct": ["hub"],
            "cloud": []
        },
        {
            "manufacturer": "roximo",
            "model": "SZBTN01-2G",
            "name": "Умный Zigbee выключатель ROXIMO, двухкнопочный, золотой. Подключается через Zigbee-колонку Sber или хаб Sber",
            "categories": ["relay"],
            "protocols": ["zigbee"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/roximo-szbtn01-2g.png",
            "link": "https://iot.roximo.ru/smarthome/zigbee-button-switch-szbtn01-2g",
            "direct": ["hub"],
            "cloud": []
        },
        {
            "manufacturer": "roximo",
            "model": "SZBTN01-2C",
            "name": "Умный Zigbee выключатель ROXIMO, двухкнопочный, бронзовый. Подключается через Zigbee-колонку Sber или хаб Sber",
            "categories": ["relay"],
            "protocols": ["zigbee"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/roximo-szbtn01-2c.png",
            "link": "https://iot.roximo.ru/smarthome/zigbee-button-switch-szbtn01-2c",
            "direct": ["hub"],
            "cloud": []
        },
        {
            "manufacturer": "roximo",
            "model": "SZBTN01-3",
            "name": "Умный Zigbee выключатель ROXIMO, трёхкнопочный. Подключается через Zigbee-колонку Sber или хаб Sber",
            "categories": ["relay"],
            "protocols": ["zigbee"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/roximo-szbtn01-3.png",
            "link": "https://iot.roximo.ru/smarthome/zigbee-button-switch-3",
            "direct": ["hub"],
            "cloud": []
        },
        {
            "manufacturer": "roximo",
            "model": "SZBTN01-3W",
            "name": "Умный Zigbee выключатель ROXIMO, трёхкнопочный, белый. Подключается через Zigbee-колонку Sber или хаб Sber",
            "categories": ["relay"],
            "protocols": ["zigbee"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/roximo-szbtn01-3w.png",
            "link": "https://iot.roximo.ru/smarthome/zigbee-button-switch-szbtn01-3w",
            "direct": ["hub"],
            "cloud": []
        },
        {
            "manufacturer": "roximo",
            "model": "SZBTN01-3S",
            "name": "Умный Zigbee выключатель ROXIMO, трёхкнопочный, серый. Подключается через Zigbee-колонку Sber или хаб Sber",
            "categories": ["relay"],
            "protocols": ["zigbee"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/roximo-szbtn01-3s.png",
            "link": "https://iot.roximo.ru/smarthome/zigbee-button-switch-szbtn01-3s",
            "direct": ["hub"],
            "cloud": []
        },
        {
            "manufacturer": "roximo",
            "model": "SZBTN01-3B",
            "name": "Умный Zigbee выключатель ROXIMO, трёхкнопочный, чёрный. Подключается через Zigbee-колонку Sber или хаб Sber",
            "categories": ["relay"],
            "protocols": ["zigbee"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/roximo-szbtn01-3b.png",
            "link": "https://iot.roximo.ru/smarthome/zigbee-button-switch-szbtn01-3b",
            "direct": ["hub"],
            "cloud": []
        },
        {
            "manufacturer": "roximo",
            "model": "SZBTN01-3P",
            "name": "Умный Zigbee выключатель ROXIMO, трёхкнопочный, платиновый. Подключается через Zigbee-колонку Sber или хаб Sber",
            "categories": ["relay"],
            "protocols": ["zigbee"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/roximo-szbtn01-3p.png",
            "link": "https://iot.roximo.ru/smarthome/zigbee-button-switch-szbtn01-3p",
            "direct": ["hub"],
            "cloud": []
        },
        {
            "manufacturer": "roximo",
            "model": "SZBTN01-3G",
            "name": "Умный Zigbee выключатель ROXIMO, трёхкнопочный, золотой. Подключается через Zigbee-колонку Sber или хаб Sber",
            "categories": ["relay"],
            "protocols": ["zigbee"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/roximo-szbtn01-3g.png",
            "link": "https://iot.roximo.ru/smarthome/zigbee-button-switch-szbtn01-3g",
            "direct": ["hub"],
            "cloud": []
        },
        {
            "manufacturer": "roximo",
            "model": "SZBTN01-3C",
            "name": "Умный Zigbee выключатель ROXIMO, трёхкнопочный, бронзовый. Подключается через Zigbee-колонку Sber или хаб Sber",
            "categories": ["relay"],
            "protocols": ["zigbee"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/roximo-szbtn01-3c.png",
            "link": "https://iot.roximo.ru/smarthome/zigbee-button-switch-szbtn01-3c",
            "direct": ["hub"],
            "cloud": []
        },
        {
            "manufacturer": "roximo",
            "model": "SZSEN01-2W",
            "name": "Умный Zigbee выключатель ROXIMO сенсорный, двухкнопочный, белый. Подключается через Zigbee-колонку Sber или хаб Sber",
            "categories": ["relay"],
            "protocols": ["zigbee"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/roximo-szsen01-2w.png",
            "link": "https://iot.roximo.ru/smarthome/zigbee-touch-switch-szsen01-2w",
            "direct": ["hub"],
            "cloud": []
        },
        {
            "manufacturer": "roximo",
            "model": "SZSEN01-2B",
            "name": "Умный Zigbee выключатель ROXIMO сенсорный, двухкнопочный, чёрный. Подключается через Zigbee-колонку Sber или хаб Sber",
            "categories": ["relay"],
            "protocols": ["zigbee"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/roximo-szsen01-2b.png",
            "link": "https://iot.roximo.ru/smarthome/zigbee-touch-switch-szsen01-2b",
            "direct": ["hub"],
            "cloud": []
        },
        {
            "manufacturer": "roximo",
            "model": "SZSEN01-2G",
            "name": "Умный Zigbee выключатель ROXIMO сенсорный, двухкнопочный, золотой. Подключается через Zigbee-колонку Sber или хаб Sber",
            "categories": ["relay"],
            "protocols": ["zigbee"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/roximo-szsen01-2g.png",
            "link": "https://iot.roximo.ru/smarthome/zigbee-touch-switch-szsen01-2g",
            "direct": ["hub"],
            "cloud": []
        },
        {
            "manufacturer": "roximo",
            "model": "SZSEN01-2S",
            "name": "Умный Zigbee выключатель ROXIMO сенсорный, двухкнопочный, серый. Подключается через Zigbee-колонку Sber или хаб Sber",
            "categories": ["relay"],
            "protocols": ["zigbee"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/roximo-szsen01-2s.png",
            "link": "https://iot.roximo.ru/smarthome/zigbee-touch-switch-szsen01-2s",
            "direct": ["hub"],
            "cloud": []
        },
        {
            "manufacturer": "roximo",
            "model": "SZSEN01-3W",
            "name": "Умный Zigbee выключатель ROXIMO сенсорный, трехкнопочный, белый. Подключается через Zigbee-колонку Sber или хаб Sber",
            "categories": ["relay"],
            "protocols": ["zigbee"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/roximo-szsen01-3w.png",
            "link": "https://iot.roximo.ru/smarthome/zigbee-touch-switch-szsen01-3w",
            "direct": ["hub"],
            "cloud": []
        },
        {
            "manufacturer": "roximo",
            "model": "SZSEN01-3B",
            "name": "Умный Zigbee выключатель ROXIMO сенсорный, трехкнопочный, чёрный. Подключается через Zigbee-колонку Sber или хаб Sber",
            "categories": ["relay"],
            "protocols": ["zigbee"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/roximo-szsen01-3b.png",
            "link": "https://iot.roximo.ru/smarthome/zigbee-touch-switch-szsen01-3b",
            "direct": ["hub"],
            "cloud": []
        },
        {
            "manufacturer": "roximo",
            "model": "SZSEN01-3G",
            "name": "Умный Zigbee выключатель ROXIMO сенсорный, трехкнопочный, золотой. Подключается через Zigbee-колонку Sber или хаб Sber",
            "categories": ["relay"],
            "protocols": ["zigbee"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/roximo-szsen01-3g.png",
            "link": "https://iot.roximo.ru/smarthome/zigbee-touch-switch-szsen01-3g",
            "direct": ["hub"],
            "cloud": []
        },
        {
            "manufacturer": "roximo",
            "model": "SZSEN01-3S",
            "name": "Умный Zigbee выключатель ROXIMO сенсорный, трехкнопочный, серый. Подключается через Zigbee-колонку Sber или хаб Sber",
            "categories": ["relay"],
            "protocols": ["zigbee"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/roximo-szsen01-3s.png",
            "link": "https://iot.roximo.ru/smarthome/zigbee-touch-switch-szsen01-3s",
            "direct": ["hub"],
            "cloud": []
        },
        {
            "manufacturer": "roximo",
            "model": "SZSEN01-4W",
            "name": "Умный Zigbee выключатель ROXIMO сенсорный, четырёхкнопочный, белый. Подключается через Zigbee-колонку Sber или хаб Sber",
            "categories": ["relay"],
            "protocols": ["zigbee"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/roximo-szsen01-4w.png",
            "link": "https://iot.roximo.ru/smarthome/zigbee-touch-switch-szsen01-4w",
            "direct": ["hub"],
            "cloud": []
        },
        {
            "manufacturer": "roximo",
            "model": "SZSEN01-4B",
            "name": "Умный Zigbee выключатель ROXIMO сенсорный, четырёхкнопочный, чёрный. Подключается через Zigbee-колонку Sber или хаб Sber",
            "categories": ["relay"],
            "protocols": ["zigbee"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/roximo-szsen01-4b.png",
            "link": "https://iot.roximo.ru/smarthome/zigbee-touch-switch-szsen01-4b",
            "direct": ["hub"],
            "cloud": []
        },
        {
            "manufacturer": "roximo",
            "model": "SZSEN01-4G",
            "name": "Умный Zigbee выключатель ROXIMO сенсорный, четырёхкнопочный, золотой. Подключается через Zigbee-колонку Sber или хаб Sber",
            "categories": ["relay"],
            "protocols": ["zigbee"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/roximo-szsen01-4g.png",
            "link": "https://iot.roximo.ru/smarthome/zigbee-touch-switch-szsen01-4g",
            "direct": ["hub"],
            "cloud": []
        },
        {
            "manufacturer": "roximo",
            "model": "SZSEN01-4S",
            "name": "Умный Zigbee выключатель ROXIMO сенсорный, четырёхкнопочный, серый. Подключается через Zigbee-колонку Sber или хаб Sber",
            "categories": ["relay"],
            "protocols": ["zigbee"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/roximo-szsen01-4s.png",
            "link": "https://iot.roximo.ru/smarthome/zigbee-touch-switch-szsen01-4s",
            "direct": ["hub"],
            "cloud": []
        },
        {
            "manufacturer": "roximo",
            "model": "SZD08",
            "name": "Умный Zigbee датчик открытия дверей и окон ROXIMO. Подключается через Zigbee-колонку Sber или хаб Sber",
            "categories": ["sensor-door"],
            "protocols": ["zigbee"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/roximo-szd08-3.png",
            "link": "https://iot.roximo.ru/smarthome/zigbee-door-window-sensor",
            "direct": ["hub"],
            "cloud": []
        },
        {
            "manufacturer": "roximo",
            "model": "SZW08",
            "name": "Умный Zigbee датчик протечки воды ROXIMO. Подключается через Zigbee-колонку Sber или хаб Sber",
            "categories": ["sensor-water-leak"],
            "protocols": ["zigbee"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/roximo-szw08-3.png",
            "link": "https://iot.roximo.ru/smarthome/zigbee-water-leak-sensor",
            "direct": ["hub"],
            "cloud": []
        },
        {
            "manufacturer": "royal-premium",
            "model": "",
            "name": "Кондиционеры ROYAL Premium. Для управления по Wi-Fi должны быть оснащены модулем Daichi DW21-B/DW22-B",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/royal-premium-hvac-ac.svg",
            "link": "https://daichicloud.ru/split-lineup/search.php?search=ROYAL+Premium",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "sber",
            "model": "SBDV-00050",
            "name": "Умное реле Sber (один канал)",
            "categories": ["relay"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/f1d0fee7_sber-relay-1.svg",
            "link": "https://sberdevices.ru/help/smarthome/relay/about-relay",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "sber",
            "model": "SBDV-00027",
            "name": "Умное реле Sber (один канал)",
            "categories": ["relay"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/f1d0fee7_sber-relay-1.svg",
            "link": "https://sberdevices.ru/help/smarthome/relay/about-relay",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "sber",
            "model": "SBDV-00028",
            "name": "Умное реле Sber (два канала)",
            "categories": ["relay"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/82120961_sber-relay-2.svg",
            "link": "https://sberdevices.ru/help/smarthome/two-channel-relay/about-relay",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "sber",
            "model": "SBDV-00029",
            "name": "Умный датчик движения Sber. Подключается через Zigbee-колонку Sber или хаб Sber",
            "categories": ["sensor-pir"],
            "protocols": ["zigbee"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/345dd31d_sber-sensor-pir-1.png",
            "link": "https://sberdevices.ru/shop/product/smart_motion_sensor/",
            "direct": ["hub"],
            "cloud": []
        },
        {
            "manufacturer": "sber",
            "model": "SBDV-00030",
            "name": "Умный датчик открытия Sber. Подключается через Zigbee-колонку Sber или хаб Sber",
            "categories": ["sensor-door"],
            "protocols": ["zigbee"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/800b5b1a_sber-sensor-door-1.png",
            "link": "https://sberdevices.ru/shop/product/smart_open_sensor/",
            "direct": ["hub"],
            "cloud": []
        },
        {
            "manufacturer": "sber",
            "model": "SBDV-00154",
            "name": "Умный датчик протечки Sber. Подключается через Zigbee-колонку Sber или хаб Sber",
            "categories": ["sensor-water-leak"],
            "protocols": ["zigbee"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/bbcb087d_sber-sensor-water-leak-1.png",
            "link": "https://sberdevices.ru/shop/product/smart_leak_sensor/",
            "direct": ["hub"],
            "cloud": []
        },
        {
            "manufacturer": "sber",
            "model": "SBDV-00079",
            "name": "Умный датчик температуры и влажности Sber. Подключается через Zigbee-колонку Sber или хаб Sber",
            "categories": ["sensor-temp"],
            "protocols": ["zigbee"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/6b78ee1d_sber-sensor-temp-1.png",
            "link": "https://sberdevices.ru/shop/product/smart_climat_sensor/",
            "direct": ["hub"],
            "cloud": []
        },
        {
            "manufacturer": "sber",
            "model": "SBDV-00090",
            "name": "Акустическая система SberBoom",
            "categories": ["controller"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/12da9093_sber-sberboom.png",
            "link": "https://sberdevices.ru/shop/product/sberboom_black/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "sber",
            "model": "SBDV-00171",
            "name": "Акустическая система SberBoom Home",
            "categories": ["controller"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/84d354e9_sberboomhome-main.png",
            "link": "https://sberdevices.ru/shop/product/sberboom_home_gray/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "sber",
            "model": "SBDV-00095",
            "name": "Акустическая система SberBoom Mini",
            "categories": ["controller"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/3d1a7dca_sber-sberboommini.png",
            "link": "https://sberdevices.ru/shop/product/sberboom_mini_lightblue/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "sber",
            "model": "SBDV-00250",
            "name": "Акустическая система SberBoom Mini 2",
            "categories": ["controller"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/sberboommini2.png",
            "link": "https://sberdevices.ru/shop/product/sberboom_mini_2/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "sber",
            "model": "SBDV-00026",
            "name": "Медиаколонка SberBox Time",
            "categories": ["controller"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/6bfe79f5_sber-sberboxtime.png",
            "link": "https://sberdevices.ru/shop/product/sberboxtime_black/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "sber",
            "model": "SBDV-00002",
            "name": "Приставка SberBox",
            "categories": ["controller"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/7951d6d9_sber-sberbox.png",
            "link": "https://sberdevices.ru/shop/product/sberbox_3/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "sber",
            "model": "SBDV-00006",
            "name": "Приставка SberBox 2",
            "categories": ["controller"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/sberbox2.png",
            "link": "https://sberdevices.ru/shop/product/sberbox2/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "sber",
            "model": "SBDV-00010",
            "name": "Смарт-дисплей SberPortal",
            "categories": ["controller"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/95086473_sber-sberportal-black.png",
            "link": "https://sberdevices.ru/shop/product/sberportal_black/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "sber",
            "model": "SBDV-00013",
            "name": "ТВ-медиацентр SberBox Top",
            "categories": ["controller"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/a1a10179_sber-sberboxtop.png",
            "link": "https://sberdevices.ru/shop/product/sberboxtop/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "sber",
            "model": "",
            "name": "Телевизоры Sber",
            "categories": ["controller"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/c20958a0_sber-sbertv.png",
            "link": "https://sberdevices.ru/shop/category/tv/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "sber",
            "model": "",
            "name": "Телевизоры сторонних производителей, работающие на платформе Салют ТВ от Sber",
            "categories": ["controller"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/869f2f18_sber-salutetv.png",
            "link": null,
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "sber",
            "model": "SWB-CW7RA",
            "name": "Контроллер умного дома Sber",
            "categories": ["controller"],
            "protocols": ["mqtt"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/sber-controller.png",
            "link": "https://developers.sber.ru/docs/ru/smarthome-plc/overview",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "sber",
            "model": "SBDV-00115",
            "name": "Умная лампа Sber SBDV-00115. Цоколь E27, колба А60 («груша»)",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/b76d9eab_sber-light-1-2.svg",
            "link": "https://sberdevices.ru/shop/product/smart_lamp_sber/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "sber",
            "model": "SBDV-00019",
            "name": "Умная лампа Sber SBDV-00019. Цоколь E27, колба А60 («груша»)",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/99f71734_sber-light-2-1.svg",
            "link": "https://sberdevices.ru/help/smarthome/bulb/bulb-features",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "sber",
            "model": "SBDV-00117",
            "name": "Умная лампа Sber SBDV-00117. Цоколь E14, колба С37 («свеча»)",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/2b2ba17b_sber-light-3.svg",
            "link": "https://sberdevices.ru/shop/product/smartlamp_e14/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "sber",
            "model": "SBDV-00020",
            "name": "Умная лампа Sber SBDV-00020. Цоколь E14, колба С37 («свеча»)",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/ed1a0cca_sber-light-4-1.svg",
            "link": "https://sberdevices.ru/help/smarthome/bulb/bulb-features",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "sber",
            "model": "SBDV-00024",
            "name": "Умная лампа Sber SBDV-00024. Цоколь GU10, колба MR16 («рефлектор»)",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/18a35218_sber-light-5-1.svg",
            "link": "https://sberdevices.ru/help/smarthome/bulb/bulb-features",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "sber",
            "model": "SBDV-00125",
            "name": "Умный ночник Sber «Кошка»",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/472fd814_sber-light-6-1.png",
            "link": "https://sberdevices.ru/shop/product/detskiy_nochnik_koshka/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "sber",
            "model": "SBDV-00127",
            "name": "Умный ночник Sber «Грибочек»",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/c096e36f_sber-light-7.png",
            "link": "https://sberdevices.ru/shop/product/detskiy_nochnik_grib/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "sber",
            "model": "SBDV-00126",
            "name": "Умный ночник Sber «Космонавт»",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/6cb36ca4_sber-light-8.png",
            "link": "https://sberdevices.ru/shop/product/detskiy_nochnik_kosmonavt/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "sber",
            "model": "SBDV-00123",
            "name": "Умная розетка Sber SBDV-00123",
            "categories": ["socket"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/8eb314bf_sber-socket-1.png",
            "link": "https://sberdevices.ru/shop/product/smartsocket/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "sber",
            "model": "SBDV-00025",
            "name": "Умная розетка Sber SBDV-00025",
            "categories": ["socket"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/920e6945_sber-socket-2.svg",
            "link": "https://sbermarket.ru/products/19786079-umnaya-rozetka-sber-sbdv-00025-bluetooth",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "sber",
            "model": "SBDV-00018",
            "name": "Умная розетка Sber SBDV-00018",
            "categories": ["socket"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/920e6945_sber-socket-2.svg",
            "link": "https://sbermarket.ru/products/16183943-umnaya-rozetka-sber-sbdv-00018-belaya-16183943",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "sber",
            "model": "SBDV-00033R",
            "name": "Умная cветодиодная лента Sber",
            "categories": ["led-strip"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/79483e2c_sber-led-strip-1.png",
            "link": "https://sberdevices.ru/shop/product/smart_led/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "sber",
            "model": "SBDV-00038",
            "name": "Удлинитель светодиодной ленты Sber, 1 м",
            "categories": ["led-strip"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/79483e2c_sber-led-strip-1.png",
            "link": "https://sberdevices.ru/shop/product/smart_led_extension/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "sber",
            "model": "SBDV-00032",
            "name": "Умная кнопка Sber. Подключается через Zigbee-колонку Sber или хаб Sber",
            "categories": ["scenario-button"],
            "protocols": ["zigbee"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/23a78cf2_sber-scenario-button-1.png",
            "link": "https://sberdevices.ru/shop/product/smart_button/",
            "direct": ["hub"],
            "cloud": []
        },
        {
            "manufacturer": "sber",
            "model": "SBDV-00185",
            "name": "Умный терморегулятор Sber. Подключается через Zigbee-колонку Sber или хаб Sber",
            "categories": ["hvac-radiator"],
            "protocols": ["zigbee"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/e780e9b6_trv.png",
            "link": "https://sberdevices.ru/shop/product/smart_thermoregulator/",
            "direct": ["hub"],
            "cloud": []
        },
        {
            "manufacturer": "sber",
            "model": "SBDV-00068",
            "name": "Умный хаб Sber",
            "categories": ["hub"],
            "protocols": ["zigbee"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/d348ff21_sber-hub-1.png",
            "link": "https://sberdevices.ru/shop/product/smart_hab/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "scarlett",
            "model": "SC-VC80RW01",
            "name": "Робот-пылесос SCARLETT SC-VC80RW01",
            "categories": ["vacuum-cleaner"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/a2871ca8_scarlett-itm-sc-vc80rw01-3.svg",
            "link": "https://www.scarlett.ru/catalog/dlya-doma/pylesosy/itm_sc-vc80rw01/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "shuft",
            "model": "",
            "name": "Кондиционеры Shuft. Для управления по Wi-Fi должны быть оснащены модулем Daichi DW21-B/DW22-B",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/shuft-hvac-ac.svg",
            "link": "https://daichicloud.ru/split-lineup/search.php?search=Shuft",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "sibling",
            "model": "Powerlite-L (8w)",
            "name": "Умная RGBW лампочка Sibling Е27 (8 Вт)",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/sibling-light-1.svg",
            "link": "https://sibling.ru/all/light/umnaya-lampochka-87/umnaya-rgbw-lampochka-e27-8-vt-110/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "sibling",
            "model": "Powerlite-L (12w)",
            "name": "Умная RGBW лампочка Sibling Е27 (12 Вт)",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/sibling-light-2.svg",
            "link": "https://sibling.ru/all/light/umnaya-lampochka-87/umnaya-rgbw-lampochka-e27-12-vt-137/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "sibling",
            "model": "Powerlite-L (15w)",
            "name": "Умная RGBW лампочка Sibling Е27 (15 Вт)",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/sibling-light-3.svg",
            "link": "https://sibling.ru/all/light/umnaya-lampochka-87/umnaya-rgbw-lampochka-e27-15-vt-138/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "sibling",
            "model": "Powerlite-L (G45)",
            "name": "Умная RGBW лампочка-шар Sibling Е14 (5 Вт)",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/sibling-light-4.svg",
            "link": "https://sibling.ru/all/light/umnaya-lampochka-87/umnaya-rgbw-lampochka-shar-e14-5-vt-152/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "sibling",
            "model": "Powerlite-L (С37)",
            "name": "Умная RGBW лампочка-свеча Sibling Е14 (5 Вт)",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/sibling-light-5.svg",
            "link": "https://sibling.ru/all/light/umnaya-lampochka-87/umnaya-rgbw-lampochka-svecha-e14-5-vt-127/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "sibling",
            "model": "Powerlite-L (GU10)",
            "name": "Умная RGBW лампочка Sibling GU10 (5 Вт)",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/sibling-light-6.svg",
            "link": "https://sibling.ru/all/light/umnaya-lampochka-87/umnaya-rgbw-lampochka-gu10-5-vt-151/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "sibling",
            "model": "Powerlite-L (G95)",
            "name": "Умная RGBW лампа-шар Sibling Е27 (13 Вт)",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/sibling-light-7.svg",
            "link": "https://sibling.ru/all/light/umnaya-lampochka-87/umnaya-rgbw-lampa-shar-e27-13-vt-122/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "sibling",
            "model": "Powerspace-OHB25",
            "name": "Умный масляный обогреватель Sibling (чёрный)",
            "categories": ["hvac-heater"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/sibling-radiator-1.svg",
            "link": "https://sibling.ru/all/tehnika-129/obogrevateli-131/umnyiy-obogrevatel-maslyanyiy-chёrnyiy-259/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "sibling",
            "model": "Powerspace-OHW25",
            "name": "Умный масляный обогреватель Sibling (белый)",
            "categories": ["hvac-heater"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/sibling-radiator-2.svg",
            "link": "https://sibling.ru/all/tehnika-129/obogrevateli-131/umnyiy-obogrevatel-maslyanyiy-belyiy-260/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "sibling",
            "model": "Powerspace-GHB10",
            "name": "Умный конвекционный обогреватель Sibling (чёрный)",
            "categories": ["hvac-heater"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/sibling-radiator-3.svg",
            "link": "https://sibling.ru/all/tehnika-129/obogrevateli-131/umnyiy-obogrevatel-elektricheskiy-chёrnyiy-258/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "sibling",
            "model": "Powerspace-GHW10",
            "name": "Умный конвекционный обогреватель Sibling (белый)",
            "categories": ["hvac-heater"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/sibling-radiator-4.svg",
            "link": "https://sibling.ru/all/tehnika-129/obogrevateli-131/umnyiy-obogrevatel-elektricheskiy-belyiy-263/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "sibling",
            "model": "Powerspace-VC",
            "name": "Умный робот-пылесос Sibling",
            "categories": ["vacuum-cleaner"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/sibling-vacuum-1-1.svg",
            "link": "https://sibling.ru/all/tehnika-129/pyilesosyi-132/umnyiy-robot-pyilesos-262/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "sibling",
            "model": "Powerspace-VCL",
            "name": "Умный робот-пылесос Sibling с лидаром",
            "categories": ["vacuum-cleaner"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/sibling-vacuum-2-1.svg",
            "link": "https://sibling.ru/all/tehnika-129/pyilesosyi-132/umnyiy-robot-pyilesos-s-lidarom-362/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "sibling",
            "model": "Powerswitch 16А",
            "name": "Умная розетка Sibling с мониторингом электроэнергии",
            "categories": ["socket"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/sibling-socket-1.svg",
            "link": "https://sibling.ru/all/smart220/rozetki-110/umnaya-rozetka-115/rozetka-s-rashodomerom-66/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "sibling",
            "model": "Powerswitch-F",
            "name": "Умная розетка Sibling люкс с мониторингом электроэнергии",
            "categories": ["socket"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/sibling-socket-2.svg",
            "link": "https://sibling.ru/all/rozetka-s-rashodomerom-lyuks-115/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "sibling",
            "model": "Powersocket-1PB(S)",
            "name": "Умная розетка Sibling в рамку (черная)",
            "categories": ["socket"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/sibling-socket-3.svg",
            "link": "https://sibling.ru/all/electromodule/rozetka-v-ramku-umnaya-chernaya-289/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "sibling",
            "model": "Powersocket-1PW(S)",
            "name": "Умная розетка Sibling в рамку (белая)",
            "categories": ["socket"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/sibling-socket-4.svg",
            "link": "https://sibling.ru/all/electromodule/rozetka-v-ramku-umnaya-belaya-298/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "sibling",
            "model": "Powerspace-SK",
            "name": "Умный электрический чайник Sibling (бело-серый)",
            "categories": ["kettle"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/sibling-kettle-1.svg",
            "link": "https://sibling.ru/all/tehnika-129/chayniki-130/umnyiy-elektricheskiy-chaynik-belo-seryiy-271/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "sibling",
            "model": "Powerspace-SK3",
            "name": "Умный электрический чайник Sibling (сине-зеленый)",
            "categories": ["kettle"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/sibling-kettle-2.svg",
            "link": "https://sibling.ru/all/tehnika-129/chayniki-130/umnyiy-elektricheskiy-chaynik-sine-zelёnyiy-270/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "sibling",
            "model": "Powerspace-SK1",
            "name": "Умный электрический чайник Sibling (черный)",
            "categories": ["kettle"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/sibling-kettle-3.svg",
            "link": "https://sibling.ru/all/tehnika-129/chayniki-130/umnyiy-elektricheskiy-chaynik-chёrnyiy-216/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "tricolor",
            "model": "GS SMHM-I1",
            "name": "Датчик движения SmartHome Tricolor. Подключается через хаб Tricolor",
            "categories": ["sensor-pir"],
            "protocols": ["zigbee"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/f6ab7593_tricolor-sensor-pir-1.svg",
            "link": "https://home.tricolor.ru/equipment/datchik-dvizheniya-gs-smhm-i1/",
            "direct": [],
            "cloud": ["с2с-hub"]
        },
        {
            "manufacturer": "tricolor",
            "model": "GS SOHM-I1",
            "name": "Датчик открытия и закрытия SmartHome Tricolor. Подключается через хаб Tricolor",
            "categories": ["sensor-door"],
            "protocols": ["zigbee"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/cad8d506_tricolor-sensor-door-1.jpeg",
            "link": "https://home.tricolor.ru/equipment/datchik-otkrytiya-i-zakrytiya-gs-sohm-i1/",
            "direct": [],
            "cloud": ["с2с-hub"]
        },
        {
            "manufacturer": "tricolor",
            "model": "GS STHM-I1H",
            "name": "Датчик температуры и влажности SmartHome Tricolor. Подключается через хаб Tricolor",
            "categories": ["sensor-temp"],
            "protocols": ["zigbee"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/cd0b59da_tricolor-sensor-temp-1.svg",
            "link": null,
            "direct": [],
            "cloud": ["с2с-hub"]
        },
        {
            "manufacturer": "tricolor",
            "model": "GS BDHM8E27W70-I1",
            "name": "Умная лампа SmartHome Tricolor. Подключается через хаб Tricolor, Zigbee-колонку Sber или хаб Sber",
            "categories": ["light"],
            "protocols": ["zigbee"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/260d87ed_tricolor-light-1.jpeg",
            "link": "https://home.tricolor.ru/equipment/umnaya-lampa-gs-bdhm8e27w70-i1/",
            "direct": ["hub"],
            "cloud": ["с2с-hub"]
        },
        {
            "manufacturer": "tricolor",
            "model": "GS BRHM8E27W70-I1",
            "name": "Умная лампа цветная SmartHome Tricolor. Подключается через хаб Tricolor, Zigbee-колонку Sber или хаб Sber",
            "categories": ["light"],
            "protocols": ["zigbee"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/52618993_tricolor-light-2.jpeg",
            "link": "https://home.tricolor.ru/equipment/umnaya-lampa-tsvetnaya-gs-brhm8e27w70-i1/",
            "direct": ["hub"],
            "cloud": ["с2с-hub"]
        },
        {
            "manufacturer": "tricolor",
            "model": "GS SKHMP30-I1",
            "name": "Умная розетка SmartHome Tricolor. Подключается через хаб Tricolor",
            "categories": ["socket"],
            "protocols": ["zigbee"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/eec17458_tricolor-socker-1.jpeg",
            "link": "https://home.tricolor.ru/equipment/umnaya-rozetka-gs-skhmp30-i1/",
            "direct": [],
            "cloud": ["с2с-hub"]
        },
        {
            "manufacturer": "sonoff",
            "model": "BASICZBR3",
            "name": "Реле Sonoff BASICZBR3. Подключается через Zigbee-колонку Sber или хаб Sber",
            "categories": ["relay"],
            "protocols": ["zigbee"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/8b7165cc_sonoff-relay-2.svg",
            "link": null,
            "direct": ["hub"],
            "cloud": []
        },
        {
            "manufacturer": "sonoff",
            "model": "ZBMINI",
            "name": "Реле с нулём Sonoff ZBMINI. Подключается через Zigbee-колонку Sber или хаб Sber",
            "categories": ["relay"],
            "protocols": ["zigbee"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/a59a60d2_sonoff-01MINIZB.png",
            "link": "https://sonoff.tech/product/diy-smart-switches/zbmini/",
            "direct": ["hub"],
            "cloud": []
        },
        {
            "manufacturer": "sonoff",
            "model": "01MINIZB",
            "name": "Реле с нулём Sonoff 01MINIZB. Подключается через Zigbee-колонку Sber или хаб Sber",
            "categories": ["relay"],
            "protocols": ["zigbee"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/a59a60d2_sonoff-01MINIZB.png",
            "link": "",
            "direct": ["hub"],
            "cloud": []
        },
        {
            "manufacturer": "sonoff",
            "model": "ZBMINI-L",
            "name": "Реле без нуля Sonoff ZBMINIL. Подключается через Zigbee-колонку Sber или хаб Sber",
            "categories": ["relay"],
            "protocols": ["zigbee"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/sonoff-zbminil.png",
            "link": "https://sonoff.ru/catalog/zigbee-ustroystva/rele-bez-nulya-sonoff-zbmini-l-zigbee-3-0-smart-switch",
            "direct": ["hub"],
            "cloud": []
        },
        {
            "manufacturer": "sonoff",
            "model": "ZBMINI-L2",
            "name": "Реле без нуля Sonoff ZBMINIL2 Extreme. Подключается через Zigbee-колонку Sber или хаб Sber",
            "categories": ["relay"],
            "protocols": ["zigbee"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/sonoff-ZBMINIL2.png",
            "link": "https://sonoff.tech/product/diy-smart-switches/zbmini-l2/",
            "direct": ["hub"],
            "cloud": []
        },
        {
            "manufacturer": "sonoff",
            "model": "DW2-Wi-Fi",
            "name": "Датчик открытия двери или окна Sonoff DW2-Wi-Fi",
            "categories": ["sensor-door"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/sonoff-sl1000.png",
            "link": "https://sonoff.tech/product-document/gateway-and-sensors-doc/dw2-wifi-doc/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "sonoff",
            "model": "B02/B05-BL",
            "name": "Умная Wi-Fi лампа Sonoff B02-BL/B05-BL",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/sonoff-b05-01.png",
            "link": "https://sonoff.tech/product-document/smart-lighting-doc/b02-b05-bl-doc/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "sonoff",
            "model": "S26R2ZB",
            "name": "Умная Zigbee розетка Sonoff S26R2ZB. Подключается через Zigbee-колонку Sber или хаб Sber",
            "categories": ["socket"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/sonoff-S26R2ZB.png",
            "link": "https://www.sonoff.ru/product/rozetka-sonoff-s26r2zb-zigbee",
            "direct": ["hub"],
            "cloud": []
        },
        {
            "manufacturer": "sonoff",
            "model": "S31/S31LITE",
            "name": "Умная Wi-Fi розетка Sonoff S31/S31 LITE",
            "categories": ["socket"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/sonoff-s31litezb.png",
            "link": "https://sonoff.tech/product-document/smart-plugs-doc/s31-s31lite-doc/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "tesla",
            "model": "",
            "name": "Кондиционеры Tesla. Для управления по Wi-Fi должны быть оснащены модулем Daichi DW21-B/DW22-B",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/tesla-hvac-ac.svg",
            "link": "https://daichicloud.ru/split-lineup/search.php?search=TESLA",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "timberk",
            "model": "T-VCR-53WI-TBN",
            "name": "Робот-пылесос с Wi-Fi Timberk T-VCR-53WI-TBN",
            "categories": ["vacuum-cleaner"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/28c40fec_timberk-vcr-53wi-tbn-s-2.svg",
            "link": "https://www.timberk.ru/catalog/product/t-vcr-53wi-tbn/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "timberk",
            "model": "T-EK21S104WF",
            "name": "Чайник электрический с Wi-Fi Timberk T-EK21S104WF",
            "categories": ["kettle"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/timberk-kettle-1.svg",
            "link": "https://www.timberk.ru/catalog/product/t-ek21s104wf/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "timberk",
            "model": "T-EK21S103WF",
            "name": "Чайник электрический с Wi-Fi Timberk T-EK21S103WF",
            "categories": ["kettle"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/timberk-kettle-2.svg",
            "link": "https://www.timberk.ru/catalog/product/t-ek21s103wf/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "tuya",
            "model": "WHD02",
            "name": "Реле Zigbee WHD02 Tuya. Подключается через Zigbee-колонку Sber или хаб Sber",
            "categories": ["relay"],
            "protocols": ["zigbee"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/1637f100_tuya-WHD02.png",
            "link": null,
            "direct": ["hub"],
            "cloud": []
        },
        {
            "manufacturer": "ujin",
            "model": "UCI-W-1С16A-M",
            "name": "Коммутатор встраиваемый Ujin Connect-in",
            "categories": ["relay"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/uci-w-116a-m-01.svg",
            "link": "https://ujin.tech/ujin-connect-in/uci-w-1c16a-m",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "ujin",
            "model": "UCD-WBG-2С16+16A-2D",
            "name": "Коммутатор на DIN-рейку Ujin Connect-din",
            "categories": ["relay"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/ucd-wbg-216-16a-2d-0.svg",
            "link": "https://ujin.tech/ujin-connect-din/ucd-wbg-2c1616a-2d",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "ujin",
            "model": "UCD-W-2С16+16A-2D",
            "name": "Коммутатор на DIN-рейку Ujin Connect-din",
            "categories": ["relay"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/ucd-w-216-16a-2d-216-01.svg",
            "link": "https://ujin.tech/ujin-connect-din/ucd-w-2c1616a-2d",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "ujin",
            "model": "UAS-B",
            "name": "Датчик протечки Ujin Aqua-sense. Подключается через хаб Ujin",
            "categories": ["sensor-water-leak"],
            "protocols": ["bluetooth"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/uas-b-01.svg",
            "link": "https://ujin.tech/ujin-aqua-sense/uas-b",
            "direct": [],
            "cloud": ["с2с-hub"]
        },
        {
            "manufacturer": "ujin",
            "model": "EP-WB-T-MI",
            "name": "Мультисенсор Ujin Pulse",
            "categories": ["sensor-temp"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/ep-wb-t-mi-01.svg",
            "link": "https://ujin.tech/ujin-pulse/ep-wb-t-mi",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "ujin",
            "model": "EH-WB-T-I-16А",
            "name": "Термостат Ujin Heat",
            "categories": ["hvac-boiler", "hvac-underfloor-heating"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/eh-wb-t-i-16-01.svg",
            "link": "https://ujin.tech/ujin-heat/eh-wb-t-i-16a",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "ujin",
            "model": "EA-WB-12V-CR1",
            "name": "Контроллер протечки Ujin Aqua",
            "categories": ["controller"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/ea-wb-12v-cr1-01.svg",
            "link": "https://ujin.tech/ujin-aqua/ea-wb-12v-cr1",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "ujin",
            "model": "EA-WBZ-12V-CR1",
            "name": "Контроллер протечки Ujin Aqua",
            "categories": ["controller"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/ea-wb-12v-cr1-01.svg",
            "link": "https://ujin.tech/ujin-aqua/ea-wb-12v-cr1",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "ujin",
            "model": "EA-WB-220V-CR2",
            "name": "Контроллер протечки Ujin Aqua",
            "categories": ["controller"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/ea-wb-12v-cr1-01.svg",
            "link": "https://ujin.tech/ujin-aqua/ea-wb-220v-cr2",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "ujin",
            "model": "EA-WB-220V-CR3",
            "name": "Контроллер протечки Ujin Aqua",
            "categories": ["controller"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/ea-wb-12v-cr1-01.svg",
            "link": "https://ujin.tech/ujin-aqua/ea-wb-220v-cr3",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "ujin",
            "model": "ULI-WB-2С",
            "name": "Диммер электрический Ujin Lume-in",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/uli-wb-2-01.svg",
            "link": "https://ujin.tech/ujin-lume-in/uli-wb-2c",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "vixion",
            "model": "",
            "name": "О совместимости устройств Vixion Home с умным домом Sber уточняйте у продавца или производителя",
            "categories": ["other"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/04adc0ef_unknown.svg",
            "link": null,
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "welrok",
            "model": "welrok_az",
            "name": "Умный Wi-Fi терморегулятор для тёплых полов Welrok AZ",
            "categories": ["hvac-underfloor-heating"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/62dde3db_welrok-az-white.svg",
            "link": "https://megamarket.ru/catalog/details/termoregulyator-welrok-az-dlya-upravleniya-teplym-polom-cherez-smartfon-ili-kompyuter-600014018330/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "wisehome",
            "model": "1",
            "name": "Wise Home Контроллер Wise Leak Protector",
            "categories": ["controller"],
            "protocols": ["mqtt"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/d25f4ba7_wisehome-controller-1.svg",
            "link": null,
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "wisehome",
            "model": "2",
            "name": " Wise Home Контроллер Wise Electro",
            "categories": ["controller"],
            "protocols": ["mqtt"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/21f4c7f9_wisehome-controller-2.svg",
            "link": null,
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "yeelight",
            "model": "YLCT01YL",
            "name": "Yeelight LED Bedside Lamp D2",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/ebe3590b_yeelight-light-1.jpeg",
            "link": "https://ru.yeelight.com/catalog/umnaya-prikrovatnaya-lampa-yeelight-led-bedside-lamp-d2/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "yeelight",
            "model": "YLCT03YL",
            "name": "Yeelight Star Smart Desk Table Lamp Pro",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/51ac5ae5_yeelight-light-2.jpeg",
            "link": "https://ru.yeelight.com/catalog/svetilnik-nastolnyy-yeelight-staria-bedside-lamp-pro/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "yeelight",
            "model": "YLTS02YL",
            "name": "Yeelight Mesh Downlight M2",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/8cfeaf95_yeelight-light-3.jpeg",
            "link": "https://ru.yeelight.com/catalog/umnyy-vstraivaemyy-svetilnik-yeelight-mesh-downlight-m2/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "yeelight",
            "model": "YLTS04YL",
            "name": "Yeelight Mesh Spotlight M2",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/8847b331_yeelight-light-4.jpeg",
            "link": "https://ru.yeelight.com/catalog/umnyy-vstraivaemyy-svetilnik-yeelight-mesh-spotlight-m2/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "yeelight",
            "model": "YLDL01YL",
            "name": "Yeelight Crystal Pendant Lamp",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/b38cef90_yeelight-light-5.jpeg",
            "link": "https://ru.yeelight.com/catalog/umnyy-vstraivaemyy-svetilnik-yeelight-mesh-spotlight-m2/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "yeelight",
            "model": "YLDP004",
            "name": "Yeelight GU10 Smart bulb W1 (Dimmable)",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/21a5f567_yeelight-light-6.jpeg",
            "link": "https://ru.yeelight.com/catalog/umnaya-lampochka-yeelight-gu10-smart-bulb-w1-dimmable/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "yeelight",
            "model": "YLDP004-A",
            "name": "Yeelight GU10 Smart bulb (Multicolor)",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/510ce54e_yeelight-light-7.svg",
            "link": "https://ru.yeelight.com/catalog/umnaya-lampochka-yeelight-gu10-smart-bulb-multicolor/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "yeelight",
            "model": "YLDP005",
            "name": "Yeelight Smart LED Bulb W3 (Multiple color)",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/80358f7c_yeelight-light-8.jpeg",
            "link": "https://ru.yeelight.com/catalog/umnaya-led-lampochka-yeelight-smart-led-bulb-w3-multicolor/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "yeelight",
            "model": "YLDP007",
            "name": "Yeelight Smart LED Bulb W3 (White)",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/e5174330_yeelight-light-9.svg",
            "link": "https://ru.yeelight.com/catalog/umnaya-led-lampochka-yeelight-smart-led-bulb-w3-white-dimmable/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "yeelight",
            "model": "YLDP11YL / YLDP12YL",
            "name": "Yeelight LED Filament Light",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/e8d4b113_yeelight-light-10.jpeg",
            "link": "https://ru.yeelight.com/catalog/umnaya-svetodiodnaya-filamentnaya-lampa-yeelight-led-filament-bulb/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "yeelight",
            "model": "YLDP13YL",
            "name": "Yeelight Smart LED Bulb 1S (Color)",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/af3136da_yeelight-light-11.jpeg",
            "link": "https://ru.yeelight.com/catalog/umnaya-led-lampochka-yeelight-smart-led-bulb-1s-color/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "yeelight",
            "model": "YLFWD-012",
            "name": "Yeelight Smart Light Panels",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/36c89de4_yeelight-light-12.jpeg",
            "link": "https://ru.yeelight.com/catalog/svetodiodnaya-panel-yeelight-smart-light-panels-6-sht/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "yeelight",
            "model": "YLTD003",
            "name": "Yeelight LED Screen Light bar Pro",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/6035cfcc_yeelight-light-13.svg",
            "link": "https://ru.yeelight.com/catalog/svetodiodnaya-panel-yeelight-led-screen-light-bar-pro/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "yeelight",
            "model": "YLTD04YL",
            "name": "Yeelight Serene Eye-friendly Desk Lamp Pro",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/298a7516_yeelight-light-14.jpeg",
            "link": "https://ru.yeelight.com/catalog/svetodiodnaya-nastolnaya-lampa-yeelight-serene-eye-friendly-lamp-pro/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "yeelight",
            "model": "YLTD08YL",
            "name": "Yeelight LED Light-sensitive desk lamp V1 Pro",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/40a4bc0d_yeelight-light-15.jpeg",
            "link": "https://ru.yeelight.com/catalog/svetodiodnaya-nastolnaya-lampa-yeelight-led-vision-desk-lamp-v1-pro-s-osnovaniem/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "yeelight",
            "model": "YLTD13YL",
            "name": "Yeelight LED Light-sensitive desk lamp V1 Pro",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/24e0a55d_yeelight-light-16.jpeg",
            "link": "https://ru.yeelight.com/catalog/svetodiodnaya-nastolnaya-lampa-yeelight-led-vision-desk-lamp-v1-pro-s-zazhimom/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "yeelight",
            "model": "YLXD013",
            "name": "Yeelight Arwen Ceiling Light 450S",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/2247130a_yeelight-light-17.svg",
            "link": "https://ru.yeelight.com/catalog/umnyy-potolochnyy-svetilnik-yeelight-arwen-ceiling-light-450s/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "yeelight",
            "model": "YLXD013-С",
            "name": "Yeelight Arwen Ceiling Light 550C",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/652adcc8_yeelight-light-18.svg",
            "link": "https://ru.yeelight.com/catalog/umnyy-potolochnyy-svetilnik-yeelight-arwen-ceiling-light-550c/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "yeelight",
            "model": "YLXD013-A",
            "name": "Yeelight Arwen Ceiling Light 550S",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/93e2d001_yeelight-light-19.svg",
            "link": "https://ru.yeelight.com/catalog/umnyy-potolochnyy-svetilnik-yeelight-arwen-ceiling-light-550s-ylxd013-a/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "yeelight",
            "model": "YLXD013-B",
            "name": "Yeelight Arwen Ceiling Light 450C",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/5f01a48d_yeelight-light-20.svg",
            "link": "https://ru.yeelight.com/catalog/umnyy-potolochnyy-svetilnik-yeelight-arwen-ceiling-light-450c/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "yeelight",
            "model": "YLXD031",
            "name": "Yeelight A2001C550 Ceiling Light",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/d3350dd1_yeelight-light-21.svg",
            "link": "https://ru.yeelight.com/catalog/umnyy-potolochnyy-svetilnik-yeelight-a2001c550-ceiling-light/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "yeelight",
            "model": "YLXD032",
            "name": "Yeelight A2001C450 Ceiling Light",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/9c7dddb5_yeelight-light-22.svg",
            "link": "https://ru.yeelight.com/catalog/umnyy-potolochnyy-svetilnik-yeelight-a2001c450-ceiling-light/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "yeelight",
            "model": "YLXD033",
            "name": "Yeelight A2001R900 Ceiling Light",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/bbb795d2_yeelight-light-23.svg",
            "link": "https://ru.yeelight.com/catalog/umnyy-potolochnyy-svetilnik-yeelight-a2001r900-ceiling-light/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "yeelight",
            "model": "YLXD036",
            "name": "Yeelight C2001C450 Ceiling Light (450 mm)",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/808261c0_yeelight-light-24.jpeg",
            "link": "https://ru.yeelight.com/catalog/umnyy-potolochnyy-svetilnik-yeelight-c2001c450-ceiling-light/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "yeelight",
            "model": "YLXD037",
            "name": "Yeelight C2001C550 Ceiling Light (550 mm)",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/616dea97_yeelight-light-25.jpeg",
            "link": "https://ru.yeelight.com/catalog/umnyy-potolochnyy-svetilnik-yeelight-c2001c550-ceiling-light/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "yeelight",
            "model": "YLXD038",
            "name": "Yeelight C2001S500 Ceiling Light (500 mm)",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/91050580_yeelight-light-26.svg",
            "link": "https://ru.yeelight.com/catalog/umnyy-potolochnyy-svetilnik-yeelight-c2001s500-ceiling-light/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "yeelight",
            "model": "YLXD039",
            "name": "Yeelight C2001R900 Ceiling Light (900 mm)",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/23458a2f_yeelight-light-27.svg",
            "link": "https://ru.yeelight.com/catalog/umnyy-potolochnyy-svetilnik-yeelight-c2001r900-ceiling-light/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "yeelight",
            "model": "YLXD49YL",
            "name": "Yeelight Halo Ceiling Light Pro",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/1991b76b_yeelight-light-28.svg",
            "link": "https://ru.yeelight.com/catalog/umnyy-potolochnyy-svetilnik-yeelight-halo-ceiling-light-pro/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "yeelight",
            "model": "YLXD50YL",
            "name": "Yeelight Halo Ceiling Light",
            "categories": ["light"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/ada65aa1_yeelight-light-29.svg",
            "link": "https://ru.yeelight.com/catalog/umnyy-potolochnyy-svetilnik-yeelight-halo/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "yeelight",
            "model": "YLDJ001",
            "name": "Yeelight Smart Curtain Motor & Stitching Track",
            "categories": ["curtain"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/a4f9e789_yeelight-curtain-1.svg",
            "link": "https://ru.yeelight.com/catalog/motor-dlya-karniza-yeelight-smart-curtain-motor-karniz-yeelight-stitching-track/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "yeelight",
            "model": "YLAI002",
            "name": "Yeelight Smart Dual Control Module",
            "categories": ["socket"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/b2869f13_yeelight-socker-1.jpeg",
            "link": "https://ru.yeelight.com/catalog/rele-yeelight-smart-dual-control-module/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "yeelight",
            "model": "YLDD005",
            "name": "Yeelight Lightstrip Pro",
            "categories": ["led-strip"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/1cbfbf6c_yeelight-led-strip-1.svg",
            "link": "https://ru.yeelight.com/catalog/umnaya-svetodiodnaya-lenta-yeelight-lightstrip-pro/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "yeelight",
            "model": "YLDD05YL",
            "name": "Yeelight LED Lightstrip Plus 1S",
            "categories": ["led-strip"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/983ef68d_yeelight-led-strip-2-1.svg",
            "link": "https://ru.yeelight.com/catalog/umnaya-svetodiodnaya-lenta-yeelight-lightstrip-plus-1s/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "yeelightpro",
            "model": "YLP013",
            "name": "Yeelight Pro S20 Умный настенный переключатель. Подключается через хаб Yeelight",
            "categories": ["relay"],
            "protocols": ["bluetooth"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/1dbae18c_yeelightpro-relay-1.svg",
            "link": null,
            "direct": [],
            "cloud": ["с2с-hub"]
        },
        {
            "manufacturer": "yeelightpro",
            "model": "YP-0052",
            "name": "Yeelight Pro E20 Умный настенный переключатель трёхканальный. Подключается через хаб Yeelight",
            "categories": ["relay"],
            "protocols": ["bluetooth"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/68459235_yeelightpro-relay-2.svg",
            "link": "https://yeelightpro.ru/catalog/knopochnaya-panel-yeelight-pro-e20-smart-wall-switch-6-klavish/",
            "direct": [],
            "cloud": ["с2с-hub"]
        },
        {
            "manufacturer": "yeelightpro",
            "model": "YP-0046",
            "name": "Yeelight Pro E20 Умная лампа E27. Подключается через хаб Yeelight",
            "categories": ["light"],
            "protocols": ["bluetooth"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/fa048893_yeelightpro-light-1.png",
            "link": "https://yeelightpro.ru/catalog/umnaya-svetodiodnaya-lampa-yeelight-pro-e20-smart-led-bulb/",
            "direct": [],
            "cloud": ["с2с-hub"]
        },
        {
            "manufacturer": "yeelightpro",
            "model": "YLP040",
            "name": "Yeelight Pro M20 Умная лампа E27. Подключается через хаб Yeelight",
            "categories": ["light"],
            "protocols": ["bluetooth"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/c9f5d7dd_yeelightpro-light-2.png",
            "link": "https://yeelightpro.ru/catalog/umnaya-svetodiodnaya-lampa-yeelight-pro-m20-smart-led-bulb-tunable-white/",
            "direct": [],
            "cloud": ["с2с-hub"]
        },
        {
            "manufacturer": "yeelightpro",
            "model": "YLP033",
            "name": "Yeelight Pro M20 Умный встраиваемый светильник. Подключается через хаб Yeelight",
            "categories": ["light"],
            "protocols": ["bluetooth"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/ff6c98c3_yeelightpro-light-3-1.svg",
            "link": null,
            "direct": [],
            "cloud": ["с2с-hub"]
        },
        {
            "manufacturer": "zanussi",
            "model": "",
            "name": "Кондиционеры Zanussi. Для управления по Wi-Fi должны быть оснащены модулем Daichi DW21-B/DW22-B",
            "categories": ["hvac-ac"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/zanussi-hvac-ac.svg",
            "link": "https://daichicloud.ru/split-lineup/search.php?search=Zanussi",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "zont",
            "model": "ML00005454",
            "name": "ZONT H-1V.02",
            "categories": ["hvac-boiler"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/77dbb785_zont-hvac-boiler-1.svg",
            "link": "https://zont-online.ru/product/zont-h-1v-02/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "zont",
            "model": "ML00004479",
            "name": "ZONT SMART 2.0",
            "categories": ["hvac-boiler"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/bd65841d_zont-hvac-boiler-2.svg",
            "link": "https://zont-online.ru/product/zont-smart-2-0/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "zont",
            "model": "ML00005890",
            "name": "ZONT H-1V NEW",
            "categories": ["hvac-boiler"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/70367c93_zont-hvac-boiler-3.svg",
            "link": "https://zont-online.ru/product/otopitelnyj-termostat-zont-h-1v-new/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "zont",
            "model": "ML00005886",
            "name": "ZONT SMART NEW",
            "categories": ["hvac-boiler"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/bd65841d_zont-hvac-boiler-2.svg",
            "link": "https://zont-online.ru/product/otopitelnyj-termostat-zont-smart-new/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "zont",
            "model": "ML00005559",
            "name": "ZONT H2000+ PRO",
            "categories": ["hvac-boiler"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/04d6d88d_zont-hvac-boiler-5.svg",
            "link": "https://zont-online.ru/product/zont-h2000-pro/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "zont",
            "model": "ML00005968",
            "name": "ZONT H1500+ PRO",
            "categories": ["hvac-boiler"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/4c7b76d9_zont-hvac-boiler-6.svg",
            "link": "https://zont-online.ru/product/zont-h1500-pro/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "zont",
            "model": "ML00005558",
            "name": "ZONT H1000+ PRO",
            "categories": ["hvac-boiler"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/a8f7f49c_zont-hvac-boiler-7.svg",
            "link": "https://zont-online.ru/product/zont-h1000-pro/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "zont",
            "model": "ML00005557",
            "name": "ZONT H700+ PRO",
            "categories": ["hvac-boiler"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/98ffc269_zont-hvac-boiler-8.svg",
            "link": "https://zont-online.ru/product/zont-h700-pro/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "zont",
            "model": "ML00004511",
            "name": "ZONT Climatic 1.1",
            "categories": ["hvac-boiler"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/242ad355_zont-hvac-boiler-9.svg",
            "link": "https://zont-online.ru/product/zont-climatic-1-1/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "zont",
            "model": "ML00004510",
            "name": "ZONT Climatic 1.2",
            "categories": ["hvac-boiler"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/242ad355_zont-hvac-boiler-9.svg",
            "link": "https://zont-online.ru/product/zont-climatic-1-2/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "zont",
            "model": "ML00004486",
            "name": "ZONT Climatic 1.3",
            "categories": ["hvac-boiler"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/242ad355_zont-hvac-boiler-9.svg",
            "link": "https://zont-online.ru/product/zont-climatic-1-3/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "majordom",
            "model": "",
            "name": "Платформа Мажордом. Управляет выключателями, реле, жалюзи, рулонными шторами, розетками, термоголовками и терморегуляторами для радиаторов",
            "categories": ["other", "relay", "window-blind", "socket", "hvac-radiator"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/41c95985_magordomo-logo.svg",
            "link": "https://majord.ru/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "mts",
            "model": "JY-GZ-03AQ",
            "name": "Датчик дыма МТС JY-GZ-03AQ. Подключается через Zigbee-колонку Sber или хаб Sber",
            "categories": ["sensor-smoke"],
            "protocols": ["zigbee"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/aeb1ee2b_mts-JY-GZ-03AQ.png",
            "link": "https://shop.mts.ru/product/datchik-dyma-mts-jy-gz-03aq-1-belyj",
            "direct": ["hub"],
            "cloud": []
        },
        {
            "manufacturer": "mts",
            "model": "MCCGQ11LM",
            "name": "Датчик открытия МТС MCCGQ11LM. Подключается через Zigbee-колонку Sber или хаб Sber",
            "categories": ["sensor-door"],
            "protocols": ["zigbee"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/4e334ad3_mts-MCCGQ11LM.png",
            "link": "https://shop.mts.ru/product/datchik-otkrytija-mts-mccgq11lm-1-belyj",
            "direct": ["hub"],
            "cloud": []
        },
        {
            "manufacturer": "mts",
            "model": "WSDCGQ11LM",
            "name": "Датчик температуры и влажности МТС WSDCGQ11LM. Подключается через Zigbee-колонку Sber или хаб Sber",
            "categories": ["sensor-temp"],
            "protocols": ["zigbee"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/8d849bfe_mts-WSDCGQ11LM.png",
            "link": "https://shop.mts.ru/product/datchik-temperatury-i-vlazhnosti-mts-wsdcgq11lm-1-belyj",
            "direct": ["hub"],
            "cloud": []
        },
        {
            "manufacturer": "mts",
            "model": "SP-EUC01",
            "name": "Умная розетка МТС SP-EUC01. Подключается через Zigbee-колонку Sber или хаб Sber",
            "categories": ["socket"],
            "protocols": ["zigbee"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/b44b3b77_mts-SP-EUC01.png",
            "link": "https://shop.mts.ru/product/umnaja-rozetka-mts-sp-euc01-1-belaja",
            "direct": ["hub"],
            "cloud": []
        },
        {
            "manufacturer": "centersvet",
            "model": "VOICE CONVERTER",
            "name": "ЦЕНТРСВЕТ VOICE CONVERTER",
            "categories": ["controller"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/centersvet-dali-converter.svg",
            "link": "https://www.centersvet.ru/catalog/components/bluetooth_dali_converter/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "centersvet",
            "model": "INF VOICE CONTROL BK",
            "name": "ЦЕНТРСВЕТ INFINITY VOICE CONTROL",
            "categories": ["controller"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/centersvet-inf-voice-control.svg",
            "link": "https://www.centersvet.ru/catalog/search/inf_voice_control/",
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "era",
            "model": "5050-30-RGB-IP65-Wifi-5m (12V)",
            "name": "ЭРА Комплект светодиодной ленты",
            "categories": ["led-strip"],
            "protocols": ["wifi"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/0a973805_era-led-strip-1.svg",
            "link": null,
            "direct": [],
            "cloud": ["с2с-app"]
        },
        {
            "manufacturer": "yandex",
            "model": "YNDX-00522",
            "name": "Датчик движения и освещения Яндекс. Подключается через Zigbee-колонку Sber или хаб Sber",
            "categories": ["sensor-pir"],
            "protocols": ["zigbee"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/yandex-00522.png",
            "link": "https://market.yandex.ru/product--yndx-00522/1819516413",
            "direct": ["hub"],
            "cloud": []
        },
        {
            "manufacturer": "yandex",
            "model": "YNDX-00520",
            "name": "Датчик открытия дверей и окон Яндекс. Подключается через Zigbee-колонку Sber или хаб Sber",
            "categories": ["sensor-door"],
            "protocols": ["zigbee"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/94501430_yndx-00520.png",
            "link": "https://market.yandex.ru/product--yndx-00520/1818999628",
            "direct": ["hub"],
            "cloud": []
        },
        {
            "manufacturer": "yandex",
            "model": "YNDX-00521",
            "name": "Датчик протечки Яндекс. Подключается через Zigbee-колонку Sber или хаб Sber",
            "categories": ["sensor-water-leak"],
            "protocols": ["zigbee"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/e084a7fb_yndx-00521.png",
            "link": "https://market.yandex.ru/product--yndx-00521/1821677519",
            "direct": ["hub"],
            "cloud": []
        },
        {
            "manufacturer": "yandex",
            "model": "YNDX-00523",
            "name": "Датчик температуры и влажности Яндекс. Подключается через Zigbee-колонку Sber или хаб Sber",
            "categories": ["sensor-temp"],
            "protocols": ["zigbee"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/2c1a9c1e_yndx-00523.png",
            "link": "https://market.yandex.ru/product--yndx-00523/1819006718",
            "direct": ["hub"],
            "cloud": []
        },
        {
            "manufacturer": "yandex",
            "model": "YNDX-00524",
            "name": "Беспроводная кнопка Яндекс. Подключается через Zigbee-колонку Sber или хаб Sber",
            "categories": ["scenario-button"],
            "protocols": ["zigbee"],
            "image": "https://cdn-app.sberdevices.ru/misc/0.0.0/assets/sd-help/49a4063e_yndx-00524.png",
            "link": "https://market.yandex.ru/product--yndx-00524/1819523517",
            "direct": ["hub"],
            "cloud": []
        }
    ]
}`

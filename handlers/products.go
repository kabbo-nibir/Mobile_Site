package handlers

import (
	"context"
	"encoding/json"
	"first/dbiface"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/go-playground/validator.v9"
)

var (
	v = validator.New()
)

//Mobile data
type Mobile struct {
	ID                     primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	BrandName              string             `json:"brand" bson:"brand"`
	ModelName              string             `json:"model_name" bson:"model_name"`
	NetworkBand            string             `json:"network" bson:"network"`
	ReleaseDate            time.Time          `json:"release_date" bson:"release_date"`
	BodyDimensions         []int              `json:"dimension" bson:"dimension"`
	BodyDimensionsDisplay  string             `json:"body_dimension_display" bson:"body_dimension_display"`
	BodyWeight             int                `json:"weight" bson:"weight"`
	BodyBuild              BodyBuild          `json:"body_build" bson:"body_build"`
	BodyBuildDisplay       string             `json:"body_build_display" bson:"body_build_display"`
	IP68Rating             bool               `json:"ip68_ration" bson:"ip68_ration"`
	SimSlot                string             `json:"sim" bson:"sim"`
	Display                DisplayType        `json:"display" bson:"display"`
	DisplayScreen          string             `json:"display_screen" bson:"display_screen"`
	Os                     Os                 `json:"os" bson:"os"`
	OsDisplay              string             `json:"os_display" bson:"os_display"`
	Processor              Processor          `json:"processor" bson:"processor"`
	ProcessorDisplay       string             `json:"processor_display" bson:"processor_display"`
	GPU                    string             `json:"gpu" bson:"gpu"`
	RAM                    RAM                `json:"ram" bson:"ram"`
	RAMDispay              string             `json:"ram_display" bson:"ram_display"`
	InternalStorage        InternalStorage    `json:"internal_storage" bson:"internal_storage"`
	InternalStorageDisplay string             `json:"internal_storage_display" bson:"internal_storage_display"`
	MemoryCardSlot         bool               `json:"memory_slot" bson:"memory_slot"`
	Camera                 []Camera           `json:"main_camera" bson:"main_camera"`
	CameraDisplay          string             `json:"camera_display" bson:"camera_display"`
	SelfiCamera            SelfiCameraType    `json:"selfi_camera" bson:"selfi_camera"`
	SelfiCameraDisplay     string             `json:"selfi_camera_display" bson:"selfi_camera_display"`
	Sound                  Sound              `json:"sound" bson:"sound"`
	SoundDisplay           string             `json:"sound_display" bson:"sound_display"`
	Wifi                   string             `json:"wifi" bson:"wifi"`
	Bluetooth              string             `json:"blutooth" bson:"blutooth"`
	NFC                    bool               `json:"nfc" bson:"nfc"`
	FMRadio                bool               `json:"fm_radio" bson:"fm_radio"`
	USB                    USB                `json:"usb" bson:"usb"`
	USBDisplay             string             `json:"usb_display" bson:"usb_display"`
	Sensor                 []string           `json:"sensor" bson:"sensor"`
	SensorDisplay          string             `json:"sensor_display" bson:"sensor_display"`
	Battery                BatteryType        `json:"battery" bson:"battery"`
	BatteryDisplay         string             `json:"battery_display" bson:"battery_display"`
	Colors                 []string           `json:"color" bson:"color"`
	ColorsDisplay          string             `json:"colors_display" bson:"colors_display"`
	UnitPrice              int                `json:"price" bson:"price"`
}

//BodyBuild from mobile struct
type BodyBuild struct {
	Front string `bson:"front,omitempty"`
	Back  string `bson:"back,omitempty"`
	Body  string `bson:"body"`
}

//DisplayType from mobile struct
type DisplayType struct {
	DisplayType        string `bson:"display_type"`
	Hdr                string `bson:"hdr,omitempty"`
	Ppi                int    `bson:"ppi,omitempty"`
	DisplayRefreshRate int    `bson:"display_refresh_rate"`
	DisplaySize        []int  `bson:"display_size"`
	DisplayResolution  []int  `bson:"display_resultion"`
}

// Os from mobile struct
type Os struct {
	Type       string `bson:"type"`
	Custom     string `bson:"custom,omitempty"`
	Upgradable string `bson:"upgradble"`
}

//Processor from mobile struct
type Processor struct {
	Chipset         string   `bson:"chipset"`
	ChipsetDiaMeter int      `bson:"chipset_dia_meter"`
	NumberOfCore    int      `bson:"numbers_of_core"`
	CoreInfo        []string `bson:"core_info"`
}

//RAM from mobile struct
type RAM struct {
	RAM  int    `bson:"ram"`
	Type string `bson:"type,omitempty"`
}

//InternalStorage from mobile struct
type InternalStorage struct {
	Capacity int     `bson:"capacity"`
	Ufs      float32 `bson:"ufs,omitempty"`
}

//Camera from mobile struct
type Camera struct {
	MegaPixel     int      `bson:"megapixel"`
	LensType      string   `bson:"lens_type"`
	FocaLength    int      `bson:"focal_length"`
	Aperture      string   `bson:"aperture"`
	SensorSize    string   `bson:"sensor_size"`
	Stabilization string   `bson:"stabilization"`
	Autofocus     []string `bson:"autofocus"`
}

//SelfiCameraType from mobile struct
type SelfiCameraType struct {
	MegaPixel     int      `bson:"megapixel"`
	LensType      string   `bson:"lens_type"`
	FocaLength    int      `bson:"focal_length"`
	Aperture      string   `bson:"aperture"`
	SensorSize    string   `bson:"sensor_size"`
	Stabilization string   `bson:"stabilization"`
	Autofocus     []string `bson:"autofocus"`
}

//Sound from mobile struct
type Sound struct {
	SteroSpeaker    bool `bson:"stero_speaker"`
	AudioJack       bool `bson:"3.5mm audio jack"`
	DolbyAtmosSound bool `bson:"dolby_atmos_sound,omitempty"`
	Bit             int  `bson:"bit,omitempty"`
	Khz             int  `bson:"khz,omitempty"`
}

//USB from mobile struct
type USB struct {
	USB float32 `bson:"usb"`
	OTG bool    `bson:"otg"`
}

//BatteryType from mobile struct
type BatteryType struct {
	BatteryType                 string  `bson:"batter_type"`
	BatterCapacity              int     `bson:"battery_capacity"`
	Removable                   bool    `bson:"removable"`
	ChargingType                string  `bson:"charging_type"`
	ChargingWatt                int     `bson:"bson:Charging_watt"`
	WirelessChargingWatt        int     `bson:"wireless_charging_watt,omitempty"`
	ReverseWirelessChargingWatt float32 `bson:"reverse_wireless_charging_watt,omitempty"`
}

//MobileHandler a mobile handler
type MobileHandler struct {
	Col dbiface.CollectionAPI
}

//insertMobiles for CreateMobiles function
func insertMobiles(ctx context.Context, mobiles []Mobile, collection dbiface.CollectionAPI) ([]interface{}, error) {
	var insertedIds []interface{}
	for _, mobile := range mobiles {
		mobile.ID = primitive.NewObjectID()
		insertID, err := collection.InsertOne(ctx, mobile)
		if err != nil {
			log.Fatalf("unable to insert : %v", err)
			return nil, err
		}
		insertedIds = append(insertedIds, insertID.InsertedID)
	}
	return insertedIds, nil
}

//CreateMobiles create mobiles on mongodb database
func (h *MobileHandler) CreateMobiles(c echo.Context) error {
	var mobiles []Mobile
	if err := c.Bind(&mobiles); err != nil {
		log.Printf("unable to bind : %v", err)
		return err
	}
	IDs, err := insertMobiles(context.Background(), mobiles, h.Col)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, IDs)
}

//findMobile for GetMobile function
func findMobile(ctx context.Context, id string, collection dbiface.CollectionAPI) (Mobile, error) {
	var mobile Mobile
	docID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return mobile, err
	}

	res := collection.FindOne(ctx, bson.M{"_id": docID})
	err = res.Decode(&mobile)
	if err != nil {
		return mobile, err
	}
	return mobile, nil
}

//GetMobile gets a single mobile
func (h *MobileHandler) GetMobile(c echo.Context) error {
	mobile, err := findMobile(context.Background(), c.Param("id"), h.Col)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, mobile)
}

func deleteMobile(ctx context.Context, id string, collection dbiface.CollectionAPI) (int64, error) {
	docID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Printf("Unable convert to ObjectID: %v", err)
		return 0, err
	}

	res, err := collection.DeleteOne(ctx, bson.M{"_id": docID})
	if err != nil {
		log.Printf("unable to delete mobile: %v", err)
		return 0, err
	}
	return res.DeletedCount, nil
}

//DeleteMobile to delete mobile from mongodb
func (h *MobileHandler) DeleteMobile(c echo.Context) error {
	delCount, err := deleteMobile(context.Background(), c.Param("id"), h.Col)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, delCount)
}

func modifyMobile(ctx context.Context, id string, reqBody io.ReadCloser, collection dbiface.CollectionAPI) (Mobile, error) {
	var mobile Mobile
	//find if the mobile exits, if error return 404
	docID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return mobile, err
	}
	filter := bson.M{"_id": docID}
	res := collection.FindOne(ctx, filter)

	if err := res.Decode(&mobile); err != nil {
		return mobile, err
	}
	//decode the req payload, if error return 500
	if err := json.NewDecoder(reqBody).Decode(&mobile); err != nil {
		return mobile, err
	}

	//validate the request, if err return 400
	if err := v.Struct(mobile); err != nil {
		return mobile, err
	}

	//update the product,if err return 500
	_, err = collection.UpdateOne(ctx, filter, bson.M{"$set": mobile})
	if err != nil {
		return mobile, err
	}
	return mobile, nil
}

//UpdateMobile form mongodb
func (h *MobileHandler) UpdateMobile(c echo.Context) error {
	mobile, err := modifyMobile(context.Background(), c.Param("id"), c.Request().Body, h.Col)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, mobile)
}

func findMobiles(ctx context.Context, collection dbiface.CollectionAPI) ([]Mobile, error) {
	var mobiles []Mobile
	opts := options.Find()
	opts.SetLimit(2)

	sortCursor, err := collection.Find(ctx, bson.M{}, opts)
	if err != nil {
		log.Printf("unable to find the mobiles : %v", err)
	}

	err = sortCursor.All(ctx, &mobiles)
	if err != nil {
		log.Printf("unable to read the cursor : %v", err)
	}
	return mobiles, nil
}

//GetMobiles gets a list of mobiles
func (h *MobileHandler) GetMobiles(c echo.Context) error {

	mobiles, err := findMobiles(context.Background(), h.Col)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, mobiles)
}

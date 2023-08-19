package wireless

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/project-flotta/flotta-device-worker/internal/common"
	"github.com/project-flotta/flotta-operator/models"
)

func GetConnectedWirelessDevices(db *sql.DB) ([]*models.WirelessDevice, error) {
	rows, err := db.Query("SELECT wireless_device_name, wireless_device_manufacturer, wireless_device_model, wireless_device_sw_version, wireless_device_identifier, wireless_device_protocol, wireless_device_connection,wireless_device_battery, wireless_device_availability, wireless_device_description, wireless_device_last_seen FROM wireless_device ORDER BY wireless_device_id DESC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []*models.WirelessDevice

	for rows.Next() {
		// fmt.Println("ITEMS")
		var deviceProperties []*models.DeviceProperty
		// Create a new instance of models.WirelessDevice
		item := &models.WirelessDevice{}
		err := rows.Scan(&item.WirelessDeviceName, &item.WirelessDeviceManufacturer, &item.WirelessDeviceModel, &item.WirelessDeviceSwVersion, &item.WirelessDeviceIdentifier, &item.WirelessDeviceProtocol, &item.WirelessDeviceConnection, &item.WirelessDeviceBattery, &item.WirelessDeviceAvailability, &item.WirelessDeviceDescription, &item.WirelessDeviceLastSeen)
		if err != nil {
			return nil, err
		}

		//get device properties
		// rowProperties, err := db.Query("SELECT wireless_device_identifier, property_identifier, property_service_uuid, property_name, property_access_mode, property_reading, property_state, property_unit, property_description,  property_last_seen FROM device_property WHERE wireless_device_identifier = '" + item.WirelessDeviceIdentifier + "' GROUP BY property_identifier ORDER BY device_property_id DESC")
		rowProperties, err := db.Query("SELECT dp.wireless_device_identifier, dp.property_identifier, dp.property_service_uuid, dp.property_name, dp.property_access_mode, dp.property_reading, dp.property_state, dp.property_unit, dp.property_description,  dp.property_last_seen FROM device_property dp JOIN (SELECT property_identifier,MAX(device_property_id) AS max_device_property_id FROM device_property WHERE wireless_device_identifier = '" + item.WirelessDeviceIdentifier + "' GROUP BY property_identifier) max_ids ON dp.property_identifier = max_ids.property_identifier AND dp.device_property_id = max_ids.max_device_property_id;")
		if err != nil {
			return nil, err
		}
		defer rowProperties.Close()
		for rowProperties.Next() {
			property := &models.DeviceProperty{}
			err := rowProperties.Scan(&property.WirelessDeviceIdentifier, &property.PropertyIdentifier, &property.PropertyServiceUUID, &property.PropertyName, &property.PropertyAccessMode, &property.PropertyReading, &property.PropertyState, &property.PropertyUnit, &property.PropertyDescription, &property.PropertyLastSeen)
			if err != nil {
				return nil, err
			}

			deviceProperties = append(deviceProperties, property)
		}

		item.DeviceProperties = deviceProperties
		items = append(items, item)

	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return items, nil
}

func ActionForDownStream(db *sql.DB, wirelessDeviceConfiguration models.WirelessDevice) error {

	client, err := common.MQTT_Connect()
	if err != nil {
		fmt.Println("ERROR CONNECT MQTT: ", err.Error())
		return err
	}

	// topic, err := GetEndNodeDeviceTopic(db, wirelessDeviceConfiguration.WirelessDeviceIdentifier)
	// if err != nil {
	// 	return err
	// }

	var devicePropertiesData []map[string]interface{}
	for _, property := range wirelessDeviceConfiguration.DeviceProperties {
		propertyData := map[string]interface{}{
			"property_access_mode":       property.PropertyAccessMode,
			"property_description":       property.PropertyDescription,
			"property_identifier":        property.PropertyIdentifier,
			"wireless_device_identifier": property.WirelessDeviceIdentifier,
			"property_last_seen":         property.PropertyLastSeen,
			"property_name":              property.PropertyName,
			"property_reading":           property.PropertyReading,
			"property_service_uuid":      property.PropertyServiceUUID,
			"property_state":             property.PropertyState,
			"property_unit":              property.PropertyUnit,
		}
		devicePropertiesData = append(devicePropertiesData, propertyData)
	}

	data := map[string]interface{}{
		"wireless_device_identifier": wirelessDeviceConfiguration.WirelessDeviceIdentifier,
		"wireless_device_name":       wirelessDeviceConfiguration.WirelessDeviceName,
		"device_properties":          devicePropertiesData,
	}

	// fmt.Println("success")
	// fmt.Println("success")
	// fmt.Println("success")
	// fmt.Println("Data in the 'data' map:")
	// fmt.Println("Wireless Device Identifier:", data["wireless_device_identifier"])
	// fmt.Println("Wireless Device Name:", data["wireless_device_name"])

	// deviceProperties := data["device_properties"].([]map[string]interface{})
	// if deviceProperties != nil {
	// 	fmt.Println("Device Properties:")
	// 	for idx, propertyData := range deviceProperties {
	// 		fmt.Printf("Property %d:\n", idx+1)
	// 		for key, value := range propertyData {
	// 			fmt.Printf("  %s: %v\n", key, value)
	// 		}
	// 	}
	// } else {
	// 	fmt.Println("No device properties found in the data.")
	// }

	if wirelessDeviceConfiguration.WirelessDeviceConnection == strings.ToLower("wi-fi") || wirelessDeviceConfiguration.WirelessDeviceConnection == strings.ToLower("WIFI") {
		err = PublishMQTT(client, "cloud/plugin/downstream/wifi", data)
		if err != nil {
			return err
		}
		return nil
	} else {
		err = PublishMQTT(client, "cloud/device/downstream", data)
		if err != nil {
			return err
		}
		return nil
	}

}

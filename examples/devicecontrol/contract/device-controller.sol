pragma solidity ^0.8.6;

contract DeviceController {

    struct Device {
        bool ready;
        uint32 metric;
        string deviceType;
    }

    mapping (string => Device) public devices;

    function register(string calldata _deviceId, string calldata _type) public {
        require(bytes(devices[_deviceId].deviceType).length == 0);
        devices[_deviceId] = Device(true, 0, _type);
    }

    function updateReadyState(string calldata _deviceId, bool _ready) public {
        require(bytes(devices[_deviceId].deviceType).length > 0);
        devices[_deviceId].ready = _ready;
    }

    function updateMetric(string calldata _deviceId, uint32 _delta) public returns (uint32 metric) {
        require(bytes(devices[_deviceId].deviceType).length > 0);
        devices[_deviceId].metric = devices[_deviceId].metric + _delta;
        return devices[_deviceId].metric;
    }

    function getReadyState(string calldata _deviceId) public view returns (bool readyState) {
        require(bytes(devices[_deviceId].deviceType).length > 0);
        return devices[_deviceId].ready;
    }
}

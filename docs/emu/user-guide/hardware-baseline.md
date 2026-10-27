# Hardware Baselines

## Get a Hardware Baseline
```bash
emu get hardware-baseline \
  -pageIndex 0 \
  -pageSize 2000
```

## Upload a Hardware Baseline
```bash
emu upload hardware-baseline \
  -approvalStatus="In Progress" \
  -assetIpAddress="Test Asset IP Address" \
  -assetName="Test Asset Name" \
  -componentType="Firewall" \
  -criticalAsset=true \
  -location="Test location (P/C/S & Building) 11" \
  -manufacturer="Test Manufacturer 11" \
  -memorySizeType="Test Memory Size / Type 11" \
  -modelNumber="Test Model Number 11" \
  -nickname="Test Nickname 11" \
  -osIosFwVersion="Test OS/iOS/FW Version 11" \
  -publicFacing=true \
  -publicFacingFqdn="Test Public Facing FQDN 11" \
  -publicFacingIpAddress="Test Public Facing IP Address 11" \
  -publicFacingUrls="Test Public Facing URL(s) 11" \
  -serialNumber="Test Serial Number 11" \
  -virtualAsset=false
```

## Update a Hardware Baseline
```bash
emu update hardware-baseline \
  -hardwareId="61dc2390-13c5-49ff-a102-52ae4f540a6e" \
  -componentType="IDS/IPS" \
  -assetName="Test Asset Modified with EMU Update Function" \
  -nickname="Test Nickname 11" \
  -assetIpAddress="Test Asset IP Address" \
  -publicFacing=true \
  -publicFacingFqdn="Test Public Facing FQDN 11" \
  -publicFacingIpAddress="Test Public Facing IP Address 11" \
  -publicFacingUrls="Test Public Facing URL(s) 11" \
  -virtualAsset=false \
  -manufacturer="Test Manufacturer 11" \
  -modelNumber="Test Model Number 11" \
  -serialNumber="Test Serial Number 11" \
  -osIosFwVersion="Test OS/iOS/FW Version 11" \
  -memorySizeType="Test Memory Size / Type 11" \
  -location="Test location223 (P/C/S & Building) 11" \
  -approvalStatus="In Progress" \
  -criticalAsset=true
```

## Delete a Hardware Baseline
```bash
emu delete hardware-baseline \
  -hardwareId 61dc2390-13c5-49ff-a102-52ae4f540a6e
```

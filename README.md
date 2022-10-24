# Smart Grid Pi

This repository contains the code that runs on a **Raspberry Pi** to control the **Smart Grid Ready** input of a heat pump, based on the current electricity production of a **solar power system**.

What is **Smart Grid Ready**? [SG Ready](https://www.ecodan.de/en/knowledge/sg-ready/) heat pumps can be controlled to increase the temperature set-point when the electricity is cheaper or when there is a surplus of power from a PV system.

The program collects every `REFRESH_TIME_SECONDS` seconds the measurements of domestic power generated and consumed from the SolarEdge monitoring API. When the power surplus is greater than `MIN_POWER_SURPLUS_WATT` then the heat pump is set in "Recommended ON" state.



### Environment variables
```bash
REFRESH_TIME_SECONDS = 15
MIN_POWER_SURPLUS_WATT = 2000
SOLAREDGE_API_BASE_URL = https://monitoringapi.solaredge.com
SOLAREDGE_API_KEY = your-solaredge-api-key
SOLAREDGE_SITE_ID = your-solaredge-site-id
```

You can find your SolarEdge Api Key and site ID from the [SolarEdge web dashboard](https://monitoring.solaredge.com/solaredge-web/p/login) settings.


### My components

- PV Inverter: SolarEdge SE5000H
- Heat pump: Mitsubishi Ecodan SUZ-SWM80VA (external) - Hydrotank ERST20D (internal)
- Board: Raspberry Pi 3 Model B
- Relay: AZDelivery KY-019


# Smart Grid Pi

This repository contains the code that runs on a **Raspberry Pi** to control the **Smart Grid Ready** input of a heat pump based on the current electricity production of a **solar power system**.

What is **Smart Grid Ready**? [SG Ready](http://heatpumpingtechnologies.org/archive/hpc2017/wp-content/uploads/2017/08/O.2.3.2-Flexibility-of-heat-pump-pools-The-use-of-SG-Ready-from-an-aggregators-perspective.pdf) heat pumps can be controlled to increase the temperature set-point when the electricity is cheaper or when there is a surplus of power from a PV system.

The program collects every 5 minutes the measurements of domestic power generated and consumed from the SolarEdge monitoring API. When the power surplus is greater than 2 kW then the heat pump is set in "Recommended ON" state.



### Environment variables
```bash
REFRESH_TIME_SECONDS = 15
SOLAREDGE_API_BASE_URL = https://monitoringapi.solaredge.com
SOLAREDGE_SITE_ID = 1234
SOLAREDGE_API_KEY = XXXX
```

You can find your SolarEdge Api Key and site ID from the [SolarEdge web dashboard](https://monitoring.solaredge.com/solaredge-web/p/login) settings.


### My appliances

- PV Inverter: SolarEdge SE5000H
- Heat pump: Mitsubishi Ecodan PUZ-WM85VAA
- Board: Raspberry Pi 3 Model B
- Relay: AZDelivery KY-019


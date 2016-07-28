# BLE sensor scanner


[BLE Sensors]  ----> [pdex-cli app] ---> Pd Exchange


Step 1:
	configure the pdex-cli

Step 2:
	create device group
		```
			pdex cr dg
		```

	creare application
		```
			pdex cr apps --app-name-suffix APP-NAME-SUFFIX
		```

Step 3:
	update the PD Exchange application and device-group information in the app.js file

	```
	const devicegroup = "01.72da6d"
	const app_id      = "4817e8ee00814e93af7a59c80b8625f9"
	```

Step 4:
	Install dependencies
	```
		npm install
	```

Step 5:
	execute the application
	```
		node .
	```


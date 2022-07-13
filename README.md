# WiFi-export-import
--- EN ---

Program for exporting and importing WiFi networks profiles on Windows systems

--- PL ---

Program do importowania i eksportowania profili sieci WiFi w systemach Windows



<img src="/screenshots/WiFi-export-import.png" alt="WiFi-export-import" height="100%" width="100%">


## Usage
### Export all WiFi profiles.
- Download [WiFi-export-import-v2.0.0-x86.exe](https://github.com/jtaczanowski/WiFi-export-import/releases/download/v2.0.0/WiFi-export-import-v2.0.0-x86.exe)
- Run ```WiFi-export-import-v2.0.0-x86.exe``` as ⚠️administrator. (Administrator privileges are needed to export all WiFi password as clear text, without administrator privileges only current user WiFi profiles will export with password as clear text - which is needed to import them on other computers)
- click "Export WiFi profiles" button and choose directory to which you want to save exported WiFi profiles.

### Import WiFi profiles from directory.
- Run ```WWiFi-export-import-v2.0.0-x86.exe```.
- click "Import WiFi profiles for current user" button and choose directory from which you want to import WiFi profiles for current user.
or
- click "Import WiFi profiles for all users" button and choose directory from which you want to import WiFi profiles for all users. (⚠️administrator privileges are needed)
## About WiFi-export-import
--- EN ---

WiFi-export-import is a simple tool for exporting and importing WiFi networks profiles from Windows systems. Underthehood, for exporting and importing WiFi network profiles program is using ```netsh``` commands.

--- PL ---

WiFi-export-import to prosty program pozwalający na eksportowanie i importowanie profili sieci w systemach Windows. Program do eksportowania i importowania profili wykorzystuje polecenia ```netsh```.
### System requirements
 - ✅  tested on Windows versions from 7 to 11

### Other
 - program was developed in Golang, GUI was builded with [github.com/rodrigocfd/windigo](https://github.com/rodrigocfd/windigo) library.
 - program uses netsh commands like ```netsh wlan export profile key=clear folder=C:\wifi``` ```wlan add profile filename=C:\wifi\WiFi-network.xml```.

